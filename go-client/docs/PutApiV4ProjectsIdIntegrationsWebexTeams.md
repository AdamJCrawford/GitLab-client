# PutApiV4ProjectsIdIntegrationsWebexTeams

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Webhook** | **string** | The Webex Teams webhook. For example, https://api.ciscospark.com/v1/webhooks/incoming/... | [default to null]
**NotifyOnlyBrokenPipelines** | **bool** | Send notifications for broken pipelines. | [optional] [default to null]
**BranchesToBeNotified** | **string** | Branches to send notifications for. Valid options are &#x60;all&#x60;, &#x60;default&#x60;, &#x60;protected&#x60;, and &#x60;default_and_protected&#x60;. The default value is &#x60;default&#x60;. | [optional] [default to null]
**PushEvents** | **bool** | Trigger event for pushes to the repository. | [optional] [default to null]
**IssuesEvents** | **bool** | Trigger event when an issue is created, updated, or closed. | [optional] [default to null]
**ConfidentialIssuesEvents** | **bool** | Trigger event when a confidential issue is created, updated, or closed. | [optional] [default to null]
**MergeRequestsEvents** | **bool** | Trigger event when a merge request is created, updated, or merged. | [optional] [default to null]
**NoteEvents** | **bool** | Trigger event for new comments. | [optional] [default to null]
**ConfidentialNoteEvents** | **bool** | Trigger event for new comments on confidential issues. | [optional] [default to null]
**TagPushEvents** | **bool** | Trigger event for new tags pushed to the repository. | [optional] [default to null]
**PipelineEvents** | **bool** | Trigger event when a pipeline status changes. | [optional] [default to null]
**WikiPageEvents** | **bool** | Trigger event when a wiki page is created or updated. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


