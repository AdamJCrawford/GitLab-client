/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ApiEntitiesGitlabSubscriptionPlan struct {
	Code          string `json:"code,omitempty"`
	Name          string `json:"name,omitempty"`
	Trial         string `json:"trial,omitempty"`
	AutoRenew     string `json:"auto_renew,omitempty"`
	Upgradable    string `json:"upgradable,omitempty"`
	ExcludeGuests string `json:"exclude_guests,omitempty"`
}
