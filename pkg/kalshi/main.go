package kalshi

import (
	"context"
	"fmt"
	"log"
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
		fmt.Printf("%v:\n", k)
		for _, market := range v {
			fmt.Printf("  %v (%v)\n", market.Title, market.TickerName)
		}
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

	fmt.Printf("Yes\n")
	fmt.Printf("  Bid:  %.2f\n", float64(userGetMarketResponse.Market.YesBid)/100)
	fmt.Printf("  Ask:  %.2f\n\n", float64(userGetMarketResponse.Market.YesAsk)/100)

	getMarketOrderBookResponse, _, err := kc.swaggerApiClient.MarketApi.GetMarketOrderBook(ctx, id)
	if err != nil {
		log.Fatalf("Error in PrintMarket: '%v'\n", err)
		return
	}
	yesOrderBook := getMarketOrderBookResponse.OrderBook.Yes
	noOrderBook := getMarketOrderBookResponse.OrderBook.No
	fmt.Printf("Yes order book (bids)\n")
	for _, row := range yesOrderBook {
		price := fmt.Sprintf("%.2f", float64(row[0])/100)
		qty := row[1]
		fmt.Printf("  %v (%v contracts)\n", price, qty)
	}
	fmt.Printf("No order book (bids)\n")
	for _, row := range noOrderBook {
		price := fmt.Sprintf("%.2f", float64(row[0])/100)
		qty := row[1]
		fmt.Printf("  %v (%v contracts)\n", price, qty)
	}
}

/*

type bidAskQty struct {
	bidQty int64
	askQty int64
}

func (kc *KalshiClient) PrintMarket2(ctx context.Context, id string) {
	userGetMarketResponse, _, err := kc.swaggerApiClient.MarketApi.GetMarket(ctx, id)
	if err != nil {
		log.Fatalf("Error in PrintMarket: '%v'\n", err)
		return
	}
	fmt.Printf("(%v) %v\n\n", userGetMarketResponse.Market.TickerName, userGetMarketResponse.Market.Title)

	fmt.Printf("Yes\n")
	fmt.Printf("  Bid:  %.2f\n", float64(userGetMarketResponse.Market.YesBid)/100)
	fmt.Printf("  Ask:  %.2f\n\n", float64(userGetMarketResponse.Market.YesAsk)/100)

	getMarketOrderBookResponse, _, err := kc.swaggerApiClient.MarketApi.GetMarketOrderBook(ctx, id)
	if err != nil {
		log.Fatalf("Error in PrintMarket: '%v'\n", err)
		return
	}

	yesOrderBookRows := make(map[string]bidAskQty)
	populateAllPrices(&yesOrderBookRows)
	yesOrderBookResponseBids := getMarketOrderBookResponse.OrderBook.Yes
	transformApiOrderBook(yesOrderBookResponseBids, &yesOrderBookRows)
	fmt.Printf("-- YES ORDER BOOK --\n")
	fmt.Printf("Price\tBid Qty\tAsk Qty\n")
	keys := make([]string, 0, len(yesOrderBookRows))
	for k := range yesOrderBookRows {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := yesOrderBookRows[k]
		fmt.Printf("%v\t%v\t%v\n", k, v.bidQty, v.askQty)
	}
	fmt.Printf("\n")

	noOrderBookRows := make(map[string]bidAskQty)
	populateAllPrices(&noOrderBookRows)
	noOrderBookResponseBids := getMarketOrderBookResponse.OrderBook.No
	transformApiOrderBook(noOrderBookResponseBids, &noOrderBookRows)
	fmt.Printf("-- NO ORDER BOOK --\n")
	fmt.Printf("Price\tBid Qty\tAsk Qty\n")
	keys = make([]string, 0, len(noOrderBookRows))
	for k := range noOrderBookRows {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := noOrderBookRows[k]
		fmt.Printf("%v\t%v\t%v\n", k, v.bidQty, v.askQty)
	}
}

func populateAllPrices(orderBookRows *map[string]bidAskQty) {
	for i := 1; i <= 99; i++ {
		iAsFloat := float64(i) / 100
		iAsStr := fmt.Sprintf("%.2f", iAsFloat)
		(*orderBookRows)[iAsStr] = bidAskQty{
			bidQty: int64(0),
			askQty: int64(0),
		}
	}
}

func transformApiOrderBook(apiOrderBookBids [][]int32, orderBookRows *map[string]bidAskQty) {
	for _, row := range apiOrderBookBids {
		priceBid := float64(row[0]) / 100
		priceBidAsStr := fmt.Sprintf("%.2f", priceBid)
		bidQty := row[1]
		if val, ok := (*orderBookRows)[priceBidAsStr]; ok {
			(*orderBookRows)[priceBidAsStr] = bidAskQty{
				bidQty: int64(bidQty),
				askQty: val.askQty,
			}
		}

		priceAsk := 1.0 - priceBid
		priceAskAsStr := fmt.Sprintf("%.2f", priceAsk)
		askQty := row[1]
		if val, ok := (*orderBookRows)[priceAskAsStr]; ok {
			(*orderBookRows)[priceAskAsStr] = bidAskQty{
				bidQty: val.bidQty,
				askQty: int64(askQty),
			}
		}
	}
}
*/
