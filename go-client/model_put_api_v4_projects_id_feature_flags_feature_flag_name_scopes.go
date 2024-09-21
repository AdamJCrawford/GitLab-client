/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type PutApiV4ProjectsIdFeatureFlagsFeatureFlagNameScopes struct {
	// The scope id
	Id int32 `json:"id,omitempty"`
	// The environment scope of the scope
	EnvironmentScope string `json:"environment_scope,omitempty"`
	// Delete the scope when true
	Destroy bool `json:"_destroy,omitempty"`
}
