/*
 * Kalshi API.
 *
 * This documentation describes Kalshi's rest API for market makers  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Watchlist is the list of markets that you have some activity on, this is used mostly by the UI.
type Watchlist struct {
	MarketIds []string `json:"market_ids,omitempty"`
}
