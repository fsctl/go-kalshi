package kalshi

import (
	"fmt"
	"log"
)

// define enum of market sides
type MarketSide int64

const (
	Yes = iota
	No
)

func (s MarketSide) String() string {
	switch s {
	case Yes:
		return "yes"
	case No:
		return "no"
	default:
		log.Fatalf("Error: enum value not found in lookup table '%v'", int(s))
		return fmt.Sprintf("%d", int(s))
	}
}

func (s MarketSide) Opposite() MarketSide {
	switch s {
	case Yes:
		return No
	case No:
		return Yes
	default:
		log.Fatalf("Error: enum value not found in lookup table '%v'", int(s))
		return No // arbitrary
	}
}
