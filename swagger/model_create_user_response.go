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
type CreateUserResponse struct {
	// swagger: ignore
	Code string `json:"code,omitempty"`
	// user_id for the created user.
	UserId string `json:"user_id"`
}
