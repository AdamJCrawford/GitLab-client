/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Create/Edit Pivotaltracker integration
type PutApiV4ProjectsIdIntegrationsPivotaltracker struct {
	// The Pivotaltracker token
	Token string `json:"token"`
	// Comma-separated list of branches which will be automatically inspected. Leave blank to include all branches.
	RestrictToBranch string `json:"restrict_to_branch,omitempty"`
	// Trigger event for pushes to the repository.
	PushEvents bool `json:"push_events,omitempty"`
}
