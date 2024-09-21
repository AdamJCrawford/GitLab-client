# \GenericPackagesApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4ProjectsIdPackagesGenericPackageNamepackageVersionPathFileName**](GenericPackagesApi.md#GetApiV4ProjectsIdPackagesGenericPackageNamepackageVersionPathFileName) | **Get** /api/v4/projects/{id}/packages/generic/{package_name}/*package_version/(*path/){file_name} | Download package file
[**PutApiV4ProjectsIdPackagesGenericPackageNamepackageVersionPathFileName**](GenericPackagesApi.md#PutApiV4ProjectsIdPackagesGenericPackageNamepackageVersionPathFileName) | **Put** /api/v4/projects/{id}/packages/generic/{package_name}/*package_version/(*path/){file_name} | Upload package file
[**PutApiV4ProjectsIdPackagesGenericPackageNamepackageVersionPathFileNameAuthorize**](GenericPackagesApi.md#PutApiV4ProjectsIdPackagesGenericPackageNamepackageVersionPathFileNameAuthorize) | **Put** /api/v4/projects/{id}/packages/generic/{package_name}/*package_version/(*path/){file_name}/authorize | Workhorse authorize generic package file


# **GetApiV4ProjectsIdPackagesGenericPackageNamepackageVersionPathFileName**
> GetApiV4ProjectsIdPackagesGenericPackageNamepackageVersionPathFileName(ctx, id, packageName, packageVersion, fileName, optional)
Download package file

This feature was introduced in GitLab 13.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **packageName** | **string**| Package name | 
  **packageVersion** | **string**| Package version | 
  **fileName** | **string**| Package file name | 
 **optional** | ***GenericPackagesApiGetApiV4ProjectsIdPackagesGenericPackageNamepackageVersionPathFileNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GenericPackagesApiGetApiV4ProjectsIdPackagesGenericPackageNamepackageVersionPathFileNameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **path** | **optional.String**| File directory path | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdPackagesGenericPackageNamepackageVersionPathFileName**
> PutApiV4ProjectsIdPackagesGenericPackageNamepackageVersionPathFileName(ctx, id, packageName, fileName, putApiV4ProjectsIdPackagesGenericPackageNamepackageVersionPathFileName)
Upload package file

This feature was introduced in GitLab 13.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **packageName** | **string**| Package name | 
  **fileName** | **string**| Package file name | 
  **putApiV4ProjectsIdPackagesGenericPackageNamepackageVersionPathFileName** | [**PutApiV4ProjectsIdPackagesGenericPackageNamepackageVersionPathFileName**](PutApiV4ProjectsIdPackagesGenericPackageNamepackageVersionPathFileName.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdPackagesGenericPackageNamepackageVersionPathFileNameAuthorize**
> PutApiV4ProjectsIdPackagesGenericPackageNamepackageVersionPathFileNameAuthorize(ctx, id, packageName, fileName, putApiV4ProjectsIdPackagesGenericPackageNamepackageVersionPathFileNameAuthorize)
Workhorse authorize generic package file

This feature was introduced in GitLab 13.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **packageName** | **string**| Package name | 
  **fileName** | **string**| Package file name | 
  **putApiV4ProjectsIdPackagesGenericPackageNamepackageVersionPathFileNameAuthorize** | [**PutApiV4ProjectsIdPackagesGenericPackageNamepackageVersionPathFileNameAuthorize**](PutApiV4ProjectsIdPackagesGenericPackageNamepackageVersionPathFileNameAuthorize.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

