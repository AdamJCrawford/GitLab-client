# \ProjectsJobTokenScopeApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4ProjectsIdJobTokenScopeAllowlistTargetProjectId**](ProjectsJobTokenScopeApi.md#DeleteApiV4ProjectsIdJobTokenScopeAllowlistTargetProjectId) | **Delete** /api/v4/projects/{id}/job_token_scope/allowlist/{target_project_id} | 
[**DeleteApiV4ProjectsIdJobTokenScopeGroupsAllowlistTargetGroupId**](ProjectsJobTokenScopeApi.md#DeleteApiV4ProjectsIdJobTokenScopeGroupsAllowlistTargetGroupId) | **Delete** /api/v4/projects/{id}/job_token_scope/groups_allowlist/{target_group_id} | 
[**GetApiV4ProjectsIdJobTokenScope**](ProjectsJobTokenScopeApi.md#GetApiV4ProjectsIdJobTokenScope) | **Get** /api/v4/projects/{id}/job_token_scope | 
[**GetApiV4ProjectsIdJobTokenScopeAllowlist**](ProjectsJobTokenScopeApi.md#GetApiV4ProjectsIdJobTokenScopeAllowlist) | **Get** /api/v4/projects/{id}/job_token_scope/allowlist | 
[**GetApiV4ProjectsIdJobTokenScopeGroupsAllowlist**](ProjectsJobTokenScopeApi.md#GetApiV4ProjectsIdJobTokenScopeGroupsAllowlist) | **Get** /api/v4/projects/{id}/job_token_scope/groups_allowlist | 
[**PatchApiV4ProjectsIdJobTokenScope**](ProjectsJobTokenScopeApi.md#PatchApiV4ProjectsIdJobTokenScope) | **Patch** /api/v4/projects/{id}/job_token_scope | 
[**PostApiV4ProjectsIdJobTokenScopeAllowlist**](ProjectsJobTokenScopeApi.md#PostApiV4ProjectsIdJobTokenScopeAllowlist) | **Post** /api/v4/projects/{id}/job_token_scope/allowlist | 
[**PostApiV4ProjectsIdJobTokenScopeGroupsAllowlist**](ProjectsJobTokenScopeApi.md#PostApiV4ProjectsIdJobTokenScopeGroupsAllowlist) | **Post** /api/v4/projects/{id}/job_token_scope/groups_allowlist | 


# **DeleteApiV4ProjectsIdJobTokenScopeAllowlistTargetProjectId**
> DeleteApiV4ProjectsIdJobTokenScopeAllowlistTargetProjectId(ctx, id, targetProjectId)


Delete project from allowlist.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| ID of user project | 
  **targetProjectId** | **int32**| ID of the project to be removed from the allowlist | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdJobTokenScopeGroupsAllowlistTargetGroupId**
> DeleteApiV4ProjectsIdJobTokenScopeGroupsAllowlistTargetGroupId(ctx, id, targetGroupId)


Delete target group from allowlist.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| ID of user project | 
  **targetGroupId** | **int32**| ID of the group to be removed from the allowlist | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdJobTokenScope**
> ApiEntitiesProjectJobTokenScope GetApiV4ProjectsIdJobTokenScope(ctx, id)


Fetch CI_JOB_TOKEN access settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**ApiEntitiesProjectJobTokenScope**](API_Entities_ProjectJobTokenScope.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdJobTokenScopeAllowlist**
> ApiEntitiesBasicProjectDetails GetApiV4ProjectsIdJobTokenScopeAllowlist(ctx, id, optional)


Fetch project inbound allowlist for CI_JOB_TOKEN access settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***ProjectsJobTokenScopeApiGetApiV4ProjectsIdJobTokenScopeAllowlistOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsJobTokenScopeApiGetApiV4ProjectsIdJobTokenScopeAllowlistOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**ApiEntitiesBasicProjectDetails**](API_Entities_BasicProjectDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdJobTokenScopeGroupsAllowlist**
> ApiEntitiesBasicProjectDetails GetApiV4ProjectsIdJobTokenScopeGroupsAllowlist(ctx, id, optional)


Fetch project groups allowlist for CI_JOB_TOKEN access settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***ProjectsJobTokenScopeApiGetApiV4ProjectsIdJobTokenScopeGroupsAllowlistOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsJobTokenScopeApiGetApiV4ProjectsIdJobTokenScopeGroupsAllowlistOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**ApiEntitiesBasicProjectDetails**](API_Entities_BasicProjectDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchApiV4ProjectsIdJobTokenScope**
> PatchApiV4ProjectsIdJobTokenScope(ctx, id, patchApiV4ProjectsIdJobTokenScope)


Patch CI_JOB_TOKEN access settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **patchApiV4ProjectsIdJobTokenScope** | [**PatchApiV4ProjectsIdJobTokenScope**](PatchApiV4ProjectsIdJobTokenScope.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdJobTokenScopeAllowlist**
> ApiEntitiesBasicProjectDetails PostApiV4ProjectsIdJobTokenScopeAllowlist(ctx, id, postApiV4ProjectsIdJobTokenScopeAllowlist)


Add target project to allowlist.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| ID of user project | 
  **postApiV4ProjectsIdJobTokenScopeAllowlist** | [**PostApiV4ProjectsIdJobTokenScopeAllowlist**](PostApiV4ProjectsIdJobTokenScopeAllowlist.md)|  | 

### Return type

[**ApiEntitiesBasicProjectDetails**](API_Entities_BasicProjectDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdJobTokenScopeGroupsAllowlist**
> ApiEntitiesBasicGroupDetails PostApiV4ProjectsIdJobTokenScopeGroupsAllowlist(ctx, id, postApiV4ProjectsIdJobTokenScopeGroupsAllowlist)


Add target group to allowlist.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| ID of user project | 
  **postApiV4ProjectsIdJobTokenScopeGroupsAllowlist** | [**PostApiV4ProjectsIdJobTokenScopeGroupsAllowlist**](PostApiV4ProjectsIdJobTokenScopeGroupsAllowlist.md)|  | 

### Return type

[**ApiEntitiesBasicGroupDetails**](API_Entities_BasicGroupDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

