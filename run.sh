#!/bin/sh

# Make sure '.env' file exists
if [[ ! -f ".env" ]]; then
	echo "Error: missing '.env' file that sets KALSHI_USERNAME and KALSHI_PASSWORD"
	exit 1
fi

# Expose username/password in the environment
source .env
export KALSHI_USERNAME
export KALSHI_PASSWORD

# Make sure username/password are non-blank
if [[ -z "$KALSHI_USERNAME" ]] || [[ -z "$KALSHI_PASSWORD" ]]; then
	echo "Error: .env file needs to set KALSHI_USERNAME and KALSHI_PASSWORD variables"
	exit 1
fi

# Run test binary
./list-markets
