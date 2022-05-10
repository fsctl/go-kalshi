package kalshi

import (
	"context"
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
