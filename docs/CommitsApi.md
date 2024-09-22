# \CommitsApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4ProjectsIdRepositoryCommits**](CommitsApi.md#GetApiV4ProjectsIdRepositoryCommits) | **Get** /api/v4/projects/{id}/repository/commits | 
[**GetApiV4ProjectsIdRepositoryCommitsSha**](CommitsApi.md#GetApiV4ProjectsIdRepositoryCommitsSha) | **Get** /api/v4/projects/{id}/repository/commits/{sha} | 
[**GetApiV4ProjectsIdRepositoryCommitsShaComments**](CommitsApi.md#GetApiV4ProjectsIdRepositoryCommitsShaComments) | **Get** /api/v4/projects/{id}/repository/commits/{sha}/comments | 
[**GetApiV4ProjectsIdRepositoryCommitsShaDiff**](CommitsApi.md#GetApiV4ProjectsIdRepositoryCommitsShaDiff) | **Get** /api/v4/projects/{id}/repository/commits/{sha}/diff | 
[**GetApiV4ProjectsIdRepositoryCommitsShaMergeRequests**](CommitsApi.md#GetApiV4ProjectsIdRepositoryCommitsShaMergeRequests) | **Get** /api/v4/projects/{id}/repository/commits/{sha}/merge_requests | 
[**GetApiV4ProjectsIdRepositoryCommitsShaRefs**](CommitsApi.md#GetApiV4ProjectsIdRepositoryCommitsShaRefs) | **Get** /api/v4/projects/{id}/repository/commits/{sha}/refs | Get all references a commit is pushed to
[**GetApiV4ProjectsIdRepositoryCommitsShaSequence**](CommitsApi.md#GetApiV4ProjectsIdRepositoryCommitsShaSequence) | **Get** /api/v4/projects/{id}/repository/commits/{sha}/sequence | 
[**GetApiV4ProjectsIdRepositoryCommitsShaSignature**](CommitsApi.md#GetApiV4ProjectsIdRepositoryCommitsShaSignature) | **Get** /api/v4/projects/{id}/repository/commits/{sha}/signature | 
[**PostApiV4ProjectsIdRepositoryCommits**](CommitsApi.md#PostApiV4ProjectsIdRepositoryCommits) | **Post** /api/v4/projects/{id}/repository/commits | Commit multiple file changes as one commit
[**PostApiV4ProjectsIdRepositoryCommitsShaCherryPick**](CommitsApi.md#PostApiV4ProjectsIdRepositoryCommitsShaCherryPick) | **Post** /api/v4/projects/{id}/repository/commits/{sha}/cherry_pick | Cherry pick commit into a branch
[**PostApiV4ProjectsIdRepositoryCommitsShaComments**](CommitsApi.md#PostApiV4ProjectsIdRepositoryCommitsShaComments) | **Post** /api/v4/projects/{id}/repository/commits/{sha}/comments | 
[**PostApiV4ProjectsIdRepositoryCommitsShaRevert**](CommitsApi.md#PostApiV4ProjectsIdRepositoryCommitsShaRevert) | **Post** /api/v4/projects/{id}/repository/commits/{sha}/revert | Revert a commit in a branch


# **GetApiV4ProjectsIdRepositoryCommits**
> []ApiEntitiesCommit GetApiV4ProjectsIdRepositoryCommits(ctx, id, optional)


Get a project repository commits

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
 **optional** | ***CommitsApiGetApiV4ProjectsIdRepositoryCommitsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommitsApiGetApiV4ProjectsIdRepositoryCommitsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refName** | **optional.String**| The name of a repository branch or tag, if not given the default branch is used | 
 **since** | **optional.Time**| Only commits after or on this date will be returned | 
 **until** | **optional.Time**| Only commits before or on this date will be returned | 
 **path** | **optional.String**| The file path | 
 **author** | **optional.String**| Search commits by commit author | 
 **all** | **optional.Bool**| Every commit will be returned | 
 **withStats** | **optional.Bool**| Stats about each commit will be added to the response | 
 **firstParent** | **optional.Bool**| Only include the first parent of merges | 
 **order** | **optional.String**| List commits in order | [default to default]
 **trailers** | **optional.Bool**| Parse and include Git trailers for every commit | [default to false]
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesCommit**](API_Entities_Commit.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdRepositoryCommitsSha**
> ApiEntitiesCommitDetail GetApiV4ProjectsIdRepositoryCommitsSha(ctx, id, sha, optional)


Get a specific commit of a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **sha** | **string**| A commit sha, or the name of a branch or tag | 
 **optional** | ***CommitsApiGetApiV4ProjectsIdRepositoryCommitsShaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommitsApiGetApiV4ProjectsIdRepositoryCommitsShaOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **stats** | **optional.Bool**| Include commit stats | [default to true]

### Return type

[**ApiEntitiesCommitDetail**](API_Entities_CommitDetail.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdRepositoryCommitsShaComments**
> []ApiEntitiesCommitNote GetApiV4ProjectsIdRepositoryCommitsShaComments(ctx, id, sha, optional)


Get a commit's comments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **sha** | **string**| A commit sha, or the name of a branch or tag | 
 **optional** | ***CommitsApiGetApiV4ProjectsIdRepositoryCommitsShaCommentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommitsApiGetApiV4ProjectsIdRepositoryCommitsShaCommentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesCommitNote**](API_Entities_CommitNote.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdRepositoryCommitsShaDiff**
> []ApiEntitiesDiff GetApiV4ProjectsIdRepositoryCommitsShaDiff(ctx, id, sha, optional)


Get the diff for a specific commit of a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **sha** | **string**| A commit sha, or the name of a branch or tag | 
 **optional** | ***CommitsApiGetApiV4ProjectsIdRepositoryCommitsShaDiffOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommitsApiGetApiV4ProjectsIdRepositoryCommitsShaDiffOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **unidiff** | **optional.Bool**| A diff in a Unified diff format | [default to false]

### Return type

[**[]ApiEntitiesDiff**](API_Entities_Diff.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdRepositoryCommitsShaMergeRequests**
> []ApiEntitiesMergeRequestBasic GetApiV4ProjectsIdRepositoryCommitsShaMergeRequests(ctx, id, sha, optional)


Get Merge Requests associated with a commit

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **sha** | **string**| A commit sha, or the name of a branch or tag on which to find Merge Requests | 
 **optional** | ***CommitsApiGetApiV4ProjectsIdRepositoryCommitsShaMergeRequestsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommitsApiGetApiV4ProjectsIdRepositoryCommitsShaMergeRequestsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesMergeRequestBasic**](API_Entities_MergeRequestBasic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdRepositoryCommitsShaRefs**
> []ApiEntitiesBasicRef GetApiV4ProjectsIdRepositoryCommitsShaRefs(ctx, id, sha, optional)
Get all references a commit is pushed to

This feature was introduced in GitLab 10.6

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **sha** | **string**| A commit sha | 
 **optional** | ***CommitsApiGetApiV4ProjectsIdRepositoryCommitsShaRefsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommitsApiGetApiV4ProjectsIdRepositoryCommitsShaRefsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | **optional.String**| Scope | [default to all]
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesBasicRef**](API_Entities_BasicRef.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdRepositoryCommitsShaSequence**
> ApiEntitiesCommitSequence GetApiV4ProjectsIdRepositoryCommitsShaSequence(ctx, id, sha, optional)


Get the sequence count of a commit SHA

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **sha** | **string**| A commit SHA | 
 **optional** | ***CommitsApiGetApiV4ProjectsIdRepositoryCommitsShaSequenceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommitsApiGetApiV4ProjectsIdRepositoryCommitsShaSequenceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **firstParent** | **optional.Bool**| Only include the first parent of merges | [default to false]

### Return type

[**ApiEntitiesCommitSequence**](API_Entities_CommitSequence.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdRepositoryCommitsShaSignature**
> ApiEntitiesCommitSignature GetApiV4ProjectsIdRepositoryCommitsShaSignature(ctx, id, sha)


Get a commit's signature

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **sha** | **string**| A commit sha, or the name of a branch or tag | 

### Return type

[**ApiEntitiesCommitSignature**](API_Entities_CommitSignature.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdRepositoryCommits**
> ApiEntitiesCommitDetail PostApiV4ProjectsIdRepositoryCommits(ctx, id, postApiV4ProjectsIdRepositoryCommits)
Commit multiple file changes as one commit

This feature was introduced in GitLab 8.13

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **postApiV4ProjectsIdRepositoryCommits** | [**PostApiV4ProjectsIdRepositoryCommits**](PostApiV4ProjectsIdRepositoryCommits.md)|  | 

### Return type

[**ApiEntitiesCommitDetail**](API_Entities_CommitDetail.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdRepositoryCommitsShaCherryPick**
> ApiEntitiesCommit PostApiV4ProjectsIdRepositoryCommitsShaCherryPick(ctx, id, sha, postApiV4ProjectsIdRepositoryCommitsShaCherryPick)
Cherry pick commit into a branch

This feature was introduced in GitLab 8.15

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **sha** | **string**| A commit sha, or the name of a branch or tag to be cherry picked | 
  **postApiV4ProjectsIdRepositoryCommitsShaCherryPick** | [**PostApiV4ProjectsIdRepositoryCommitsShaCherryPick**](PostApiV4ProjectsIdRepositoryCommitsShaCherryPick.md)|  | 

### Return type

[**ApiEntitiesCommit**](API_Entities_Commit.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdRepositoryCommitsShaComments**
> ApiEntitiesCommitNote PostApiV4ProjectsIdRepositoryCommitsShaComments(ctx, id, sha, postApiV4ProjectsIdRepositoryCommitsShaComments)


Post comment to commit

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **sha** | **string**| A commit sha, or the name of a branch or tag on which to post a comment | 
  **postApiV4ProjectsIdRepositoryCommitsShaComments** | [**PostApiV4ProjectsIdRepositoryCommitsShaComments**](PostApiV4ProjectsIdRepositoryCommitsShaComments.md)|  | 

### Return type

[**ApiEntitiesCommitNote**](API_Entities_CommitNote.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdRepositoryCommitsShaRevert**
> ApiEntitiesCommit PostApiV4ProjectsIdRepositoryCommitsShaRevert(ctx, id, sha, postApiV4ProjectsIdRepositoryCommitsShaRevert)
Revert a commit in a branch

This feature was introduced in GitLab 11.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **sha** | **string**| Commit SHA to revert | 
  **postApiV4ProjectsIdRepositoryCommitsShaRevert** | [**PostApiV4ProjectsIdRepositoryCommitsShaRevert**](PostApiV4ProjectsIdRepositoryCommitsShaRevert.md)|  | 

### Return type

[**ApiEntitiesCommit**](API_Entities_Commit.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

