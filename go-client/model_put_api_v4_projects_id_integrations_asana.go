/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Create/Edit Asana integration
type PutApiV4ProjectsIdIntegrationsAsana struct {
	// User API token. The user must have access to the task. All comments are attributed to this user.
	ApiKey string `json:"api_key"`
	// Comma-separated list of branches to be automatically inspected. Leave blank to include all branches.
	RestrictToBranch string `json:"restrict_to_branch,omitempty"`
	// Trigger event for pushes to the repository.
	PushEvents bool `json:"push_events,omitempty"`
}
