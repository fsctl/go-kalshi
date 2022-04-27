package main

import (
	"context"

	"github.com/fsctl/go-kalshi/pkg/kalshi"
)

func main() {
	ctx := context.Background()

	//fmt.Printf("hello, world!\n")
	//Main2Hello()
	kalshi.Printer(ctx)

	//////////////////////////////////////////////////////////////////

	// fmt.Printf("KALSHI_USERNAME=%s\n", os.Getenv("KALSHI_USERNAME"))
	// fmt.Printf("KALSHI_PASSWORD=%s\n", os.Getenv("KALSHI_PASSWORD"))
	// fmt.Printf("\n")

	// username := os.Getenv("KALSHI_USERNAME")
	// password := os.Getenv("KALSHI_PASSWORD")

	// client, err := kalshi.NewClient(ctx, username, password)
	// if err != nil {
	// 	log.Fatalf("Error:  could not instantiate client (wrong user/pass?)\n")
	// 	return
	// }

	// fmt.Printf("Success!\n%v\n", client.Text())
}
