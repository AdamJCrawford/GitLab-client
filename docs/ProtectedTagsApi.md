# \ProtectedTagsApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4ProjectsIdProtectedTagsName**](ProtectedTagsApi.md#DeleteApiV4ProjectsIdProtectedTagsName) | **Delete** /api/v4/projects/{id}/protected_tags/{name} | Unprotect a single tag
[**GetApiV4ProjectsIdProtectedTags**](ProtectedTagsApi.md#GetApiV4ProjectsIdProtectedTags) | **Get** /api/v4/projects/{id}/protected_tags | Get a project&#39;s protected tags
[**GetApiV4ProjectsIdProtectedTagsName**](ProtectedTagsApi.md#GetApiV4ProjectsIdProtectedTagsName) | **Get** /api/v4/projects/{id}/protected_tags/{name} | Get a single protected tag
[**PostApiV4ProjectsIdProtectedTags**](ProtectedTagsApi.md#PostApiV4ProjectsIdProtectedTags) | **Post** /api/v4/projects/{id}/protected_tags | Protect a single tag or wildcard


# **DeleteApiV4ProjectsIdProtectedTagsName**
> DeleteApiV4ProjectsIdProtectedTagsName(ctx, id, name)
Unprotect a single tag

This feature was introduced in GitLab 11.3.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **name** | **string**| The name of the protected tag | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdProtectedTags**
> []ApiEntitiesProtectedTag GetApiV4ProjectsIdProtectedTags(ctx, id, optional)
Get a project's protected tags

This feature was introduced in GitLab 11.3.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
 **optional** | ***ProtectedTagsApiGetApiV4ProjectsIdProtectedTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProtectedTagsApiGetApiV4ProjectsIdProtectedTagsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesProtectedTag**](API_Entities_ProtectedTag.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdProtectedTagsName**
> ApiEntitiesProtectedTag GetApiV4ProjectsIdProtectedTagsName(ctx, id, name)
Get a single protected tag

This feature was introduced in GitLab 11.3.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **name** | **string**| The name of the tag or wildcard | 

### Return type

[**ApiEntitiesProtectedTag**](API_Entities_ProtectedTag.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdProtectedTags**
> ApiEntitiesProtectedTag PostApiV4ProjectsIdProtectedTags(ctx, id, postApiV4ProjectsIdProtectedTags)
Protect a single tag or wildcard

This feature was introduced in GitLab 11.3.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **postApiV4ProjectsIdProtectedTags** | [**PostApiV4ProjectsIdProtectedTags**](PostApiV4ProjectsIdProtectedTags.md)|  | 

### Return type

[**ApiEntitiesProtectedTag**](API_Entities_ProtectedTag.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

