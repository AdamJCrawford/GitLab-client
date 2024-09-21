# \InvitationsApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4GroupsIdInvitationsEmail**](InvitationsApi.md#DeleteApiV4GroupsIdInvitationsEmail) | **Delete** /api/v4/groups/{id}/invitations/{email} | 
[**DeleteApiV4ProjectsIdInvitationsEmail**](InvitationsApi.md#DeleteApiV4ProjectsIdInvitationsEmail) | **Delete** /api/v4/projects/{id}/invitations/{email} | 
[**GetApiV4GroupsIdInvitations**](InvitationsApi.md#GetApiV4GroupsIdInvitations) | **Get** /api/v4/groups/{id}/invitations | Get a list of group or project invitations viewable by the authenticated user
[**GetApiV4ProjectsIdInvitations**](InvitationsApi.md#GetApiV4ProjectsIdInvitations) | **Get** /api/v4/projects/{id}/invitations | Get a list of group or project invitations viewable by the authenticated user
[**PostApiV4GroupsIdInvitations**](InvitationsApi.md#PostApiV4GroupsIdInvitations) | **Post** /api/v4/groups/{id}/invitations | Invite non-members by email address to a group or project.
[**PostApiV4ProjectsIdInvitations**](InvitationsApi.md#PostApiV4ProjectsIdInvitations) | **Post** /api/v4/projects/{id}/invitations | Invite non-members by email address to a group or project.
[**PutApiV4GroupsIdInvitationsEmail**](InvitationsApi.md#PutApiV4GroupsIdInvitationsEmail) | **Put** /api/v4/groups/{id}/invitations/{email} | 
[**PutApiV4ProjectsIdInvitationsEmail**](InvitationsApi.md#PutApiV4ProjectsIdInvitationsEmail) | **Put** /api/v4/projects/{id}/invitations/{email} | 


# **DeleteApiV4GroupsIdInvitationsEmail**
> DeleteApiV4GroupsIdInvitationsEmail(ctx, id, email)


Removes an invitation from a group or project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The group ID | 
  **email** | **string**| The email address of the invitation | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdInvitationsEmail**
> DeleteApiV4ProjectsIdInvitationsEmail(ctx, id, email)


Removes an invitation from a group or project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID | 
  **email** | **string**| The email address of the invitation | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdInvitations**
> []ApiEntitiesInvitation GetApiV4GroupsIdInvitations(ctx, id, optional)
Get a list of group or project invitations viewable by the authenticated user

This feature was introduced in GitLab 13.6

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The group ID | 
 **optional** | ***InvitationsApiGetApiV4GroupsIdInvitationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InvitationsApiGetApiV4GroupsIdInvitationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **query** | **optional.String**| A query string to search for members | 

### Return type

[**[]ApiEntitiesInvitation**](API_Entities_Invitation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdInvitations**
> []ApiEntitiesInvitation GetApiV4ProjectsIdInvitations(ctx, id, optional)
Get a list of group or project invitations viewable by the authenticated user

This feature was introduced in GitLab 13.6

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID | 
 **optional** | ***InvitationsApiGetApiV4ProjectsIdInvitationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InvitationsApiGetApiV4ProjectsIdInvitationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **query** | **optional.String**| A query string to search for members | 

### Return type

[**[]ApiEntitiesInvitation**](API_Entities_Invitation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4GroupsIdInvitations**
> ApiEntitiesInvitation PostApiV4GroupsIdInvitations(ctx, id, postApiV4GroupsIdInvitations)
Invite non-members by email address to a group or project.

This feature was introduced in GitLab 13.6

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The group ID | 
  **postApiV4GroupsIdInvitations** | [**PostApiV4GroupsIdInvitations**](PostApiV4GroupsIdInvitations.md)|  | 

### Return type

[**ApiEntitiesInvitation**](API_Entities_Invitation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdInvitations**
> ApiEntitiesInvitation PostApiV4ProjectsIdInvitations(ctx, id, postApiV4ProjectsIdInvitations)
Invite non-members by email address to a group or project.

This feature was introduced in GitLab 13.6

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID | 
  **postApiV4ProjectsIdInvitations** | [**PostApiV4ProjectsIdInvitations**](PostApiV4ProjectsIdInvitations.md)|  | 

### Return type

[**ApiEntitiesInvitation**](API_Entities_Invitation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4GroupsIdInvitationsEmail**
> ApiEntitiesInvitation PutApiV4GroupsIdInvitationsEmail(ctx, id, email, putApiV4GroupsIdInvitationsEmail)


Updates a group or project invitation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The group ID | 
  **email** | **string**| The email address of the invitation | 
  **putApiV4GroupsIdInvitationsEmail** | [**PutApiV4GroupsIdInvitationsEmail**](PutApiV4GroupsIdInvitationsEmail.md)|  | 

### Return type

[**ApiEntitiesInvitation**](API_Entities_Invitation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdInvitationsEmail**
> ApiEntitiesInvitation PutApiV4ProjectsIdInvitationsEmail(ctx, id, email, putApiV4ProjectsIdInvitationsEmail)


Updates a group or project invitation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID | 
  **email** | **string**| The email address of the invitation | 
  **putApiV4ProjectsIdInvitationsEmail** | [**PutApiV4ProjectsIdInvitationsEmail**](PutApiV4ProjectsIdInvitationsEmail.md)|  | 

### Return type

[**ApiEntitiesInvitation**](API_Entities_Invitation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

