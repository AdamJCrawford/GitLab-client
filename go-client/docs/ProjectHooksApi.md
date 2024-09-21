# \ProjectHooksApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4ProjectsIdHooksHookId**](ProjectHooksApi.md#DeleteApiV4ProjectsIdHooksHookId) | **Delete** /api/v4/projects/{id}/hooks/{hook_id} | Delete a project hook
[**GetApiV4ProjectsIdHooks**](ProjectHooksApi.md#GetApiV4ProjectsIdHooks) | **Get** /api/v4/projects/{id}/hooks | List project hooks
[**GetApiV4ProjectsIdHooksHookId**](ProjectHooksApi.md#GetApiV4ProjectsIdHooksHookId) | **Get** /api/v4/projects/{id}/hooks/{hook_id} | Get project hook
[**PostApiV4ProjectsIdHooks**](ProjectHooksApi.md#PostApiV4ProjectsIdHooks) | **Post** /api/v4/projects/{id}/hooks | Add project hook
[**PutApiV4ProjectsIdHooksHookId**](ProjectHooksApi.md#PutApiV4ProjectsIdHooksHookId) | **Put** /api/v4/projects/{id}/hooks/{hook_id} | Edit project hook


# **DeleteApiV4ProjectsIdHooksHookId**
> ApiEntitiesProjectHook DeleteApiV4ProjectsIdHooksHookId(ctx, id, hookId)
Delete a project hook

Removes a hook from a project. This is an idempotent method and can be called multiple times. Either the hook is available or not.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **hookId** | **int32**| The ID of the project hook | 

### Return type

[**ApiEntitiesProjectHook**](API_Entities_ProjectHook.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdHooks**
> []ApiEntitiesProjectHook GetApiV4ProjectsIdHooks(ctx, id, optional)
List project hooks

Get a list of project hooks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
 **optional** | ***ProjectHooksApiGetApiV4ProjectsIdHooksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectHooksApiGetApiV4ProjectsIdHooksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesProjectHook**](API_Entities_ProjectHook.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdHooksHookId**
> ApiEntitiesProjectHook GetApiV4ProjectsIdHooksHookId(ctx, id, hookId)
Get project hook

Get a specific hook for a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **hookId** | **int32**| The ID of a project hook | 

### Return type

[**ApiEntitiesProjectHook**](API_Entities_ProjectHook.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdHooks**
> ApiEntitiesProjectHook PostApiV4ProjectsIdHooks(ctx, id, postApiV4ProjectsIdHooks)
Add project hook

Adds a hook to a specified project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **postApiV4ProjectsIdHooks** | [**PostApiV4ProjectsIdHooks**](PostApiV4ProjectsIdHooks.md)|  | 

### Return type

[**ApiEntitiesProjectHook**](API_Entities_ProjectHook.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdHooksHookId**
> ApiEntitiesProjectHook PutApiV4ProjectsIdHooksHookId(ctx, id, hookId, putApiV4ProjectsIdHooksHookId)
Edit project hook

Edits a hook for a specified project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **hookId** | **int32**| The ID of the project hook | 
  **putApiV4ProjectsIdHooksHookId** | [**PutApiV4ProjectsIdHooksHookId**](PutApiV4ProjectsIdHooksHookId.md)|  | 

### Return type

[**ApiEntitiesProjectHook**](API_Entities_ProjectHook.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

