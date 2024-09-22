# \ProjectSnippetsApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4ProjectsIdSnippetsSnippetId**](ProjectSnippetsApi.md#DeleteApiV4ProjectsIdSnippetsSnippetId) | **Delete** /api/v4/projects/{id}/snippets/{snippet_id} | 
[**GetApiV4ProjectsIdSnippets**](ProjectSnippetsApi.md#GetApiV4ProjectsIdSnippets) | **Get** /api/v4/projects/{id}/snippets | 
[**GetApiV4ProjectsIdSnippetsSnippetId**](ProjectSnippetsApi.md#GetApiV4ProjectsIdSnippetsSnippetId) | **Get** /api/v4/projects/{id}/snippets/{snippet_id} | 
[**GetApiV4ProjectsIdSnippetsSnippetIdFilesRefFilePathRaw**](ProjectSnippetsApi.md#GetApiV4ProjectsIdSnippetsSnippetIdFilesRefFilePathRaw) | **Get** /api/v4/projects/{id}/snippets/{snippet_id}/files/{ref}/{file_path}/raw | 
[**GetApiV4ProjectsIdSnippetsSnippetIdRaw**](ProjectSnippetsApi.md#GetApiV4ProjectsIdSnippetsSnippetIdRaw) | **Get** /api/v4/projects/{id}/snippets/{snippet_id}/raw | 
[**GetApiV4ProjectsIdSnippetsSnippetIdUserAgentDetail**](ProjectSnippetsApi.md#GetApiV4ProjectsIdSnippetsSnippetIdUserAgentDetail) | **Get** /api/v4/projects/{id}/snippets/{snippet_id}/user_agent_detail | 
[**PostApiV4ProjectsIdSnippets**](ProjectSnippetsApi.md#PostApiV4ProjectsIdSnippets) | **Post** /api/v4/projects/{id}/snippets | 
[**PutApiV4ProjectsIdSnippetsSnippetId**](ProjectSnippetsApi.md#PutApiV4ProjectsIdSnippetsSnippetId) | **Put** /api/v4/projects/{id}/snippets/{snippet_id} | 


# **DeleteApiV4ProjectsIdSnippetsSnippetId**
> DeleteApiV4ProjectsIdSnippetsSnippetId(ctx, id, snippetId)


Delete a project snippet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **snippetId** | **int32**| The ID of a project snippet | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdSnippets**
> []ApiEntitiesProjectSnippet GetApiV4ProjectsIdSnippets(ctx, id, optional)


Get all project snippets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
 **optional** | ***ProjectSnippetsApiGetApiV4ProjectsIdSnippetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectSnippetsApiGetApiV4ProjectsIdSnippetsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesProjectSnippet**](API_Entities_ProjectSnippet.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdSnippetsSnippetId**
> ApiEntitiesProjectSnippet GetApiV4ProjectsIdSnippetsSnippetId(ctx, id, snippetId)


Get a single project snippet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **snippetId** | **int32**| The ID of a project snippet | 

### Return type

[**ApiEntitiesProjectSnippet**](API_Entities_ProjectSnippet.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdSnippetsSnippetIdFilesRefFilePathRaw**
> ApiEntitiesProjectSnippet GetApiV4ProjectsIdSnippetsSnippetIdFilesRefFilePathRaw(ctx, id, filePath, ref, snippetId)


Get raw project snippet file contents from the repository

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **filePath** | **string**| The url encoded path to the file, e.g. lib%2Fclass%2Erb | 
  **ref** | **string**| The name of branch, tag or commit | 
  **snippetId** | **int32**|  | 

### Return type

[**ApiEntitiesProjectSnippet**](API_Entities_ProjectSnippet.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdSnippetsSnippetIdRaw**
> ApiEntitiesProjectSnippet GetApiV4ProjectsIdSnippetsSnippetIdRaw(ctx, id, snippetId)


Get a raw project snippet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **snippetId** | **int32**| The ID of a project snippet | 

### Return type

[**ApiEntitiesProjectSnippet**](API_Entities_ProjectSnippet.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdSnippetsSnippetIdUserAgentDetail**
> ApiEntitiesUserAgentDetail GetApiV4ProjectsIdSnippetsSnippetIdUserAgentDetail(ctx, id, snippetId)


Get the user agent details for a project snippet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **snippetId** | **int32**| The ID of a project snippet | 

### Return type

[**ApiEntitiesUserAgentDetail**](API_Entities_UserAgentDetail.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdSnippets**
> ApiEntitiesProjectSnippet PostApiV4ProjectsIdSnippets(ctx, id, postApiV4ProjectsIdSnippets)


Create a new project snippet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **postApiV4ProjectsIdSnippets** | [**PostApiV4ProjectsIdSnippets**](PostApiV4ProjectsIdSnippets.md)|  | 

### Return type

[**ApiEntitiesProjectSnippet**](API_Entities_ProjectSnippet.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdSnippetsSnippetId**
> ApiEntitiesProjectSnippet PutApiV4ProjectsIdSnippetsSnippetId(ctx, id, snippetId, putApiV4ProjectsIdSnippetsSnippetId)


Update an existing project snippet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **snippetId** | **int32**| The ID of a project snippet | 
  **putApiV4ProjectsIdSnippetsSnippetId** | [**PutApiV4ProjectsIdSnippetsSnippetId**](PutApiV4ProjectsIdSnippetsSnippetId.md)|  | 

### Return type

[**ApiEntitiesProjectSnippet**](API_Entities_ProjectSnippet.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

