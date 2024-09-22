# \ProjectImportGithubApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostApiV4ImportGithub**](ProjectImportGithubApi.md#PostApiV4ImportGithub) | **Post** /api/v4/import/github | Import a GitHub project
[**PostApiV4ImportGithubCancel**](ProjectImportGithubApi.md#PostApiV4ImportGithubCancel) | **Post** /api/v4/import/github/cancel | Cancel GitHub project import


# **PostApiV4ImportGithub**
> ProjectEntity PostApiV4ImportGithub(ctx, postApiV4ImportGithub)
Import a GitHub project

This feature was introduced in GitLab 11.3.4.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4ImportGithub** | [**PostApiV4ImportGithub**](PostApiV4ImportGithub.md)|  | 

### Return type

[**ProjectEntity**](ProjectEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ImportGithubCancel**
> ProjectImportEntity PostApiV4ImportGithubCancel(ctx, postApiV4ImportGithubCancel)
Cancel GitHub project import

This feature was introduced in GitLab 15.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4ImportGithubCancel** | [**PostApiV4ImportGithubCancel**](PostApiV4ImportGithubCancel.md)|  | 

### Return type

[**ProjectImportEntity**](ProjectImportEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

