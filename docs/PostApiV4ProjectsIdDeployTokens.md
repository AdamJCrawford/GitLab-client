# PostApiV4ProjectsIdDeployTokens

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | New deploy token&#39;s name | [default to null]
**Scopes** | **[]string** | Indicates the deploy token scopes. Must be at least one of &#x60;read_repository&#x60;, &#x60;read_registry&#x60;, &#x60;write_registry&#x60;, &#x60;read_package_registry&#x60;, or &#x60;write_package_registry&#x60;. | [default to null]
**ExpiresAt** | [**time.Time**](time.Time.md) | Expiration date for the deploy token. Does not expire if no value is provided. Expected in ISO 8601 format (&#x60;2019-03-15T08:00:00Z&#x60;). | [optional] [default to null]
**Username** | **string** | Username for deploy token. Default is &#x60;gitlab+deploy-token-{n}&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


