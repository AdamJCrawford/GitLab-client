/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Create/Edit Jira integration
type PutApiV4ProjectsIdServicesJira struct {
	// The base URL to the Jira instance web interface which is being linked to this GitLab project. E.g., https://jira.example.com
	Url string `json:"url"`
	// The base URL to the Jira instance API. Web URL value will be used if not set. E.g., https://jira-api.example.com
	ApiUrl string `json:"api_url,omitempty"`
	// The authentication method to be used with Jira. `0` means Basic Authentication. `1` means Jira personal access token. Defaults to `0`
	JiraAuthType int32 `json:"jira_auth_type,omitempty"`
	// The email or username to be used with Jira. For Jira Cloud use an email, for Jira Data Center and Jira Server use a username. Required when using Basic authentication (`jira_auth_type` is `0`)
	Username string `json:"username,omitempty"`
	// The Jira API token, password, or personal access token to be used with Jira. When your authentication method is Basic (`jira_auth_type` is `0`) use an API token for Jira Cloud, or a password for Jira Data Center or Jira Server. When your authentication method is Jira personal access token (`jira_auth_type` is `1`) use a personal access token
	Password string `json:"password"`
	// Enable automatic issue transitions
	JiraIssueTransitionAutomatic bool `json:"jira_issue_transition_automatic,omitempty"`
	// The ID of one or more transitions for custom issue transitions
	JiraIssueTransitionId string `json:"jira_issue_transition_id,omitempty"`
	// Prefix to match Jira issue keys
	JiraIssuePrefix string `json:"jira_issue_prefix,omitempty"`
	// Regular expression to match Jira issue keys
	JiraIssueRegex string `json:"jira_issue_regex,omitempty"`
	// Enable viewing Jira issues in GitLab
	IssuesEnabled bool `json:"issues_enabled,omitempty"`
	// Keys of Jira projects to view issues from in GitLab
	ProjectKeys []string `json:"project_keys,omitempty"`
	// Enable comments inside Jira issues on each GitLab event (commit / merge request)
	CommentOnEventEnabled bool `json:"comment_on_event_enabled,omitempty"`
	// Trigger event when a commit is created or updated.
	CommitEvents bool `json:"commit_events,omitempty"`
	// Trigger event when a merge request is created, updated, or merged.
	MergeRequestsEvents bool `json:"merge_requests_events,omitempty"`
}