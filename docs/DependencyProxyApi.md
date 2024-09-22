# \DependencyProxyApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4GroupsIdDependencyProxyCache**](DependencyProxyApi.md#DeleteApiV4GroupsIdDependencyProxyCache) | **Delete** /api/v4/groups/{id}/dependency_proxy/cache | Purge the dependency proxy for a group


# **DeleteApiV4GroupsIdDependencyProxyCache**
> DeleteApiV4GroupsIdDependencyProxyCache(ctx, id)
Purge the dependency proxy for a group

Schedules for deletion the cached manifests and blobs for a group.This endpoint requires the Owner role for the group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the group owned by the authenticated user | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

