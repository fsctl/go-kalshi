package kalshi

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/fsctl/go-kalshi/swagger"
)

// This constant represents the Kalshi API endpoint URL to use.
const (
	BaseURLV1 = "https://trading-api.kalshi.com/v1"
)

// KalshiClient is a higher level API client than the Swagger generated client.
// It exposes methods to do complex actions that involve multiple underlying
// API calls.
type KalshiClient struct {
	authToken        authToken
	swaggerApiClient *swagger.APIClient
}

// NewKalshiClient() constructs a new KalshiClient struct that is initialized
// and ready for use, using the specified Kalshi credentials.
func NewKalshiClient(ctx context.Context, kalshiUsername string, kalshiPassword string) (*KalshiClient, error) {
	authToken, err := getAuthToken(ctx, kalshiUsername, kalshiPassword)
	if err != nil {
		log.Fatalf("ERROR:  getAuthToken failed:  '%v'\n", err)
		return nil, err
	}

	cfg := swagger.NewConfiguration()
	cfg.AddDefaultHeader("Authorization", fmt.Sprintf("%s %s", authToken.UserId, authToken.Token))

	swaggerApiClient := swagger.NewAPIClient(cfg)

	return &KalshiClient{
		authToken:        *authToken,
		swaggerApiClient: swaggerApiClient,
	}, nil
}

func (kc *KalshiClient) doAuthenticatedRequest(method string, url string, reqBody *strings.Reader) (*http.Response, []byte) {
	req, _ := http.NewRequest(method, url, reqBody)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("%s %s", kc.authToken.UserId, kc.authToken.Token))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error: http.DefaultClient.Do failed: '%v'\n", err)
	}

	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	return res, data
}

// This function opens a new position on `side` in the specified market for
// 'contracts' number of contracts with a limit price of limit (must be between
// 0.01 and 0.99 inclusive).
//
// Returns the order id (for cancelation if the order becomes a resting order),
// a boolean indicating whether the order executed or has become a resting order,
// and an error value that is non-nil if the function failed.
func (kc *KalshiClient) OrderOpenPosition(ctx context.Context, marketId string, side MarketSide, contracts int, limit float64) (orderId string, isResting bool, err error) {
	// serialize arguments to json
	userOrderCreateRequest := swagger.UserOrderCreateRequest{
		Count:              int32(contracts),
		ExpirationUnixTs:   0,
		MarketId:           marketId,
		MaxCostCents:       0,
		Price:              int64(limit * 100),
		SellPositionCapped: false,
		Side:               side.String(),
	}
	createReqJson, err := json.Marshal(userOrderCreateRequest)
	if err != nil {
		log.Fatalf("Error: json.Marshal(): %v\n", err)
	}
	reqBody := strings.NewReader(string(createReqJson))

	// construct url
	url := fmt.Sprintf("%s/users/%s/orders", BaseURLV1, kc.authToken.UserId)

	// perform request
	res, data := kc.doAuthenticatedRequest("POST", url, reqBody)
	if res.StatusCode != 201 {
		return "", false, fmt.Errorf("error: orderopenposition res.statuscode is not 201 (it's %d)", res.StatusCode)
	}

	// deserialize json
	var userOrderCreateResponse swagger.UserOrderCreateResponse
	if err := json.Unmarshal(data, &userOrderCreateResponse); err != nil {
		log.Fatalf("Error: failed to unmarshal: %v", err)
	}

	return userOrderCreateResponse.Order.OrderId, (userOrderCreateResponse.Order.Status == "resting"), nil
}

// This function checks that user has enough contracts on `side`, and if so
// opens a new position on opposite side to close out the user's current position.
//
// Parameters have the same meaning as in OrderOpenPosition. Note that `limit` is
// the maximum you will pay to open the contract(s) on the other side.
//
// Return values have the same meaning as in OrderOpenPosition.
func (kc *KalshiClient) OrderClosePosition(ctx context.Context, marketId string, side MarketSide, contracts int, limit float64) (orderId string, isResting bool, err error) {
	// Return with error if we have no contracts on 'side' to close
	ticker := kc.GetMarketTicker(ctx, marketId)
	portfolio, err := kc.GetPortfolio(ctx)
	if err != nil {
		log.Fatalf("Error in GetPortfolio: %v\n", err)
	}

	var contractsOnSide int64 = 0
	for _, item := range portfolio.Items {
		if item.Ticker == ticker {
			contractsOnSide = item.Contracts
		}
	}

	if contractsOnSide < int64(contracts) {
		return "", false, fmt.Errorf("error: ordercloseposition: cannot close %s position in %v because you have only %d %s contracts (need at least %d)", side.String(), ticker, contractsOnSide, side.String(), contracts)
	}

	// Buy 'contracts' contracts on opposite of 'side'
	openOnSide := side.Opposite()
	orderId, isResting, err = kc.OrderOpenPosition(ctx, marketId, openOnSide, contracts, limit)

	return orderId, isResting, err
}

// This function cancels a resting order given its order id.
func (kc *KalshiClient) CancelRestingOrder(ctx context.Context, orderId string) error {
	// construct url
	url := fmt.Sprintf("%s/users/%s/orders/%s", BaseURLV1, kc.authToken.UserId, orderId)

	// perform request
	reqBody := strings.NewReader("")
	res, _ := kc.doAuthenticatedRequest("DELETE", url, reqBody)
	if res.StatusCode != 200 {
		return fmt.Errorf("error: cancelrestingorder response not 200 (it's %d)", res.StatusCode)
	}

	return nil
}

// TODO:
// - get list of user's resting orders with id's
