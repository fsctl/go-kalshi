/*
 * Kalshi API.
 *
 * This documentation describes Kalshi's rest API for market makers  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Olhcm struct {
	Close int64 `json:"close,omitempty"`
	High int64 `json:"high,omitempty"`
	Low int64 `json:"low,omitempty"`
	Mean int64 `json:"mean,omitempty"`
	Open int64 `json:"open,omitempty"`
}
