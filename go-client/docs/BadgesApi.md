# \BadgesApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4GroupsIdBadgesBadgeId**](BadgesApi.md#DeleteApiV4GroupsIdBadgesBadgeId) | **Delete** /api/v4/groups/{id}/badges/{badge_id} | Removes a badge from the group.
[**DeleteApiV4ProjectsIdBadgesBadgeId**](BadgesApi.md#DeleteApiV4ProjectsIdBadgesBadgeId) | **Delete** /api/v4/projects/{id}/badges/{badge_id} | Removes a badge from the project.
[**GetApiV4GroupsIdBadges**](BadgesApi.md#GetApiV4GroupsIdBadges) | **Get** /api/v4/groups/{id}/badges | Gets a list of group badges viewable by the authenticated user.
[**GetApiV4GroupsIdBadgesBadgeId**](BadgesApi.md#GetApiV4GroupsIdBadgesBadgeId) | **Get** /api/v4/groups/{id}/badges/{badge_id} | Gets a badge of a group.
[**GetApiV4GroupsIdBadgesRender**](BadgesApi.md#GetApiV4GroupsIdBadgesRender) | **Get** /api/v4/groups/{id}/badges/render | Preview a badge from a group.
[**GetApiV4ProjectsIdBadges**](BadgesApi.md#GetApiV4ProjectsIdBadges) | **Get** /api/v4/projects/{id}/badges | Gets a list of project badges viewable by the authenticated user.
[**GetApiV4ProjectsIdBadgesBadgeId**](BadgesApi.md#GetApiV4ProjectsIdBadgesBadgeId) | **Get** /api/v4/projects/{id}/badges/{badge_id} | Gets a badge of a project.
[**GetApiV4ProjectsIdBadgesRender**](BadgesApi.md#GetApiV4ProjectsIdBadgesRender) | **Get** /api/v4/projects/{id}/badges/render | Preview a badge from a project.
[**PostApiV4GroupsIdBadges**](BadgesApi.md#PostApiV4GroupsIdBadges) | **Post** /api/v4/groups/{id}/badges | Adds a badge to a group.
[**PostApiV4ProjectsIdBadges**](BadgesApi.md#PostApiV4ProjectsIdBadges) | **Post** /api/v4/projects/{id}/badges | Adds a badge to a project.
[**PutApiV4GroupsIdBadgesBadgeId**](BadgesApi.md#PutApiV4GroupsIdBadgesBadgeId) | **Put** /api/v4/groups/{id}/badges/{badge_id} | Updates a badge of a group.
[**PutApiV4ProjectsIdBadgesBadgeId**](BadgesApi.md#PutApiV4ProjectsIdBadgesBadgeId) | **Put** /api/v4/projects/{id}/badges/{badge_id} | Updates a badge of a project.


# **DeleteApiV4GroupsIdBadgesBadgeId**
> DeleteApiV4GroupsIdBadgesBadgeId(ctx, id, badgeId)
Removes a badge from the group.

This feature was introduced in GitLab 10.6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the group owned by the authenticated user. | 
  **badgeId** | **int32**| The badge ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdBadgesBadgeId**
> DeleteApiV4ProjectsIdBadgesBadgeId(ctx, id, badgeId)
Removes a badge from the project.

This feature was introduced in GitLab 10.6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user. | 
  **badgeId** | **int32**| The badge ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdBadges**
> []ApiEntitiesBadge GetApiV4GroupsIdBadges(ctx, id, optional)
Gets a list of group badges viewable by the authenticated user.

This feature was introduced in GitLab 10.6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the group owned by the authenticated user. | 
 **optional** | ***BadgesApiGetApiV4GroupsIdBadgesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BadgesApiGetApiV4GroupsIdBadgesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **name** | **optional.String**| Name for the badge | 

### Return type

[**[]ApiEntitiesBadge**](API_Entities_Badge.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdBadgesBadgeId**
> ApiEntitiesBadge GetApiV4GroupsIdBadgesBadgeId(ctx, id, badgeId)
Gets a badge of a group.

This feature was introduced in GitLab 10.6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the group owned by the authenticated user. | 
  **badgeId** | **int32**| The badge ID | 

### Return type

[**ApiEntitiesBadge**](API_Entities_Badge.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdBadgesRender**
> ApiEntitiesBasicBadgeDetails GetApiV4GroupsIdBadgesRender(ctx, id, linkUrl, imageUrl)
Preview a badge from a group.

This feature was introduced in GitLab 10.6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the group owned by the authenticated user. | 
  **linkUrl** | **string**| URL of the badge link | 
  **imageUrl** | **string**| URL of the badge image | 

### Return type

[**ApiEntitiesBasicBadgeDetails**](API_Entities_BasicBadgeDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdBadges**
> []ApiEntitiesBadge GetApiV4ProjectsIdBadges(ctx, id, optional)
Gets a list of project badges viewable by the authenticated user.

This feature was introduced in GitLab 10.6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user. | 
 **optional** | ***BadgesApiGetApiV4ProjectsIdBadgesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BadgesApiGetApiV4ProjectsIdBadgesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **name** | **optional.String**| Name for the badge | 

### Return type

[**[]ApiEntitiesBadge**](API_Entities_Badge.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdBadgesBadgeId**
> ApiEntitiesBadge GetApiV4ProjectsIdBadgesBadgeId(ctx, id, badgeId)
Gets a badge of a project.

This feature was introduced in GitLab 10.6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user. | 
  **badgeId** | **int32**| The badge ID | 

### Return type

[**ApiEntitiesBadge**](API_Entities_Badge.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdBadgesRender**
> ApiEntitiesBasicBadgeDetails GetApiV4ProjectsIdBadgesRender(ctx, id, linkUrl, imageUrl)
Preview a badge from a project.

This feature was introduced in GitLab 10.6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user. | 
  **linkUrl** | **string**| URL of the badge link | 
  **imageUrl** | **string**| URL of the badge image | 

### Return type

[**ApiEntitiesBasicBadgeDetails**](API_Entities_BasicBadgeDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4GroupsIdBadges**
> ApiEntitiesBadge PostApiV4GroupsIdBadges(ctx, id, postApiV4GroupsIdBadges)
Adds a badge to a group.

This feature was introduced in GitLab 10.6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the group owned by the authenticated user. | 
  **postApiV4GroupsIdBadges** | [**PostApiV4GroupsIdBadges**](PostApiV4GroupsIdBadges.md)|  | 

### Return type

[**ApiEntitiesBadge**](API_Entities_Badge.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdBadges**
> ApiEntitiesBadge PostApiV4ProjectsIdBadges(ctx, id, postApiV4ProjectsIdBadges)
Adds a badge to a project.

This feature was introduced in GitLab 10.6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user. | 
  **postApiV4ProjectsIdBadges** | [**PostApiV4ProjectsIdBadges**](PostApiV4ProjectsIdBadges.md)|  | 

### Return type

[**ApiEntitiesBadge**](API_Entities_Badge.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4GroupsIdBadgesBadgeId**
> ApiEntitiesBadge PutApiV4GroupsIdBadgesBadgeId(ctx, id, badgeId, putApiV4GroupsIdBadgesBadgeId)
Updates a badge of a group.

This feature was introduced in GitLab 10.6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the group owned by the authenticated user. | 
  **badgeId** | **int32**|  | 
  **putApiV4GroupsIdBadgesBadgeId** | [**PutApiV4GroupsIdBadgesBadgeId**](PutApiV4GroupsIdBadgesBadgeId.md)|  | 

### Return type

[**ApiEntitiesBadge**](API_Entities_Badge.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdBadgesBadgeId**
> ApiEntitiesBadge PutApiV4ProjectsIdBadgesBadgeId(ctx, id, badgeId, putApiV4ProjectsIdBadgesBadgeId)
Updates a badge of a project.

This feature was introduced in GitLab 10.6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user. | 
  **badgeId** | **int32**|  | 
  **putApiV4ProjectsIdBadgesBadgeId** | [**PutApiV4ProjectsIdBadgesBadgeId**](PutApiV4ProjectsIdBadgesBadgeId.md)|  | 

### Return type

[**ApiEntitiesBadge**](API_Entities_Badge.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

