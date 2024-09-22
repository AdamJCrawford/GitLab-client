# \JobApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4Job**](JobApi.md#GetApiV4Job) | **Get** /api/v4/job | 
[**GetApiV4JobAllowedAgents**](JobApi.md#GetApiV4JobAllowedAgents) | **Get** /api/v4/job/allowed_agents | Get current agents


# **GetApiV4Job**
> ApiEntitiesCiJob GetApiV4Job(ctx, )


Get current job using job token

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiEntitiesCiJob**](API_Entities_Ci_Job.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4JobAllowedAgents**
> ApiEntitiesCiJob GetApiV4JobAllowedAgents(ctx, )
Get current agents

Retrieves a list of agents for the given job token

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiEntitiesCiJob**](API_Entities_Ci_Job.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

