# \UsageDataApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4UsageDataMetricDefinitions**](UsageDataApi.md#GetApiV4UsageDataMetricDefinitions) | **Get** /api/v4/usage_data/metric_definitions | Get a list of all metric definitions
[**GetApiV4UsageDataNonSqlMetrics**](UsageDataApi.md#GetApiV4UsageDataNonSqlMetrics) | **Get** /api/v4/usage_data/non_sql_metrics | Get Non SQL usage ping metrics
[**GetApiV4UsageDataQueries**](UsageDataApi.md#GetApiV4UsageDataQueries) | **Get** /api/v4/usage_data/queries | Get raw SQL queries for usage data SQL metrics
[**GetApiV4UsageDataServicePing**](UsageDataApi.md#GetApiV4UsageDataServicePing) | **Get** /api/v4/usage_data/service_ping | Get the latest ServicePing payload
[**PostApiV4UsageDataIncrementCounter**](UsageDataApi.md#PostApiV4UsageDataIncrementCounter) | **Post** /api/v4/usage_data/increment_counter | Track usage data event
[**PostApiV4UsageDataIncrementUniqueUsers**](UsageDataApi.md#PostApiV4UsageDataIncrementUniqueUsers) | **Post** /api/v4/usage_data/increment_unique_users | 
[**PostApiV4UsageDataTrackEvent**](UsageDataApi.md#PostApiV4UsageDataTrackEvent) | **Post** /api/v4/usage_data/track_event | Track gitlab internal events


# **GetApiV4UsageDataMetricDefinitions**
> GetApiV4UsageDataMetricDefinitions(ctx, )
Get a list of all metric definitions

This feature was introduced in GitLab 13.11.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4UsageDataNonSqlMetrics**
> GetApiV4UsageDataNonSqlMetrics(ctx, )
Get Non SQL usage ping metrics

This feature was introduced in GitLab 13.11.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4UsageDataQueries**
> GetApiV4UsageDataQueries(ctx, )
Get raw SQL queries for usage data SQL metrics

This feature was introduced in GitLab 13.11.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4UsageDataServicePing**
> GetApiV4UsageDataServicePing(ctx, )
Get the latest ServicePing payload

Introduces in Gitlab 16.9. Requires Personal Access Token with read_service_ping scope.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4UsageDataIncrementCounter**
> PostApiV4UsageDataIncrementCounter(ctx, postApiV4UsageDataIncrementCounter)
Track usage data event

This feature was introduced in GitLab 13.4.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4UsageDataIncrementCounter** | [**PostApiV4UsageDataIncrementCounter**](PostApiV4UsageDataIncrementCounter.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4UsageDataIncrementUniqueUsers**
> PostApiV4UsageDataIncrementUniqueUsers(ctx, postApiV4UsageDataIncrementUniqueUsers)


Track usage data event for the current user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4UsageDataIncrementUniqueUsers** | [**PostApiV4UsageDataIncrementUniqueUsers**](PostApiV4UsageDataIncrementUniqueUsers.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4UsageDataTrackEvent**
> PostApiV4UsageDataTrackEvent(ctx, postApiV4UsageDataTrackEvent)
Track gitlab internal events

This feature was introduced in GitLab 16.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4UsageDataTrackEvent** | [**PostApiV4UsageDataTrackEvent**](PostApiV4UsageDataTrackEvent.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

