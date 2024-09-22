# \UsersApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4UsersId**](UsersApi.md#DeleteApiV4UsersId) | **Delete** /api/v4/users/{id} | 
[**DeleteApiV4UsersIdCustomAttributesKey**](UsersApi.md#DeleteApiV4UsersIdCustomAttributesKey) | **Delete** /api/v4/users/{id}/custom_attributes/{key} | 
[**DeleteApiV4UsersIdEmailsEmailId**](UsersApi.md#DeleteApiV4UsersIdEmailsEmailId) | **Delete** /api/v4/users/{id}/emails/{email_id} | 
[**DeleteApiV4UsersIdGpgKeysKeyId**](UsersApi.md#DeleteApiV4UsersIdGpgKeysKeyId) | **Delete** /api/v4/users/{id}/gpg_keys/{key_id} | Delete an existing GPG key from a specified user. Available only for admins.
[**DeleteApiV4UsersIdIdentitiesProvider**](UsersApi.md#DeleteApiV4UsersIdIdentitiesProvider) | **Delete** /api/v4/users/{id}/identities/{provider} | 
[**DeleteApiV4UsersIdKeysKeyId**](UsersApi.md#DeleteApiV4UsersIdKeysKeyId) | **Delete** /api/v4/users/{id}/keys/{key_id} | 
[**DeleteApiV4UsersUserIdImpersonationTokensImpersonationTokenId**](UsersApi.md#DeleteApiV4UsersUserIdImpersonationTokensImpersonationTokenId) | **Delete** /api/v4/users/{user_id}/impersonation_tokens/{impersonation_token_id} | Revoke a impersonation token. Available only for admins.
[**GetApiV4Users**](UsersApi.md#GetApiV4Users) | **Get** /api/v4/users | 
[**GetApiV4UsersId**](UsersApi.md#GetApiV4UsersId) | **Get** /api/v4/users/{id} | 
[**GetApiV4UsersIdAssociationsCount**](UsersApi.md#GetApiV4UsersIdAssociationsCount) | **Get** /api/v4/users/{id}/associations_count | 
[**GetApiV4UsersIdCustomAttributes**](UsersApi.md#GetApiV4UsersIdCustomAttributes) | **Get** /api/v4/users/{id}/custom_attributes | 
[**GetApiV4UsersIdCustomAttributesKey**](UsersApi.md#GetApiV4UsersIdCustomAttributesKey) | **Get** /api/v4/users/{id}/custom_attributes/{key} | 
[**GetApiV4UsersIdEmails**](UsersApi.md#GetApiV4UsersIdEmails) | **Get** /api/v4/users/{id}/emails | 
[**GetApiV4UsersIdFollowers**](UsersApi.md#GetApiV4UsersIdFollowers) | **Get** /api/v4/users/{id}/followers | 
[**GetApiV4UsersIdFollowing**](UsersApi.md#GetApiV4UsersIdFollowing) | **Get** /api/v4/users/{id}/following | 
[**GetApiV4UsersIdGpgKeys**](UsersApi.md#GetApiV4UsersIdGpgKeys) | **Get** /api/v4/users/{id}/gpg_keys | Get the GPG keys of a specified user.
[**GetApiV4UsersIdGpgKeysKeyId**](UsersApi.md#GetApiV4UsersIdGpgKeysKeyId) | **Get** /api/v4/users/{id}/gpg_keys/{key_id} | Get a specific GPG key for a given user.
[**GetApiV4UsersIdKeysKeyId**](UsersApi.md#GetApiV4UsersIdKeysKeyId) | **Get** /api/v4/users/{id}/keys/{key_id} | 
[**GetApiV4UsersUserIdImpersonationTokens**](UsersApi.md#GetApiV4UsersUserIdImpersonationTokens) | **Get** /api/v4/users/{user_id}/impersonation_tokens | Retrieve impersonation tokens. Available only for admins.
[**GetApiV4UsersUserIdImpersonationTokensImpersonationTokenId**](UsersApi.md#GetApiV4UsersUserIdImpersonationTokensImpersonationTokenId) | **Get** /api/v4/users/{user_id}/impersonation_tokens/{impersonation_token_id} | Retrieve impersonation token. Available only for admins.
[**GetApiV4UsersUserIdKeys**](UsersApi.md#GetApiV4UsersUserIdKeys) | **Get** /api/v4/users/{user_id}/keys | 
[**GetApiV4UsersUserIdMemberships**](UsersApi.md#GetApiV4UsersUserIdMemberships) | **Get** /api/v4/users/{user_id}/memberships | 
[**GetApiV4UsersUserIdProjectDeployKeys**](UsersApi.md#GetApiV4UsersUserIdProjectDeployKeys) | **Get** /api/v4/users/{user_id}/project_deploy_keys | 
[**GetApiV4UsersUserIdStatus**](UsersApi.md#GetApiV4UsersUserIdStatus) | **Get** /api/v4/users/{user_id}/status | 
[**PatchApiV4UsersIdDisableTwoFactor**](UsersApi.md#PatchApiV4UsersIdDisableTwoFactor) | **Patch** /api/v4/users/{id}/disable_two_factor | Disable two factor authentication for a user. Available only for admins
[**PostApiV4Users**](UsersApi.md#PostApiV4Users) | **Post** /api/v4/users | 
[**PostApiV4UsersIdActivate**](UsersApi.md#PostApiV4UsersIdActivate) | **Post** /api/v4/users/{id}/activate | 
[**PostApiV4UsersIdApprove**](UsersApi.md#PostApiV4UsersIdApprove) | **Post** /api/v4/users/{id}/approve | 
[**PostApiV4UsersIdBan**](UsersApi.md#PostApiV4UsersIdBan) | **Post** /api/v4/users/{id}/ban | 
[**PostApiV4UsersIdBlock**](UsersApi.md#PostApiV4UsersIdBlock) | **Post** /api/v4/users/{id}/block | 
[**PostApiV4UsersIdDeactivate**](UsersApi.md#PostApiV4UsersIdDeactivate) | **Post** /api/v4/users/{id}/deactivate | 
[**PostApiV4UsersIdEmails**](UsersApi.md#PostApiV4UsersIdEmails) | **Post** /api/v4/users/{id}/emails | 
[**PostApiV4UsersIdFollow**](UsersApi.md#PostApiV4UsersIdFollow) | **Post** /api/v4/users/{id}/follow | 
[**PostApiV4UsersIdGpgKeys**](UsersApi.md#PostApiV4UsersIdGpgKeys) | **Post** /api/v4/users/{id}/gpg_keys | Add a GPG key to a specified user. Available only for admins.
[**PostApiV4UsersIdGpgKeysKeyIdRevoke**](UsersApi.md#PostApiV4UsersIdGpgKeysKeyIdRevoke) | **Post** /api/v4/users/{id}/gpg_keys/{key_id}/revoke | Revokes an existing GPG key from a specified user. Available only for admins.
[**PostApiV4UsersIdReject**](UsersApi.md#PostApiV4UsersIdReject) | **Post** /api/v4/users/{id}/reject | 
[**PostApiV4UsersIdUnban**](UsersApi.md#PostApiV4UsersIdUnban) | **Post** /api/v4/users/{id}/unban | 
[**PostApiV4UsersIdUnblock**](UsersApi.md#PostApiV4UsersIdUnblock) | **Post** /api/v4/users/{id}/unblock | 
[**PostApiV4UsersIdUnfollow**](UsersApi.md#PostApiV4UsersIdUnfollow) | **Post** /api/v4/users/{id}/unfollow | 
[**PostApiV4UsersUserIdImpersonationTokens**](UsersApi.md#PostApiV4UsersUserIdImpersonationTokens) | **Post** /api/v4/users/{user_id}/impersonation_tokens | Create a impersonation token. Available only for admins.
[**PostApiV4UsersUserIdKeys**](UsersApi.md#PostApiV4UsersUserIdKeys) | **Post** /api/v4/users/{user_id}/keys | 
[**PostApiV4UsersUserIdPersonalAccessTokens**](UsersApi.md#PostApiV4UsersUserIdPersonalAccessTokens) | **Post** /api/v4/users/{user_id}/personal_access_tokens | Create a personal access token. Available only for admins.
[**PutApiV4UsersId**](UsersApi.md#PutApiV4UsersId) | **Put** /api/v4/users/{id} | 
[**PutApiV4UsersIdCustomAttributesKey**](UsersApi.md#PutApiV4UsersIdCustomAttributesKey) | **Put** /api/v4/users/{id}/custom_attributes/{key} | 


# **DeleteApiV4UsersId**
> ApiEntitiesEmail DeleteApiV4UsersId(ctx, id, optional)


Delete a user. Available only for admins.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the user | 
 **optional** | ***UsersApiDeleteApiV4UsersIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiDeleteApiV4UsersIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hardDelete** | **optional.Bool**| Whether to remove a user&#39;s contributions | 

### Return type

[**ApiEntitiesEmail**](API_Entities_Email.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4UsersIdCustomAttributesKey**
> DeleteApiV4UsersIdCustomAttributesKey(ctx, key, id)


Delete a custom attribute on a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| The key of the custom attribute | 
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4UsersIdEmailsEmailId**
> ApiEntitiesEmail DeleteApiV4UsersIdEmailsEmailId(ctx, id, emailId)


Delete an email address of a specified user. Available only for admins.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the user | 
  **emailId** | **int32**| The ID of the email | 

### Return type

[**ApiEntitiesEmail**](API_Entities_Email.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4UsersIdGpgKeysKeyId**
> DeleteApiV4UsersIdGpgKeysKeyId(ctx, id, keyId)
Delete an existing GPG key from a specified user. Available only for admins.

This feature was added in GitLab 10.0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the user | 
  **keyId** | **int32**| The ID of the GPG key | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4UsersIdIdentitiesProvider**
> ApiEntitiesUserWithAdmin DeleteApiV4UsersIdIdentitiesProvider(ctx, id, provider)


Delete a user's identity. Available only for admins

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the user | 
  **provider** | **string**| The external provider | 

### Return type

[**ApiEntitiesUserWithAdmin**](API_Entities_UserWithAdmin.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4UsersIdKeysKeyId**
> ApiEntitiesSshKey DeleteApiV4UsersIdKeysKeyId(ctx, id, keyId)


Delete an existing SSH key from a specified user. Available only for admins.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the user | 
  **keyId** | **int32**| The ID of the SSH key | 

### Return type

[**ApiEntitiesSshKey**](API_Entities_SSHKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4UsersUserIdImpersonationTokensImpersonationTokenId**
> DeleteApiV4UsersUserIdImpersonationTokensImpersonationTokenId(ctx, userId, impersonationTokenId)
Revoke a impersonation token. Available only for admins.

This feature was introduced in GitLab 9.0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int32**| The ID of the user | 
  **impersonationTokenId** | **int32**| The ID of the impersonation token | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4Users**
> ApiEntitiesUserBasic GetApiV4Users(ctx, optional)


Get the list of users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiGetApiV4UsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiGetApiV4UsersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **optional.String**| Get a single user with a specific username | 
 **externUid** | **optional.String**| Get a single user with a specific external authentication provider UID | 
 **provider** | **optional.String**| The external provider | 
 **search** | **optional.String**| Search for a username | 
 **active** | **optional.Bool**| Filters only active users | [default to false]
 **external** | **optional.Bool**| Filters only external users | [default to false]
 **excludeExternal** | **optional.Bool**| Filters only non external users | [default to false]
 **blocked** | **optional.Bool**| Filters only blocked users | [default to false]
 **createdAfter** | **optional.Time**| Return users created after the specified time | 
 **createdBefore** | **optional.Time**| Return users created before the specified time | 
 **withoutProjects** | **optional.Bool**| Filters only users without projects | [default to false]
 **excludeInternal** | **optional.Bool**| Filters only non internal users | [default to false]
 **withoutProjectBots** | **optional.Bool**| Filters users without project bots | [default to false]
 **admins** | **optional.Bool**| Filters only admin users | [default to false]
 **twoFactor** | **optional.String**| Filter users by Two-factor authentication. | 
 **orderBy** | **optional.String**| Return users ordered by a field | 
 **sort** | **optional.String**| Return users sorted in ascending and descending order | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **withCustomAttributes** | **optional.Bool**| Include custom attributes in the response | [default to false]
 **skipLdap** | **optional.Bool**| Skip LDAP users | [default to false]
 **samlProviderId** | **optional.Int32**| Return only users from the specified SAML provider Id | 
 **auditors** | **optional.Bool**| Filters only auditor users | [default to false]

### Return type

[**ApiEntitiesUserBasic**](API_Entities_UserBasic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4UsersId**
> ApiEntitiesUser GetApiV4UsersId(ctx, id, optional)


Get a single user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the user | 
 **optional** | ***UsersApiGetApiV4UsersIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiGetApiV4UsersIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withCustomAttributes** | **optional.Bool**| Include custom attributes in the response | [default to false]

### Return type

[**ApiEntitiesUser**](API_Entities_User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4UsersIdAssociationsCount**
> GetApiV4UsersIdAssociationsCount(ctx, id)


Returns a list of a specified user's count of projects, groups, issues and merge requests.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| ID of the user to query. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4UsersIdCustomAttributes**
> ApiEntitiesCustomAttribute GetApiV4UsersIdCustomAttributes(ctx, id)


Get all custom attributes on a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**ApiEntitiesCustomAttribute**](API_Entities_CustomAttribute.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4UsersIdCustomAttributesKey**
> ApiEntitiesCustomAttribute GetApiV4UsersIdCustomAttributesKey(ctx, key, id)


Get a custom attribute on a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| The key of the custom attribute | 
  **id** | **int32**|  | 

### Return type

[**ApiEntitiesCustomAttribute**](API_Entities_CustomAttribute.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4UsersIdEmails**
> ApiEntitiesEmail GetApiV4UsersIdEmails(ctx, id, optional)


Get the emails addresses of a specified user. Available only for admins.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the user | 
 **optional** | ***UsersApiGetApiV4UsersIdEmailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiGetApiV4UsersIdEmailsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**ApiEntitiesEmail**](API_Entities_Email.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4UsersIdFollowers**
> ApiEntitiesUserBasic GetApiV4UsersIdFollowers(ctx, id, optional)


Get the followers of a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the user | 
 **optional** | ***UsersApiGetApiV4UsersIdFollowersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiGetApiV4UsersIdFollowersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**ApiEntitiesUserBasic**](API_Entities_UserBasic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4UsersIdFollowing**
> ApiEntitiesUserBasic GetApiV4UsersIdFollowing(ctx, id, optional)


Get the users who follow a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the user | 
 **optional** | ***UsersApiGetApiV4UsersIdFollowingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiGetApiV4UsersIdFollowingOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**ApiEntitiesUserBasic**](API_Entities_UserBasic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4UsersIdGpgKeys**
> ApiEntitiesGpgKey GetApiV4UsersIdGpgKeys(ctx, id, optional)
Get the GPG keys of a specified user.

This feature was added in GitLab 10.0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the user | 
 **optional** | ***UsersApiGetApiV4UsersIdGpgKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiGetApiV4UsersIdGpgKeysOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**ApiEntitiesGpgKey**](API_Entities_GpgKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4UsersIdGpgKeysKeyId**
> ApiEntitiesGpgKey GetApiV4UsersIdGpgKeysKeyId(ctx, id, keyId)
Get a specific GPG key for a given user.

This feature was added in GitLab 13.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the user | 
  **keyId** | **int32**| The ID of the GPG key | 

### Return type

[**ApiEntitiesGpgKey**](API_Entities_GpgKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4UsersIdKeysKeyId**
> ApiEntitiesSshKey GetApiV4UsersIdKeysKeyId(ctx, id, keyId)


Get a SSH key of a specified user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the user | 
  **keyId** | **int32**| The ID of the SSH key | 

### Return type

[**ApiEntitiesSshKey**](API_Entities_SSHKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4UsersUserIdImpersonationTokens**
> ApiEntitiesImpersonationToken GetApiV4UsersUserIdImpersonationTokens(ctx, userId, optional)
Retrieve impersonation tokens. Available only for admins.

This feature was introduced in GitLab 9.0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int32**| The ID of the user | 
 **optional** | ***UsersApiGetApiV4UsersUserIdImpersonationTokensOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiGetApiV4UsersUserIdImpersonationTokensOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **state** | **optional.String**| Filters (all|active|inactive) impersonation_tokens | [default to all]

### Return type

[**ApiEntitiesImpersonationToken**](API_Entities_ImpersonationToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4UsersUserIdImpersonationTokensImpersonationTokenId**
> ApiEntitiesImpersonationToken GetApiV4UsersUserIdImpersonationTokensImpersonationTokenId(ctx, userId, impersonationTokenId)
Retrieve impersonation token. Available only for admins.

This feature was introduced in GitLab 9.0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int32**| The ID of the user | 
  **impersonationTokenId** | **int32**| The ID of the impersonation token | 

### Return type

[**ApiEntitiesImpersonationToken**](API_Entities_ImpersonationToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4UsersUserIdKeys**
> ApiEntitiesSshKey GetApiV4UsersUserIdKeys(ctx, userId, optional)


Get the SSH keys of a specified user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| The ID or username of the user | 
 **optional** | ***UsersApiGetApiV4UsersUserIdKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiGetApiV4UsersUserIdKeysOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**ApiEntitiesSshKey**](API_Entities_SSHKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4UsersUserIdMemberships**
> ApiEntitiesMembership GetApiV4UsersUserIdMemberships(ctx, userId, optional)


Get memberships

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int32**| The ID of the user | 
 **optional** | ***UsersApiGetApiV4UsersUserIdMembershipsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiGetApiV4UsersUserIdMembershipsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**|  | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**ApiEntitiesMembership**](API_Entities_Membership.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4UsersUserIdProjectDeployKeys**
> ApiEntitiesDeployKey GetApiV4UsersUserIdProjectDeployKeys(ctx, userId, optional)


Get the project-level Deploy keys that a specified user can access to.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| The ID or username of the user | 
 **optional** | ***UsersApiGetApiV4UsersUserIdProjectDeployKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiGetApiV4UsersUserIdProjectDeployKeysOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**ApiEntitiesDeployKey**](API_Entities_DeployKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4UsersUserIdStatus**
> GetApiV4UsersUserIdStatus(ctx, userId)


Get the status of a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| The ID or username of the user | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchApiV4UsersIdDisableTwoFactor**
> ApiEntitiesUserWithAdmin PatchApiV4UsersIdDisableTwoFactor(ctx, id)
Disable two factor authentication for a user. Available only for admins

This feature was added in GitLab 15.2

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the user | 

### Return type

[**ApiEntitiesUserWithAdmin**](API_Entities_UserWithAdmin.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4Users**
> ApiEntitiesUserWithAdmin PostApiV4Users(ctx, postApiV4Users)


Create a user. Available only for admins.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4Users** | [**PostApiV4Users**](PostApiV4Users.md)|  | 

### Return type

[**ApiEntitiesUserWithAdmin**](API_Entities_UserWithAdmin.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4UsersIdActivate**
> PostApiV4UsersIdActivate(ctx, id)


Activate a deactivated user. Available only for admins.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the user | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4UsersIdApprove**
> PostApiV4UsersIdApprove(ctx, id)


Approve a pending user. Available only for admins.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the user | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4UsersIdBan**
> PostApiV4UsersIdBan(ctx, id)


Ban a user. Available only for admins.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the user | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4UsersIdBlock**
> PostApiV4UsersIdBlock(ctx, id)


Block a user. Available only for admins.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the user | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4UsersIdDeactivate**
> PostApiV4UsersIdDeactivate(ctx, id)


Deactivate an active user. Available only for admins.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the user | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4UsersIdEmails**
> ApiEntitiesEmail PostApiV4UsersIdEmails(ctx, id, postApiV4UsersIdEmails)


Add an email address to a specified user. Available only for admins.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the user | 
  **postApiV4UsersIdEmails** | [**PostApiV4UsersIdEmails**](PostApiV4UsersIdEmails.md)|  | 

### Return type

[**ApiEntitiesEmail**](API_Entities_Email.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4UsersIdFollow**
> ApiEntitiesUser PostApiV4UsersIdFollow(ctx, id)


Follow a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the user | 

### Return type

[**ApiEntitiesUser**](API_Entities_User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4UsersIdGpgKeys**
> ApiEntitiesGpgKey PostApiV4UsersIdGpgKeys(ctx, id, postApiV4UsersIdGpgKeys)
Add a GPG key to a specified user. Available only for admins.

This feature was added in GitLab 10.0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the user | 
  **postApiV4UsersIdGpgKeys** | [**PostApiV4UsersIdGpgKeys**](PostApiV4UsersIdGpgKeys.md)|  | 

### Return type

[**ApiEntitiesGpgKey**](API_Entities_GpgKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4UsersIdGpgKeysKeyIdRevoke**
> PostApiV4UsersIdGpgKeysKeyIdRevoke(ctx, id, keyId)
Revokes an existing GPG key from a specified user. Available only for admins.

This feature was added in GitLab 10.0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the user | 
  **keyId** | **int32**| The ID of the GPG key | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4UsersIdReject**
> PostApiV4UsersIdReject(ctx, id)


Reject a pending user. Available only for admins.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the user | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4UsersIdUnban**
> PostApiV4UsersIdUnban(ctx, id)


Unban a user. Available only for admins.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the user | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4UsersIdUnblock**
> PostApiV4UsersIdUnblock(ctx, id)


Unblock a user. Available only for admins.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the user | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4UsersIdUnfollow**
> ApiEntitiesUser PostApiV4UsersIdUnfollow(ctx, id)


Unfollow a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the user | 

### Return type

[**ApiEntitiesUser**](API_Entities_User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4UsersUserIdImpersonationTokens**
> ApiEntitiesImpersonationTokenWithToken PostApiV4UsersUserIdImpersonationTokens(ctx, userId, postApiV4UsersUserIdImpersonationTokens)
Create a impersonation token. Available only for admins.

This feature was introduced in GitLab 9.0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int32**| The ID of the user | 
  **postApiV4UsersUserIdImpersonationTokens** | [**PostApiV4UsersUserIdImpersonationTokens**](PostApiV4UsersUserIdImpersonationTokens.md)|  | 

### Return type

[**ApiEntitiesImpersonationTokenWithToken**](API_Entities_ImpersonationTokenWithToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4UsersUserIdKeys**
> ApiEntitiesSshKey PostApiV4UsersUserIdKeys(ctx, userId, postApiV4UsersUserIdKeys)


Add an SSH key to a specified user. Available only for admins.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int32**| The ID of the user | 
  **postApiV4UsersUserIdKeys** | [**PostApiV4UsersUserIdKeys**](PostApiV4UsersUserIdKeys.md)|  | 

### Return type

[**ApiEntitiesSshKey**](API_Entities_SSHKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4UsersUserIdPersonalAccessTokens**
> ApiEntitiesPersonalAccessTokenWithToken PostApiV4UsersUserIdPersonalAccessTokens(ctx, userId, postApiV4UsersUserIdPersonalAccessTokens)
Create a personal access token. Available only for admins.

This feature was introduced in GitLab 13.6

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int32**| The ID of the user | 
  **postApiV4UsersUserIdPersonalAccessTokens** | [**PostApiV4UsersUserIdPersonalAccessTokens**](PostApiV4UsersUserIdPersonalAccessTokens.md)|  | 

### Return type

[**ApiEntitiesPersonalAccessTokenWithToken**](API_Entities_PersonalAccessTokenWithToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4UsersId**
> ApiEntitiesUserWithAdmin PutApiV4UsersId(ctx, id, putApiV4UsersId)


Update a user. Available only for admins.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the user | 
  **putApiV4UsersId** | [**PutApiV4UsersId**](PutApiV4UsersId.md)|  | 

### Return type

[**ApiEntitiesUserWithAdmin**](API_Entities_UserWithAdmin.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4UsersIdCustomAttributesKey**
> PutApiV4UsersIdCustomAttributesKey(ctx, key, id, putApiV4UsersIdCustomAttributesKey)


Set a custom attribute on a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| The key of the custom attribute | 
  **id** | **int32**|  | 
  **putApiV4UsersIdCustomAttributesKey** | [**PutApiV4UsersIdCustomAttributesKey**](PutApiV4UsersIdCustomAttributesKey.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

