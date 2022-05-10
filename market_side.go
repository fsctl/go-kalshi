package kalshi

import (
	"fmt"
	"log"
)

// MarketSide is an enum of the two market sides.
type MarketSide int64

// These are the two market side values that a MarketSide
// variable can be set to.
const (
	Yes = iota
	No
)

// Converts a MarketSide value to its string representation.
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

// Returns the opposite side of the MarketSide value. The opposite of Yes is No
// and vice-versa.
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
