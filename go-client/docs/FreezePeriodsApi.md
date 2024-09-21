# \FreezePeriodsApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4ProjectsIdFreezePeriodsFreezePeriodId**](FreezePeriodsApi.md#DeleteApiV4ProjectsIdFreezePeriodsFreezePeriodId) | **Delete** /api/v4/projects/{id}/freeze_periods/{freeze_period_id} | Delete a freeze period
[**GetApiV4ProjectsIdFreezePeriods**](FreezePeriodsApi.md#GetApiV4ProjectsIdFreezePeriods) | **Get** /api/v4/projects/{id}/freeze_periods | List freeze periods
[**GetApiV4ProjectsIdFreezePeriodsFreezePeriodId**](FreezePeriodsApi.md#GetApiV4ProjectsIdFreezePeriodsFreezePeriodId) | **Get** /api/v4/projects/{id}/freeze_periods/{freeze_period_id} | Get a freeze period
[**PostApiV4ProjectsIdFreezePeriods**](FreezePeriodsApi.md#PostApiV4ProjectsIdFreezePeriods) | **Post** /api/v4/projects/{id}/freeze_periods | Create a freeze period
[**PutApiV4ProjectsIdFreezePeriodsFreezePeriodId**](FreezePeriodsApi.md#PutApiV4ProjectsIdFreezePeriodsFreezePeriodId) | **Put** /api/v4/projects/{id}/freeze_periods/{freeze_period_id} | Update a freeze period


# **DeleteApiV4ProjectsIdFreezePeriodsFreezePeriodId**
> ApiEntitiesFreezePeriod DeleteApiV4ProjectsIdFreezePeriodsFreezePeriodId(ctx, id, freezePeriodId)
Delete a freeze period

Deletes a freeze period for the given `freeze_period_id`. This feature was introduced in GitLab 13.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **freezePeriodId** | **int32**| The ID of the freeze period | 

### Return type

[**ApiEntitiesFreezePeriod**](API_Entities_FreezePeriod.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdFreezePeriods**
> []ApiEntitiesFreezePeriod GetApiV4ProjectsIdFreezePeriods(ctx, id, optional)
List freeze periods

Paginated list of Freeze Periods, sorted by created_at in ascending order. This feature was introduced in GitLab 13.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
 **optional** | ***FreezePeriodsApiGetApiV4ProjectsIdFreezePeriodsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FreezePeriodsApiGetApiV4ProjectsIdFreezePeriodsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesFreezePeriod**](API_Entities_FreezePeriod.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdFreezePeriodsFreezePeriodId**
> ApiEntitiesFreezePeriod GetApiV4ProjectsIdFreezePeriodsFreezePeriodId(ctx, id, freezePeriodId)
Get a freeze period

Get a freeze period for the given `freeze_period_id`. This feature was introduced in GitLab 13.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **freezePeriodId** | **int32**| The ID of the freeze period | 

### Return type

[**ApiEntitiesFreezePeriod**](API_Entities_FreezePeriod.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdFreezePeriods**
> ApiEntitiesFreezePeriod PostApiV4ProjectsIdFreezePeriods(ctx, id, postApiV4ProjectsIdFreezePeriods)
Create a freeze period

Creates a freeze period. This feature was introduced in GitLab 13.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **postApiV4ProjectsIdFreezePeriods** | [**PostApiV4ProjectsIdFreezePeriods**](PostApiV4ProjectsIdFreezePeriods.md)|  | 

### Return type

[**ApiEntitiesFreezePeriod**](API_Entities_FreezePeriod.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdFreezePeriodsFreezePeriodId**
> ApiEntitiesFreezePeriod PutApiV4ProjectsIdFreezePeriodsFreezePeriodId(ctx, id, freezePeriodId, putApiV4ProjectsIdFreezePeriodsFreezePeriodId)
Update a freeze period

Updates a freeze period for the given `freeze_period_id`. This feature was introduced in GitLab 13.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **freezePeriodId** | **int32**|  | 
  **putApiV4ProjectsIdFreezePeriodsFreezePeriodId** | [**PutApiV4ProjectsIdFreezePeriodsFreezePeriodId**](PutApiV4ProjectsIdFreezePeriodsFreezePeriodId.md)|  | 

### Return type

[**ApiEntitiesFreezePeriod**](API_Entities_FreezePeriod.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

