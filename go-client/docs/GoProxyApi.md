# \GoProxyApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4ProjectsIdPackagesGomoduleNamevList**](GoProxyApi.md#GetApiV4ProjectsIdPackagesGomoduleNamevList) | **Get** /api/v4/projects/{id}/packages/go/*module_name/@v/list | List
[**GetApiV4ProjectsIdPackagesGomoduleNamevModuleVersionInfo**](GoProxyApi.md#GetApiV4ProjectsIdPackagesGomoduleNamevModuleVersionInfo) | **Get** /api/v4/projects/{id}/packages/go/*module_name/@v/{module_version}.info | Version metadata
[**GetApiV4ProjectsIdPackagesGomoduleNamevModuleVersionMod**](GoProxyApi.md#GetApiV4ProjectsIdPackagesGomoduleNamevModuleVersionMod) | **Get** /api/v4/projects/{id}/packages/go/*module_name/@v/{module_version}.mod | Download module file
[**GetApiV4ProjectsIdPackagesGomoduleNamevModuleVersionZip**](GoProxyApi.md#GetApiV4ProjectsIdPackagesGomoduleNamevModuleVersionZip) | **Get** /api/v4/projects/{id}/packages/go/*module_name/@v/{module_version}.zip | Download module source


# **GetApiV4ProjectsIdPackagesGomoduleNamevList**
> GetApiV4ProjectsIdPackagesGomoduleNamevList(ctx, id, moduleName)
List

Get all tagged versions for a given Go module.See `go help goproxy`, GET $GOPROXY/<module>/@v/list. This feature was introduced in GitLab 13.1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID or full path of a project | 
  **moduleName** | **string**| The name of the Go module | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesGomoduleNamevModuleVersionInfo**
> ApiEntitiesGoModuleVersion GetApiV4ProjectsIdPackagesGomoduleNamevModuleVersionInfo(ctx, id, moduleName, moduleVersion)
Version metadata

Get all tagged versions for a given Go module.See `go help goproxy`, GET $GOPROXY/<module>/@v/<version>.info. This feature was introduced in GitLab 13.1

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID or full path of a project | 
  **moduleName** | **string**| The name of the Go module | 
  **moduleVersion** | **string**| The version of the Go module | 

### Return type

[**ApiEntitiesGoModuleVersion**](API_Entities_GoModuleVersion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesGomoduleNamevModuleVersionMod**
> GetApiV4ProjectsIdPackagesGomoduleNamevModuleVersionMod(ctx, id, moduleName, moduleVersion)
Download module file

Get the module file of a given module version.See `go help goproxy`, GET $GOPROXY/<module>/@v/<version>.mod. This feature was introduced in GitLab 13.1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID or full path of a project | 
  **moduleName** | **string**| The name of the Go module | 
  **moduleVersion** | **string**| The version of the Go module | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesGomoduleNamevModuleVersionZip**
> GetApiV4ProjectsIdPackagesGomoduleNamevModuleVersionZip(ctx, id, moduleName, moduleVersion)
Download module source

Get a zip of the source of the given module version.See `go help goproxy`, GET $GOPROXY/<module>/@v/<version>.zip. This feature was introduced in GitLab 13.1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID or full path of a project | 
  **moduleName** | **string**| The name of the Go module | 
  **moduleVersion** | **string**| The version of the Go module | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

