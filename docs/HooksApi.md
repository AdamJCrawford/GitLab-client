# \HooksApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4HooksHookIdUrlVariablesKey**](HooksApi.md#DeleteApiV4HooksHookIdUrlVariablesKey) | **Delete** /api/v4/hooks/{hook_id}/url_variables/{key} | 
[**PostApiV4HooksHookId**](HooksApi.md#PostApiV4HooksHookId) | **Post** /api/v4/hooks/{hook_id} | 
[**PutApiV4HooksHookIdUrlVariablesKey**](HooksApi.md#PutApiV4HooksHookIdUrlVariablesKey) | **Put** /api/v4/hooks/{hook_id}/url_variables/{key} | 


# **DeleteApiV4HooksHookIdUrlVariablesKey**
> DeleteApiV4HooksHookIdUrlVariablesKey(ctx, hookId, key)


Un-Set a url variable

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hookId** | **int32**| The ID of the hook | 
  **key** | **string**| The key of the variable | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4HooksHookId**
> PostApiV4HooksHookId(ctx, hookId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hookId** | **int32**| The ID of the hook | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4HooksHookIdUrlVariablesKey**
> PutApiV4HooksHookIdUrlVariablesKey(ctx, hookId, key, putApiV4HooksHookIdUrlVariablesKey)


Set a url variable

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hookId** | **int32**| The ID of the hook | 
  **key** | **string**| The key of the variable | 
  **putApiV4HooksHookIdUrlVariablesKey** | [**PutApiV4HooksHookIdUrlVariablesKey**](PutApiV4HooksHookIdUrlVariablesKey.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

