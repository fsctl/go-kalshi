/*
 * Kalshi API.
 *
 * This documentation describes Kalshi's rest API for market makers  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type LoginRequest struct {
	// Email should be used as login identification credentials.
	Email string `json:"email"`
	// Password defined in the first step of the sign-up.
	Password string `json:"password"`
}
