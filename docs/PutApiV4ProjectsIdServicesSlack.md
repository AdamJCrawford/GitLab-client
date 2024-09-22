# PutApiV4ProjectsIdServicesSlack

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Webhook** | **string** | Slack notifications webhook (for example, &#x60;https://hooks.slack.com/services/...&#x60;). | [default to null]
**Username** | **string** | Slack notifications username. | [optional] [default to null]
**Channel** | **string** | Default channel to use if no other channel is configured. | [optional] [default to null]
**NotifyOnlyBrokenPipelines** | **bool** | Send notifications for broken pipelines. | [optional] [default to null]
**BranchesToBeNotified** | **string** | Branches to send notifications for. Valid options are &#x60;all&#x60;, &#x60;default&#x60;, &#x60;protected&#x60;, and &#x60;default_and_protected&#x60;. The default value is &#x60;default&#x60;. | [optional] [default to null]
**LabelsToBeNotified** | **string** | Labels to send notifications for. Leave blank to receive notifications for all events. | [optional] [default to null]
**LabelsToBeNotifiedBehavior** | **string** | Labels to be notified for. Valid options are &#x60;match_any&#x60; and &#x60;match_all&#x60;. The default value is &#x60;match_any&#x60;. | [optional] [default to null]
**PushChannel** | **string** | The name of the channel to receive push_events notifications | [optional] [default to null]
**IssueChannel** | **string** | The name of the channel to receive issues_events notifications | [optional] [default to null]
**IncidentChannel** | **string** | The name of the channel to receive incident_events notifications | [optional] [default to null]
**AlertChannel** | **string** | The name of the channel to receive alert_events notifications | [optional] [default to null]
**ConfidentialIssueChannel** | **string** | The name of the channel to receive confidential_issues_events notifications | [optional] [default to null]
**MergeRequestChannel** | **string** | The name of the channel to receive merge_requests_events notifications | [optional] [default to null]
**NoteChannel** | **string** | The name of the channel to receive note_events notifications | [optional] [default to null]
**ConfidentialNoteChannel** | **string** | The name of the channel to receive confidential_note_events notifications | [optional] [default to null]
**TagPushChannel** | **string** | The name of the channel to receive tag_push_events notifications | [optional] [default to null]
**DeploymentChannel** | **string** | The name of the channel to receive deployment_events notifications | [optional] [default to null]
**PipelineChannel** | **string** | The name of the channel to receive pipeline_events notifications | [optional] [default to null]
**WikiPageChannel** | **string** | The name of the channel to receive wiki_page_events notifications | [optional] [default to null]
**VulnerabilityChannel** | **string** | The name of the channel to receive vulnerability_events notifications | [optional] [default to null]
**PushEvents** | **bool** | Trigger event for pushes to the repository. | [optional] [default to null]
**IssuesEvents** | **bool** | Trigger event when an issue is created, updated, or closed. | [optional] [default to null]
**ConfidentialIssuesEvents** | **bool** | Trigger event when a confidential issue is created, updated, or closed. | [optional] [default to null]
**MergeRequestsEvents** | **bool** | Trigger event when a merge request is created, updated, or merged. | [optional] [default to null]
**NoteEvents** | **bool** | Trigger event for new comments. | [optional] [default to null]
**ConfidentialNoteEvents** | **bool** | Trigger event for new comments on confidential issues. | [optional] [default to null]
**TagPushEvents** | **bool** | Trigger event for new tags pushed to the repository. | [optional] [default to null]
**PipelineEvents** | **bool** | Trigger event when a pipeline status changes. | [optional] [default to null]
**WikiPageEvents** | **bool** | Trigger event when a wiki page is created or updated. | [optional] [default to null]
**DeploymentEvents** | **bool** | Trigger event when a deployment starts or finishes. | [optional] [default to null]
**IncidentEvents** | **bool** | Trigger event when an incident is created. | [optional] [default to null]
**VulnerabilityEvents** | **bool** |  | [optional] [default to null]
**AlertEvents** | **bool** | Trigger event when a new, unique alert is recorded. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


