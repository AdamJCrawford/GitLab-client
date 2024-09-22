# \AvatarApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4Avatar**](AvatarApi.md#GetApiV4Avatar) | **Get** /api/v4/avatar | 


# **GetApiV4Avatar**
> ApiEntitiesAvatar GetApiV4Avatar(ctx, email, optional)


Return avatar url for a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **email** | **string**| Public email address of the user | 
 **optional** | ***AvatarApiGetApiV4AvatarOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AvatarApiGetApiV4AvatarOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **optional.Int32**| Single pixel dimension for Gravatar images | 

### Return type

[**ApiEntitiesAvatar**](API_Entities_Avatar.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

