# \DebianDistributionApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4GroupsIdDebianDistributionsCodename**](DebianDistributionApi.md#DeleteApiV4GroupsIdDebianDistributionsCodename) | **Delete** /api/v4/groups/{id}/-/debian_distributions/{codename} | Delete a Debian Distribution
[**DeleteApiV4ProjectsIdDebianDistributionsCodename**](DebianDistributionApi.md#DeleteApiV4ProjectsIdDebianDistributionsCodename) | **Delete** /api/v4/projects/{id}/debian_distributions/{codename} | Delete a Debian Distribution
[**GetApiV4GroupsIdDebianDistributions**](DebianDistributionApi.md#GetApiV4GroupsIdDebianDistributions) | **Get** /api/v4/groups/{id}/-/debian_distributions | Get a list of Debian Distributions
[**GetApiV4GroupsIdDebianDistributionsCodename**](DebianDistributionApi.md#GetApiV4GroupsIdDebianDistributionsCodename) | **Get** /api/v4/groups/{id}/-/debian_distributions/{codename} | Get a Debian Distribution
[**GetApiV4GroupsIdDebianDistributionsCodenameKeyAsc**](DebianDistributionApi.md#GetApiV4GroupsIdDebianDistributionsCodenameKeyAsc) | **Get** /api/v4/groups/{id}/-/debian_distributions/{codename}/key.asc | Get a Debian Distribution Key
[**GetApiV4ProjectsIdDebianDistributions**](DebianDistributionApi.md#GetApiV4ProjectsIdDebianDistributions) | **Get** /api/v4/projects/{id}/debian_distributions | Get a list of Debian Distributions
[**GetApiV4ProjectsIdDebianDistributionsCodename**](DebianDistributionApi.md#GetApiV4ProjectsIdDebianDistributionsCodename) | **Get** /api/v4/projects/{id}/debian_distributions/{codename} | Get a Debian Distribution
[**GetApiV4ProjectsIdDebianDistributionsCodenameKeyAsc**](DebianDistributionApi.md#GetApiV4ProjectsIdDebianDistributionsCodenameKeyAsc) | **Get** /api/v4/projects/{id}/debian_distributions/{codename}/key.asc | Get a Debian Distribution Key
[**PostApiV4GroupsIdDebianDistributions**](DebianDistributionApi.md#PostApiV4GroupsIdDebianDistributions) | **Post** /api/v4/groups/{id}/-/debian_distributions | Create a Debian Distribution
[**PostApiV4ProjectsIdDebianDistributions**](DebianDistributionApi.md#PostApiV4ProjectsIdDebianDistributions) | **Post** /api/v4/projects/{id}/debian_distributions | Create a Debian Distribution
[**PutApiV4GroupsIdDebianDistributionsCodename**](DebianDistributionApi.md#PutApiV4GroupsIdDebianDistributionsCodename) | **Put** /api/v4/groups/{id}/-/debian_distributions/{codename} | Update a Debian Distribution
[**PutApiV4ProjectsIdDebianDistributionsCodename**](DebianDistributionApi.md#PutApiV4ProjectsIdDebianDistributionsCodename) | **Put** /api/v4/projects/{id}/debian_distributions/{codename} | Update a Debian Distribution


# **DeleteApiV4GroupsIdDebianDistributionsCodename**
> DeleteApiV4GroupsIdDebianDistributionsCodename(ctx, id, codename, optional)
Delete a Debian Distribution

This feature was introduced in 14.0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the group | 
  **codename** | **string**| The Debian Codename | 
 **optional** | ***DebianDistributionApiDeleteApiV4GroupsIdDebianDistributionsCodenameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DebianDistributionApiDeleteApiV4GroupsIdDebianDistributionsCodenameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **suite** | **optional.String**| The Debian Suite | 
 **origin** | **optional.String**| The Debian Origin | 
 **label** | **optional.String**| The Debian Label | 
 **version** | **optional.String**| The Debian Version | 
 **description** | **optional.String**| The Debian Description | 
 **validTimeDurationSeconds** | **optional.Int32**| The duration before the Release file should be considered expired by the client | 
 **components** | [**optional.Interface of []string**](string.md)| The list of Components | 
 **architectures** | [**optional.Interface of []string**](string.md)| The list of Architectures | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdDebianDistributionsCodename**
> DeleteApiV4ProjectsIdDebianDistributionsCodename(ctx, id, codename, optional)
Delete a Debian Distribution

This feature was introduced in 14.0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **codename** | **string**| The Debian Codename | 
 **optional** | ***DebianDistributionApiDeleteApiV4ProjectsIdDebianDistributionsCodenameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DebianDistributionApiDeleteApiV4ProjectsIdDebianDistributionsCodenameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **suite** | **optional.String**| The Debian Suite | 
 **origin** | **optional.String**| The Debian Origin | 
 **label** | **optional.String**| The Debian Label | 
 **version** | **optional.String**| The Debian Version | 
 **description** | **optional.String**| The Debian Description | 
 **validTimeDurationSeconds** | **optional.Int32**| The duration before the Release file should be considered expired by the client | 
 **components** | [**optional.Interface of []string**](string.md)| The list of Components | 
 **architectures** | [**optional.Interface of []string**](string.md)| The list of Architectures | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdDebianDistributions**
> ApiEntitiesPackagesDebianDistribution GetApiV4GroupsIdDebianDistributions(ctx, id, optional)
Get a list of Debian Distributions

This feature was introduced in 14.0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the group | 
 **optional** | ***DebianDistributionApiGetApiV4GroupsIdDebianDistributionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DebianDistributionApiGetApiV4GroupsIdDebianDistributionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **codename** | **optional.String**| The Debian Codename | 
 **suite** | **optional.String**| The Debian Suite | 
 **origin** | **optional.String**| The Debian Origin | 
 **label** | **optional.String**| The Debian Label | 
 **version** | **optional.String**| The Debian Version | 
 **description** | **optional.String**| The Debian Description | 
 **validTimeDurationSeconds** | **optional.Int32**| The duration before the Release file should be considered expired by the client | 
 **components** | [**optional.Interface of []string**](string.md)| The list of Components | 
 **architectures** | [**optional.Interface of []string**](string.md)| The list of Architectures | 

### Return type

[**ApiEntitiesPackagesDebianDistribution**](API_Entities_Packages_Debian_Distribution.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdDebianDistributionsCodename**
> ApiEntitiesPackagesDebianDistribution GetApiV4GroupsIdDebianDistributionsCodename(ctx, id, codename)
Get a Debian Distribution

This feature was introduced in 14.0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the group | 
  **codename** | **string**| The Debian Codename | 

### Return type

[**ApiEntitiesPackagesDebianDistribution**](API_Entities_Packages_Debian_Distribution.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdDebianDistributionsCodenameKeyAsc**
> ApiEntitiesPackagesDebianDistribution GetApiV4GroupsIdDebianDistributionsCodenameKeyAsc(ctx, id, codename)
Get a Debian Distribution Key

This feature was introduced in 14.4

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the group | 
  **codename** | **string**| The Debian Codename | 

### Return type

[**ApiEntitiesPackagesDebianDistribution**](API_Entities_Packages_Debian_Distribution.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdDebianDistributions**
> ApiEntitiesPackagesDebianDistribution GetApiV4ProjectsIdDebianDistributions(ctx, id, optional)
Get a list of Debian Distributions

This feature was introduced in 14.0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
 **optional** | ***DebianDistributionApiGetApiV4ProjectsIdDebianDistributionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DebianDistributionApiGetApiV4ProjectsIdDebianDistributionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **codename** | **optional.String**| The Debian Codename | 
 **suite** | **optional.String**| The Debian Suite | 
 **origin** | **optional.String**| The Debian Origin | 
 **label** | **optional.String**| The Debian Label | 
 **version** | **optional.String**| The Debian Version | 
 **description** | **optional.String**| The Debian Description | 
 **validTimeDurationSeconds** | **optional.Int32**| The duration before the Release file should be considered expired by the client | 
 **components** | [**optional.Interface of []string**](string.md)| The list of Components | 
 **architectures** | [**optional.Interface of []string**](string.md)| The list of Architectures | 

### Return type

[**ApiEntitiesPackagesDebianDistribution**](API_Entities_Packages_Debian_Distribution.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdDebianDistributionsCodename**
> ApiEntitiesPackagesDebianDistribution GetApiV4ProjectsIdDebianDistributionsCodename(ctx, id, codename)
Get a Debian Distribution

This feature was introduced in 14.0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **codename** | **string**| The Debian Codename | 

### Return type

[**ApiEntitiesPackagesDebianDistribution**](API_Entities_Packages_Debian_Distribution.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdDebianDistributionsCodenameKeyAsc**
> ApiEntitiesPackagesDebianDistribution GetApiV4ProjectsIdDebianDistributionsCodenameKeyAsc(ctx, id, codename)
Get a Debian Distribution Key

This feature was introduced in 14.4

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **codename** | **string**| The Debian Codename | 

### Return type

[**ApiEntitiesPackagesDebianDistribution**](API_Entities_Packages_Debian_Distribution.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4GroupsIdDebianDistributions**
> ApiEntitiesPackagesDebianDistribution PostApiV4GroupsIdDebianDistributions(ctx, id, postApiV4GroupsIdDebianDistributions)
Create a Debian Distribution

This feature was introduced in 14.0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the group | 
  **postApiV4GroupsIdDebianDistributions** | [**PostApiV4GroupsIdDebianDistributions**](PostApiV4GroupsIdDebianDistributions.md)|  | 

### Return type

[**ApiEntitiesPackagesDebianDistribution**](API_Entities_Packages_Debian_Distribution.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdDebianDistributions**
> ApiEntitiesPackagesDebianDistribution PostApiV4ProjectsIdDebianDistributions(ctx, id, postApiV4ProjectsIdDebianDistributions)
Create a Debian Distribution

This feature was introduced in 14.0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **postApiV4ProjectsIdDebianDistributions** | [**PostApiV4ProjectsIdDebianDistributions**](PostApiV4ProjectsIdDebianDistributions.md)|  | 

### Return type

[**ApiEntitiesPackagesDebianDistribution**](API_Entities_Packages_Debian_Distribution.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4GroupsIdDebianDistributionsCodename**
> ApiEntitiesPackagesDebianDistribution PutApiV4GroupsIdDebianDistributionsCodename(ctx, id, codename, putApiV4GroupsIdDebianDistributionsCodename)
Update a Debian Distribution

This feature was introduced in 14.0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the group | 
  **codename** | **string**| The Debian Codename | 
  **putApiV4GroupsIdDebianDistributionsCodename** | [**PutApiV4GroupsIdDebianDistributionsCodename**](PutApiV4GroupsIdDebianDistributionsCodename.md)|  | 

### Return type

[**ApiEntitiesPackagesDebianDistribution**](API_Entities_Packages_Debian_Distribution.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdDebianDistributionsCodename**
> ApiEntitiesPackagesDebianDistribution PutApiV4ProjectsIdDebianDistributionsCodename(ctx, id, codename, putApiV4ProjectsIdDebianDistributionsCodename)
Update a Debian Distribution

This feature was introduced in 14.0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **codename** | **string**| The Debian Codename | 
  **putApiV4ProjectsIdDebianDistributionsCodename** | [**PutApiV4ProjectsIdDebianDistributionsCodename**](PutApiV4ProjectsIdDebianDistributionsCodename.md)|  | 

### Return type

[**ApiEntitiesPackagesDebianDistribution**](API_Entities_Packages_Debian_Distribution.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

