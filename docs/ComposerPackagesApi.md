# \ComposerPackagesApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4GroupIdPackagesComposerP2packageName**](ComposerPackagesApi.md#GetApiV4GroupIdPackagesComposerP2packageName) | **Get** /api/v4/group/{id}/-/packages/composer/p2/*package_name | Composer v2 packages p2 endpoint at group level for package versions metadata
[**GetApiV4GroupIdPackagesComposerPSha**](ComposerPackagesApi.md#GetApiV4GroupIdPackagesComposerPSha) | **Get** /api/v4/group/{id}/-/packages/composer/p/{sha} | Composer packages endpoint at group level for packages list
[**GetApiV4GroupIdPackagesComposerPackages**](ComposerPackagesApi.md#GetApiV4GroupIdPackagesComposerPackages) | **Get** /api/v4/group/{id}/-/packages/composer/packages | Composer packages endpoint at group level
[**GetApiV4GroupIdPackagesComposerpackageName**](ComposerPackagesApi.md#GetApiV4GroupIdPackagesComposerpackageName) | **Get** /api/v4/group/{id}/-/packages/composer/*package_name | Composer packages endpoint at group level for package versions metadata
[**GetApiV4ProjectsIdPackagesComposerArchivespackageName**](ComposerPackagesApi.md#GetApiV4ProjectsIdPackagesComposerArchivespackageName) | **Get** /api/v4/projects/{id}/packages/composer/archives/*package_name | Composer package endpoint to download a package archive
[**PostApiV4ProjectsIdPackagesComposer**](ComposerPackagesApi.md#PostApiV4ProjectsIdPackagesComposer) | **Post** /api/v4/projects/{id}/packages/composer | Composer packages endpoint for registering packages


# **GetApiV4GroupIdPackagesComposerP2packageName**
> GetApiV4GroupIdPackagesComposerP2packageName(ctx, id, packageName)
Composer v2 packages p2 endpoint at group level for package versions metadata

This feature was introduced in GitLab 13.1

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of a group | 
  **packageName** | **string**| The Composer package name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupIdPackagesComposerPSha**
> GetApiV4GroupIdPackagesComposerPSha(ctx, id, sha)
Composer packages endpoint at group level for packages list

This feature was introduced in GitLab 13.1

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of a group | 
  **sha** | **string**| Shasum of current json | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupIdPackagesComposerPackages**
> GetApiV4GroupIdPackagesComposerPackages(ctx, id)
Composer packages endpoint at group level

This feature was introduced in GitLab 13.1

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of a group | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupIdPackagesComposerpackageName**
> GetApiV4GroupIdPackagesComposerpackageName(ctx, id, packageName)
Composer packages endpoint at group level for package versions metadata

This feature was introduced in GitLab 12.1

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of a group | 
  **packageName** | **string**| The Composer package name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesComposerArchivespackageName**
> GetApiV4ProjectsIdPackagesComposerArchivespackageName(ctx, id, sha, packageName)
Composer package endpoint to download a package archive

This feature was introduced in GitLab 13.1

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of a project | 
  **sha** | **string**| Shasum of current json | 
  **packageName** | **string**| The Composer package name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdPackagesComposer**
> PostApiV4ProjectsIdPackagesComposer(ctx, id, postApiV4ProjectsIdPackagesComposer)
Composer packages endpoint for registering packages

This feature was introduced in GitLab 13.1

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of a project | 
  **postApiV4ProjectsIdPackagesComposer** | [**PostApiV4ProjectsIdPackagesComposer**](PostApiV4ProjectsIdPackagesComposer.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

