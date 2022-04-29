#!/bin/sh

# Make sure user specified a program to run
PROG=$1
if [[ -z "$PROG" ]]; then
	echo "Error:   missing name of program to run"
	echo "Usage:   ./run.sh [program]"
    echo "Example: ./run.sh kalshi-tool"
	exit 1
fi

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
./$PROG
