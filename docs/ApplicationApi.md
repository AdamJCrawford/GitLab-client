# \ApplicationApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4ApplicationAppearance**](ApplicationApi.md#GetApiV4ApplicationAppearance) | **Get** /api/v4/application/appearance | 
[**GetApiV4ApplicationStatistics**](ApplicationApi.md#GetApiV4ApplicationStatistics) | **Get** /api/v4/application/statistics | 
[**PutApiV4ApplicationAppearance**](ApplicationApi.md#PutApiV4ApplicationAppearance) | **Put** /api/v4/application/appearance | 


# **GetApiV4ApplicationAppearance**
> ApiEntitiesAppearance GetApiV4ApplicationAppearance(ctx, )


Get the current appearance

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiEntitiesAppearance**](API_Entities_Appearance.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ApplicationStatistics**
> ApiEntitiesApplicationStatistics GetApiV4ApplicationStatistics(ctx, )


Get the current application statistics

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiEntitiesApplicationStatistics**](API_Entities_ApplicationStatistics.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ApplicationAppearance**
> ApiEntitiesAppearance PutApiV4ApplicationAppearance(ctx, optional)


Modify appearance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationApiPutApiV4ApplicationAppearanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiPutApiV4ApplicationAppearanceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **title** | **optional.String**| Instance title on the sign in / sign up page | 
 **description** | **optional.String**| Markdown text shown on the sign in / sign up page | 
 **pwaName** | **optional.String**| Name of the Progressive Web App | 
 **pwaShortName** | **optional.String**| Optional, short name for Progressive Web App | 
 **pwaDescription** | **optional.String**| An explanation of what the Progressive Web App does | 
 **logo** | **optional.Interface of *os.File**| Instance image used on the sign in / sign up page | 
 **pwaIcon** | **optional.Interface of *os.File**| Icon used for Progressive Web App | 
 **headerLogo** | **optional.Interface of *os.File**| Instance image used for the main navigation bar | 
 **favicon** | **optional.Interface of *os.File**| Instance favicon in .ico/.png format | 
 **memberGuidelines** | **optional.String**| Markdown text shown on the members page of a group or project | 
 **newProjectGuidelines** | **optional.String**| Markdown text shown on the new project page | 
 **profileImageGuidelines** | **optional.String**| Markdown text shown on the profile page below Public Avatar | 
 **headerMessage** | **optional.String**| Message within the system header bar | 
 **footerMessage** | **optional.String**| Message within the system footer bar | 
 **messageBackgroundColor** | **optional.String**| Background color for the system header / footer bar | 
 **messageFontColor** | **optional.String**| Font color for the system header / footer bar | 
 **emailHeaderAndFooterEnabled** | **optional.Bool**| Add header and footer to all outgoing emails if enabled | 

### Return type

[**ApiEntitiesAppearance**](API_Entities_Appearance.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

