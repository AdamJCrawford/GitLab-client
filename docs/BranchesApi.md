# \BranchesApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4ProjectsIdRepositoryBranchesBranch**](BranchesApi.md#DeleteApiV4ProjectsIdRepositoryBranchesBranch) | **Delete** /api/v4/projects/{id}/repository/branches/{branch} | 
[**DeleteApiV4ProjectsIdRepositoryMergedBranches**](BranchesApi.md#DeleteApiV4ProjectsIdRepositoryMergedBranches) | **Delete** /api/v4/projects/{id}/repository/merged_branches | 
[**GetApiV4ProjectsIdRepositoryBranches**](BranchesApi.md#GetApiV4ProjectsIdRepositoryBranches) | **Get** /api/v4/projects/{id}/repository/branches | 
[**GetApiV4ProjectsIdRepositoryBranchesBranch**](BranchesApi.md#GetApiV4ProjectsIdRepositoryBranchesBranch) | **Get** /api/v4/projects/{id}/repository/branches/{branch} | 
[**HeadApiV4ProjectsIdRepositoryBranchesBranch**](BranchesApi.md#HeadApiV4ProjectsIdRepositoryBranchesBranch) | **Head** /api/v4/projects/{id}/repository/branches/{branch} | 
[**PostApiV4ProjectsIdRepositoryBranches**](BranchesApi.md#PostApiV4ProjectsIdRepositoryBranches) | **Post** /api/v4/projects/{id}/repository/branches | 
[**PutApiV4ProjectsIdRepositoryBranchesBranchProtect**](BranchesApi.md#PutApiV4ProjectsIdRepositoryBranchesBranchProtect) | **Put** /api/v4/projects/{id}/repository/branches/{branch}/protect | 
[**PutApiV4ProjectsIdRepositoryBranchesBranchUnprotect**](BranchesApi.md#PutApiV4ProjectsIdRepositoryBranchesBranchUnprotect) | **Put** /api/v4/projects/{id}/repository/branches/{branch}/unprotect | 


# **DeleteApiV4ProjectsIdRepositoryBranchesBranch**
> DeleteApiV4ProjectsIdRepositoryBranchesBranch(ctx, id, branch)


Delete a branch

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **branch** | **string**| The name of the branch | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdRepositoryMergedBranches**
> DeleteApiV4ProjectsIdRepositoryMergedBranches(ctx, id)


Delete all merged branches

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdRepositoryBranches**
> []ApiEntitiesBranch GetApiV4ProjectsIdRepositoryBranches(ctx, id, optional)


Get a project repository branches

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
 **optional** | ***BranchesApiGetApiV4ProjectsIdRepositoryBranchesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BranchesApiGetApiV4ProjectsIdRepositoryBranchesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **search** | **optional.String**| Return list of branches matching the search criteria | 
 **regex** | **optional.String**| Return list of branches matching the regex | 
 **sort** | **optional.String**| Return list of branches sorted by the given field | 
 **pageToken** | **optional.String**| Name of branch to start the pagination from | 

### Return type

[**[]ApiEntitiesBranch**](API_Entities_Branch.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdRepositoryBranchesBranch**
> ApiEntitiesBranch GetApiV4ProjectsIdRepositoryBranchesBranch(ctx, id, branch)


Get a single repository branch

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **branch** | **int32**|  | 

### Return type

[**ApiEntitiesBranch**](API_Entities_Branch.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HeadApiV4ProjectsIdRepositoryBranchesBranch**
> HeadApiV4ProjectsIdRepositoryBranchesBranch(ctx, id, branch)


Check if a branch exists

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **branch** | **string**| The name of the branch | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdRepositoryBranches**
> ApiEntitiesBranch PostApiV4ProjectsIdRepositoryBranches(ctx, id, postApiV4ProjectsIdRepositoryBranches)


Create branch

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **postApiV4ProjectsIdRepositoryBranches** | [**PostApiV4ProjectsIdRepositoryBranches**](PostApiV4ProjectsIdRepositoryBranches.md)|  | 

### Return type

[**ApiEntitiesBranch**](API_Entities_Branch.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdRepositoryBranchesBranchProtect**
> ApiEntitiesBranch PutApiV4ProjectsIdRepositoryBranchesBranchProtect(ctx, id, branch, putApiV4ProjectsIdRepositoryBranchesBranchProtect)


Protect a single branch

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **branch** | **string**| The name of the branch | 
  **putApiV4ProjectsIdRepositoryBranchesBranchProtect** | [**PutApiV4ProjectsIdRepositoryBranchesBranchProtect**](PutApiV4ProjectsIdRepositoryBranchesBranchProtect.md)|  | 

### Return type

[**ApiEntitiesBranch**](API_Entities_Branch.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdRepositoryBranchesBranchUnprotect**
> ApiEntitiesBranch PutApiV4ProjectsIdRepositoryBranchesBranchUnprotect(ctx, id, branch)


Unprotect a single branch

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **branch** | **string**| The name of the branch | 

### Return type

[**ApiEntitiesBranch**](API_Entities_Branch.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

