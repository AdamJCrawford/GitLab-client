# \FeatureFlagsUserListsApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4ProjectsIdFeatureFlagsUserListsIid**](FeatureFlagsUserListsApi.md#DeleteApiV4ProjectsIdFeatureFlagsUserListsIid) | **Delete** /api/v4/projects/{id}/feature_flags_user_lists/{iid} | Delete feature flag user list
[**GetApiV4ProjectsIdFeatureFlagsUserLists**](FeatureFlagsUserListsApi.md#GetApiV4ProjectsIdFeatureFlagsUserLists) | **Get** /api/v4/projects/{id}/feature_flags_user_lists | List all feature flag user lists for a project
[**GetApiV4ProjectsIdFeatureFlagsUserListsIid**](FeatureFlagsUserListsApi.md#GetApiV4ProjectsIdFeatureFlagsUserListsIid) | **Get** /api/v4/projects/{id}/feature_flags_user_lists/{iid} | Get a feature flag user list
[**PostApiV4ProjectsIdFeatureFlagsUserLists**](FeatureFlagsUserListsApi.md#PostApiV4ProjectsIdFeatureFlagsUserLists) | **Post** /api/v4/projects/{id}/feature_flags_user_lists | Create a feature flag user list
[**PutApiV4ProjectsIdFeatureFlagsUserListsIid**](FeatureFlagsUserListsApi.md#PutApiV4ProjectsIdFeatureFlagsUserListsIid) | **Put** /api/v4/projects/{id}/feature_flags_user_lists/{iid} | Update a feature flag user list


# **DeleteApiV4ProjectsIdFeatureFlagsUserListsIid**
> DeleteApiV4ProjectsIdFeatureFlagsUserListsIid(ctx, id, iid)
Delete feature flag user list

Deletes a feature flag user list. This feature was introduced in GitLab 12.10.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **iid** | **string**| The internal ID of the project&#39;s feature flag user list | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdFeatureFlagsUserLists**
> []ApiEntitiesFeatureFlagUserList GetApiV4ProjectsIdFeatureFlagsUserLists(ctx, id, optional)
List all feature flag user lists for a project

Gets all feature flag user lists for the requested project. This feature was introduced in GitLab 12.10.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
 **optional** | ***FeatureFlagsUserListsApiGetApiV4ProjectsIdFeatureFlagsUserListsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeatureFlagsUserListsApiGetApiV4ProjectsIdFeatureFlagsUserListsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| Return user lists matching the search criteria | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesFeatureFlagUserList**](API_Entities_FeatureFlag_UserList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdFeatureFlagsUserListsIid**
> ApiEntitiesFeatureFlagUserList GetApiV4ProjectsIdFeatureFlagsUserListsIid(ctx, id, iid)
Get a feature flag user list

Gets a feature flag user list. This feature was introduced in GitLab 12.10.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **iid** | **string**| The internal ID of the project&#39;s feature flag user list | 

### Return type

[**ApiEntitiesFeatureFlagUserList**](API_Entities_FeatureFlag_UserList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdFeatureFlagsUserLists**
> ApiEntitiesFeatureFlagUserList PostApiV4ProjectsIdFeatureFlagsUserLists(ctx, id, postApiV4ProjectsIdFeatureFlagsUserLists)
Create a feature flag user list

Creates a feature flag user list. This feature was introduced in GitLab 12.10.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **postApiV4ProjectsIdFeatureFlagsUserLists** | [**PostApiV4ProjectsIdFeatureFlagsUserLists**](PostApiV4ProjectsIdFeatureFlagsUserLists.md)|  | 

### Return type

[**ApiEntitiesFeatureFlagUserList**](API_Entities_FeatureFlag_UserList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdFeatureFlagsUserListsIid**
> ApiEntitiesFeatureFlagUserList PutApiV4ProjectsIdFeatureFlagsUserListsIid(ctx, id, iid, putApiV4ProjectsIdFeatureFlagsUserListsIid)
Update a feature flag user list

Updates a feature flag user list. This feature was introduced in GitLab 12.10.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **iid** | **string**| The internal ID of the project&#39;s feature flag user list | 
  **putApiV4ProjectsIdFeatureFlagsUserListsIid** | [**PutApiV4ProjectsIdFeatureFlagsUserListsIid**](PutApiV4ProjectsIdFeatureFlagsUserListsIid.md)|  | 

### Return type

[**ApiEntitiesFeatureFlagUserList**](API_Entities_FeatureFlag_UserList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

