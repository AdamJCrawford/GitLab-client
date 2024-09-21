# \ProjectExportApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4ProjectsIdExport**](ProjectExportApi.md#GetApiV4ProjectsIdExport) | **Get** /api/v4/projects/{id}/export | Get export status
[**GetApiV4ProjectsIdExportDownload**](ProjectExportApi.md#GetApiV4ProjectsIdExportDownload) | **Get** /api/v4/projects/{id}/export/download | Download export
[**GetApiV4ProjectsIdExportRelationsDownload**](ProjectExportApi.md#GetApiV4ProjectsIdExportRelationsDownload) | **Get** /api/v4/projects/{id}/export_relations/download | Download relations export
[**GetApiV4ProjectsIdExportRelationsStatus**](ProjectExportApi.md#GetApiV4ProjectsIdExportRelationsStatus) | **Get** /api/v4/projects/{id}/export_relations/status | Relations export status
[**PostApiV4ProjectsIdExport**](ProjectExportApi.md#PostApiV4ProjectsIdExport) | **Post** /api/v4/projects/{id}/export | Start export
[**PostApiV4ProjectsIdExportRelations**](ProjectExportApi.md#PostApiV4ProjectsIdExportRelations) | **Post** /api/v4/projects/{id}/export_relations | Start relations export


# **GetApiV4ProjectsIdExport**
> ApiEntitiesProjectExportStatus GetApiV4ProjectsIdExport(ctx, id)
Get export status

This feature was introduced in GitLab 10.6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 

### Return type

[**ApiEntitiesProjectExportStatus**](API_Entities_ProjectExportStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdExportDownload**
> GetApiV4ProjectsIdExportDownload(ctx, id)
Download export

This feature was introduced in GitLab 10.6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdExportRelationsDownload**
> GetApiV4ProjectsIdExportRelationsDownload(ctx, id, relation, optional)
Download relations export

This feature was introduced in GitLab 14.4

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **relation** | **string**| Project relation name | 
 **optional** | ***ProjectExportApiGetApiV4ProjectsIdExportRelationsDownloadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectExportApiGetApiV4ProjectsIdExportRelationsDownloadOpts struct

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
 - **Accept**: application/octet-stream, application/gzip, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdExportRelationsStatus**
> []ApiEntitiesBulkImportsExportStatus GetApiV4ProjectsIdExportRelationsStatus(ctx, id, optional)
Relations export status

This feature was introduced in GitLab 14.4

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
 **optional** | ***ProjectExportApiGetApiV4ProjectsIdExportRelationsStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectExportApiGetApiV4ProjectsIdExportRelationsStatusOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **relation** | **optional.String**| Project relation name | 

### Return type

[**[]ApiEntitiesBulkImportsExportStatus**](API_Entities_BulkImports_ExportStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdExport**
> PostApiV4ProjectsIdExport(ctx, id, postApiV4ProjectsIdExport)
Start export

This feature was introduced in GitLab 10.6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **postApiV4ProjectsIdExport** | [**PostApiV4ProjectsIdExport**](PostApiV4ProjectsIdExport.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdExportRelations**
> PostApiV4ProjectsIdExportRelations(ctx, id, postApiV4ProjectsIdExportRelations)
Start relations export

This feature was introduced in GitLab 14.4

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **postApiV4ProjectsIdExportRelations** | [**PostApiV4ProjectsIdExportRelations**](PostApiV4ProjectsIdExportRelations.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

