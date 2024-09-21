# \CiVariablesApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4AdminCiVariablesKey**](CiVariablesApi.md#DeleteApiV4AdminCiVariablesKey) | **Delete** /api/v4/admin/ci/variables/{key} | 
[**DeleteApiV4GroupsIdVariablesKey**](CiVariablesApi.md#DeleteApiV4GroupsIdVariablesKey) | **Delete** /api/v4/groups/{id}/variables/{key} | 
[**DeleteApiV4ProjectsIdVariablesKey**](CiVariablesApi.md#DeleteApiV4ProjectsIdVariablesKey) | **Delete** /api/v4/projects/{id}/variables/{key} | 
[**GetApiV4AdminCiVariables**](CiVariablesApi.md#GetApiV4AdminCiVariables) | **Get** /api/v4/admin/ci/variables | 
[**GetApiV4AdminCiVariablesKey**](CiVariablesApi.md#GetApiV4AdminCiVariablesKey) | **Get** /api/v4/admin/ci/variables/{key} | 
[**GetApiV4GroupsIdVariables**](CiVariablesApi.md#GetApiV4GroupsIdVariables) | **Get** /api/v4/groups/{id}/variables | 
[**GetApiV4GroupsIdVariablesKey**](CiVariablesApi.md#GetApiV4GroupsIdVariablesKey) | **Get** /api/v4/groups/{id}/variables/{key} | 
[**GetApiV4ProjectsIdVariables**](CiVariablesApi.md#GetApiV4ProjectsIdVariables) | **Get** /api/v4/projects/{id}/variables | 
[**GetApiV4ProjectsIdVariablesKey**](CiVariablesApi.md#GetApiV4ProjectsIdVariablesKey) | **Get** /api/v4/projects/{id}/variables/{key} | 
[**PostApiV4AdminCiVariables**](CiVariablesApi.md#PostApiV4AdminCiVariables) | **Post** /api/v4/admin/ci/variables | 
[**PostApiV4GroupsIdVariables**](CiVariablesApi.md#PostApiV4GroupsIdVariables) | **Post** /api/v4/groups/{id}/variables | 
[**PostApiV4ProjectsIdVariables**](CiVariablesApi.md#PostApiV4ProjectsIdVariables) | **Post** /api/v4/projects/{id}/variables | 
[**PutApiV4AdminCiVariablesKey**](CiVariablesApi.md#PutApiV4AdminCiVariablesKey) | **Put** /api/v4/admin/ci/variables/{key} | 
[**PutApiV4GroupsIdVariablesKey**](CiVariablesApi.md#PutApiV4GroupsIdVariablesKey) | **Put** /api/v4/groups/{id}/variables/{key} | 
[**PutApiV4ProjectsIdVariablesKey**](CiVariablesApi.md#PutApiV4ProjectsIdVariablesKey) | **Put** /api/v4/projects/{id}/variables/{key} | 


# **DeleteApiV4AdminCiVariablesKey**
> ApiEntitiesCiVariable DeleteApiV4AdminCiVariablesKey(ctx, key)


Delete an existing instance-level variable

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| The key of a variable | 

### Return type

[**ApiEntitiesCiVariable**](API_Entities_Ci_Variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4GroupsIdVariablesKey**
> ApiEntitiesCiVariable DeleteApiV4GroupsIdVariablesKey(ctx, id, key)


Delete an existing variable from a group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group or URL-encoded path of the group owned by the authenticated       user | 
  **key** | **string**| The key of a variable | 

### Return type

[**ApiEntitiesCiVariable**](API_Entities_Ci_Variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdVariablesKey**
> ApiEntitiesCiVariable DeleteApiV4ProjectsIdVariablesKey(ctx, id, key, optional)


Delete an existing variable from a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a project or URL-encoded NAMESPACE/PROJECT_NAME of the project owned by the authenticated user | 
  **key** | **string**| The key of a variable | 
 **optional** | ***CiVariablesApiDeleteApiV4ProjectsIdVariablesKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CiVariablesApiDeleteApiV4ProjectsIdVariablesKeyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filterEnvironmentScope** | **optional.String**| The environment scope of the variable | 

### Return type

[**ApiEntitiesCiVariable**](API_Entities_Ci_Variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4AdminCiVariables**
> ApiEntitiesCiVariable GetApiV4AdminCiVariables(ctx, optional)


List all instance-level variables

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CiVariablesApiGetApiV4AdminCiVariablesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CiVariablesApiGetApiV4AdminCiVariablesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**ApiEntitiesCiVariable**](API_Entities_Ci_Variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4AdminCiVariablesKey**
> ApiEntitiesCiVariable GetApiV4AdminCiVariablesKey(ctx, key)


Get the details of a specific instance-level variable

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| The key of a variable | 

### Return type

[**ApiEntitiesCiVariable**](API_Entities_Ci_Variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdVariables**
> ApiEntitiesCiVariable GetApiV4GroupsIdVariables(ctx, id, optional)


Get a list of group-level variables

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group or URL-encoded path of the group owned by the authenticated       user | 
 **optional** | ***CiVariablesApiGetApiV4GroupsIdVariablesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CiVariablesApiGetApiV4GroupsIdVariablesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**ApiEntitiesCiVariable**](API_Entities_Ci_Variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdVariablesKey**
> ApiEntitiesCiVariable GetApiV4GroupsIdVariablesKey(ctx, id, key)


Get the details of a group’s specific variable

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group or URL-encoded path of the group owned by the authenticated       user | 
  **key** | **string**| The key of the variable | 

### Return type

[**ApiEntitiesCiVariable**](API_Entities_Ci_Variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdVariables**
> ApiEntitiesCiVariable GetApiV4ProjectsIdVariables(ctx, id, optional)


Get project variables

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a project or URL-encoded NAMESPACE/PROJECT_NAME of the project owned by the authenticated user | 
 **optional** | ***CiVariablesApiGetApiV4ProjectsIdVariablesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CiVariablesApiGetApiV4ProjectsIdVariablesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**ApiEntitiesCiVariable**](API_Entities_Ci_Variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdVariablesKey**
> ApiEntitiesCiVariable GetApiV4ProjectsIdVariablesKey(ctx, id, key, optional)


Get the details of a single variable from a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a project or URL-encoded NAMESPACE/PROJECT_NAME of the project owned by the authenticated user | 
  **key** | **string**| The key of a variable | 
 **optional** | ***CiVariablesApiGetApiV4ProjectsIdVariablesKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CiVariablesApiGetApiV4ProjectsIdVariablesKeyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filterEnvironmentScope** | **optional.String**| The environment scope of a variable | 

### Return type

[**ApiEntitiesCiVariable**](API_Entities_Ci_Variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4AdminCiVariables**
> ApiEntitiesCiVariable PostApiV4AdminCiVariables(ctx, postApiV4AdminCiVariables)


Create a new instance-level variable

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4AdminCiVariables** | [**PostApiV4AdminCiVariables**](PostApiV4AdminCiVariables.md)|  | 

### Return type

[**ApiEntitiesCiVariable**](API_Entities_Ci_Variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4GroupsIdVariables**
> ApiEntitiesCiVariable PostApiV4GroupsIdVariables(ctx, id, postApiV4GroupsIdVariables)


Create a new variable in a group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group or URL-encoded path of the group owned by the authenticated       user | 
  **postApiV4GroupsIdVariables** | [**PostApiV4GroupsIdVariables**](PostApiV4GroupsIdVariables.md)|  | 

### Return type

[**ApiEntitiesCiVariable**](API_Entities_Ci_Variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdVariables**
> ApiEntitiesCiVariable PostApiV4ProjectsIdVariables(ctx, id, postApiV4ProjectsIdVariables)


Create a new variable in a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a project or URL-encoded NAMESPACE/PROJECT_NAME of the project owned by the authenticated user | 
  **postApiV4ProjectsIdVariables** | [**PostApiV4ProjectsIdVariables**](PostApiV4ProjectsIdVariables.md)|  | 

### Return type

[**ApiEntitiesCiVariable**](API_Entities_Ci_Variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4AdminCiVariablesKey**
> ApiEntitiesCiVariable PutApiV4AdminCiVariablesKey(ctx, key, putApiV4AdminCiVariablesKey)


Update an instance-level variable

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| The key of a variable | 
  **putApiV4AdminCiVariablesKey** | [**PutApiV4AdminCiVariablesKey**](PutApiV4AdminCiVariablesKey.md)|  | 

### Return type

[**ApiEntitiesCiVariable**](API_Entities_Ci_Variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4GroupsIdVariablesKey**
> ApiEntitiesCiVariable PutApiV4GroupsIdVariablesKey(ctx, id, key, putApiV4GroupsIdVariablesKey)


Update an existing variable from a group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group or URL-encoded path of the group owned by the authenticated       user | 
  **key** | **string**| The key of a variable | 
  **putApiV4GroupsIdVariablesKey** | [**PutApiV4GroupsIdVariablesKey**](PutApiV4GroupsIdVariablesKey.md)|  | 

### Return type

[**ApiEntitiesCiVariable**](API_Entities_Ci_Variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdVariablesKey**
> ApiEntitiesCiVariable PutApiV4ProjectsIdVariablesKey(ctx, id, key, putApiV4ProjectsIdVariablesKey)


Update an existing variable from a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a project or URL-encoded NAMESPACE/PROJECT_NAME of the project owned by the authenticated user | 
  **key** | **string**| The key of a variable | 
  **putApiV4ProjectsIdVariablesKey** | [**PutApiV4ProjectsIdVariablesKey**](PutApiV4ProjectsIdVariablesKey.md)|  | 

### Return type

[**ApiEntitiesCiVariable**](API_Entities_Ci_Variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

