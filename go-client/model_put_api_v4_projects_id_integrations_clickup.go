/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Create/Edit Clickup integration
type PutApiV4ProjectsIdIntegrationsClickup struct {
	// URL of the project.
	ProjectUrl string `json:"project_url"`
	// URL of the issue.
	IssuesUrl string `json:"issues_url"`
	// Trigger event for pushes to the repository.
	PushEvents bool `json:"push_events,omitempty"`
}
