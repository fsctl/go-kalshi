package kalshi

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/fsctl/go-kalshi/swagger"
)

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

// RestingOrder represents a single resting order for the logged in user.
type RestingOrder struct {
	Id           string
	MarketTicker string
	Side         MarketSide
	Contracts    int64
	Price        float64
}

// GetRestingOrders returns an array of all resting orders for the logged in user.
func (kc *KalshiClient) GetRestingOrders(ctx context.Context) ([]RestingOrder, error) {
	restingOrders := make([]RestingOrder, 0)

	// construct url
	url := fmt.Sprintf("%s/users/%s/orders?status=resting", BaseURLV1, kc.authToken.UserId)

	// perform request
	reqBody := strings.NewReader("")
	res, data := kc.doAuthenticatedRequest("GET", url, reqBody)
	if res.StatusCode != 200 {
		return restingOrders, fmt.Errorf("error: orderopenposition res.statuscode is not 200 (it's %d)", res.StatusCode)
	}

	// deserialize json
	var userOrdersGetResponse swagger.UserOrdersGetResponse
	if err := json.Unmarshal(data, &userOrdersGetResponse); err != nil {
		log.Fatalf("Error: failed to unmarshal: %v", err)
	}

	// build return array
	for _, apiOrder := range *userOrdersGetResponse.Orders {
		id := apiOrder.OrderId
		marketId := apiOrder.MarketId
		marketTicker := kc.GetMarketTicker(ctx, marketId)
		var side MarketSide
		if apiOrder.IsYes {
			side = Yes
		} else {
			side = No
		}
		contracts := apiOrder.RemainingCount
		price := float64(apiOrder.Price) / 100
		restingOrder := RestingOrder{
			Id:           id,
			MarketTicker: marketTicker,
			Side:         side,
			Contracts:    int64(contracts),
			Price:        price,
		}
		restingOrders = append(restingOrders, restingOrder)
	}

	return restingOrders, nil
}
