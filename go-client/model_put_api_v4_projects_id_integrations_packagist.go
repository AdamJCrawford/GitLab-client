/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Create/Edit Packagist integration
type PutApiV4ProjectsIdIntegrationsPackagist struct {
	// The username
	Username string `json:"username"`
	// The Packagist API token
	Token string `json:"token"`
	// The server
	Server string `json:"server,omitempty"`
	// Trigger event for pushes to the repository.
	PushEvents bool `json:"push_events,omitempty"`
	// Trigger event when a merge request is created, updated, or merged.
	MergeRequestsEvents bool `json:"merge_requests_events,omitempty"`
	// Trigger event for new tags pushed to the repository.
	TagPushEvents bool `json:"tag_push_events,omitempty"`
}