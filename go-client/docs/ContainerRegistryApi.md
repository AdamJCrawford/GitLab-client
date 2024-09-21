# \ContainerRegistryApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4ProjectsIdRegistryRepositoriesRepositoryId**](ContainerRegistryApi.md#DeleteApiV4ProjectsIdRegistryRepositoriesRepositoryId) | **Delete** /api/v4/projects/{id}/registry/repositories/{repository_id} | Delete repository
[**DeleteApiV4ProjectsIdRegistryRepositoriesRepositoryIdTags**](ContainerRegistryApi.md#DeleteApiV4ProjectsIdRegistryRepositoriesRepositoryIdTags) | **Delete** /api/v4/projects/{id}/registry/repositories/{repository_id}/tags | Delete repository tags (in bulk)
[**DeleteApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsTagName**](ContainerRegistryApi.md#DeleteApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsTagName) | **Delete** /api/v4/projects/{id}/registry/repositories/{repository_id}/tags/{tag_name} | Delete repository tag
[**GetApiV4GroupsIdRegistryRepositories**](ContainerRegistryApi.md#GetApiV4GroupsIdRegistryRepositories) | **Get** /api/v4/groups/{id}/registry/repositories | List registry repositories within a group
[**GetApiV4ProjectsIdRegistryRepositories**](ContainerRegistryApi.md#GetApiV4ProjectsIdRegistryRepositories) | **Get** /api/v4/projects/{id}/registry/repositories | List container repositories within a project
[**GetApiV4ProjectsIdRegistryRepositoriesRepositoryIdTags**](ContainerRegistryApi.md#GetApiV4ProjectsIdRegistryRepositoriesRepositoryIdTags) | **Get** /api/v4/projects/{id}/registry/repositories/{repository_id}/tags | List tags of a repository
[**GetApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsTagName**](ContainerRegistryApi.md#GetApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsTagName) | **Get** /api/v4/projects/{id}/registry/repositories/{repository_id}/tags/{tag_name} | Get details about a repository tag
[**GetApiV4RegistryRepositoriesId**](ContainerRegistryApi.md#GetApiV4RegistryRepositoriesId) | **Get** /api/v4/registry/repositories/{id} | Get a container repository


# **DeleteApiV4ProjectsIdRegistryRepositoriesRepositoryId**
> DeleteApiV4ProjectsIdRegistryRepositoriesRepositoryId(ctx, id, repositoryId)
Delete repository

This feature was introduced in GitLab 11.8.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **repositoryId** | **int32**| The ID of the repository | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdRegistryRepositoriesRepositoryIdTags**
> DeleteApiV4ProjectsIdRegistryRepositoriesRepositoryIdTags(ctx, id, repositoryId, optional)
Delete repository tags (in bulk)

This feature was introduced in GitLab 11.8.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **repositoryId** | **int32**| The ID of the repository | 
 **optional** | ***ContainerRegistryApiDeleteApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainerRegistryApiDeleteApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nameRegexDelete** | **optional.String**| The tag name regexp to delete, specify .* to delete all | 
 **nameRegex** | **optional.String**| The tag name regexp to delete, specify .* to delete all | 
 **nameRegexKeep** | **optional.String**| The tag name regexp to retain | 
 **keepN** | **optional.Int32**| Keep n of latest tags with matching name | 
 **olderThan** | **optional.String**| Delete older than: 1h, 1d, 1month | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsTagName**
> DeleteApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsTagName(ctx, id, repositoryId, tagName)
Delete repository tag

This feature was introduced in GitLab 11.8.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **repositoryId** | **int32**| The ID of the repository | 
  **tagName** | **string**| The name of the tag | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdRegistryRepositories**
> []ApiEntitiesContainerRegistryRepository GetApiV4GroupsIdRegistryRepositories(ctx, id, optional)
List registry repositories within a group

Get a list of registry repositories in a group. This feature was introduced in GitLab 12.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the group accessible by the authenticated user | 
 **optional** | ***ContainerRegistryApiGetApiV4GroupsIdRegistryRepositoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainerRegistryApiGetApiV4GroupsIdRegistryRepositoriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesContainerRegistryRepository**](API_Entities_ContainerRegistry_Repository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdRegistryRepositories**
> []ApiEntitiesContainerRegistryRepository GetApiV4ProjectsIdRegistryRepositories(ctx, id, optional)
List container repositories within a project

This feature was introduced in GitLab 11.8.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
 **optional** | ***ContainerRegistryApiGetApiV4ProjectsIdRegistryRepositoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainerRegistryApiGetApiV4ProjectsIdRegistryRepositoriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **tags** | **optional.Bool**| Determines if tags should be included | [default to false]
 **tagsCount** | **optional.Bool**| Determines if the tags count should be included | [default to false]

### Return type

[**[]ApiEntitiesContainerRegistryRepository**](API_Entities_ContainerRegistry_Repository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdRegistryRepositoriesRepositoryIdTags**
> []ApiEntitiesContainerRegistryTag GetApiV4ProjectsIdRegistryRepositoriesRepositoryIdTags(ctx, id, repositoryId, optional)
List tags of a repository

This feature was introduced in GitLab 11.8.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **repositoryId** | **int32**| The ID of the repository | 
 **optional** | ***ContainerRegistryApiGetApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainerRegistryApiGetApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesContainerRegistryTag**](API_Entities_ContainerRegistry_Tag.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsTagName**
> ApiEntitiesContainerRegistryTagDetails GetApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsTagName(ctx, id, repositoryId, tagName)
Get details about a repository tag

This feature was introduced in GitLab 11.8.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **repositoryId** | **int32**| The ID of the repository | 
  **tagName** | **string**| The name of the tag | 

### Return type

[**ApiEntitiesContainerRegistryTagDetails**](API_Entities_ContainerRegistry_TagDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4RegistryRepositoriesId**
> ApiEntitiesContainerRegistryRepository GetApiV4RegistryRepositoriesId(ctx, id, optional)
Get a container repository

This feature was introduced in GitLab 13.6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the repository | 
 **optional** | ***ContainerRegistryApiGetApiV4RegistryRepositoriesIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainerRegistryApiGetApiV4RegistryRepositoriesIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tags** | **optional.Bool**| Determines if tags should be included | [default to false]
 **tagsCount** | **optional.Bool**| Determines if the tags count should be included | [default to false]
 **size** | **optional.Bool**| Determines if the size should be included | [default to false]

### Return type

[**ApiEntitiesContainerRegistryRepository**](API_Entities_ContainerRegistry_Repository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

