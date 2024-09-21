# \AlertManagementApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesMetricImageId**](AlertManagementApi.md#DeleteApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesMetricImageId) | **Delete** /api/v4/projects/{id}/alert_management_alerts/{alert_iid}/metric_images/{metric_image_id} | 
[**GetApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImages**](AlertManagementApi.md#GetApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImages) | **Get** /api/v4/projects/{id}/alert_management_alerts/{alert_iid}/metric_images | 
[**PostApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImages**](AlertManagementApi.md#PostApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImages) | **Post** /api/v4/projects/{id}/alert_management_alerts/{alert_iid}/metric_images | 
[**PostApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesAuthorize**](AlertManagementApi.md#PostApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesAuthorize) | **Post** /api/v4/projects/{id}/alert_management_alerts/{alert_iid}/metric_images/authorize | 
[**PutApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesMetricImageId**](AlertManagementApi.md#PutApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesMetricImageId) | **Put** /api/v4/projects/{id}/alert_management_alerts/{alert_iid}/metric_images/{metric_image_id} | 


# **DeleteApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesMetricImageId**
> ApiEntitiesMetricImage DeleteApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesMetricImageId(ctx, id, alertIid, metricImageId)


Remove a metric image for an alert

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **alertIid** | **int32**| The IID of the Alert | 
  **metricImageId** | **int32**| The ID of metric image | 

### Return type

[**ApiEntitiesMetricImage**](API_Entities_MetricImage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImages**
> []ApiEntitiesMetricImage GetApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImages(ctx, id, alertIid)


Metric Images for alert

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **alertIid** | **int32**| The IID of the Alert | 

### Return type

[**[]ApiEntitiesMetricImage**](API_Entities_MetricImage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImages**
> ApiEntitiesMetricImage PostApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImages(ctx, id, alertIid, file, optional)


Upload a metric image for an alert

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **alertIid** | **int32**| The IID of the Alert | 
  **file** | ***os.File**| The image file to be uploaded | 
 **optional** | ***AlertManagementApiPostApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertManagementApiPostApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **url** | **optional.String**| The url to view more metric info | 
 **urlText** | **optional.String**| A description of the image or URL | 

### Return type

[**ApiEntitiesMetricImage**](API_Entities_MetricImage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesAuthorize**
> PostApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesAuthorize(ctx, id, alertIid)


Workhorse authorize metric image file upload

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **alertIid** | **int32**| The IID of the Alert | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesMetricImageId**
> ApiEntitiesMetricImage PutApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesMetricImageId(ctx, id, alertIid, metricImageId, optional)


Update a metric image for an alert

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **alertIid** | **int32**| The IID of the Alert | 
  **metricImageId** | **int32**| The ID of metric image | 
 **optional** | ***AlertManagementApiPutApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesMetricImageIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertManagementApiPutApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesMetricImageIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **url** | **optional.String**| The url to view more metric info | 
 **urlText** | **optional.String**| A description of the image or URL | 

### Return type

[**ApiEntitiesMetricImage**](API_Entities_MetricImage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

