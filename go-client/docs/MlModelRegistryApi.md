# \MlModelRegistryApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4ProjectsIdPackagesMlModelsModelVersionIdFilesPathFileName**](MlModelRegistryApi.md#GetApiV4ProjectsIdPackagesMlModelsModelVersionIdFilesPathFileName) | **Get** /api/v4/projects/{id}/packages/ml_models/{model_version_id}/files/(*path/){file_name} | Download an ml_model package file
[**PutApiV4ProjectsIdPackagesMlModelsModelVersionIdFilesPathFileName**](MlModelRegistryApi.md#PutApiV4ProjectsIdPackagesMlModelsModelVersionIdFilesPathFileName) | **Put** /api/v4/projects/{id}/packages/ml_models/{model_version_id}/files/(*path/){file_name} | Workhorse upload model package file
[**PutApiV4ProjectsIdPackagesMlModelsModelVersionIdFilesPathFileNameAuthorize**](MlModelRegistryApi.md#PutApiV4ProjectsIdPackagesMlModelsModelVersionIdFilesPathFileNameAuthorize) | **Put** /api/v4/projects/{id}/packages/ml_models/{model_version_id}/files/(*path/){file_name}/authorize | Workhorse authorize model package file


# **GetApiV4ProjectsIdPackagesMlModelsModelVersionIdFilesPathFileName**
> GetApiV4ProjectsIdPackagesMlModelsModelVersionIdFilesPathFileName(ctx, id, modelVersionId, fileName, optional)
Download an ml_model package file

This feature was introduced in GitLab 16.8

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **modelVersionId** | **int32**| Model version id | 
  **fileName** | **string**| File name | 
 **optional** | ***MlModelRegistryApiGetApiV4ProjectsIdPackagesMlModelsModelVersionIdFilesPathFileNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MlModelRegistryApiGetApiV4ProjectsIdPackagesMlModelsModelVersionIdFilesPathFileNameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **path** | **optional.String**| File directory path | 
 **status** | **optional.String**| Package status | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdPackagesMlModelsModelVersionIdFilesPathFileName**
> PutApiV4ProjectsIdPackagesMlModelsModelVersionIdFilesPathFileName(ctx, id, modelVersionId, fileName, putApiV4ProjectsIdPackagesMlModelsModelVersionIdFilesPathFileName)
Workhorse upload model package file

Introduced in GitLab 16.8

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **modelVersionId** | **int32**| Model version id | 
  **fileName** | **string**| File name | 
  **putApiV4ProjectsIdPackagesMlModelsModelVersionIdFilesPathFileName** | [**PutApiV4ProjectsIdPackagesMlModelsModelVersionIdFilesPathFileName**](PutApiV4ProjectsIdPackagesMlModelsModelVersionIdFilesPathFileName.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdPackagesMlModelsModelVersionIdFilesPathFileNameAuthorize**
> PutApiV4ProjectsIdPackagesMlModelsModelVersionIdFilesPathFileNameAuthorize(ctx, id, modelVersionId, fileName, putApiV4ProjectsIdPackagesMlModelsModelVersionIdFilesPathFileNameAuthorize)
Workhorse authorize model package file

Introduced in GitLab 16.8

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **modelVersionId** | **int32**| Model version id | 
  **fileName** | **string**| File name | 
  **putApiV4ProjectsIdPackagesMlModelsModelVersionIdFilesPathFileNameAuthorize** | [**PutApiV4ProjectsIdPackagesMlModelsModelVersionIdFilesPathFileNameAuthorize**](PutApiV4ProjectsIdPackagesMlModelsModelVersionIdFilesPathFileNameAuthorize.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

