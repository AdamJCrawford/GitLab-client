/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type ApiEntitiesGitlabSubscriptionBilling struct {
	SubscriptionStartDate string `json:"subscription_start_date,omitempty"`
	SubscriptionEndDate   string `json:"subscription_end_date,omitempty"`
	TrialEndsOn           string `json:"trial_ends_on,omitempty"`
}
