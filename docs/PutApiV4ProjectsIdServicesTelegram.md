# PutApiV4ProjectsIdServicesTelegram

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** | Custom hostname of the Telegram API. The default value is &#x60;https://api.telegram.org&#x60;. | [optional] [default to null]
**Token** | **string** | The Telegram chat token. For example, 123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11 | [default to null]
**Room** | **string** | Unique identifier for the target chat or username of the target channel (in the format @channelusername) | [default to null]
**Thread** | **int32** | Unique identifier for the target message thread (topic in a forum supergroup) | [optional] [default to null]
**BranchesToBeNotified** | **string** | Branches for which notifications are to be sent. | [optional] [default to null]
**NotifyOnlyBrokenPipelines** | **bool** | Send notifications for broken pipelines | [optional] [default to null]
**PushEvents** | **bool** | Trigger event for pushes to the repository. | [optional] [default to null]
**IssuesEvents** | **bool** | Trigger event when an issue is created, updated, or closed. | [optional] [default to null]
**ConfidentialIssuesEvents** | **bool** | Trigger event when a confidential issue is created, updated, or closed. | [optional] [default to null]
**MergeRequestsEvents** | **bool** | Trigger event when a merge request is created, updated, or merged. | [optional] [default to null]
**NoteEvents** | **bool** | Trigger event for new comments. | [optional] [default to null]
**ConfidentialNoteEvents** | **bool** | Trigger event for new comments on confidential issues. | [optional] [default to null]
**TagPushEvents** | **bool** | Trigger event for new tags pushed to the repository. | [optional] [default to null]
**PipelineEvents** | **bool** | Trigger event when a pipeline status changes. | [optional] [default to null]
**WikiPageEvents** | **bool** | Trigger event when a wiki page is created or updated. | [optional] [default to null]
**IncidentEvents** | **bool** | Trigger event when an incident is created. | [optional] [default to null]
**VulnerabilityEvents** | **bool** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


