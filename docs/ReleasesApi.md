# \ReleasesApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4ProjectsIdReleasesTagName**](ReleasesApi.md#DeleteApiV4ProjectsIdReleasesTagName) | **Delete** /api/v4/projects/{id}/releases/{tag_name} | Delete a release
[**GetApiV4GroupsIdReleases**](ReleasesApi.md#GetApiV4GroupsIdReleases) | **Get** /api/v4/groups/{id}/releases | List group releases
[**GetApiV4ProjectsIdReleases**](ReleasesApi.md#GetApiV4ProjectsIdReleases) | **Get** /api/v4/projects/{id}/releases | List Releases
[**GetApiV4ProjectsIdReleasesPermalinkLatestSuffixPath**](ReleasesApi.md#GetApiV4ProjectsIdReleasesPermalinkLatestSuffixPath) | **Get** /api/v4/projects/{id}/releases/permalink/latest(/)(*suffix_path) | Get the latest project release
[**GetApiV4ProjectsIdReleasesTagName**](ReleasesApi.md#GetApiV4ProjectsIdReleasesTagName) | **Get** /api/v4/projects/{id}/releases/{tag_name} | Get a release by a tag name
[**GetApiV4ProjectsIdReleasesTagNameDownloadsdirectAssetPath**](ReleasesApi.md#GetApiV4ProjectsIdReleasesTagNameDownloadsdirectAssetPath) | **Get** /api/v4/projects/{id}/releases/{tag_name}/downloads/*direct_asset_path | Download a project release asset file
[**PostApiV4ProjectsIdReleases**](ReleasesApi.md#PostApiV4ProjectsIdReleases) | **Post** /api/v4/projects/{id}/releases | Create a release
[**PostApiV4ProjectsIdReleasesTagNameEvidence**](ReleasesApi.md#PostApiV4ProjectsIdReleasesTagNameEvidence) | **Post** /api/v4/projects/{id}/releases/{tag_name}/evidence | Collect release evidence
[**PutApiV4ProjectsIdReleasesTagName**](ReleasesApi.md#PutApiV4ProjectsIdReleasesTagName) | **Put** /api/v4/projects/{id}/releases/{tag_name} | Update a release


# **DeleteApiV4ProjectsIdReleasesTagName**
> ApiEntitiesRelease DeleteApiV4ProjectsIdReleasesTagName(ctx, id, tagName)
Delete a release

Delete a release. Deleting a release doesn't delete the associated tag. Maintainer level access to the project is required to delete a release. This feature was introduced in GitLab 11.7.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **tagName** | **string**| The Git tag the release is associated with | 

### Return type

[**ApiEntitiesRelease**](API_Entities_Release.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdReleases**
> []ApiEntitiesRelease GetApiV4GroupsIdReleases(ctx, id, optional)
List group releases

Returns a list of group releases.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the group owned by the authenticated user | 
 **optional** | ***ReleasesApiGetApiV4GroupsIdReleasesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReleasesApiGetApiV4GroupsIdReleasesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **optional.String**| The direction of the order. Either &#x60;desc&#x60; (default) for descending order or &#x60;asc&#x60; for ascending order | [default to desc]
 **simple** | **optional.Bool**| Return only limited fields for each release | [default to false]
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesRelease**](API_Entities_Release.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdReleases**
> []ApiEntitiesRelease GetApiV4ProjectsIdReleases(ctx, id, optional)
List Releases

Returns a paginated list of releases. This feature was introduced in GitLab 11.7.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
 **optional** | ***ReleasesApiGetApiV4ProjectsIdReleasesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReleasesApiGetApiV4ProjectsIdReleasesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **orderBy** | **optional.String**| The field to use as order. Either &#x60;released_at&#x60; (default) or &#x60;created_at&#x60; | [default to released_at]
 **sort** | **optional.String**| The direction of the order. Either &#x60;desc&#x60; (default) for descending order or &#x60;asc&#x60; for ascending order | [default to desc]
 **includeHtmlDescription** | **optional.Bool**| If &#x60;true&#x60;, a response includes HTML rendered markdown of the release description | 
 **updatedBefore** | **optional.Time**| Return releases updated before the specified datetime. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ | 
 **updatedAfter** | **optional.Time**| Return releases updated after the specified datetime. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ | 

### Return type

[**[]ApiEntitiesRelease**](API_Entities_Release.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdReleasesPermalinkLatestSuffixPath**
> GetApiV4ProjectsIdReleasesPermalinkLatestSuffixPath(ctx, id, suffixPath)
Get the latest project release

This feature was introduced in GitLab 15.4.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **suffixPath** | **string**| The path to be suffixed to the latest release | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdReleasesTagName**
> ApiEntitiesRelease GetApiV4ProjectsIdReleasesTagName(ctx, id, tagName, optional)
Get a release by a tag name

Gets a release for the given tag. This feature was introduced in GitLab 11.7.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **tagName** | **string**| The Git tag the release is associated with | 
 **optional** | ***ReleasesApiGetApiV4ProjectsIdReleasesTagNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReleasesApiGetApiV4ProjectsIdReleasesTagNameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeHtmlDescription** | **optional.Bool**| If &#x60;true&#x60;, a response includes HTML rendered markdown of the release description | 

### Return type

[**ApiEntitiesRelease**](API_Entities_Release.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdReleasesTagNameDownloadsdirectAssetPath**
> GetApiV4ProjectsIdReleasesTagNameDownloadsdirectAssetPath(ctx, id, tagName, directAssetPath)
Download a project release asset file

This feature was introduced in GitLab 15.4.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **tagName** | **string**| The Git tag the release is associated with | 
  **directAssetPath** | **string**| The path to the file to download, as specified when creating the release asset | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdReleases**
> ApiEntitiesRelease PostApiV4ProjectsIdReleases(ctx, id, postApiV4ProjectsIdReleases)
Create a release

Creates a release. Developer level access to the project is required to create a release. This feature was introduced in GitLab 11.7.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **postApiV4ProjectsIdReleases** | [**PostApiV4ProjectsIdReleases**](PostApiV4ProjectsIdReleases.md)|  | 

### Return type

[**ApiEntitiesRelease**](API_Entities_Release.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdReleasesTagNameEvidence**
> ApiEntitiesRelease PostApiV4ProjectsIdReleasesTagNameEvidence(ctx, tagName, id)
Collect release evidence

Creates an evidence for an existing Release. This feature was introduced in GitLab 12.10.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tagName** | **string**| The Git tag the release is associated with | 
  **id** | **int32**|  | 

### Return type

[**ApiEntitiesRelease**](API_Entities_Release.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdReleasesTagName**
> ApiEntitiesRelease PutApiV4ProjectsIdReleasesTagName(ctx, id, tagName, putApiV4ProjectsIdReleasesTagName)
Update a release

Updates a release. Developer level access to the project is required to update a release. This feature was introduced in GitLab 11.7.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **tagName** | **string**| The Git tag the release is associated with | 
  **putApiV4ProjectsIdReleasesTagName** | [**PutApiV4ProjectsIdReleasesTagName**](PutApiV4ProjectsIdReleasesTagName.md)|  | 

### Return type

[**ApiEntitiesRelease**](API_Entities_Release.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

