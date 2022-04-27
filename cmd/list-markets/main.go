package main

import (
	"context"
	"os"

	"github.com/fsctl/go-kalshi/pkg/kalshi"
)

func main() {
	ctx := context.Background()

	Main2Hello()

	username := os.Getenv("KALSHI_USERNAME")
	password := os.Getenv("KALSHI_PASSWORD")
	kalshi.PrintMarketsList(ctx, username, password)
}
