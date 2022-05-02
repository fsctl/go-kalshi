package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fsctl/go-kalshi/pkg/kalshi"
	"github.com/fsctl/go-kalshi/pkg/kalshi/helpers"
)

func usage() {
	usage := `kalshi-tool

Usage

'kalshi-tool [--help]'
Prints this usage information.

'kalshi-tool list-markets'
The list-markets subcommand prints a list of all markets, grouped by ranged market set, with their tickers.

'kalshi-tool market --ticker ACPI-22-C3.0'
'kalshi-tool market --ticker MOON-25'
The market subcommand takes a ticker for a single binary option and prints its Yes and No order books.
`

	fmt.Println(usage)
}

func main() {
	ctx := context.Background()

	kalshiUsername, kalshiPassword := helpers.ReadEnvFile()

	kc, err := kalshi.NewKalshiClient(ctx, kalshiUsername, kalshiPassword)
	if err != nil {
		log.Fatalf("Error:  NewKalshiClient failed: '%v'\n", err)
		return
	}

	// Trap --help manually
	if (len(os.Args) < 2) || (len(os.Args) == 2 && (os.Args[1] == "--help" || os.Args[1] == "-h")) {
		usage()
		os.Exit(0)
	}

	// 'market' subcommand flags
	marketSubCmd := flag.NewFlagSet("market", flag.ExitOnError)
	marketTicker := marketSubCmd.String("ticker", "", "ticker")

	subcommand := os.Args[1]
	switch subcommand {
	case "list-markets":
		kc.PrintMarketsList(ctx)
	case "market":
		marketSubCmd.Parse(os.Args[2:])
		marketId := kc.GetMarketId(ctx, *marketTicker)
		kc.PrintMarket(ctx, marketId)
	default:
		log.Fatalf("Error: did not recognize subcommand '%v'\n", subcommand)
	}
}
