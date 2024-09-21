# \BroadcastMessagesApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4BroadcastMessagesId**](BroadcastMessagesApi.md#DeleteApiV4BroadcastMessagesId) | **Delete** /api/v4/broadcast_messages/{id} | Delete a broadcast message
[**GetApiV4BroadcastMessages**](BroadcastMessagesApi.md#GetApiV4BroadcastMessages) | **Get** /api/v4/broadcast_messages | Get all broadcast messages
[**GetApiV4BroadcastMessagesId**](BroadcastMessagesApi.md#GetApiV4BroadcastMessagesId) | **Get** /api/v4/broadcast_messages/{id} | Get a specific broadcast message
[**PostApiV4BroadcastMessages**](BroadcastMessagesApi.md#PostApiV4BroadcastMessages) | **Post** /api/v4/broadcast_messages | Create a broadcast message
[**PutApiV4BroadcastMessagesId**](BroadcastMessagesApi.md#PutApiV4BroadcastMessagesId) | **Put** /api/v4/broadcast_messages/{id} | Update a broadcast message


# **DeleteApiV4BroadcastMessagesId**
> ApiEntitiesSystemBroadcastMessage DeleteApiV4BroadcastMessagesId(ctx, id)
Delete a broadcast message

This feature was introduced in GitLab 8.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Broadcast message ID | 

### Return type

[**ApiEntitiesSystemBroadcastMessage**](API_Entities_System_BroadcastMessage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4BroadcastMessages**
> ApiEntitiesSystemBroadcastMessage GetApiV4BroadcastMessages(ctx, optional)
Get all broadcast messages

This feature was introduced in GitLab 8.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BroadcastMessagesApiGetApiV4BroadcastMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BroadcastMessagesApiGetApiV4BroadcastMessagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**ApiEntitiesSystemBroadcastMessage**](API_Entities_System_BroadcastMessage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4BroadcastMessagesId**
> ApiEntitiesSystemBroadcastMessage GetApiV4BroadcastMessagesId(ctx, id)
Get a specific broadcast message

This feature was introduced in GitLab 8.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Broadcast message ID | 

### Return type

[**ApiEntitiesSystemBroadcastMessage**](API_Entities_System_BroadcastMessage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4BroadcastMessages**
> ApiEntitiesSystemBroadcastMessage PostApiV4BroadcastMessages(ctx, postApiV4BroadcastMessages)
Create a broadcast message

This feature was introduced in GitLab 8.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4BroadcastMessages** | [**PostApiV4BroadcastMessages**](PostApiV4BroadcastMessages.md)|  | 

### Return type

[**ApiEntitiesSystemBroadcastMessage**](API_Entities_System_BroadcastMessage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4BroadcastMessagesId**
> ApiEntitiesSystemBroadcastMessage PutApiV4BroadcastMessagesId(ctx, id, putApiV4BroadcastMessagesId)
Update a broadcast message

This feature was introduced in GitLab 8.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Broadcast message ID | 
  **putApiV4BroadcastMessagesId** | [**PutApiV4BroadcastMessagesId**](PutApiV4BroadcastMessagesId.md)|  | 

### Return type

[**ApiEntitiesSystemBroadcastMessage**](API_Entities_System_BroadcastMessage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

