# \ProjectAccessTokensApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4ProjectsIdAccessTokensTokenId**](ProjectAccessTokensApi.md#DeleteApiV4ProjectsIdAccessTokensTokenId) | **Delete** /api/v4/projects/{id}/access_tokens/{token_id} | Revoke a resource access token
[**GetApiV4ProjectsIdAccessTokens**](ProjectAccessTokensApi.md#GetApiV4ProjectsIdAccessTokens) | **Get** /api/v4/projects/{id}/access_tokens | Get list of all access tokens for the specified resource
[**GetApiV4ProjectsIdAccessTokensTokenId**](ProjectAccessTokensApi.md#GetApiV4ProjectsIdAccessTokensTokenId) | **Get** /api/v4/projects/{id}/access_tokens/{token_id} | Get an access token for the specified resource by ID
[**PostApiV4ProjectsIdAccessTokens**](ProjectAccessTokensApi.md#PostApiV4ProjectsIdAccessTokens) | **Post** /api/v4/projects/{id}/access_tokens | Create a resource access token
[**PostApiV4ProjectsIdAccessTokensTokenIdRotate**](ProjectAccessTokensApi.md#PostApiV4ProjectsIdAccessTokensTokenIdRotate) | **Post** /api/v4/projects/{id}/access_tokens/{token_id}/rotate | Rotate a resource access token


# **DeleteApiV4ProjectsIdAccessTokensTokenId**
> DeleteApiV4ProjectsIdAccessTokensTokenId(ctx, id, tokenId)
Revoke a resource access token

This feature was introduced in GitLab 13.9.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID | 
  **tokenId** | **string**| The ID of the token | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdAccessTokens**
> []ApiEntitiesResourceAccessToken GetApiV4ProjectsIdAccessTokens(ctx, id)
Get list of all access tokens for the specified resource

This feature was introduced in GitLab 13.9.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID or URL-encoded path of the project | 

### Return type

[**[]ApiEntitiesResourceAccessToken**](API_Entities_ResourceAccessToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdAccessTokensTokenId**
> ApiEntitiesResourceAccessToken GetApiV4ProjectsIdAccessTokensTokenId(ctx, id, tokenId)
Get an access token for the specified resource by ID

This feature was introduced in GitLab 14.10.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID or URL-encoded path of the project | 
  **tokenId** | **string**| The ID of the token | 

### Return type

[**ApiEntitiesResourceAccessToken**](API_Entities_ResourceAccessToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdAccessTokens**
> ApiEntitiesResourceAccessTokenWithToken PostApiV4ProjectsIdAccessTokens(ctx, id, postApiV4ProjectsIdAccessTokens)
Create a resource access token

This feature was introduced in GitLab 13.9.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID | 
  **postApiV4ProjectsIdAccessTokens** | [**PostApiV4ProjectsIdAccessTokens**](PostApiV4ProjectsIdAccessTokens.md)|  | 

### Return type

[**ApiEntitiesResourceAccessTokenWithToken**](API_Entities_ResourceAccessTokenWithToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdAccessTokensTokenIdRotate**
> ApiEntitiesResourceAccessTokenWithToken PostApiV4ProjectsIdAccessTokensTokenIdRotate(ctx, id, tokenId, postApiV4ProjectsIdAccessTokensTokenIdRotate)
Rotate a resource access token

This feature was introduced in GitLab 16.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID | 
  **tokenId** | **string**| The ID of the token | 
  **postApiV4ProjectsIdAccessTokensTokenIdRotate** | [**PostApiV4ProjectsIdAccessTokensTokenIdRotate**](PostApiV4ProjectsIdAccessTokensTokenIdRotate.md)|  | 

### Return type

[**ApiEntitiesResourceAccessTokenWithToken**](API_Entities_ResourceAccessTokenWithToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

