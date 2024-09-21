# \SecureFilesApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4ProjectsIdSecureFilesSecureFileId**](SecureFilesApi.md#DeleteApiV4ProjectsIdSecureFilesSecureFileId) | **Delete** /api/v4/projects/{id}/secure_files/{secure_file_id} | 
[**GetApiV4ProjectsIdSecureFiles**](SecureFilesApi.md#GetApiV4ProjectsIdSecureFiles) | **Get** /api/v4/projects/{id}/secure_files | 
[**GetApiV4ProjectsIdSecureFilesSecureFileId**](SecureFilesApi.md#GetApiV4ProjectsIdSecureFilesSecureFileId) | **Get** /api/v4/projects/{id}/secure_files/{secure_file_id} | 
[**GetApiV4ProjectsIdSecureFilesSecureFileIdDownload**](SecureFilesApi.md#GetApiV4ProjectsIdSecureFilesSecureFileIdDownload) | **Get** /api/v4/projects/{id}/secure_files/{secure_file_id}/download | 
[**PostApiV4ProjectsIdSecureFiles**](SecureFilesApi.md#PostApiV4ProjectsIdSecureFiles) | **Post** /api/v4/projects/{id}/secure_files | 


# **DeleteApiV4ProjectsIdSecureFilesSecureFileId**
> DeleteApiV4ProjectsIdSecureFilesSecureFileId(ctx, id, secureFileId)


Remove a secure file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the         authenticated user | 
  **secureFileId** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdSecureFiles**
> ApiEntitiesCiSecureFile GetApiV4ProjectsIdSecureFiles(ctx, id, optional)


Get list of secure files in a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the         authenticated user | 
 **optional** | ***SecureFilesApiGetApiV4ProjectsIdSecureFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecureFilesApiGetApiV4ProjectsIdSecureFilesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**ApiEntitiesCiSecureFile**](API_Entities_Ci_SecureFile.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdSecureFilesSecureFileId**
> ApiEntitiesCiSecureFile GetApiV4ProjectsIdSecureFilesSecureFileId(ctx, id, secureFileId)


Get the details of a specific secure file in a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of a secure file | 
  **secureFileId** | **int32**|  | 

### Return type

[**ApiEntitiesCiSecureFile**](API_Entities_Ci_SecureFile.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdSecureFilesSecureFileIdDownload**
> GetApiV4ProjectsIdSecureFilesSecureFileIdDownload(ctx, id, secureFileId)


Download secure file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the         authenticated user | 
  **secureFileId** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdSecureFiles**
> ApiEntitiesCiSecureFile PostApiV4ProjectsIdSecureFiles(ctx, id, postApiV4ProjectsIdSecureFiles)


Create a secure file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the         authenticated user | 
  **postApiV4ProjectsIdSecureFiles** | [**PostApiV4ProjectsIdSecureFiles**](PostApiV4ProjectsIdSecureFiles.md)|  | 

### Return type

[**ApiEntitiesCiSecureFile**](API_Entities_Ci_SecureFile.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

