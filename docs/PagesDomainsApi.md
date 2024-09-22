# \PagesDomainsApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4ProjectsIdPagesDomains**](PagesDomainsApi.md#GetApiV4ProjectsIdPagesDomains) | **Get** /api/v4/projects/{id}/pages/domains | 


# **GetApiV4ProjectsIdPagesDomains**
> []ApiEntitiesPagesDomain GetApiV4ProjectsIdPagesDomains(ctx, id, optional)


Get all pages domains

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
 **optional** | ***PagesDomainsApiGetApiV4ProjectsIdPagesDomainsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PagesDomainsApiGetApiV4ProjectsIdPagesDomainsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesPagesDomain**](API_Entities_PagesDomain.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

