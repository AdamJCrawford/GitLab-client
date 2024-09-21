# \UserApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4UserEmailsEmailId**](UserApi.md#DeleteApiV4UserEmailsEmailId) | **Delete** /api/v4/user/emails/{email_id} | 
[**DeleteApiV4UserGpgKeysKeyId**](UserApi.md#DeleteApiV4UserGpgKeysKeyId) | **Delete** /api/v4/user/gpg_keys/{key_id} | Delete a GPG key from the currently authenticated user
[**DeleteApiV4UserKeysKeyId**](UserApi.md#DeleteApiV4UserKeysKeyId) | **Delete** /api/v4/user/keys/{key_id} | 
[**GetApiV3User**](UserApi.md#GetApiV3User) | **Get** /api/v3/user | 
[**GetApiV4UserActivities**](UserApi.md#GetApiV4UserActivities) | **Get** /api/v4/user/activities | 
[**GetApiV4UserEmails**](UserApi.md#GetApiV4UserEmails) | **Get** /api/v4/user/emails | 
[**GetApiV4UserEmailsEmailId**](UserApi.md#GetApiV4UserEmailsEmailId) | **Get** /api/v4/user/emails/{email_id} | 
[**GetApiV4UserGpgKeys**](UserApi.md#GetApiV4UserGpgKeys) | **Get** /api/v4/user/gpg_keys | Get the currently authenticated user&#39;s GPG keys
[**GetApiV4UserGpgKeysKeyId**](UserApi.md#GetApiV4UserGpgKeysKeyId) | **Get** /api/v4/user/gpg_keys/{key_id} | Get a single GPG key owned by currently authenticated user
[**GetApiV4UserKeys**](UserApi.md#GetApiV4UserKeys) | **Get** /api/v4/user/keys | 
[**GetApiV4UserKeysKeyId**](UserApi.md#GetApiV4UserKeysKeyId) | **Get** /api/v4/user/keys/{key_id} | 
[**GetApiV4UserPreferences**](UserApi.md#GetApiV4UserPreferences) | **Get** /api/v4/user/preferences | Get the current user&#39;s preferences
[**GetApiV4UserStatus**](UserApi.md#GetApiV4UserStatus) | **Get** /api/v4/user/status | 
[**PatchApiV4UserStatus**](UserApi.md#PatchApiV4UserStatus) | **Patch** /api/v4/user/status | Set the status of the current user
[**PostApiV4UserEmails**](UserApi.md#PostApiV4UserEmails) | **Post** /api/v4/user/emails | 
[**PostApiV4UserGpgKeys**](UserApi.md#PostApiV4UserGpgKeys) | **Post** /api/v4/user/gpg_keys | Add a new GPG key to the currently authenticated user
[**PostApiV4UserGpgKeysKeyIdRevoke**](UserApi.md#PostApiV4UserGpgKeysKeyIdRevoke) | **Post** /api/v4/user/gpg_keys/{key_id}/revoke | Revoke a GPG key owned by currently authenticated user
[**PostApiV4UserKeys**](UserApi.md#PostApiV4UserKeys) | **Post** /api/v4/user/keys | 
[**PostApiV4UserPersonalAccessTokens**](UserApi.md#PostApiV4UserPersonalAccessTokens) | **Post** /api/v4/user/personal_access_tokens | Create a personal access token with limited scopes for the currently authenticated user
[**PostApiV4UserRunners**](UserApi.md#PostApiV4UserRunners) | **Post** /api/v4/user/runners | Create a runner owned by currently authenticated user
[**PutApiV4UserAvatar**](UserApi.md#PutApiV4UserAvatar) | **Put** /api/v4/user/avatar | Set the avatar of the current user
[**PutApiV4UserPreferences**](UserApi.md#PutApiV4UserPreferences) | **Put** /api/v4/user/preferences | Update the current user&#39;s preferences
[**PutApiV4UserStatus**](UserApi.md#PutApiV4UserStatus) | **Put** /api/v4/user/status | Set the status of the current user
[**PutApiV4UserUserIdCreditCardValidation**](UserApi.md#PutApiV4UserUserIdCreditCardValidation) | **Put** /api/v4/user/{user_id}/credit_card_validation | 


# **DeleteApiV4UserEmailsEmailId**
> DeleteApiV4UserEmailsEmailId(ctx, emailId)


Delete an email address from the currently authenticated user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **emailId** | **int32**| The ID of the email | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4UserGpgKeysKeyId**
> DeleteApiV4UserGpgKeysKeyId(ctx, keyId)
Delete a GPG key from the currently authenticated user

This feature was added in GitLab 10.0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **keyId** | **int32**| The ID of the SSH key | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4UserKeysKeyId**
> ApiEntitiesSshKey DeleteApiV4UserKeysKeyId(ctx, keyId)


Delete an SSH key from the currently authenticated user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **keyId** | **int32**| The ID of the SSH key | 

### Return type

[**ApiEntitiesSshKey**](API_Entities_SSHKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV3User**
> ApiEntitiesUserPublic GetApiV3User(ctx, )


Get the currently authenticated user

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiEntitiesUserPublic**](API_Entities_UserPublic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4UserActivities**
> GetApiV4UserActivities(ctx, optional)


Get a list of user activities

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiGetApiV4UserActivitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiGetApiV4UserActivitiesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **optional.Time**| Date string in the format YEAR-MONTH-DAY | [default to 2023-11-23T21:13:09.055Z]
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4UserEmails**
> ApiEntitiesEmail GetApiV4UserEmails(ctx, optional)


Get the currently authenticated user's email addresses

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiGetApiV4UserEmailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiGetApiV4UserEmailsOpts struct

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

# **GetApiV4UserEmailsEmailId**
> ApiEntitiesEmail GetApiV4UserEmailsEmailId(ctx, emailId)


Get a single email address owned by the currently authenticated user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **emailId** | **int32**| The ID of the email | 

### Return type

[**ApiEntitiesEmail**](API_Entities_Email.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4UserGpgKeys**
> ApiEntitiesGpgKey GetApiV4UserGpgKeys(ctx, optional)
Get the currently authenticated user's GPG keys

This feature was added in GitLab 10.0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiGetApiV4UserGpgKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiGetApiV4UserGpgKeysOpts struct

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

# **GetApiV4UserGpgKeysKeyId**
> ApiEntitiesGpgKey GetApiV4UserGpgKeysKeyId(ctx, keyId)
Get a single GPG key owned by currently authenticated user

This feature was added in GitLab 10.0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **keyId** | **int32**| The ID of the GPG key | 

### Return type

[**ApiEntitiesGpgKey**](API_Entities_GpgKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4UserKeys**
> ApiEntitiesSshKey GetApiV4UserKeys(ctx, optional)


Get the currently authenticated user's SSH keys

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiGetApiV4UserKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiGetApiV4UserKeysOpts struct

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

# **GetApiV4UserKeysKeyId**
> ApiEntitiesSshKey GetApiV4UserKeysKeyId(ctx, keyId)


Get a single key owned by currently authenticated user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **keyId** | **int32**| The ID of the SSH key | 

### Return type

[**ApiEntitiesSshKey**](API_Entities_SSHKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4UserPreferences**
> ApiEntitiesUserPreferences GetApiV4UserPreferences(ctx, )
Get the current user's preferences

This feature was introduced in GitLab 14.0.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiEntitiesUserPreferences**](API_Entities_UserPreferences.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4UserStatus**
> ApiEntitiesUserStatus GetApiV4UserStatus(ctx, )


get the status of the current user

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiEntitiesUserStatus**](API_Entities_UserStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchApiV4UserStatus**
> ApiEntitiesUserStatus PatchApiV4UserStatus(ctx, patchApiV4UserStatus)
Set the status of the current user

Any parameters that are not passed will be ignored.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **patchApiV4UserStatus** | [**PatchApiV4UserStatus**](PatchApiV4UserStatus.md)|  | 

### Return type

[**ApiEntitiesUserStatus**](API_Entities_UserStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4UserEmails**
> ApiEntitiesEmail PostApiV4UserEmails(ctx, postApiV4UserEmails)


Add new email address to the currently authenticated user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4UserEmails** | [**PostApiV4UserEmails**](PostApiV4UserEmails.md)|  | 

### Return type

[**ApiEntitiesEmail**](API_Entities_Email.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4UserGpgKeys**
> ApiEntitiesGpgKey PostApiV4UserGpgKeys(ctx, postApiV4UserGpgKeys)
Add a new GPG key to the currently authenticated user

This feature was added in GitLab 10.0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4UserGpgKeys** | [**PostApiV4UserGpgKeys**](PostApiV4UserGpgKeys.md)|  | 

### Return type

[**ApiEntitiesGpgKey**](API_Entities_GpgKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4UserGpgKeysKeyIdRevoke**
> PostApiV4UserGpgKeysKeyIdRevoke(ctx, keyId)
Revoke a GPG key owned by currently authenticated user

This feature was added in GitLab 10.0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **keyId** | **int32**| The ID of the GPG key | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4UserKeys**
> ApiEntitiesSshKey PostApiV4UserKeys(ctx, postApiV4UserKeys)


Add a new SSH key to the currently authenticated user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4UserKeys** | [**PostApiV4UserKeys**](PostApiV4UserKeys.md)|  | 

### Return type

[**ApiEntitiesSshKey**](API_Entities_SSHKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4UserPersonalAccessTokens**
> ApiEntitiesPersonalAccessTokenWithToken PostApiV4UserPersonalAccessTokens(ctx, postApiV4UserPersonalAccessTokens)
Create a personal access token with limited scopes for the currently authenticated user

This feature was introduced in GitLab 16.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4UserPersonalAccessTokens** | [**PostApiV4UserPersonalAccessTokens**](PostApiV4UserPersonalAccessTokens.md)|  | 

### Return type

[**ApiEntitiesPersonalAccessTokenWithToken**](API_Entities_PersonalAccessTokenWithToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4UserRunners**
> ApiEntitiesCiRunnerRegistrationDetails PostApiV4UserRunners(ctx, postApiV4UserRunners)
Create a runner owned by currently authenticated user

Create a new runner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4UserRunners** | [**PostApiV4UserRunners**](PostApiV4UserRunners.md)|  | 

### Return type

[**ApiEntitiesCiRunnerRegistrationDetails**](API_Entities_Ci_RunnerRegistrationDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4UserAvatar**
> ApiEntitiesAvatar PutApiV4UserAvatar(ctx, putApiV4UserAvatar)
Set the avatar of the current user

This feature was introduced in GitLab 17.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **putApiV4UserAvatar** | [**PutApiV4UserAvatar**](PutApiV4UserAvatar.md)|  | 

### Return type

[**ApiEntitiesAvatar**](API_Entities_Avatar.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4UserPreferences**
> ApiEntitiesUserPreferences PutApiV4UserPreferences(ctx, putApiV4UserPreferences)
Update the current user's preferences

This feature was introduced in GitLab 13.10.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **putApiV4UserPreferences** | [**PutApiV4UserPreferences**](PutApiV4UserPreferences.md)|  | 

### Return type

[**ApiEntitiesUserPreferences**](API_Entities_UserPreferences.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4UserStatus**
> ApiEntitiesUserStatus PutApiV4UserStatus(ctx, putApiV4UserStatus)
Set the status of the current user

Any parameters that are not passed will be nullified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **putApiV4UserStatus** | [**PutApiV4UserStatus**](PutApiV4UserStatus.md)|  | 

### Return type

[**ApiEntitiesUserStatus**](API_Entities_UserStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4UserUserIdCreditCardValidation**
> ApiEntitiesUserCreditCardValidations PutApiV4UserUserIdCreditCardValidation(ctx, userId, putApiV4UserUserIdCreditCardValidation)


Update a user's credit_card_validation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| The ID or username of the user | 
  **putApiV4UserUserIdCreditCardValidation** | [**PutApiV4UserUserIdCreditCardValidation**](PutApiV4UserUserIdCreditCardValidation.md)|  | 

### Return type

[**ApiEntitiesUserCreditCardValidations**](API_Entities_UserCreditCardValidations.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

