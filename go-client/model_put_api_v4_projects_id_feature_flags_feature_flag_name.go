/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Update a feature flag
type PutApiV4ProjectsIdFeatureFlagsFeatureFlagName struct {
	// The new name of the feature flag. Supported in GitLab 13.3 and later
	Name string `json:"name,omitempty"`
	// The description of the feature flag
	Description string `json:"description,omitempty"`
	// The active state of the flag. Supported in GitLab 13.3 and later
	Active     bool                                                      `json:"active,omitempty"`
	Strategies []PutApiV4ProjectsIdFeatureFlagsFeatureFlagNameStrategies `json:"strategies,omitempty"`
}
