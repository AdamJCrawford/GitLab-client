# \RunnersApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4ProjectsIdRunnersRunnerId**](RunnersApi.md#DeleteApiV4ProjectsIdRunnersRunnerId) | **Delete** /api/v4/projects/{id}/runners/{runner_id} | Disable a project runner from the project
[**DeleteApiV4Runners**](RunnersApi.md#DeleteApiV4Runners) | **Delete** /api/v4/runners | Delete a runner by authentication token
[**DeleteApiV4RunnersId**](RunnersApi.md#DeleteApiV4RunnersId) | **Delete** /api/v4/runners/{id} | Delete a runner
[**DeleteApiV4RunnersManagers**](RunnersApi.md#DeleteApiV4RunnersManagers) | **Delete** /api/v4/runners/managers | Internal endpoint that deletes a runner manager by authentication token and system ID.
[**GetApiV4GroupsIdRunners**](RunnersApi.md#GetApiV4GroupsIdRunners) | **Get** /api/v4/groups/{id}/runners | List group&#39;s runners
[**GetApiV4ProjectsIdRunners**](RunnersApi.md#GetApiV4ProjectsIdRunners) | **Get** /api/v4/projects/{id}/runners | List project&#39;s runners
[**GetApiV4Runners**](RunnersApi.md#GetApiV4Runners) | **Get** /api/v4/runners | List owned runners
[**GetApiV4RunnersAll**](RunnersApi.md#GetApiV4RunnersAll) | **Get** /api/v4/runners/all | List all runners
[**GetApiV4RunnersId**](RunnersApi.md#GetApiV4RunnersId) | **Get** /api/v4/runners/{id} | Get runner&#39;s details
[**GetApiV4RunnersIdJobs**](RunnersApi.md#GetApiV4RunnersIdJobs) | **Get** /api/v4/runners/{id}/jobs | List runner&#39;s jobs
[**PostApiV4GroupsIdRunnersResetRegistrationToken**](RunnersApi.md#PostApiV4GroupsIdRunnersResetRegistrationToken) | **Post** /api/v4/groups/{id}/runners/reset_registration_token | Reset the runner registration token for a group
[**PostApiV4ProjectsIdRunners**](RunnersApi.md#PostApiV4ProjectsIdRunners) | **Post** /api/v4/projects/{id}/runners | Enable a runner in project
[**PostApiV4ProjectsIdRunnersResetRegistrationToken**](RunnersApi.md#PostApiV4ProjectsIdRunnersResetRegistrationToken) | **Post** /api/v4/projects/{id}/runners/reset_registration_token | Reset the runner registration token for a project
[**PostApiV4Runners**](RunnersApi.md#PostApiV4Runners) | **Post** /api/v4/runners | Register a new runner
[**PostApiV4RunnersIdResetAuthenticationToken**](RunnersApi.md#PostApiV4RunnersIdResetAuthenticationToken) | **Post** /api/v4/runners/{id}/reset_authentication_token | Reset runner&#39;s authentication token
[**PostApiV4RunnersResetAuthenticationToken**](RunnersApi.md#PostApiV4RunnersResetAuthenticationToken) | **Post** /api/v4/runners/reset_authentication_token | 
[**PostApiV4RunnersResetRegistrationToken**](RunnersApi.md#PostApiV4RunnersResetRegistrationToken) | **Post** /api/v4/runners/reset_registration_token | Reset instance&#39;s runner registration token
[**PostApiV4RunnersVerify**](RunnersApi.md#PostApiV4RunnersVerify) | **Post** /api/v4/runners/verify | Verify authentication for a registered runner
[**PostApiV4UserRunners**](RunnersApi.md#PostApiV4UserRunners) | **Post** /api/v4/user/runners | Create a runner owned by currently authenticated user
[**PutApiV4RunnersId**](RunnersApi.md#PutApiV4RunnersId) | **Put** /api/v4/runners/{id} | Update details of a runner


# **DeleteApiV4ProjectsIdRunnersRunnerId**
> ApiEntitiesCiRunner DeleteApiV4ProjectsIdRunnersRunnerId(ctx, id, runnerId)
Disable a project runner from the project

It works only if the project isn't the only project associated with the specified runner. If so, an error is returned. Use the call to delete a runner instead.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **runnerId** | **int32**| The ID of a runner | 

### Return type

[**ApiEntitiesCiRunner**](API_Entities_Ci_Runner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4Runners**
> DeleteApiV4Runners(ctx, token)
Delete a runner by authentication token

Delete a registered runner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| The runner&#39;s authentication token | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4RunnersId**
> ApiEntitiesCiRunner DeleteApiV4RunnersId(ctx, id)
Delete a runner

Remove a runner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of a runner | 

### Return type

[**ApiEntitiesCiRunner**](API_Entities_Ci_Runner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4RunnersManagers**
> DeleteApiV4RunnersManagers(ctx, token, systemId)
Internal endpoint that deletes a runner manager by authentication token and system ID.

Delete a registered runner manager

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| The runner&#39;s authentication token | 
  **systemId** | **string**| The runner&#39;s system identifier. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdRunners**
> ApiEntitiesCiRunner GetApiV4GroupsIdRunners(ctx, id, optional)
List group's runners

List all runners available in the group as well as its ancestor groups, including any allowed shared runners.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group | 
 **optional** | ***RunnersApiGetApiV4GroupsIdRunnersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RunnersApiGetApiV4GroupsIdRunnersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**| The type of runners to return | 
 **paused** | **optional.Bool**| Whether to include only runners that are accepting or ignoring new jobs | 
 **status** | **optional.String**| The status of runners to return | 
 **tagList** | [**optional.Interface of []string**](string.md)| A list of runner tags | 
 **versionPrefix** | **optional.String**| The version prefix of runners to return | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**ApiEntitiesCiRunner**](API_Entities_Ci_Runner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdRunners**
> ApiEntitiesCiRunner GetApiV4ProjectsIdRunners(ctx, id, optional)
List project's runners

List all runners available in the project, including from ancestor groups and any allowed shared runners.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
 **optional** | ***RunnersApiGetApiV4ProjectsIdRunnersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RunnersApiGetApiV4ProjectsIdRunnersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scope** | **optional.String**| Deprecated: Use &#x60;type&#x60; or &#x60;status&#x60; instead. The scope of runners to return | 
 **type_** | **optional.String**| The type of runners to return | 
 **paused** | **optional.Bool**| Whether to include only runners that are accepting or ignoring new jobs | 
 **status** | **optional.String**| The status of runners to return | 
 **tagList** | [**optional.Interface of []string**](string.md)| A list of runner tags | 
 **versionPrefix** | **optional.String**| The version prefix of runners to return | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**ApiEntitiesCiRunner**](API_Entities_Ci_Runner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4Runners**
> ApiEntitiesCiRunner GetApiV4Runners(ctx, optional)
List owned runners

Get runners available for user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RunnersApiGetApiV4RunnersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RunnersApiGetApiV4RunnersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **optional.String**| Deprecated: Use &#x60;type&#x60; or &#x60;status&#x60; instead. The scope of runners to return | 
 **type_** | **optional.String**| The type of runners to return | 
 **paused** | **optional.Bool**| Whether to include only runners that are accepting or ignoring new jobs | 
 **status** | **optional.String**| The status of runners to return | 
 **tagList** | [**optional.Interface of []string**](string.md)| A list of runner tags | 
 **versionPrefix** | **optional.String**| The version prefix of runners to return | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**ApiEntitiesCiRunner**](API_Entities_Ci_Runner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4RunnersAll**
> ApiEntitiesCiRunner GetApiV4RunnersAll(ctx, optional)
List all runners

Get a list of all runners in the GitLab instance (shared and project). Access is restricted to users with administrator access.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RunnersApiGetApiV4RunnersAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RunnersApiGetApiV4RunnersAllOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **optional.String**| Deprecated: Use &#x60;type&#x60; or &#x60;status&#x60; instead. The scope of runners to return | 
 **type_** | **optional.String**| The type of runners to return | 
 **paused** | **optional.Bool**| Whether to include only runners that are accepting or ignoring new jobs | 
 **status** | **optional.String**| The status of runners to return | 
 **tagList** | [**optional.Interface of []string**](string.md)| A list of runner tags | 
 **versionPrefix** | **optional.String**| The version prefix of runners to return | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**ApiEntitiesCiRunner**](API_Entities_Ci_Runner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4RunnersId**
> ApiEntitiesCiRunnerDetails GetApiV4RunnersId(ctx, id)
Get runner's details

At least the Maintainer role is required to get runner details at the project and group level. Instance-level runner details via this endpoint are available to all signed in users.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of a runner | 

### Return type

[**ApiEntitiesCiRunnerDetails**](API_Entities_Ci_RunnerDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4RunnersIdJobs**
> ApiEntitiesCiJobBasicWithProject GetApiV4RunnersIdJobs(ctx, id, optional)
List runner's jobs

List jobs that are being processed or were processed by the specified runner. The list of jobs is limited to projects where the user has at least the Reporter role.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of a runner | 
 **optional** | ***RunnersApiGetApiV4RunnersIdJobsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RunnersApiGetApiV4RunnersIdJobsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemId** | **optional.String**| System ID associated with the runner manager | 
 **status** | **optional.String**| Status of the job | 
 **orderBy** | **optional.String**| Order by &#x60;id&#x60; | 
 **sort** | **optional.String**| Sort by &#x60;asc&#x60; or &#x60;desc&#x60; order. Specify &#x60;order_by&#x60; as well, including for &#x60;id&#x60; | [default to desc]
 **cursor** | **optional.String**| Cursor for obtaining the next set of records | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**ApiEntitiesCiJobBasicWithProject**](API_Entities_Ci_JobBasicWithProject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4GroupsIdRunnersResetRegistrationToken**
> ApiEntitiesCiResetTokenResult PostApiV4GroupsIdRunnersResetRegistrationToken(ctx, id)
Reset the runner registration token for a group

Reset runner registration token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group | 

### Return type

[**ApiEntitiesCiResetTokenResult**](API_Entities_Ci_ResetTokenResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdRunners**
> ApiEntitiesCiRunner PostApiV4ProjectsIdRunners(ctx, id, postApiV4ProjectsIdRunners)
Enable a runner in project

Enable an available project runner in the project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **postApiV4ProjectsIdRunners** | [**PostApiV4ProjectsIdRunners**](PostApiV4ProjectsIdRunners.md)|  | 

### Return type

[**ApiEntitiesCiRunner**](API_Entities_Ci_Runner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdRunnersResetRegistrationToken**
> ApiEntitiesCiResetTokenResult PostApiV4ProjectsIdRunnersResetRegistrationToken(ctx, id)
Reset the runner registration token for a project

Reset runner registration token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a project | 

### Return type

[**ApiEntitiesCiResetTokenResult**](API_Entities_Ci_ResetTokenResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4Runners**
> ApiEntitiesCiRunnerRegistrationDetails PostApiV4Runners(ctx, postApiV4Runners)
Register a new runner

Register a new runner for the instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4Runners** | [**PostApiV4Runners**](PostApiV4Runners.md)|  | 

### Return type

[**ApiEntitiesCiRunnerRegistrationDetails**](API_Entities_Ci_RunnerRegistrationDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4RunnersIdResetAuthenticationToken**
> ApiEntitiesCiResetTokenResult PostApiV4RunnersIdResetAuthenticationToken(ctx, id)
Reset runner's authentication token

Reset runner authentication token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the runner | 

### Return type

[**ApiEntitiesCiResetTokenResult**](API_Entities_Ci_ResetTokenResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4RunnersResetAuthenticationToken**
> ApiEntitiesCiResetTokenResult PostApiV4RunnersResetAuthenticationToken(ctx, postApiV4RunnersResetAuthenticationToken)


Reset runner authentication token with current token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4RunnersResetAuthenticationToken** | [**PostApiV4RunnersResetAuthenticationToken**](PostApiV4RunnersResetAuthenticationToken.md)|  | 

### Return type

[**ApiEntitiesCiResetTokenResult**](API_Entities_Ci_ResetTokenResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4RunnersResetRegistrationToken**
> ApiEntitiesCiResetTokenResult PostApiV4RunnersResetRegistrationToken(ctx, )
Reset instance's runner registration token

Reset runner registration token

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiEntitiesCiResetTokenResult**](API_Entities_Ci_ResetTokenResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4RunnersVerify**
> PostApiV4RunnersVerify(ctx, postApiV4RunnersVerify)
Verify authentication for a registered runner

Validate authentication credentials

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4RunnersVerify** | [**PostApiV4RunnersVerify**](PostApiV4RunnersVerify.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4UserRunners**
> ApiEntitiesCiRunnerRegistrationDetails PostApiV4UserRunners(ctx, postApiV4UserRunners)
Create a runner owned by currently authenticated user

Create a new runner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4UserRunners** | [**PostApiV4UserRunners**](PostApiV4UserRunners.md)|  | 

### Return type

[**ApiEntitiesCiRunnerRegistrationDetails**](API_Entities_Ci_RunnerRegistrationDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4RunnersId**
> ApiEntitiesCiRunnerDetails PutApiV4RunnersId(ctx, id, putApiV4RunnersId)
Update details of a runner

Update runner's details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of a runner | 
  **putApiV4RunnersId** | [**PutApiV4RunnersId**](PutApiV4RunnersId.md)|  | 

### Return type

[**ApiEntitiesCiRunnerDetails**](API_Entities_Ci_RunnerDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

