/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Create/Edit Mock Ci integration
type PutApiV4ProjectsIdServicesMockCi struct {
	// Enable SSL verification. Defaults to `true` (enabled).
	EnableSslVerification bool `json:"enable_ssl_verification,omitempty"`
	// URL of the Mock CI integration.
	MockServiceUrl string `json:"mock_service_url"`
	// Trigger event for pushes to the repository.
	PushEvents bool `json:"push_events,omitempty"`
}
