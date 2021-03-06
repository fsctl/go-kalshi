/*
 * Kalshi API.
 *
 * This documentation describes Kalshi's rest API for market makers  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RangedMarket struct {
	Id string `json:"id,omitempty"`
	MiniTitle string `json:"mini_title,omitempty"`
	MutuallyExclusiveSide string `json:"mutually_exclusive_side,omitempty"`
	Ticker string `json:"ticker,omitempty"`
	Title string `json:"title,omitempty"`
}
