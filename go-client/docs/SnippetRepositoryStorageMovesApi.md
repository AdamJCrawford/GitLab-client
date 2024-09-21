# \SnippetRepositoryStorageMovesApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4SnippetRepositoryStorageMoves**](SnippetRepositoryStorageMovesApi.md#GetApiV4SnippetRepositoryStorageMoves) | **Get** /api/v4/snippet_repository_storage_moves | Get a list of all snippet repository storage moves
[**GetApiV4SnippetRepositoryStorageMovesRepositoryStorageMoveId**](SnippetRepositoryStorageMovesApi.md#GetApiV4SnippetRepositoryStorageMovesRepositoryStorageMoveId) | **Get** /api/v4/snippet_repository_storage_moves/{repository_storage_move_id} | Get a snippet repository storage move
[**PostApiV4SnippetRepositoryStorageMoves**](SnippetRepositoryStorageMovesApi.md#PostApiV4SnippetRepositoryStorageMoves) | **Post** /api/v4/snippet_repository_storage_moves | Schedule bulk snippet repository storage moves


# **GetApiV4SnippetRepositoryStorageMoves**
> []ApiEntitiesSnippetsRepositoryStorageMove GetApiV4SnippetRepositoryStorageMoves(ctx, optional)
Get a list of all snippet repository storage moves

This feature was introduced in GitLab 13.8.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SnippetRepositoryStorageMovesApiGetApiV4SnippetRepositoryStorageMovesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnippetRepositoryStorageMovesApiGetApiV4SnippetRepositoryStorageMovesOpts struct

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

# **GetApiV4SnippetRepositoryStorageMovesRepositoryStorageMoveId**
> ApiEntitiesSnippetsRepositoryStorageMove GetApiV4SnippetRepositoryStorageMovesRepositoryStorageMoveId(ctx, repositoryStorageMoveId)
Get a snippet repository storage move

This feature was introduced in GitLab 13.8.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryStorageMoveId** | **int32**| The ID of a snippet repository storage move | 

### Return type

[**ApiEntitiesSnippetsRepositoryStorageMove**](API_Entities_Snippets_RepositoryStorageMove.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4SnippetRepositoryStorageMoves**
> PostApiV4SnippetRepositoryStorageMoves(ctx, postApiV4SnippetRepositoryStorageMoves)
Schedule bulk snippet repository storage moves

This feature was introduced in GitLab 13.8.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4SnippetRepositoryStorageMoves** | [**PostApiV4SnippetRepositoryStorageMoves**](PostApiV4SnippetRepositoryStorageMoves.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

