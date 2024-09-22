# \MigrationsApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostApiV4AdminMigrationsTimestampMark**](MigrationsApi.md#PostApiV4AdminMigrationsTimestampMark) | **Post** /api/v4/admin/migrations/{timestamp}/mark | 


# **PostApiV4AdminMigrationsTimestampMark**
> PostApiV4AdminMigrationsTimestampMark(ctx, timestamp, postApiV4AdminMigrationsTimestampMark)


Mark the migration as successfully executed

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **timestamp** | **int32**| The migration version timestamp | 
  **postApiV4AdminMigrationsTimestampMark** | [**PostApiV4AdminMigrationsTimestampMark**](PostApiV4AdminMigrationsTimestampMark.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

