package kalshi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
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

func Printer() {
	fmt.Printf("hello, world from kalshi package!\n")
}
