# \MetadataApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4Metadata**](MetadataApi.md#GetApiV4Metadata) | **Get** /api/v4/metadata | Retrieve metadata information for this GitLab instance
[**GetApiV4Version**](MetadataApi.md#GetApiV4Version) | **Get** /api/v4/version | Retrieves version information for the GitLab instance


# **GetApiV4Metadata**
> ApiEntitiesMetadata GetApiV4Metadata(ctx, )
Retrieve metadata information for this GitLab instance

This feature was introduced in GitLab 15.2.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiEntitiesMetadata**](API_Entities_Metadata.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4Version**
> ApiEntitiesMetadata GetApiV4Version(ctx, )
Retrieves version information for the GitLab instance

This feature was introduced in GitLab 8.13 and deprecated in 15.5. We recommend you instead use the Metadata API.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiEntitiesMetadata**](API_Entities_Metadata.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

