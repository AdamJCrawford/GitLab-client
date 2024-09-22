# \TopicsApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4TopicsId**](TopicsApi.md#DeleteApiV4TopicsId) | **Delete** /api/v4/topics/{id} | Delete a topic
[**GetApiV4Topics**](TopicsApi.md#GetApiV4Topics) | **Get** /api/v4/topics | Get topics
[**GetApiV4TopicsId**](TopicsApi.md#GetApiV4TopicsId) | **Get** /api/v4/topics/{id} | Get topic
[**PostApiV4Topics**](TopicsApi.md#PostApiV4Topics) | **Post** /api/v4/topics | Create a topic
[**PostApiV4TopicsMerge**](TopicsApi.md#PostApiV4TopicsMerge) | **Post** /api/v4/topics/merge | Merge topics
[**PutApiV4TopicsId**](TopicsApi.md#PutApiV4TopicsId) | **Put** /api/v4/topics/{id} | Update a topic


# **DeleteApiV4TopicsId**
> DeleteApiV4TopicsId(ctx, id)
Delete a topic

This feature was introduced in GitLab 14.9.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| ID of project topic | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4Topics**
> ApiEntitiesProjectsTopic GetApiV4Topics(ctx, optional)
Get topics

This feature was introduced in GitLab 14.5.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TopicsApiGetApiV4TopicsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TopicsApiGetApiV4TopicsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| Return list of topics matching the search criteria | 
 **withoutProjects** | **optional.Bool**| Return list of topics without assigned projects | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**ApiEntitiesProjectsTopic**](API_Entities_Projects_Topic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4TopicsId**
> ApiEntitiesProjectsTopic GetApiV4TopicsId(ctx, id)
Get topic

This feature was introduced in GitLab 14.5.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| ID of project topic | 

### Return type

[**ApiEntitiesProjectsTopic**](API_Entities_Projects_Topic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4Topics**
> ApiEntitiesProjectsTopic PostApiV4Topics(ctx, postApiV4Topics)
Create a topic

This feature was introduced in GitLab 14.5.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4Topics** | [**PostApiV4Topics**](PostApiV4Topics.md)|  | 

### Return type

[**ApiEntitiesProjectsTopic**](API_Entities_Projects_Topic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4TopicsMerge**
> ApiEntitiesProjectsTopic PostApiV4TopicsMerge(ctx, postApiV4TopicsMerge)
Merge topics

This feature was introduced in GitLab 15.4.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4TopicsMerge** | [**PostApiV4TopicsMerge**](PostApiV4TopicsMerge.md)|  | 

### Return type

[**ApiEntitiesProjectsTopic**](API_Entities_Projects_Topic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4TopicsId**
> ApiEntitiesProjectsTopic PutApiV4TopicsId(ctx, id, putApiV4TopicsId)
Update a topic

This feature was introduced in GitLab 14.5.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| ID of project topic | 
  **putApiV4TopicsId** | [**PutApiV4TopicsId**](PutApiV4TopicsId.md)|  | 

### Return type

[**ApiEntitiesProjectsTopic**](API_Entities_Projects_Topic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

