# PutApiV4ProjectsIdIntegrationsJira

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The base URL to the Jira instance web interface which is being linked to this GitLab project. E.g., https://jira.example.com | [default to null]
**ApiUrl** | **string** | The base URL to the Jira instance API. Web URL value will be used if not set. E.g., https://jira-api.example.com | [optional] [default to null]
**JiraAuthType** | **int32** | The authentication method to be used with Jira. &#x60;0&#x60; means Basic Authentication. &#x60;1&#x60; means Jira personal access token. Defaults to &#x60;0&#x60; | [optional] [default to null]
**Username** | **string** | The email or username to be used with Jira. For Jira Cloud use an email, for Jira Data Center and Jira Server use a username. Required when using Basic authentication (&#x60;jira_auth_type&#x60; is &#x60;0&#x60;) | [optional] [default to null]
**Password** | **string** | The Jira API token, password, or personal access token to be used with Jira. When your authentication method is Basic (&#x60;jira_auth_type&#x60; is &#x60;0&#x60;) use an API token for Jira Cloud, or a password for Jira Data Center or Jira Server. When your authentication method is Jira personal access token (&#x60;jira_auth_type&#x60; is &#x60;1&#x60;) use a personal access token | [default to null]
**JiraIssueTransitionAutomatic** | **bool** | Enable automatic issue transitions | [optional] [default to null]
**JiraIssueTransitionId** | **string** | The ID of one or more transitions for custom issue transitions | [optional] [default to null]
**JiraIssuePrefix** | **string** | Prefix to match Jira issue keys | [optional] [default to null]
**JiraIssueRegex** | **string** | Regular expression to match Jira issue keys | [optional] [default to null]
**IssuesEnabled** | **bool** | Enable viewing Jira issues in GitLab | [optional] [default to null]
**ProjectKeys** | **[]string** | Keys of Jira projects to view issues from in GitLab | [optional] [default to null]
**CommentOnEventEnabled** | **bool** | Enable comments inside Jira issues on each GitLab event (commit / merge request) | [optional] [default to null]
**CommitEvents** | **bool** | Trigger event when a commit is created or updated. | [optional] [default to null]
**MergeRequestsEvents** | **bool** | Trigger event when a merge request is created, updated, or merged. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


