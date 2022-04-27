package kalshi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

	var authToken AuthToken
	if err := json.NewDecoder(resp.Body).Decode(&authToken); err != nil {
		log.Fatalf("ERROR:  decoding auth token response:  '%v'\n", err)
		return nil, err
	}

	return &authToken, nil
}
