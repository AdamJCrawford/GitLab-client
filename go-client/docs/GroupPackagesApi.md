# \GroupPackagesApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4GroupsIdPackages**](GroupPackagesApi.md#GetApiV4GroupsIdPackages) | **Get** /api/v4/groups/{id}/packages | List packages within a group


# **GetApiV4GroupsIdPackages**
> []ApiEntitiesPackage GetApiV4GroupsIdPackages(ctx, id, optional)
List packages within a group

Get a list of project packages at the group level. This feature was introduced in GitLab 12.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID or URL-encoded path of the group | 
 **optional** | ***GroupPackagesApiGetApiV4GroupsIdPackagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupPackagesApiGetApiV4GroupsIdPackagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **excludeSubgroups** | **optional.Bool**| Determines if subgroups should be excluded | [default to false]
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

