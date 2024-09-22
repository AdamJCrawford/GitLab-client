# \ClustersApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4AdminClustersClusterId**](ClustersApi.md#DeleteApiV4AdminClustersClusterId) | **Delete** /api/v4/admin/clusters/{cluster_id} | Delete instance cluster
[**DeleteApiV4GroupsIdClustersClusterId**](ClustersApi.md#DeleteApiV4GroupsIdClustersClusterId) | **Delete** /api/v4/groups/{id}/clusters/{cluster_id} | Delete group cluster
[**DeleteApiV4ProjectsIdClustersClusterId**](ClustersApi.md#DeleteApiV4ProjectsIdClustersClusterId) | **Delete** /api/v4/projects/{id}/clusters/{cluster_id} | Delete project cluster
[**GetApiV4AdminClusters**](ClustersApi.md#GetApiV4AdminClusters) | **Get** /api/v4/admin/clusters | List instance clusters
[**GetApiV4AdminClustersClusterId**](ClustersApi.md#GetApiV4AdminClustersClusterId) | **Get** /api/v4/admin/clusters/{cluster_id} | Get a single instance cluster
[**GetApiV4GroupsIdClusters**](ClustersApi.md#GetApiV4GroupsIdClusters) | **Get** /api/v4/groups/{id}/clusters | List group clusters
[**GetApiV4GroupsIdClustersClusterId**](ClustersApi.md#GetApiV4GroupsIdClustersClusterId) | **Get** /api/v4/groups/{id}/clusters/{cluster_id} | Get a single group cluster
[**GetApiV4ProjectsIdClusters**](ClustersApi.md#GetApiV4ProjectsIdClusters) | **Get** /api/v4/projects/{id}/clusters | List project clusters
[**GetApiV4ProjectsIdClustersClusterId**](ClustersApi.md#GetApiV4ProjectsIdClustersClusterId) | **Get** /api/v4/projects/{id}/clusters/{cluster_id} | Get a single project cluster
[**PostApiV4AdminClustersAdd**](ClustersApi.md#PostApiV4AdminClustersAdd) | **Post** /api/v4/admin/clusters/add | Add existing instance cluster
[**PostApiV4GroupsIdClustersUser**](ClustersApi.md#PostApiV4GroupsIdClustersUser) | **Post** /api/v4/groups/{id}/clusters/user | Add existing cluster to group
[**PostApiV4ProjectsIdClustersUser**](ClustersApi.md#PostApiV4ProjectsIdClustersUser) | **Post** /api/v4/projects/{id}/clusters/user | Add existing cluster to project
[**PutApiV4AdminClustersClusterId**](ClustersApi.md#PutApiV4AdminClustersClusterId) | **Put** /api/v4/admin/clusters/{cluster_id} | Edit instance cluster
[**PutApiV4GroupsIdClustersClusterId**](ClustersApi.md#PutApiV4GroupsIdClustersClusterId) | **Put** /api/v4/groups/{id}/clusters/{cluster_id} | Edit group cluster
[**PutApiV4ProjectsIdClustersClusterId**](ClustersApi.md#PutApiV4ProjectsIdClustersClusterId) | **Put** /api/v4/projects/{id}/clusters/{cluster_id} | Edit project cluster


# **DeleteApiV4AdminClustersClusterId**
> ApiEntitiesCluster DeleteApiV4AdminClustersClusterId(ctx, clusterId)
Delete instance cluster

This feature was introduced in GitLab 13.2. Deletes an existing instance cluster. Does not remove existing resources within the connected Kubernetes cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **int32**| The cluster ID | 

### Return type

[**ApiEntitiesCluster**](API_Entities_Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4GroupsIdClustersClusterId**
> ApiEntitiesClusterGroup DeleteApiV4GroupsIdClustersClusterId(ctx, id, clusterId)
Delete group cluster

This feature was introduced in GitLab 12.1. Deletes an existing group cluster. Does not remove existing resources within the connected Kubernetes cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the group | 
  **clusterId** | **int32**| The Cluster ID | 

### Return type

[**ApiEntitiesClusterGroup**](API_Entities_ClusterGroup.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdClustersClusterId**
> ApiEntitiesClusterProject DeleteApiV4ProjectsIdClustersClusterId(ctx, id, clusterId)
Delete project cluster

This feature was introduced in GitLab 11.7. Deletes an existing project cluster. Does not remove existing resources within the connected Kubernetes cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **clusterId** | **int32**| The Cluster ID | 

### Return type

[**ApiEntitiesClusterProject**](API_Entities_ClusterProject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4AdminClusters**
> []ApiEntitiesCluster GetApiV4AdminClusters(ctx, )
List instance clusters

This feature was introduced in GitLab 13.2. Returns a list of instance clusters.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ApiEntitiesCluster**](API_Entities_Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4AdminClustersClusterId**
> ApiEntitiesCluster GetApiV4AdminClustersClusterId(ctx, clusterId)
Get a single instance cluster

This feature was introduced in GitLab 13.2. Returns a single instance cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **int32**| The cluster ID | 

### Return type

[**ApiEntitiesCluster**](API_Entities_Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdClusters**
> []ApiEntitiesCluster GetApiV4GroupsIdClusters(ctx, id, optional)
List group clusters

This feature was introduced in GitLab 12.1. Returns a list of group clusters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the group | 
 **optional** | ***ClustersApiGetApiV4GroupsIdClustersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiGetApiV4GroupsIdClustersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesCluster**](API_Entities_Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdClustersClusterId**
> ApiEntitiesClusterGroup GetApiV4GroupsIdClustersClusterId(ctx, id, clusterId)
Get a single group cluster

This feature was introduced in GitLab 12.1. Gets a single group cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the group | 
  **clusterId** | **int32**| The cluster ID | 

### Return type

[**ApiEntitiesClusterGroup**](API_Entities_ClusterGroup.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdClusters**
> []ApiEntitiesCluster GetApiV4ProjectsIdClusters(ctx, id, optional)
List project clusters

This feature was introduced in GitLab 11.7. Returns a list of project clusters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
 **optional** | ***ClustersApiGetApiV4ProjectsIdClustersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiGetApiV4ProjectsIdClustersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesCluster**](API_Entities_Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdClustersClusterId**
> ApiEntitiesClusterProject GetApiV4ProjectsIdClustersClusterId(ctx, id, clusterId)
Get a single project cluster

This feature was introduced in GitLab 11.7. Gets a single project cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **clusterId** | **int32**| The cluster ID | 

### Return type

[**ApiEntitiesClusterProject**](API_Entities_ClusterProject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4AdminClustersAdd**
> ApiEntitiesCluster PostApiV4AdminClustersAdd(ctx, postApiV4AdminClustersAdd)
Add existing instance cluster

This feature was introduced in GitLab 13.2. Adds an existing Kubernetes instance cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4AdminClustersAdd** | [**PostApiV4AdminClustersAdd**](PostApiV4AdminClustersAdd.md)|  | 

### Return type

[**ApiEntitiesCluster**](API_Entities_Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4GroupsIdClustersUser**
> ApiEntitiesClusterGroup PostApiV4GroupsIdClustersUser(ctx, id, postApiV4GroupsIdClustersUser)
Add existing cluster to group

This feature was introduced in GitLab 12.1. Adds an existing Kubernetes cluster to the group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the group | 
  **postApiV4GroupsIdClustersUser** | [**PostApiV4GroupsIdClustersUser**](PostApiV4GroupsIdClustersUser.md)|  | 

### Return type

[**ApiEntitiesClusterGroup**](API_Entities_ClusterGroup.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdClustersUser**
> ApiEntitiesClusterProject PostApiV4ProjectsIdClustersUser(ctx, id, postApiV4ProjectsIdClustersUser)
Add existing cluster to project

This feature was introduced in GitLab 11.7. Adds an existing Kubernetes cluster to the project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **postApiV4ProjectsIdClustersUser** | [**PostApiV4ProjectsIdClustersUser**](PostApiV4ProjectsIdClustersUser.md)|  | 

### Return type

[**ApiEntitiesClusterProject**](API_Entities_ClusterProject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4AdminClustersClusterId**
> ApiEntitiesCluster PutApiV4AdminClustersClusterId(ctx, clusterId, putApiV4AdminClustersClusterId)
Edit instance cluster

This feature was introduced in GitLab 13.2. Updates an existing instance cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **int32**| The cluster ID | 
  **putApiV4AdminClustersClusterId** | [**PutApiV4AdminClustersClusterId**](PutApiV4AdminClustersClusterId.md)|  | 

### Return type

[**ApiEntitiesCluster**](API_Entities_Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4GroupsIdClustersClusterId**
> ApiEntitiesClusterGroup PutApiV4GroupsIdClustersClusterId(ctx, id, clusterId, putApiV4GroupsIdClustersClusterId)
Edit group cluster

This feature was introduced in GitLab 12.1. Updates an existing group cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the group | 
  **clusterId** | **int32**| The cluster ID | 
  **putApiV4GroupsIdClustersClusterId** | [**PutApiV4GroupsIdClustersClusterId**](PutApiV4GroupsIdClustersClusterId.md)|  | 

### Return type

[**ApiEntitiesClusterGroup**](API_Entities_ClusterGroup.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdClustersClusterId**
> ApiEntitiesClusterProject PutApiV4ProjectsIdClustersClusterId(ctx, id, clusterId, putApiV4ProjectsIdClustersClusterId)
Edit project cluster

This feature was introduced in GitLab 11.7. Updates an existing project cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **clusterId** | **int32**| The cluster ID | 
  **putApiV4ProjectsIdClustersClusterId** | [**PutApiV4ProjectsIdClustersClusterId**](PutApiV4ProjectsIdClustersClusterId.md)|  | 

### Return type

[**ApiEntitiesClusterProject**](API_Entities_ClusterProject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

