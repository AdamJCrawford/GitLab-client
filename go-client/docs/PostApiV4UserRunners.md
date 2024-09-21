# PostApiV4UserRunners

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunnerType** | **string** | Specifies the scope of the runner | [default to null]
**GroupId** | **int32** | The ID of the group that the runner is created in | [default to null]
**ProjectId** | **int32** | The ID of the project that the runner is created in | [default to null]
**Description** | **string** | Description of the runner | [optional] [default to null]
**MaintenanceNote** | **string** | Free-form maintenance notes for the runner (1024 characters) | [optional] [default to null]
**Paused** | **bool** | Specifies if the runner should ignore new jobs (defaults to false) | [optional] [default to null]
**Locked** | **bool** | Specifies if the runner should be locked for the current project (defaults to false) | [optional] [default to null]
**AccessLevel** | **string** | The access level of the runner | [optional] [default to null]
**RunUntagged** | **bool** | Specifies if the runner should handle untagged jobs  (defaults to true) | [optional] [default to null]
**TagList** | **[]string** | A list of runner tags | [optional] [default to null]
**MaximumTimeout** | **int32** | Maximum timeout that limits the amount of time (in seconds) that runners can run jobs | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


