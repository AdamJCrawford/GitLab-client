# \SystemHooksApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4HooksHookId**](SystemHooksApi.md#DeleteApiV4HooksHookId) | **Delete** /api/v4/hooks/{hook_id} | Delete system hook
[**GetApiV4Hooks**](SystemHooksApi.md#GetApiV4Hooks) | **Get** /api/v4/hooks | List system hooks
[**GetApiV4HooksHookId**](SystemHooksApi.md#GetApiV4HooksHookId) | **Get** /api/v4/hooks/{hook_id} | Get system hook
[**PostApiV4Hooks**](SystemHooksApi.md#PostApiV4Hooks) | **Post** /api/v4/hooks | Add new system hook
[**PutApiV4HooksHookId**](SystemHooksApi.md#PutApiV4HooksHookId) | **Put** /api/v4/hooks/{hook_id} | Edit system hook


# **DeleteApiV4HooksHookId**
> ApiEntitiesHook DeleteApiV4HooksHookId(ctx, hookId)
Delete system hook

Deletes a system hook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hookId** | **int32**| The ID of the system hook | 

### Return type

[**ApiEntitiesHook**](API_Entities_Hook.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4Hooks**
> []ApiEntitiesHook GetApiV4Hooks(ctx, optional)
List system hooks

Get a list of all system hooks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemHooksApiGetApiV4HooksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemHooksApiGetApiV4HooksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesHook**](API_Entities_Hook.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4HooksHookId**
> ApiEntitiesHook GetApiV4HooksHookId(ctx, hookId)
Get system hook

Get a system hook by its ID. Introduced in GitLab 14.9.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hookId** | **int32**| The ID of the system hook | 

### Return type

[**ApiEntitiesHook**](API_Entities_Hook.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4Hooks**
> ApiEntitiesHook PostApiV4Hooks(ctx, postApiV4Hooks)
Add new system hook

Add a new system hook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4Hooks** | [**PostApiV4Hooks**](PostApiV4Hooks.md)|  | 

### Return type

[**ApiEntitiesHook**](API_Entities_Hook.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4HooksHookId**
> ApiEntitiesHook PutApiV4HooksHookId(ctx, hookId, putApiV4HooksHookId)
Edit system hook

Edits a system hook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hookId** | **int32**| The ID of the system hook | 
  **putApiV4HooksHookId** | [**PutApiV4HooksHookId**](PutApiV4HooksHookId.md)|  | 

### Return type

[**ApiEntitiesHook**](API_Entities_Hook.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

