# \GeoApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4GeoProxy**](GeoApi.md#GetApiV4GeoProxy) | **Get** /api/v4/geo/proxy | Determine if a Geo site should proxy requests
[**GetApiV4GeoRepositoriesGlRepositoryPipelineRefs**](GeoApi.md#GetApiV4GeoRepositoriesGlRepositoryPipelineRefs) | **Get** /api/v4/geo/repositories/{gl_repository}/pipeline_refs | Used by secondary runners to verify the secondary instance has the very latest version
[**GetApiV4GeoRetrieveReplicableNameReplicableId**](GeoApi.md#GetApiV4GeoRetrieveReplicableNameReplicableId) | **Get** /api/v4/geo/retrieve/{replicable_name}/{replicable_id} | Internal endpoint that returns a replicable file
[**PostApiV4GeoNodeProxyIdGraphql**](GeoApi.md#PostApiV4GeoNodeProxyIdGraphql) | **Post** /api/v4/geo/node_proxy/{id}/graphql | Query the GraphQL endpoint of an existing Geo node
[**PostApiV4GeoProxyGitSshInfoRefsReceivePack**](GeoApi.md#PostApiV4GeoProxyGitSshInfoRefsReceivePack) | **Post** /api/v4/geo/proxy_git_ssh/info_refs_receive_pack | Internal endpoint that returns git-received-pack output for git push
[**PostApiV4GeoProxyGitSshInfoRefsUploadPack**](GeoApi.md#PostApiV4GeoProxyGitSshInfoRefsUploadPack) | **Post** /api/v4/geo/proxy_git_ssh/info_refs_upload_pack | Internal endpoint that returns info refs upload pack for git clone/pull
[**PostApiV4GeoProxyGitSshReceivePack**](GeoApi.md#PostApiV4GeoProxyGitSshReceivePack) | **Post** /api/v4/geo/proxy_git_ssh/receive_pack | Internal endpoint that posts git-receive-pack for git push
[**PostApiV4GeoProxyGitSshUploadPack**](GeoApi.md#PostApiV4GeoProxyGitSshUploadPack) | **Post** /api/v4/geo/proxy_git_ssh/upload_pack | Internal endpoint that posts git-upload-pack for git clone/pull
[**PostApiV4GeoStatus**](GeoApi.md#PostApiV4GeoStatus) | **Post** /api/v4/geo/status | Internal endpoint that posts the current node status


# **GetApiV4GeoProxy**
> GetApiV4GeoProxy(ctx, )
Determine if a Geo site should proxy requests

Returns a Geo proxy response

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

# **GetApiV4GeoRepositoriesGlRepositoryPipelineRefs**
> []EeApiEntitiesGeoPipelineRefs GetApiV4GeoRepositoriesGlRepositoryPipelineRefs(ctx, glRepository)
Used by secondary runners to verify the secondary instance has the very latest version

Returns the list of pipeline refs for the project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **glRepository** | **string**| The repository to check | 

### Return type

[**[]EeApiEntitiesGeoPipelineRefs**](EE_API_Entities_Geo_PipelineRefs.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GeoRetrieveReplicableNameReplicableId**
> GetApiV4GeoRetrieveReplicableNameReplicableId(ctx, replicableName, replicableId)
Internal endpoint that returns a replicable file

Returns a replicable file from store (via CDN or sendfile)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **replicableName** | **string**| The replicable name of a replicator instance | 
  **replicableId** | **int32**| The replicable ID of a replicable instance | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4GeoNodeProxyIdGraphql**
> PostApiV4GeoNodeProxyIdGraphql(ctx, id)
Query the GraphQL endpoint of an existing Geo node

Query the GraphQL endpoint of an existing Geo node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the Geo node | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4GeoProxyGitSshInfoRefsReceivePack**
> PostApiV4GeoProxyGitSshInfoRefsReceivePack(ctx, postApiV4GeoProxyGitSshInfoRefsReceivePack)
Internal endpoint that returns git-received-pack output for git push

Responsible for making HTTP GET /repo.git/info/refs?service=git-receive-pack                   request from secondary gitlab-shell to primary

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4GeoProxyGitSshInfoRefsReceivePack** | [**PostApiV4GeoProxyGitSshInfoRefsReceivePack**](PostApiV4GeoProxyGitSshInfoRefsReceivePack.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4GeoProxyGitSshInfoRefsUploadPack**
> PostApiV4GeoProxyGitSshInfoRefsUploadPack(ctx, postApiV4GeoProxyGitSshInfoRefsUploadPack)
Internal endpoint that returns info refs upload pack for git clone/pull

Responsible for making HTTP GET /repo.git/info/refs?service=git-upload-pack                   request from secondary gitlab-shell to primary

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4GeoProxyGitSshInfoRefsUploadPack** | [**PostApiV4GeoProxyGitSshInfoRefsUploadPack**](PostApiV4GeoProxyGitSshInfoRefsUploadPack.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4GeoProxyGitSshReceivePack**
> PostApiV4GeoProxyGitSshReceivePack(ctx, postApiV4GeoProxyGitSshReceivePack)
Internal endpoint that posts git-receive-pack for git push

Responsible for making HTTP POST /repo.git/info/refs?service=git-receive-pack                   request from secondary gitlab-shell to primary

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4GeoProxyGitSshReceivePack** | [**PostApiV4GeoProxyGitSshReceivePack**](PostApiV4GeoProxyGitSshReceivePack.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4GeoProxyGitSshUploadPack**
> PostApiV4GeoProxyGitSshUploadPack(ctx, postApiV4GeoProxyGitSshUploadPack)
Internal endpoint that posts git-upload-pack for git clone/pull

Responsible for making HTTP POST /repo.git/git-upload-pack                   request from secondary gitlab-shell to primary

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4GeoProxyGitSshUploadPack** | [**PostApiV4GeoProxyGitSshUploadPack**](PostApiV4GeoProxyGitSshUploadPack.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4GeoStatus**
> EeApiEntitiesGeoNodeStatus PostApiV4GeoStatus(ctx, postApiV4GeoStatus)
Internal endpoint that posts the current node status

Posts the current node status to the primary site

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4GeoStatus** | [**PostApiV4GeoStatus**](PostApiV4GeoStatus.md)|  | 

### Return type

[**EeApiEntitiesGeoNodeStatus**](EE_API_Entities_GeoNodeStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

