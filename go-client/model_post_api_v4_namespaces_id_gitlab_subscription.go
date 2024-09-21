/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Create a subscription for the namespace
type PostApiV4NamespacesIdGitlabSubscription struct {
	// The date when subscription was started
	StartDate string `json:"start_date"`
	// Number of seats in subscription
	Seats int32 `json:"seats,omitempty"`
	// Highest number of active users in the last month
	MaxSeatsUsed int32 `json:"max_seats_used,omitempty"`
	// Subscription tier code
	PlanCode string `json:"plan_code,omitempty"`
	// End date of subscription
	EndDate string `json:"end_date,omitempty"`
	// Whether subscription will auto renew on end date
	AutoRenew bool `json:"auto_renew,omitempty"`
	// Whether the subscription is a trial
	Trial bool `json:"trial,omitempty"`
	// End date of trial
	TrialEndsOn string `json:"trial_ends_on,omitempty"`
	// Start date of trial
	TrialStartsOn string `json:"trial_starts_on,omitempty"`
	// Whether subscription is an extended or reactivated trial
	TrialExtensionType int32 `json:"trial_extension_type,omitempty"`
}
