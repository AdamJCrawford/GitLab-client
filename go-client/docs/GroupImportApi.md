# \GroupImportApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostApiV4GroupsImport**](GroupImportApi.md#PostApiV4GroupsImport) | **Post** /api/v4/groups/import | Create a new group import
[**PostApiV4GroupsImportAuthorize**](GroupImportApi.md#PostApiV4GroupsImportAuthorize) | **Post** /api/v4/groups/import/authorize | Workhorse authorize the group import upload


# **PostApiV4GroupsImport**
> PostApiV4GroupsImport(ctx, path, name, file, optional)
Create a new group import

This feature was introduced in GitLab 12.8

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| Group path | 
  **name** | **string**| Group name | 
  **file** | ***os.File**| The group export file to be imported | 
 **optional** | ***GroupImportApiPostApiV4GroupsImportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupImportApiPostApiV4GroupsImportOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **parentId** | **optional.Int32**| The ID of the parent group that the group will be imported into. Defaults to the current user&#39;s namespace. | 
 **organizationId** | **optional.Int32**| The ID of the organization that the group will be part of.  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4GroupsImportAuthorize**
> PostApiV4GroupsImportAuthorize(ctx, )
Workhorse authorize the group import upload

This feature was introduced in GitLab 12.8

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

