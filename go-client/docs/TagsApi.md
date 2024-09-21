# \TagsApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4ProjectsIdRepositoryTagsTagName**](TagsApi.md#DeleteApiV4ProjectsIdRepositoryTagsTagName) | **Delete** /api/v4/projects/{id}/repository/tags/{tag_name} | 
[**GetApiV4ProjectsIdRepositoryTags**](TagsApi.md#GetApiV4ProjectsIdRepositoryTags) | **Get** /api/v4/projects/{id}/repository/tags | 
[**GetApiV4ProjectsIdRepositoryTagsTagName**](TagsApi.md#GetApiV4ProjectsIdRepositoryTagsTagName) | **Get** /api/v4/projects/{id}/repository/tags/{tag_name} | 
[**GetApiV4ProjectsIdRepositoryTagsTagNameSignature**](TagsApi.md#GetApiV4ProjectsIdRepositoryTagsTagNameSignature) | **Get** /api/v4/projects/{id}/repository/tags/{tag_name}/signature | 
[**PostApiV4ProjectsIdRepositoryTags**](TagsApi.md#PostApiV4ProjectsIdRepositoryTags) | **Post** /api/v4/projects/{id}/repository/tags | 


# **DeleteApiV4ProjectsIdRepositoryTagsTagName**
> DeleteApiV4ProjectsIdRepositoryTagsTagName(ctx, id, tagName)


Delete a repository tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **tagName** | **string**| The name of the tag | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdRepositoryTags**
> []ApiEntitiesTag GetApiV4ProjectsIdRepositoryTags(ctx, id, optional)


Get a project repository tags

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
 **optional** | ***TagsApiGetApiV4ProjectsIdRepositoryTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TagsApiGetApiV4ProjectsIdRepositoryTagsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **optional.String**| Return tags sorted in updated by &#x60;asc&#x60; or &#x60;desc&#x60; order. | [default to desc]
 **orderBy** | **optional.String**| Return tags ordered by &#x60;name&#x60;, &#x60;updated&#x60;, &#x60;version&#x60; fields. | [default to updated]
 **search** | **optional.String**| Return list of tags matching the search criteria | 
 **pageToken** | **optional.String**| Name of tag to start the paginaition from | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesTag**](API_Entities_Tag.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdRepositoryTagsTagName**
> ApiEntitiesTag GetApiV4ProjectsIdRepositoryTagsTagName(ctx, id, tagName)


Get a single repository tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **tagName** | **string**| The name of the tag | 

### Return type

[**ApiEntitiesTag**](API_Entities_Tag.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdRepositoryTagsTagNameSignature**
> ApiEntitiesTagSignature GetApiV4ProjectsIdRepositoryTagsTagNameSignature(ctx, id, tagName)


Get a tag's signature

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **tagName** | **string**| The name of the tag | 

### Return type

[**ApiEntitiesTagSignature**](API_Entities_TagSignature.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdRepositoryTags**
> ApiEntitiesTag PostApiV4ProjectsIdRepositoryTags(ctx, id, postApiV4ProjectsIdRepositoryTags)


Create a new repository tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **postApiV4ProjectsIdRepositoryTags** | [**PostApiV4ProjectsIdRepositoryTags**](PostApiV4ProjectsIdRepositoryTags.md)|  | 

### Return type

[**ApiEntitiesTag**](API_Entities_Tag.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

