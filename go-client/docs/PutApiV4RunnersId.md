# PutApiV4RunnersId

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The description of the runner | [optional] [default to null]
**Active** | **bool** | Deprecated: Use &#x60;paused&#x60; instead. Flag indicating whether the runner is allowed to receive jobs | [optional] [default to null]
**Paused** | **bool** | Specifies if the runner should ignore new jobs | [optional] [default to null]
**TagList** | **[]string** | The list of tags for a runner | [optional] [default to null]
**RunUntagged** | **bool** | Specifies if the runner can execute untagged jobs | [optional] [default to null]
**Locked** | **bool** | Specifies if the runner is locked | [optional] [default to null]
**AccessLevel** | **string** | The access level of the runner | [optional] [default to null]
**MaximumTimeout** | **int32** | Maximum timeout that limits the amount of time (in seconds) that runners can run jobs | [optional] [default to null]
**MaintenanceNote** | **string** | Free-form maintenance notes for the runner (1024 characters) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


