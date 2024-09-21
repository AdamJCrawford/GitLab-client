# \ErrorTrackingClientKeysApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4ProjectsIdErrorTrackingClientKeysKeyId**](ErrorTrackingClientKeysApi.md#DeleteApiV4ProjectsIdErrorTrackingClientKeysKeyId) | **Delete** /api/v4/projects/{id}/error_tracking/client_keys/{key_id} | Delete a client key
[**GetApiV4ProjectsIdErrorTrackingClientKeys**](ErrorTrackingClientKeysApi.md#GetApiV4ProjectsIdErrorTrackingClientKeys) | **Get** /api/v4/projects/{id}/error_tracking/client_keys | List project client keys
[**PostApiV4ProjectsIdErrorTrackingClientKeys**](ErrorTrackingClientKeysApi.md#PostApiV4ProjectsIdErrorTrackingClientKeys) | **Post** /api/v4/projects/{id}/error_tracking/client_keys | Create a client key


# **DeleteApiV4ProjectsIdErrorTrackingClientKeysKeyId**
> ApiEntitiesErrorTrackingClientKey DeleteApiV4ProjectsIdErrorTrackingClientKeysKeyId(ctx, id, keyId)
Delete a client key

Removes a client key from the project. This feature was introduced in GitLab 14.3.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **keyId** | **int32**|  | 

### Return type

[**ApiEntitiesErrorTrackingClientKey**](API_Entities_ErrorTracking_ClientKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdErrorTrackingClientKeys**
> []ApiEntitiesErrorTrackingClientKey GetApiV4ProjectsIdErrorTrackingClientKeys(ctx, id)
List project client keys

List all client keys. This feature was introduced in GitLab 14.3.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 

### Return type

[**[]ApiEntitiesErrorTrackingClientKey**](API_Entities_ErrorTracking_ClientKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdErrorTrackingClientKeys**
> ApiEntitiesErrorTrackingClientKey PostApiV4ProjectsIdErrorTrackingClientKeys(ctx, id)
Create a client key

Creates a new client key for a project. The public key attribute is generated automatically.This feature was introduced in GitLab 14.3.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 

### Return type

[**ApiEntitiesErrorTrackingClientKey**](API_Entities_ErrorTracking_ClientKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

