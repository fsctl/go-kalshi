package kalshi

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

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

func (kc *KalshiClient) PrintMarketsList(ctx context.Context) {
	userGetMarketsResponse, _, err := kc.swaggerApiClient.MarketApi.GetMarkets(ctx)
	if err != nil {
		log.Fatalf("Error from apiClient.MarketApi.GetMarkets:  '%v'", err)
		return
	}

	ranges := make(map[string][]swagger.Market)
	for _, mkt := range userGetMarketsResponse.Markets {
		var t time.Time = mkt.CloseDate
		if !t.Before(time.Now()) {
			ranges[mkt.RangedGroupTicker] = append(ranges[mkt.RangedGroupTicker], mkt)
		}
	}
	for k, v := range ranges {
		for _, market := range v {
			fmt.Printf("[%v] %v (%v)\n", k, market.Title, market.TickerName)
		}
		fmt.Printf("\n")
	}
}

func (kc *KalshiClient) GetMarketId(ctx context.Context, ticker string) string {
	userGetMarketsResponse, _, err := kc.swaggerApiClient.MarketApi.GetMarkets(ctx)
	if err != nil {
		log.Fatalf("Error from apiClient.MarketApi.GetMarkets: '%v'", err)
		return ""
	}

	for _, market := range userGetMarketsResponse.Markets {
		if market.TickerName == ticker {
			return market.Id
		}
	}
	log.Fatalf("Error: could not find market '%v'\n", ticker)
	return ""
}

func (kc *KalshiClient) PrintMarket(ctx context.Context, id string) {
	userGetMarketResponse, _, err := kc.swaggerApiClient.MarketApi.GetMarket(ctx, id)
	if err != nil {
		log.Fatalf("Error in PrintMarket: '%v'\n", err)
		return
	}
	fmt.Printf("(%v) %v\n\n", userGetMarketResponse.Market.TickerName, userGetMarketResponse.Market.Title)

	getMarketOrderBookResponse, _, err := kc.swaggerApiClient.MarketApi.GetMarketOrderBook(ctx, id)
	if err != nil {
		log.Fatalf("Error in PrintMarket: '%v'\n", err)
		return
	}
	yesOrderBook := getMarketOrderBookResponse.OrderBook.Yes
	noOrderBook := getMarketOrderBookResponse.OrderBook.No
	mostExpensiveYesBid := float64(0.00)
	mostExpensiveNoBid := float64(0.00)
	fmt.Printf("Yes order book (bids)\n")
	for _, row := range yesOrderBook {
		price := fmt.Sprintf("%.2f", float64(row[0])/100)
		if float64(row[0])/100 > mostExpensiveYesBid {
			mostExpensiveYesBid = float64(row[0]) / 100
		}
		qty := row[1]
		fmt.Printf("  %v (%v contracts)\n", price, qty)
	}
	fmt.Printf("No order book (bids)\n")
	for _, row := range noOrderBook {
		price := fmt.Sprintf("%.2f", float64(row[0])/100)
		if float64(row[0])/100 > mostExpensiveNoBid {
			mostExpensiveNoBid = float64(row[0]) / 100
		}
		qty := row[1]
		fmt.Printf("  %v (%v contracts)\n", price, qty)
	}

	fmt.Printf("\n")
	if mostExpensiveYesBid > 0.00 && mostExpensiveNoBid > 0.00 {
		fmt.Printf("If you predict Yes:\n")
		fmt.Printf("  - You can open a Yes position for (1-%.2f) = %.2f and close it (buy No) for (1-%.2f) = %.2f\n", mostExpensiveNoBid, 1-mostExpensiveNoBid, mostExpensiveYesBid, 1-mostExpensiveYesBid)
		fmt.Printf("  - Open:  kalshi-tool order --action open-yes  --ticker=%v --contracts 1 --limit %.2f\n", userGetMarketResponse.Market.TickerName, 1-mostExpensiveNoBid)
		fmt.Printf("  - Close: kalshi-tool order --action close-yes --ticker=%v --contracts 1 --limit %.2f\n", userGetMarketResponse.Market.TickerName, 1-mostExpensiveYesBid)
		fmt.Printf("If you predict No:\n")
		fmt.Printf("  - You can open a No position for (1-%.2f) = %.2f and close it (buy Yes) for (1-%.2f) = %.2f\n", mostExpensiveYesBid, 1-mostExpensiveYesBid, mostExpensiveNoBid, 1-mostExpensiveNoBid)
		fmt.Printf("  - Open:  kalshi-tool order --action open-no  --ticker=%v --contracts 1 --limit %.2f\n", userGetMarketResponse.Market.TickerName, 1-mostExpensiveYesBid)
		fmt.Printf("  - Close: kalshi-tool order --action close-no --ticker=%v --contracts 1 --limit %.2f\n", userGetMarketResponse.Market.TickerName, 1-mostExpensiveNoBid)
	}
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
func (kc *KalshiClient) OrderOpenPosition(ctx context.Context, marketId string, side MarketSide, contracts int, limit float64) (string, error) {
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
		return "", err
	}
	reqBody := strings.NewReader(string(createReqJson))

	// construct url
	url := fmt.Sprintf("%s/users/%s/orders", BaseURLV1, kc.authToken.UserId)

	// perform request
	res, data := kc.doAuthenticatedRequest("POST", url, reqBody)
	if res.StatusCode != 201 {
		return "", fmt.Errorf("error: orderopenposition res.statuscode is not 201 (it's %d)", res.StatusCode)
	}

	// deserialize json
	var userOrderCreateResponse swagger.UserOrderCreateResponse
	if err := json.Unmarshal(data, &userOrderCreateResponse); err != nil {
		log.Fatalf("Error: failed to unmarshal: %v", err)
	}
	//fmt.Printf("\nSuccess: order id: %v (status: %s)\n", userOrderCreateResponse.Order.OrderId, userOrderCreateResponse.Order.Status)

	return userOrderCreateResponse.Order.OrderId, nil
}

// check that user has enough contracts on side, then open a new position on opposite side
func (kc *KalshiClient) OrderClosePosition(ctx context.Context, marketId string, side MarketSide, contracts int, limit float64) (string, error) {
	return "", fmt.Errorf("error: subcommand not implemented")
}

// TODO:
// - cancel resting order by ID --> return whether it was canceled successfully (fully)
// - get user's portfolio (needed to see if OrderClosePosition can be called with likely success)
//		- how can you tell from UserGetMarketPositions whether contracts are on Yes or No side?
// - OrderClosePosition
