# \GroupExportApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4GroupsIdExportDownload**](GroupExportApi.md#GetApiV4GroupsIdExportDownload) | **Get** /api/v4/groups/{id}/export/download | Download export
[**GetApiV4GroupsIdExportRelationsDownload**](GroupExportApi.md#GetApiV4GroupsIdExportRelationsDownload) | **Get** /api/v4/groups/{id}/export_relations/download | Download relations export
[**GetApiV4GroupsIdExportRelationsStatus**](GroupExportApi.md#GetApiV4GroupsIdExportRelationsStatus) | **Get** /api/v4/groups/{id}/export_relations/status | Relations export status
[**PostApiV4GroupsIdExport**](GroupExportApi.md#PostApiV4GroupsIdExport) | **Post** /api/v4/groups/{id}/export | Start export
[**PostApiV4GroupsIdExportRelations**](GroupExportApi.md#PostApiV4GroupsIdExportRelations) | **Post** /api/v4/groups/{id}/export_relations | Start relations export


# **GetApiV4GroupsIdExportDownload**
> GetApiV4GroupsIdExportDownload(ctx, id)
Download export

This feature was introduced in GitLab 12.5.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdExportRelationsDownload**
> GetApiV4GroupsIdExportRelationsDownload(ctx, id, relation, optional)
Download relations export

This feature was introduced in GitLab 13.12

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group | 
  **relation** | **string**| Group relation name | 
 **optional** | ***GroupExportApiGetApiV4GroupsIdExportRelationsDownloadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupExportApiGetApiV4GroupsIdExportRelationsDownloadOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batched** | **optional.Bool**| Whether to download in batches | 
 **batchNumber** | **optional.Int32**| Batch number to download | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdExportRelationsStatus**
> []ApiEntitiesBulkImportsExportStatus GetApiV4GroupsIdExportRelationsStatus(ctx, id, optional)
Relations export status

This feature was introduced in GitLab 13.12

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group | 
 **optional** | ***GroupExportApiGetApiV4GroupsIdExportRelationsStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupExportApiGetApiV4GroupsIdExportRelationsStatusOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **relation** | **optional.String**| Group relation name | 

### Return type

[**[]ApiEntitiesBulkImportsExportStatus**](API_Entities_BulkImports_ExportStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4GroupsIdExport**
> PostApiV4GroupsIdExport(ctx, id)
Start export

This feature was introduced in GitLab 12.5.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4GroupsIdExportRelations**
> PostApiV4GroupsIdExportRelations(ctx, id, postApiV4GroupsIdExportRelations)
Start relations export

This feature was introduced in GitLab 13.12

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group | 
  **postApiV4GroupsIdExportRelations** | [**PostApiV4GroupsIdExportRelations**](PostApiV4GroupsIdExportRelations.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

