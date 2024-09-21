# \PagesApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4ProjectsIdPages**](PagesApi.md#DeleteApiV4ProjectsIdPages) | **Delete** /api/v4/projects/{id}/pages | Unpublish pages
[**GetApiV4PagesDomains**](PagesApi.md#GetApiV4PagesDomains) | **Get** /api/v4/pages/domains | 
[**GetApiV4ProjectsIdPages**](PagesApi.md#GetApiV4ProjectsIdPages) | **Get** /api/v4/projects/{id}/pages | Get pages settings
[**PatchApiV4ProjectsIdPages**](PagesApi.md#PatchApiV4ProjectsIdPages) | **Patch** /api/v4/projects/{id}/pages | Update pages settings


# **DeleteApiV4ProjectsIdPages**
> DeleteApiV4ProjectsIdPages(ctx, id)
Unpublish pages

Remove pages. The user must have administrator access. This feature was introduced in GitLab 12.6

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4PagesDomains**
> ApiEntitiesPagesDomainBasic GetApiV4PagesDomains(ctx, optional)


Get all pages domains

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PagesApiGetApiV4PagesDomainsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PagesApiGetApiV4PagesDomainsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**ApiEntitiesPagesDomainBasic**](API_Entities_PagesDomainBasic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPages**
> GetApiV4ProjectsIdPages(ctx, id)
Get pages settings

Get pages URL and other settings. This feature was introduced in Gitlab 16.8

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchApiV4ProjectsIdPages**
> PatchApiV4ProjectsIdPages(ctx, id, patchApiV4ProjectsIdPages)
Update pages settings

Update page settings for a project. User must have administrative access.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **patchApiV4ProjectsIdPages** | [**PatchApiV4ProjectsIdPages**](PatchApiV4ProjectsIdPages.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

