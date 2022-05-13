# go-kalshi

![build workflow](https://github.com/fsctl/go-kalshi/actions/workflows/makefile.yml/badge.svg)

Experimental Go wrapper for Kalshi API.

This package does not expose all of the features of the Kalshi API. However, it does:

 * Demonstrate how to codegen Go code from Kalshi's swagger file (available [here](https://kalshi-public-docs.s3.amazonaws.com/KalshiAPI.html)). See `Makefile`.
 * Show how to write higher level abstractions atop the Kalshi API
 * Create a simple CLI tool `kalshi-tool` that performs actions like listing markets, displaying market order books, etc.

## Screenshots

Listing all markets containing "CPI":

![screenshot list-markets](/img/screenshot-list-markets.png)

Placing an order to open a new No position in the [moon market](https://kalshi.com/events/MOON-25/markets/MOON-25):

![screenshot order](/img/screenshot-order.png)

`kalshi-tool` usage text:

```
$ ./kalshi-tool --help
kalshi-tool

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

'kalshi-tool list-resting'
The list-resting subcommand prints a list of all resting orders and their ids.

'kalshi-tool cancel-resting --order 1250e322-c586-5ffb-bb6e-14a9503b8997'
The cancel-resting subcommand cancels an order by its id.

'kalshi-tool portfolio'
The portfolio subcommand prints your current open positions.
```

<!-- ![screenshot help](/img/screenshot-help.png) -->

## Instructions

### Using the `kalshi-tool` command in this repo

This repo contains an example CLI command called `kalshi-tool` that uses the packages in this repo to interact with Kalshi.  It performs tasks like listing available markets, displaying the order book for a chosen market, etc.  

To run the tool, clone and create a `.env` file modeled after `.env.example` in the root of this repo.  Then, build and run:

```
make
./kalshi-tool
```

### Using go-kalshi from another repo

For an example application that depends on `go-kalshi`, check out [`fsctl/go-kalshi-usage-example`](https://github.com/fsctl/go-kalshi-usage-example).  Follow the build instructions in that repo and it will pull in `go-kalshi` automatically.
