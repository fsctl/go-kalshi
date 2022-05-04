/*
 * Kalshi API.
 *
 * This documentation describes Kalshi's rest API for market makers  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type UserUpdateProfileRequest struct {
	// User's phone area code.
	AreaCode string `json:"area_code"`
	BirthDate string `json:"birth_date,omitempty"`
	City string `json:"city,omitempty"`
	// User's country 2 digits code
	Country string `json:"country"`
	// User's phone country code. Should be 1 for now because only USA accounts are accepted.
	CountryCode string `json:"country_code"`
	FinishedFre bool `json:"finished_fre,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName string `json:"last_name,omitempty"`
	// User's phone number.
	PhoneNumber string `json:"phone_number"`
	// User's address postal code
	PostalCode string `json:"postal_code"`
	// User's state 2 digits code
	State string `json:"state"`
	Street1 string `json:"street1,omitempty"`
	Street2 string `json:"street2,omitempty"`
	UseBidAsk bool `json:"use_bid_ask,omitempty"`
	Watchlist []string `json:"watchlist,omitempty"`
}