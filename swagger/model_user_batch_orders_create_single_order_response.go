/*
 * Kalshi API.
 *
 * This documentation describes Kalshi's rest API for market makers  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Response for submitting an order
type UserBatchOrdersCreateSingleOrderResponse struct {
	Error_ *JsonError `json:"error,omitempty"`
	Order *Order `json:"order,omitempty"`
	// Status of the order submit operation
	Status string `json:"status,omitempty"`
}
