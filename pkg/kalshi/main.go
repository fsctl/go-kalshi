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
			fmt.Printf("  %v\n", market.Title)
		}
	}
	fmt.Printf("\n----------------------------------------------------\n\n")
}

func (kc *KalshiClient) PrintMarket(ctx context.Context, id string) {
	_ = ctx
	_ = id
	fmt.Printf("TODO TODO TODO...\n\n")
}
