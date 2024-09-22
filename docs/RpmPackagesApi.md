# \RpmPackagesApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4ProjectsIdPackagesRpmRepodatafileName**](RpmPackagesApi.md#GetApiV4ProjectsIdPackagesRpmRepodatafileName) | **Get** /api/v4/projects/{id}/packages/rpm/repodata/*file_name | Download repository metadata files
[**GetApiV4ProjectsIdPackagesRpmpackageFileIdfileName**](RpmPackagesApi.md#GetApiV4ProjectsIdPackagesRpmpackageFileIdfileName) | **Get** /api/v4/projects/{id}/packages/rpm/*package_file_id/*file_name | Download RPM package files
[**PostApiV4ProjectsIdPackagesRpm**](RpmPackagesApi.md#PostApiV4ProjectsIdPackagesRpm) | **Post** /api/v4/projects/{id}/packages/rpm | Upload a RPM package
[**PostApiV4ProjectsIdPackagesRpmAuthorize**](RpmPackagesApi.md#PostApiV4ProjectsIdPackagesRpmAuthorize) | **Post** /api/v4/projects/{id}/packages/rpm/authorize | Authorize package upload from workhorse


# **GetApiV4ProjectsIdPackagesRpmRepodatafileName**
> GetApiV4ProjectsIdPackagesRpmRepodatafileName(ctx, id, fileName)
Download repository metadata files

This feature was introduced in GitLab 15.7

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **fileName** | **string**| Repository metadata file name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesRpmpackageFileIdfileName**
> GetApiV4ProjectsIdPackagesRpmpackageFileIdfileName(ctx, id, packageFileId, fileName)
Download RPM package files

This feature was introduced in GitLab 15.7

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **packageFileId** | **int32**| RPM package file id | 
  **fileName** | **string**| RPM package file name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdPackagesRpm**
> PostApiV4ProjectsIdPackagesRpm(ctx, id)
Upload a RPM package

This feature was introduced in GitLab 15.7

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

# **PostApiV4ProjectsIdPackagesRpmAuthorize**
> PostApiV4ProjectsIdPackagesRpmAuthorize(ctx, id)
Authorize package upload from workhorse

This feature was introduced in GitLab 15.7

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

