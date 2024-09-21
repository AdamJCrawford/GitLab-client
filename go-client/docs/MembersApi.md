# \MembersApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4GroupsIdMembersUserId**](MembersApi.md#DeleteApiV4GroupsIdMembersUserId) | **Delete** /api/v4/groups/{id}/members/{user_id} | 
[**DeleteApiV4ProjectsIdMembersUserId**](MembersApi.md#DeleteApiV4ProjectsIdMembersUserId) | **Delete** /api/v4/projects/{id}/members/{user_id} | 
[**GetApiV4GroupsIdMembers**](MembersApi.md#GetApiV4GroupsIdMembers) | **Get** /api/v4/groups/{id}/members | 
[**GetApiV4GroupsIdMembersAll**](MembersApi.md#GetApiV4GroupsIdMembersAll) | **Get** /api/v4/groups/{id}/members/all | 
[**GetApiV4GroupsIdMembersAllUserId**](MembersApi.md#GetApiV4GroupsIdMembersAllUserId) | **Get** /api/v4/groups/{id}/members/all/{user_id} | 
[**GetApiV4GroupsIdMembersUserId**](MembersApi.md#GetApiV4GroupsIdMembersUserId) | **Get** /api/v4/groups/{id}/members/{user_id} | 
[**GetApiV4ProjectsIdMembers**](MembersApi.md#GetApiV4ProjectsIdMembers) | **Get** /api/v4/projects/{id}/members | 
[**GetApiV4ProjectsIdMembersAll**](MembersApi.md#GetApiV4ProjectsIdMembersAll) | **Get** /api/v4/projects/{id}/members/all | 
[**GetApiV4ProjectsIdMembersAllUserId**](MembersApi.md#GetApiV4ProjectsIdMembersAllUserId) | **Get** /api/v4/projects/{id}/members/all/{user_id} | 
[**GetApiV4ProjectsIdMembersUserId**](MembersApi.md#GetApiV4ProjectsIdMembersUserId) | **Get** /api/v4/projects/{id}/members/{user_id} | 
[**PostApiV4GroupsIdMembers**](MembersApi.md#PostApiV4GroupsIdMembers) | **Post** /api/v4/groups/{id}/members | 
[**PostApiV4ProjectsIdMembers**](MembersApi.md#PostApiV4ProjectsIdMembers) | **Post** /api/v4/projects/{id}/members | 
[**PutApiV4GroupsIdMembersUserId**](MembersApi.md#PutApiV4GroupsIdMembersUserId) | **Put** /api/v4/groups/{id}/members/{user_id} | 
[**PutApiV4ProjectsIdMembersUserId**](MembersApi.md#PutApiV4ProjectsIdMembersUserId) | **Put** /api/v4/projects/{id}/members/{user_id} | 


# **DeleteApiV4GroupsIdMembersUserId**
> DeleteApiV4GroupsIdMembersUserId(ctx, id, userId, optional)


Removes a user from a group or project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The group ID | 
  **userId** | **int32**| The user ID of the member | 
 **optional** | ***MembersApiDeleteApiV4GroupsIdMembersUserIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MembersApiDeleteApiV4GroupsIdMembersUserIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **skipSubresources** | **optional.Bool**| Flag indicating if the deletion of direct memberships of the removed member in subgroups and projects should be skipped | [default to false]
 **unassignIssuables** | **optional.Bool**| Flag indicating if the removed member should be unassigned from any issues or merge requests within given group or project | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdMembersUserId**
> DeleteApiV4ProjectsIdMembersUserId(ctx, id, userId, optional)


Removes a user from a group or project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID | 
  **userId** | **int32**| The user ID of the member | 
 **optional** | ***MembersApiDeleteApiV4ProjectsIdMembersUserIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MembersApiDeleteApiV4ProjectsIdMembersUserIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **skipSubresources** | **optional.Bool**| Flag indicating if the deletion of direct memberships of the removed member in subgroups and projects should be skipped | [default to false]
 **unassignIssuables** | **optional.Bool**| Flag indicating if the removed member should be unassigned from any issues or merge requests within given group or project | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdMembers**
> []ApiEntitiesMember GetApiV4GroupsIdMembers(ctx, id, optional)


Gets a list of group or project members viewable by the authenticated user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The group ID | 
 **optional** | ***MembersApiGetApiV4GroupsIdMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MembersApiGetApiV4GroupsIdMembersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **optional.String**| A query string to search for members | 
 **userIds** | [**optional.Interface of []int32**](int32.md)| Array of user ids to look up for membership | 
 **skipUsers** | [**optional.Interface of []int32**](int32.md)| Array of user ids to be skipped for membership | 
 **showSeatInfo** | **optional.Bool**| Show seat information for members | 
 **withSamlIdentity** | **optional.Bool**| List only members with linked SAML identity | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesMember**](API_Entities_Member.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdMembersAll**
> []ApiEntitiesMember GetApiV4GroupsIdMembersAll(ctx, id, optional)


Gets a list of group or project members viewable by the authenticated user, including those who gained membership through ancestor group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The group ID | 
 **optional** | ***MembersApiGetApiV4GroupsIdMembersAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MembersApiGetApiV4GroupsIdMembersAllOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **optional.String**| A query string to search for members | 
 **userIds** | [**optional.Interface of []int32**](int32.md)| Array of user ids to look up for membership | 
 **showSeatInfo** | **optional.Bool**| Show seat information for members | 
 **state** | **optional.String**| Filter results by member state | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesMember**](API_Entities_Member.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdMembersAllUserId**
> ApiEntitiesMember GetApiV4GroupsIdMembersAllUserId(ctx, id, userId)


Gets a member of a group or project, including those who gained membership through ancestor group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The group ID | 
  **userId** | **int32**| The user ID of the member | 

### Return type

[**ApiEntitiesMember**](API_Entities_Member.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdMembersUserId**
> ApiEntitiesMember GetApiV4GroupsIdMembersUserId(ctx, id, userId)


Gets a member of a group or project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The group ID | 
  **userId** | **int32**| The user ID of the member | 

### Return type

[**ApiEntitiesMember**](API_Entities_Member.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdMembers**
> []ApiEntitiesMember GetApiV4ProjectsIdMembers(ctx, id, optional)


Gets a list of group or project members viewable by the authenticated user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID | 
 **optional** | ***MembersApiGetApiV4ProjectsIdMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MembersApiGetApiV4ProjectsIdMembersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **optional.String**| A query string to search for members | 
 **userIds** | [**optional.Interface of []int32**](int32.md)| Array of user ids to look up for membership | 
 **skipUsers** | [**optional.Interface of []int32**](int32.md)| Array of user ids to be skipped for membership | 
 **showSeatInfo** | **optional.Bool**| Show seat information for members | 
 **withSamlIdentity** | **optional.Bool**| List only members with linked SAML identity | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesMember**](API_Entities_Member.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdMembersAll**
> []ApiEntitiesMember GetApiV4ProjectsIdMembersAll(ctx, id, optional)


Gets a list of group or project members viewable by the authenticated user, including those who gained membership through ancestor group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID | 
 **optional** | ***MembersApiGetApiV4ProjectsIdMembersAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MembersApiGetApiV4ProjectsIdMembersAllOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **optional.String**| A query string to search for members | 
 **userIds** | [**optional.Interface of []int32**](int32.md)| Array of user ids to look up for membership | 
 **showSeatInfo** | **optional.Bool**| Show seat information for members | 
 **state** | **optional.String**| Filter results by member state | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesMember**](API_Entities_Member.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdMembersAllUserId**
> ApiEntitiesMember GetApiV4ProjectsIdMembersAllUserId(ctx, id, userId)


Gets a member of a group or project, including those who gained membership through ancestor group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID | 
  **userId** | **int32**| The user ID of the member | 

### Return type

[**ApiEntitiesMember**](API_Entities_Member.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdMembersUserId**
> ApiEntitiesMember GetApiV4ProjectsIdMembersUserId(ctx, id, userId)


Gets a member of a group or project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID | 
  **userId** | **int32**| The user ID of the member | 

### Return type

[**ApiEntitiesMember**](API_Entities_Member.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4GroupsIdMembers**
> ApiEntitiesMember PostApiV4GroupsIdMembers(ctx, id, postApiV4GroupsIdMembers)


Adds a member to a group or project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The group ID | 
  **postApiV4GroupsIdMembers** | [**PostApiV4GroupsIdMembers**](PostApiV4GroupsIdMembers.md)|  | 

### Return type

[**ApiEntitiesMember**](API_Entities_Member.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdMembers**
> ApiEntitiesMember PostApiV4ProjectsIdMembers(ctx, id, postApiV4ProjectsIdMembers)


Adds a member to a group or project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID | 
  **postApiV4ProjectsIdMembers** | [**PostApiV4ProjectsIdMembers**](PostApiV4ProjectsIdMembers.md)|  | 

### Return type

[**ApiEntitiesMember**](API_Entities_Member.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4GroupsIdMembersUserId**
> ApiEntitiesMember PutApiV4GroupsIdMembersUserId(ctx, id, userId, putApiV4GroupsIdMembersUserId)


Updates a member of a group or project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The group ID | 
  **userId** | **int32**| The user ID of the new member | 
  **putApiV4GroupsIdMembersUserId** | [**PutApiV4GroupsIdMembersUserId**](PutApiV4GroupsIdMembersUserId.md)|  | 

### Return type

[**ApiEntitiesMember**](API_Entities_Member.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdMembersUserId**
> ApiEntitiesMember PutApiV4ProjectsIdMembersUserId(ctx, id, userId, putApiV4ProjectsIdMembersUserId)


Updates a member of a group or project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID | 
  **userId** | **int32**| The user ID of the new member | 
  **putApiV4ProjectsIdMembersUserId** | [**PutApiV4ProjectsIdMembersUserId**](PutApiV4ProjectsIdMembersUserId.md)|  | 

### Return type

[**ApiEntitiesMember**](API_Entities_Member.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

