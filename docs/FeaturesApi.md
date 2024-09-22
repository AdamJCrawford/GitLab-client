# \FeaturesApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4FeaturesName**](FeaturesApi.md#DeleteApiV4FeaturesName) | **Delete** /api/v4/features/{name} | Delete a feature
[**GetApiV4Features**](FeaturesApi.md#GetApiV4Features) | **Get** /api/v4/features | List all features
[**GetApiV4FeaturesDefinitions**](FeaturesApi.md#GetApiV4FeaturesDefinitions) | **Get** /api/v4/features/definitions | List all feature definitions
[**PostApiV4FeaturesName**](FeaturesApi.md#PostApiV4FeaturesName) | **Post** /api/v4/features/{name} | Set or create a feature


# **DeleteApiV4FeaturesName**
> DeleteApiV4FeaturesName(ctx, name)
Delete a feature

Removes a feature gate. Response is equal when the gate exists, or doesn't.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4Features**
> []ApiEntitiesFeature GetApiV4Features(ctx, )
List all features

Get a list of all persisted features, with its gate values.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ApiEntitiesFeature**](API_Entities_Feature.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4FeaturesDefinitions**
> []ApiEntitiesFeatureDefinition GetApiV4FeaturesDefinitions(ctx, )
List all feature definitions

Get a list of all feature definitions.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ApiEntitiesFeatureDefinition**](API_Entities_Feature_Definition.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4FeaturesName**
> ApiEntitiesFeature PostApiV4FeaturesName(ctx, name, postApiV4FeaturesName)
Set or create a feature

Set a feature's gate value. If a feature with the given name doesn't exist yet, it's created. The value can be a boolean, or an integer to indicate percentage of time.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **int32**|  | 
  **postApiV4FeaturesName** | [**PostApiV4FeaturesName**](PostApiV4FeaturesName.md)|  | 

### Return type

[**ApiEntitiesFeature**](API_Entities_Feature.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

