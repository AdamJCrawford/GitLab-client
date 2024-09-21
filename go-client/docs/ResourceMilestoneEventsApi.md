# \ResourceMilestoneEventsApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4ProjectsIdIssuesEventableIdResourceMilestoneEvents**](ResourceMilestoneEventsApi.md#GetApiV4ProjectsIdIssuesEventableIdResourceMilestoneEvents) | **Get** /api/v4/projects/{id}/issues/{eventable_id}/resource_milestone_events | List project Issue milestone events
[**GetApiV4ProjectsIdIssuesEventableIdResourceMilestoneEventsEventId**](ResourceMilestoneEventsApi.md#GetApiV4ProjectsIdIssuesEventableIdResourceMilestoneEventsEventId) | **Get** /api/v4/projects/{id}/issues/{eventable_id}/resource_milestone_events/{event_id} | Get single Issue milestone event
[**GetApiV4ProjectsIdMergeRequestsEventableIdResourceMilestoneEvents**](ResourceMilestoneEventsApi.md#GetApiV4ProjectsIdMergeRequestsEventableIdResourceMilestoneEvents) | **Get** /api/v4/projects/{id}/merge_requests/{eventable_id}/resource_milestone_events | List project Merge request milestone events
[**GetApiV4ProjectsIdMergeRequestsEventableIdResourceMilestoneEventsEventId**](ResourceMilestoneEventsApi.md#GetApiV4ProjectsIdMergeRequestsEventableIdResourceMilestoneEventsEventId) | **Get** /api/v4/projects/{id}/merge_requests/{eventable_id}/resource_milestone_events/{event_id} | Get single Merge request milestone event


# **GetApiV4ProjectsIdIssuesEventableIdResourceMilestoneEvents**
> []ApiEntitiesResourceMilestoneEvent GetApiV4ProjectsIdIssuesEventableIdResourceMilestoneEvents(ctx, id, eventableId, optional)
List project Issue milestone events

Gets a list of all milestone events for a single Issue

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **eventableId** | **int32**| The ID of the eventable | 
 **optional** | ***ResourceMilestoneEventsApiGetApiV4ProjectsIdIssuesEventableIdResourceMilestoneEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ResourceMilestoneEventsApiGetApiV4ProjectsIdIssuesEventableIdResourceMilestoneEventsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesResourceMilestoneEvent**](API_Entities_ResourceMilestoneEvent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdIssuesEventableIdResourceMilestoneEventsEventId**
> ApiEntitiesResourceMilestoneEvent GetApiV4ProjectsIdIssuesEventableIdResourceMilestoneEventsEventId(ctx, id, eventId, eventableId)
Get single Issue milestone event

Returns a single milestone event for a specific project Issue

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **eventId** | **string**| The ID of a resource milestone event | 
  **eventableId** | **int32**| The ID of the eventable | 

### Return type

[**ApiEntitiesResourceMilestoneEvent**](API_Entities_ResourceMilestoneEvent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdMergeRequestsEventableIdResourceMilestoneEvents**
> []ApiEntitiesResourceMilestoneEvent GetApiV4ProjectsIdMergeRequestsEventableIdResourceMilestoneEvents(ctx, id, eventableId, optional)
List project Merge request milestone events

Gets a list of all milestone events for a single Merge request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **eventableId** | **int32**| The ID of the eventable | 
 **optional** | ***ResourceMilestoneEventsApiGetApiV4ProjectsIdMergeRequestsEventableIdResourceMilestoneEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ResourceMilestoneEventsApiGetApiV4ProjectsIdMergeRequestsEventableIdResourceMilestoneEventsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesResourceMilestoneEvent**](API_Entities_ResourceMilestoneEvent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdMergeRequestsEventableIdResourceMilestoneEventsEventId**
> ApiEntitiesResourceMilestoneEvent GetApiV4ProjectsIdMergeRequestsEventableIdResourceMilestoneEventsEventId(ctx, id, eventId, eventableId)
Get single Merge request milestone event

Returns a single milestone event for a specific project Merge request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **eventId** | **string**| The ID of a resource milestone event | 
  **eventableId** | **int32**| The ID of the eventable | 

### Return type

[**ApiEntitiesResourceMilestoneEvent**](API_Entities_ResourceMilestoneEvent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

