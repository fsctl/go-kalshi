/*
go-kalshi is a Kalshi API client library.

go-kalshi exports functions that wrap Kalshi API calls, but provide higher level
abstractions for common operations.  For example, OrderClosePosition() closes a
position in a market by buying the opposite side after first checking that you
actually have an open position in the specified market.

go-kalshi is intended to be a toolkit on which trading bots can be built. Its
embedded `cmd/kalshi-tool` is a simple CLI application that demonstrates the
functionality of the library.

Under the hood, go-kalshi uses Swagger Codegen to generate the types used in
making Kalshi API calls. Unfortunately, the codegen'd code does not work fully,
so we only use the structs generated and make most API calls manually (see
`doAuthenticatedRequest` internal function).
*/
package kalshi
