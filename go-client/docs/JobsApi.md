# \JobsApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4JobsIdArtifacts**](JobsApi.md#GetApiV4JobsIdArtifacts) | **Get** /api/v4/jobs/{id}/artifacts | 
[**GetApiV4RunnersIdJobs**](JobsApi.md#GetApiV4RunnersIdJobs) | **Get** /api/v4/runners/{id}/jobs | List runner&#39;s jobs
[**PatchApiV4JobsIdTrace**](JobsApi.md#PatchApiV4JobsIdTrace) | **Patch** /api/v4/jobs/{id}/trace | 
[**PostApiV4JobsIdArtifacts**](JobsApi.md#PostApiV4JobsIdArtifacts) | **Post** /api/v4/jobs/{id}/artifacts | 
[**PostApiV4JobsIdArtifactsAuthorize**](JobsApi.md#PostApiV4JobsIdArtifactsAuthorize) | **Post** /api/v4/jobs/{id}/artifacts/authorize | 
[**PostApiV4JobsRequest**](JobsApi.md#PostApiV4JobsRequest) | **Post** /api/v4/jobs/request | 
[**PutApiV4JobsId**](JobsApi.md#PutApiV4JobsId) | **Put** /api/v4/jobs/{id} | 


# **GetApiV4JobsIdArtifacts**
> GetApiV4JobsIdArtifacts(ctx, id, optional)


Download the artifacts file for job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Job&#39;s ID | 
 **optional** | ***JobsApiGetApiV4JobsIdArtifactsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobsApiGetApiV4JobsIdArtifactsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **optional.String**| Job&#39;s authentication token | 
 **directDownload** | **optional.Bool**| Perform direct download from remote storage instead of proxying artifacts | [default to false]

### Return type

 (empty response body)

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
 **optional** | ***JobsApiGetApiV4RunnersIdJobsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobsApiGetApiV4RunnersIdJobsOpts struct

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

# **PatchApiV4JobsIdTrace**
> PatchApiV4JobsIdTrace(ctx, id, patchApiV4JobsIdTrace)


Append a patch to the job trace

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Job&#39;s ID | 
  **patchApiV4JobsIdTrace** | [**PatchApiV4JobsIdTrace**](PatchApiV4JobsIdTrace.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4JobsIdArtifacts**
> PostApiV4JobsIdArtifacts(ctx, id, postApiV4JobsIdArtifacts)


Upload a job artifact

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Job&#39;s ID | 
  **postApiV4JobsIdArtifacts** | [**PostApiV4JobsIdArtifacts**](PostApiV4JobsIdArtifacts.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4JobsIdArtifactsAuthorize**
> PostApiV4JobsIdArtifactsAuthorize(ctx, id, postApiV4JobsIdArtifactsAuthorize)


Authorize uploading job artifact

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Job&#39;s ID | 
  **postApiV4JobsIdArtifactsAuthorize** | [**PostApiV4JobsIdArtifactsAuthorize**](PostApiV4JobsIdArtifactsAuthorize.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4JobsRequest**
> PostApiV4JobsRequest(ctx, postApiV4JobsRequest)


Request a job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4JobsRequest** | [**PostApiV4JobsRequest**](PostApiV4JobsRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4JobsId**
> PutApiV4JobsId(ctx, id, putApiV4JobsId)


Update a job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Job&#39;s ID | 
  **putApiV4JobsId** | [**PutApiV4JobsId**](PutApiV4JobsId.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

