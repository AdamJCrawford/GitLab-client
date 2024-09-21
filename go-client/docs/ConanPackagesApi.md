# \ConanPackagesApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel**](ConanPackagesApi.md#DeleteApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel) | **Delete** /api/v4/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel} | Delete Package
[**DeleteApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel**](ConanPackagesApi.md#DeleteApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel) | **Delete** /api/v4/projects/{id}/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel} | Delete Package
[**GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel**](ConanPackagesApi.md#GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel) | **Get** /api/v4/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel} | Recipe Snapshot
[**GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigest**](ConanPackagesApi.md#GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigest) | **Get** /api/v4/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel}/digest | Recipe Digest
[**GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrls**](ConanPackagesApi.md#GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrls) | **Get** /api/v4/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel}/download_urls | Recipe Download Urls
[**GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReference**](ConanPackagesApi.md#GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReference) | **Get** /api/v4/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel}/packages/{conan_package_reference} | Package Snapshot
[**GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigest**](ConanPackagesApi.md#GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigest) | **Get** /api/v4/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel}/packages/{conan_package_reference}/digest | Package Digest
[**GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrls**](ConanPackagesApi.md#GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrls) | **Get** /api/v4/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel}/packages/{conan_package_reference}/download_urls | Package Download Urls
[**GetApiV4PackagesConanV1ConansSearch**](ConanPackagesApi.md#GetApiV4PackagesConanV1ConansSearch) | **Get** /api/v4/packages/conan/v1/conans/search | Search for packages
[**GetApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName**](ConanPackagesApi.md#GetApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName) | **Get** /api/v4/packages/conan/v1/files/{package_name}/{package_version}/{package_username}/{package_channel}/{recipe_revision}/export/{file_name} | Download recipe files
[**GetApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName**](ConanPackagesApi.md#GetApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName) | **Get** /api/v4/packages/conan/v1/files/{package_name}/{package_version}/{package_username}/{package_channel}/{recipe_revision}/package/{conan_package_reference}/{package_revision}/{file_name} | Download package files
[**GetApiV4PackagesConanV1Ping**](ConanPackagesApi.md#GetApiV4PackagesConanV1Ping) | **Get** /api/v4/packages/conan/v1/ping | Ping the Conan API
[**GetApiV4PackagesConanV1UsersAuthenticate**](ConanPackagesApi.md#GetApiV4PackagesConanV1UsersAuthenticate) | **Get** /api/v4/packages/conan/v1/users/authenticate | Authenticate user against conan CLI
[**GetApiV4PackagesConanV1UsersCheckCredentials**](ConanPackagesApi.md#GetApiV4PackagesConanV1UsersCheckCredentials) | **Get** /api/v4/packages/conan/v1/users/check_credentials | Check for valid user credentials per conan CLI
[**GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel**](ConanPackagesApi.md#GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel) | **Get** /api/v4/projects/{id}/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel} | Recipe Snapshot
[**GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigest**](ConanPackagesApi.md#GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigest) | **Get** /api/v4/projects/{id}/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel}/digest | Recipe Digest
[**GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrls**](ConanPackagesApi.md#GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrls) | **Get** /api/v4/projects/{id}/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel}/download_urls | Recipe Download Urls
[**GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReference**](ConanPackagesApi.md#GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReference) | **Get** /api/v4/projects/{id}/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel}/packages/{conan_package_reference} | Package Snapshot
[**GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigest**](ConanPackagesApi.md#GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigest) | **Get** /api/v4/projects/{id}/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel}/packages/{conan_package_reference}/digest | Package Digest
[**GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrls**](ConanPackagesApi.md#GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrls) | **Get** /api/v4/projects/{id}/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel}/packages/{conan_package_reference}/download_urls | Package Download Urls
[**GetApiV4ProjectsIdPackagesConanV1ConansSearch**](ConanPackagesApi.md#GetApiV4ProjectsIdPackagesConanV1ConansSearch) | **Get** /api/v4/projects/{id}/packages/conan/v1/conans/search | Search for packages
[**GetApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName**](ConanPackagesApi.md#GetApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName) | **Get** /api/v4/projects/{id}/packages/conan/v1/files/{package_name}/{package_version}/{package_username}/{package_channel}/{recipe_revision}/export/{file_name} | Download recipe files
[**GetApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName**](ConanPackagesApi.md#GetApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName) | **Get** /api/v4/projects/{id}/packages/conan/v1/files/{package_name}/{package_version}/{package_username}/{package_channel}/{recipe_revision}/package/{conan_package_reference}/{package_revision}/{file_name} | Download package files
[**GetApiV4ProjectsIdPackagesConanV1Ping**](ConanPackagesApi.md#GetApiV4ProjectsIdPackagesConanV1Ping) | **Get** /api/v4/projects/{id}/packages/conan/v1/ping | Ping the Conan API
[**GetApiV4ProjectsIdPackagesConanV1UsersAuthenticate**](ConanPackagesApi.md#GetApiV4ProjectsIdPackagesConanV1UsersAuthenticate) | **Get** /api/v4/projects/{id}/packages/conan/v1/users/authenticate | Authenticate user against conan CLI
[**GetApiV4ProjectsIdPackagesConanV1UsersCheckCredentials**](ConanPackagesApi.md#GetApiV4ProjectsIdPackagesConanV1UsersCheckCredentials) | **Get** /api/v4/projects/{id}/packages/conan/v1/users/check_credentials | Check for valid user credentials per conan CLI
[**PostApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceUploadUrls**](ConanPackagesApi.md#PostApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceUploadUrls) | **Post** /api/v4/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel}/packages/{conan_package_reference}/upload_urls | Package Upload Urls
[**PostApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelUploadUrls**](ConanPackagesApi.md#PostApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelUploadUrls) | **Post** /api/v4/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel}/upload_urls | Recipe Upload Urls
[**PostApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceUploadUrls**](ConanPackagesApi.md#PostApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceUploadUrls) | **Post** /api/v4/projects/{id}/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel}/packages/{conan_package_reference}/upload_urls | Package Upload Urls
[**PostApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelUploadUrls**](ConanPackagesApi.md#PostApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelUploadUrls) | **Post** /api/v4/projects/{id}/packages/conan/v1/conans/{package_name}/{package_version}/{package_username}/{package_channel}/upload_urls | Recipe Upload Urls
[**PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName**](ConanPackagesApi.md#PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName) | **Put** /api/v4/packages/conan/v1/files/{package_name}/{package_version}/{package_username}/{package_channel}/{recipe_revision}/export/{file_name} | Upload recipe package files
[**PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameAuthorize**](ConanPackagesApi.md#PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameAuthorize) | **Put** /api/v4/packages/conan/v1/files/{package_name}/{package_version}/{package_username}/{package_channel}/{recipe_revision}/export/{file_name}/authorize | Workhorse authorize the conan recipe file
[**PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName**](ConanPackagesApi.md#PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName) | **Put** /api/v4/packages/conan/v1/files/{package_name}/{package_version}/{package_username}/{package_channel}/{recipe_revision}/package/{conan_package_reference}/{package_revision}/{file_name} | Upload package files
[**PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameAuthorize**](ConanPackagesApi.md#PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameAuthorize) | **Put** /api/v4/packages/conan/v1/files/{package_name}/{package_version}/{package_username}/{package_channel}/{recipe_revision}/package/{conan_package_reference}/{package_revision}/{file_name}/authorize | Workhorse authorize the conan package file
[**PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName**](ConanPackagesApi.md#PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName) | **Put** /api/v4/projects/{id}/packages/conan/v1/files/{package_name}/{package_version}/{package_username}/{package_channel}/{recipe_revision}/export/{file_name} | Upload recipe package files
[**PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameAuthorize**](ConanPackagesApi.md#PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameAuthorize) | **Put** /api/v4/projects/{id}/packages/conan/v1/files/{package_name}/{package_version}/{package_username}/{package_channel}/{recipe_revision}/export/{file_name}/authorize | Workhorse authorize the conan recipe file
[**PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName**](ConanPackagesApi.md#PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName) | **Put** /api/v4/projects/{id}/packages/conan/v1/files/{package_name}/{package_version}/{package_username}/{package_channel}/{recipe_revision}/package/{conan_package_reference}/{package_revision}/{file_name} | Upload package files
[**PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameAuthorize**](ConanPackagesApi.md#PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameAuthorize) | **Put** /api/v4/projects/{id}/packages/conan/v1/files/{package_name}/{package_version}/{package_username}/{package_channel}/{recipe_revision}/package/{conan_package_reference}/{package_revision}/{file_name}/authorize | Workhorse authorize the conan package file


# **DeleteApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel**
> DeleteApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel(ctx, packageName, packageVersion, packageUsername, packageChannel)
Delete Package

This feature was introduced in GitLab 12.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **packageName** | **string**| Package name | 
  **packageVersion** | **string**| Package version | 
  **packageUsername** | **string**| Package username | 
  **packageChannel** | **string**| Package channel | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel**
> DeleteApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel(ctx, id, packageName, packageVersion, packageUsername, packageChannel)
Delete Package

This feature was introduced in GitLab 12.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **packageName** | **string**| Package name | 
  **packageVersion** | **string**| Package version | 
  **packageUsername** | **string**| Package username | 
  **packageChannel** | **string**| Package channel | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel**
> ApiEntitiesConanPackageConanRecipeSnapshot GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel(ctx, packageName, packageVersion, packageUsername, packageChannel)
Recipe Snapshot

This feature was introduced in GitLab 12.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **packageName** | **string**| Package name | 
  **packageVersion** | **string**| Package version | 
  **packageUsername** | **string**| Package username | 
  **packageChannel** | **string**| Package channel | 

### Return type

[**ApiEntitiesConanPackageConanRecipeSnapshot**](API_Entities_ConanPackage_ConanRecipeSnapshot.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigest**
> ApiEntitiesConanPackageConanRecipeManifest GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigest(ctx, packageName, packageVersion, packageUsername, packageChannel)
Recipe Digest

This feature was introduced in GitLab 12.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **packageName** | **string**| Package name | 
  **packageVersion** | **string**| Package version | 
  **packageUsername** | **string**| Package username | 
  **packageChannel** | **string**| Package channel | 

### Return type

[**ApiEntitiesConanPackageConanRecipeManifest**](API_Entities_ConanPackage_ConanRecipeManifest.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrls**
> ApiEntitiesConanPackageConanRecipeManifest GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrls(ctx, packageName, packageVersion, packageUsername, packageChannel)
Recipe Download Urls

This feature was introduced in GitLab 12.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **packageName** | **string**| Package name | 
  **packageVersion** | **string**| Package version | 
  **packageUsername** | **string**| Package username | 
  **packageChannel** | **string**| Package channel | 

### Return type

[**ApiEntitiesConanPackageConanRecipeManifest**](API_Entities_ConanPackage_ConanRecipeManifest.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReference**
> ApiEntitiesConanPackageConanPackageSnapshot GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReference(ctx, packageName, packageVersion, packageUsername, packageChannel, conanPackageReference)
Package Snapshot

This feature was introduced in GitLab 12.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **packageName** | **string**| Package name | 
  **packageVersion** | **string**| Package version | 
  **packageUsername** | **string**| Package username | 
  **packageChannel** | **string**| Package channel | 
  **conanPackageReference** | **string**| Conan package ID | 

### Return type

[**ApiEntitiesConanPackageConanPackageSnapshot**](API_Entities_ConanPackage_ConanPackageSnapshot.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigest**
> ApiEntitiesConanPackageConanPackageManifest GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigest(ctx, packageName, packageVersion, packageUsername, packageChannel, conanPackageReference)
Package Digest

This feature was introduced in GitLab 12.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **packageName** | **string**| Package name | 
  **packageVersion** | **string**| Package version | 
  **packageUsername** | **string**| Package username | 
  **packageChannel** | **string**| Package channel | 
  **conanPackageReference** | **string**| Conan package ID | 

### Return type

[**ApiEntitiesConanPackageConanPackageManifest**](API_Entities_ConanPackage_ConanPackageManifest.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrls**
> ApiEntitiesConanPackageConanPackageManifest GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrls(ctx, packageName, packageVersion, packageUsername, packageChannel, conanPackageReference)
Package Download Urls

This feature was introduced in GitLab 12.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **packageName** | **string**| Package name | 
  **packageVersion** | **string**| Package version | 
  **packageUsername** | **string**| Package username | 
  **packageChannel** | **string**| Package channel | 
  **conanPackageReference** | **string**| Conan package ID | 

### Return type

[**ApiEntitiesConanPackageConanPackageManifest**](API_Entities_ConanPackage_ConanPackageManifest.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4PackagesConanV1ConansSearch**
> GetApiV4PackagesConanV1ConansSearch(ctx, q)
Search for packages

This feature was introduced in GitLab 12.4

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **q** | **string**| Search query | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName**
> GetApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName(ctx, packageName, packageVersion, packageUsername, packageChannel, recipeRevision, fileName)
Download recipe files

This feature was introduced in GitLab 12.6

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **packageName** | **string**| Package name | 
  **packageVersion** | **string**| Package version | 
  **packageUsername** | **string**| Package username | 
  **packageChannel** | **string**| Package channel | 
  **recipeRevision** | **string**| Conan Recipe Revision | 
  **fileName** | **string**| Package file name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName**
> GetApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName(ctx, packageName, packageVersion, packageUsername, packageChannel, recipeRevision, conanPackageReference, packageRevision, fileName)
Download package files

This feature was introduced in GitLab 12.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **packageName** | **string**| Package name | 
  **packageVersion** | **string**| Package version | 
  **packageUsername** | **string**| Package username | 
  **packageChannel** | **string**| Package channel | 
  **recipeRevision** | **string**| Conan Recipe Revision | 
  **conanPackageReference** | **string**| Conan Package ID | 
  **packageRevision** | **string**| Conan Package Revision | 
  **fileName** | **string**| Package file name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4PackagesConanV1Ping**
> GetApiV4PackagesConanV1Ping(ctx, )
Ping the Conan API

This feature was introduced in GitLab 12.2

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4PackagesConanV1UsersAuthenticate**
> GetApiV4PackagesConanV1UsersAuthenticate(ctx, )
Authenticate user against conan CLI

This feature was introduced in GitLab 12.2

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4PackagesConanV1UsersCheckCredentials**
> GetApiV4PackagesConanV1UsersCheckCredentials(ctx, )
Check for valid user credentials per conan CLI

This feature was introduced in GitLab 12.4

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel**
> ApiEntitiesConanPackageConanRecipeSnapshot GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel(ctx, id, packageName, packageVersion, packageUsername, packageChannel)
Recipe Snapshot

This feature was introduced in GitLab 12.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **packageName** | **string**| Package name | 
  **packageVersion** | **string**| Package version | 
  **packageUsername** | **string**| Package username | 
  **packageChannel** | **string**| Package channel | 

### Return type

[**ApiEntitiesConanPackageConanRecipeSnapshot**](API_Entities_ConanPackage_ConanRecipeSnapshot.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigest**
> ApiEntitiesConanPackageConanRecipeManifest GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigest(ctx, id, packageName, packageVersion, packageUsername, packageChannel)
Recipe Digest

This feature was introduced in GitLab 12.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **packageName** | **string**| Package name | 
  **packageVersion** | **string**| Package version | 
  **packageUsername** | **string**| Package username | 
  **packageChannel** | **string**| Package channel | 

### Return type

[**ApiEntitiesConanPackageConanRecipeManifest**](API_Entities_ConanPackage_ConanRecipeManifest.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrls**
> ApiEntitiesConanPackageConanRecipeManifest GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrls(ctx, id, packageName, packageVersion, packageUsername, packageChannel)
Recipe Download Urls

This feature was introduced in GitLab 12.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **packageName** | **string**| Package name | 
  **packageVersion** | **string**| Package version | 
  **packageUsername** | **string**| Package username | 
  **packageChannel** | **string**| Package channel | 

### Return type

[**ApiEntitiesConanPackageConanRecipeManifest**](API_Entities_ConanPackage_ConanRecipeManifest.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReference**
> ApiEntitiesConanPackageConanPackageSnapshot GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReference(ctx, id, packageName, packageVersion, packageUsername, packageChannel, conanPackageReference)
Package Snapshot

This feature was introduced in GitLab 12.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **packageName** | **string**| Package name | 
  **packageVersion** | **string**| Package version | 
  **packageUsername** | **string**| Package username | 
  **packageChannel** | **string**| Package channel | 
  **conanPackageReference** | **string**| Conan package ID | 

### Return type

[**ApiEntitiesConanPackageConanPackageSnapshot**](API_Entities_ConanPackage_ConanPackageSnapshot.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigest**
> ApiEntitiesConanPackageConanPackageManifest GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigest(ctx, id, packageName, packageVersion, packageUsername, packageChannel, conanPackageReference)
Package Digest

This feature was introduced in GitLab 12.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **packageName** | **string**| Package name | 
  **packageVersion** | **string**| Package version | 
  **packageUsername** | **string**| Package username | 
  **packageChannel** | **string**| Package channel | 
  **conanPackageReference** | **string**| Conan package ID | 

### Return type

[**ApiEntitiesConanPackageConanPackageManifest**](API_Entities_ConanPackage_ConanPackageManifest.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrls**
> ApiEntitiesConanPackageConanPackageManifest GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrls(ctx, id, packageName, packageVersion, packageUsername, packageChannel, conanPackageReference)
Package Download Urls

This feature was introduced in GitLab 12.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **packageName** | **string**| Package name | 
  **packageVersion** | **string**| Package version | 
  **packageUsername** | **string**| Package username | 
  **packageChannel** | **string**| Package channel | 
  **conanPackageReference** | **string**| Conan package ID | 

### Return type

[**ApiEntitiesConanPackageConanPackageManifest**](API_Entities_ConanPackage_ConanPackageManifest.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesConanV1ConansSearch**
> GetApiV4ProjectsIdPackagesConanV1ConansSearch(ctx, id, q)
Search for packages

This feature was introduced in GitLab 12.4

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **q** | **string**| Search query | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName**
> GetApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName(ctx, id, packageName, packageVersion, packageUsername, packageChannel, recipeRevision, fileName)
Download recipe files

This feature was introduced in GitLab 12.6

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **packageName** | **string**| Package name | 
  **packageVersion** | **string**| Package version | 
  **packageUsername** | **string**| Package username | 
  **packageChannel** | **string**| Package channel | 
  **recipeRevision** | **string**| Conan Recipe Revision | 
  **fileName** | **string**| Package file name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName**
> GetApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName(ctx, id, packageName, packageVersion, packageUsername, packageChannel, recipeRevision, conanPackageReference, packageRevision, fileName)
Download package files

This feature was introduced in GitLab 12.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **packageName** | **string**| Package name | 
  **packageVersion** | **string**| Package version | 
  **packageUsername** | **string**| Package username | 
  **packageChannel** | **string**| Package channel | 
  **recipeRevision** | **string**| Conan Recipe Revision | 
  **conanPackageReference** | **string**| Conan Package ID | 
  **packageRevision** | **string**| Conan Package Revision | 
  **fileName** | **string**| Package file name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesConanV1Ping**
> GetApiV4ProjectsIdPackagesConanV1Ping(ctx, id)
Ping the Conan API

This feature was introduced in GitLab 12.2

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

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesConanV1UsersAuthenticate**
> GetApiV4ProjectsIdPackagesConanV1UsersAuthenticate(ctx, id)
Authenticate user against conan CLI

This feature was introduced in GitLab 12.2

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

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesConanV1UsersCheckCredentials**
> GetApiV4ProjectsIdPackagesConanV1UsersCheckCredentials(ctx, id)
Check for valid user credentials per conan CLI

This feature was introduced in GitLab 12.4

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

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceUploadUrls**
> ApiEntitiesConanPackageConanUploadUrls PostApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceUploadUrls(ctx, packageName, packageVersion, packageUsername, packageChannel, conanPackageReference)
Package Upload Urls

This feature was introduced in GitLab 12.4

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **packageName** | **string**| Package name | 
  **packageVersion** | **string**| Package version | 
  **packageUsername** | **string**| Package username | 
  **packageChannel** | **string**| Package channel | 
  **conanPackageReference** | **string**| Conan package ID | 

### Return type

[**ApiEntitiesConanPackageConanUploadUrls**](API_Entities_ConanPackage_ConanUploadUrls.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelUploadUrls**
> ApiEntitiesConanPackageConanUploadUrls PostApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelUploadUrls(ctx, packageName, packageVersion, packageUsername, packageChannel)
Recipe Upload Urls

This feature was introduced in GitLab 12.4

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **packageName** | **string**| Package name | 
  **packageVersion** | **string**| Package version | 
  **packageUsername** | **string**| Package username | 
  **packageChannel** | **string**| Package channel | 

### Return type

[**ApiEntitiesConanPackageConanUploadUrls**](API_Entities_ConanPackage_ConanUploadUrls.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceUploadUrls**
> ApiEntitiesConanPackageConanUploadUrls PostApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceUploadUrls(ctx, id, packageName, packageVersion, packageUsername, packageChannel, conanPackageReference)
Package Upload Urls

This feature was introduced in GitLab 12.4

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **packageName** | **string**| Package name | 
  **packageVersion** | **string**| Package version | 
  **packageUsername** | **string**| Package username | 
  **packageChannel** | **string**| Package channel | 
  **conanPackageReference** | **string**| Conan package ID | 

### Return type

[**ApiEntitiesConanPackageConanUploadUrls**](API_Entities_ConanPackage_ConanUploadUrls.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelUploadUrls**
> ApiEntitiesConanPackageConanUploadUrls PostApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelUploadUrls(ctx, id, packageName, packageVersion, packageUsername, packageChannel)
Recipe Upload Urls

This feature was introduced in GitLab 12.4

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **packageName** | **string**| Package name | 
  **packageVersion** | **string**| Package version | 
  **packageUsername** | **string**| Package username | 
  **packageChannel** | **string**| Package channel | 

### Return type

[**ApiEntitiesConanPackageConanUploadUrls**](API_Entities_ConanPackage_ConanUploadUrls.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName**
> PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName(ctx, packageName, packageVersion, packageUsername, packageChannel, recipeRevision, fileName, putApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName)
Upload recipe package files

This feature was introduced in GitLab 12.6

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **packageName** | **string**| Package name | 
  **packageVersion** | **string**| Package version | 
  **packageUsername** | **string**| Package username | 
  **packageChannel** | **string**| Package channel | 
  **recipeRevision** | **string**| Conan Recipe Revision | 
  **fileName** | **string**| Package file name | 
  **putApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName** | [**PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName**](PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameAuthorize**
> PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameAuthorize(ctx, packageName, packageVersion, packageUsername, packageChannel, recipeRevision, fileName)
Workhorse authorize the conan recipe file

This feature was introduced in GitLab 12.6

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **packageName** | **string**| Package name | 
  **packageVersion** | **string**| Package version | 
  **packageUsername** | **string**| Package username | 
  **packageChannel** | **string**| Package channel | 
  **recipeRevision** | **string**| Conan Recipe Revision | 
  **fileName** | **string**| Package file name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName**
> PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName(ctx, packageName, packageVersion, packageUsername, packageChannel, recipeRevision, conanPackageReference, packageRevision, fileName, putApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName)
Upload package files

This feature was introduced in GitLab 12.6

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **packageName** | **string**| Package name | 
  **packageVersion** | **string**| Package version | 
  **packageUsername** | **string**| Package username | 
  **packageChannel** | **string**| Package channel | 
  **recipeRevision** | **string**| Conan Recipe Revision | 
  **conanPackageReference** | **string**| Conan Package ID | 
  **packageRevision** | **string**| Conan Package Revision | 
  **fileName** | **string**| Package file name | 
  **putApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName** | [**PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName**](PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameAuthorize**
> PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameAuthorize(ctx, packageName, packageVersion, packageUsername, packageChannel, recipeRevision, conanPackageReference, packageRevision, fileName)
Workhorse authorize the conan package file

This feature was introduced in GitLab 12.6

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **packageName** | **string**| Package name | 
  **packageVersion** | **string**| Package version | 
  **packageUsername** | **string**| Package username | 
  **packageChannel** | **string**| Package channel | 
  **recipeRevision** | **string**| Conan Recipe Revision | 
  **conanPackageReference** | **string**| Conan Package ID | 
  **packageRevision** | **string**| Conan Package Revision | 
  **fileName** | **string**| Package file name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName**
> PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName(ctx, id, packageName, packageVersion, packageUsername, packageChannel, recipeRevision, fileName, putApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName)
Upload recipe package files

This feature was introduced in GitLab 12.6

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **packageName** | **string**| Package name | 
  **packageVersion** | **string**| Package version | 
  **packageUsername** | **string**| Package username | 
  **packageChannel** | **string**| Package channel | 
  **recipeRevision** | **string**| Conan Recipe Revision | 
  **fileName** | **string**| Package file name | 
  **putApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName** | [**PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName**](PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameAuthorize**
> PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameAuthorize(ctx, id, packageName, packageVersion, packageUsername, packageChannel, recipeRevision, fileName)
Workhorse authorize the conan recipe file

This feature was introduced in GitLab 12.6

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **packageName** | **string**| Package name | 
  **packageVersion** | **string**| Package version | 
  **packageUsername** | **string**| Package username | 
  **packageChannel** | **string**| Package channel | 
  **recipeRevision** | **string**| Conan Recipe Revision | 
  **fileName** | **string**| Package file name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName**
> PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName(ctx, id, packageName, packageVersion, packageUsername, packageChannel, recipeRevision, conanPackageReference, packageRevision, fileName, putApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName)
Upload package files

This feature was introduced in GitLab 12.6

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **packageName** | **string**| Package name | 
  **packageVersion** | **string**| Package version | 
  **packageUsername** | **string**| Package username | 
  **packageChannel** | **string**| Package channel | 
  **recipeRevision** | **string**| Conan Recipe Revision | 
  **conanPackageReference** | **string**| Conan Package ID | 
  **packageRevision** | **string**| Conan Package Revision | 
  **fileName** | **string**| Package file name | 
  **putApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName** | [**PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName**](PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameAuthorize**
> PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameAuthorize(ctx, id, packageName, packageVersion, packageUsername, packageChannel, recipeRevision, conanPackageReference, packageRevision, fileName)
Workhorse authorize the conan package file

This feature was introduced in GitLab 12.6

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **packageName** | **string**| Package name | 
  **packageVersion** | **string**| Package version | 
  **packageUsername** | **string**| Package username | 
  **packageChannel** | **string**| Package channel | 
  **recipeRevision** | **string**| Conan Recipe Revision | 
  **conanPackageReference** | **string**| Conan Package ID | 
  **packageRevision** | **string**| Conan Package Revision | 
  **fileName** | **string**| Package file name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

