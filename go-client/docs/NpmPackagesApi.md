# \NpmPackagesApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4GroupsIdPackagesNpmPackagepackageNameDistTagsTag**](NpmPackagesApi.md#DeleteApiV4GroupsIdPackagesNpmPackagepackageNameDistTagsTag) | **Delete** /api/v4/groups/{id}/-/packages/npm/-/package/*package_name/dist-tags/{tag} | Deletes the given tag
[**DeleteApiV4PackagesNpmPackagepackageNameDistTagsTag**](NpmPackagesApi.md#DeleteApiV4PackagesNpmPackagepackageNameDistTagsTag) | **Delete** /api/v4/packages/npm/-/package/*package_name/dist-tags/{tag} | Deletes the given tag
[**DeleteApiV4ProjectsIdPackagesNpmPackagepackageNameDistTagsTag**](NpmPackagesApi.md#DeleteApiV4ProjectsIdPackagesNpmPackagepackageNameDistTagsTag) | **Delete** /api/v4/projects/{id}/packages/npm/-/package/*package_name/dist-tags/{tag} | Deletes the given tag
[**GetApiV4GroupsIdPackagesNpmPackagepackageNameDistTags**](NpmPackagesApi.md#GetApiV4GroupsIdPackagesNpmPackagepackageNameDistTags) | **Get** /api/v4/groups/{id}/-/packages/npm/-/package/*package_name/dist-tags | Get all tags for a given an NPM package
[**GetApiV4GroupsIdPackagesNpmpackageName**](NpmPackagesApi.md#GetApiV4GroupsIdPackagesNpmpackageName) | **Get** /api/v4/groups/{id}/-/packages/npm/*package_name | NPM registry metadata endpoint
[**GetApiV4PackagesNpmPackagepackageNameDistTags**](NpmPackagesApi.md#GetApiV4PackagesNpmPackagepackageNameDistTags) | **Get** /api/v4/packages/npm/-/package/*package_name/dist-tags | Get all tags for a given an NPM package
[**GetApiV4PackagesNpmpackageName**](NpmPackagesApi.md#GetApiV4PackagesNpmpackageName) | **Get** /api/v4/packages/npm/*package_name | NPM registry metadata endpoint
[**GetApiV4ProjectsIdPackagesNpmPackagepackageNameDistTags**](NpmPackagesApi.md#GetApiV4ProjectsIdPackagesNpmPackagepackageNameDistTags) | **Get** /api/v4/projects/{id}/packages/npm/-/package/*package_name/dist-tags | Get all tags for a given an NPM package
[**GetApiV4ProjectsIdPackagesNpmpackageName**](NpmPackagesApi.md#GetApiV4ProjectsIdPackagesNpmpackageName) | **Get** /api/v4/projects/{id}/packages/npm/*package_name | NPM registry metadata endpoint
[**GetApiV4ProjectsIdPackagesNpmpackageNamefileName**](NpmPackagesApi.md#GetApiV4ProjectsIdPackagesNpmpackageNamefileName) | **Get** /api/v4/projects/{id}/packages/npm/*package_name/-/*file_name | Download the NPM tarball
[**PostApiV4GroupsIdPackagesNpmNpmV1SecurityAdvisoriesBulk**](NpmPackagesApi.md#PostApiV4GroupsIdPackagesNpmNpmV1SecurityAdvisoriesBulk) | **Post** /api/v4/groups/{id}/-/packages/npm/-/npm/v1/security/advisories/bulk | NPM registry bulk advisory endpoint
[**PostApiV4GroupsIdPackagesNpmNpmV1SecurityAuditsQuick**](NpmPackagesApi.md#PostApiV4GroupsIdPackagesNpmNpmV1SecurityAuditsQuick) | **Post** /api/v4/groups/{id}/-/packages/npm/-/npm/v1/security/audits/quick | NPM registry quick audit endpoint
[**PostApiV4PackagesNpmNpmV1SecurityAdvisoriesBulk**](NpmPackagesApi.md#PostApiV4PackagesNpmNpmV1SecurityAdvisoriesBulk) | **Post** /api/v4/packages/npm/-/npm/v1/security/advisories/bulk | NPM registry bulk advisory endpoint
[**PostApiV4PackagesNpmNpmV1SecurityAuditsQuick**](NpmPackagesApi.md#PostApiV4PackagesNpmNpmV1SecurityAuditsQuick) | **Post** /api/v4/packages/npm/-/npm/v1/security/audits/quick | NPM registry quick audit endpoint
[**PostApiV4ProjectsIdPackagesNpmNpmV1SecurityAdvisoriesBulk**](NpmPackagesApi.md#PostApiV4ProjectsIdPackagesNpmNpmV1SecurityAdvisoriesBulk) | **Post** /api/v4/projects/{id}/packages/npm/-/npm/v1/security/advisories/bulk | NPM registry bulk advisory endpoint
[**PostApiV4ProjectsIdPackagesNpmNpmV1SecurityAuditsQuick**](NpmPackagesApi.md#PostApiV4ProjectsIdPackagesNpmNpmV1SecurityAuditsQuick) | **Post** /api/v4/projects/{id}/packages/npm/-/npm/v1/security/audits/quick | NPM registry quick audit endpoint
[**PutApiV4GroupsIdPackagesNpmPackagepackageNameDistTagsTag**](NpmPackagesApi.md#PutApiV4GroupsIdPackagesNpmPackagepackageNameDistTagsTag) | **Put** /api/v4/groups/{id}/-/packages/npm/-/package/*package_name/dist-tags/{tag} | Create or Update the given tag for the given NPM package and version
[**PutApiV4PackagesNpmPackagepackageNameDistTagsTag**](NpmPackagesApi.md#PutApiV4PackagesNpmPackagepackageNameDistTagsTag) | **Put** /api/v4/packages/npm/-/package/*package_name/dist-tags/{tag} | Create or Update the given tag for the given NPM package and version
[**PutApiV4ProjectsIdPackagesNpmPackageName**](NpmPackagesApi.md#PutApiV4ProjectsIdPackagesNpmPackageName) | **Put** /api/v4/projects/{id}/packages/npm/{package_name} | Create or deprecate NPM package
[**PutApiV4ProjectsIdPackagesNpmPackagepackageNameDistTagsTag**](NpmPackagesApi.md#PutApiV4ProjectsIdPackagesNpmPackagepackageNameDistTagsTag) | **Put** /api/v4/projects/{id}/packages/npm/-/package/*package_name/dist-tags/{tag} | Create or Update the given tag for the given NPM package and version


# **DeleteApiV4GroupsIdPackagesNpmPackagepackageNameDistTagsTag**
> DeleteApiV4GroupsIdPackagesNpmPackagepackageNameDistTagsTag(ctx, id, packageName, tag)
Deletes the given tag

This feature was introduced in GitLab 12.7

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the group | 
  **packageName** | **string**| Package name | 
  **tag** | **string**| Package dist-tag | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4PackagesNpmPackagepackageNameDistTagsTag**
> DeleteApiV4PackagesNpmPackagepackageNameDistTagsTag(ctx, packageName, tag)
Deletes the given tag

This feature was introduced in GitLab 12.7

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **packageName** | **string**| Package name | 
  **tag** | **string**| Package dist-tag | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdPackagesNpmPackagepackageNameDistTagsTag**
> DeleteApiV4ProjectsIdPackagesNpmPackagepackageNameDistTagsTag(ctx, id, packageName, tag)
Deletes the given tag

This feature was introduced in GitLab 12.7

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **packageName** | **string**| Package name | 
  **tag** | **string**| Package dist-tag | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdPackagesNpmPackagepackageNameDistTags**
> ApiEntitiesNpmPackageTag GetApiV4GroupsIdPackagesNpmPackagepackageNameDistTags(ctx, id, packageName)
Get all tags for a given an NPM package

This feature was introduced in GitLab 12.7

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the group | 
  **packageName** | **string**| Package name | 

### Return type

[**ApiEntitiesNpmPackageTag**](API_Entities_NpmPackageTag.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdPackagesNpmpackageName**
> ApiEntitiesNpmPackage GetApiV4GroupsIdPackagesNpmpackageName(ctx, id, packageName)
NPM registry metadata endpoint

This feature was introduced in GitLab 11.8

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the group | 
  **packageName** | **string**| Package name | 

### Return type

[**ApiEntitiesNpmPackage**](API_Entities_NpmPackage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4PackagesNpmPackagepackageNameDistTags**
> ApiEntitiesNpmPackageTag GetApiV4PackagesNpmPackagepackageNameDistTags(ctx, packageName)
Get all tags for a given an NPM package

This feature was introduced in GitLab 12.7

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **packageName** | **string**| Package name | 

### Return type

[**ApiEntitiesNpmPackageTag**](API_Entities_NpmPackageTag.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4PackagesNpmpackageName**
> ApiEntitiesNpmPackage GetApiV4PackagesNpmpackageName(ctx, packageName)
NPM registry metadata endpoint

This feature was introduced in GitLab 11.8

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **packageName** | **string**| Package name | 

### Return type

[**ApiEntitiesNpmPackage**](API_Entities_NpmPackage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesNpmPackagepackageNameDistTags**
> ApiEntitiesNpmPackageTag GetApiV4ProjectsIdPackagesNpmPackagepackageNameDistTags(ctx, id, packageName)
Get all tags for a given an NPM package

This feature was introduced in GitLab 12.7

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **packageName** | **string**| Package name | 

### Return type

[**ApiEntitiesNpmPackageTag**](API_Entities_NpmPackageTag.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesNpmpackageName**
> ApiEntitiesNpmPackage GetApiV4ProjectsIdPackagesNpmpackageName(ctx, id, packageName)
NPM registry metadata endpoint

This feature was introduced in GitLab 11.8

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **packageName** | **string**| Package name | 

### Return type

[**ApiEntitiesNpmPackage**](API_Entities_NpmPackage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesNpmpackageNamefileName**
> GetApiV4ProjectsIdPackagesNpmpackageNamefileName(ctx, id, packageName, fileName)
Download the NPM tarball

This feature was introduced in GitLab 11.8

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **packageName** | **string**| Package name | 
  **fileName** | **string**| Package file name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4GroupsIdPackagesNpmNpmV1SecurityAdvisoriesBulk**
> PostApiV4GroupsIdPackagesNpmNpmV1SecurityAdvisoriesBulk(ctx, id)
NPM registry bulk advisory endpoint

This feature was introduced in GitLab 15.6

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the group | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4GroupsIdPackagesNpmNpmV1SecurityAuditsQuick**
> PostApiV4GroupsIdPackagesNpmNpmV1SecurityAuditsQuick(ctx, id)
NPM registry quick audit endpoint

This feature was introduced in GitLab 15.6

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the group | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4PackagesNpmNpmV1SecurityAdvisoriesBulk**
> PostApiV4PackagesNpmNpmV1SecurityAdvisoriesBulk(ctx, )
NPM registry bulk advisory endpoint

This feature was introduced in GitLab 15.6

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

# **PostApiV4PackagesNpmNpmV1SecurityAuditsQuick**
> PostApiV4PackagesNpmNpmV1SecurityAuditsQuick(ctx, )
NPM registry quick audit endpoint

This feature was introduced in GitLab 15.6

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

# **PostApiV4ProjectsIdPackagesNpmNpmV1SecurityAdvisoriesBulk**
> PostApiV4ProjectsIdPackagesNpmNpmV1SecurityAdvisoriesBulk(ctx, id)
NPM registry bulk advisory endpoint

This feature was introduced in GitLab 15.6

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdPackagesNpmNpmV1SecurityAuditsQuick**
> PostApiV4ProjectsIdPackagesNpmNpmV1SecurityAuditsQuick(ctx, id)
NPM registry quick audit endpoint

This feature was introduced in GitLab 15.6

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4GroupsIdPackagesNpmPackagepackageNameDistTagsTag**
> PutApiV4GroupsIdPackagesNpmPackagepackageNameDistTagsTag(ctx, id, tag, putApiV4GroupsIdPackagesNpmPackagepackageNameDistTagsTag)
Create or Update the given tag for the given NPM package and version

This feature was introduced in GitLab 12.7

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the group | 
  **tag** | **string**| Package dist-tag | 
  **putApiV4GroupsIdPackagesNpmPackagepackageNameDistTagsTag** | [**PutApiV4GroupsIdPackagesNpmPackagepackageNameDistTagsTag**](PutApiV4GroupsIdPackagesNpmPackagepackageNameDistTagsTag.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4PackagesNpmPackagepackageNameDistTagsTag**
> PutApiV4PackagesNpmPackagepackageNameDistTagsTag(ctx, tag, putApiV4PackagesNpmPackagepackageNameDistTagsTag)
Create or Update the given tag for the given NPM package and version

This feature was introduced in GitLab 12.7

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tag** | **string**| Package dist-tag | 
  **putApiV4PackagesNpmPackagepackageNameDistTagsTag** | [**PutApiV4PackagesNpmPackagepackageNameDistTagsTag**](PutApiV4PackagesNpmPackagepackageNameDistTagsTag.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdPackagesNpmPackageName**
> PutApiV4ProjectsIdPackagesNpmPackageName(ctx, id, packageName, putApiV4ProjectsIdPackagesNpmPackageName)
Create or deprecate NPM package

Create was introduced in GitLab 11.8 & deprecate suppport was added in 16.0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **packageName** | **string**| Package name | 
  **putApiV4ProjectsIdPackagesNpmPackageName** | [**PutApiV4ProjectsIdPackagesNpmPackageName**](PutApiV4ProjectsIdPackagesNpmPackageName.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdPackagesNpmPackagepackageNameDistTagsTag**
> PutApiV4ProjectsIdPackagesNpmPackagepackageNameDistTagsTag(ctx, id, tag, putApiV4ProjectsIdPackagesNpmPackagepackageNameDistTagsTag)
Create or Update the given tag for the given NPM package and version

This feature was introduced in GitLab 12.7

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **tag** | **string**| Package dist-tag | 
  **putApiV4ProjectsIdPackagesNpmPackagepackageNameDistTagsTag** | [**PutApiV4ProjectsIdPackagesNpmPackagepackageNameDistTagsTag**](PutApiV4ProjectsIdPackagesNpmPackagepackageNameDistTagsTag.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

