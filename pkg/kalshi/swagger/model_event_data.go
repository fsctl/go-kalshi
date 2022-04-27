/*
 * Kalshi API.
 *
 * This documentation describes Kalshi's rest API for market makers  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type EventData struct {
	Category string `json:"category,omitempty"`
	DescriptionContext string `json:"description_context,omitempty"`
	Markets []EventChildMarket `json:"markets,omitempty"`
	MetricsTags []string `json:"metrics_tags,omitempty"`
	MinTickSize string `json:"min_tick_size,omitempty"`
	MiniTitle string `json:"mini_title,omitempty"`
	MutuallyExclusive bool `json:"mutually_exclusive,omitempty"`
	MutuallyExclusiveSide string `json:"mutually_exclusive_side,omitempty"`
	SeriesTicker string `json:"series_ticker,omitempty"`
	SettleDetails string `json:"settle_details,omitempty"`
	SettlementSources []SettlementSource `json:"settlement_sources,omitempty"`
	SubTitle string `json:"sub_title,omitempty"`
	Tags []string `json:"tags,omitempty"`
	TargetDatetime time.Time `json:"target_datetime,omitempty"`
	Ticker string `json:"ticker,omitempty"`
	Title string `json:"title,omitempty"`
	Underlying string `json:"underlying,omitempty"`
}
