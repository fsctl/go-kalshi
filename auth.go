package kalshi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Kalshi API authentication token.
type authToken struct {
	Token       string `json:"token"`
	UserId      string `json:"user_id"`
	AccessLevel string `json:"access_level"`
}

// Debugging function to print out a Kalshi API authentication token in its
// string representation.
func (a authToken) Text() string {
	p := fmt.Sprintf(
		"Token: %s\nUserId : %s\nAccessLevel: %s\n",
		a.Token, a.UserId, a.AccessLevel)
	return p
}

// Uses the specified Kalshi credentials to return an authToken struct.
func getAuthToken(ctx context.Context, kalshiUsername string, kalshiPassword string) (*authToken, error) {
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

	var authToken authToken
	if err := json.NewDecoder(resp.Body).Decode(&authToken); err != nil {
		log.Fatalf("ERROR:  decoding auth token response:  '%v'\n", err)
		return nil, err
	}

	return &authToken, nil
}
