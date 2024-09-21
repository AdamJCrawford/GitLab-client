# \RemoteMirrorsApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4ProjectsIdRemoteMirrorsMirrorId**](RemoteMirrorsApi.md#DeleteApiV4ProjectsIdRemoteMirrorsMirrorId) | **Delete** /api/v4/projects/{id}/remote_mirrors/{mirror_id} | Delete a single remote mirror
[**GetApiV4ProjectsIdRemoteMirrors**](RemoteMirrorsApi.md#GetApiV4ProjectsIdRemoteMirrors) | **Get** /api/v4/projects/{id}/remote_mirrors | 
[**GetApiV4ProjectsIdRemoteMirrorsMirrorId**](RemoteMirrorsApi.md#GetApiV4ProjectsIdRemoteMirrorsMirrorId) | **Get** /api/v4/projects/{id}/remote_mirrors/{mirror_id} | 
[**PostApiV4ProjectsIdRemoteMirrors**](RemoteMirrorsApi.md#PostApiV4ProjectsIdRemoteMirrors) | **Post** /api/v4/projects/{id}/remote_mirrors | 
[**PostApiV4ProjectsIdRemoteMirrorsMirrorIdSync**](RemoteMirrorsApi.md#PostApiV4ProjectsIdRemoteMirrorsMirrorIdSync) | **Post** /api/v4/projects/{id}/remote_mirrors/{mirror_id}/sync | 
[**PutApiV4ProjectsIdRemoteMirrorsMirrorId**](RemoteMirrorsApi.md#PutApiV4ProjectsIdRemoteMirrorsMirrorId) | **Put** /api/v4/projects/{id}/remote_mirrors/{mirror_id} | 


# **DeleteApiV4ProjectsIdRemoteMirrorsMirrorId**
> DeleteApiV4ProjectsIdRemoteMirrorsMirrorId(ctx, id, mirrorId)
Delete a single remote mirror

This feature was introduced in GitLab 14.10

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **mirrorId** | **string**| The ID of a remote mirror | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdRemoteMirrors**
> []ApiEntitiesRemoteMirror GetApiV4ProjectsIdRemoteMirrors(ctx, id, optional)


List the project's remote mirrors

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
 **optional** | ***RemoteMirrorsApiGetApiV4ProjectsIdRemoteMirrorsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RemoteMirrorsApiGetApiV4ProjectsIdRemoteMirrorsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesRemoteMirror**](API_Entities_RemoteMirror.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdRemoteMirrorsMirrorId**
> ApiEntitiesRemoteMirror GetApiV4ProjectsIdRemoteMirrorsMirrorId(ctx, id, mirrorId)


Get a single remote mirror

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **mirrorId** | **string**| The ID of a remote mirror | 

### Return type

[**ApiEntitiesRemoteMirror**](API_Entities_RemoteMirror.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdRemoteMirrors**
> ApiEntitiesRemoteMirror PostApiV4ProjectsIdRemoteMirrors(ctx, id, postApiV4ProjectsIdRemoteMirrors)


Create remote mirror for a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **postApiV4ProjectsIdRemoteMirrors** | [**PostApiV4ProjectsIdRemoteMirrors**](PostApiV4ProjectsIdRemoteMirrors.md)|  | 

### Return type

[**ApiEntitiesRemoteMirror**](API_Entities_RemoteMirror.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdRemoteMirrorsMirrorIdSync**
> PostApiV4ProjectsIdRemoteMirrorsMirrorIdSync(ctx, id, mirrorId)


Triggers a push mirror operation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **mirrorId** | **string**| The ID of a remote mirror | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdRemoteMirrorsMirrorId**
> ApiEntitiesRemoteMirror PutApiV4ProjectsIdRemoteMirrorsMirrorId(ctx, id, mirrorId, putApiV4ProjectsIdRemoteMirrorsMirrorId)


Update the attributes of a single remote mirror

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **mirrorId** | **string**| The ID of a remote mirror | 
  **putApiV4ProjectsIdRemoteMirrorsMirrorId** | [**PutApiV4ProjectsIdRemoteMirrorsMirrorId**](PutApiV4ProjectsIdRemoteMirrorsMirrorId.md)|  | 

### Return type

[**ApiEntitiesRemoteMirror**](API_Entities_RemoteMirror.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

