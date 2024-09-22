# \EnvironmentsApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4ProjectsIdEnvironmentsEnvironmentId**](EnvironmentsApi.md#DeleteApiV4ProjectsIdEnvironmentsEnvironmentId) | **Delete** /api/v4/projects/{id}/environments/{environment_id} | Delete an environment
[**DeleteApiV4ProjectsIdEnvironmentsReviewApps**](EnvironmentsApi.md#DeleteApiV4ProjectsIdEnvironmentsReviewApps) | **Delete** /api/v4/projects/{id}/environments/review_apps | Delete multiple stopped review apps
[**GetApiV4ProjectsIdEnvironments**](EnvironmentsApi.md#GetApiV4ProjectsIdEnvironments) | **Get** /api/v4/projects/{id}/environments | List environments
[**GetApiV4ProjectsIdEnvironmentsEnvironmentId**](EnvironmentsApi.md#GetApiV4ProjectsIdEnvironmentsEnvironmentId) | **Get** /api/v4/projects/{id}/environments/{environment_id} | 
[**PostApiV4ProjectsIdEnvironments**](EnvironmentsApi.md#PostApiV4ProjectsIdEnvironments) | **Post** /api/v4/projects/{id}/environments | Create a new environment
[**PostApiV4ProjectsIdEnvironmentsEnvironmentIdStop**](EnvironmentsApi.md#PostApiV4ProjectsIdEnvironmentsEnvironmentIdStop) | **Post** /api/v4/projects/{id}/environments/{environment_id}/stop | Stop an environment
[**PostApiV4ProjectsIdEnvironmentsStopStale**](EnvironmentsApi.md#PostApiV4ProjectsIdEnvironmentsStopStale) | **Post** /api/v4/projects/{id}/environments/stop_stale | Stop stale environments
[**PutApiV4ProjectsIdEnvironmentsEnvironmentId**](EnvironmentsApi.md#PutApiV4ProjectsIdEnvironmentsEnvironmentId) | **Put** /api/v4/projects/{id}/environments/{environment_id} | Update an existing environment


# **DeleteApiV4ProjectsIdEnvironmentsEnvironmentId**
> ApiEntitiesEnvironment DeleteApiV4ProjectsIdEnvironmentsEnvironmentId(ctx, id, environmentId)
Delete an environment

It returns 204 if the environment was successfully deleted, and 404 if the environment does not exist. This feature was introduced in GitLab 8.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **environmentId** | **int32**| The ID of the environment | 

### Return type

[**ApiEntitiesEnvironment**](API_Entities_Environment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdEnvironmentsReviewApps**
> ApiEntitiesEnvironmentBasic DeleteApiV4ProjectsIdEnvironmentsReviewApps(ctx, id, optional)
Delete multiple stopped review apps

It schedules for deletion multiple environments that have already been stopped and are in the review app folder. The actual deletion is performed after 1 week from the time of execution. By default, it only deletes environments 30 days or older. You can change this default using the `before` parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
 **optional** | ***EnvironmentsApiDeleteApiV4ProjectsIdEnvironmentsReviewAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnvironmentsApiDeleteApiV4ProjectsIdEnvironmentsReviewAppsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **before** | **optional.Time**| The date before which environments can be deleted. Defaults to 30 days ago. Expected in ISO 8601 format (&#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;) | [default to {}]
 **limit** | **optional.Int32**| Maximum number of environments to delete. Defaults to 100 | [default to 100]
 **dryRun** | **optional.Bool**| Defaults to true for safety reasons. It performs a dry run where no actual deletion will be performed. Set to false to actually delete the environment | [default to true]

### Return type

[**ApiEntitiesEnvironmentBasic**](API_Entities_EnvironmentBasic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdEnvironments**
> []ApiEntitiesEnvironment GetApiV4ProjectsIdEnvironments(ctx, id, optional)
List environments

Get all environments for a given project. This feature was introduced in GitLab 8.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
 **optional** | ***EnvironmentsApiGetApiV4ProjectsIdEnvironmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnvironmentsApiGetApiV4ProjectsIdEnvironmentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **name** | **optional.String**| Return the environment with this name. Mutually exclusive with search | 
 **search** | **optional.String**| Return list of environments matching the search criteria. Mutually exclusive with name. Must be at least 3 characters. | 
 **states** | **optional.String**| List all environments that match a specific state. Accepted values: &#x60;available&#x60;, &#x60;stopping&#x60;, or &#x60;stopped&#x60;. If no state value given, returns all environments | 

### Return type

[**[]ApiEntitiesEnvironment**](API_Entities_Environment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdEnvironmentsEnvironmentId**
> ApiEntitiesEnvironment GetApiV4ProjectsIdEnvironmentsEnvironmentId(ctx, id, environmentId)


Get a specific environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **environmentId** | **int32**| The ID of the environment | 

### Return type

[**ApiEntitiesEnvironment**](API_Entities_Environment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdEnvironments**
> ApiEntitiesEnvironment PostApiV4ProjectsIdEnvironments(ctx, id, postApiV4ProjectsIdEnvironments)
Create a new environment

Creates a new environment with the given name and `external_url`. It returns `201` if the environment was successfully created, `400` for wrong parameters. This feature was introduced in GitLab 8.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **postApiV4ProjectsIdEnvironments** | [**PostApiV4ProjectsIdEnvironments**](PostApiV4ProjectsIdEnvironments.md)|  | 

### Return type

[**ApiEntitiesEnvironment**](API_Entities_Environment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdEnvironmentsEnvironmentIdStop**
> ApiEntitiesEnvironment PostApiV4ProjectsIdEnvironmentsEnvironmentIdStop(ctx, id, environmentId, postApiV4ProjectsIdEnvironmentsEnvironmentIdStop)
Stop an environment

It returns 200 if the environment was successfully stopped, and 404 if the environment does not exist.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **environmentId** | **int32**| The ID of the environment | 
  **postApiV4ProjectsIdEnvironmentsEnvironmentIdStop** | [**PostApiV4ProjectsIdEnvironmentsEnvironmentIdStop**](PostApiV4ProjectsIdEnvironmentsEnvironmentIdStop.md)|  | 

### Return type

[**ApiEntitiesEnvironment**](API_Entities_Environment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdEnvironmentsStopStale**
> PostApiV4ProjectsIdEnvironmentsStopStale(ctx, id, postApiV4ProjectsIdEnvironmentsStopStale)
Stop stale environments

It returns `200` if stale environment check was scheduled successfully

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **postApiV4ProjectsIdEnvironmentsStopStale** | [**PostApiV4ProjectsIdEnvironmentsStopStale**](PostApiV4ProjectsIdEnvironmentsStopStale.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdEnvironmentsEnvironmentId**
> ApiEntitiesEnvironment PutApiV4ProjectsIdEnvironmentsEnvironmentId(ctx, id, environmentId, putApiV4ProjectsIdEnvironmentsEnvironmentId)
Update an existing environment

Updates an existing environment name and/or `external_url`. It returns `200` if the environment was successfully updated. In case of an error, a status code `400` is returned. This feature was introduced in GitLab 8.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **environmentId** | **int32**| The ID of the environment | 
  **putApiV4ProjectsIdEnvironmentsEnvironmentId** | [**PutApiV4ProjectsIdEnvironmentsEnvironmentId**](PutApiV4ProjectsIdEnvironmentsEnvironmentId.md)|  | 

### Return type

[**ApiEntitiesEnvironment**](API_Entities_Environment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

