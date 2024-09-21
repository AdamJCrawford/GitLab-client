# \TerraformRegistryApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystem**](TerraformRegistryApi.md#GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystem) | **Get** /api/v4/packages/terraform/modules/v1/{module_namespace}/{module_name}/{module_system} | Get details about the latest version of a module
[**GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemDownload**](TerraformRegistryApi.md#GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemDownload) | **Get** /api/v4/packages/terraform/modules/v1/{module_namespace}/{module_name}/{module_system}/download | Get download location for the latest version of a module
[**GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemVersions**](TerraformRegistryApi.md#GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemVersions) | **Get** /api/v4/packages/terraform/modules/v1/{module_namespace}/{module_name}/{module_system}/versions | List versions for a module
[**GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemmoduleVersion**](TerraformRegistryApi.md#GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemmoduleVersion) | **Get** /api/v4/packages/terraform/modules/v1/{module_namespace}/{module_name}/{module_system}/*module_version | Get details about specific version of a module
[**GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemmoduleVersionDownload**](TerraformRegistryApi.md#GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemmoduleVersionDownload) | **Get** /api/v4/packages/terraform/modules/v1/{module_namespace}/{module_name}/{module_system}/*module_version/download | Get download location for specific version of a module
[**GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemmoduleVersionFile**](TerraformRegistryApi.md#GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemmoduleVersionFile) | **Get** /api/v4/packages/terraform/modules/v1/{module_namespace}/{module_name}/{module_system}/*module_version/file | Download specific version of a module
[**GetApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystem**](TerraformRegistryApi.md#GetApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystem) | **Get** /api/v4/projects/{id}/packages/terraform/modules/{module_name}/{module_system} | Download the latest version of a module
[**GetApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersion**](TerraformRegistryApi.md#GetApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersion) | **Get** /api/v4/projects/{id}/packages/terraform/modules/{module_name}/{module_system}/*module_version | Download a specific version of a module
[**PutApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersionFile**](TerraformRegistryApi.md#PutApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersionFile) | **Put** /api/v4/projects/{id}/packages/terraform/modules/{module_name}/{module_system}/*module_version/file | Upload Terraform Module package file
[**PutApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersionFileAuthorize**](TerraformRegistryApi.md#PutApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersionFileAuthorize) | **Put** /api/v4/projects/{id}/packages/terraform/modules/{module_name}/{module_system}/*module_version/file/authorize | Workhorse authorize Terraform Module package file


# **GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystem**
> ApiEntitiesTerraformModuleVersion GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystem(ctx, moduleNamespace, moduleName, moduleSystem)
Get details about the latest version of a module

Get details about the latest version of a module

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **moduleNamespace** | **string**| Group&#39;s ID or slug | 
  **moduleName** | **string**|  | 
  **moduleSystem** | **string**|  | 

### Return type

[**ApiEntitiesTerraformModuleVersion**](API_Entities_Terraform_ModuleVersion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemDownload**
> GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemDownload(ctx, moduleNamespace, moduleName, moduleSystem)
Get download location for the latest version of a module

Download the latest version of a module

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **moduleNamespace** | **string**| Group&#39;s ID or slug | 
  **moduleName** | **string**|  | 
  **moduleSystem** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemVersions**
> []ApiEntitiesTerraformModuleVersions GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemVersions(ctx, moduleNamespace, moduleName, moduleSystem)
List versions for a module

List versions for a module

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **moduleNamespace** | **string**| Group&#39;s ID or slug | 
  **moduleName** | **string**|  | 
  **moduleSystem** | **string**|  | 

### Return type

[**[]ApiEntitiesTerraformModuleVersions**](API_Entities_Terraform_ModuleVersions.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemmoduleVersion**
> ApiEntitiesTerraformModuleVersion GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemmoduleVersion(ctx, moduleNamespace, moduleName, moduleSystem, moduleVersion)
Get details about specific version of a module

Get details about specific version of a module

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **moduleNamespace** | **string**| Group&#39;s ID or slug | 
  **moduleName** | **string**|  | 
  **moduleSystem** | **string**|  | 
  **moduleVersion** | **string**| Module version | 

### Return type

[**ApiEntitiesTerraformModuleVersion**](API_Entities_Terraform_ModuleVersion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemmoduleVersionDownload**
> GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemmoduleVersionDownload(ctx, moduleNamespace, moduleName, moduleSystem, moduleVersion)
Get download location for specific version of a module

Download specific version of a module

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **moduleNamespace** | **string**| Group&#39;s ID or slug | 
  **moduleName** | **string**|  | 
  **moduleSystem** | **string**|  | 
  **moduleVersion** | **string**| Module version | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemmoduleVersionFile**
> *os.File GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemmoduleVersionFile(ctx, moduleNamespace, moduleName, moduleSystem, moduleVersion)
Download specific version of a module

Download specific version of a module

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **moduleNamespace** | **string**| Group&#39;s ID or slug | 
  **moduleName** | **string**|  | 
  **moduleSystem** | **string**|  | 
  **moduleVersion** | **string**| Module version | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystem**
> GetApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystem(ctx, id, moduleName, moduleSystem, optional)
Download the latest version of a module

This feature was introduced in GitLab 16.7

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or full path of a project | 
  **moduleName** | **string**| Module name | 
  **moduleSystem** | **string**| Module system | 
 **optional** | ***TerraformRegistryApiGetApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TerraformRegistryApiGetApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **terraformGet** | **optional.String**| Terraform get redirection flag | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersion**
> GetApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersion(ctx, id, moduleName, moduleSystem, moduleVersion, optional)
Download a specific version of a module

This feature was introduced in GitLab 16.7

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or full path of a project | 
  **moduleName** | **string**| Module name | 
  **moduleSystem** | **string**| Module system | 
  **moduleVersion** | **string**| Module version | 
 **optional** | ***TerraformRegistryApiGetApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TerraformRegistryApiGetApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **terraformGet** | **optional.String**| Terraform get redirection flag | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersionFile**
> PutApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersionFile(ctx, id, moduleName, moduleSystem, moduleVersion, file)
Upload Terraform Module package file

This feature was introduced in GitLab 13.11

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or full path of a project | 
  **moduleName** | **string**| Module name | 
  **moduleSystem** | **string**| Module system | 
  **moduleVersion** | **string**| Module version | 
  **file** | ***os.File**| The package file to be published (generated by Multipart middleware) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersionFileAuthorize**
> PutApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersionFileAuthorize(ctx, id, moduleName, moduleSystem, putApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersionFileAuthorize)
Workhorse authorize Terraform Module package file

This feature was introduced in GitLab 13.11

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or full path of a project | 
  **moduleName** | **string**| Module name | 
  **moduleSystem** | **string**| Module system | 
  **putApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersionFileAuthorize** | [**PutApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersionFileAuthorize**](PutApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersionFileAuthorize.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

