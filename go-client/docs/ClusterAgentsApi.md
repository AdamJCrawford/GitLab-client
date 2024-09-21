# \ClusterAgentsApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4ProjectsIdClusterAgentsAgentId**](ClusterAgentsApi.md#DeleteApiV4ProjectsIdClusterAgentsAgentId) | **Delete** /api/v4/projects/{id}/cluster_agents/{agent_id} | Delete a registered agent
[**DeleteApiV4ProjectsIdClusterAgentsAgentIdTokensTokenId**](ClusterAgentsApi.md#DeleteApiV4ProjectsIdClusterAgentsAgentIdTokensTokenId) | **Delete** /api/v4/projects/{id}/cluster_agents/{agent_id}/tokens/{token_id} | Revoke an agent token
[**GetApiV4ProjectsIdClusterAgents**](ClusterAgentsApi.md#GetApiV4ProjectsIdClusterAgents) | **Get** /api/v4/projects/{id}/cluster_agents | List the agents for a project
[**GetApiV4ProjectsIdClusterAgentsAgentId**](ClusterAgentsApi.md#GetApiV4ProjectsIdClusterAgentsAgentId) | **Get** /api/v4/projects/{id}/cluster_agents/{agent_id} | Get details about an agent
[**GetApiV4ProjectsIdClusterAgentsAgentIdTokens**](ClusterAgentsApi.md#GetApiV4ProjectsIdClusterAgentsAgentIdTokens) | **Get** /api/v4/projects/{id}/cluster_agents/{agent_id}/tokens | List tokens for an agent
[**GetApiV4ProjectsIdClusterAgentsAgentIdTokensTokenId**](ClusterAgentsApi.md#GetApiV4ProjectsIdClusterAgentsAgentIdTokensTokenId) | **Get** /api/v4/projects/{id}/cluster_agents/{agent_id}/tokens/{token_id} | Get a single agent token
[**PostApiV4ProjectsIdClusterAgents**](ClusterAgentsApi.md#PostApiV4ProjectsIdClusterAgents) | **Post** /api/v4/projects/{id}/cluster_agents | Register an agent with a project
[**PostApiV4ProjectsIdClusterAgentsAgentIdTokens**](ClusterAgentsApi.md#PostApiV4ProjectsIdClusterAgentsAgentIdTokens) | **Post** /api/v4/projects/{id}/cluster_agents/{agent_id}/tokens | Create an agent token


# **DeleteApiV4ProjectsIdClusterAgentsAgentId**
> DeleteApiV4ProjectsIdClusterAgentsAgentId(ctx, id, agentId)
Delete a registered agent

This feature was introduced in GitLab 14.10. Deletes an existing agent registration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **agentId** | **int32**| The ID of an agent | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdClusterAgentsAgentIdTokensTokenId**
> DeleteApiV4ProjectsIdClusterAgentsAgentIdTokensTokenId(ctx, id, agentId, tokenId)
Revoke an agent token

This feature was introduced in GitLab 15.0. Revokes an agent token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **agentId** | **int32**| The ID of an agent | 
  **tokenId** | **int32**| The ID of the agent token | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdClusterAgents**
> ApiEntitiesClustersAgent GetApiV4ProjectsIdClusterAgents(ctx, id, optional)
List the agents for a project

This feature was introduced in GitLab 14.10. Returns the list of agents registered for the project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
 **optional** | ***ClusterAgentsApiGetApiV4ProjectsIdClusterAgentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClusterAgentsApiGetApiV4ProjectsIdClusterAgentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**ApiEntitiesClustersAgent**](API_Entities_Clusters_Agent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdClusterAgentsAgentId**
> ApiEntitiesClustersAgent GetApiV4ProjectsIdClusterAgentsAgentId(ctx, id, agentId)
Get details about an agent

This feature was introduced in GitLab 14.10. Gets a single agent details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **agentId** | **int32**| The ID of an agent | 

### Return type

[**ApiEntitiesClustersAgent**](API_Entities_Clusters_Agent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdClusterAgentsAgentIdTokens**
> ApiEntitiesClustersAgentTokenBasic GetApiV4ProjectsIdClusterAgentsAgentIdTokens(ctx, id, agentId, optional)
List tokens for an agent

This feature was introduced in GitLab 15.0. Returns a list of tokens for an agent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **agentId** | **int32**| The ID of an agent | 
 **optional** | ***ClusterAgentsApiGetApiV4ProjectsIdClusterAgentsAgentIdTokensOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClusterAgentsApiGetApiV4ProjectsIdClusterAgentsAgentIdTokensOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**ApiEntitiesClustersAgentTokenBasic**](API_Entities_Clusters_AgentTokenBasic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdClusterAgentsAgentIdTokensTokenId**
> ApiEntitiesClustersAgentToken GetApiV4ProjectsIdClusterAgentsAgentIdTokensTokenId(ctx, id, agentId, tokenId)
Get a single agent token

This feature was introduced in GitLab 15.0. Gets a single agent token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **agentId** | **int32**| The ID of an agent | 
  **tokenId** | **int32**| The ID of the agent token | 

### Return type

[**ApiEntitiesClustersAgentToken**](API_Entities_Clusters_AgentToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdClusterAgents**
> ApiEntitiesClustersAgent PostApiV4ProjectsIdClusterAgents(ctx, id, postApiV4ProjectsIdClusterAgents)
Register an agent with a project

This feature was introduced in GitLab 14.10. Registers an agent to the project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **postApiV4ProjectsIdClusterAgents** | [**PostApiV4ProjectsIdClusterAgents**](PostApiV4ProjectsIdClusterAgents.md)|  | 

### Return type

[**ApiEntitiesClustersAgent**](API_Entities_Clusters_Agent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdClusterAgentsAgentIdTokens**
> ApiEntitiesClustersAgentTokenWithToken PostApiV4ProjectsIdClusterAgentsAgentIdTokens(ctx, id, agentId, postApiV4ProjectsIdClusterAgentsAgentIdTokens)
Create an agent token

This feature was introduced in GitLab 15.0. Creates a new token for an agent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **agentId** | **int32**| The ID of an agent | 
  **postApiV4ProjectsIdClusterAgentsAgentIdTokens** | [**PostApiV4ProjectsIdClusterAgentsAgentIdTokens**](PostApiV4ProjectsIdClusterAgentsAgentIdTokens.md)|  | 

### Return type

[**ApiEntitiesClustersAgentTokenWithToken**](API_Entities_Clusters_AgentTokenWithToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

