# \WikisApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4GroupsIdWikisSlug**](WikisApi.md#DeleteApiV4GroupsIdWikisSlug) | **Delete** /api/v4/groups/{id}/wikis/{slug} | 
[**DeleteApiV4ProjectsIdWikisSlug**](WikisApi.md#DeleteApiV4ProjectsIdWikisSlug) | **Delete** /api/v4/projects/{id}/wikis/{slug} | 
[**GetApiV4GroupsIdWikis**](WikisApi.md#GetApiV4GroupsIdWikis) | **Get** /api/v4/groups/{id}/wikis | 
[**GetApiV4GroupsIdWikisSlug**](WikisApi.md#GetApiV4GroupsIdWikisSlug) | **Get** /api/v4/groups/{id}/wikis/{slug} | 
[**GetApiV4ProjectsIdWikis**](WikisApi.md#GetApiV4ProjectsIdWikis) | **Get** /api/v4/projects/{id}/wikis | 
[**GetApiV4ProjectsIdWikisSlug**](WikisApi.md#GetApiV4ProjectsIdWikisSlug) | **Get** /api/v4/projects/{id}/wikis/{slug} | 
[**PostApiV4GroupsIdWikis**](WikisApi.md#PostApiV4GroupsIdWikis) | **Post** /api/v4/groups/{id}/wikis | 
[**PostApiV4GroupsIdWikisAttachments**](WikisApi.md#PostApiV4GroupsIdWikisAttachments) | **Post** /api/v4/groups/{id}/wikis/attachments | Upload an attachment to the wiki repository
[**PostApiV4ProjectsIdWikis**](WikisApi.md#PostApiV4ProjectsIdWikis) | **Post** /api/v4/projects/{id}/wikis | 
[**PostApiV4ProjectsIdWikisAttachments**](WikisApi.md#PostApiV4ProjectsIdWikisAttachments) | **Post** /api/v4/projects/{id}/wikis/attachments | Upload an attachment to the wiki repository
[**PutApiV4GroupsIdWikisSlug**](WikisApi.md#PutApiV4GroupsIdWikisSlug) | **Put** /api/v4/groups/{id}/wikis/{slug} | 
[**PutApiV4ProjectsIdWikisSlug**](WikisApi.md#PutApiV4ProjectsIdWikisSlug) | **Put** /api/v4/projects/{id}/wikis/{slug} | 


# **DeleteApiV4GroupsIdWikisSlug**
> DeleteApiV4GroupsIdWikisSlug(ctx, slug, id)


Delete a wiki page

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **slug** | **string**| The slug of a wiki page | 
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdWikisSlug**
> DeleteApiV4ProjectsIdWikisSlug(ctx, slug, id)


Delete a wiki page

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **slug** | **string**| The slug of a wiki page | 
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdWikis**
> []ApiEntitiesWikiPageBasic GetApiV4GroupsIdWikis(ctx, id, optional)


Get a list of wiki pages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***WikisApiGetApiV4GroupsIdWikisOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WikisApiGetApiV4GroupsIdWikisOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withContent** | **optional.Bool**| Include pages&#39; content | [default to false]

### Return type

[**[]ApiEntitiesWikiPageBasic**](API_Entities_WikiPageBasic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdWikisSlug**
> ApiEntitiesWikiPage GetApiV4GroupsIdWikisSlug(ctx, slug, id, optional)


Get a wiki page

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **slug** | **string**| The slug of a wiki page | 
  **id** | **int32**|  | 
 **optional** | ***WikisApiGetApiV4GroupsIdWikisSlugOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WikisApiGetApiV4GroupsIdWikisSlugOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **optional.String**| The version hash of a wiki page | 
 **renderHtml** | **optional.Bool**| Render content to HTML | [default to false]

### Return type

[**ApiEntitiesWikiPage**](API_Entities_WikiPage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdWikis**
> []ApiEntitiesWikiPageBasic GetApiV4ProjectsIdWikis(ctx, id, optional)


Get a list of wiki pages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***WikisApiGetApiV4ProjectsIdWikisOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WikisApiGetApiV4ProjectsIdWikisOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withContent** | **optional.Bool**| Include pages&#39; content | [default to false]

### Return type

[**[]ApiEntitiesWikiPageBasic**](API_Entities_WikiPageBasic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdWikisSlug**
> ApiEntitiesWikiPage GetApiV4ProjectsIdWikisSlug(ctx, slug, id, optional)


Get a wiki page

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **slug** | **string**| The slug of a wiki page | 
  **id** | **int32**|  | 
 **optional** | ***WikisApiGetApiV4ProjectsIdWikisSlugOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WikisApiGetApiV4ProjectsIdWikisSlugOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **optional.String**| The version hash of a wiki page | 
 **renderHtml** | **optional.Bool**| Render content to HTML | [default to false]

### Return type

[**ApiEntitiesWikiPage**](API_Entities_WikiPage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4GroupsIdWikis**
> ApiEntitiesWikiPage PostApiV4GroupsIdWikis(ctx, id, postApiV4GroupsIdWikis)


Create a wiki page

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **postApiV4GroupsIdWikis** | [**PostApiV4GroupsIdWikis**](PostApiV4GroupsIdWikis.md)|  | 

### Return type

[**ApiEntitiesWikiPage**](API_Entities_WikiPage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4GroupsIdWikisAttachments**
> ApiEntitiesWikiAttachment PostApiV4GroupsIdWikisAttachments(ctx, id, postApiV4GroupsIdWikisAttachments)
Upload an attachment to the wiki repository

This feature was introduced in GitLab 11.3.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **postApiV4GroupsIdWikisAttachments** | [**PostApiV4GroupsIdWikisAttachments**](PostApiV4GroupsIdWikisAttachments.md)|  | 

### Return type

[**ApiEntitiesWikiAttachment**](API_Entities_WikiAttachment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdWikis**
> ApiEntitiesWikiPage PostApiV4ProjectsIdWikis(ctx, id, postApiV4ProjectsIdWikis)


Create a wiki page

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **postApiV4ProjectsIdWikis** | [**PostApiV4ProjectsIdWikis**](PostApiV4ProjectsIdWikis.md)|  | 

### Return type

[**ApiEntitiesWikiPage**](API_Entities_WikiPage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdWikisAttachments**
> ApiEntitiesWikiAttachment PostApiV4ProjectsIdWikisAttachments(ctx, id, postApiV4ProjectsIdWikisAttachments)
Upload an attachment to the wiki repository

This feature was introduced in GitLab 11.3.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **postApiV4ProjectsIdWikisAttachments** | [**PostApiV4ProjectsIdWikisAttachments**](PostApiV4ProjectsIdWikisAttachments.md)|  | 

### Return type

[**ApiEntitiesWikiAttachment**](API_Entities_WikiAttachment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4GroupsIdWikisSlug**
> ApiEntitiesWikiPage PutApiV4GroupsIdWikisSlug(ctx, id, slug, putApiV4GroupsIdWikisSlug)


Update a wiki page

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **slug** | **int32**|  | 
  **putApiV4GroupsIdWikisSlug** | [**PutApiV4GroupsIdWikisSlug**](PutApiV4GroupsIdWikisSlug.md)|  | 

### Return type

[**ApiEntitiesWikiPage**](API_Entities_WikiPage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdWikisSlug**
> ApiEntitiesWikiPage PutApiV4ProjectsIdWikisSlug(ctx, id, slug, putApiV4ProjectsIdWikisSlug)


Update a wiki page

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **slug** | **int32**|  | 
  **putApiV4ProjectsIdWikisSlug** | [**PutApiV4ProjectsIdWikisSlug**](PutApiV4ProjectsIdWikisSlug.md)|  | 

### Return type

[**ApiEntitiesWikiPage**](API_Entities_WikiPage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

