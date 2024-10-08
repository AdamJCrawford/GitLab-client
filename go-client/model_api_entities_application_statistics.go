/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// API_Entities_ApplicationStatistics model
type ApiEntitiesApplicationStatistics struct {
	// Approximate number of repo forks
	Forks int32 `json:"forks,omitempty"`
	// Approximate number of issues
	Issues int32 `json:"issues,omitempty"`
	// Approximate number of merge requests
	MergeRequests int32 `json:"merge_requests,omitempty"`
	// Approximate number of notes
	Notes int32 `json:"notes,omitempty"`
	// Approximate number of snippets
	Snippets int32 `json:"snippets,omitempty"`
	// Approximate number of SSH keys
	SshKeys int32 `json:"ssh_keys,omitempty"`
	// Approximate number of milestones
	Milestones int32 `json:"milestones,omitempty"`
	// Approximate number of users
	Users int32 `json:"users,omitempty"`
	// Approximate number of projects
	Projects int32 `json:"projects,omitempty"`
	// Approximate number of projects
	Groups int32 `json:"groups,omitempty"`
	// Number of active users
	ActiveUsers int32 `json:"active_users,omitempty"`
}
