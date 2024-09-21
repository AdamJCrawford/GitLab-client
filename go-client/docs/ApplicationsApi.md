# \ApplicationsApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4ApplicationsId**](ApplicationsApi.md#DeleteApiV4ApplicationsId) | **Delete** /api/v4/applications/{id} | Delete an application
[**GetApiV4Applications**](ApplicationsApi.md#GetApiV4Applications) | **Get** /api/v4/applications | Get applications
[**PostApiV4Applications**](ApplicationsApi.md#PostApiV4Applications) | **Post** /api/v4/applications | Create a new application
[**PostApiV4ApplicationsIdRenewSecret**](ApplicationsApi.md#PostApiV4ApplicationsIdRenewSecret) | **Post** /api/v4/applications/{id}/renew-secret | Renew an application secret


# **DeleteApiV4ApplicationsId**
> DeleteApiV4ApplicationsId(ctx, id)
Delete an application

Delete a specific application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the application (not the application_id) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4Applications**
> []ApiEntitiesApplication GetApiV4Applications(ctx, )
Get applications

List all registered applications

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ApiEntitiesApplication**](API_Entities_Application.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4Applications**
> ApiEntitiesApplicationWithSecret PostApiV4Applications(ctx, postApiV4Applications)
Create a new application

This feature was introduced in GitLab 10.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4Applications** | [**PostApiV4Applications**](PostApiV4Applications.md)|  | 

### Return type

[**ApiEntitiesApplicationWithSecret**](API_Entities_ApplicationWithSecret.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ApplicationsIdRenewSecret**
> ApiEntitiesApplicationWithSecret PostApiV4ApplicationsIdRenewSecret(ctx, id)
Renew an application secret

Renew the secret of a specific application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the application (not the application_id) | 

### Return type

[**ApiEntitiesApplicationWithSecret**](API_Entities_ApplicationWithSecret.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

