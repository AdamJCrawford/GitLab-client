# \FeatureFlagsApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4ProjectsIdFeatureFlagsFeatureFlagName**](FeatureFlagsApi.md#DeleteApiV4ProjectsIdFeatureFlagsFeatureFlagName) | **Delete** /api/v4/projects/{id}/feature_flags/{feature_flag_name} | Delete a feature flag
[**GetApiV4FeatureFlagsUnleashProjectId**](FeatureFlagsApi.md#GetApiV4FeatureFlagsUnleashProjectId) | **Get** /api/v4/feature_flags/unleash/{project_id} | 
[**GetApiV4ProjectsIdFeatureFlags**](FeatureFlagsApi.md#GetApiV4ProjectsIdFeatureFlags) | **Get** /api/v4/projects/{id}/feature_flags | List feature flags for a project
[**GetApiV4ProjectsIdFeatureFlagsFeatureFlagName**](FeatureFlagsApi.md#GetApiV4ProjectsIdFeatureFlagsFeatureFlagName) | **Get** /api/v4/projects/{id}/feature_flags/{feature_flag_name} | Get a single feature flag
[**PostApiV4FeatureFlagsUnleashProjectIdClientMetrics**](FeatureFlagsApi.md#PostApiV4FeatureFlagsUnleashProjectIdClientMetrics) | **Post** /api/v4/feature_flags/unleash/{project_id}/client/metrics | 
[**PostApiV4FeatureFlagsUnleashProjectIdClientRegister**](FeatureFlagsApi.md#PostApiV4FeatureFlagsUnleashProjectIdClientRegister) | **Post** /api/v4/feature_flags/unleash/{project_id}/client/register | 
[**PostApiV4ProjectsIdFeatureFlags**](FeatureFlagsApi.md#PostApiV4ProjectsIdFeatureFlags) | **Post** /api/v4/projects/{id}/feature_flags | Create a new feature flag
[**PutApiV4ProjectsIdFeatureFlagsFeatureFlagName**](FeatureFlagsApi.md#PutApiV4ProjectsIdFeatureFlagsFeatureFlagName) | **Put** /api/v4/projects/{id}/feature_flags/{feature_flag_name} | Update a feature flag


# **DeleteApiV4ProjectsIdFeatureFlagsFeatureFlagName**
> ApiEntitiesFeatureFlag DeleteApiV4ProjectsIdFeatureFlagsFeatureFlagName(ctx, id, featureFlagName)
Delete a feature flag

Deletes a feature flag. This feature was introduced in GitLab 12.5.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **featureFlagName** | **string**| The name of the feature flag | 

### Return type

[**ApiEntitiesFeatureFlag**](API_Entities_FeatureFlag.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4FeatureFlagsUnleashProjectId**
> GetApiV4FeatureFlagsUnleashProjectId(ctx, projectId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The ID of a project | 
 **optional** | ***FeatureFlagsApiGetApiV4FeatureFlagsUnleashProjectIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeatureFlagsApiGetApiV4FeatureFlagsUnleashProjectIdOpts struct

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

# **GetApiV4ProjectsIdFeatureFlags**
> []ApiEntitiesFeatureFlag GetApiV4ProjectsIdFeatureFlags(ctx, id, optional)
List feature flags for a project

Gets all feature flags of the requested project. This feature was introduced in GitLab 12.5.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
 **optional** | ***FeatureFlagsApiGetApiV4ProjectsIdFeatureFlagsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeatureFlagsApiGetApiV4ProjectsIdFeatureFlagsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scope** | **optional.String**| The scope of feature flags, one of: &#x60;enabled&#x60;, &#x60;disabled&#x60; | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesFeatureFlag**](API_Entities_FeatureFlag.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdFeatureFlagsFeatureFlagName**
> ApiEntitiesFeatureFlag GetApiV4ProjectsIdFeatureFlagsFeatureFlagName(ctx, id, featureFlagName)
Get a single feature flag

Gets a single feature flag. This feature was introduced in GitLab 12.5.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **featureFlagName** | **string**| The name of the feature flag | 

### Return type

[**ApiEntitiesFeatureFlag**](API_Entities_FeatureFlag.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4FeatureFlagsUnleashProjectIdClientMetrics**
> PostApiV4FeatureFlagsUnleashProjectIdClientMetrics(ctx, projectId, postApiV4FeatureFlagsUnleashProjectIdClientMetrics)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The ID of a project | 
  **postApiV4FeatureFlagsUnleashProjectIdClientMetrics** | [**PostApiV4FeatureFlagsUnleashProjectIdClientMetrics**](PostApiV4FeatureFlagsUnleashProjectIdClientMetrics.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4FeatureFlagsUnleashProjectIdClientRegister**
> PostApiV4FeatureFlagsUnleashProjectIdClientRegister(ctx, projectId, postApiV4FeatureFlagsUnleashProjectIdClientRegister)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The ID of a project | 
  **postApiV4FeatureFlagsUnleashProjectIdClientRegister** | [**PostApiV4FeatureFlagsUnleashProjectIdClientRegister**](PostApiV4FeatureFlagsUnleashProjectIdClientRegister.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdFeatureFlags**
> ApiEntitiesFeatureFlag PostApiV4ProjectsIdFeatureFlags(ctx, id, postApiV4ProjectsIdFeatureFlags)
Create a new feature flag

Creates a new feature flag. This feature was introduced in GitLab 12.5.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **postApiV4ProjectsIdFeatureFlags** | [**PostApiV4ProjectsIdFeatureFlags**](PostApiV4ProjectsIdFeatureFlags.md)|  | 

### Return type

[**ApiEntitiesFeatureFlag**](API_Entities_FeatureFlag.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdFeatureFlagsFeatureFlagName**
> ApiEntitiesFeatureFlag PutApiV4ProjectsIdFeatureFlagsFeatureFlagName(ctx, id, featureFlagName, putApiV4ProjectsIdFeatureFlagsFeatureFlagName)
Update a feature flag

Updates a feature flag. This feature was introduced in GitLab 13.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **featureFlagName** | **string**| The name of the feature flag | 
  **putApiV4ProjectsIdFeatureFlagsFeatureFlagName** | [**PutApiV4ProjectsIdFeatureFlagsFeatureFlagName**](PutApiV4ProjectsIdFeatureFlagsFeatureFlagName.md)|  | 

### Return type

[**ApiEntitiesFeatureFlag**](API_Entities_FeatureFlag.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

