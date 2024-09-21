# \UnleashApiApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4FeatureFlagsUnleashProjectIdClientFeatures**](UnleashApiApi.md#GetApiV4FeatureFlagsUnleashProjectIdClientFeatures) | **Get** /api/v4/feature_flags/unleash/{project_id}/client/features | 
[**GetApiV4FeatureFlagsUnleashProjectIdFeatures**](UnleashApiApi.md#GetApiV4FeatureFlagsUnleashProjectIdFeatures) | **Get** /api/v4/feature_flags/unleash/{project_id}/features | 


# **GetApiV4FeatureFlagsUnleashProjectIdClientFeatures**
> GetApiV4FeatureFlagsUnleashProjectIdClientFeatures(ctx, projectId, optional)


Get a list of features

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The ID of a project | 
 **optional** | ***UnleashApiApiGetApiV4FeatureFlagsUnleashProjectIdClientFeaturesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UnleashApiApiGetApiV4FeatureFlagsUnleashProjectIdClientFeaturesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instanceId** | **optional.String**| The instance ID of Unleash Client | 
 **appName** | **optional.String**| The application name of Unleash Client | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4FeatureFlagsUnleashProjectIdFeatures**
> GetApiV4FeatureFlagsUnleashProjectIdFeatures(ctx, projectId, optional)


Get a list of features (deprecated, v2 client support)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The ID of a project | 
 **optional** | ***UnleashApiApiGetApiV4FeatureFlagsUnleashProjectIdFeaturesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UnleashApiApiGetApiV4FeatureFlagsUnleashProjectIdFeaturesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instanceId** | **optional.String**| The instance ID of Unleash Client | 
 **appName** | **optional.String**| The application name of Unleash Client | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

