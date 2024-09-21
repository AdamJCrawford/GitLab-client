# \KeysApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4Keys**](KeysApi.md#GetApiV4Keys) | **Get** /api/v4/keys | Get user by fingerprint of SSH key
[**GetApiV4KeysId**](KeysApi.md#GetApiV4KeysId) | **Get** /api/v4/keys/{id} | Get single ssh key by id. Only available to admin users


# **GetApiV4Keys**
> ApiEntitiesUserWithAdmin GetApiV4Keys(ctx, fingerprint)
Get user by fingerprint of SSH key

You can search for a user that owns a specific SSH key. Note only administrators can lookup SSH key\\         with the fingerprint of an SSH key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fingerprint** | **string**| The fingerprint of an SSH key | 

### Return type

[**ApiEntitiesUserWithAdmin**](API_Entities_UserWithAdmin.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4KeysId**
> ApiEntitiesSshKeyWithUser GetApiV4KeysId(ctx, id)
Get single ssh key by id. Only available to admin users

Get SSH key with user by ID of an SSH key. Note only administrators can lookup SSH key with user by ID\\         of an SSH key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of an SSH key | 

### Return type

[**ApiEntitiesSshKeyWithUser**](API_Entities_SSHKeyWithUser.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

