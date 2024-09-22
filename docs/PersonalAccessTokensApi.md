# \PersonalAccessTokensApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4PersonalAccessTokensId**](PersonalAccessTokensApi.md#DeleteApiV4PersonalAccessTokensId) | **Delete** /api/v4/personal_access_tokens/{id} | Revoke a personal access token
[**DeleteApiV4PersonalAccessTokensSelf**](PersonalAccessTokensApi.md#DeleteApiV4PersonalAccessTokensSelf) | **Delete** /api/v4/personal_access_tokens/self | Revoke a personal access token
[**GetApiV4PersonalAccessTokens**](PersonalAccessTokensApi.md#GetApiV4PersonalAccessTokens) | **Get** /api/v4/personal_access_tokens | List personal access tokens
[**GetApiV4PersonalAccessTokensId**](PersonalAccessTokensApi.md#GetApiV4PersonalAccessTokensId) | **Get** /api/v4/personal_access_tokens/{id} | Get single personal access token
[**GetApiV4PersonalAccessTokensSelf**](PersonalAccessTokensApi.md#GetApiV4PersonalAccessTokensSelf) | **Get** /api/v4/personal_access_tokens/self | Get single personal access token
[**PostApiV4PersonalAccessTokensIdRotate**](PersonalAccessTokensApi.md#PostApiV4PersonalAccessTokensIdRotate) | **Post** /api/v4/personal_access_tokens/{id}/rotate | Rotate personal access token
[**PostApiV4PersonalAccessTokensSelfRotate**](PersonalAccessTokensApi.md#PostApiV4PersonalAccessTokensSelfRotate) | **Post** /api/v4/personal_access_tokens/self/rotate | Rotate a personal access token


# **DeleteApiV4PersonalAccessTokensId**
> DeleteApiV4PersonalAccessTokensId(ctx, id)
Revoke a personal access token

Revoke a personal access token by using the ID of the personal access token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4PersonalAccessTokensSelf**
> DeleteApiV4PersonalAccessTokensSelf(ctx, )
Revoke a personal access token

Revoke a personal access token by passing it to the API in a header

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4PersonalAccessTokens**
> []ApiEntitiesPersonalAccessToken GetApiV4PersonalAccessTokens(ctx, optional)
List personal access tokens

Get all personal access tokens the authenticated user has access to.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PersonalAccessTokensApiGetApiV4PersonalAccessTokensOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PersonalAccessTokensApiGetApiV4PersonalAccessTokensOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **optional.Int32**| Filter PATs by User ID | 
 **revoked** | **optional.Bool**| Filter PATs where revoked state matches parameter | 
 **state** | **optional.String**| Filter PATs which are either active or not | 
 **createdBefore** | **optional.Time**| Filter PATs which were created before given datetime | 
 **createdAfter** | **optional.Time**| Filter PATs which were created after given datetime | 
 **lastUsedBefore** | **optional.Time**| Filter PATs which were used before given datetime | 
 **lastUsedAfter** | **optional.Time**| Filter PATs which were used after given datetime | 
 **search** | **optional.String**| Filters PATs by its name | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesPersonalAccessToken**](API_Entities_PersonalAccessToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4PersonalAccessTokensId**
> ApiEntitiesPersonalAccessToken GetApiV4PersonalAccessTokensId(ctx, id)
Get single personal access token

Get a personal access token by using the ID of the personal access token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**ApiEntitiesPersonalAccessToken**](API_Entities_PersonalAccessToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4PersonalAccessTokensSelf**
> ApiEntitiesPersonalAccessToken GetApiV4PersonalAccessTokensSelf(ctx, )
Get single personal access token

Get the details of a personal access token by passing it to the API in a header

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiEntitiesPersonalAccessToken**](API_Entities_PersonalAccessToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4PersonalAccessTokensIdRotate**
> ApiEntitiesPersonalAccessTokenWithToken PostApiV4PersonalAccessTokensIdRotate(ctx, id, postApiV4PersonalAccessTokensIdRotate)
Rotate personal access token

Roates a personal access token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **postApiV4PersonalAccessTokensIdRotate** | [**PostApiV4PersonalAccessTokensIdRotate**](PostApiV4PersonalAccessTokensIdRotate.md)|  | 

### Return type

[**ApiEntitiesPersonalAccessTokenWithToken**](API_Entities_PersonalAccessTokenWithToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4PersonalAccessTokensSelfRotate**
> ApiEntitiesPersonalAccessTokenWithToken PostApiV4PersonalAccessTokensSelfRotate(ctx, postApiV4PersonalAccessTokensSelfRotate)
Rotate a personal access token

Rotates a personal access token by passing it to the API in a header

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4PersonalAccessTokensSelfRotate** | [**PostApiV4PersonalAccessTokensSelfRotate**](PostApiV4PersonalAccessTokensSelfRotate.md)|  | 

### Return type

[**ApiEntitiesPersonalAccessTokenWithToken**](API_Entities_PersonalAccessTokenWithToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

