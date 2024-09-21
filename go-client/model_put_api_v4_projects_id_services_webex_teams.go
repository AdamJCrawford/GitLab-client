/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Create/Edit Webex Teams integration
type PutApiV4ProjectsIdServicesWebexTeams struct {
	// The Webex Teams webhook. For example, https://api.ciscospark.com/v1/webhooks/incoming/...
	Webhook string `json:"webhook"`
	// Send notifications for broken pipelines.
	NotifyOnlyBrokenPipelines bool `json:"notify_only_broken_pipelines,omitempty"`
	// Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `default_and_protected`. The default value is `default`.
	BranchesToBeNotified string `json:"branches_to_be_notified,omitempty"`
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
