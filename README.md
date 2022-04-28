# go-kalshi

Experimental Go wrapper for Kalshi API.

This is just a **work-in-progress**/toy and doesn't do anything useful yet.

It does demonstrate how to codegen Go code from Kalshi's swagger file (available [here](https://kalshi-public-docs.s3.amazonaws.com/KalshiAPI.html)).  See `pkg/kalshi/Makefile`.

## Instructions

### Build test program

```
git clone <this repo>
cd go-kalshi
make
```

### Run test program

Create a `.env` file that looks like this:

```
KALSHI_USERNAME="your kalshi login (email)"
KALSHI_PASSWORD="your kalshi password"
```

Now run the test program:

```
./run.sh
```

