# \PlanLimitsApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4ApplicationPlanLimits**](PlanLimitsApi.md#GetApiV4ApplicationPlanLimits) | **Get** /api/v4/application/plan_limits | Get current plan limits
[**PutApiV4ApplicationPlanLimits**](PlanLimitsApi.md#PutApiV4ApplicationPlanLimits) | **Put** /api/v4/application/plan_limits | Change plan limits


# **GetApiV4ApplicationPlanLimits**
> ApiEntitiesPlanLimit GetApiV4ApplicationPlanLimits(ctx, optional)
Get current plan limits

List the current limits of a plan on the GitLab instance.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PlanLimitsApiGetApiV4ApplicationPlanLimitsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlanLimitsApiGetApiV4ApplicationPlanLimitsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **planName** | **optional.String**| Name of the plan to get the limits from. Default: default. | [default to default]

### Return type

[**ApiEntitiesPlanLimit**](API_Entities_PlanLimit.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ApplicationPlanLimits**
> ApiEntitiesPlanLimit PutApiV4ApplicationPlanLimits(ctx, putApiV4ApplicationPlanLimits)
Change plan limits

Modify the limits of a plan on the GitLab instance.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **putApiV4ApplicationPlanLimits** | [**PutApiV4ApplicationPlanLimits**](PutApiV4ApplicationPlanLimits.md)|  | 

### Return type

[**ApiEntitiesPlanLimit**](API_Entities_PlanLimit.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

