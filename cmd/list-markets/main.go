package main

import (
	"fmt"
	"os"
	"github.com/fsctl/go-kalshi/pkg/kalshi"
)

func main() {
	fmt.Printf("hello, world!\n")
	Main2Hello()
	kalshi.Printer()

	fmt.Printf("KALSHI_USERNAME=%s\n",os.Getenv("KALSHI_USERNAME"))
	fmt.Printf("KALSHI_PASSWORD=%s\n",os.Getenv("KALSHI_PASSWORD"))
}
