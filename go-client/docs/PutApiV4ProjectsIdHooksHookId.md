# PutApiV4ProjectsIdHooksHookId

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The URL to send the request to | [optional] [default to null]
**Name** | **string** | Name of the hook | [optional] [default to null]
**Description** | **string** | Description of the hook | [optional] [default to null]
**PushEvents** | **bool** | Trigger hook on push events | [optional] [default to null]
**IssuesEvents** | **bool** | Trigger hook on issues events | [optional] [default to null]
**ConfidentialIssuesEvents** | **bool** | Trigger hook on confidential issues events | [optional] [default to null]
**MergeRequestsEvents** | **bool** | Trigger hook on merge request events | [optional] [default to null]
**TagPushEvents** | **bool** | Trigger hook on tag push events | [optional] [default to null]
**NoteEvents** | **bool** | Trigger hook on note (comment) events | [optional] [default to null]
**ConfidentialNoteEvents** | **bool** | Trigger hook on confidential note (comment) events | [optional] [default to null]
**JobEvents** | **bool** | Trigger hook on job events | [optional] [default to null]
**PipelineEvents** | **bool** | Trigger hook on pipeline events | [optional] [default to null]
**WikiPageEvents** | **bool** | Trigger hook on wiki events | [optional] [default to null]
**DeploymentEvents** | **bool** | Trigger hook on deployment events | [optional] [default to null]
**ReleasesEvents** | **bool** | Trigger hook on release events | [optional] [default to null]
**EmojiEvents** | **bool** | Trigger hook on emoji events | [optional] [default to null]
**ResourceAccessTokenEvents** | **bool** | Trigger hook on project access token expiry events | [optional] [default to null]
**EnableSslVerification** | **bool** | Do SSL verification when triggering the hook | [optional] [default to null]
**Token** | **string** | Secret token to validate received payloads; this will not be returned in the response | [optional] [default to null]
**PushEventsBranchFilter** | **string** | Trigger hook on specified branch only | [optional] [default to null]
**CustomWebhookTemplate** | **string** | Custom template for the request payload | [optional] [default to null]
**UrlVariables** | [**[]PostApiV4ProjectsIdHooksUrlVariables**](postApiV4ProjectsIdHooks_url_variables.md) | URL variables for interpolation | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


