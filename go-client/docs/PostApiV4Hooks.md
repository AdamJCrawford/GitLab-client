# PostApiV4Hooks

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The URL to send the request to | [default to null]
**Name** | **string** | Name of the hook | [optional] [default to null]
**Description** | **string** | Description of the hook | [optional] [default to null]
**Token** | **string** | Secret token to validate received payloads; this isn&#39;t returned in the response | [optional] [default to null]
**PushEvents** | **bool** | When true, the hook fires on push events | [optional] [default to null]
**TagPushEvents** | **bool** | When true, the hook fires on new tags being pushed | [optional] [default to null]
**MergeRequestsEvents** | **bool** | Trigger hook on merge requests events | [optional] [default to null]
**RepositoryUpdateEvents** | **bool** | Trigger hook on repository update events | [optional] [default to null]
**EnableSslVerification** | **bool** | Do SSL verification when triggering the hook | [optional] [default to null]
**UrlVariables** | [**[]PostApiV4ProjectsIdHooksUrlVariables**](postApiV4ProjectsIdHooks_url_variables.md) | URL variables for interpolation | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


