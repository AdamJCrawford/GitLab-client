# \AccessRequestsApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4GroupsIdAccessRequestsUserId**](AccessRequestsApi.md#DeleteApiV4GroupsIdAccessRequestsUserId) | **Delete** /api/v4/groups/{id}/access_requests/{user_id} | Denies an access request for the given user.
[**DeleteApiV4ProjectsIdAccessRequestsUserId**](AccessRequestsApi.md#DeleteApiV4ProjectsIdAccessRequestsUserId) | **Delete** /api/v4/projects/{id}/access_requests/{user_id} | Denies an access request for the given user.
[**GetApiV4GroupsIdAccessRequests**](AccessRequestsApi.md#GetApiV4GroupsIdAccessRequests) | **Get** /api/v4/groups/{id}/access_requests | Gets a list of access requests for a group.
[**GetApiV4ProjectsIdAccessRequests**](AccessRequestsApi.md#GetApiV4ProjectsIdAccessRequests) | **Get** /api/v4/projects/{id}/access_requests | Gets a list of access requests for a project.
[**PostApiV4GroupsIdAccessRequests**](AccessRequestsApi.md#PostApiV4GroupsIdAccessRequests) | **Post** /api/v4/groups/{id}/access_requests | Requests access for the authenticated user to a group.
[**PostApiV4ProjectsIdAccessRequests**](AccessRequestsApi.md#PostApiV4ProjectsIdAccessRequests) | **Post** /api/v4/projects/{id}/access_requests | Requests access for the authenticated user to a project.
[**PutApiV4GroupsIdAccessRequestsUserIdApprove**](AccessRequestsApi.md#PutApiV4GroupsIdAccessRequestsUserIdApprove) | **Put** /api/v4/groups/{id}/access_requests/{user_id}/approve | Approves an access request for the given user.
[**PutApiV4ProjectsIdAccessRequestsUserIdApprove**](AccessRequestsApi.md#PutApiV4ProjectsIdAccessRequestsUserIdApprove) | **Put** /api/v4/projects/{id}/access_requests/{user_id}/approve | Approves an access request for the given user.


# **DeleteApiV4GroupsIdAccessRequestsUserId**
> DeleteApiV4GroupsIdAccessRequestsUserId(ctx, id, userId)
Denies an access request for the given user.

This feature was introduced in GitLab 8.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the group owned by the authenticated user | 
  **userId** | **int32**| The user ID of the access requester | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdAccessRequestsUserId**
> DeleteApiV4ProjectsIdAccessRequestsUserId(ctx, id, userId)
Denies an access request for the given user.

This feature was introduced in GitLab 8.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **userId** | **int32**| The user ID of the access requester | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdAccessRequests**
> ApiEntitiesAccessRequester GetApiV4GroupsIdAccessRequests(ctx, id, optional)
Gets a list of access requests for a group.

This feature was introduced in GitLab 8.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the group owned by the authenticated user | 
 **optional** | ***AccessRequestsApiGetApiV4GroupsIdAccessRequestsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccessRequestsApiGetApiV4GroupsIdAccessRequestsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**ApiEntitiesAccessRequester**](API_Entities_AccessRequester.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdAccessRequests**
> ApiEntitiesAccessRequester GetApiV4ProjectsIdAccessRequests(ctx, id, optional)
Gets a list of access requests for a project.

This feature was introduced in GitLab 8.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
 **optional** | ***AccessRequestsApiGetApiV4ProjectsIdAccessRequestsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccessRequestsApiGetApiV4ProjectsIdAccessRequestsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**ApiEntitiesAccessRequester**](API_Entities_AccessRequester.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4GroupsIdAccessRequests**
> ApiEntitiesAccessRequester PostApiV4GroupsIdAccessRequests(ctx, id)
Requests access for the authenticated user to a group.

This feature was introduced in GitLab 8.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the group owned by the authenticated user | 

### Return type

[**ApiEntitiesAccessRequester**](API_Entities_AccessRequester.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdAccessRequests**
> ApiEntitiesAccessRequester PostApiV4ProjectsIdAccessRequests(ctx, id)
Requests access for the authenticated user to a project.

This feature was introduced in GitLab 8.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 

### Return type

[**ApiEntitiesAccessRequester**](API_Entities_AccessRequester.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4GroupsIdAccessRequestsUserIdApprove**
> ApiEntitiesAccessRequester PutApiV4GroupsIdAccessRequestsUserIdApprove(ctx, id, userId, putApiV4GroupsIdAccessRequestsUserIdApprove)
Approves an access request for the given user.

This feature was introduced in GitLab 8.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the group owned by the authenticated user | 
  **userId** | **int32**| The user ID of the access requester | 
  **putApiV4GroupsIdAccessRequestsUserIdApprove** | [**PutApiV4GroupsIdAccessRequestsUserIdApprove**](PutApiV4GroupsIdAccessRequestsUserIdApprove.md)|  | 

### Return type

[**ApiEntitiesAccessRequester**](API_Entities_AccessRequester.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdAccessRequestsUserIdApprove**
> ApiEntitiesAccessRequester PutApiV4ProjectsIdAccessRequestsUserIdApprove(ctx, id, userId, putApiV4ProjectsIdAccessRequestsUserIdApprove)
Approves an access request for the given user.

This feature was introduced in GitLab 8.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **userId** | **int32**| The user ID of the access requester | 
  **putApiV4ProjectsIdAccessRequestsUserIdApprove** | [**PutApiV4ProjectsIdAccessRequestsUserIdApprove**](PutApiV4ProjectsIdAccessRequestsUserIdApprove.md)|  | 

### Return type

[**ApiEntitiesAccessRequester**](API_Entities_AccessRequester.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

