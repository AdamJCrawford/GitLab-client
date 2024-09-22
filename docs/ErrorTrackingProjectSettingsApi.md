# \ErrorTrackingProjectSettingsApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4ProjectsIdErrorTrackingSettings**](ErrorTrackingProjectSettingsApi.md#GetApiV4ProjectsIdErrorTrackingSettings) | **Get** /api/v4/projects/{id}/error_tracking/settings | Get Error Tracking settings
[**PatchApiV4ProjectsIdErrorTrackingSettings**](ErrorTrackingProjectSettingsApi.md#PatchApiV4ProjectsIdErrorTrackingSettings) | **Patch** /api/v4/projects/{id}/error_tracking/settings | Enable or disable the Error Tracking project settings
[**PutApiV4ProjectsIdErrorTrackingSettings**](ErrorTrackingProjectSettingsApi.md#PutApiV4ProjectsIdErrorTrackingSettings) | **Put** /api/v4/projects/{id}/error_tracking/settings | Update Error Tracking project settings. Available in GitLab 15.10 and later.


# **GetApiV4ProjectsIdErrorTrackingSettings**
> ApiEntitiesErrorTrackingProjectSetting GetApiV4ProjectsIdErrorTrackingSettings(ctx, id)
Get Error Tracking settings

Get error tracking settings for the project. This feature was introduced in GitLab 12.7.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 

### Return type

[**ApiEntitiesErrorTrackingProjectSetting**](API_Entities_ErrorTracking_ProjectSetting.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchApiV4ProjectsIdErrorTrackingSettings**
> ApiEntitiesErrorTrackingProjectSetting PatchApiV4ProjectsIdErrorTrackingSettings(ctx, id, patchApiV4ProjectsIdErrorTrackingSettings)
Enable or disable the Error Tracking project settings

The API allows you to enable or disable the Error Tracking settings for a project.Only for users with the Maintainer role for the project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **patchApiV4ProjectsIdErrorTrackingSettings** | [**PatchApiV4ProjectsIdErrorTrackingSettings**](PatchApiV4ProjectsIdErrorTrackingSettings.md)|  | 

### Return type

[**ApiEntitiesErrorTrackingProjectSetting**](API_Entities_ErrorTracking_ProjectSetting.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdErrorTrackingSettings**
> ApiEntitiesErrorTrackingProjectSetting PutApiV4ProjectsIdErrorTrackingSettings(ctx, id, putApiV4ProjectsIdErrorTrackingSettings)
Update Error Tracking project settings. Available in GitLab 15.10 and later.

Update Error Tracking settings for a project. Only for users with Maintainer role for the project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **putApiV4ProjectsIdErrorTrackingSettings** | [**PutApiV4ProjectsIdErrorTrackingSettings**](PutApiV4ProjectsIdErrorTrackingSettings.md)|  | 

### Return type

[**ApiEntitiesErrorTrackingProjectSetting**](API_Entities_ErrorTracking_ProjectSetting.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

