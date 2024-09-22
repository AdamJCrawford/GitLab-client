# \ProjectPackagesApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4ProjectsIdPackagesPackageId**](ProjectPackagesApi.md#DeleteApiV4ProjectsIdPackagesPackageId) | **Delete** /api/v4/projects/{id}/packages/{package_id} | Delete a project package
[**GetApiV4ProjectsIdPackages**](ProjectPackagesApi.md#GetApiV4ProjectsIdPackages) | **Get** /api/v4/projects/{id}/packages | Get a list of project packages
[**GetApiV4ProjectsIdPackagesPackageId**](ProjectPackagesApi.md#GetApiV4ProjectsIdPackagesPackageId) | **Get** /api/v4/projects/{id}/packages/{package_id} | Get a single project package
[**GetApiV4ProjectsIdPackagesPackageIdPipelines**](ProjectPackagesApi.md#GetApiV4ProjectsIdPackagesPackageIdPipelines) | **Get** /api/v4/projects/{id}/packages/{package_id}/pipelines | Get the pipelines for a single project package


# **DeleteApiV4ProjectsIdPackagesPackageId**
> DeleteApiV4ProjectsIdPackagesPackageId(ctx, id, packageId)
Delete a project package

This feature was introduced in GitLab 11.9

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **packageId** | **int32**| The ID of a package | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackages**
> []ApiEntitiesPackage GetApiV4ProjectsIdPackages(ctx, id, optional)
Get a list of project packages

This feature was introduced in GitLab 11.8

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
 **optional** | ***ProjectPackagesApiGetApiV4ProjectsIdPackagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectPackagesApiGetApiV4ProjectsIdPackagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **orderBy** | **optional.String**| Return packages ordered by &#x60;created_at&#x60;, &#x60;name&#x60;, &#x60;version&#x60; or &#x60;type&#x60; fields. | [default to created_at]
 **sort** | **optional.String**| Return packages sorted in &#x60;asc&#x60; or &#x60;desc&#x60; order. | [default to asc]
 **packageType** | **optional.String**| Return packages of a certain type | 
 **packageName** | **optional.String**| Return packages with this name | 
 **packageVersion** | **optional.String**| Return packages with this version | 
 **includeVersionless** | **optional.Bool**| Returns packages without a version | 
 **status** | **optional.String**| Return packages with specified status | 

### Return type

[**[]ApiEntitiesPackage**](API_Entities_Package.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesPackageId**
> ApiEntitiesPackage GetApiV4ProjectsIdPackagesPackageId(ctx, id, packageId)
Get a single project package

This feature was introduced in GitLab 11.9

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **packageId** | **int32**| The ID of a package | 

### Return type

[**ApiEntitiesPackage**](API_Entities_Package.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesPackageIdPipelines**
> ApiEntitiesPackagePipeline GetApiV4ProjectsIdPackagesPackageIdPipelines(ctx, id, packageId, optional)
Get the pipelines for a single project package

This feature was introduced in GitLab 16.1

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **packageId** | **int32**| The ID of a package | 
 **optional** | ***ProjectPackagesApiGetApiV4ProjectsIdPackagesPackageIdPipelinesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectPackagesApiGetApiV4ProjectsIdPackagesPackageIdPipelinesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **cursor** | **optional.String**| Cursor for obtaining the next set of records | 

### Return type

[**ApiEntitiesPackagePipeline**](API_Entities_Package_Pipeline.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

