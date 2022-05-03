package kalshi

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/fsctl/go-kalshi/pkg/kalshi/swagger"
)

type PortfolioItem struct {
	Ticker    string
	Contracts int64
	Side      MarketSide
}

type Portfolio struct {
	Items []PortfolioItem
}

func (p *Portfolio) Print() {
	fmt.Printf("Portfolio:\n")
	for _, item := range p.Items {
		fmt.Printf("  %s (%s): %d contracts\n", item.Ticker, item.Side.String(), item.Contracts)
	}
}

func (kc *KalshiClient) GetPortfolio(ctx context.Context) (*Portfolio, error) {
	p := Portfolio{
		Items: make([]PortfolioItem, 0),
	}

	// construct url
	url := fmt.Sprintf("%s/users/%s/positions", BaseURLV1, kc.authToken.UserId)

	// perform request
	reqBody := strings.NewReader("")
	res, data := kc.doAuthenticatedRequest("GET", url, reqBody)
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("error: getportfolio response not 200 (it's %d)", res.StatusCode)
	}

	// deserialize json
	var userGetMarketPositionsResponse swagger.UserGetMarketPositionsResponse
	if err := json.Unmarshal(data, &userGetMarketPositionsResponse); err != nil {
		log.Fatalf("Error: failed to unmarshal: %v", err)
	}

	for _, marketPosition := range userGetMarketPositionsResponse.MarketPositions {
		if marketPosition.Position != 0 {
			marketId := marketPosition.MarketId
			ticker := kc.GetMarketTicker(ctx, marketId)
			position := marketPosition.Position
			portfolioItem := PortfolioItem{
				Ticker:    ticker,
				Contracts: int64(math.Abs(float64(position))),
			}
			if position < 0 {
				portfolioItem.Side = No
			} else {
				portfolioItem.Side = Yes
			}
			p.Items = append(p.Items, portfolioItem)
		}
	}

	return &p, nil
}
