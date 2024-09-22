# \BulkImportsApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4BulkImports**](BulkImportsApi.md#GetApiV4BulkImports) | **Get** /api/v4/bulk_imports | List all GitLab Migrations
[**GetApiV4BulkImportsEntities**](BulkImportsApi.md#GetApiV4BulkImportsEntities) | **Get** /api/v4/bulk_imports/entities | List all GitLab Migrations&#39; entities
[**GetApiV4BulkImportsImportId**](BulkImportsApi.md#GetApiV4BulkImportsImportId) | **Get** /api/v4/bulk_imports/{import_id} | Get GitLab Migration details
[**GetApiV4BulkImportsImportIdEntities**](BulkImportsApi.md#GetApiV4BulkImportsImportIdEntities) | **Get** /api/v4/bulk_imports/{import_id}/entities | List GitLab Migration entities
[**GetApiV4BulkImportsImportIdEntitiesEntityId**](BulkImportsApi.md#GetApiV4BulkImportsImportIdEntitiesEntityId) | **Get** /api/v4/bulk_imports/{import_id}/entities/{entity_id} | Get GitLab Migration entity details
[**GetApiV4BulkImportsImportIdEntitiesEntityIdFailures**](BulkImportsApi.md#GetApiV4BulkImportsImportIdEntitiesEntityIdFailures) | **Get** /api/v4/bulk_imports/{import_id}/entities/{entity_id}/failures | Get GitLab Migration entity failures
[**PostApiV4BulkImports**](BulkImportsApi.md#PostApiV4BulkImports) | **Post** /api/v4/bulk_imports | Start a new GitLab Migration
[**PostApiV4BulkImportsImportIdCancel**](BulkImportsApi.md#PostApiV4BulkImportsImportIdCancel) | **Post** /api/v4/bulk_imports/{import_id}/cancel | Cancel GitLab Migration


# **GetApiV4BulkImports**
> []ApiEntitiesBulkImport GetApiV4BulkImports(ctx, optional)
List all GitLab Migrations

This feature was introduced in GitLab 14.1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BulkImportsApiGetApiV4BulkImportsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BulkImportsApiGetApiV4BulkImportsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **sort** | **optional.String**| Return GitLab Migrations sorted in created by &#x60;asc&#x60; or &#x60;desc&#x60; order. | [default to desc]
 **status** | **optional.String**| Return GitLab Migrations with specified status | 

### Return type

[**[]ApiEntitiesBulkImport**](API_Entities_BulkImport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4BulkImportsEntities**
> []ApiEntitiesBulkImports GetApiV4BulkImportsEntities(ctx, optional)
List all GitLab Migrations' entities

This feature was introduced in GitLab 14.1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BulkImportsApiGetApiV4BulkImportsEntitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BulkImportsApiGetApiV4BulkImportsEntitiesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **sort** | **optional.String**| Return GitLab Migrations sorted in created by &#x60;asc&#x60; or &#x60;desc&#x60; order. | [default to desc]
 **status** | **optional.String**| Return all GitLab Migrations&#39; entities with specified status | 

### Return type

[**[]ApiEntitiesBulkImports**](API_Entities_BulkImports.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4BulkImportsImportId**
> ApiEntitiesBulkImport GetApiV4BulkImportsImportId(ctx, importId)
Get GitLab Migration details

This feature was introduced in GitLab 14.1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **importId** | **int32**| The ID of user&#39;s GitLab Migration | 

### Return type

[**ApiEntitiesBulkImport**](API_Entities_BulkImport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4BulkImportsImportIdEntities**
> []ApiEntitiesBulkImports GetApiV4BulkImportsImportIdEntities(ctx, importId, optional)
List GitLab Migration entities

This feature was introduced in GitLab 14.1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **importId** | **int32**| The ID of user&#39;s GitLab Migration | 
 **optional** | ***BulkImportsApiGetApiV4BulkImportsImportIdEntitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BulkImportsApiGetApiV4BulkImportsImportIdEntitiesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **optional.String**| Return import entities with specified status | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesBulkImports**](API_Entities_BulkImports.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4BulkImportsImportIdEntitiesEntityId**
> ApiEntitiesBulkImports GetApiV4BulkImportsImportIdEntitiesEntityId(ctx, importId, entityId)
Get GitLab Migration entity details

This feature was introduced in GitLab 14.1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **importId** | **int32**| The ID of user&#39;s GitLab Migration | 
  **entityId** | **int32**| The ID of GitLab Migration entity | 

### Return type

[**ApiEntitiesBulkImports**](API_Entities_BulkImports.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4BulkImportsImportIdEntitiesEntityIdFailures**
> ApiEntitiesBulkImportsEntityFailure GetApiV4BulkImportsImportIdEntitiesEntityIdFailures(ctx, importId, entityId)
Get GitLab Migration entity failures

This feature was introduced in GitLab 16.6

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **importId** | **int32**| The ID of user&#39;s GitLab Migration | 
  **entityId** | **int32**| The ID of GitLab Migration entity | 

### Return type

[**ApiEntitiesBulkImportsEntityFailure**](API_Entities_BulkImports_EntityFailure.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4BulkImports**
> ApiEntitiesBulkImport PostApiV4BulkImports(ctx, configurationUrl, configurationAccessToken, entitiesSourceType, entitiesSourceFullPath, entitiesDestinationNamespace, optional)
Start a new GitLab Migration

This feature was introduced in GitLab 14.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configurationUrl** | **string**| Source GitLab instance URL | 
  **configurationAccessToken** | **string**| Access token to the source GitLab instance | 
  **entitiesSourceType** | [**[]string**](string.md)| Source entity type | 
  **entitiesSourceFullPath** | [**[]string**](string.md)| Relative path of the source entity to import | 
  **entitiesDestinationNamespace** | [**[]string**](string.md)| Destination namespace for the entity | 
 **optional** | ***BulkImportsApiPostApiV4BulkImportsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BulkImportsApiPostApiV4BulkImportsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **entitiesDestinationSlug** | [**optional.Interface of []string**](string.md)| Destination slug for the entity | 
 **entitiesDestinationName** | [**optional.Interface of []string**](string.md)| Deprecated: Use :destination_slug instead. Destination slug for the entity | 
 **entitiesMigrateProjects** | [**optional.Interface of []bool**](bool.md)| Indicates group migration should include nested projects | [default to true]

### Return type

[**ApiEntitiesBulkImport**](API_Entities_BulkImport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4BulkImportsImportIdCancel**
> ApiEntitiesBulkImport PostApiV4BulkImportsImportIdCancel(ctx, importId)
Cancel GitLab Migration

This feature was introduced in GitLab 17.1

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **importId** | **int32**| The ID of user&#39;s GitLab Migration | 

### Return type

[**ApiEntitiesBulkImport**](API_Entities_BulkImport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

