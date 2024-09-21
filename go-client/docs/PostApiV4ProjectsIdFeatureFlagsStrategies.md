# PostApiV4ProjectsIdFeatureFlagsStrategies

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The strategy name. Can be &#x60;default&#x60;, &#x60;gradualRolloutUserId&#x60;, &#x60;userWithId&#x60;, or &#x60;gitlabUserList&#x60;. In GitLab 13.5 and later, can be &#x60;flexibleRollout&#x60; | [default to null]
**Parameters** | **string** | The strategy parameters as a JSON-formatted string e.g. &#x60;{\&quot;userIds\&quot;:\&quot;user1\&quot;}&#x60; | [optional] [default to null]
**UserListId** | **int32** | The ID of the feature flag user list. If strategy is &#x60;gitlabUserList&#x60;. | [optional] [default to null]
**Scopes** | [**[]PostApiV4ProjectsIdFeatureFlagsScopes**](postApiV4ProjectsIdFeatureFlags_scopes.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


