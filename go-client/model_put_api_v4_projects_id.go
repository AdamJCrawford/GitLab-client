/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"os"
)

// Update an existing project
type PutApiV4ProjectsId struct {
	// The name of the project
	Name string `json:"name,omitempty"`
	// The default branch of the project
	DefaultBranch string `json:"default_branch,omitempty"`
	// The path of the repository
	Path string `json:"path,omitempty"`
	// The description of the project
	Description string `json:"description,omitempty"`
	// The Git strategy. Defaults to `fetch`
	BuildGitStrategy string `json:"build_git_strategy,omitempty"`
	// Build timeout
	BuildTimeout int32 `json:"build_timeout,omitempty"`
	// Auto-cancel pending pipelines
	AutoCancelPendingPipelines string `json:"auto_cancel_pending_pipelines,omitempty"`
	// The path to CI config file. Defaults to `.gitlab-ci.yml`
	CiConfigPath string `json:"ci_config_path,omitempty"`
	// Disable or enable the service desk
	ServiceDeskEnabled bool `json:"service_desk_enabled,omitempty"`
	// Flag indication if the issue tracker is enabled
	IssuesEnabled bool `json:"issues_enabled,omitempty"`
	// Flag indication if merge requests are enabled
	MergeRequestsEnabled bool `json:"merge_requests_enabled,omitempty"`
	// Flag indication if the wiki is enabled
	WikiEnabled bool `json:"wiki_enabled,omitempty"`
	// Flag indication if jobs are enabled
	JobsEnabled bool `json:"jobs_enabled,omitempty"`
	// Flag indication if snippets are enabled
	SnippetsEnabled bool `json:"snippets_enabled,omitempty"`
	// Issues access level. One of `disabled`, `private` or `enabled`
	IssuesAccessLevel string `json:"issues_access_level,omitempty"`
	// Repository access level. One of `disabled`, `private` or `enabled`
	RepositoryAccessLevel string `json:"repository_access_level,omitempty"`
	// Merge requests access level. One of `disabled`, `private` or `enabled`
	MergeRequestsAccessLevel string `json:"merge_requests_access_level,omitempty"`
	// Forks access level. One of `disabled`, `private` or `enabled`
	ForkingAccessLevel string `json:"forking_access_level,omitempty"`
	// Wiki access level. One of `disabled`, `private` or `enabled`
	WikiAccessLevel string `json:"wiki_access_level,omitempty"`
	// Builds access level. One of `disabled`, `private` or `enabled`
	BuildsAccessLevel string `json:"builds_access_level,omitempty"`
	// Snippets access level. One of `disabled`, `private` or `enabled`
	SnippetsAccessLevel string `json:"snippets_access_level,omitempty"`
	// Pages access level. One of `disabled`, `private`, `enabled` or `public`
	PagesAccessLevel string `json:"pages_access_level,omitempty"`
	// Analytics access level. One of `disabled`, `private` or `enabled`
	AnalyticsAccessLevel string `json:"analytics_access_level,omitempty"`
	// Controls visibility of the container registry. One of `disabled`, `private` or `enabled`. `private` will make the container registry accessible only to project members (reporter role and above). `enabled` will make the container registry accessible to everyone who has access to the project. `disabled` will disable the container registry
	ContainerRegistryAccessLevel string `json:"container_registry_access_level,omitempty"`
	// Security and compliance access level. One of `disabled`, `private` or `enabled`
	SecurityAndComplianceAccessLevel string `json:"security_and_compliance_access_level,omitempty"`
	// Releases access level. One of `disabled`, `private` or `enabled`
	ReleasesAccessLevel string `json:"releases_access_level,omitempty"`
	// Environments access level. One of `disabled`, `private` or `enabled`
	EnvironmentsAccessLevel string `json:"environments_access_level,omitempty"`
	// Feature flags access level. One of `disabled`, `private` or `enabled`
	FeatureFlagsAccessLevel string `json:"feature_flags_access_level,omitempty"`
	// Infrastructure access level. One of `disabled`, `private` or `enabled`
	InfrastructureAccessLevel string `json:"infrastructure_access_level,omitempty"`
	// Monitor access level. One of `disabled`, `private` or `enabled`
	MonitorAccessLevel string `json:"monitor_access_level,omitempty"`
	// Model experiments access level. One of `disabled`, `private` or `enabled`
	ModelExperimentsAccessLevel string `json:"model_experiments_access_level,omitempty"`
	// Model registry access level. One of `disabled`, `private` or `enabled`
	ModelRegistryAccessLevel string `json:"model_registry_access_level,omitempty"`
	// Deprecated: Use emails_enabled instead.
	EmailsDisabled bool `json:"emails_disabled,omitempty"`
	// Enable email notifications
	EmailsEnabled bool `json:"emails_enabled,omitempty"`
	// Show default award emojis
	ShowDefaultAwardEmojis bool `json:"show_default_award_emojis,omitempty"`
	// Include the code diff preview in merge request notification emails
	ShowDiffPreviewInEmail bool `json:"show_diff_preview_in_email,omitempty"`
	// Warn about potentially unwanted characters
	WarnAboutPotentiallyUnwantedCharacters bool `json:"warn_about_potentially_unwanted_characters,omitempty"`
	// Enforce auth check on uploads
	EnforceAuthChecksOnUploads bool `json:"enforce_auth_checks_on_uploads,omitempty"`
	// Flag indication if shared runners are enabled for that project
	SharedRunnersEnabled bool `json:"shared_runners_enabled,omitempty"`
	// Flag indication if group runners are enabled for that project
	GroupRunnersEnabled bool `json:"group_runners_enabled,omitempty"`
	// Automatically resolve merge request diff threads on lines changed with a push
	ResolveOutdatedDiffDiscussions bool `json:"resolve_outdated_diff_discussions,omitempty"`
	// Remove the source branch by default after merge
	RemoveSourceBranchAfterMerge bool `json:"remove_source_branch_after_merge,omitempty"`
	// Deprecated: Use :container_registry_access_level instead. Flag indication if the container registry is enabled for that project
	ContainerRegistryEnabled            bool                                                  `json:"container_registry_enabled,omitempty"`
	ContainerExpirationPolicyAttributes *PostApiV4ProjectsContainerExpirationPolicyAttributes `json:"container_expiration_policy_attributes,omitempty"`
	// Flag indication if Git LFS is enabled for that project
	LfsEnabled bool `json:"lfs_enabled,omitempty"`
	// The visibility of the project.
	Visibility string `json:"visibility,omitempty"`
	// Deprecated: Use public_jobs instead.
	PublicBuilds bool `json:"public_builds,omitempty"`
	// Perform public builds
	PublicJobs bool `json:"public_jobs,omitempty"`
	// Allow users to request member access
	RequestAccessEnabled bool `json:"request_access_enabled,omitempty"`
	// Only allow to merge if builds succeed
	OnlyAllowMergeIfPipelineSucceeds bool `json:"only_allow_merge_if_pipeline_succeeds,omitempty"`
	// Allow to merge if pipeline is skipped
	AllowMergeOnSkippedPipeline bool `json:"allow_merge_on_skipped_pipeline,omitempty"`
	// Only allow to merge if all threads are resolved
	OnlyAllowMergeIfAllDiscussionsAreResolved bool `json:"only_allow_merge_if_all_discussions_are_resolved,omitempty"`
	// Deprecated: Use :topics instead
	TagList []string `json:"tag_list,omitempty"`
	// The list of topics for a project
	Topics []string `json:"topics,omitempty"`
	// Avatar image for project
	Avatar **os.File `json:"avatar,omitempty"`
	// Show link to create/view merge request when pushing from the command line
	PrintingMergeRequestLinkEnabled bool `json:"printing_merge_request_link_enabled,omitempty"`
	// The merge method used when merging merge requests
	MergeMethod string `json:"merge_method,omitempty"`
	// The commit message used to apply merge request suggestions
	SuggestionCommitMessage string `json:"suggestion_commit_message,omitempty"`
	// Template used to create merge commit message
	MergeCommitTemplate string `json:"merge_commit_template,omitempty"`
	// Template used to create squash commit message
	SquashCommitTemplate string `json:"squash_commit_template,omitempty"`
	// Template used to create a branch from an issue
	IssueBranchTemplate string `json:"issue_branch_template,omitempty"`
	// Initialize a project with a README.md
	InitializeWithReadme bool `json:"initialize_with_readme,omitempty"`
	// Flag indication if Auto DevOps is enabled
	AutoDevopsEnabled bool `json:"auto_devops_enabled,omitempty"`
	// Auto Deploy strategy
	AutoDevopsDeployStrategy string `json:"auto_devops_deploy_strategy,omitempty"`
	// Flag indication if referenced issues auto-closing is enabled
	AutocloseReferencedIssues bool `json:"autoclose_referenced_issues,omitempty"`
	// Which storage shard the repository is on. Available only to admins
	RepositoryStorage string `json:"repository_storage,omitempty"`
	// Enable project packages feature
	PackagesEnabled bool `json:"packages_enabled,omitempty"`
	// Squash default for project. One of `never`, `always`, `default_on`, or `default_off`.
	SquashOption string `json:"squash_option,omitempty"`
	// Merge requests of this forked project targets itself by default
	MrDefaultTargetSelf bool `json:"mr_default_target_self,omitempty"`
	// Blocks merge requests from merging unless all status checks have passed
	OnlyAllowMergeIfAllStatusChecksPassed bool `json:"only_allow_merge_if_all_status_checks_passed,omitempty"`
	// How many approvers should approve merge request by default
	ApprovalsBeforeMerge int32 `json:"approvals_before_merge,omitempty"`
	// Enables pull mirroring in a project
	Mirror bool `json:"mirror,omitempty"`
	// Pull mirroring triggers builds
	MirrorTriggerBuilds bool `json:"mirror_trigger_builds,omitempty"`
	// The classification label for the project
	ExternalAuthorizationClassificationLabel string `json:"external_authorization_classification_label,omitempty"`
	// Requirements feature access level. One of `disabled`, `private` or `enabled`
	RequirementsAccessLevel string `json:"requirements_access_level,omitempty"`
	// Require an associated issue from Jira
	PreventMergeWithoutJiraIssue bool `json:"prevent_merge_without_jira_issue,omitempty"`
	// Default number of revisions for shallow cloning
	CiDefaultGitDepth int32 `json:"ci_default_git_depth,omitempty"`
	// Indicates if the latest artifact should be kept for this project.
	KeepLatestArtifact bool `json:"keep_latest_artifact,omitempty"`
	// Prevent older deployment jobs that are still pending
	CiForwardDeploymentEnabled bool `json:"ci_forward_deployment_enabled,omitempty"`
	// Allow job retries for rollback deployments
	CiForwardDeploymentRollbackAllowed bool `json:"ci_forward_deployment_rollback_allowed,omitempty"`
	// Allow fork merge request pipelines to run in parent project
	CiAllowForkPipelinesToRunInParentProject bool `json:"ci_allow_fork_pipelines_to_run_in_parent_project,omitempty"`
	// Enable or disable separated caches based on branch protection.
	CiSeparatedCaches bool `json:"ci_separated_caches,omitempty"`
	// Restrict ability to override variables when triggering a pipeline
	RestrictUserDefinedVariables bool `json:"restrict_user_defined_variables,omitempty"`
	// Limit ability to override CI/CD variables when triggering a pipeline to only users with at least the set minimum role
	CiPipelineVariablesMinimumOverrideRole string `json:"ci_pipeline_variables_minimum_override_role,omitempty"`
	// Allow pipeline triggerer to approve deployments
	AllowPipelineTriggerApproveDeployment bool `json:"allow_pipeline_trigger_approve_deployment,omitempty"`
	// User responsible for all the activity surrounding a pull mirror event. Can only be set by admins
	MirrorUserId int32 `json:"mirror_user_id,omitempty"`
	// Only mirror protected branches
	OnlyMirrorProtectedBranches bool `json:"only_mirror_protected_branches,omitempty"`
	// Only mirror branches match regex
	MirrorBranchRegex string `json:"mirror_branch_regex,omitempty"`
	// Pull mirror overwrites diverged branches
	MirrorOverwritesDivergedBranches bool `json:"mirror_overwrites_diverged_branches,omitempty"`
	// URL from which the project is imported
	ImportUrl string `json:"import_url,omitempty"`
	// Overall approvals required when no rule is present
	FallbackApprovalsRequired int32 `json:"fallback_approvals_required,omitempty"`
	// Default description for Issues. Description is parsed with GitLab Flavored Markdown.
	IssuesTemplate string `json:"issues_template,omitempty"`
	// Default description for merge requests. Description is parsed with GitLab Flavored Markdown.
	MergeRequestsTemplate string `json:"merge_requests_template,omitempty"`
	// Enable merged results pipelines.
	MergePipelinesEnabled bool `json:"merge_pipelines_enabled,omitempty"`
	// Enable merge trains.
	MergeTrainsEnabled bool `json:"merge_trains_enabled,omitempty"`
	// Allow merge train merge requests to be merged without waiting for pipelines to finish.
	MergeTrainsSkipTrainAllowed bool `json:"merge_trains_skip_train_allowed,omitempty"`
	// Roles allowed to cancel pipelines and jobs.
	CiRestrictPipelineCancellationRole string `json:"ci_restrict_pipeline_cancellation_role,omitempty"`
}
