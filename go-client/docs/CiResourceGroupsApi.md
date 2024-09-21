# \CiResourceGroupsApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4ProjectsIdResourceGroups**](CiResourceGroupsApi.md#GetApiV4ProjectsIdResourceGroups) | **Get** /api/v4/projects/{id}/resource_groups | 
[**GetApiV4ProjectsIdResourceGroupsKey**](CiResourceGroupsApi.md#GetApiV4ProjectsIdResourceGroupsKey) | **Get** /api/v4/projects/{id}/resource_groups/{key} | 
[**GetApiV4ProjectsIdResourceGroupsKeyUpcomingJobs**](CiResourceGroupsApi.md#GetApiV4ProjectsIdResourceGroupsKeyUpcomingJobs) | **Get** /api/v4/projects/{id}/resource_groups/{key}/upcoming_jobs | 
[**PutApiV4ProjectsIdResourceGroupsKey**](CiResourceGroupsApi.md#PutApiV4ProjectsIdResourceGroupsKey) | **Put** /api/v4/projects/{id}/resource_groups/{key} | Edit an existing resource group


# **GetApiV4ProjectsIdResourceGroups**
> []ApiEntitiesCiResourceGroup GetApiV4ProjectsIdResourceGroups(ctx, id, optional)


Get all resource groups for a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
 **optional** | ***CiResourceGroupsApiGetApiV4ProjectsIdResourceGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CiResourceGroupsApiGetApiV4ProjectsIdResourceGroupsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesCiResourceGroup**](API_Entities_Ci_ResourceGroup.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdResourceGroupsKey**
> ApiEntitiesCiResourceGroup GetApiV4ProjectsIdResourceGroupsKey(ctx, id, key)


Get a specific resource group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **key** | **string**| The key of the resource group | 

### Return type

[**ApiEntitiesCiResourceGroup**](API_Entities_Ci_ResourceGroup.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdResourceGroupsKeyUpcomingJobs**
> []ApiEntitiesCiJobBasic GetApiV4ProjectsIdResourceGroupsKeyUpcomingJobs(ctx, id, key, optional)


List upcoming jobs for a specific resource group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **key** | **string**| The key of the resource group | 
 **optional** | ***CiResourceGroupsApiGetApiV4ProjectsIdResourceGroupsKeyUpcomingJobsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CiResourceGroupsApiGetApiV4ProjectsIdResourceGroupsKeyUpcomingJobsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesCiJobBasic**](API_Entities_Ci_JobBasic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdResourceGroupsKey**
> ApiEntitiesCiResourceGroup PutApiV4ProjectsIdResourceGroupsKey(ctx, id, key, putApiV4ProjectsIdResourceGroupsKey)
Edit an existing resource group

Updates an existing resource group's properties.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **key** | **string**| The key of the resource group | 
  **putApiV4ProjectsIdResourceGroupsKey** | [**PutApiV4ProjectsIdResourceGroupsKey**](PutApiV4ProjectsIdResourceGroupsKey.md)|  | 

### Return type

[**ApiEntitiesCiResourceGroup**](API_Entities_Ci_ResourceGroup.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

