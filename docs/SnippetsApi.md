# \SnippetsApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4SnippetsId**](SnippetsApi.md#DeleteApiV4SnippetsId) | **Delete** /api/v4/snippets/{id} | Remove snippet
[**GetApiV4Snippets**](SnippetsApi.md#GetApiV4Snippets) | **Get** /api/v4/snippets | Get a snippets list for an authenticated user
[**GetApiV4SnippetsAll**](SnippetsApi.md#GetApiV4SnippetsAll) | **Get** /api/v4/snippets/all | List all snippets current_user has access to
[**GetApiV4SnippetsId**](SnippetsApi.md#GetApiV4SnippetsId) | **Get** /api/v4/snippets/{id} | Get a single snippet
[**GetApiV4SnippetsIdFilesRefFilePathRaw**](SnippetsApi.md#GetApiV4SnippetsIdFilesRefFilePathRaw) | **Get** /api/v4/snippets/{id}/files/{ref}/{file_path}/raw | 
[**GetApiV4SnippetsIdRaw**](SnippetsApi.md#GetApiV4SnippetsIdRaw) | **Get** /api/v4/snippets/{id}/raw | Get a raw snippet
[**GetApiV4SnippetsIdRepositoryStorageMoves**](SnippetsApi.md#GetApiV4SnippetsIdRepositoryStorageMoves) | **Get** /api/v4/snippets/{id}/repository_storage_moves | Get a list of all snippets repository storage moves
[**GetApiV4SnippetsIdRepositoryStorageMovesRepositoryStorageMoveId**](SnippetsApi.md#GetApiV4SnippetsIdRepositoryStorageMovesRepositoryStorageMoveId) | **Get** /api/v4/snippets/{id}/repository_storage_moves/{repository_storage_move_id} | Get a snippet repository storage move
[**GetApiV4SnippetsIdUserAgentDetail**](SnippetsApi.md#GetApiV4SnippetsIdUserAgentDetail) | **Get** /api/v4/snippets/{id}/user_agent_detail | 
[**GetApiV4SnippetsPublic**](SnippetsApi.md#GetApiV4SnippetsPublic) | **Get** /api/v4/snippets/public | List all public personal snippets current_user has access to
[**PostApiV4Snippets**](SnippetsApi.md#PostApiV4Snippets) | **Post** /api/v4/snippets | Create new snippet
[**PostApiV4SnippetsIdRepositoryStorageMoves**](SnippetsApi.md#PostApiV4SnippetsIdRepositoryStorageMoves) | **Post** /api/v4/snippets/{id}/repository_storage_moves | Schedule a snippet repository storage move
[**PutApiV4SnippetsId**](SnippetsApi.md#PutApiV4SnippetsId) | **Put** /api/v4/snippets/{id} | Update an existing snippet


# **DeleteApiV4SnippetsId**
> ApiEntitiesPersonalSnippet DeleteApiV4SnippetsId(ctx, id)
Remove snippet

This feature was introduced in GitLab 8.15.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of a snippet | 

### Return type

[**ApiEntitiesPersonalSnippet**](API_Entities_PersonalSnippet.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4Snippets**
> []ApiEntitiesSnippet GetApiV4Snippets(ctx, optional)
Get a snippets list for an authenticated user

This feature was introduced in GitLab 8.15.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SnippetsApiGetApiV4SnippetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnippetsApiGetApiV4SnippetsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdAfter** | **optional.Time**| Return snippets created after the specified time | 
 **createdBefore** | **optional.Time**| Return snippets created before the specified time | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesSnippet**](API_Entities_Snippet.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4SnippetsAll**
> []ApiEntitiesSnippet GetApiV4SnippetsAll(ctx, optional)
List all snippets current_user has access to

This feature was introduced in GitLab 16.3.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SnippetsApiGetApiV4SnippetsAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnippetsApiGetApiV4SnippetsAllOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdAfter** | **optional.Time**| Return snippets created after the specified time | 
 **createdBefore** | **optional.Time**| Return snippets created before the specified time | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **repositoryStorage** | **optional.String**| Filter by repository storage used by the snippet | 

### Return type

[**[]ApiEntitiesSnippet**](API_Entities_Snippet.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4SnippetsId**
> ApiEntitiesPersonalSnippet GetApiV4SnippetsId(ctx, id)
Get a single snippet

This feature was introduced in GitLab 8.15.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of a snippet | 

### Return type

[**ApiEntitiesPersonalSnippet**](API_Entities_PersonalSnippet.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4SnippetsIdFilesRefFilePathRaw**
> GetApiV4SnippetsIdFilesRefFilePathRaw(ctx, filePath, ref, id)


Get raw snippet file contents from the repository

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **filePath** | **string**| The url encoded path to the file, e.g. lib%2Fclass%2Erb | 
  **ref** | **string**| The name of branch, tag or commit | 
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4SnippetsIdRaw**
> GetApiV4SnippetsIdRaw(ctx, id)
Get a raw snippet

This feature was introduced in GitLab 8.15.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of a snippet | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4SnippetsIdRepositoryStorageMoves**
> []ApiEntitiesSnippetsRepositoryStorageMove GetApiV4SnippetsIdRepositoryStorageMoves(ctx, id, optional)
Get a list of all snippets repository storage moves

This feature was introduced in GitLab 13.8.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a snippet | 
 **optional** | ***SnippetsApiGetApiV4SnippetsIdRepositoryStorageMovesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnippetsApiGetApiV4SnippetsIdRepositoryStorageMovesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesSnippetsRepositoryStorageMove**](API_Entities_Snippets_RepositoryStorageMove.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4SnippetsIdRepositoryStorageMovesRepositoryStorageMoveId**
> ApiEntitiesSnippetsRepositoryStorageMove GetApiV4SnippetsIdRepositoryStorageMovesRepositoryStorageMoveId(ctx, id, repositoryStorageMoveId)
Get a snippet repository storage move

This feature was introduced in GitLab 13.8.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a snippet | 
  **repositoryStorageMoveId** | **int32**| The ID of a snippet repository storage move | 

### Return type

[**ApiEntitiesSnippetsRepositoryStorageMove**](API_Entities_Snippets_RepositoryStorageMove.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4SnippetsIdUserAgentDetail**
> ApiEntitiesUserAgentDetail GetApiV4SnippetsIdUserAgentDetail(ctx, id)


Get the user agent details for a snippet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of a snippet | 

### Return type

[**ApiEntitiesUserAgentDetail**](API_Entities_UserAgentDetail.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4SnippetsPublic**
> []ApiEntitiesPersonalSnippet GetApiV4SnippetsPublic(ctx, optional)
List all public personal snippets current_user has access to

This feature was introduced in GitLab 8.15.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SnippetsApiGetApiV4SnippetsPublicOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnippetsApiGetApiV4SnippetsPublicOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdAfter** | **optional.Time**| Return snippets created after the specified time | 
 **createdBefore** | **optional.Time**| Return snippets created before the specified time | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesPersonalSnippet**](API_Entities_PersonalSnippet.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4Snippets**
> ApiEntitiesPersonalSnippet PostApiV4Snippets(ctx, postApiV4Snippets)
Create new snippet

This feature was introduced in GitLab 8.15.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4Snippets** | [**PostApiV4Snippets**](PostApiV4Snippets.md)|  | 

### Return type

[**ApiEntitiesPersonalSnippet**](API_Entities_PersonalSnippet.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4SnippetsIdRepositoryStorageMoves**
> ApiEntitiesSnippetsRepositoryStorageMove PostApiV4SnippetsIdRepositoryStorageMoves(ctx, id, postApiV4SnippetsIdRepositoryStorageMoves)
Schedule a snippet repository storage move

This feature was introduced in GitLab 13.8.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a snippet | 
  **postApiV4SnippetsIdRepositoryStorageMoves** | [**PostApiV4SnippetsIdRepositoryStorageMoves**](PostApiV4SnippetsIdRepositoryStorageMoves.md)|  | 

### Return type

[**ApiEntitiesSnippetsRepositoryStorageMove**](API_Entities_Snippets_RepositoryStorageMove.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4SnippetsId**
> ApiEntitiesPersonalSnippet PutApiV4SnippetsId(ctx, id, putApiV4SnippetsId)
Update an existing snippet

This feature was introduced in GitLab 8.15.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of a snippet | 
  **putApiV4SnippetsId** | [**PutApiV4SnippetsId**](PutApiV4SnippetsId.md)|  | 

### Return type

[**ApiEntitiesPersonalSnippet**](API_Entities_PersonalSnippet.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

