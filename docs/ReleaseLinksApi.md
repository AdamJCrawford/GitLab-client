# \ReleaseLinksApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4ProjectsIdReleasesTagNameAssetsLinksLinkId**](ReleaseLinksApi.md#DeleteApiV4ProjectsIdReleasesTagNameAssetsLinksLinkId) | **Delete** /api/v4/projects/{id}/releases/{tag_name}/assets/links/{link_id} | Delete a release link
[**GetApiV4ProjectsIdReleasesTagNameAssetsLinks**](ReleaseLinksApi.md#GetApiV4ProjectsIdReleasesTagNameAssetsLinks) | **Get** /api/v4/projects/{id}/releases/{tag_name}/assets/links | List links of a release
[**GetApiV4ProjectsIdReleasesTagNameAssetsLinksLinkId**](ReleaseLinksApi.md#GetApiV4ProjectsIdReleasesTagNameAssetsLinksLinkId) | **Get** /api/v4/projects/{id}/releases/{tag_name}/assets/links/{link_id} | Get a release link
[**PostApiV4ProjectsIdReleasesTagNameAssetsLinks**](ReleaseLinksApi.md#PostApiV4ProjectsIdReleasesTagNameAssetsLinks) | **Post** /api/v4/projects/{id}/releases/{tag_name}/assets/links | Create a release link
[**PutApiV4ProjectsIdReleasesTagNameAssetsLinksLinkId**](ReleaseLinksApi.md#PutApiV4ProjectsIdReleasesTagNameAssetsLinksLinkId) | **Put** /api/v4/projects/{id}/releases/{tag_name}/assets/links/{link_id} | Update a release link


# **DeleteApiV4ProjectsIdReleasesTagNameAssetsLinksLinkId**
> ApiEntitiesReleasesLink DeleteApiV4ProjectsIdReleasesTagNameAssetsLinksLinkId(ctx, id, tagName, linkId)
Delete a release link

Deletes an asset as a link from a release. This feature was introduced in GitLab 11.7.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **tagName** | **string**| The tag associated with the release | 
  **linkId** | **int32**| The ID of the link | 

### Return type

[**ApiEntitiesReleasesLink**](API_Entities_Releases_Link.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdReleasesTagNameAssetsLinks**
> []ApiEntitiesReleasesLink GetApiV4ProjectsIdReleasesTagNameAssetsLinks(ctx, id, tagName, optional)
List links of a release

Get assets as links from a release. This feature was introduced in GitLab 11.7.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **tagName** | **string**| The tag associated with the release | 
 **optional** | ***ReleaseLinksApiGetApiV4ProjectsIdReleasesTagNameAssetsLinksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReleaseLinksApiGetApiV4ProjectsIdReleasesTagNameAssetsLinksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesReleasesLink**](API_Entities_Releases_Link.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdReleasesTagNameAssetsLinksLinkId**
> ApiEntitiesReleasesLink GetApiV4ProjectsIdReleasesTagNameAssetsLinksLinkId(ctx, id, tagName, linkId)
Get a release link

Get an asset as a link from a release. This feature was introduced in GitLab 11.7.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **tagName** | **string**| The tag associated with the release | 
  **linkId** | **int32**| The ID of the link | 

### Return type

[**ApiEntitiesReleasesLink**](API_Entities_Releases_Link.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdReleasesTagNameAssetsLinks**
> ApiEntitiesReleasesLink PostApiV4ProjectsIdReleasesTagNameAssetsLinks(ctx, id, tagName, postApiV4ProjectsIdReleasesTagNameAssetsLinks)
Create a release link

Create an asset as a link from a release. This feature was introduced in GitLab 11.7.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **tagName** | **string**| The tag associated with the release | 
  **postApiV4ProjectsIdReleasesTagNameAssetsLinks** | [**PostApiV4ProjectsIdReleasesTagNameAssetsLinks**](PostApiV4ProjectsIdReleasesTagNameAssetsLinks.md)|  | 

### Return type

[**ApiEntitiesReleasesLink**](API_Entities_Releases_Link.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdReleasesTagNameAssetsLinksLinkId**
> ApiEntitiesReleasesLink PutApiV4ProjectsIdReleasesTagNameAssetsLinksLinkId(ctx, id, tagName, linkId, putApiV4ProjectsIdReleasesTagNameAssetsLinksLinkId)
Update a release link

Update an asset as a link from a release. This feature was introduced in GitLab 11.7.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **tagName** | **string**| The tag associated with the release | 
  **linkId** | **int32**| The ID of the link | 
  **putApiV4ProjectsIdReleasesTagNameAssetsLinksLinkId** | [**PutApiV4ProjectsIdReleasesTagNameAssetsLinksLinkId**](PutApiV4ProjectsIdReleasesTagNameAssetsLinksLinkId.md)|  | 

### Return type

[**ApiEntitiesReleasesLink**](API_Entities_Releases_Link.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

