# go-kalshi

Experimental Go wrapper for Kalshi API.

This is just a **work-in-progress** and doesn't do anything useful yet.

It does demonstrate how to codegen Go code from Kalshi's swagger file (available [here](https://kalshi-public-docs.s3.amazonaws.com/KalshiAPI.html)).  See `pkg/kalshi/Makefile`.

## Instructions

This is currently a package-only repo, so you don't need to clone and build it.

Instead, clone and build a dependent repo like [`fsctl/kot`](https://github.com/fsctl/kot) and it will pull this module in automatically.

#### Regenerating the swagger code

```
make clean
make
```

