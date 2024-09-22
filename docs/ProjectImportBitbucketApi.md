# \ProjectImportBitbucketApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostApiV4ImportBitbucket**](ProjectImportBitbucketApi.md#PostApiV4ImportBitbucket) | **Post** /api/v4/import/bitbucket | Import a BitBucket Cloud repository
[**PostApiV4ImportBitbucketServer**](ProjectImportBitbucketApi.md#PostApiV4ImportBitbucketServer) | **Post** /api/v4/import/bitbucket_server | Import a BitBucket Server repository


# **PostApiV4ImportBitbucket**
> ProjectImportEntity PostApiV4ImportBitbucket(ctx, postApiV4ImportBitbucket)
Import a BitBucket Cloud repository

This feature was introduced in GitLab 17.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4ImportBitbucket** | [**PostApiV4ImportBitbucket**](PostApiV4ImportBitbucket.md)|  | 

### Return type

[**ProjectImportEntity**](ProjectImportEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ImportBitbucketServer**
> ProjectEntity PostApiV4ImportBitbucketServer(ctx, postApiV4ImportBitbucketServer)
Import a BitBucket Server repository

This feature was introduced in GitLab 13.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4ImportBitbucketServer** | [**PostApiV4ImportBitbucketServer**](PostApiV4ImportBitbucketServer.md)|  | 

### Return type

[**ProjectEntity**](ProjectEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

