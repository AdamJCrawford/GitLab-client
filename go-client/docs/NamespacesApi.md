# \NamespacesApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4NamespacesIdStorageLimitExclusion**](NamespacesApi.md#DeleteApiV4NamespacesIdStorageLimitExclusion) | **Delete** /api/v4/namespaces/{id}/storage/limit_exclusion | Removes a storage limit exclusion for a Namespace
[**GetApiV4Namespaces**](NamespacesApi.md#GetApiV4Namespaces) | **Get** /api/v4/namespaces | List namespaces
[**GetApiV4NamespacesId**](NamespacesApi.md#GetApiV4NamespacesId) | **Get** /api/v4/namespaces/{id} | Get namespace by ID
[**GetApiV4NamespacesIdExists**](NamespacesApi.md#GetApiV4NamespacesIdExists) | **Get** /api/v4/namespaces/{id}/exists | Get existence of a namespace
[**GetApiV4NamespacesIdGitlabSubscription**](NamespacesApi.md#GetApiV4NamespacesIdGitlabSubscription) | **Get** /api/v4/namespaces/{id}/gitlab_subscription | 
[**GetApiV4NamespacesStorageLimitExclusions**](NamespacesApi.md#GetApiV4NamespacesStorageLimitExclusions) | **Get** /api/v4/namespaces/storage/limit_exclusions | Retrieve all limit exclusions
[**PostApiV4NamespacesIdGitlabSubscription**](NamespacesApi.md#PostApiV4NamespacesIdGitlabSubscription) | **Post** /api/v4/namespaces/{id}/gitlab_subscription | 
[**PostApiV4NamespacesIdStorageLimitExclusion**](NamespacesApi.md#PostApiV4NamespacesIdStorageLimitExclusion) | **Post** /api/v4/namespaces/{id}/storage/limit_exclusion | Creates a storage limit exclusion for a Namespace
[**PutApiV4NamespacesId**](NamespacesApi.md#PutApiV4NamespacesId) | **Put** /api/v4/namespaces/{id} | 
[**PutApiV4NamespacesIdGitlabSubscription**](NamespacesApi.md#PutApiV4NamespacesIdGitlabSubscription) | **Put** /api/v4/namespaces/{id}/gitlab_subscription | 


# **DeleteApiV4NamespacesIdStorageLimitExclusion**
> DeleteApiV4NamespacesIdStorageLimitExclusion(ctx, id)
Removes a storage limit exclusion for a Namespace

Removes a Namespaces::Storage::LimitExclusion

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

# **GetApiV4Namespaces**
> []ApiEntitiesNamespace GetApiV4Namespaces(ctx, optional)
List namespaces

Get a list of the namespaces of the authenticated user. If the user is an administrator, a list of all namespaces in the GitLab instance is shown.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NamespacesApiGetApiV4NamespacesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NamespacesApiGetApiV4NamespacesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| Returns a list of namespaces the user is authorized to view based on the search criteria | 
 **ownedOnly** | **optional.Bool**| In GitLab 14.2 and later, returns a list of owned namespaces only | 
 **topLevelOnly** | **optional.Bool**| Only include top level namespaces | [default to false]
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **requestedHostedPlan** | **optional.String**| Name of the hosted plan requested by the customer | 

### Return type

[**[]ApiEntitiesNamespace**](API_Entities_Namespace.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4NamespacesId**
> ApiEntitiesNamespace GetApiV4NamespacesId(ctx, id)
Get namespace by ID

Get a namespace by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID or URL-encoded path of the namespace | 

### Return type

[**ApiEntitiesNamespace**](API_Entities_Namespace.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4NamespacesIdExists**
> ApiEntitiesNamespaceExistence GetApiV4NamespacesIdExists(ctx, id, optional)
Get existence of a namespace

Get existence of a namespace by path. Suggests a new namespace path that does not already exist.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Namespace’s path | 
 **optional** | ***NamespacesApiGetApiV4NamespacesIdExistsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NamespacesApiGetApiV4NamespacesIdExistsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentId** | **optional.Int32**| The ID of the parent namespace. If no ID is specified, only top-level namespaces are considered. | 

### Return type

[**ApiEntitiesNamespaceExistence**](API_Entities_NamespaceExistence.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4NamespacesIdGitlabSubscription**
> ApiEntitiesGitlabSubscription GetApiV4NamespacesIdGitlabSubscription(ctx, id)


Returns the subscription for the namespace

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**ApiEntitiesGitlabSubscription**](API_Entities_GitlabSubscription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4NamespacesStorageLimitExclusions**
> ApiEntitiesNamespacesStorageLimitExclusion GetApiV4NamespacesStorageLimitExclusions(ctx, optional)
Retrieve all limit exclusions

Gets all records for namespaces that have been excluded

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NamespacesApiGetApiV4NamespacesStorageLimitExclusionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NamespacesApiGetApiV4NamespacesStorageLimitExclusionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**ApiEntitiesNamespacesStorageLimitExclusion**](API_Entities_Namespaces_Storage_LimitExclusion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4NamespacesIdGitlabSubscription**
> ApiEntitiesGitlabSubscription PostApiV4NamespacesIdGitlabSubscription(ctx, id, postApiV4NamespacesIdGitlabSubscription)


Create a subscription for the namespace

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **postApiV4NamespacesIdGitlabSubscription** | [**PostApiV4NamespacesIdGitlabSubscription**](PostApiV4NamespacesIdGitlabSubscription.md)|  | 

### Return type

[**ApiEntitiesGitlabSubscription**](API_Entities_GitlabSubscription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4NamespacesIdStorageLimitExclusion**
> ApiEntitiesNamespacesStorageLimitExclusion PostApiV4NamespacesIdStorageLimitExclusion(ctx, id, postApiV4NamespacesIdStorageLimitExclusion)
Creates a storage limit exclusion for a Namespace

Creates a Namespaces::Storage::LimitExclusion

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **postApiV4NamespacesIdStorageLimitExclusion** | [**PostApiV4NamespacesIdStorageLimitExclusion**](PostApiV4NamespacesIdStorageLimitExclusion.md)|  | 

### Return type

[**ApiEntitiesNamespacesStorageLimitExclusion**](API_Entities_Namespaces_Storage_LimitExclusion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4NamespacesId**
> ApiEntitiesNamespace PutApiV4NamespacesId(ctx, id, putApiV4NamespacesId)


Update a namespace

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **putApiV4NamespacesId** | [**PutApiV4NamespacesId**](PutApiV4NamespacesId.md)|  | 

### Return type

[**ApiEntitiesNamespace**](API_Entities_Namespace.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4NamespacesIdGitlabSubscription**
> ApiEntitiesGitlabSubscription PutApiV4NamespacesIdGitlabSubscription(ctx, id, putApiV4NamespacesIdGitlabSubscription)


Update the subscription for the namespace

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **putApiV4NamespacesIdGitlabSubscription** | [**PutApiV4NamespacesIdGitlabSubscription**](PutApiV4NamespacesIdGitlabSubscription.md)|  | 

### Return type

[**ApiEntitiesGitlabSubscription**](API_Entities_GitlabSubscription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

