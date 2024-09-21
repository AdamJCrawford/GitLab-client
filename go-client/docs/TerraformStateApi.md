# \TerraformStateApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4ProjectsIdTerraformStateName**](TerraformStateApi.md#DeleteApiV4ProjectsIdTerraformStateName) | **Delete** /api/v4/projects/{id}/terraform/state/{name} | Delete a Terraform state of a certain name
[**DeleteApiV4ProjectsIdTerraformStateNameLock**](TerraformStateApi.md#DeleteApiV4ProjectsIdTerraformStateNameLock) | **Delete** /api/v4/projects/{id}/terraform/state/{name}/lock | Unlock a Terraform state of a certain name
[**DeleteApiV4ProjectsIdTerraformStateNameVersionsSerial**](TerraformStateApi.md#DeleteApiV4ProjectsIdTerraformStateNameVersionsSerial) | **Delete** /api/v4/projects/{id}/terraform/state/{name}/versions/{serial} | Delete a Terraform state version
[**GetApiV4ProjectsIdTerraformStateName**](TerraformStateApi.md#GetApiV4ProjectsIdTerraformStateName) | **Get** /api/v4/projects/{id}/terraform/state/{name} | Get a Terraform state by its name
[**GetApiV4ProjectsIdTerraformStateNameVersionsSerial**](TerraformStateApi.md#GetApiV4ProjectsIdTerraformStateNameVersionsSerial) | **Get** /api/v4/projects/{id}/terraform/state/{name}/versions/{serial} | Get a Terraform state version
[**PostApiV4ProjectsIdTerraformStateName**](TerraformStateApi.md#PostApiV4ProjectsIdTerraformStateName) | **Post** /api/v4/projects/{id}/terraform/state/{name} | Add a new Terraform state or update an existing one
[**PostApiV4ProjectsIdTerraformStateNameLock**](TerraformStateApi.md#PostApiV4ProjectsIdTerraformStateNameLock) | **Post** /api/v4/projects/{id}/terraform/state/{name}/lock | Lock a Terraform state of a certain name


# **DeleteApiV4ProjectsIdTerraformStateName**
> DeleteApiV4ProjectsIdTerraformStateName(ctx, id, name)
Delete a Terraform state of a certain name

Delete a Terraform state of a certain name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **name** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdTerraformStateNameLock**
> DeleteApiV4ProjectsIdTerraformStateNameLock(ctx, id, name, optional)
Unlock a Terraform state of a certain name

Unlock a Terraform state of a certain name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **name** | **int32**|  | 
 **optional** | ***TerraformStateApiDeleteApiV4ProjectsIdTerraformStateNameLockOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TerraformStateApiDeleteApiV4ProjectsIdTerraformStateNameLockOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iD** | **optional.String**| Terraform state lock ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdTerraformStateNameVersionsSerial**
> DeleteApiV4ProjectsIdTerraformStateNameVersionsSerial(ctx, id, name, serial)
Delete a Terraform state version

Delete a Terraform state version

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **name** | **int32**|  | 
  **serial** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdTerraformStateName**
> GetApiV4ProjectsIdTerraformStateName(ctx, id, name, optional)
Get a Terraform state by its name

Get a Terraform state by its name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **name** | **string**| The name of a Terraform state | 
 **optional** | ***TerraformStateApiGetApiV4ProjectsIdTerraformStateNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TerraformStateApiGetApiV4ProjectsIdTerraformStateNameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iD** | **optional.String**| Terraform state lock ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdTerraformStateNameVersionsSerial**
> *os.File GetApiV4ProjectsIdTerraformStateNameVersionsSerial(ctx, id, name, serial)
Get a Terraform state version

Get a Terraform state version

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **name** | **string**| The name of a Terraform state | 
  **serial** | **int32**| The version number of the state | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdTerraformStateName**
> PostApiV4ProjectsIdTerraformStateName(ctx, id, name)
Add a new Terraform state or update an existing one

Add a new Terraform state or update an existing one

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **name** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdTerraformStateNameLock**
> PostApiV4ProjectsIdTerraformStateNameLock(ctx, id, name, postApiV4ProjectsIdTerraformStateNameLock)
Lock a Terraform state of a certain name

Lock a Terraform state of a certain name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **name** | **int32**|  | 
  **postApiV4ProjectsIdTerraformStateNameLock** | [**PostApiV4ProjectsIdTerraformStateNameLock**](PostApiV4ProjectsIdTerraformStateNameLock.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

