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

USAGE

'kalshi-tool [--help]'
Prints this usage information.

'kalshi-tool list-markets'
The list-markets subcommand prints a list of all markets, grouped by ranged market set, with their tickers.

'kalshi-tool market --ticker ACPI-22-C3.0'
'kalshi-tool market --ticker MOON-25'
The market subcommand takes a ticker for a single binary option and prints its Yes and No order books.

'kalshi-tool order --action open-yes  --ticker=MOON-25 --contracts 1 --limit 0.20'
'kalshi-tool order --action close-yes --ticker=MOON-25 --contracts 1 --limit 0.50'
The order subcommand places an order for the specified ticker and amount with limit price.  Actions are:
  'open-yes' (buy yes)
  'close-yes' (buy no, with check that you have enough yes to cancel out)
  'open-no' (buy no)
  'close-no' (buy yes, with analogous portfolio check)
The first command opens a Yes position (buys Yes), assuming there is currently a No bid of 1-0.20 = 0.80 or higher.
If there's not, it creates a resting order for Yes with a bid price of 0.20.
The second command closes a Yes position. First it verifies that you have a Yes position of that size or larger,
then it buys a No position in that amount with a limit price (max you'll pay for No) of 0.50.
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

	// 'market' and 'order' subcommand flags
	marketSubCmd := flag.NewFlagSet("market", flag.ExitOnError)
	marketTicker := marketSubCmd.String("ticker", "", "ticker, e.g., MOON-25 (case sensitive)")
	orderSubCmd := flag.NewFlagSet("order", flag.ExitOnError)
	orderTicker := orderSubCmd.String("ticker", "", "ticker, e.g., MOON-25 (case sensitive)")
	orderAction := orderSubCmd.String("action", "", "one of:  open-yes, open-no, close-yes, close-no")
	orderContracts := orderSubCmd.Int("contracts", 0, "number of contracts")
	orderLimit := orderSubCmd.Float64("limit", 0.0, "limit price")

	subcommand := os.Args[1]
	switch subcommand {
	case "list-markets":
		kc.PrintMarketsList(ctx)
	case "market":
		marketSubCmd.Parse(os.Args[2:])
		marketId := kc.GetMarketId(ctx, *marketTicker)
		kc.PrintMarket(ctx, marketId)
	case "order":
		orderSubCmd.Parse(os.Args[2:])

		// check required args are set that won't be caught elsewhere
		if *orderContracts < 1 {
			log.Fatalf("Error: --contracts not set\n")
		}
		if *orderLimit < 0.01 || *orderLimit > 0.99 {
			log.Fatalf("Error: --limit not set or out of bounds (0.01 <= limit <= 0.99)\n")
		}

		// invalid ticker will halt program
		marketId := kc.GetMarketId(ctx, *orderTicker)

		fmt.Printf("Order:\n")
		fmt.Printf("  Ticker:    %v\n", *orderTicker)
		fmt.Printf("  Market Id: %v\n", marketId)
		fmt.Printf("  Action:    %v\n", *orderAction)
		fmt.Printf("  Contracts: %v\n", *orderContracts)
		fmt.Printf("  Limit:     %.2f\n", *orderLimit)

		switch *orderAction {
		case "open-yes":
			kc.OrderOpenPosition(ctx, marketId, kalshi.Yes, *orderContracts, *orderLimit)
			if err != nil {
				log.Fatalf("Error: open-yes order failed: %v\n", err)
			}
		case "open-no":
			kc.OrderOpenPosition(ctx, marketId, kalshi.No, *orderContracts, *orderLimit)
			if err != nil {
				log.Fatalf("Error: open-no order failed: %v\n", err)
			}
		case "close-yes":
			kc.OrderClosePosition(ctx, marketId, kalshi.Yes, *orderContracts, *orderLimit)
			if err != nil {
				log.Fatalf("Error: close-yes order failed: %v\n", err)
			}
		case "close-no":
			kc.OrderClosePosition(ctx, marketId, kalshi.No, *orderContracts, *orderLimit)
			if err != nil {
				log.Fatalf("Error: close-no order failed: %v\n", err)
			}
		default:
			log.Fatalf("Error: unrecognized action '%v'\n", *orderAction)
		}
	default:
		log.Fatalf("Error: did not recognize subcommand '%v'\n", subcommand)
	}
}
