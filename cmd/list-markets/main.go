package main

import (
	"context"
	"log"
	"os"

	"github.com/fsctl/go-kalshi/pkg/kalshi"
)

func main() {
	ctx := context.Background()

	username := os.Getenv("KALSHI_USERNAME")
	password := os.Getenv("KALSHI_PASSWORD")
	if username == "" || password == "" {
		log.Fatalf("Error:  username/password environment variables not set\n")
		return
	}

	kc, err := kalshi.NewKalshiClient(ctx, username, password)
	if err != nil {
		log.Fatalf("Error:  NewKalshiClient failed: '%v'\n", err)
		return
	}

	kc.PrintMarketsList(ctx)
}
