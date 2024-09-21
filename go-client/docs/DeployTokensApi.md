# \DeployTokensApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4GroupsIdDeployTokensTokenId**](DeployTokensApi.md#DeleteApiV4GroupsIdDeployTokensTokenId) | **Delete** /api/v4/groups/{id}/deploy_tokens/{token_id} | Delete a group deploy token
[**DeleteApiV4ProjectsIdDeployTokensTokenId**](DeployTokensApi.md#DeleteApiV4ProjectsIdDeployTokensTokenId) | **Delete** /api/v4/projects/{id}/deploy_tokens/{token_id} | Delete a project deploy token
[**GetApiV4DeployTokens**](DeployTokensApi.md#GetApiV4DeployTokens) | **Get** /api/v4/deploy_tokens | List all deploy tokens
[**GetApiV4GroupsIdDeployTokens**](DeployTokensApi.md#GetApiV4GroupsIdDeployTokens) | **Get** /api/v4/groups/{id}/deploy_tokens | List group deploy tokens
[**GetApiV4GroupsIdDeployTokensTokenId**](DeployTokensApi.md#GetApiV4GroupsIdDeployTokensTokenId) | **Get** /api/v4/groups/{id}/deploy_tokens/{token_id} | Get a group deploy token
[**GetApiV4ProjectsIdDeployTokens**](DeployTokensApi.md#GetApiV4ProjectsIdDeployTokens) | **Get** /api/v4/projects/{id}/deploy_tokens | List project deploy tokens
[**GetApiV4ProjectsIdDeployTokensTokenId**](DeployTokensApi.md#GetApiV4ProjectsIdDeployTokensTokenId) | **Get** /api/v4/projects/{id}/deploy_tokens/{token_id} | Get a project deploy token
[**PostApiV4GroupsIdDeployTokens**](DeployTokensApi.md#PostApiV4GroupsIdDeployTokens) | **Post** /api/v4/groups/{id}/deploy_tokens | Create a group deploy token
[**PostApiV4ProjectsIdDeployTokens**](DeployTokensApi.md#PostApiV4ProjectsIdDeployTokens) | **Post** /api/v4/projects/{id}/deploy_tokens | Create a project deploy token


# **DeleteApiV4GroupsIdDeployTokensTokenId**
> DeleteApiV4GroupsIdDeployTokensTokenId(ctx, id, tokenId)
Delete a group deploy token

Removes a deploy token from the group. This feature was introduced in GitLab 12.9.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID or URL-encoded path of the group owned by the authenticated user | 
  **tokenId** | **int32**| The ID of the deploy token | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdDeployTokensTokenId**
> DeleteApiV4ProjectsIdDeployTokensTokenId(ctx, id, tokenId)
Delete a project deploy token

This feature was introduced in GitLab 12.9.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **tokenId** | **int32**| The ID of the deploy token | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4DeployTokens**
> []ApiEntitiesDeployToken GetApiV4DeployTokens(ctx, optional)
List all deploy tokens

Get a list of all deploy tokens across the GitLab instance. This endpoint requires administrator access. This feature was introduced in GitLab 12.9.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeployTokensApiGetApiV4DeployTokensOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeployTokensApiGetApiV4DeployTokensOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **active** | **optional.Bool**| Limit by active status | 

### Return type

[**[]ApiEntitiesDeployToken**](API_Entities_DeployToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdDeployTokens**
> []ApiEntitiesDeployToken GetApiV4GroupsIdDeployTokens(ctx, id, optional)
List group deploy tokens

Get a list of a group's deploy tokens. This feature was introduced in GitLab 12.9.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID or URL-encoded path of the group owned by the authenticated user | 
 **optional** | ***DeployTokensApiGetApiV4GroupsIdDeployTokensOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeployTokensApiGetApiV4GroupsIdDeployTokensOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **active** | **optional.Bool**| Limit by active status | 

### Return type

[**[]ApiEntitiesDeployToken**](API_Entities_DeployToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdDeployTokensTokenId**
> ApiEntitiesDeployToken GetApiV4GroupsIdDeployTokensTokenId(ctx, id, tokenId)
Get a group deploy token

Get a single group's deploy token by ID. This feature was introduced in GitLab 14.9. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID or URL-encoded path of the group owned by the authenticated user | 
  **tokenId** | **int32**| The ID of the deploy token | 

### Return type

[**ApiEntitiesDeployToken**](API_Entities_DeployToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdDeployTokens**
> []ApiEntitiesDeployToken GetApiV4ProjectsIdDeployTokens(ctx, id, optional)
List project deploy tokens

Get a list of a project's deploy tokens. This feature was introduced in GitLab 12.9.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
 **optional** | ***DeployTokensApiGetApiV4ProjectsIdDeployTokensOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeployTokensApiGetApiV4ProjectsIdDeployTokensOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **active** | **optional.Bool**| Limit by active status | 

### Return type

[**[]ApiEntitiesDeployToken**](API_Entities_DeployToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdDeployTokensTokenId**
> ApiEntitiesDeployToken GetApiV4ProjectsIdDeployTokensTokenId(ctx, id, tokenId)
Get a project deploy token

Get a single project's deploy token by ID. This feature was introduced in GitLab 14.9.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **tokenId** | **int32**| The ID of the deploy token | 

### Return type

[**ApiEntitiesDeployToken**](API_Entities_DeployToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4GroupsIdDeployTokens**
> ApiEntitiesDeployTokenWithToken PostApiV4GroupsIdDeployTokens(ctx, id, postApiV4GroupsIdDeployTokens)
Create a group deploy token

Creates a new deploy token for a group. This feature was introduced in GitLab 12.9.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID or URL-encoded path of the group owned by the authenticated user | 
  **postApiV4GroupsIdDeployTokens** | [**PostApiV4GroupsIdDeployTokens**](PostApiV4GroupsIdDeployTokens.md)|  | 

### Return type

[**ApiEntitiesDeployTokenWithToken**](API_Entities_DeployTokenWithToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdDeployTokens**
> ApiEntitiesDeployTokenWithToken PostApiV4ProjectsIdDeployTokens(ctx, id, postApiV4ProjectsIdDeployTokens)
Create a project deploy token

Creates a new deploy token for a project. This feature was introduced in GitLab 12.9.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **postApiV4ProjectsIdDeployTokens** | [**PostApiV4ProjectsIdDeployTokens**](PostApiV4ProjectsIdDeployTokens.md)|  | 

### Return type

[**ApiEntitiesDeployTokenWithToken**](API_Entities_DeployTokenWithToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

