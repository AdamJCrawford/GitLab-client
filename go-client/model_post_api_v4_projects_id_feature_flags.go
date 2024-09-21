/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Create a new feature flag
type PostApiV4ProjectsIdFeatureFlags struct {
	// The name of the feature flag
	Name string `json:"name"`
	// The description of the feature flag
	Description string `json:"description,omitempty"`
	// The active state of the flag. Defaults to `true`. Supported in GitLab 13.3 and later
	Active bool `json:"active,omitempty"`
	// The version of the feature flag. Must be `new_version_flag`. Omit to create a Legacy feature flag.
	Version    string                                      `json:"version,omitempty"`
	Strategies []PostApiV4ProjectsIdFeatureFlagsStrategies `json:"strategies,omitempty"`
}
