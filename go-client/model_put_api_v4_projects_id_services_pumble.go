/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Create/Edit Pumble integration
type PutApiV4ProjectsIdServicesPumble struct {
	// The Pumble chat webhook. For example, https://api.pumble.com/workspaces/x/...
	Webhook string `json:"webhook"`
	// Trigger event for pushes to the repository.
	PushEvents bool `json:"push_events,omitempty"`
	// Trigger event when an issue is created, updated, or closed.
	IssuesEvents bool `json:"issues_events,omitempty"`
	// Trigger event when a confidential issue is created, updated, or closed.
	ConfidentialIssuesEvents bool `json:"confidential_issues_events,omitempty"`
	// Trigger event when a merge request is created, updated, or merged.
	MergeRequestsEvents bool `json:"merge_requests_events,omitempty"`
	// Trigger event for new comments.
	NoteEvents bool `json:"note_events,omitempty"`
	// Trigger event for new comments on confidential issues.
	ConfidentialNoteEvents bool `json:"confidential_note_events,omitempty"`
	// Trigger event for new tags pushed to the repository.
	TagPushEvents bool `json:"tag_push_events,omitempty"`
	// Trigger event when a pipeline status changes.
	PipelineEvents bool `json:"pipeline_events,omitempty"`
	// Trigger event when a wiki page is created or updated.
	WikiPageEvents bool `json:"wiki_page_events,omitempty"`
}
