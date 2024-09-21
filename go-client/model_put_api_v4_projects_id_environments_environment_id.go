/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Update an existing environment
type PutApiV4ProjectsIdEnvironmentsEnvironmentId struct {
	// The new URL on which this deployment is viewable
	ExternalUrl string `json:"external_url,omitempty"`
	// The tier of the new environment. Allowed values are `production`, `staging`, `testing`, `development`, and `other`
	Tier string `json:"tier,omitempty"`
}
