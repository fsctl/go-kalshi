package main

import (
	"context"
	"fmt"
	"log"

	"github.com/fsctl/go-kalshi/pkg/kalshi"
	"github.com/spf13/viper"
)

func readEnvFile() (string, string) {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error in readEnvFile: '%v'\n", err)
	}
	kalshiUsername, ok := viper.Get("KALSHI_USERNAME").(string)
	if !ok {
		log.Fatalf("Error: KALSHI_USERNAME was not a string")
	}
	kalshiPassword, ok := viper.Get("KALSHI_PASSWORD").(string)
	if !ok {
		log.Fatalf("Error: KALSHI_PASSWORD was not a string")
	}
	return kalshiUsername, kalshiPassword
}

func main() {
	ctx := context.Background()

	fmt.Printf("Kalshi Tool\n")

	kalshiUsername, kalshiPassword := readEnvFile()

	kc, err := kalshi.NewKalshiClient(ctx, kalshiUsername, kalshiPassword)
	if err != nil {
		log.Fatalf("Error:  NewKalshiClient failed: '%v'\n", err)
		return
	}

	kc.PrintMarketsList(ctx)

}
