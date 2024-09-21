# \ProjectImportApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4ProjectsIdImport**](ProjectImportApi.md#GetApiV4ProjectsIdImport) | **Get** /api/v4/projects/{id}/import | Get a project import status
[**GetApiV4ProjectsIdRelationImports**](ProjectImportApi.md#GetApiV4ProjectsIdRelationImports) | **Get** /api/v4/projects/{id}/relation-imports | Get the statuses of relation imports for specified project
[**PostApiV4ProjectsImport**](ProjectImportApi.md#PostApiV4ProjectsImport) | **Post** /api/v4/projects/import | Create a new project import
[**PostApiV4ProjectsImportAuthorize**](ProjectImportApi.md#PostApiV4ProjectsImportAuthorize) | **Post** /api/v4/projects/import/authorize | Workhorse authorize the project import upload
[**PostApiV4ProjectsImportRelation**](ProjectImportApi.md#PostApiV4ProjectsImportRelation) | **Post** /api/v4/projects/import-relation | Re-import a relation into a project
[**PostApiV4ProjectsImportRelationAuthorize**](ProjectImportApi.md#PostApiV4ProjectsImportRelationAuthorize) | **Post** /api/v4/projects/import-relation/authorize | Workhorse authorize the project relation import upload
[**PostApiV4ProjectsRemoteImport**](ProjectImportApi.md#PostApiV4ProjectsRemoteImport) | **Post** /api/v4/projects/remote-import | Create a new project import using a remote object storage path
[**PostApiV4ProjectsRemoteImportS3**](ProjectImportApi.md#PostApiV4ProjectsRemoteImportS3) | **Post** /api/v4/projects/remote-import-s3 | Create a new project import using a file from AWS S3


# **GetApiV4ProjectsIdImport**
> ApiEntitiesProjectImportStatus GetApiV4ProjectsIdImport(ctx, id)
Get a project import status

This feature was introduced in GitLab 10.6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 

### Return type

[**ApiEntitiesProjectImportStatus**](API_Entities_ProjectImportStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdRelationImports**
> ApiEntitiesProjectImportStatus GetApiV4ProjectsIdRelationImports(ctx, id)
Get the statuses of relation imports for specified project

This feature was introduced in GitLab 16.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 

### Return type

[**ApiEntitiesProjectImportStatus**](API_Entities_ProjectImportStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsImport**
> ApiEntitiesProjectImportStatus PostApiV4ProjectsImport(ctx, path, file, optional)
Create a new project import

This feature was introduced in GitLab 10.6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| The new project path and name | 
  **file** | ***os.File**| The project export file to be imported | 
 **optional** | ***ProjectImportApiPostApiV4ProjectsImportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectImportApiPostApiV4ProjectsImportOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **optional.String**| The name of the project to be imported. Defaults to the path of the project if not provided. | 
 **namespace** | **optional.String**| The ID or name of the namespace that the project will be imported into. Defaults to the current user&#39;s namespace. | 
 **overwrite** | **optional.Bool**| If there is a project in the same namespace and with the same name overwrite it | [default to false]
 **overrideParamsDescription** | **optional.String**| The description of the project | 
 **overrideParamsBuildGitStrategy** | **optional.String**| The Git strategy. Defaults to &#x60;fetch&#x60; | 
 **overrideParamsBuildTimeout** | **optional.Int32**| Build timeout | 
 **overrideParamsAutoCancelPendingPipelines** | **optional.String**| Auto-cancel pending pipelines | 
 **overrideParamsCiConfigPath** | **optional.String**| The path to CI config file. Defaults to &#x60;.gitlab-ci.yml&#x60; | 
 **overrideParamsServiceDeskEnabled** | **optional.Bool**| Disable or enable the service desk | 
 **overrideParamsIssuesEnabled** | **optional.Bool**| Flag indication if the issue tracker is enabled | 
 **overrideParamsMergeRequestsEnabled** | **optional.Bool**| Flag indication if merge requests are enabled | 
 **overrideParamsWikiEnabled** | **optional.Bool**| Flag indication if the wiki is enabled | 
 **overrideParamsJobsEnabled** | **optional.Bool**| Flag indication if jobs are enabled | 
 **overrideParamsSnippetsEnabled** | **optional.Bool**| Flag indication if snippets are enabled | 
 **overrideParamsIssuesAccessLevel** | **optional.String**| Issues access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsRepositoryAccessLevel** | **optional.String**| Repository access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsMergeRequestsAccessLevel** | **optional.String**| Merge requests access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsForkingAccessLevel** | **optional.String**| Forks access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsWikiAccessLevel** | **optional.String**| Wiki access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsBuildsAccessLevel** | **optional.String**| Builds access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsSnippetsAccessLevel** | **optional.String**| Snippets access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsPagesAccessLevel** | **optional.String**| Pages access level. One of &#x60;disabled&#x60;, &#x60;private&#x60;, &#x60;enabled&#x60; or &#x60;public&#x60; | 
 **overrideParamsAnalyticsAccessLevel** | **optional.String**| Analytics access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsContainerRegistryAccessLevel** | **optional.String**| Controls visibility of the container registry. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60;. &#x60;private&#x60; will make the container registry accessible only to project members (reporter role and above). &#x60;enabled&#x60; will make the container registry accessible to everyone who has access to the project. &#x60;disabled&#x60; will disable the container registry | 
 **overrideParamsSecurityAndComplianceAccessLevel** | **optional.String**| Security and compliance access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsReleasesAccessLevel** | **optional.String**| Releases access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsEnvironmentsAccessLevel** | **optional.String**| Environments access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsFeatureFlagsAccessLevel** | **optional.String**| Feature flags access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsInfrastructureAccessLevel** | **optional.String**| Infrastructure access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsMonitorAccessLevel** | **optional.String**| Monitor access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsModelExperimentsAccessLevel** | **optional.String**| Model experiments access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsModelRegistryAccessLevel** | **optional.String**| Model registry access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsEmailsDisabled** | **optional.Bool**| Deprecated: Use emails_enabled instead. | 
 **overrideParamsEmailsEnabled** | **optional.Bool**| Enable email notifications | 
 **overrideParamsShowDefaultAwardEmojis** | **optional.Bool**| Show default award emojis | 
 **overrideParamsShowDiffPreviewInEmail** | **optional.Bool**| Include the code diff preview in merge request notification emails | 
 **overrideParamsWarnAboutPotentiallyUnwantedCharacters** | **optional.Bool**| Warn about potentially unwanted characters | 
 **overrideParamsEnforceAuthChecksOnUploads** | **optional.Bool**| Enforce auth check on uploads | 
 **overrideParamsSharedRunnersEnabled** | **optional.Bool**| Flag indication if shared runners are enabled for that project | 
 **overrideParamsGroupRunnersEnabled** | **optional.Bool**| Flag indication if group runners are enabled for that project | 
 **overrideParamsResolveOutdatedDiffDiscussions** | **optional.Bool**| Automatically resolve merge request diff threads on lines changed with a push | 
 **overrideParamsRemoveSourceBranchAfterMerge** | **optional.Bool**| Remove the source branch by default after merge | 
 **overrideParamsContainerRegistryEnabled** | **optional.Bool**| Deprecated: Use :container_registry_access_level instead. Flag indication if the container registry is enabled for that project | 
 **overrideParamsContainerExpirationPolicyAttributesCadence** | **optional.String**| Container expiration policy cadence for recurring job | 
 **overrideParamsContainerExpirationPolicyAttributesKeepN** | **optional.Int32**| Container expiration policy number of images to keep | 
 **overrideParamsContainerExpirationPolicyAttributesOlderThan** | **optional.String**| Container expiration policy remove images older than value | 
 **overrideParamsContainerExpirationPolicyAttributesNameRegex** | **optional.String**| Container expiration policy regex for image removal | 
 **overrideParamsContainerExpirationPolicyAttributesNameRegexKeep** | **optional.String**| Container expiration policy regex for image retention | 
 **overrideParamsContainerExpirationPolicyAttributesEnabled** | **optional.Bool**| Flag indication if container expiration policy is enabled | 
 **overrideParamsLfsEnabled** | **optional.Bool**| Flag indication if Git LFS is enabled for that project | 
 **overrideParamsVisibility** | **optional.String**| The visibility of the project. | 
 **overrideParamsPublicBuilds** | **optional.Bool**| Deprecated: Use public_jobs instead. | 
 **overrideParamsPublicJobs** | **optional.Bool**| Perform public builds | 
 **overrideParamsRequestAccessEnabled** | **optional.Bool**| Allow users to request member access | 
 **overrideParamsOnlyAllowMergeIfPipelineSucceeds** | **optional.Bool**| Only allow to merge if builds succeed | 
 **overrideParamsAllowMergeOnSkippedPipeline** | **optional.Bool**| Allow to merge if pipeline is skipped | 
 **overrideParamsOnlyAllowMergeIfAllDiscussionsAreResolved** | **optional.Bool**| Only allow to merge if all threads are resolved | 
 **overrideParamsTagList** | [**optional.Interface of []string**](string.md)| Deprecated: Use :topics instead | 
 **overrideParamsTopics** | [**optional.Interface of []string**](string.md)| The list of topics for a project | 
 **overrideParamsAvatar** | **optional.Interface of *os.File**| Avatar image for project | 
 **overrideParamsPrintingMergeRequestLinkEnabled** | **optional.Bool**| Show link to create/view merge request when pushing from the command line | 
 **overrideParamsMergeMethod** | **optional.String**| The merge method used when merging merge requests | 
 **overrideParamsSuggestionCommitMessage** | **optional.String**| The commit message used to apply merge request suggestions | 
 **overrideParamsMergeCommitTemplate** | **optional.String**| Template used to create merge commit message | 
 **overrideParamsSquashCommitTemplate** | **optional.String**| Template used to create squash commit message | 
 **overrideParamsIssueBranchTemplate** | **optional.String**| Template used to create a branch from an issue | 
 **overrideParamsInitializeWithReadme** | **optional.Bool**| Initialize a project with a README.md | 
 **overrideParamsAutoDevopsEnabled** | **optional.Bool**| Flag indication if Auto DevOps is enabled | 
 **overrideParamsAutoDevopsDeployStrategy** | **optional.String**| Auto Deploy strategy | 
 **overrideParamsAutocloseReferencedIssues** | **optional.Bool**| Flag indication if referenced issues auto-closing is enabled | 
 **overrideParamsRepositoryStorage** | **optional.String**| Which storage shard the repository is on. Available only to admins | 
 **overrideParamsPackagesEnabled** | **optional.Bool**| Enable project packages feature | 
 **overrideParamsSquashOption** | **optional.String**| Squash default for project. One of &#x60;never&#x60;, &#x60;always&#x60;, &#x60;default_on&#x60;, or &#x60;default_off&#x60;. | 
 **overrideParamsMrDefaultTargetSelf** | **optional.Bool**| Merge requests of this forked project targets itself by default | 
 **overrideParamsOnlyAllowMergeIfAllStatusChecksPassed** | **optional.Bool**| Blocks merge requests from merging unless all status checks have passed | 
 **overrideParamsApprovalsBeforeMerge** | **optional.Int32**| How many approvers should approve merge request by default | 
 **overrideParamsMirror** | **optional.Bool**| Enables pull mirroring in a project | 
 **overrideParamsMirrorTriggerBuilds** | **optional.Bool**| Pull mirroring triggers builds | 
 **overrideParamsExternalAuthorizationClassificationLabel** | **optional.String**| The classification label for the project | 
 **overrideParamsRequirementsAccessLevel** | **optional.String**| Requirements feature access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsPreventMergeWithoutJiraIssue** | **optional.Bool**| Require an associated issue from Jira | 
 **filePath** | **optional.String**| Path to locally stored body (generated by Workhorse) | 
 **fileName** | **optional.String**| Real filename as send in Content-Disposition (generated by Workhorse) | 
 **fileType** | **optional.String**| Real content type as send in Content-Type (generated by Workhorse) | 
 **fileSize** | **optional.Int32**| Real size of file (generated by Workhorse) | 
 **fileMd5** | **optional.String**| MD5 checksum of the file (generated by Workhorse) | 
 **fileSha1** | **optional.String**| SHA1 checksum of the file (generated by Workhorse) | 
 **fileSha256** | **optional.String**| SHA256 checksum of the file (generated by Workhorse) | 
 **fileEtag** | **optional.String**| Etag of the file (generated by Workhorse) | 
 **fileRemoteId** | **optional.String**| Remote_id of the file (generated by Workhorse) | 
 **fileRemoteUrl** | **optional.String**| Remote_url of the file (generated by Workhorse) | 

### Return type

[**ApiEntitiesProjectImportStatus**](API_Entities_ProjectImportStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsImportAuthorize**
> PostApiV4ProjectsImportAuthorize(ctx, )
Workhorse authorize the project import upload

This feature was introduced in GitLab 12.9

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsImportRelation**
> ApiEntitiesRelationImportTracker PostApiV4ProjectsImportRelation(ctx, path, file, relation, optional)
Re-import a relation into a project

This feature was introduced in GitLab 16.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| The project path and name | 
  **file** | ***os.File**| The project export file from which to extract the relation. | 
  **relation** | **string**| The relation to import. Must be one of issues, merge_requests, ci_pipelines, or milestones. | 
 **optional** | ***ProjectImportApiPostApiV4ProjectsImportRelationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectImportApiPostApiV4ProjectsImportRelationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filePath** | **optional.String**| Path to locally stored body (generated by Workhorse) | 
 **fileName** | **optional.String**| Real filename as sent in Content-Disposition (generated by Workhorse) | 
 **fileType** | **optional.String**| Real content type as send in Content-Type (generated by Workhorse) | 
 **fileSize** | **optional.Int32**| Real size of file (generated by Workhorse) | 
 **fileMd5** | **optional.String**| MD5 checksum of the file (generated by Workhorse) | 
 **fileSha1** | **optional.String**| SHA1 checksum of the file (generated by Workhorse) | 
 **fileSha256** | **optional.String**| SHA256 checksum of the file (generated by Workhorse) | 
 **fileEtag** | **optional.String**| Etag of the file (generated by Workhorse) | 
 **fileRemoteId** | **optional.String**| Remote_id of the file (generated by Workhorse) | 
 **fileRemoteUrl** | **optional.String**| Remote_url of the file (generated by Workhorse) | 

### Return type

[**ApiEntitiesRelationImportTracker**](API_Entities_RelationImportTracker.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsImportRelationAuthorize**
> PostApiV4ProjectsImportRelationAuthorize(ctx, )
Workhorse authorize the project relation import upload

This feature was introduced in GitLab 16.11

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsRemoteImport**
> ApiEntitiesProjectImportStatus PostApiV4ProjectsRemoteImport(ctx, url, path, optional)
Create a new project import using a remote object storage path

This feature was introduced in GitLab 13.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **url** | **string**| The URL for the file. | 
  **path** | **string**| The new project path and name | 
 **optional** | ***ProjectImportApiPostApiV4ProjectsRemoteImportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectImportApiPostApiV4ProjectsRemoteImportOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **optional.String**| The name of the project to be imported. Defaults to the path of the project if not provided. | 
 **namespace** | **optional.String**| The ID or name of the namespace that the project will be imported into. Defaults to the current user&#39;s namespace. | 
 **overwrite** | **optional.Bool**| If there is a project in the same namespace and with the same name overwrite it | [default to false]
 **overrideParamsDescription** | **optional.String**| The description of the project | 
 **overrideParamsBuildGitStrategy** | **optional.String**| The Git strategy. Defaults to &#x60;fetch&#x60; | 
 **overrideParamsBuildTimeout** | **optional.Int32**| Build timeout | 
 **overrideParamsAutoCancelPendingPipelines** | **optional.String**| Auto-cancel pending pipelines | 
 **overrideParamsCiConfigPath** | **optional.String**| The path to CI config file. Defaults to &#x60;.gitlab-ci.yml&#x60; | 
 **overrideParamsServiceDeskEnabled** | **optional.Bool**| Disable or enable the service desk | 
 **overrideParamsIssuesEnabled** | **optional.Bool**| Flag indication if the issue tracker is enabled | 
 **overrideParamsMergeRequestsEnabled** | **optional.Bool**| Flag indication if merge requests are enabled | 
 **overrideParamsWikiEnabled** | **optional.Bool**| Flag indication if the wiki is enabled | 
 **overrideParamsJobsEnabled** | **optional.Bool**| Flag indication if jobs are enabled | 
 **overrideParamsSnippetsEnabled** | **optional.Bool**| Flag indication if snippets are enabled | 
 **overrideParamsIssuesAccessLevel** | **optional.String**| Issues access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsRepositoryAccessLevel** | **optional.String**| Repository access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsMergeRequestsAccessLevel** | **optional.String**| Merge requests access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsForkingAccessLevel** | **optional.String**| Forks access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsWikiAccessLevel** | **optional.String**| Wiki access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsBuildsAccessLevel** | **optional.String**| Builds access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsSnippetsAccessLevel** | **optional.String**| Snippets access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsPagesAccessLevel** | **optional.String**| Pages access level. One of &#x60;disabled&#x60;, &#x60;private&#x60;, &#x60;enabled&#x60; or &#x60;public&#x60; | 
 **overrideParamsAnalyticsAccessLevel** | **optional.String**| Analytics access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsContainerRegistryAccessLevel** | **optional.String**| Controls visibility of the container registry. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60;. &#x60;private&#x60; will make the container registry accessible only to project members (reporter role and above). &#x60;enabled&#x60; will make the container registry accessible to everyone who has access to the project. &#x60;disabled&#x60; will disable the container registry | 
 **overrideParamsSecurityAndComplianceAccessLevel** | **optional.String**| Security and compliance access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsReleasesAccessLevel** | **optional.String**| Releases access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsEnvironmentsAccessLevel** | **optional.String**| Environments access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsFeatureFlagsAccessLevel** | **optional.String**| Feature flags access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsInfrastructureAccessLevel** | **optional.String**| Infrastructure access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsMonitorAccessLevel** | **optional.String**| Monitor access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsModelExperimentsAccessLevel** | **optional.String**| Model experiments access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsModelRegistryAccessLevel** | **optional.String**| Model registry access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsEmailsDisabled** | **optional.Bool**| Deprecated: Use emails_enabled instead. | 
 **overrideParamsEmailsEnabled** | **optional.Bool**| Enable email notifications | 
 **overrideParamsShowDefaultAwardEmojis** | **optional.Bool**| Show default award emojis | 
 **overrideParamsShowDiffPreviewInEmail** | **optional.Bool**| Include the code diff preview in merge request notification emails | 
 **overrideParamsWarnAboutPotentiallyUnwantedCharacters** | **optional.Bool**| Warn about potentially unwanted characters | 
 **overrideParamsEnforceAuthChecksOnUploads** | **optional.Bool**| Enforce auth check on uploads | 
 **overrideParamsSharedRunnersEnabled** | **optional.Bool**| Flag indication if shared runners are enabled for that project | 
 **overrideParamsGroupRunnersEnabled** | **optional.Bool**| Flag indication if group runners are enabled for that project | 
 **overrideParamsResolveOutdatedDiffDiscussions** | **optional.Bool**| Automatically resolve merge request diff threads on lines changed with a push | 
 **overrideParamsRemoveSourceBranchAfterMerge** | **optional.Bool**| Remove the source branch by default after merge | 
 **overrideParamsContainerRegistryEnabled** | **optional.Bool**| Deprecated: Use :container_registry_access_level instead. Flag indication if the container registry is enabled for that project | 
 **overrideParamsContainerExpirationPolicyAttributesCadence** | **optional.String**| Container expiration policy cadence for recurring job | 
 **overrideParamsContainerExpirationPolicyAttributesKeepN** | **optional.Int32**| Container expiration policy number of images to keep | 
 **overrideParamsContainerExpirationPolicyAttributesOlderThan** | **optional.String**| Container expiration policy remove images older than value | 
 **overrideParamsContainerExpirationPolicyAttributesNameRegex** | **optional.String**| Container expiration policy regex for image removal | 
 **overrideParamsContainerExpirationPolicyAttributesNameRegexKeep** | **optional.String**| Container expiration policy regex for image retention | 
 **overrideParamsContainerExpirationPolicyAttributesEnabled** | **optional.Bool**| Flag indication if container expiration policy is enabled | 
 **overrideParamsLfsEnabled** | **optional.Bool**| Flag indication if Git LFS is enabled for that project | 
 **overrideParamsVisibility** | **optional.String**| The visibility of the project. | 
 **overrideParamsPublicBuilds** | **optional.Bool**| Deprecated: Use public_jobs instead. | 
 **overrideParamsPublicJobs** | **optional.Bool**| Perform public builds | 
 **overrideParamsRequestAccessEnabled** | **optional.Bool**| Allow users to request member access | 
 **overrideParamsOnlyAllowMergeIfPipelineSucceeds** | **optional.Bool**| Only allow to merge if builds succeed | 
 **overrideParamsAllowMergeOnSkippedPipeline** | **optional.Bool**| Allow to merge if pipeline is skipped | 
 **overrideParamsOnlyAllowMergeIfAllDiscussionsAreResolved** | **optional.Bool**| Only allow to merge if all threads are resolved | 
 **overrideParamsTagList** | [**optional.Interface of []string**](string.md)| Deprecated: Use :topics instead | 
 **overrideParamsTopics** | [**optional.Interface of []string**](string.md)| The list of topics for a project | 
 **overrideParamsAvatar** | **optional.Interface of *os.File**| Avatar image for project | 
 **overrideParamsPrintingMergeRequestLinkEnabled** | **optional.Bool**| Show link to create/view merge request when pushing from the command line | 
 **overrideParamsMergeMethod** | **optional.String**| The merge method used when merging merge requests | 
 **overrideParamsSuggestionCommitMessage** | **optional.String**| The commit message used to apply merge request suggestions | 
 **overrideParamsMergeCommitTemplate** | **optional.String**| Template used to create merge commit message | 
 **overrideParamsSquashCommitTemplate** | **optional.String**| Template used to create squash commit message | 
 **overrideParamsIssueBranchTemplate** | **optional.String**| Template used to create a branch from an issue | 
 **overrideParamsInitializeWithReadme** | **optional.Bool**| Initialize a project with a README.md | 
 **overrideParamsAutoDevopsEnabled** | **optional.Bool**| Flag indication if Auto DevOps is enabled | 
 **overrideParamsAutoDevopsDeployStrategy** | **optional.String**| Auto Deploy strategy | 
 **overrideParamsAutocloseReferencedIssues** | **optional.Bool**| Flag indication if referenced issues auto-closing is enabled | 
 **overrideParamsRepositoryStorage** | **optional.String**| Which storage shard the repository is on. Available only to admins | 
 **overrideParamsPackagesEnabled** | **optional.Bool**| Enable project packages feature | 
 **overrideParamsSquashOption** | **optional.String**| Squash default for project. One of &#x60;never&#x60;, &#x60;always&#x60;, &#x60;default_on&#x60;, or &#x60;default_off&#x60;. | 
 **overrideParamsMrDefaultTargetSelf** | **optional.Bool**| Merge requests of this forked project targets itself by default | 
 **overrideParamsOnlyAllowMergeIfAllStatusChecksPassed** | **optional.Bool**| Blocks merge requests from merging unless all status checks have passed | 
 **overrideParamsApprovalsBeforeMerge** | **optional.Int32**| How many approvers should approve merge request by default | 
 **overrideParamsMirror** | **optional.Bool**| Enables pull mirroring in a project | 
 **overrideParamsMirrorTriggerBuilds** | **optional.Bool**| Pull mirroring triggers builds | 
 **overrideParamsExternalAuthorizationClassificationLabel** | **optional.String**| The classification label for the project | 
 **overrideParamsRequirementsAccessLevel** | **optional.String**| Requirements feature access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsPreventMergeWithoutJiraIssue** | **optional.Bool**| Require an associated issue from Jira | 

### Return type

[**ApiEntitiesProjectImportStatus**](API_Entities_ProjectImportStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsRemoteImportS3**
> ApiEntitiesProjectImportStatus PostApiV4ProjectsRemoteImportS3(ctx, region, bucketName, fileKey, accessKeyId, secretAccessKey, path, optional)
Create a new project import using a file from AWS S3

This feature was introduced in GitLab 14.9.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **region** | **string**| AWS region | 
  **bucketName** | **string**| Bucket name | 
  **fileKey** | **string**| File key | 
  **accessKeyId** | **string**| Access key id | 
  **secretAccessKey** | **string**| Secret access key | 
  **path** | **string**| The new project path and name | 
 **optional** | ***ProjectImportApiPostApiV4ProjectsRemoteImportS3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectImportApiPostApiV4ProjectsRemoteImportS3Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **name** | **optional.String**| The name of the project to be imported. Defaults to the path of the project if not provided. | 
 **namespace** | **optional.String**| The ID or name of the namespace that the project will be imported into. Defaults to the current user&#39;s namespace. | 
 **overwrite** | **optional.Bool**| If there is a project in the same namespace and with the same name overwrite it | [default to false]
 **overrideParamsDescription** | **optional.String**| The description of the project | 
 **overrideParamsBuildGitStrategy** | **optional.String**| The Git strategy. Defaults to &#x60;fetch&#x60; | 
 **overrideParamsBuildTimeout** | **optional.Int32**| Build timeout | 
 **overrideParamsAutoCancelPendingPipelines** | **optional.String**| Auto-cancel pending pipelines | 
 **overrideParamsCiConfigPath** | **optional.String**| The path to CI config file. Defaults to &#x60;.gitlab-ci.yml&#x60; | 
 **overrideParamsServiceDeskEnabled** | **optional.Bool**| Disable or enable the service desk | 
 **overrideParamsIssuesEnabled** | **optional.Bool**| Flag indication if the issue tracker is enabled | 
 **overrideParamsMergeRequestsEnabled** | **optional.Bool**| Flag indication if merge requests are enabled | 
 **overrideParamsWikiEnabled** | **optional.Bool**| Flag indication if the wiki is enabled | 
 **overrideParamsJobsEnabled** | **optional.Bool**| Flag indication if jobs are enabled | 
 **overrideParamsSnippetsEnabled** | **optional.Bool**| Flag indication if snippets are enabled | 
 **overrideParamsIssuesAccessLevel** | **optional.String**| Issues access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsRepositoryAccessLevel** | **optional.String**| Repository access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsMergeRequestsAccessLevel** | **optional.String**| Merge requests access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsForkingAccessLevel** | **optional.String**| Forks access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsWikiAccessLevel** | **optional.String**| Wiki access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsBuildsAccessLevel** | **optional.String**| Builds access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsSnippetsAccessLevel** | **optional.String**| Snippets access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsPagesAccessLevel** | **optional.String**| Pages access level. One of &#x60;disabled&#x60;, &#x60;private&#x60;, &#x60;enabled&#x60; or &#x60;public&#x60; | 
 **overrideParamsAnalyticsAccessLevel** | **optional.String**| Analytics access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsContainerRegistryAccessLevel** | **optional.String**| Controls visibility of the container registry. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60;. &#x60;private&#x60; will make the container registry accessible only to project members (reporter role and above). &#x60;enabled&#x60; will make the container registry accessible to everyone who has access to the project. &#x60;disabled&#x60; will disable the container registry | 
 **overrideParamsSecurityAndComplianceAccessLevel** | **optional.String**| Security and compliance access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsReleasesAccessLevel** | **optional.String**| Releases access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsEnvironmentsAccessLevel** | **optional.String**| Environments access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsFeatureFlagsAccessLevel** | **optional.String**| Feature flags access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsInfrastructureAccessLevel** | **optional.String**| Infrastructure access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsMonitorAccessLevel** | **optional.String**| Monitor access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsModelExperimentsAccessLevel** | **optional.String**| Model experiments access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsModelRegistryAccessLevel** | **optional.String**| Model registry access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsEmailsDisabled** | **optional.Bool**| Deprecated: Use emails_enabled instead. | 
 **overrideParamsEmailsEnabled** | **optional.Bool**| Enable email notifications | 
 **overrideParamsShowDefaultAwardEmojis** | **optional.Bool**| Show default award emojis | 
 **overrideParamsShowDiffPreviewInEmail** | **optional.Bool**| Include the code diff preview in merge request notification emails | 
 **overrideParamsWarnAboutPotentiallyUnwantedCharacters** | **optional.Bool**| Warn about potentially unwanted characters | 
 **overrideParamsEnforceAuthChecksOnUploads** | **optional.Bool**| Enforce auth check on uploads | 
 **overrideParamsSharedRunnersEnabled** | **optional.Bool**| Flag indication if shared runners are enabled for that project | 
 **overrideParamsGroupRunnersEnabled** | **optional.Bool**| Flag indication if group runners are enabled for that project | 
 **overrideParamsResolveOutdatedDiffDiscussions** | **optional.Bool**| Automatically resolve merge request diff threads on lines changed with a push | 
 **overrideParamsRemoveSourceBranchAfterMerge** | **optional.Bool**| Remove the source branch by default after merge | 
 **overrideParamsContainerRegistryEnabled** | **optional.Bool**| Deprecated: Use :container_registry_access_level instead. Flag indication if the container registry is enabled for that project | 
 **overrideParamsContainerExpirationPolicyAttributesCadence** | **optional.String**| Container expiration policy cadence for recurring job | 
 **overrideParamsContainerExpirationPolicyAttributesKeepN** | **optional.Int32**| Container expiration policy number of images to keep | 
 **overrideParamsContainerExpirationPolicyAttributesOlderThan** | **optional.String**| Container expiration policy remove images older than value | 
 **overrideParamsContainerExpirationPolicyAttributesNameRegex** | **optional.String**| Container expiration policy regex for image removal | 
 **overrideParamsContainerExpirationPolicyAttributesNameRegexKeep** | **optional.String**| Container expiration policy regex for image retention | 
 **overrideParamsContainerExpirationPolicyAttributesEnabled** | **optional.Bool**| Flag indication if container expiration policy is enabled | 
 **overrideParamsLfsEnabled** | **optional.Bool**| Flag indication if Git LFS is enabled for that project | 
 **overrideParamsVisibility** | **optional.String**| The visibility of the project. | 
 **overrideParamsPublicBuilds** | **optional.Bool**| Deprecated: Use public_jobs instead. | 
 **overrideParamsPublicJobs** | **optional.Bool**| Perform public builds | 
 **overrideParamsRequestAccessEnabled** | **optional.Bool**| Allow users to request member access | 
 **overrideParamsOnlyAllowMergeIfPipelineSucceeds** | **optional.Bool**| Only allow to merge if builds succeed | 
 **overrideParamsAllowMergeOnSkippedPipeline** | **optional.Bool**| Allow to merge if pipeline is skipped | 
 **overrideParamsOnlyAllowMergeIfAllDiscussionsAreResolved** | **optional.Bool**| Only allow to merge if all threads are resolved | 
 **overrideParamsTagList** | [**optional.Interface of []string**](string.md)| Deprecated: Use :topics instead | 
 **overrideParamsTopics** | [**optional.Interface of []string**](string.md)| The list of topics for a project | 
 **overrideParamsAvatar** | **optional.Interface of *os.File**| Avatar image for project | 
 **overrideParamsPrintingMergeRequestLinkEnabled** | **optional.Bool**| Show link to create/view merge request when pushing from the command line | 
 **overrideParamsMergeMethod** | **optional.String**| The merge method used when merging merge requests | 
 **overrideParamsSuggestionCommitMessage** | **optional.String**| The commit message used to apply merge request suggestions | 
 **overrideParamsMergeCommitTemplate** | **optional.String**| Template used to create merge commit message | 
 **overrideParamsSquashCommitTemplate** | **optional.String**| Template used to create squash commit message | 
 **overrideParamsIssueBranchTemplate** | **optional.String**| Template used to create a branch from an issue | 
 **overrideParamsInitializeWithReadme** | **optional.Bool**| Initialize a project with a README.md | 
 **overrideParamsAutoDevopsEnabled** | **optional.Bool**| Flag indication if Auto DevOps is enabled | 
 **overrideParamsAutoDevopsDeployStrategy** | **optional.String**| Auto Deploy strategy | 
 **overrideParamsAutocloseReferencedIssues** | **optional.Bool**| Flag indication if referenced issues auto-closing is enabled | 
 **overrideParamsRepositoryStorage** | **optional.String**| Which storage shard the repository is on. Available only to admins | 
 **overrideParamsPackagesEnabled** | **optional.Bool**| Enable project packages feature | 
 **overrideParamsSquashOption** | **optional.String**| Squash default for project. One of &#x60;never&#x60;, &#x60;always&#x60;, &#x60;default_on&#x60;, or &#x60;default_off&#x60;. | 
 **overrideParamsMrDefaultTargetSelf** | **optional.Bool**| Merge requests of this forked project targets itself by default | 
 **overrideParamsOnlyAllowMergeIfAllStatusChecksPassed** | **optional.Bool**| Blocks merge requests from merging unless all status checks have passed | 
 **overrideParamsApprovalsBeforeMerge** | **optional.Int32**| How many approvers should approve merge request by default | 
 **overrideParamsMirror** | **optional.Bool**| Enables pull mirroring in a project | 
 **overrideParamsMirrorTriggerBuilds** | **optional.Bool**| Pull mirroring triggers builds | 
 **overrideParamsExternalAuthorizationClassificationLabel** | **optional.String**| The classification label for the project | 
 **overrideParamsRequirementsAccessLevel** | **optional.String**| Requirements feature access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | 
 **overrideParamsPreventMergeWithoutJiraIssue** | **optional.Bool**| Require an associated issue from Jira | 

### Return type

[**ApiEntitiesProjectImportStatus**](API_Entities_ProjectImportStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

