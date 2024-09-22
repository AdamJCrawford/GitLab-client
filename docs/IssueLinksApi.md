# \IssueLinksApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4ProjectsIdIssuesIssueIidLinksIssueLinkId**](IssueLinksApi.md#DeleteApiV4ProjectsIdIssuesIssueIidLinksIssueLinkId) | **Delete** /api/v4/projects/{id}/issues/{issue_iid}/links/{issue_link_id} | Delete an issue link
[**GetApiV4ProjectsIdIssuesIssueIidLinks**](IssueLinksApi.md#GetApiV4ProjectsIdIssuesIssueIidLinks) | **Get** /api/v4/projects/{id}/issues/{issue_iid}/links | List issue relations
[**GetApiV4ProjectsIdIssuesIssueIidLinksIssueLinkId**](IssueLinksApi.md#GetApiV4ProjectsIdIssuesIssueIidLinksIssueLinkId) | **Get** /api/v4/projects/{id}/issues/{issue_iid}/links/{issue_link_id} | Get an issue link
[**PostApiV4ProjectsIdIssuesIssueIidLinks**](IssueLinksApi.md#PostApiV4ProjectsIdIssuesIssueIidLinks) | **Post** /api/v4/projects/{id}/issues/{issue_iid}/links | Create an issue link


# **DeleteApiV4ProjectsIdIssuesIssueIidLinksIssueLinkId**
> ApiEntitiesIssueLink DeleteApiV4ProjectsIdIssuesIssueIidLinksIssueLinkId(ctx, id, issueIid, issueLinkId)
Delete an issue link

Deletes an issue link, thus removes the two-way relationship.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **issueIid** | **int32**| The internal ID of a project’s issue | 
  **issueLinkId** | **string**| The ID of an issue relationship | 

### Return type

[**ApiEntitiesIssueLink**](API_Entities_IssueLink.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdIssuesIssueIidLinks**
> []ApiEntitiesRelatedIssue GetApiV4ProjectsIdIssuesIssueIidLinks(ctx, id, issueIid)
List issue relations

Get a list of a given issue’s linked issues, sorted by the relationship creation datetime (ascending).Issues are filtered according to the user authorizations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **issueIid** | **int32**| The internal ID of a project’s issue | 

### Return type

[**[]ApiEntitiesRelatedIssue**](API_Entities_RelatedIssue.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdIssuesIssueIidLinksIssueLinkId**
> ApiEntitiesIssueLink GetApiV4ProjectsIdIssuesIssueIidLinksIssueLinkId(ctx, id, issueIid, issueLinkId)
Get an issue link

Gets details about an issue link. This feature was introduced in GitLab 15.1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **issueIid** | **int32**| The internal ID of a project’s issue | 
  **issueLinkId** | **string**| ID of an issue relationship | 

### Return type

[**ApiEntitiesIssueLink**](API_Entities_IssueLink.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdIssuesIssueIidLinks**
> ApiEntitiesIssueLink PostApiV4ProjectsIdIssuesIssueIidLinks(ctx, id, issueIid, postApiV4ProjectsIdIssuesIssueIidLinks)
Create an issue link

Creates a two-way relation between two issues.The user must be allowed to update both issues to succeed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **issueIid** | **int32**| The internal ID of a project’s issue | 
  **postApiV4ProjectsIdIssuesIssueIidLinks** | [**PostApiV4ProjectsIdIssuesIssueIidLinks**](PostApiV4ProjectsIdIssuesIssueIidLinks.md)|  | 

### Return type

[**ApiEntitiesIssueLink**](API_Entities_IssueLink.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

