# \DeploymentsApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4ProjectsIdDeploymentsDeploymentId**](DeploymentsApi.md#DeleteApiV4ProjectsIdDeploymentsDeploymentId) | **Delete** /api/v4/projects/{id}/deployments/{deployment_id} | Delete a specific deployment
[**GetApiV4ProjectsIdDeployments**](DeploymentsApi.md#GetApiV4ProjectsIdDeployments) | **Get** /api/v4/projects/{id}/deployments | List project deployments
[**GetApiV4ProjectsIdDeploymentsDeploymentId**](DeploymentsApi.md#GetApiV4ProjectsIdDeploymentsDeploymentId) | **Get** /api/v4/projects/{id}/deployments/{deployment_id} | Get a specific deployment
[**GetApiV4ProjectsIdDeploymentsDeploymentIdMergeRequests**](DeploymentsApi.md#GetApiV4ProjectsIdDeploymentsDeploymentIdMergeRequests) | **Get** /api/v4/projects/{id}/deployments/{deployment_id}/merge_requests | List of merge requests associated with a deployment
[**PostApiV4ProjectsIdDeployments**](DeploymentsApi.md#PostApiV4ProjectsIdDeployments) | **Post** /api/v4/projects/{id}/deployments | Create a deployment
[**PostApiV4ProjectsIdDeploymentsDeploymentIdApproval**](DeploymentsApi.md#PostApiV4ProjectsIdDeploymentsDeploymentIdApproval) | **Post** /api/v4/projects/{id}/deployments/{deployment_id}/approval | Approve or reject a blocked deployment
[**PutApiV4ProjectsIdDeploymentsDeploymentId**](DeploymentsApi.md#PutApiV4ProjectsIdDeploymentsDeploymentId) | **Put** /api/v4/projects/{id}/deployments/{deployment_id} | Update a deployment


# **DeleteApiV4ProjectsIdDeploymentsDeploymentId**
> DeleteApiV4ProjectsIdDeploymentsDeploymentId(ctx, id, deploymentId)
Delete a specific deployment

Delete a specific deployment that is not currently the last deployment for an environment or in a running state. This feature was introduced in GitLab 15.3.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **deploymentId** | **int32**| The ID of the deployment | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdDeployments**
> []ApiEntitiesDeployment GetApiV4ProjectsIdDeployments(ctx, id, optional)
List project deployments

Get a list of deployments in a project. This feature was introduced in GitLab 8.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
 **optional** | ***DeploymentsApiGetApiV4ProjectsIdDeploymentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeploymentsApiGetApiV4ProjectsIdDeploymentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **orderBy** | **optional.String**| Return deployments ordered by either one of &#x60;id&#x60;, &#x60;iid&#x60;, &#x60;created_at&#x60;, &#x60;updated_at&#x60; or &#x60;ref&#x60; fields. Default is &#x60;id&#x60; | [default to id]
 **sort** | **optional.String**| Return deployments sorted in &#x60;asc&#x60; or &#x60;desc&#x60; order. Default is &#x60;asc&#x60; | [default to asc]
 **updatedAfter** | **optional.Time**| Return deployments updated after the specified date. Expected in ISO 8601 format (&#x60;2019-03-15T08:00:00Z&#x60;) | 
 **updatedBefore** | **optional.Time**| Return deployments updated before the specified date. Expected in ISO 8601 format (&#x60;2019-03-15T08:00:00Z&#x60;) | 
 **finishedAfter** | **optional.Time**| Return deployments finished after the specified date. Expected in ISO 8601 format (&#x60;2019-03-15T08:00:00Z&#x60;) | 
 **finishedBefore** | **optional.Time**| Return deployments finished before the specified date. Expected in ISO 8601 format (&#x60;2019-03-15T08:00:00Z&#x60;) | 
 **environment** | **optional.String**| The name of the environment to filter deployments by | 
 **status** | **optional.String**| The status to filter deployments by. One of &#x60;created&#x60;, &#x60;running&#x60;, &#x60;success&#x60;, &#x60;failed&#x60;, &#x60;canceled&#x60;, or &#x60;blocked&#x60; | 

### Return type

[**[]ApiEntitiesDeployment**](API_Entities_Deployment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdDeploymentsDeploymentId**
> ApiEntitiesDeploymentExtended GetApiV4ProjectsIdDeploymentsDeploymentId(ctx, id, deploymentId)
Get a specific deployment

This feature was introduced in GitLab 8.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **deploymentId** | **int32**| The ID of the deployment | 

### Return type

[**ApiEntitiesDeploymentExtended**](API_Entities_DeploymentExtended.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdDeploymentsDeploymentIdMergeRequests**
> []ApiEntitiesMergeRequestBasic GetApiV4ProjectsIdDeploymentsDeploymentIdMergeRequests(ctx, id, deploymentId, optional)
List of merge requests associated with a deployment

Retrieves the list of merge requests shipped with a given deployment. This feature was introduced in GitLab 12.7.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **deploymentId** | **int32**| The ID of the deployment | 
 **optional** | ***DeploymentsApiGetApiV4ProjectsIdDeploymentsDeploymentIdMergeRequestsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeploymentsApiGetApiV4ProjectsIdDeploymentsDeploymentIdMergeRequestsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **authorId** | **optional.Int32**| Returns merge requests created by the given user &#x60;id&#x60;. Mutually exclusive with &#x60;author_username&#x60;. Combine with &#x60;scope&#x3D;all&#x60; or &#x60;scope&#x3D;assigned_to_me&#x60;. | 
 **authorUsername** | **optional.String**| Returns merge requests created by the given &#x60;username&#x60;. Mutually exclusive with &#x60;author_id&#x60;. | 
 **assigneeId** | **optional.Int32**| Returns merge requests assigned to the given user &#x60;id&#x60;. &#x60;None&#x60; returns unassigned merge requests. &#x60;Any&#x60; returns merge requests with an assignee. | 
 **assigneeUsername** | [**optional.Interface of []string**](string.md)| Returns merge requests created by the given &#x60;username&#x60;. Mutually exclusive with &#x60;author_id&#x60;. | 
 **reviewerUsername** | **optional.String**| Returns merge requests which have the user as a reviewer with the given &#x60;username&#x60;. &#x60;None&#x60; returns merge requests with no reviewers. &#x60;Any&#x60; returns merge requests with any reviewer. Mutually exclusive with &#x60;reviewer_id&#x60;. Introduced in GitLab 13.8. | 
 **labels** | [**optional.Interface of []string**](string.md)| Returns merge requests matching a comma-separated list of labels. &#x60;None&#x60; lists all merge requests with no labels. &#x60;Any&#x60; lists all merge requests with at least one label. Predefined names are case-insensitive. | 
 **milestone** | **optional.String**| Returns merge requests for a specific milestone. &#x60;None&#x60; returns merge requests with no milestone. &#x60;Any&#x60; returns merge requests that have an assigned milestone. | 
 **myReactionEmoji** | **optional.String**| Returns merge requests reacted by the authenticated user by the given &#x60;emoji&#x60;. &#x60;None&#x60; returns issues not given a reaction. &#x60;Any&#x60; returns issues given at least one reaction. | 
 **reviewerId** | **optional.Int32**| Returns merge requests which have the user as a reviewer with the given user &#x60;id&#x60;. &#x60;None&#x60; returns merge requests with no reviewers. &#x60;Any&#x60; returns merge requests with any reviewer. Mutually exclusive with &#x60;reviewer_username&#x60;. | 
 **state** | **optional.String**| Returns &#x60;all&#x60; merge requests or just those that are &#x60;opened&#x60;, &#x60;closed&#x60;, &#x60;locked&#x60;, or &#x60;merged&#x60;. | [default to all]
 **orderBy** | **optional.String**| Returns merge requests ordered by &#x60;created_at&#x60;, &#x60;label_priority&#x60;, &#x60;milestone_due&#x60;, &#x60;popularity&#x60;, &#x60;priority&#x60;, &#x60;title&#x60; or &#x60;updated_at&#x60; fields. Introduced in GitLab 14.8. | [default to created_at]
 **sort** | **optional.String**| Returns merge requests sorted in &#x60;asc&#x60; or &#x60;desc&#x60; order. | [default to desc]
 **withLabelsDetails** | **optional.Bool**| If &#x60;true&#x60;, response returns more details for each label in labels field: &#x60;:name&#x60;,&#x60;:color&#x60;, &#x60;:description&#x60;, &#x60;:description_html&#x60;, &#x60;:text_color&#x60; | [default to false]
 **withMergeStatusRecheck** | **optional.Bool**| If &#x60;true&#x60;, this projection requests (but does not guarantee) that the &#x60;merge_status&#x60; field be recalculated asynchronously. Introduced in GitLab 13.0. | [default to false]
 **createdAfter** | **optional.Time**| Returns merge requests created on or after the given time. Expected in ISO 8601 format. | 
 **createdBefore** | **optional.Time**| Returns merge requests created on or before the given time. Expected in ISO 8601 format. | 
 **updatedAfter** | **optional.Time**| Returns merge requests updated on or after the given time. Expected in ISO 8601 format. | 
 **updatedBefore** | **optional.Time**| Returns merge requests updated on or before the given time. Expected in ISO 8601 format. | 
 **view** | **optional.String**| If simple, returns the &#x60;iid&#x60;, URL, title, description, and basic state of merge request | 
 **scope** | **optional.String**| Returns merge requests for the given scope: &#x60;created_by_me&#x60;, &#x60;assigned_to_me&#x60; or &#x60;all&#x60; | 
 **sourceBranch** | **optional.String**| Returns merge requests with the given source branch | 
 **sourceProjectId** | **optional.Int32**| Returns merge requests with the given source project id | 
 **targetBranch** | **optional.String**| Returns merge requests with the given target branch | 
 **search** | **optional.String**| Search merge requests against their &#x60;title&#x60; and &#x60;description&#x60;. | 
 **in** | **optional.String**| Modify the scope of the search attribute. &#x60;title&#x60;, &#x60;description&#x60;, or a string joining them with comma. | 
 **wip** | **optional.String**| Filter merge requests against their &#x60;wip&#x60; status. &#x60;yes&#x60; to return only draft merge requests, &#x60;no&#x60; to return non-draft merge requests. | 
 **notAuthorId** | **optional.Int32**| &#x60;&lt;Negated&gt;&#x60; Returns merge requests created by the given user &#x60;id&#x60;. Mutually exclusive with &#x60;author_username&#x60;. Combine with &#x60;scope&#x3D;all&#x60; or &#x60;scope&#x3D;assigned_to_me&#x60;. | 
 **notAuthorUsername** | **optional.String**| &#x60;&lt;Negated&gt;&#x60; Returns merge requests created by the given &#x60;username&#x60;. Mutually exclusive with &#x60;author_id&#x60;. | 
 **notAssigneeId** | **optional.Int32**| &#x60;&lt;Negated&gt;&#x60; Returns merge requests assigned to the given user &#x60;id&#x60;. &#x60;None&#x60; returns unassigned merge requests. &#x60;Any&#x60; returns merge requests with an assignee. | 
 **notAssigneeUsername** | [**optional.Interface of []string**](string.md)| &#x60;&lt;Negated&gt;&#x60; Returns merge requests created by the given &#x60;username&#x60;. Mutually exclusive with &#x60;author_id&#x60;. | 
 **notReviewerUsername** | **optional.String**| &#x60;&lt;Negated&gt;&#x60; Returns merge requests which have the user as a reviewer with the given &#x60;username&#x60;. &#x60;None&#x60; returns merge requests with no reviewers. &#x60;Any&#x60; returns merge requests with any reviewer. Mutually exclusive with &#x60;reviewer_id&#x60;. Introduced in GitLab 13.8. | 
 **notLabels** | [**optional.Interface of []string**](string.md)| &#x60;&lt;Negated&gt;&#x60; Returns merge requests matching a comma-separated list of labels. &#x60;None&#x60; lists all merge requests with no labels. &#x60;Any&#x60; lists all merge requests with at least one label. Predefined names are case-insensitive. | 
 **notMilestone** | **optional.String**| &#x60;&lt;Negated&gt;&#x60; Returns merge requests for a specific milestone. &#x60;None&#x60; returns merge requests with no milestone. &#x60;Any&#x60; returns merge requests that have an assigned milestone. | 
 **notMyReactionEmoji** | **optional.String**| &#x60;&lt;Negated&gt;&#x60; Returns merge requests reacted by the authenticated user by the given &#x60;emoji&#x60;. &#x60;None&#x60; returns issues not given a reaction. &#x60;Any&#x60; returns issues given at least one reaction. | 
 **notReviewerId** | **optional.Int32**| &#x60;&lt;Negated&gt;&#x60; Returns merge requests which have the user as a reviewer with the given user &#x60;id&#x60;. &#x60;None&#x60; returns merge requests with no reviewers. &#x60;Any&#x60; returns merge requests with any reviewer. Mutually exclusive with &#x60;reviewer_username&#x60;. | 
 **deployedBefore** | **optional.String**| Returns merge requests deployed before the given date/time. Expected in ISO 8601 format. | 
 **deployedAfter** | **optional.String**| Returns merge requests deployed after the given date/time. Expected in ISO 8601 format | 
 **environment** | **optional.String**| Returns merge requests deployed to the given environment | 
 **approved** | **optional.String**| Filters merge requests by their &#x60;approved&#x60; status. &#x60;yes&#x60; returns only approved merge requests. &#x60;no&#x60; returns only non-approved merge requests. | 
 **mergeUserId** | **optional.Int32**| Returns merge requests which have been merged by the user with the given user &#x60;id&#x60;. Mutually exclusive with &#x60;merge_user_username&#x60;. | 
 **mergeUserUsername** | **optional.String**| Returns merge requests which have been merged by the user with the given &#x60;username&#x60;. Mutually exclusive with &#x60;merge_user_id&#x60;. | 

### Return type

[**[]ApiEntitiesMergeRequestBasic**](API_Entities_MergeRequestBasic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdDeployments**
> ApiEntitiesDeploymentExtended PostApiV4ProjectsIdDeployments(ctx, id, postApiV4ProjectsIdDeployments)
Create a deployment

This feature was introduced in GitLab 12.4.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **postApiV4ProjectsIdDeployments** | [**PostApiV4ProjectsIdDeployments**](PostApiV4ProjectsIdDeployments.md)|  | 

### Return type

[**ApiEntitiesDeploymentExtended**](API_Entities_DeploymentExtended.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdDeploymentsDeploymentIdApproval**
> ApiEntitiesDeploymentsApproval PostApiV4ProjectsIdDeploymentsDeploymentIdApproval(ctx, id, deploymentId, postApiV4ProjectsIdDeploymentsDeploymentIdApproval)
Approve or reject a blocked deployment

This feature was introduced in GitLab 14.8.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **deploymentId** | **int32**| The ID of the deployment | 
  **postApiV4ProjectsIdDeploymentsDeploymentIdApproval** | [**PostApiV4ProjectsIdDeploymentsDeploymentIdApproval**](PostApiV4ProjectsIdDeploymentsDeploymentIdApproval.md)|  | 

### Return type

[**ApiEntitiesDeploymentsApproval**](API_Entities_Deployments_Approval.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdDeploymentsDeploymentId**
> ApiEntitiesDeploymentExtended PutApiV4ProjectsIdDeploymentsDeploymentId(ctx, id, deploymentId, putApiV4ProjectsIdDeploymentsDeploymentId)
Update a deployment

This feature was introduced in GitLab 12.4.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **deploymentId** | **int32**|  | 
  **putApiV4ProjectsIdDeploymentsDeploymentId** | [**PutApiV4ProjectsIdDeploymentsDeploymentId**](PutApiV4ProjectsIdDeploymentsDeploymentId.md)|  | 

### Return type

[**ApiEntitiesDeploymentExtended**](API_Entities_DeploymentExtended.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

