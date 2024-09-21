# \MavenPackagesApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4GroupsIdPackagesMavenpathFileName**](MavenPackagesApi.md#GetApiV4GroupsIdPackagesMavenpathFileName) | **Get** /api/v4/groups/{id}/-/packages/maven/*path/{file_name} | Download the maven package file at a group level
[**GetApiV4PackagesMavenpathFileName**](MavenPackagesApi.md#GetApiV4PackagesMavenpathFileName) | **Get** /api/v4/packages/maven/*path/{file_name} | Download the maven package file at instance level
[**GetApiV4ProjectsIdPackagesMavenpathFileName**](MavenPackagesApi.md#GetApiV4ProjectsIdPackagesMavenpathFileName) | **Get** /api/v4/projects/{id}/packages/maven/*path/{file_name} | Download the maven package file at a project level
[**PutApiV4ProjectsIdPackagesMavenpathFileName**](MavenPackagesApi.md#PutApiV4ProjectsIdPackagesMavenpathFileName) | **Put** /api/v4/projects/{id}/packages/maven/*path/{file_name} | Upload the maven package file
[**PutApiV4ProjectsIdPackagesMavenpathFileNameAuthorize**](MavenPackagesApi.md#PutApiV4ProjectsIdPackagesMavenpathFileNameAuthorize) | **Put** /api/v4/projects/{id}/packages/maven/*path/{file_name}/authorize | Workhorse authorize the maven package file upload


# **GetApiV4GroupsIdPackagesMavenpathFileName**
> GetApiV4GroupsIdPackagesMavenpathFileName(ctx, id, path, fileName)
Download the maven package file at a group level

This feature was introduced in GitLab 11.7

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the group | 
  **path** | **string**| Package path | 
  **fileName** | **string**| Package file name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4PackagesMavenpathFileName**
> GetApiV4PackagesMavenpathFileName(ctx, path, fileName)
Download the maven package file at instance level

This feature was introduced in GitLab 11.6

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| Package path | 
  **fileName** | **string**| Package file name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesMavenpathFileName**
> GetApiV4ProjectsIdPackagesMavenpathFileName(ctx, id, path, fileName)
Download the maven package file at a project level

This feature was introduced in GitLab 11.3

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **path** | **string**| Package path | 
  **fileName** | **string**| Package file name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdPackagesMavenpathFileName**
> PutApiV4ProjectsIdPackagesMavenpathFileName(ctx, id, fileName, putApiV4ProjectsIdPackagesMavenpathFileName)
Upload the maven package file

This feature was introduced in GitLab 11.3

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **fileName** | **string**| Package file name | 
  **putApiV4ProjectsIdPackagesMavenpathFileName** | [**PutApiV4ProjectsIdPackagesMavenpathFileName**](PutApiV4ProjectsIdPackagesMavenpathFileName.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdPackagesMavenpathFileNameAuthorize**
> PutApiV4ProjectsIdPackagesMavenpathFileNameAuthorize(ctx, id, fileName, putApiV4ProjectsIdPackagesMavenpathFileNameAuthorize)
Workhorse authorize the maven package file upload

This feature was introduced in GitLab 11.3

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **fileName** | **string**| Package file name | 
  **putApiV4ProjectsIdPackagesMavenpathFileNameAuthorize** | [**PutApiV4ProjectsIdPackagesMavenpathFileNameAuthorize**](PutApiV4ProjectsIdPackagesMavenpathFileNameAuthorize.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

