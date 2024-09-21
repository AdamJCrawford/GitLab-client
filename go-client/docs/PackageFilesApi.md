# \PackageFilesApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4ProjectsIdPackagesPackageIdPackageFilesPackageFileId**](PackageFilesApi.md#DeleteApiV4ProjectsIdPackagesPackageIdPackageFilesPackageFileId) | **Delete** /api/v4/projects/{id}/packages/{package_id}/package_files/{package_file_id} | Delete a package file
[**GetApiV4ProjectsIdPackagesPackageIdPackageFiles**](PackageFilesApi.md#GetApiV4ProjectsIdPackagesPackageIdPackageFiles) | **Get** /api/v4/projects/{id}/packages/{package_id}/package_files | List package files


# **DeleteApiV4ProjectsIdPackagesPackageIdPackageFilesPackageFileId**
> DeleteApiV4ProjectsIdPackagesPackageIdPackageFilesPackageFileId(ctx, id, packageId, packageFileId)
Delete a package file

This feature was introduced in GitLab 13.12

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID or URL-encoded path of the project | 
  **packageId** | **int32**| ID of a package | 
  **packageFileId** | **int32**| ID of a package file | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesPackageIdPackageFiles**
> []ApiEntitiesPackageFile GetApiV4ProjectsIdPackagesPackageIdPackageFiles(ctx, id, packageId, optional)
List package files

Get a list of package files of a single package

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID or URL-encoded path of the project | 
  **packageId** | **int32**| ID of a package | 
 **optional** | ***PackageFilesApiGetApiV4ProjectsIdPackagesPackageIdPackageFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PackageFilesApiGetApiV4ProjectsIdPackagesPackageIdPackageFilesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesPackageFile**](API_Entities_PackageFile.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

