# \RubygemPackagesApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4ProjectsIdPackagesRubygemsApiV1Dependencies**](RubygemPackagesApi.md#GetApiV4ProjectsIdPackagesRubygemsApiV1Dependencies) | **Get** /api/v4/projects/{id}/packages/rubygems/api/v1/dependencies | Fetch a list of dependencies
[**GetApiV4ProjectsIdPackagesRubygemsFileName**](RubygemPackagesApi.md#GetApiV4ProjectsIdPackagesRubygemsFileName) | **Get** /api/v4/projects/{id}/packages/rubygems/{file_name} | Download the spec index file
[**GetApiV4ProjectsIdPackagesRubygemsGemsFileName**](RubygemPackagesApi.md#GetApiV4ProjectsIdPackagesRubygemsGemsFileName) | **Get** /api/v4/projects/{id}/packages/rubygems/gems/{file_name} | Download the .gem package
[**GetApiV4ProjectsIdPackagesRubygemsQuickMarshal48FileName**](RubygemPackagesApi.md#GetApiV4ProjectsIdPackagesRubygemsQuickMarshal48FileName) | **Get** /api/v4/projects/{id}/packages/rubygems/quick/Marshal.4.8/{file_name} | Download the gemspec file
[**PostApiV4ProjectsIdPackagesRubygemsApiV1Gems**](RubygemPackagesApi.md#PostApiV4ProjectsIdPackagesRubygemsApiV1Gems) | **Post** /api/v4/projects/{id}/packages/rubygems/api/v1/gems | Upload a gem
[**PostApiV4ProjectsIdPackagesRubygemsApiV1GemsAuthorize**](RubygemPackagesApi.md#PostApiV4ProjectsIdPackagesRubygemsApiV1GemsAuthorize) | **Post** /api/v4/projects/{id}/packages/rubygems/api/v1/gems/authorize | Authorize a gem upload from workhorse


# **GetApiV4ProjectsIdPackagesRubygemsApiV1Dependencies**
> GetApiV4ProjectsIdPackagesRubygemsApiV1Dependencies(ctx, id, optional)
Fetch a list of dependencies

This feature was introduced in GitLab 13.9

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID or URL-encoded path of the project | 
 **optional** | ***RubygemPackagesApiGetApiV4ProjectsIdPackagesRubygemsApiV1DependenciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RubygemPackagesApiGetApiV4ProjectsIdPackagesRubygemsApiV1DependenciesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gems** | [**optional.Interface of []string**](string.md)| Comma delimited gem names | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesRubygemsFileName**
> GetApiV4ProjectsIdPackagesRubygemsFileName(ctx, id, fileName)
Download the spec index file

This feature was introduced in GitLab 13.9

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID or URL-encoded path of the project | 
  **fileName** | ***os.File**| Spec file name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesRubygemsGemsFileName**
> GetApiV4ProjectsIdPackagesRubygemsGemsFileName(ctx, id, fileName)
Download the .gem package

This feature was introduced in GitLab 13.9

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID or URL-encoded path of the project | 
  **fileName** | ***os.File**| Package file name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesRubygemsQuickMarshal48FileName**
> GetApiV4ProjectsIdPackagesRubygemsQuickMarshal48FileName(ctx, id, fileName)
Download the gemspec file

This feature was introduced in GitLab 13.9

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID or URL-encoded path of the project | 
  **fileName** | ***os.File**| Gemspec file name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdPackagesRubygemsApiV1Gems**
> PostApiV4ProjectsIdPackagesRubygemsApiV1Gems(ctx, id, postApiV4ProjectsIdPackagesRubygemsApiV1Gems)
Upload a gem

This feature was introduced in GitLab 13.9

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID or URL-encoded path of the project | 
  **postApiV4ProjectsIdPackagesRubygemsApiV1Gems** | [**PostApiV4ProjectsIdPackagesRubygemsApiV1Gems**](PostApiV4ProjectsIdPackagesRubygemsApiV1Gems.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdPackagesRubygemsApiV1GemsAuthorize**
> PostApiV4ProjectsIdPackagesRubygemsApiV1GemsAuthorize(ctx, id)
Authorize a gem upload from workhorse

This feature was introduced in GitLab 13.9

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID or URL-encoded path of the project | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

