# go-kalshi

Experimental Go wrapper for Kalshi API.

This is just a work-in-progress and doesn't do much that's useful yet.

It does demonstrate how to codegen Go code from Kalshi's swagger file (available [here](https://kalshi-public-docs.s3.amazonaws.com/KalshiAPI.html)) and use that to build higher level abstractions atop the Kalshi API.  See `pkg/kalshi/Makefile`.

## Instructions

### Using go-kalshi from another repo

For an example of writing an application that depends on `go-kalshi`, check out [`fsctl/go-kalshi-usage-example`](https://github.com/fsctl/go-kalshi-usage-example).  Follow the build instructions in that repo and it will pull in `go-kalshi` automatically.

### Using the `kalshi-tool` command in this repo

This repo contains an example command called `kalshi-tool` that supports interaction with Kalshi on the CLI.  First, create a `.env` file modeled after the example one in the root of this repo.  Then, build and run:

```
make
./kalshi-tool
```
