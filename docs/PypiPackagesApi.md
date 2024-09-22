# \PypiPackagesApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4GroupsIdPackagesPypiFilesSha256fileIdentifier**](PypiPackagesApi.md#GetApiV4GroupsIdPackagesPypiFilesSha256fileIdentifier) | **Get** /api/v4/groups/{id}/-/packages/pypi/files/{sha256}/*file_identifier | Download a package file from a group
[**GetApiV4GroupsIdPackagesPypiSimple**](PypiPackagesApi.md#GetApiV4GroupsIdPackagesPypiSimple) | **Get** /api/v4/groups/{id}/-/packages/pypi/simple | The PyPi Simple Group Index Endpoint
[**GetApiV4GroupsIdPackagesPypiSimplepackageName**](PypiPackagesApi.md#GetApiV4GroupsIdPackagesPypiSimplepackageName) | **Get** /api/v4/groups/{id}/-/packages/pypi/simple/*package_name | The PyPi Simple Group Package Endpoint
[**GetApiV4ProjectsIdPackagesPypiFilesSha256fileIdentifier**](PypiPackagesApi.md#GetApiV4ProjectsIdPackagesPypiFilesSha256fileIdentifier) | **Get** /api/v4/projects/{id}/packages/pypi/files/{sha256}/*file_identifier | The PyPi package download endpoint
[**GetApiV4ProjectsIdPackagesPypiSimple**](PypiPackagesApi.md#GetApiV4ProjectsIdPackagesPypiSimple) | **Get** /api/v4/projects/{id}/packages/pypi/simple | The PyPi Simple Project Index Endpoint
[**GetApiV4ProjectsIdPackagesPypiSimplepackageName**](PypiPackagesApi.md#GetApiV4ProjectsIdPackagesPypiSimplepackageName) | **Get** /api/v4/projects/{id}/packages/pypi/simple/*package_name | The PyPi Simple Project Package Endpoint
[**PostApiV4ProjectsIdPackagesPypi**](PypiPackagesApi.md#PostApiV4ProjectsIdPackagesPypi) | **Post** /api/v4/projects/{id}/packages/pypi | The PyPi Package upload endpoint
[**PostApiV4ProjectsIdPackagesPypiAuthorize**](PypiPackagesApi.md#PostApiV4ProjectsIdPackagesPypiAuthorize) | **Post** /api/v4/projects/{id}/packages/pypi/authorize | Authorize the PyPi package upload from workhorse


# **GetApiV4GroupsIdPackagesPypiFilesSha256fileIdentifier**
> GetApiV4GroupsIdPackagesPypiFilesSha256fileIdentifier(ctx, id, fileIdentifier, sha256)
Download a package file from a group

This feature was introduced in GitLab 13.12

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID or full path of the group. | 
  **fileIdentifier** | **string**| The PyPi package file identifier | 
  **sha256** | **string**| The PyPi package sha256 check sum | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdPackagesPypiSimple**
> GetApiV4GroupsIdPackagesPypiSimple(ctx, id)
The PyPi Simple Group Index Endpoint

This feature was introduced in GitLab 15.1

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID or full path of the group. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdPackagesPypiSimplepackageName**
> GetApiV4GroupsIdPackagesPypiSimplepackageName(ctx, id, packageName)
The PyPi Simple Group Package Endpoint

This feature was introduced in GitLab 12.10

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID or full path of the group. | 
  **packageName** | **string**| The PyPi package name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesPypiFilesSha256fileIdentifier**
> GetApiV4ProjectsIdPackagesPypiFilesSha256fileIdentifier(ctx, id, fileIdentifier, sha256)
The PyPi package download endpoint

This feature was introduced in GitLab 12.10

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **fileIdentifier** | **string**| The PyPi package file identifier | 
  **sha256** | **string**| The PyPi package sha256 check sum | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesPypiSimple**
> GetApiV4ProjectsIdPackagesPypiSimple(ctx, id)
The PyPi Simple Project Index Endpoint

This feature was introduced in GitLab 15.1

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
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesPypiSimplepackageName**
> GetApiV4ProjectsIdPackagesPypiSimplepackageName(ctx, id, packageName)
The PyPi Simple Project Package Endpoint

This feature was introduced in GitLab 12.10

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **packageName** | **string**| The PyPi package name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdPackagesPypi**
> PostApiV4ProjectsIdPackagesPypi(ctx, id, postApiV4ProjectsIdPackagesPypi)
The PyPi Package upload endpoint

This feature was introduced in GitLab 12.10

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **postApiV4ProjectsIdPackagesPypi** | [**PostApiV4ProjectsIdPackagesPypi**](PostApiV4ProjectsIdPackagesPypi.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdPackagesPypiAuthorize**
> PostApiV4ProjectsIdPackagesPypiAuthorize(ctx, id)
Authorize the PyPi package upload from workhorse

This feature was introduced in GitLab 12.10

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

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

