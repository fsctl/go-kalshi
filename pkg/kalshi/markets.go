package kalshi

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fsctl/go-kalshi/pkg/kalshi/swagger"
)

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

func (kc *KalshiClient) GetMarketTicker(ctx context.Context, marketId string) string {
	userGetMarketsResponse, _, err := kc.swaggerApiClient.MarketApi.GetMarkets(ctx)
	if err != nil {
		log.Fatalf("Error from apiClient.MarketApi.GetMarkets: '%v'", err)
	}

	for _, market := range userGetMarketsResponse.Markets {
		if market.Id == marketId {
			return market.TickerName
		}
	}
	log.Fatalf("Error: could not find market '%v'\n", marketId)
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
