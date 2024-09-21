# PostApiV4Runners

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | Registration token | [default to null]
**Description** | **string** | Description of the runner | [optional] [default to null]
**MaintainerNote** | **string** | Deprecated: see &#x60;maintenance_note&#x60; | [optional] [default to null]
**MaintenanceNote** | **string** | Free-form maintenance notes for the runner (1024 characters) | [optional] [default to null]
**Info** | [***PostApiV4RunnersInfo**](postApiV4Runners_info.md) |  | [optional] [default to null]
**Active** | **bool** | Deprecated: Use &#x60;paused&#x60; instead. Specifies if the runner is allowed to receive new jobs | [optional] [default to null]
**Paused** | **bool** | Specifies if the runner should ignore new jobs | [optional] [default to null]
**Locked** | **bool** | Specifies if the runner should be locked for the current project | [optional] [default to null]
**AccessLevel** | **string** | The access level of the runner | [optional] [default to null]
**RunUntagged** | **bool** | Specifies if the runner should handle untagged jobs | [optional] [default to null]
**TagList** | **[]string** | A list of runner tags | [optional] [default to null]
**MaximumTimeout** | **int32** | Maximum timeout that limits the amount of time (in seconds) that runners can run jobs | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


