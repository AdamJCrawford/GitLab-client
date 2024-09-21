# \GroupAccessTokensApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4GroupsIdAccessTokensTokenId**](GroupAccessTokensApi.md#DeleteApiV4GroupsIdAccessTokensTokenId) | **Delete** /api/v4/groups/{id}/access_tokens/{token_id} | Revoke a resource access token
[**GetApiV4GroupsIdAccessTokens**](GroupAccessTokensApi.md#GetApiV4GroupsIdAccessTokens) | **Get** /api/v4/groups/{id}/access_tokens | Get list of all access tokens for the specified resource
[**GetApiV4GroupsIdAccessTokensTokenId**](GroupAccessTokensApi.md#GetApiV4GroupsIdAccessTokensTokenId) | **Get** /api/v4/groups/{id}/access_tokens/{token_id} | Get an access token for the specified resource by ID
[**PostApiV4GroupsIdAccessTokens**](GroupAccessTokensApi.md#PostApiV4GroupsIdAccessTokens) | **Post** /api/v4/groups/{id}/access_tokens | Create a resource access token
[**PostApiV4GroupsIdAccessTokensTokenIdRotate**](GroupAccessTokensApi.md#PostApiV4GroupsIdAccessTokensTokenIdRotate) | **Post** /api/v4/groups/{id}/access_tokens/{token_id}/rotate | Rotate a resource access token


# **DeleteApiV4GroupsIdAccessTokensTokenId**
> DeleteApiV4GroupsIdAccessTokensTokenId(ctx, id, tokenId)
Revoke a resource access token

This feature was introduced in GitLab 13.9.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The group ID | 
  **tokenId** | **string**| The ID of the token | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdAccessTokens**
> []ApiEntitiesResourceAccessToken GetApiV4GroupsIdAccessTokens(ctx, id)
Get list of all access tokens for the specified resource

This feature was introduced in GitLab 13.9.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID or URL-encoded path of the group | 

### Return type

[**[]ApiEntitiesResourceAccessToken**](API_Entities_ResourceAccessToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdAccessTokensTokenId**
> ApiEntitiesResourceAccessToken GetApiV4GroupsIdAccessTokensTokenId(ctx, id, tokenId)
Get an access token for the specified resource by ID

This feature was introduced in GitLab 14.10.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID or URL-encoded path of the group | 
  **tokenId** | **string**| The ID of the token | 

### Return type

[**ApiEntitiesResourceAccessToken**](API_Entities_ResourceAccessToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4GroupsIdAccessTokens**
> ApiEntitiesResourceAccessTokenWithToken PostApiV4GroupsIdAccessTokens(ctx, id, postApiV4GroupsIdAccessTokens)
Create a resource access token

This feature was introduced in GitLab 13.9.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The group ID | 
  **postApiV4GroupsIdAccessTokens** | [**PostApiV4GroupsIdAccessTokens**](PostApiV4GroupsIdAccessTokens.md)|  | 

### Return type

[**ApiEntitiesResourceAccessTokenWithToken**](API_Entities_ResourceAccessTokenWithToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4GroupsIdAccessTokensTokenIdRotate**
> ApiEntitiesResourceAccessTokenWithToken PostApiV4GroupsIdAccessTokensTokenIdRotate(ctx, id, tokenId, postApiV4GroupsIdAccessTokensTokenIdRotate)
Rotate a resource access token

This feature was introduced in GitLab 16.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The group ID | 
  **tokenId** | **string**| The ID of the token | 
  **postApiV4GroupsIdAccessTokensTokenIdRotate** | [**PostApiV4GroupsIdAccessTokensTokenIdRotate**](PostApiV4GroupsIdAccessTokensTokenIdRotate.md)|  | 

### Return type

[**ApiEntitiesResourceAccessTokenWithToken**](API_Entities_ResourceAccessTokenWithToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

