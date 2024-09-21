# \ProjectRepositoryStorageMovesApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4ProjectRepositoryStorageMoves**](ProjectRepositoryStorageMovesApi.md#GetApiV4ProjectRepositoryStorageMoves) | **Get** /api/v4/project_repository_storage_moves | Get a list of all project repository storage moves
[**GetApiV4ProjectRepositoryStorageMovesRepositoryStorageMoveId**](ProjectRepositoryStorageMovesApi.md#GetApiV4ProjectRepositoryStorageMovesRepositoryStorageMoveId) | **Get** /api/v4/project_repository_storage_moves/{repository_storage_move_id} | Get a project repository storage move
[**PostApiV4ProjectRepositoryStorageMoves**](ProjectRepositoryStorageMovesApi.md#PostApiV4ProjectRepositoryStorageMoves) | **Post** /api/v4/project_repository_storage_moves | Schedule bulk project repository storage moves


# **GetApiV4ProjectRepositoryStorageMoves**
> []ApiEntitiesProjectsRepositoryStorageMove GetApiV4ProjectRepositoryStorageMoves(ctx, optional)
Get a list of all project repository storage moves

This feature was introduced in GitLab 13.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProjectRepositoryStorageMovesApiGetApiV4ProjectRepositoryStorageMovesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectRepositoryStorageMovesApiGetApiV4ProjectRepositoryStorageMovesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesProjectsRepositoryStorageMove**](API_Entities_Projects_RepositoryStorageMove.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectRepositoryStorageMovesRepositoryStorageMoveId**
> ApiEntitiesProjectsRepositoryStorageMove GetApiV4ProjectRepositoryStorageMovesRepositoryStorageMoveId(ctx, repositoryStorageMoveId)
Get a project repository storage move

This feature was introduced in GitLab 13.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryStorageMoveId** | **int32**| The ID of a project repository storage move | 

### Return type

[**ApiEntitiesProjectsRepositoryStorageMove**](API_Entities_Projects_RepositoryStorageMove.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectRepositoryStorageMoves**
> PostApiV4ProjectRepositoryStorageMoves(ctx, postApiV4ProjectRepositoryStorageMoves)
Schedule bulk project repository storage moves

This feature was introduced in GitLab 13.7.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4ProjectRepositoryStorageMoves** | [**PostApiV4ProjectRepositoryStorageMoves**](PostApiV4ProjectRepositoryStorageMoves.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

