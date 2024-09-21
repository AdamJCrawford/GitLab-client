# \HelmPackagesApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4ProjectsIdPackagesHelmChannelChartsFileNameTgz**](HelmPackagesApi.md#GetApiV4ProjectsIdPackagesHelmChannelChartsFileNameTgz) | **Get** /api/v4/projects/{id}/packages/helm/{channel}/charts/{file_name}.tgz | Download a chart
[**GetApiV4ProjectsIdPackagesHelmChannelIndexYaml**](HelmPackagesApi.md#GetApiV4ProjectsIdPackagesHelmChannelIndexYaml) | **Get** /api/v4/projects/{id}/packages/helm/{channel}/index.yaml | Download a chart index
[**PostApiV4ProjectsIdPackagesHelmApiChannelCharts**](HelmPackagesApi.md#PostApiV4ProjectsIdPackagesHelmApiChannelCharts) | **Post** /api/v4/projects/{id}/packages/helm/api/{channel}/charts | Upload a chart
[**PostApiV4ProjectsIdPackagesHelmApiChannelChartsAuthorize**](HelmPackagesApi.md#PostApiV4ProjectsIdPackagesHelmApiChannelChartsAuthorize) | **Post** /api/v4/projects/{id}/packages/helm/api/{channel}/charts/authorize | Authorize a chart upload from workhorse


# **GetApiV4ProjectsIdPackagesHelmChannelChartsFileNameTgz**
> GetApiV4ProjectsIdPackagesHelmChannelChartsFileNameTgz(ctx, id, channel, fileName)
Download a chart

This feature was introduced in GitLab 14.0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID or full path of a project | 
  **channel** | **string**| Helm channel | 
  **fileName** | **string**| Helm package file name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesHelmChannelIndexYaml**
> GetApiV4ProjectsIdPackagesHelmChannelIndexYaml(ctx, id, channel)
Download a chart index

This feature was introduced in GitLab 14.0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID or full path of a project | 
  **channel** | **string**| Helm channel | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdPackagesHelmApiChannelCharts**
> PostApiV4ProjectsIdPackagesHelmApiChannelCharts(ctx, id, channel, postApiV4ProjectsIdPackagesHelmApiChannelCharts)
Upload a chart

This feature was introduced in GitLab 14.0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID or full path of a project | 
  **channel** | **string**| Helm channel | 
  **postApiV4ProjectsIdPackagesHelmApiChannelCharts** | [**PostApiV4ProjectsIdPackagesHelmApiChannelCharts**](PostApiV4ProjectsIdPackagesHelmApiChannelCharts.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdPackagesHelmApiChannelChartsAuthorize**
> PostApiV4ProjectsIdPackagesHelmApiChannelChartsAuthorize(ctx, id, channel)
Authorize a chart upload from workhorse

This feature was introduced in GitLab 14.0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID or full path of a project | 
  **channel** | **string**| Helm channel | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

