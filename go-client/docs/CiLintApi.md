# \CiLintApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4ProjectsIdCiLint**](CiLintApi.md#GetApiV4ProjectsIdCiLint) | **Get** /api/v4/projects/{id}/ci/lint | Validates a CI YAML configuration with a namespace
[**PostApiV4ProjectsIdCiLint**](CiLintApi.md#PostApiV4ProjectsIdCiLint) | **Post** /api/v4/projects/{id}/ci/lint | Validate a CI YAML configuration with a namespace


# **GetApiV4ProjectsIdCiLint**
> ApiEntitiesCiLintResult GetApiV4ProjectsIdCiLint(ctx, id, optional)
Validates a CI YAML configuration with a namespace

Checks if a project’s .gitlab-ci.yml configuration in a given commit (by default HEAD of the         project’s default branch) is valid

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***CiLintApiGetApiV4ProjectsIdCiLintOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CiLintApiGetApiV4ProjectsIdCiLintOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sha** | **optional.String**| Deprecated: Use content_ref instead | 
 **contentRef** | **optional.String**| The CI/CD configuration content is taken from this commit SHA, branch or tag. Defaults to the HEAD of the project&#39;s default branch | 
 **dryRun** | **optional.Bool**| Run pipeline creation simulation, or only do static check. This is false by default | [default to false]
 **includeJobs** | **optional.Bool**| If the list of jobs that would exist in a static check or pipeline         simulation should be included in the response. This is false by default | 
 **ref** | **optional.String**| Deprecated: Use dry_run_ref instead | 
 **dryRunRef** | **optional.String**| Branch or tag used as context when executing a dry run. Defaults to the default branch of the project. Only used when dry_run is true | 

### Return type

[**ApiEntitiesCiLintResult**](API_Entities_Ci_Lint_Result.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdCiLint**
> ApiEntitiesCiLintResult PostApiV4ProjectsIdCiLint(ctx, id, postApiV4ProjectsIdCiLint)
Validate a CI YAML configuration with a namespace

Checks if CI/CD YAML configuration is valid. This endpoint has namespace specific context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **postApiV4ProjectsIdCiLint** | [**PostApiV4ProjectsIdCiLint**](PostApiV4ProjectsIdCiLint.md)|  | 

### Return type

[**ApiEntitiesCiLintResult**](API_Entities_Ci_Lint_Result.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

