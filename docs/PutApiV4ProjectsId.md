# PutApiV4ProjectsId

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the project | [optional] [default to null]
**DefaultBranch** | **string** | The default branch of the project | [optional] [default to null]
**Path** | **string** | The path of the repository | [optional] [default to null]
**Description** | **string** | The description of the project | [optional] [default to null]
**BuildGitStrategy** | **string** | The Git strategy. Defaults to &#x60;fetch&#x60; | [optional] [default to null]
**BuildTimeout** | **int32** | Build timeout | [optional] [default to null]
**AutoCancelPendingPipelines** | **string** | Auto-cancel pending pipelines | [optional] [default to null]
**CiConfigPath** | **string** | The path to CI config file. Defaults to &#x60;.gitlab-ci.yml&#x60; | [optional] [default to null]
**ServiceDeskEnabled** | **bool** | Disable or enable the service desk | [optional] [default to null]
**IssuesEnabled** | **bool** | Flag indication if the issue tracker is enabled | [optional] [default to null]
**MergeRequestsEnabled** | **bool** | Flag indication if merge requests are enabled | [optional] [default to null]
**WikiEnabled** | **bool** | Flag indication if the wiki is enabled | [optional] [default to null]
**JobsEnabled** | **bool** | Flag indication if jobs are enabled | [optional] [default to null]
**SnippetsEnabled** | **bool** | Flag indication if snippets are enabled | [optional] [default to null]
**IssuesAccessLevel** | **string** | Issues access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | [optional] [default to null]
**RepositoryAccessLevel** | **string** | Repository access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | [optional] [default to null]
**MergeRequestsAccessLevel** | **string** | Merge requests access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | [optional] [default to null]
**ForkingAccessLevel** | **string** | Forks access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | [optional] [default to null]
**WikiAccessLevel** | **string** | Wiki access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | [optional] [default to null]
**BuildsAccessLevel** | **string** | Builds access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | [optional] [default to null]
**SnippetsAccessLevel** | **string** | Snippets access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | [optional] [default to null]
**PagesAccessLevel** | **string** | Pages access level. One of &#x60;disabled&#x60;, &#x60;private&#x60;, &#x60;enabled&#x60; or &#x60;public&#x60; | [optional] [default to null]
**AnalyticsAccessLevel** | **string** | Analytics access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | [optional] [default to null]
**ContainerRegistryAccessLevel** | **string** | Controls visibility of the container registry. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60;. &#x60;private&#x60; will make the container registry accessible only to project members (reporter role and above). &#x60;enabled&#x60; will make the container registry accessible to everyone who has access to the project. &#x60;disabled&#x60; will disable the container registry | [optional] [default to null]
**SecurityAndComplianceAccessLevel** | **string** | Security and compliance access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | [optional] [default to null]
**ReleasesAccessLevel** | **string** | Releases access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | [optional] [default to null]
**EnvironmentsAccessLevel** | **string** | Environments access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | [optional] [default to null]
**FeatureFlagsAccessLevel** | **string** | Feature flags access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | [optional] [default to null]
**InfrastructureAccessLevel** | **string** | Infrastructure access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | [optional] [default to null]
**MonitorAccessLevel** | **string** | Monitor access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | [optional] [default to null]
**ModelExperimentsAccessLevel** | **string** | Model experiments access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | [optional] [default to null]
**ModelRegistryAccessLevel** | **string** | Model registry access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | [optional] [default to null]
**EmailsDisabled** | **bool** | Deprecated: Use emails_enabled instead. | [optional] [default to null]
**EmailsEnabled** | **bool** | Enable email notifications | [optional] [default to null]
**ShowDefaultAwardEmojis** | **bool** | Show default award emojis | [optional] [default to null]
**ShowDiffPreviewInEmail** | **bool** | Include the code diff preview in merge request notification emails | [optional] [default to null]
**WarnAboutPotentiallyUnwantedCharacters** | **bool** | Warn about potentially unwanted characters | [optional] [default to null]
**EnforceAuthChecksOnUploads** | **bool** | Enforce auth check on uploads | [optional] [default to null]
**SharedRunnersEnabled** | **bool** | Flag indication if shared runners are enabled for that project | [optional] [default to null]
**GroupRunnersEnabled** | **bool** | Flag indication if group runners are enabled for that project | [optional] [default to null]
**ResolveOutdatedDiffDiscussions** | **bool** | Automatically resolve merge request diff threads on lines changed with a push | [optional] [default to null]
**RemoveSourceBranchAfterMerge** | **bool** | Remove the source branch by default after merge | [optional] [default to null]
**ContainerRegistryEnabled** | **bool** | Deprecated: Use :container_registry_access_level instead. Flag indication if the container registry is enabled for that project | [optional] [default to null]
**ContainerExpirationPolicyAttributes** | [***PostApiV4ProjectsContainerExpirationPolicyAttributes**](postApiV4Projects_container_expiration_policy_attributes.md) |  | [optional] [default to null]
**LfsEnabled** | **bool** | Flag indication if Git LFS is enabled for that project | [optional] [default to null]
**Visibility** | **string** | The visibility of the project. | [optional] [default to null]
**PublicBuilds** | **bool** | Deprecated: Use public_jobs instead. | [optional] [default to null]
**PublicJobs** | **bool** | Perform public builds | [optional] [default to null]
**RequestAccessEnabled** | **bool** | Allow users to request member access | [optional] [default to null]
**OnlyAllowMergeIfPipelineSucceeds** | **bool** | Only allow to merge if builds succeed | [optional] [default to null]
**AllowMergeOnSkippedPipeline** | **bool** | Allow to merge if pipeline is skipped | [optional] [default to null]
**OnlyAllowMergeIfAllDiscussionsAreResolved** | **bool** | Only allow to merge if all threads are resolved | [optional] [default to null]
**TagList** | **[]string** | Deprecated: Use :topics instead | [optional] [default to null]
**Topics** | **[]string** | The list of topics for a project | [optional] [default to null]
**Avatar** | [****os.File**](*os.File.md) | Avatar image for project | [optional] [default to null]
**PrintingMergeRequestLinkEnabled** | **bool** | Show link to create/view merge request when pushing from the command line | [optional] [default to null]
**MergeMethod** | **string** | The merge method used when merging merge requests | [optional] [default to null]
**SuggestionCommitMessage** | **string** | The commit message used to apply merge request suggestions | [optional] [default to null]
**MergeCommitTemplate** | **string** | Template used to create merge commit message | [optional] [default to null]
**SquashCommitTemplate** | **string** | Template used to create squash commit message | [optional] [default to null]
**IssueBranchTemplate** | **string** | Template used to create a branch from an issue | [optional] [default to null]
**InitializeWithReadme** | **bool** | Initialize a project with a README.md | [optional] [default to null]
**AutoDevopsEnabled** | **bool** | Flag indication if Auto DevOps is enabled | [optional] [default to null]
**AutoDevopsDeployStrategy** | **string** | Auto Deploy strategy | [optional] [default to null]
**AutocloseReferencedIssues** | **bool** | Flag indication if referenced issues auto-closing is enabled | [optional] [default to null]
**RepositoryStorage** | **string** | Which storage shard the repository is on. Available only to admins | [optional] [default to null]
**PackagesEnabled** | **bool** | Enable project packages feature | [optional] [default to null]
**SquashOption** | **string** | Squash default for project. One of &#x60;never&#x60;, &#x60;always&#x60;, &#x60;default_on&#x60;, or &#x60;default_off&#x60;. | [optional] [default to null]
**MrDefaultTargetSelf** | **bool** | Merge requests of this forked project targets itself by default | [optional] [default to null]
**OnlyAllowMergeIfAllStatusChecksPassed** | **bool** | Blocks merge requests from merging unless all status checks have passed | [optional] [default to null]
**ApprovalsBeforeMerge** | **int32** | How many approvers should approve merge request by default | [optional] [default to null]
**Mirror** | **bool** | Enables pull mirroring in a project | [optional] [default to null]
**MirrorTriggerBuilds** | **bool** | Pull mirroring triggers builds | [optional] [default to null]
**ExternalAuthorizationClassificationLabel** | **string** | The classification label for the project | [optional] [default to null]
**RequirementsAccessLevel** | **string** | Requirements feature access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | [optional] [default to null]
**PreventMergeWithoutJiraIssue** | **bool** | Require an associated issue from Jira | [optional] [default to null]
**CiDefaultGitDepth** | **int32** | Default number of revisions for shallow cloning | [optional] [default to null]
**KeepLatestArtifact** | **bool** | Indicates if the latest artifact should be kept for this project. | [optional] [default to null]
**CiForwardDeploymentEnabled** | **bool** | Prevent older deployment jobs that are still pending | [optional] [default to null]
**CiForwardDeploymentRollbackAllowed** | **bool** | Allow job retries for rollback deployments | [optional] [default to null]
**CiAllowForkPipelinesToRunInParentProject** | **bool** | Allow fork merge request pipelines to run in parent project | [optional] [default to null]
**CiSeparatedCaches** | **bool** | Enable or disable separated caches based on branch protection. | [optional] [default to null]
**RestrictUserDefinedVariables** | **bool** | Restrict ability to override variables when triggering a pipeline | [optional] [default to null]
**CiPipelineVariablesMinimumOverrideRole** | **string** | Limit ability to override CI/CD variables when triggering a pipeline to only users with at least the set minimum role | [optional] [default to null]
**AllowPipelineTriggerApproveDeployment** | **bool** | Allow pipeline triggerer to approve deployments | [optional] [default to null]
**MirrorUserId** | **int32** | User responsible for all the activity surrounding a pull mirror event. Can only be set by admins | [optional] [default to null]
**OnlyMirrorProtectedBranches** | **bool** | Only mirror protected branches | [optional] [default to null]
**MirrorBranchRegex** | **string** | Only mirror branches match regex | [optional] [default to null]
**MirrorOverwritesDivergedBranches** | **bool** | Pull mirror overwrites diverged branches | [optional] [default to null]
**ImportUrl** | **string** | URL from which the project is imported | [optional] [default to null]
**FallbackApprovalsRequired** | **int32** | Overall approvals required when no rule is present | [optional] [default to null]
**IssuesTemplate** | **string** | Default description for Issues. Description is parsed with GitLab Flavored Markdown. | [optional] [default to null]
**MergeRequestsTemplate** | **string** | Default description for merge requests. Description is parsed with GitLab Flavored Markdown. | [optional] [default to null]
**MergePipelinesEnabled** | **bool** | Enable merged results pipelines. | [optional] [default to null]
**MergeTrainsEnabled** | **bool** | Enable merge trains. | [optional] [default to null]
**MergeTrainsSkipTrainAllowed** | **bool** | Allow merge train merge requests to be merged without waiting for pipelines to finish. | [optional] [default to null]
**CiRestrictPipelineCancellationRole** | **string** | Roles allowed to cancel pipelines and jobs. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


