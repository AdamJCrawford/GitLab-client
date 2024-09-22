# \DeployKeysApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4ProjectsIdDeployKeysKeyId**](DeployKeysApi.md#DeleteApiV4ProjectsIdDeployKeysKeyId) | **Delete** /api/v4/projects/{id}/deploy_keys/{key_id} | Delete deploy key
[**GetApiV4DeployKeys**](DeployKeysApi.md#GetApiV4DeployKeys) | **Get** /api/v4/deploy_keys | List all deploy keys
[**GetApiV4ProjectsIdDeployKeys**](DeployKeysApi.md#GetApiV4ProjectsIdDeployKeys) | **Get** /api/v4/projects/{id}/deploy_keys | List deploy keys for project
[**GetApiV4ProjectsIdDeployKeysKeyId**](DeployKeysApi.md#GetApiV4ProjectsIdDeployKeysKeyId) | **Get** /api/v4/projects/{id}/deploy_keys/{key_id} | Get a single deploy key
[**PostApiV4ProjectsIdDeployKeys**](DeployKeysApi.md#PostApiV4ProjectsIdDeployKeys) | **Post** /api/v4/projects/{id}/deploy_keys | Add deploy key
[**PostApiV4ProjectsIdDeployKeysKeyIdEnable**](DeployKeysApi.md#PostApiV4ProjectsIdDeployKeysKeyIdEnable) | **Post** /api/v4/projects/{id}/deploy_keys/{key_id}/enable | Enable a deploy key
[**PutApiV4ProjectsIdDeployKeysKeyId**](DeployKeysApi.md#PutApiV4ProjectsIdDeployKeysKeyId) | **Put** /api/v4/projects/{id}/deploy_keys/{key_id} | Update deploy key


# **DeleteApiV4ProjectsIdDeployKeysKeyId**
> DeleteApiV4ProjectsIdDeployKeysKeyId(ctx, id, keyId)
Delete deploy key

Removes a deploy key from the project. If the deploy key is used only for this project, it's deleted from the system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **keyId** | **int32**| The ID of the deploy key | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4DeployKeys**
> []ApiEntitiesDeployKey GetApiV4DeployKeys(ctx, optional)
List all deploy keys

Get a list of all deploy keys across all projects of the GitLab instance. This endpoint requires administrator access and is not available on GitLab.com.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeployKeysApiGetApiV4DeployKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeployKeysApiGetApiV4DeployKeysOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **public** | **optional.Bool**| Only return deploy keys that are public | [default to false]

### Return type

[**[]ApiEntitiesDeployKey**](API_Entities_DeployKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdDeployKeys**
> []ApiEntitiesDeployKeysProject GetApiV4ProjectsIdDeployKeys(ctx, id, optional)
List deploy keys for project

Get a list of a project's deploy keys.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
 **optional** | ***DeployKeysApiGetApiV4ProjectsIdDeployKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeployKeysApiGetApiV4ProjectsIdDeployKeysOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesDeployKeysProject**](API_Entities_DeployKeysProject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdDeployKeysKeyId**
> ApiEntitiesDeployKeysProject GetApiV4ProjectsIdDeployKeysKeyId(ctx, id, keyId)
Get a single deploy key

Get a single key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **keyId** | **int32**| The ID of the deploy key | 

### Return type

[**ApiEntitiesDeployKeysProject**](API_Entities_DeployKeysProject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdDeployKeys**
> ApiEntitiesDeployKeysProject PostApiV4ProjectsIdDeployKeys(ctx, id, postApiV4ProjectsIdDeployKeys)
Add deploy key

Creates a new deploy key for a project. If the deploy key already exists in another project, it's joined to the current project only if the original one is accessible by the same user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **postApiV4ProjectsIdDeployKeys** | [**PostApiV4ProjectsIdDeployKeys**](PostApiV4ProjectsIdDeployKeys.md)|  | 

### Return type

[**ApiEntitiesDeployKeysProject**](API_Entities_DeployKeysProject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdDeployKeysKeyIdEnable**
> ApiEntitiesDeployKey PostApiV4ProjectsIdDeployKeysKeyIdEnable(ctx, id, keyId)
Enable a deploy key

Enables a deploy key for a project so this can be used. Returns the enabled key, with a status code 201 when successful. This feature was added in GitLab 8.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **keyId** | **int32**| The ID of the deploy key | 

### Return type

[**ApiEntitiesDeployKey**](API_Entities_DeployKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdDeployKeysKeyId**
> ApiEntitiesDeployKey PutApiV4ProjectsIdDeployKeysKeyId(ctx, id, keyId, putApiV4ProjectsIdDeployKeysKeyId)
Update deploy key

Updates a deploy key for a project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **keyId** | **int32**| The ID of the deploy key | 
  **putApiV4ProjectsIdDeployKeysKeyId** | [**PutApiV4ProjectsIdDeployKeysKeyId**](PutApiV4ProjectsIdDeployKeysKeyId.md)|  | 

### Return type

[**ApiEntitiesDeployKey**](API_Entities_DeployKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

