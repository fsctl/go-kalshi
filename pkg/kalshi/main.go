package kalshi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fsctl/go-kalshi/pkg/kalshi/swagger"
	//"github.com/fsctl/go-kalshi/pkg/kalshi/KalshiSwagger"
)

const (
	BaseURLV1 = "https://trading-api.kalshi.com/v1"
)

type AuthToken struct {
	Token       string `json:"token"`
	UserId      string `json:"user_id"`
	AccessLevel string `json:"access_level"`
}

func (a AuthToken) Text() string {
	p := fmt.Sprintf(
		"Token: %s\nUserId : %s\nAccessLevel: %s\n",
		a.Token, a.UserId, a.AccessLevel)
	return p
}

type Client struct {
	BaseURL    string
	userId     string
	token      string
	HTTPClient *http.Client
}

func (c Client) Text() string {
	p := fmt.Sprintf(
		"Token: %s\nUserId : %s\nBaseURL: %s\n",
		c.token, c.userId, c.BaseURL)
	return p
}

func getAuthToken(ctx context.Context, kalshiUsername string, kalshiPassword string) (*AuthToken, error) {
	postBody, _ := json.Marshal(map[string]string{
		"email":    kalshiUsername,
		"password": kalshiPassword,
	})
	postBodyBuf := bytes.NewBuffer(postBody)
	resp, err := http.Post(fmt.Sprintf("%s/log_in", BaseURLV1), "application/json", postBodyBuf)
	if err != nil {
		log.Fatalf("Error: '%v'\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	// // Print the raw response
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalf("ERROR:  '%v'\n", err)
	//  return nil, err
	// }
	// sb := string(body)
	// fmt.Printf("Body:\n%v\n\n", sb)

	var authTokenResp AuthToken
	if err := json.NewDecoder(resp.Body).Decode(&authTokenResp); err != nil {
		log.Fatalf("ERROR:  decoding auth token response:  '%v'\n", err)
		return nil, err
	}

	return &authTokenResp, nil
}

func NewClient(ctx context.Context, kalshiUsername string, kalshiPassword string) (*Client, error) {
	authToken, err := getAuthToken(ctx, kalshiUsername, kalshiPassword)
	if err != nil {
		log.Fatalf("ERROR:  getAuthToken failed:  '%v'\n", err)
		return nil, err
	}

	return &Client{
		BaseURL: BaseURLV1,
		userId:  authToken.UserId,
		token:   authToken.Token,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}, nil
}

func Printer(ctx context.Context) {
	cfg := swagger.NewConfiguration()
	cfg.AddDefaultHeader("Authorization", "f1f9dbfa-810b-4c14-b2fd-91ce2c229bed f1f9dbfa-810b-4c14-b2fd-91ce2c229bed:nsowBaVhS1oE8MxRUG4lQtJmtagAD36WtEkUykcXtk3NKLipnf96WeKjOJ6D5g4b")

	apiClient := swagger.NewAPIClient(cfg)

	userGetMarketsResponse, _, err := apiClient.MarketApi.GetMarkets(ctx)
	if err != nil {
		log.Fatalf("Error from apiClient.MarketApi.GetMarkets:  '%v'", err)
		return
	}

	for _, market := range userGetMarketsResponse.Markets {
		fmt.Printf("Title:  %v, Category:  %v\n", market.Title, market.Category)
	}
}
