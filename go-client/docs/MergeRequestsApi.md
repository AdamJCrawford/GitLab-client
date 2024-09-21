# \MergeRequestsApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4ProjectsIdMergeRequestsMergeRequestIid**](MergeRequestsApi.md#DeleteApiV4ProjectsIdMergeRequestsMergeRequestIid) | **Delete** /api/v4/projects/{id}/merge_requests/{merge_request_iid} | Delete a merge request
[**DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommits**](MergeRequestsApi.md#DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommits) | **Delete** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/context_commits | Delete merge request context commits
[**GetApiV4GroupsIdMergeRequests**](MergeRequestsApi.md#GetApiV4GroupsIdMergeRequests) | **Get** /api/v4/groups/{id}/merge_requests | List group merge requests
[**GetApiV4MergeRequests**](MergeRequestsApi.md#GetApiV4MergeRequests) | **Get** /api/v4/merge_requests | List merge requests
[**GetApiV4ProjectsIdMergeRequests**](MergeRequestsApi.md#GetApiV4ProjectsIdMergeRequests) | **Get** /api/v4/projects/{id}/merge_requests | List project merge requests
[**GetApiV4ProjectsIdMergeRequestsMergeRequestIid**](MergeRequestsApi.md#GetApiV4ProjectsIdMergeRequestsMergeRequestIid) | **Get** /api/v4/projects/{id}/merge_requests/{merge_request_iid} | Get single merge request
[**GetApiV4ProjectsIdMergeRequestsMergeRequestIidChanges**](MergeRequestsApi.md#GetApiV4ProjectsIdMergeRequestsMergeRequestIidChanges) | **Get** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/changes | Get single merge request changes
[**GetApiV4ProjectsIdMergeRequestsMergeRequestIidClosesIssues**](MergeRequestsApi.md#GetApiV4ProjectsIdMergeRequestsMergeRequestIidClosesIssues) | **Get** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/closes_issues | List issues that close on merge
[**GetApiV4ProjectsIdMergeRequestsMergeRequestIidCommits**](MergeRequestsApi.md#GetApiV4ProjectsIdMergeRequestsMergeRequestIidCommits) | **Get** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/commits | Get single merge request commits
[**GetApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommits**](MergeRequestsApi.md#GetApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommits) | **Get** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/context_commits | List merge request context commits
[**GetApiV4ProjectsIdMergeRequestsMergeRequestIidDiffs**](MergeRequestsApi.md#GetApiV4ProjectsIdMergeRequestsMergeRequestIidDiffs) | **Get** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/diffs | Get the merge request diffs
[**GetApiV4ProjectsIdMergeRequestsMergeRequestIidMergeRef**](MergeRequestsApi.md#GetApiV4ProjectsIdMergeRequestsMergeRequestIidMergeRef) | **Get** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/merge_ref | Returns the up to date merge-ref HEAD commit
[**GetApiV4ProjectsIdMergeRequestsMergeRequestIidParticipants**](MergeRequestsApi.md#GetApiV4ProjectsIdMergeRequestsMergeRequestIidParticipants) | **Get** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/participants | Get single merge request participants
[**GetApiV4ProjectsIdMergeRequestsMergeRequestIidPipelines**](MergeRequestsApi.md#GetApiV4ProjectsIdMergeRequestsMergeRequestIidPipelines) | **Get** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/pipelines | Get single merge request pipelines
[**GetApiV4ProjectsIdMergeRequestsMergeRequestIidReviewers**](MergeRequestsApi.md#GetApiV4ProjectsIdMergeRequestsMergeRequestIidReviewers) | **Get** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/reviewers | Get single merge request reviewers
[**GetApiV4ProjectsIdMergeRequestsMergeRequestIidTimeStats**](MergeRequestsApi.md#GetApiV4ProjectsIdMergeRequestsMergeRequestIidTimeStats) | **Get** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/time_stats | Get time tracking stats
[**GetApiV4ProjectsIdMergeRequestsMergeRequestIidVersions**](MergeRequestsApi.md#GetApiV4ProjectsIdMergeRequestsMergeRequestIidVersions) | **Get** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/versions | Get a list of merge request diff versions
[**GetApiV4ProjectsIdMergeRequestsMergeRequestIidVersionsVersionId**](MergeRequestsApi.md#GetApiV4ProjectsIdMergeRequestsMergeRequestIidVersionsVersionId) | **Get** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/versions/{version_id} | Get a single merge request diff version
[**PostApiV4ProjectsIdMergeRequests**](MergeRequestsApi.md#PostApiV4ProjectsIdMergeRequests) | **Post** /api/v4/projects/{id}/merge_requests | Create merge request
[**PostApiV4ProjectsIdMergeRequestsMergeRequestIidAddSpentTime**](MergeRequestsApi.md#PostApiV4ProjectsIdMergeRequestsMergeRequestIidAddSpentTime) | **Post** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/add_spent_time | Add spent time for a merge_request
[**PostApiV4ProjectsIdMergeRequestsMergeRequestIidCancelMergeWhenPipelineSucceeds**](MergeRequestsApi.md#PostApiV4ProjectsIdMergeRequestsMergeRequestIidCancelMergeWhenPipelineSucceeds) | **Post** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/cancel_merge_when_pipeline_succeeds | Cancel Merge When Pipeline Succeeds
[**PostApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommits**](MergeRequestsApi.md#PostApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommits) | **Post** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/context_commits | Create merge request context commits
[**PostApiV4ProjectsIdMergeRequestsMergeRequestIidPipelines**](MergeRequestsApi.md#PostApiV4ProjectsIdMergeRequestsMergeRequestIidPipelines) | **Post** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/pipelines | Create merge request pipeline
[**PostApiV4ProjectsIdMergeRequestsMergeRequestIidResetSpentTime**](MergeRequestsApi.md#PostApiV4ProjectsIdMergeRequestsMergeRequestIidResetSpentTime) | **Post** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/reset_spent_time | Reset spent time for a merge_request
[**PostApiV4ProjectsIdMergeRequestsMergeRequestIidResetTimeEstimate**](MergeRequestsApi.md#PostApiV4ProjectsIdMergeRequestsMergeRequestIidResetTimeEstimate) | **Post** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/reset_time_estimate | Reset the time estimate for a project merge_request
[**PostApiV4ProjectsIdMergeRequestsMergeRequestIidTimeEstimate**](MergeRequestsApi.md#PostApiV4ProjectsIdMergeRequestsMergeRequestIidTimeEstimate) | **Post** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/time_estimate | Set a time estimate for a merge_request
[**PutApiV4ProjectsIdMergeRequestsMergeRequestIid**](MergeRequestsApi.md#PutApiV4ProjectsIdMergeRequestsMergeRequestIid) | **Put** /api/v4/projects/{id}/merge_requests/{merge_request_iid} | Update merge request
[**PutApiV4ProjectsIdMergeRequestsMergeRequestIidMerge**](MergeRequestsApi.md#PutApiV4ProjectsIdMergeRequestsMergeRequestIidMerge) | **Put** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/merge | Merge a merge request
[**PutApiV4ProjectsIdMergeRequestsMergeRequestIidRebase**](MergeRequestsApi.md#PutApiV4ProjectsIdMergeRequestsMergeRequestIidRebase) | **Put** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/rebase | Rebase a merge request
[**PutApiV4ProjectsIdMergeRequestsMergeRequestIidResetApprovals**](MergeRequestsApi.md#PutApiV4ProjectsIdMergeRequestsMergeRequestIidResetApprovals) | **Put** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/reset_approvals | Remove all merge request approvals


# **DeleteApiV4ProjectsIdMergeRequestsMergeRequestIid**
> DeleteApiV4ProjectsIdMergeRequestsMergeRequestIid(ctx, id, mergeRequestIid)
Delete a merge request

Only for administrators and project owners. Deletes the merge request in question. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project. | 
  **mergeRequestIid** | **int32**| The internal ID of the merge request. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommits**
> DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommits(ctx, id, commits, mergeRequestIid)
Delete merge request context commits

Delete a list of merge request context commits.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project. | 
  **commits** | [**[]string**](string.md)| The context commits’ SHA. | 
  **mergeRequestIid** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdMergeRequests**
> ApiEntitiesMergeRequestBasic GetApiV4GroupsIdMergeRequests(ctx, id, optional)
List group merge requests

Get all merge requests for this group and its subgroups.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the group owned by the authenticated user. | 
 **optional** | ***MergeRequestsApiGetApiV4GroupsIdMergeRequestsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MergeRequestsApiGetApiV4GroupsIdMergeRequestsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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
 **approverIds** | **optional.String**| Return merge requests which have specified the users with the given IDs as an individual approver | 
 **approvedByIds** | **optional.String**| Return merge requests which have been approved by the specified users with the given IDs | 
 **approvedByUsernames** | **optional.String**| Return merge requests which have been approved by the specified users with the given             usernames | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **nonArchived** | **optional.Bool**| Returns merge requests from non archived projects only. | [default to true]

### Return type

[**ApiEntitiesMergeRequestBasic**](API_Entities_MergeRequestBasic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4MergeRequests**
> ApiEntitiesMergeRequestBasic GetApiV4MergeRequests(ctx, optional)
List merge requests

Get all merge requests the authenticated user has access to. By default it returns only merge requests created by the current user. To get all merge requests, use parameter `scope=all`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MergeRequestsApiGetApiV4MergeRequestsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MergeRequestsApiGetApiV4MergeRequestsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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
 **scope** | **optional.String**| Returns merge requests for the given scope: &#x60;created_by_me&#x60;, &#x60;assigned_to_me&#x60; or &#x60;all&#x60; | [default to created_by_me]
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
 **approverIds** | **optional.String**| Return merge requests which have specified the users with the given IDs as an individual approver | 
 **approvedByIds** | **optional.String**| Return merge requests which have been approved by the specified users with the given IDs | 
 **approvedByUsernames** | **optional.String**| Return merge requests which have been approved by the specified users with the given             usernames | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**ApiEntitiesMergeRequestBasic**](API_Entities_MergeRequestBasic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdMergeRequests**
> ApiEntitiesMergeRequestBasic GetApiV4ProjectsIdMergeRequests(ctx, id, optional)
List project merge requests

Get all merge requests for this project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project. | 
 **optional** | ***MergeRequestsApiGetApiV4ProjectsIdMergeRequestsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MergeRequestsApiGetApiV4ProjectsIdMergeRequestsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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
 **approverIds** | **optional.String**| Return merge requests which have specified the users with the given IDs as an individual approver | 
 **approvedByIds** | **optional.String**| Return merge requests which have been approved by the specified users with the given IDs | 
 **approvedByUsernames** | **optional.String**| Return merge requests which have been approved by the specified users with the given             usernames | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **iids** | [**optional.Interface of []int32**](int32.md)| Returns the request having the given &#x60;iid&#x60;. | 

### Return type

[**ApiEntitiesMergeRequestBasic**](API_Entities_MergeRequestBasic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdMergeRequestsMergeRequestIid**
> ApiEntitiesMergeRequest GetApiV4ProjectsIdMergeRequestsMergeRequestIid(ctx, id, mergeRequestIid, optional)
Get single merge request

Shows information about a single merge request. Note: the `changes_count` value in the response is a string, not an integer. This is because when an merge request has too many changes to display and store, it is capped at 1,000. In that case, the API returns the string `\"1000+\"` for the changes count.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project. | 
  **mergeRequestIid** | **int32**| The internal ID of the merge request. | 
 **optional** | ***MergeRequestsApiGetApiV4ProjectsIdMergeRequestsMergeRequestIidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MergeRequestsApiGetApiV4ProjectsIdMergeRequestsMergeRequestIidOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **renderHtml** | **optional.Bool**| If &#x60;true&#x60;, response includes rendered HTML for title and description. | 
 **includeDivergedCommitsCount** | **optional.Bool**| If &#x60;true&#x60;, response includes the commits behind the target branch. | 
 **includeRebaseInProgress** | **optional.Bool**| If &#x60;true&#x60;, response includes whether a rebase operation is in progress. | 

### Return type

[**ApiEntitiesMergeRequest**](API_Entities_MergeRequest.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdMergeRequestsMergeRequestIidChanges**
> ApiEntitiesMergeRequestChanges GetApiV4ProjectsIdMergeRequestsMergeRequestIidChanges(ctx, id, mergeRequestIid, optional)
Get single merge request changes

Shows information about the merge request including its files and changes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project. | 
  **mergeRequestIid** | **int32**|  | 
 **optional** | ***MergeRequestsApiGetApiV4ProjectsIdMergeRequestsMergeRequestIidChangesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MergeRequestsApiGetApiV4ProjectsIdMergeRequestsMergeRequestIidChangesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **unidiff** | **optional.Bool**| A diff in a Unified diff format | [default to false]

### Return type

[**ApiEntitiesMergeRequestChanges**](API_Entities_MergeRequestChanges.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdMergeRequestsMergeRequestIidClosesIssues**
> ApiEntitiesMrNote GetApiV4ProjectsIdMergeRequestsMergeRequestIidClosesIssues(ctx, id, mergeRequestIid, optional)
List issues that close on merge

Get all the issues that would be closed by merging the provided merge request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project. | 
  **mergeRequestIid** | **int32**|  | 
 **optional** | ***MergeRequestsApiGetApiV4ProjectsIdMergeRequestsMergeRequestIidClosesIssuesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MergeRequestsApiGetApiV4ProjectsIdMergeRequestsMergeRequestIidClosesIssuesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**ApiEntitiesMrNote**](API_Entities_MRNote.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdMergeRequestsMergeRequestIidCommits**
> ApiEntitiesCommit GetApiV4ProjectsIdMergeRequestsMergeRequestIidCommits(ctx, id, mergeRequestIid)
Get single merge request commits

Get a list of merge request commits.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project. | 
  **mergeRequestIid** | **int32**|  | 

### Return type

[**ApiEntitiesCommit**](API_Entities_Commit.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommits**
> ApiEntitiesCommit GetApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommits(ctx, id, mergeRequestIid)
List merge request context commits

Get a list of merge request context commits.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project. | 
  **mergeRequestIid** | **int32**|  | 

### Return type

[**ApiEntitiesCommit**](API_Entities_Commit.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdMergeRequestsMergeRequestIidDiffs**
> ApiEntitiesDiff GetApiV4ProjectsIdMergeRequestsMergeRequestIidDiffs(ctx, id, mergeRequestIid, optional)
Get the merge request diffs

Get a list of merge request diffs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project. | 
  **mergeRequestIid** | **int32**|  | 
 **optional** | ***MergeRequestsApiGetApiV4ProjectsIdMergeRequestsMergeRequestIidDiffsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MergeRequestsApiGetApiV4ProjectsIdMergeRequestsMergeRequestIidDiffsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **unidiff** | **optional.Bool**| A diff in a Unified diff format | [default to false]

### Return type

[**ApiEntitiesDiff**](API_Entities_Diff.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdMergeRequestsMergeRequestIidMergeRef**
> GetApiV4ProjectsIdMergeRequestsMergeRequestIidMergeRef(ctx, id, mergeRequestIid)
Returns the up to date merge-ref HEAD commit

Returns the up to date merge-ref HEAD commit

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project. | 
  **mergeRequestIid** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdMergeRequestsMergeRequestIidParticipants**
> ApiEntitiesUserBasic GetApiV4ProjectsIdMergeRequestsMergeRequestIidParticipants(ctx, id, mergeRequestIid)
Get single merge request participants

Get a list of merge request participants.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project. | 
  **mergeRequestIid** | **int32**|  | 

### Return type

[**ApiEntitiesUserBasic**](API_Entities_UserBasic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdMergeRequestsMergeRequestIidPipelines**
> ApiEntitiesCiPipelineBasic GetApiV4ProjectsIdMergeRequestsMergeRequestIidPipelines(ctx, id, mergeRequestIid)
Get single merge request pipelines

Get a list of merge request pipelines.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project. | 
  **mergeRequestIid** | **int32**|  | 

### Return type

[**ApiEntitiesCiPipelineBasic**](API_Entities_Ci_PipelineBasic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdMergeRequestsMergeRequestIidReviewers**
> ApiEntitiesMergeRequestReviewer GetApiV4ProjectsIdMergeRequestsMergeRequestIidReviewers(ctx, id, mergeRequestIid)
Get single merge request reviewers

Get a list of merge request reviewers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project. | 
  **mergeRequestIid** | **int32**|  | 

### Return type

[**ApiEntitiesMergeRequestReviewer**](API_Entities_MergeRequestReviewer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdMergeRequestsMergeRequestIidTimeStats**
> ApiEntitiesIssuableTimeStats GetApiV4ProjectsIdMergeRequestsMergeRequestIidTimeStats(ctx, id, mergeRequestIid)
Get time tracking stats

Get time tracking stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project. | 
  **mergeRequestIid** | **int32**| The internal ID of the merge_request | 

### Return type

[**ApiEntitiesIssuableTimeStats**](API_Entities_IssuableTimeStats.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdMergeRequestsMergeRequestIidVersions**
> []ApiEntitiesMergeRequestDiff GetApiV4ProjectsIdMergeRequestsMergeRequestIidVersions(ctx, id, mergeRequestIid, optional)
Get a list of merge request diff versions

This feature was introduced in GitLab 8.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **mergeRequestIid** | **int32**| The internal ID of the merge request | 
 **optional** | ***MergeRequestsApiGetApiV4ProjectsIdMergeRequestsMergeRequestIidVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MergeRequestsApiGetApiV4ProjectsIdMergeRequestsMergeRequestIidVersionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesMergeRequestDiff**](API_Entities_MergeRequestDiff.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdMergeRequestsMergeRequestIidVersionsVersionId**
> ApiEntitiesMergeRequestDiffFull GetApiV4ProjectsIdMergeRequestsMergeRequestIidVersionsVersionId(ctx, id, mergeRequestIid, versionId, optional)
Get a single merge request diff version

This feature was introduced in GitLab 8.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **mergeRequestIid** | **int32**| The internal ID of the merge request | 
  **versionId** | **int32**| The ID of the merge request diff version | 
 **optional** | ***MergeRequestsApiGetApiV4ProjectsIdMergeRequestsMergeRequestIidVersionsVersionIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MergeRequestsApiGetApiV4ProjectsIdMergeRequestsMergeRequestIidVersionsVersionIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **unidiff** | **optional.Bool**| A diff in a Unified diff format | [default to false]

### Return type

[**ApiEntitiesMergeRequestDiffFull**](API_Entities_MergeRequestDiffFull.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdMergeRequests**
> ApiEntitiesMergeRequest PostApiV4ProjectsIdMergeRequests(ctx, id, postApiV4ProjectsIdMergeRequests)
Create merge request

Create a new merge request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project. | 
  **postApiV4ProjectsIdMergeRequests** | [**PostApiV4ProjectsIdMergeRequests**](PostApiV4ProjectsIdMergeRequests.md)|  | 

### Return type

[**ApiEntitiesMergeRequest**](API_Entities_MergeRequest.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdMergeRequestsMergeRequestIidAddSpentTime**
> ApiEntitiesIssuableTimeStats PostApiV4ProjectsIdMergeRequestsMergeRequestIidAddSpentTime(ctx, id, mergeRequestIid, postApiV4ProjectsIdMergeRequestsMergeRequestIidAddSpentTime)
Add spent time for a merge_request

Adds spent time for this merge_request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project. | 
  **mergeRequestIid** | **int32**| The internal ID of the merge_request. | 
  **postApiV4ProjectsIdMergeRequestsMergeRequestIidAddSpentTime** | [**PostApiV4ProjectsIdMergeRequestsMergeRequestIidAddSpentTime**](PostApiV4ProjectsIdMergeRequestsMergeRequestIidAddSpentTime.md)|  | 

### Return type

[**ApiEntitiesIssuableTimeStats**](API_Entities_IssuableTimeStats.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdMergeRequestsMergeRequestIidCancelMergeWhenPipelineSucceeds**
> ApiEntitiesMergeRequest PostApiV4ProjectsIdMergeRequestsMergeRequestIidCancelMergeWhenPipelineSucceeds(ctx, id, mergeRequestIid)
Cancel Merge When Pipeline Succeeds

Cancel merge if \"Merge When Pipeline Succeeds\" is enabled

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project. | 
  **mergeRequestIid** | **int32**|  | 

### Return type

[**ApiEntitiesMergeRequest**](API_Entities_MergeRequest.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommits**
> ApiEntitiesCommit PostApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommits(ctx, id, mergeRequestIid, postApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommits)
Create merge request context commits

Create a list of merge request context commits.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project. | 
  **mergeRequestIid** | **int32**|  | 
  **postApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommits** | [**PostApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommits**](PostApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommits.md)|  | 

### Return type

[**ApiEntitiesCommit**](API_Entities_Commit.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdMergeRequestsMergeRequestIidPipelines**
> ApiEntitiesCiPipeline PostApiV4ProjectsIdMergeRequestsMergeRequestIidPipelines(ctx, id, mergeRequestIid)
Create merge request pipeline

Create a new pipeline for a merge request. A pipeline created via this endpoint doesn’t run a regular branch/tag pipeline. It requires `.gitlab-ci.yml` to be configured with `only: [merge_requests]` to create jobs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project. | 
  **mergeRequestIid** | **int32**|  | 

### Return type

[**ApiEntitiesCiPipeline**](API_Entities_Ci_Pipeline.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdMergeRequestsMergeRequestIidResetSpentTime**
> ApiEntitiesIssuableTimeStats PostApiV4ProjectsIdMergeRequestsMergeRequestIidResetSpentTime(ctx, id, mergeRequestIid)
Reset spent time for a merge_request

Resets the total spent time for this merge_request to 0 seconds.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project. | 
  **mergeRequestIid** | **int32**| The internal ID of the merge_request | 

### Return type

[**ApiEntitiesIssuableTimeStats**](API_Entities_IssuableTimeStats.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdMergeRequestsMergeRequestIidResetTimeEstimate**
> ApiEntitiesIssuableTimeStats PostApiV4ProjectsIdMergeRequestsMergeRequestIidResetTimeEstimate(ctx, id, mergeRequestIid)
Reset the time estimate for a project merge_request

Resets the estimated time for this merge_request to 0 seconds.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project. | 
  **mergeRequestIid** | **int32**| The internal ID of the merge_request. | 

### Return type

[**ApiEntitiesIssuableTimeStats**](API_Entities_IssuableTimeStats.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdMergeRequestsMergeRequestIidTimeEstimate**
> ApiEntitiesIssuableTimeStats PostApiV4ProjectsIdMergeRequestsMergeRequestIidTimeEstimate(ctx, id, mergeRequestIid, postApiV4ProjectsIdMergeRequestsMergeRequestIidTimeEstimate)
Set a time estimate for a merge_request

Sets an estimated time of work for this merge_request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project. | 
  **mergeRequestIid** | **int32**| The internal ID of the merge_request. | 
  **postApiV4ProjectsIdMergeRequestsMergeRequestIidTimeEstimate** | [**PostApiV4ProjectsIdMergeRequestsMergeRequestIidTimeEstimate**](PostApiV4ProjectsIdMergeRequestsMergeRequestIidTimeEstimate.md)|  | 

### Return type

[**ApiEntitiesIssuableTimeStats**](API_Entities_IssuableTimeStats.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdMergeRequestsMergeRequestIid**
> ApiEntitiesMergeRequest PutApiV4ProjectsIdMergeRequestsMergeRequestIid(ctx, id, mergeRequestIid, putApiV4ProjectsIdMergeRequestsMergeRequestIid)
Update merge request

Updates an existing merge request. You can change the target branch, title, or even close the merge request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project. | 
  **mergeRequestIid** | **int32**|  | 
  **putApiV4ProjectsIdMergeRequestsMergeRequestIid** | [**PutApiV4ProjectsIdMergeRequestsMergeRequestIid**](PutApiV4ProjectsIdMergeRequestsMergeRequestIid.md)|  | 

### Return type

[**ApiEntitiesMergeRequest**](API_Entities_MergeRequest.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdMergeRequestsMergeRequestIidMerge**
> ApiEntitiesMergeRequest PutApiV4ProjectsIdMergeRequestsMergeRequestIidMerge(ctx, id, mergeRequestIid, putApiV4ProjectsIdMergeRequestsMergeRequestIidMerge)
Merge a merge request

Accept and merge changes submitted with the merge request using this API.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project. | 
  **mergeRequestIid** | **int32**|  | 
  **putApiV4ProjectsIdMergeRequestsMergeRequestIidMerge** | [**PutApiV4ProjectsIdMergeRequestsMergeRequestIidMerge**](PutApiV4ProjectsIdMergeRequestsMergeRequestIidMerge.md)|  | 

### Return type

[**ApiEntitiesMergeRequest**](API_Entities_MergeRequest.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdMergeRequestsMergeRequestIidRebase**
> PutApiV4ProjectsIdMergeRequestsMergeRequestIidRebase(ctx, id, mergeRequestIid, putApiV4ProjectsIdMergeRequestsMergeRequestIidRebase)
Rebase a merge request

Automatically rebase the `source_branch` of the merge request against its `target_branch`. This feature was added in GitLab 11.6

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project. | 
  **mergeRequestIid** | **int32**|  | 
  **putApiV4ProjectsIdMergeRequestsMergeRequestIidRebase** | [**PutApiV4ProjectsIdMergeRequestsMergeRequestIidRebase**](PutApiV4ProjectsIdMergeRequestsMergeRequestIidRebase.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdMergeRequestsMergeRequestIidResetApprovals**
> PutApiV4ProjectsIdMergeRequestsMergeRequestIidResetApprovals(ctx, id, mergeRequestIid)
Remove all merge request approvals

Clear all approvals of merge request. This feature was added in GitLab 15.4

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **mergeRequestIid** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

