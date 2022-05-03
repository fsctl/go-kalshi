package kalshi

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/fsctl/go-kalshi/pkg/kalshi/swagger"
)

const (
	BaseURLV1 = "https://trading-api.kalshi.com/v1"
)

// KalshiClient is a higher level API client than the Swagger generated client.  It
// does complex actions that involve multiple API calls.
type KalshiClient struct {
	authToken        AuthToken
	swaggerApiClient *swagger.APIClient
}

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

// open a new position on side
func (kc *KalshiClient) OrderOpenPosition(ctx context.Context, marketId string, side MarketSide, contracts int, limit float64) (string, bool, error) {
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

// check that user has enough contracts on side, then open a new position on opposite side
func (kc *KalshiClient) OrderClosePosition(ctx context.Context, marketId string, side MarketSide, contracts int, limit float64) (string, error) {
	return "", fmt.Errorf("error: subcommand not implemented")
}

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
// - get user's portfolio (needed to see if OrderClosePosition can be called with likely success)
//		- how can you tell from UserGetMarketPositions whether contracts are on Yes or No side?
// - OrderClosePosition
