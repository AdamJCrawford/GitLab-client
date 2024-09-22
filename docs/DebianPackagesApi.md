# \DebianPackagesApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4GroupsIdPackagesDebianDistsdistributionComponentBinaryArchitectureByHashSha256FileSha256**](DebianPackagesApi.md#GetApiV4GroupsIdPackagesDebianDistsdistributionComponentBinaryArchitectureByHashSha256FileSha256) | **Get** /api/v4/groups/{id}/-/packages/debian/dists/*distribution/{component}/binary-{architecture}/by-hash/SHA256/{file_sha256} | The binary files index by hash
[**GetApiV4GroupsIdPackagesDebianDistsdistributionComponentBinaryArchitecturePackages**](DebianPackagesApi.md#GetApiV4GroupsIdPackagesDebianDistsdistributionComponentBinaryArchitecturePackages) | **Get** /api/v4/groups/{id}/-/packages/debian/dists/*distribution/{component}/binary-{architecture}/Packages | The binary files index
[**GetApiV4GroupsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitectureByHashSha256FileSha256**](DebianPackagesApi.md#GetApiV4GroupsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitectureByHashSha256FileSha256) | **Get** /api/v4/groups/{id}/-/packages/debian/dists/*distribution/{component}/debian-installer/binary-{architecture}/by-hash/SHA256/{file_sha256} | The installer (udeb) binary files index by hash
[**GetApiV4GroupsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitecturePackages**](DebianPackagesApi.md#GetApiV4GroupsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitecturePackages) | **Get** /api/v4/groups/{id}/-/packages/debian/dists/*distribution/{component}/debian-installer/binary-{architecture}/Packages | The installer (udeb) binary files index
[**GetApiV4GroupsIdPackagesDebianDistsdistributionComponentSourceByHashSha256FileSha256**](DebianPackagesApi.md#GetApiV4GroupsIdPackagesDebianDistsdistributionComponentSourceByHashSha256FileSha256) | **Get** /api/v4/groups/{id}/-/packages/debian/dists/*distribution/{component}/source/by-hash/SHA256/{file_sha256} | The source files index by hash
[**GetApiV4GroupsIdPackagesDebianDistsdistributionComponentSourceSources**](DebianPackagesApi.md#GetApiV4GroupsIdPackagesDebianDistsdistributionComponentSourceSources) | **Get** /api/v4/groups/{id}/-/packages/debian/dists/*distribution/{component}/source/Sources | The source files index
[**GetApiV4GroupsIdPackagesDebianDistsdistributionInrelease**](DebianPackagesApi.md#GetApiV4GroupsIdPackagesDebianDistsdistributionInrelease) | **Get** /api/v4/groups/{id}/-/packages/debian/dists/*distribution/InRelease | The signed Release file
[**GetApiV4GroupsIdPackagesDebianDistsdistributionRelease**](DebianPackagesApi.md#GetApiV4GroupsIdPackagesDebianDistsdistributionRelease) | **Get** /api/v4/groups/{id}/-/packages/debian/dists/*distribution/Release | The unsigned Release file
[**GetApiV4GroupsIdPackagesDebianDistsdistributionReleaseGpg**](DebianPackagesApi.md#GetApiV4GroupsIdPackagesDebianDistsdistributionReleaseGpg) | **Get** /api/v4/groups/{id}/-/packages/debian/dists/*distribution/Release.gpg | The Release file signature
[**GetApiV4GroupsIdPackagesDebianPoolDistributionProjectIdLetterPackageNamePackageVersionFileName**](DebianPackagesApi.md#GetApiV4GroupsIdPackagesDebianPoolDistributionProjectIdLetterPackageNamePackageVersionFileName) | **Get** /api/v4/groups/{id}/-/packages/debian/pool/{distribution}/{project_id}/{letter}/{package_name}/{package_version}/{file_name} | Download Debian package
[**GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentBinaryArchitectureByHashSha256FileSha256**](DebianPackagesApi.md#GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentBinaryArchitectureByHashSha256FileSha256) | **Get** /api/v4/projects/{id}/packages/debian/dists/*distribution/{component}/binary-{architecture}/by-hash/SHA256/{file_sha256} | The binary files index by hash
[**GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentBinaryArchitecturePackages**](DebianPackagesApi.md#GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentBinaryArchitecturePackages) | **Get** /api/v4/projects/{id}/packages/debian/dists/*distribution/{component}/binary-{architecture}/Packages | The binary files index
[**GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitectureByHashSha256FileSha256**](DebianPackagesApi.md#GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitectureByHashSha256FileSha256) | **Get** /api/v4/projects/{id}/packages/debian/dists/*distribution/{component}/debian-installer/binary-{architecture}/by-hash/SHA256/{file_sha256} | The installer (udeb) binary files index by hash
[**GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitecturePackages**](DebianPackagesApi.md#GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitecturePackages) | **Get** /api/v4/projects/{id}/packages/debian/dists/*distribution/{component}/debian-installer/binary-{architecture}/Packages | The installer (udeb) binary files index
[**GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentSourceByHashSha256FileSha256**](DebianPackagesApi.md#GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentSourceByHashSha256FileSha256) | **Get** /api/v4/projects/{id}/packages/debian/dists/*distribution/{component}/source/by-hash/SHA256/{file_sha256} | The source files index by hash
[**GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentSourceSources**](DebianPackagesApi.md#GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentSourceSources) | **Get** /api/v4/projects/{id}/packages/debian/dists/*distribution/{component}/source/Sources | The source files index
[**GetApiV4ProjectsIdPackagesDebianDistsdistributionInrelease**](DebianPackagesApi.md#GetApiV4ProjectsIdPackagesDebianDistsdistributionInrelease) | **Get** /api/v4/projects/{id}/packages/debian/dists/*distribution/InRelease | The signed Release file
[**GetApiV4ProjectsIdPackagesDebianDistsdistributionRelease**](DebianPackagesApi.md#GetApiV4ProjectsIdPackagesDebianDistsdistributionRelease) | **Get** /api/v4/projects/{id}/packages/debian/dists/*distribution/Release | The unsigned Release file
[**GetApiV4ProjectsIdPackagesDebianDistsdistributionReleaseGpg**](DebianPackagesApi.md#GetApiV4ProjectsIdPackagesDebianDistsdistributionReleaseGpg) | **Get** /api/v4/projects/{id}/packages/debian/dists/*distribution/Release.gpg | The Release file signature
[**GetApiV4ProjectsIdPackagesDebianPoolDistributionLetterPackageNamePackageVersionFileName**](DebianPackagesApi.md#GetApiV4ProjectsIdPackagesDebianPoolDistributionLetterPackageNamePackageVersionFileName) | **Get** /api/v4/projects/{id}/packages/debian/pool/{distribution}/{letter}/{package_name}/{package_version}/{file_name} | Download Debian package
[**PutApiV4ProjectsIdPackagesDebianFileName**](DebianPackagesApi.md#PutApiV4ProjectsIdPackagesDebianFileName) | **Put** /api/v4/projects/{id}/packages/debian/{file_name} | Upload Debian package
[**PutApiV4ProjectsIdPackagesDebianFileNameAuthorize**](DebianPackagesApi.md#PutApiV4ProjectsIdPackagesDebianFileNameAuthorize) | **Put** /api/v4/projects/{id}/packages/debian/{file_name}/authorize | Authorize Debian package upload


# **GetApiV4GroupsIdPackagesDebianDistsdistributionComponentBinaryArchitectureByHashSha256FileSha256**
> GetApiV4GroupsIdPackagesDebianDistsdistributionComponentBinaryArchitectureByHashSha256FileSha256(ctx, id, distribution, component, architecture, fileSha256)
The binary files index by hash

This feature was introduced in GitLab 15.4

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The group ID or full group path. | 
  **distribution** | **string**| The Debian Codename or Suite | 
  **component** | **string**| The Debian Component | 
  **architecture** | **string**| The Debian Architecture | 
  **fileSha256** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdPackagesDebianDistsdistributionComponentBinaryArchitecturePackages**
> GetApiV4GroupsIdPackagesDebianDistsdistributionComponentBinaryArchitecturePackages(ctx, id, distribution, component, architecture)
The binary files index

This feature was introduced in GitLab 13.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The group ID or full group path. | 
  **distribution** | **string**| The Debian Codename or Suite | 
  **component** | **string**| The Debian Component | 
  **architecture** | **string**| The Debian Architecture | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitectureByHashSha256FileSha256**
> GetApiV4GroupsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitectureByHashSha256FileSha256(ctx, id, distribution, component, architecture, fileSha256)
The installer (udeb) binary files index by hash

This feature was introduced in GitLab 15.4

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The group ID or full group path. | 
  **distribution** | **string**| The Debian Codename or Suite | 
  **component** | **string**| The Debian Component | 
  **architecture** | **string**| The Debian Architecture | 
  **fileSha256** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitecturePackages**
> GetApiV4GroupsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitecturePackages(ctx, id, distribution, component, architecture)
The installer (udeb) binary files index

This feature was introduced in GitLab 15.4

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The group ID or full group path. | 
  **distribution** | **string**| The Debian Codename or Suite | 
  **component** | **string**| The Debian Component | 
  **architecture** | **string**| The Debian Architecture | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdPackagesDebianDistsdistributionComponentSourceByHashSha256FileSha256**
> GetApiV4GroupsIdPackagesDebianDistsdistributionComponentSourceByHashSha256FileSha256(ctx, id, distribution, component, fileSha256)
The source files index by hash

This feature was introduced in GitLab 15.4

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The group ID or full group path. | 
  **distribution** | **string**| The Debian Codename or Suite | 
  **component** | **string**| The Debian Component | 
  **fileSha256** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdPackagesDebianDistsdistributionComponentSourceSources**
> GetApiV4GroupsIdPackagesDebianDistsdistributionComponentSourceSources(ctx, id, distribution, component)
The source files index

This feature was introduced in GitLab 15.4

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The group ID or full group path. | 
  **distribution** | **string**| The Debian Codename or Suite | 
  **component** | **string**| The Debian Component | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdPackagesDebianDistsdistributionInrelease**
> GetApiV4GroupsIdPackagesDebianDistsdistributionInrelease(ctx, id, distribution)
The signed Release file

This feature was introduced in GitLab 13.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The group ID or full group path. | 
  **distribution** | **string**| The Debian Codename or Suite | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdPackagesDebianDistsdistributionRelease**
> GetApiV4GroupsIdPackagesDebianDistsdistributionRelease(ctx, id, distribution)
The unsigned Release file

This feature was introduced in GitLab 13.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The group ID or full group path. | 
  **distribution** | **string**| The Debian Codename or Suite | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdPackagesDebianDistsdistributionReleaseGpg**
> GetApiV4GroupsIdPackagesDebianDistsdistributionReleaseGpg(ctx, id, distribution)
The Release file signature

This feature was introduced in GitLab 13.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The group ID or full group path. | 
  **distribution** | **string**| The Debian Codename or Suite | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdPackagesDebianPoolDistributionProjectIdLetterPackageNamePackageVersionFileName**
> GetApiV4GroupsIdPackagesDebianPoolDistributionProjectIdLetterPackageNamePackageVersionFileName(ctx, id, projectId, distribution, letter, packageName, packageVersion, fileName)
Download Debian package

This feature was introduced in GitLab 14.2

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The group ID or full group path. | 
  **projectId** | **int32**| The Project Id | 
  **distribution** | **string**| The Debian Codename or Suite | 
  **letter** | **string**| The Debian Classification (first-letter or lib-first-letter) | 
  **packageName** | **string**| The Debian Source Package Name | 
  **packageVersion** | **string**| The Debian Source Package Version | 
  **fileName** | **string**| The Debian File Name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentBinaryArchitectureByHashSha256FileSha256**
> GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentBinaryArchitectureByHashSha256FileSha256(ctx, id, distribution, component, architecture, fileSha256)
The binary files index by hash

This feature was introduced in GitLab 15.4

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **distribution** | **string**| The Debian Codename or Suite | 
  **component** | **string**| The Debian Component | 
  **architecture** | **string**| The Debian Architecture | 
  **fileSha256** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentBinaryArchitecturePackages**
> GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentBinaryArchitecturePackages(ctx, id, distribution, component, architecture)
The binary files index

This feature was introduced in GitLab 13.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **distribution** | **string**| The Debian Codename or Suite | 
  **component** | **string**| The Debian Component | 
  **architecture** | **string**| The Debian Architecture | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitectureByHashSha256FileSha256**
> GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitectureByHashSha256FileSha256(ctx, id, distribution, component, architecture, fileSha256)
The installer (udeb) binary files index by hash

This feature was introduced in GitLab 15.4

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **distribution** | **string**| The Debian Codename or Suite | 
  **component** | **string**| The Debian Component | 
  **architecture** | **string**| The Debian Architecture | 
  **fileSha256** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitecturePackages**
> GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitecturePackages(ctx, id, distribution, component, architecture)
The installer (udeb) binary files index

This feature was introduced in GitLab 15.4

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **distribution** | **string**| The Debian Codename or Suite | 
  **component** | **string**| The Debian Component | 
  **architecture** | **string**| The Debian Architecture | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentSourceByHashSha256FileSha256**
> GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentSourceByHashSha256FileSha256(ctx, id, distribution, component, fileSha256)
The source files index by hash

This feature was introduced in GitLab 15.4

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **distribution** | **string**| The Debian Codename or Suite | 
  **component** | **string**| The Debian Component | 
  **fileSha256** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentSourceSources**
> GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentSourceSources(ctx, id, distribution, component)
The source files index

This feature was introduced in GitLab 15.4

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **distribution** | **string**| The Debian Codename or Suite | 
  **component** | **string**| The Debian Component | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesDebianDistsdistributionInrelease**
> GetApiV4ProjectsIdPackagesDebianDistsdistributionInrelease(ctx, id, distribution)
The signed Release file

This feature was introduced in GitLab 13.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **distribution** | **string**| The Debian Codename or Suite | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesDebianDistsdistributionRelease**
> GetApiV4ProjectsIdPackagesDebianDistsdistributionRelease(ctx, id, distribution)
The unsigned Release file

This feature was introduced in GitLab 13.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **distribution** | **string**| The Debian Codename or Suite | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesDebianDistsdistributionReleaseGpg**
> GetApiV4ProjectsIdPackagesDebianDistsdistributionReleaseGpg(ctx, id, distribution)
The Release file signature

This feature was introduced in GitLab 13.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **distribution** | **string**| The Debian Codename or Suite | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesDebianPoolDistributionLetterPackageNamePackageVersionFileName**
> GetApiV4ProjectsIdPackagesDebianPoolDistributionLetterPackageNamePackageVersionFileName(ctx, id, distribution, letter, packageName, packageVersion, fileName)
Download Debian package

This feature was introduced in GitLab 14.2

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **distribution** | **string**| The Debian Codename or Suite | 
  **letter** | **string**| The Debian Classification (first-letter or lib-first-letter) | 
  **packageName** | **string**| The Debian Source Package Name | 
  **packageVersion** | **string**| The Debian Source Package Version | 
  **fileName** | **string**| The Debian File Name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdPackagesDebianFileName**
> PutApiV4ProjectsIdPackagesDebianFileName(ctx, id, fileName, putApiV4ProjectsIdPackagesDebianFileName)
Upload Debian package

This feature was introduced in GitLab 14.0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **fileName** | **string**| The filename | 
  **putApiV4ProjectsIdPackagesDebianFileName** | [**PutApiV4ProjectsIdPackagesDebianFileName**](PutApiV4ProjectsIdPackagesDebianFileName.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdPackagesDebianFileNameAuthorize**
> PutApiV4ProjectsIdPackagesDebianFileNameAuthorize(ctx, id, fileName, putApiV4ProjectsIdPackagesDebianFileNameAuthorize)
Authorize Debian package upload

This feature was introduced in GitLab 13.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **fileName** | **string**| The filename | 
  **putApiV4ProjectsIdPackagesDebianFileNameAuthorize** | [**PutApiV4ProjectsIdPackagesDebianFileNameAuthorize**](PutApiV4ProjectsIdPackagesDebianFileNameAuthorize.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

