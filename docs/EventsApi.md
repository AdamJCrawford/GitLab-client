# \EventsApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4Events**](EventsApi.md#GetApiV4Events) | **Get** /api/v4/events | List currently authenticated user&#39;s events
[**GetApiV4UsersIdEvents**](EventsApi.md#GetApiV4UsersIdEvents) | **Get** /api/v4/users/{id}/events | Get the contribution events of a specified user


# **GetApiV4Events**
> []ApiEntitiesEvent GetApiV4Events(ctx, optional)
List currently authenticated user's events

This feature was introduced in GitLab 9.3.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventsApiGetApiV4EventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventsApiGetApiV4EventsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **optional.String**| Include all events across a user’s projects | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **action** | **optional.String**| Event action to filter on | 
 **targetType** | **optional.String**| Event target type to filter on | 
 **before** | **optional.String**| Include only events created before this date | 
 **after** | **optional.String**| Include only events created after this date | 
 **sort** | **optional.String**| Return events sorted in ascending and descending order | [default to desc]

### Return type

[**[]ApiEntitiesEvent**](API_Entities_Event.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4UsersIdEvents**
> []ApiEntitiesEvent GetApiV4UsersIdEvents(ctx, id, optional)
Get the contribution events of a specified user

This feature was introduced in GitLab 8.13.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or username of the user | 
 **optional** | ***EventsApiGetApiV4UsersIdEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventsApiGetApiV4UsersIdEventsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **action** | **optional.String**| Event action to filter on | 
 **targetType** | **optional.String**| Event target type to filter on | 
 **before** | **optional.String**| Include only events created before this date | 
 **after** | **optional.String**| Include only events created after this date | 
 **sort** | **optional.String**| Return events sorted in ascending and descending order | [default to desc]

### Return type

[**[]ApiEntitiesEvent**](API_Entities_Event.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

