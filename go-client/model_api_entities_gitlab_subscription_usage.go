/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ApiEntitiesGitlabSubscriptionUsage struct {
	SeatsInSubscription string `json:"seats_in_subscription,omitempty"`
	SeatsInUse          string `json:"seats_in_use,omitempty"`
	MaxSeatsUsed        string `json:"max_seats_used,omitempty"`
	SeatsOwed           string `json:"seats_owed,omitempty"`
}
