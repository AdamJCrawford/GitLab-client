# \NugetPackagesApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4ProjectsIdPackagesNugetpackageNamepackageVersion**](NugetPackagesApi.md#DeleteApiV4ProjectsIdPackagesNugetpackageNamepackageVersion) | **Delete** /api/v4/projects/{id}/packages/nuget/*package_name/*package_version | The NuGet Package Delete endpoint
[**GetApiV4GroupsIdPackagesNugetIndex**](NugetPackagesApi.md#GetApiV4GroupsIdPackagesNugetIndex) | **Get** /api/v4/groups/{id}/-/packages/nuget/index | The NuGet V3 Feed Service Index
[**GetApiV4GroupsIdPackagesNugetMetadatapackageNameIndex**](NugetPackagesApi.md#GetApiV4GroupsIdPackagesNugetMetadatapackageNameIndex) | **Get** /api/v4/groups/{id}/-/packages/nuget/metadata/*package_name/index | The NuGet Metadata Service - Package name level
[**GetApiV4GroupsIdPackagesNugetMetadatapackageNamepackageVersion**](NugetPackagesApi.md#GetApiV4GroupsIdPackagesNugetMetadatapackageNamepackageVersion) | **Get** /api/v4/groups/{id}/-/packages/nuget/metadata/*package_name/*package_version | The NuGet Metadata Service - Package name and version level
[**GetApiV4GroupsIdPackagesNugetQuery**](NugetPackagesApi.md#GetApiV4GroupsIdPackagesNugetQuery) | **Get** /api/v4/groups/{id}/-/packages/nuget/query | The NuGet Search Service
[**GetApiV4GroupsIdPackagesNugetSymbolfilesfileNamesignaturesameFileName**](NugetPackagesApi.md#GetApiV4GroupsIdPackagesNugetSymbolfilesfileNamesignaturesameFileName) | **Get** /api/v4/groups/{id}/-/packages/nuget/symbolfiles/*file_name/*signature/*same_file_name | The NuGet Symbol File Download Endpoint
[**GetApiV4GroupsIdPackagesNugetV2**](NugetPackagesApi.md#GetApiV4GroupsIdPackagesNugetV2) | **Get** /api/v4/groups/{id}/-/packages/nuget/v2 | The NuGet V2 Feed Service Index
[**GetApiV4GroupsIdPackagesNugetV2metadata**](NugetPackagesApi.md#GetApiV4GroupsIdPackagesNugetV2metadata) | **Get** /api/v4/groups/{id}/-/packages/nuget/v2/$metadata | The NuGet V2 Feed Package $metadata endpoint
[**GetApiV4ProjectsIdPackagesNugetDownloadpackageNameIndex**](NugetPackagesApi.md#GetApiV4ProjectsIdPackagesNugetDownloadpackageNameIndex) | **Get** /api/v4/projects/{id}/packages/nuget/download/*package_name/index | The NuGet Content Service - index request
[**GetApiV4ProjectsIdPackagesNugetDownloadpackageNamepackageVersionpackageFilename**](NugetPackagesApi.md#GetApiV4ProjectsIdPackagesNugetDownloadpackageNamepackageVersionpackageFilename) | **Get** /api/v4/projects/{id}/packages/nuget/download/*package_name/*package_version/*package_filename | The NuGet Content Service - content request
[**GetApiV4ProjectsIdPackagesNugetIndex**](NugetPackagesApi.md#GetApiV4ProjectsIdPackagesNugetIndex) | **Get** /api/v4/projects/{id}/packages/nuget/index | The NuGet V3 Feed Service Index
[**GetApiV4ProjectsIdPackagesNugetMetadatapackageNameIndex**](NugetPackagesApi.md#GetApiV4ProjectsIdPackagesNugetMetadatapackageNameIndex) | **Get** /api/v4/projects/{id}/packages/nuget/metadata/*package_name/index | The NuGet Metadata Service - Package name level
[**GetApiV4ProjectsIdPackagesNugetMetadatapackageNamepackageVersion**](NugetPackagesApi.md#GetApiV4ProjectsIdPackagesNugetMetadatapackageNamepackageVersion) | **Get** /api/v4/projects/{id}/packages/nuget/metadata/*package_name/*package_version | The NuGet Metadata Service - Package name and version level
[**GetApiV4ProjectsIdPackagesNugetQuery**](NugetPackagesApi.md#GetApiV4ProjectsIdPackagesNugetQuery) | **Get** /api/v4/projects/{id}/packages/nuget/query | The NuGet Search Service
[**GetApiV4ProjectsIdPackagesNugetSymbolfilesfileNamesignaturesameFileName**](NugetPackagesApi.md#GetApiV4ProjectsIdPackagesNugetSymbolfilesfileNamesignaturesameFileName) | **Get** /api/v4/projects/{id}/packages/nuget/symbolfiles/*file_name/*signature/*same_file_name | The NuGet Symbol File Download Endpoint
[**GetApiV4ProjectsIdPackagesNugetV2**](NugetPackagesApi.md#GetApiV4ProjectsIdPackagesNugetV2) | **Get** /api/v4/projects/{id}/packages/nuget/v2 | The NuGet V2 Feed Service Index
[**GetApiV4ProjectsIdPackagesNugetV2metadata**](NugetPackagesApi.md#GetApiV4ProjectsIdPackagesNugetV2metadata) | **Get** /api/v4/projects/{id}/packages/nuget/v2/$metadata | The NuGet V2 Feed Package $metadata endpoint
[**GetApiV4ProjectsProjectIdPackagesNugetV2Findpackagesbyid_**](NugetPackagesApi.md#GetApiV4ProjectsProjectIdPackagesNugetV2Findpackagesbyid_) | **Get** /api/v4/projects/{project_id}/packages/nuget/v2/FindPackagesById\(\) | The NuGet V2 Feed Find Packages by ID endpoint
[**GetApiV4ProjectsProjectIdPackagesNugetV2PackagesIdpackageNameversionpackageVersion**](NugetPackagesApi.md#GetApiV4ProjectsProjectIdPackagesNugetV2PackagesIdpackageNameversionpackageVersion) | **Get** /api/v4/projects/{project_id}/packages/nuget/v2/Packages\(Id&#x3D;&#39;*package_name&#39;,Version&#x3D;&#39;*package_version&#39;\) | The NuGet V2 Feed Single Package Metadata endpoint
[**GetApiV4ProjectsProjectIdPackagesNugetV2Packages_**](NugetPackagesApi.md#GetApiV4ProjectsProjectIdPackagesNugetV2Packages_) | **Get** /api/v4/projects/{project_id}/packages/nuget/v2/Packages\(\) | The NuGet V2 Feed Enumerate Packages endpoint
[**PutApiV4ProjectsIdPackagesNuget**](NugetPackagesApi.md#PutApiV4ProjectsIdPackagesNuget) | **Put** /api/v4/projects/{id}/packages/nuget | The NuGet V3 Feed Package Publish endpoint
[**PutApiV4ProjectsIdPackagesNugetAuthorize**](NugetPackagesApi.md#PutApiV4ProjectsIdPackagesNugetAuthorize) | **Put** /api/v4/projects/{id}/packages/nuget/authorize | The NuGet Package Authorize endpoint
[**PutApiV4ProjectsIdPackagesNugetSymbolpackage**](NugetPackagesApi.md#PutApiV4ProjectsIdPackagesNugetSymbolpackage) | **Put** /api/v4/projects/{id}/packages/nuget/symbolpackage | The NuGet Symbol Package Publish endpoint
[**PutApiV4ProjectsIdPackagesNugetSymbolpackageAuthorize**](NugetPackagesApi.md#PutApiV4ProjectsIdPackagesNugetSymbolpackageAuthorize) | **Put** /api/v4/projects/{id}/packages/nuget/symbolpackage/authorize | The NuGet Symbol Package Authorize endpoint
[**PutApiV4ProjectsIdPackagesNugetV2**](NugetPackagesApi.md#PutApiV4ProjectsIdPackagesNugetV2) | **Put** /api/v4/projects/{id}/packages/nuget/v2 | The NuGet V2 Feed Package Publish endpoint
[**PutApiV4ProjectsIdPackagesNugetV2Authorize**](NugetPackagesApi.md#PutApiV4ProjectsIdPackagesNugetV2Authorize) | **Put** /api/v4/projects/{id}/packages/nuget/v2/authorize | The NuGet V2 Feed Package Authorize endpoint


# **DeleteApiV4ProjectsIdPackagesNugetpackageNamepackageVersion**
> DeleteApiV4ProjectsIdPackagesNugetpackageNamepackageVersion(ctx, id, packageName, packageVersion)
The NuGet Package Delete endpoint

This feature was introduced in GitLab 16.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **packageName** | **string**| The NuGet package name | 
  **packageVersion** | **string**| The NuGet package version | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdPackagesNugetIndex**
> ApiEntitiesNugetServiceIndex GetApiV4GroupsIdPackagesNugetIndex(ctx, id)
The NuGet V3 Feed Service Index

This feature was introduced in GitLab 12.6

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The group ID or full group path. | 

### Return type

[**ApiEntitiesNugetServiceIndex**](API_Entities_Nuget_ServiceIndex.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdPackagesNugetMetadatapackageNameIndex**
> ApiEntitiesNugetPackagesMetadata GetApiV4GroupsIdPackagesNugetMetadatapackageNameIndex(ctx, id, packageName)
The NuGet Metadata Service - Package name level

This feature was introduced in GitLab 12.8

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The group ID or full group path. | 
  **packageName** | **string**| The NuGet package name | 

### Return type

[**ApiEntitiesNugetPackagesMetadata**](API_Entities_Nuget_PackagesMetadata.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdPackagesNugetMetadatapackageNamepackageVersion**
> ApiEntitiesNugetPackageMetadata GetApiV4GroupsIdPackagesNugetMetadatapackageNamepackageVersion(ctx, id, packageName, packageVersion)
The NuGet Metadata Service - Package name and version level

This feature was introduced in GitLab 12.8

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The group ID or full group path. | 
  **packageName** | **string**| The NuGet package name | 
  **packageVersion** | **string**| The NuGet package version | 

### Return type

[**ApiEntitiesNugetPackageMetadata**](API_Entities_Nuget_PackageMetadata.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdPackagesNugetQuery**
> ApiEntitiesNugetSearchResults GetApiV4GroupsIdPackagesNugetQuery(ctx, id, optional)
The NuGet Search Service

This feature was introduced in GitLab 12.8

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The group ID or full group path. | 
 **optional** | ***NugetPackagesApiGetApiV4GroupsIdPackagesNugetQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NugetPackagesApiGetApiV4GroupsIdPackagesNugetQueryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **optional.String**| The search term | 
 **skip** | **optional.Int32**| The number of results to skip | [default to 0]
 **take** | **optional.Int32**| The number of results to return | [default to 20]
 **prerelease** | **optional.Bool**| Include prerelease versions | [default to true]

### Return type

[**ApiEntitiesNugetSearchResults**](API_Entities_Nuget_SearchResults.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdPackagesNugetSymbolfilesfileNamesignaturesameFileName**
> GetApiV4GroupsIdPackagesNugetSymbolfilesfileNamesignaturesameFileName(ctx, symbolchecksum, id, fileName, signature, sameFileName)
The NuGet Symbol File Download Endpoint

This feature was introduced in GitLab 16.7

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbolchecksum** | **string**|  | 
  **id** | **int32**| The group ID or full group path. | 
  **fileName** | **string**| The symbol file name | 
  **signature** | **string**| The symbol file signature | 
  **sameFileName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdPackagesNugetV2**
> GetApiV4GroupsIdPackagesNugetV2(ctx, id)
The NuGet V2 Feed Service Index

This feature was introduced in GitLab 16.2

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The group ID or full group path. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdPackagesNugetV2metadata**
> GetApiV4GroupsIdPackagesNugetV2metadata(ctx, id)
The NuGet V2 Feed Package $metadata endpoint

This feature was introduced in GitLab 16.3

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The group ID or full group path. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesNugetDownloadpackageNameIndex**
> ApiEntitiesNugetPackagesVersions GetApiV4ProjectsIdPackagesNugetDownloadpackageNameIndex(ctx, id, packageName)
The NuGet Content Service - index request

This feature was introduced in GitLab 12.8

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **packageName** | **string**| The NuGet package name | 

### Return type

[**ApiEntitiesNugetPackagesVersions**](API_Entities_Nuget_PackagesVersions.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesNugetDownloadpackageNamepackageVersionpackageFilename**
> GetApiV4ProjectsIdPackagesNugetDownloadpackageNamepackageVersionpackageFilename(ctx, id, packageName, packageVersion, packageFilename)
The NuGet Content Service - content request

This feature was introduced in GitLab 12.8

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **packageName** | **string**| The NuGet package name | 
  **packageVersion** | **string**| The NuGet package version | 
  **packageFilename** | **string**| The NuGet package filename | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesNugetIndex**
> ApiEntitiesNugetServiceIndex GetApiV4ProjectsIdPackagesNugetIndex(ctx, id)
The NuGet V3 Feed Service Index

This feature was introduced in GitLab 12.6

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 

### Return type

[**ApiEntitiesNugetServiceIndex**](API_Entities_Nuget_ServiceIndex.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesNugetMetadatapackageNameIndex**
> ApiEntitiesNugetPackagesMetadata GetApiV4ProjectsIdPackagesNugetMetadatapackageNameIndex(ctx, id, packageName)
The NuGet Metadata Service - Package name level

This feature was introduced in GitLab 12.8

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **packageName** | **string**| The NuGet package name | 

### Return type

[**ApiEntitiesNugetPackagesMetadata**](API_Entities_Nuget_PackagesMetadata.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesNugetMetadatapackageNamepackageVersion**
> ApiEntitiesNugetPackageMetadata GetApiV4ProjectsIdPackagesNugetMetadatapackageNamepackageVersion(ctx, id, packageName, packageVersion)
The NuGet Metadata Service - Package name and version level

This feature was introduced in GitLab 12.8

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **packageName** | **string**| The NuGet package name | 
  **packageVersion** | **string**| The NuGet package version | 

### Return type

[**ApiEntitiesNugetPackageMetadata**](API_Entities_Nuget_PackageMetadata.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesNugetQuery**
> ApiEntitiesNugetSearchResults GetApiV4ProjectsIdPackagesNugetQuery(ctx, id, optional)
The NuGet Search Service

This feature was introduced in GitLab 12.8

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
 **optional** | ***NugetPackagesApiGetApiV4ProjectsIdPackagesNugetQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NugetPackagesApiGetApiV4ProjectsIdPackagesNugetQueryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **optional.String**| The search term | 
 **skip** | **optional.Int32**| The number of results to skip | [default to 0]
 **take** | **optional.Int32**| The number of results to return | [default to 20]
 **prerelease** | **optional.Bool**| Include prerelease versions | [default to true]

### Return type

[**ApiEntitiesNugetSearchResults**](API_Entities_Nuget_SearchResults.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesNugetSymbolfilesfileNamesignaturesameFileName**
> GetApiV4ProjectsIdPackagesNugetSymbolfilesfileNamesignaturesameFileName(ctx, symbolchecksum, id, fileName, signature, sameFileName)
The NuGet Symbol File Download Endpoint

This feature was introduced in GitLab 16.7

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbolchecksum** | **string**|  | 
  **id** | **string**| The ID or URL-encoded path of the project | 
  **fileName** | **string**| The symbol file name | 
  **signature** | **string**| The symbol file signature | 
  **sameFileName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesNugetV2**
> GetApiV4ProjectsIdPackagesNugetV2(ctx, id)
The NuGet V2 Feed Service Index

This feature was introduced in GitLab 16.2

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

# **GetApiV4ProjectsIdPackagesNugetV2metadata**
> GetApiV4ProjectsIdPackagesNugetV2metadata(ctx, id)
The NuGet V2 Feed Package $metadata endpoint

This feature was introduced in GitLab 16.3

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

# **GetApiV4ProjectsProjectIdPackagesNugetV2Findpackagesbyid_**
> GetApiV4ProjectsProjectIdPackagesNugetV2Findpackagesbyid_(ctx, projectId, id)
The NuGet V2 Feed Find Packages by ID endpoint

This feature was introduced in GitLab 16.4

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The ID or URL-encoded path of the project | 
  **id** | **string**| The NuGet package name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsProjectIdPackagesNugetV2PackagesIdpackageNameversionpackageVersion**
> GetApiV4ProjectsProjectIdPackagesNugetV2PackagesIdpackageNameversionpackageVersion(ctx, projectId, packageName, packageVersion)
The NuGet V2 Feed Single Package Metadata endpoint

This feature was introduced in GitLab 16.4

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The ID or URL-encoded path of the project | 
  **packageName** | **string**| The NuGet package name | 
  **packageVersion** | **string**| The NuGet package version | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsProjectIdPackagesNugetV2Packages_**
> GetApiV4ProjectsProjectIdPackagesNugetV2Packages_(ctx, projectId, filter)
The NuGet V2 Feed Enumerate Packages endpoint

This feature was introduced in GitLab 16.4

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The ID or URL-encoded path of the project | 
  **filter** | **string**| The NuGet package name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdPackagesNuget**
> PutApiV4ProjectsIdPackagesNuget(ctx, id, putApiV4ProjectsIdPackagesNuget)
The NuGet V3 Feed Package Publish endpoint

This feature was introduced in GitLab 12.6

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **putApiV4ProjectsIdPackagesNuget** | [**PutApiV4ProjectsIdPackagesNuget**](PutApiV4ProjectsIdPackagesNuget.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdPackagesNugetAuthorize**
> PutApiV4ProjectsIdPackagesNugetAuthorize(ctx, id)
The NuGet Package Authorize endpoint

This feature was introduced in GitLab 14.1

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

# **PutApiV4ProjectsIdPackagesNugetSymbolpackage**
> PutApiV4ProjectsIdPackagesNugetSymbolpackage(ctx, id, putApiV4ProjectsIdPackagesNugetSymbolpackage)
The NuGet Symbol Package Publish endpoint

This feature was introduced in GitLab 14.1

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **putApiV4ProjectsIdPackagesNugetSymbolpackage** | [**PutApiV4ProjectsIdPackagesNugetSymbolpackage**](PutApiV4ProjectsIdPackagesNugetSymbolpackage.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdPackagesNugetSymbolpackageAuthorize**
> PutApiV4ProjectsIdPackagesNugetSymbolpackageAuthorize(ctx, id)
The NuGet Symbol Package Authorize endpoint

This feature was introduced in GitLab 14.1

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

# **PutApiV4ProjectsIdPackagesNugetV2**
> PutApiV4ProjectsIdPackagesNugetV2(ctx, id, putApiV4ProjectsIdPackagesNugetV2)
The NuGet V2 Feed Package Publish endpoint

This feature was introduced in GitLab 16.2

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **putApiV4ProjectsIdPackagesNugetV2** | [**PutApiV4ProjectsIdPackagesNugetV2**](PutApiV4ProjectsIdPackagesNugetV2.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdPackagesNugetV2Authorize**
> PutApiV4ProjectsIdPackagesNugetV2Authorize(ctx, id)
The NuGet V2 Feed Package Authorize endpoint

This feature was introduced in GitLab 16.2

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

