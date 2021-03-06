/*
 * Kalshi API.
 *
 * This documentation describes Kalshi's rest API for market makers  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type UserGetReferralInfoResponse struct {
	EligibleToRefer bool `json:"eligible_to_refer,omitempty"`
	NumContractsTraded int64 `json:"num_contracts_traded,omitempty"`
	PeopleReferred *[]PeopleReferred `json:"people_referred,omitempty"`
	ReferralCode string `json:"referral_code,omitempty"`
	ReferralId string `json:"referral_id,omitempty"`
	ReferralMoneyRewarded int64 `json:"referral_money_rewarded,omitempty"`
	StageInFunnel string `json:"stage_in_funnel,omitempty"`
}
