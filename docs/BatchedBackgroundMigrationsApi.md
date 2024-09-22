# \BatchedBackgroundMigrationsApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4AdminBatchedBackgroundMigrations**](BatchedBackgroundMigrationsApi.md#GetApiV4AdminBatchedBackgroundMigrations) | **Get** /api/v4/admin/batched_background_migrations | 
[**GetApiV4AdminBatchedBackgroundMigrationsId**](BatchedBackgroundMigrationsApi.md#GetApiV4AdminBatchedBackgroundMigrationsId) | **Get** /api/v4/admin/batched_background_migrations/{id} | 
[**PutApiV4AdminBatchedBackgroundMigrationsIdPause**](BatchedBackgroundMigrationsApi.md#PutApiV4AdminBatchedBackgroundMigrationsIdPause) | **Put** /api/v4/admin/batched_background_migrations/{id}/pause | 
[**PutApiV4AdminBatchedBackgroundMigrationsIdResume**](BatchedBackgroundMigrationsApi.md#PutApiV4AdminBatchedBackgroundMigrationsIdResume) | **Put** /api/v4/admin/batched_background_migrations/{id}/resume | 


# **GetApiV4AdminBatchedBackgroundMigrations**
> []ApiEntitiesBatchedBackgroundMigration GetApiV4AdminBatchedBackgroundMigrations(ctx, optional)


Get the list of batched background migrations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BatchedBackgroundMigrationsApiGetApiV4AdminBatchedBackgroundMigrationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BatchedBackgroundMigrationsApiGetApiV4AdminBatchedBackgroundMigrationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **database** | **optional.String**| The name of the database, the default &#x60;main&#x60; | [default to main]

### Return type

[**[]ApiEntitiesBatchedBackgroundMigration**](API_Entities_BatchedBackgroundMigration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4AdminBatchedBackgroundMigrationsId**
> ApiEntitiesBatchedBackgroundMigration GetApiV4AdminBatchedBackgroundMigrationsId(ctx, id, optional)


Retrieve a batched background migration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The batched background migration id | 
 **optional** | ***BatchedBackgroundMigrationsApiGetApiV4AdminBatchedBackgroundMigrationsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BatchedBackgroundMigrationsApiGetApiV4AdminBatchedBackgroundMigrationsIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **database** | **optional.String**| The name of the database | [default to main]

### Return type

[**ApiEntitiesBatchedBackgroundMigration**](API_Entities_BatchedBackgroundMigration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4AdminBatchedBackgroundMigrationsIdPause**
> ApiEntitiesBatchedBackgroundMigration PutApiV4AdminBatchedBackgroundMigrationsIdPause(ctx, id, putApiV4AdminBatchedBackgroundMigrationsIdPause)


Pause a batched background migration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The batched background migration id | 
  **putApiV4AdminBatchedBackgroundMigrationsIdPause** | [**PutApiV4AdminBatchedBackgroundMigrationsIdPause**](PutApiV4AdminBatchedBackgroundMigrationsIdPause.md)|  | 

### Return type

[**ApiEntitiesBatchedBackgroundMigration**](API_Entities_BatchedBackgroundMigration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4AdminBatchedBackgroundMigrationsIdResume**
> ApiEntitiesBatchedBackgroundMigration PutApiV4AdminBatchedBackgroundMigrationsIdResume(ctx, id, putApiV4AdminBatchedBackgroundMigrationsIdResume)


Resume a batched background migration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The batched background migration id | 
  **putApiV4AdminBatchedBackgroundMigrationsIdResume** | [**PutApiV4AdminBatchedBackgroundMigrationsIdResume**](PutApiV4AdminBatchedBackgroundMigrationsIdResume.md)|  | 

### Return type

[**ApiEntitiesBatchedBackgroundMigration**](API_Entities_BatchedBackgroundMigration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

