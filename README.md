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

![screenshot help](/img/screenshot-help.png)

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
