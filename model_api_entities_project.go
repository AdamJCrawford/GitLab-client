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
	"time"
)

// API_Entities_Project model
type ApiEntitiesProject struct {
	Id                                        int32                                 `json:"id,omitempty"`
	Description                               string                                `json:"description,omitempty"`
	Name                                      string                                `json:"name,omitempty"`
	NameWithNamespace                         string                                `json:"name_with_namespace,omitempty"`
	Path                                      string                                `json:"path,omitempty"`
	PathWithNamespace                         string                                `json:"path_with_namespace,omitempty"`
	CreatedAt                                 time.Time                             `json:"created_at,omitempty"`
	DefaultBranch                             string                                `json:"default_branch,omitempty"`
	TagList                                   []string                              `json:"tag_list,omitempty"`
	Topics                                    []string                              `json:"topics,omitempty"`
	SshUrlToRepo                              string                                `json:"ssh_url_to_repo,omitempty"`
	HttpUrlToRepo                             string                                `json:"http_url_to_repo,omitempty"`
	WebUrl                                    string                                `json:"web_url,omitempty"`
	ReadmeUrl                                 string                                `json:"readme_url,omitempty"`
	ForksCount                                int32                                 `json:"forks_count,omitempty"`
	LicenseUrl                                string                                `json:"license_url,omitempty"`
	License                                   *ApiEntitiesLicenseBasic              `json:"license,omitempty"`
	AvatarUrl                                 string                                `json:"avatar_url,omitempty"`
	StarCount                                 int32                                 `json:"star_count,omitempty"`
	LastActivityAt                            time.Time                             `json:"last_activity_at,omitempty"`
	Namespace                                 *ApiEntitiesNamespaceBasic            `json:"namespace,omitempty"`
	CustomAttributes                          *ApiEntitiesCustomAttribute           `json:"custom_attributes,omitempty"`
	RepositoryStorage                         string                                `json:"repository_storage,omitempty"`
	ContainerRegistryImagePrefix              string                                `json:"container_registry_image_prefix,omitempty"`
	Links                                     *ApiEntitiesProjectLinks              `json:"_links,omitempty"`
	PackagesEnabled                           bool                                  `json:"packages_enabled,omitempty"`
	EmptyRepo                                 bool                                  `json:"empty_repo,omitempty"`
	Archived                                  bool                                  `json:"archived,omitempty"`
	Visibility                                string                                `json:"visibility,omitempty"`
	Owner                                     *ApiEntitiesUserBasic                 `json:"owner,omitempty"`
	ResolveOutdatedDiffDiscussions            bool                                  `json:"resolve_outdated_diff_discussions,omitempty"`
	ContainerExpirationPolicy                 *ApiEntitiesContainerExpirationPolicy `json:"container_expiration_policy,omitempty"`
	RepositoryObjectFormat                    string                                `json:"repository_object_format,omitempty"`
	IssuesEnabled                             bool                                  `json:"issues_enabled,omitempty"`
	MergeRequestsEnabled                      bool                                  `json:"merge_requests_enabled,omitempty"`
	WikiEnabled                               bool                                  `json:"wiki_enabled,omitempty"`
	JobsEnabled                               bool                                  `json:"jobs_enabled,omitempty"`
	SnippetsEnabled                           bool                                  `json:"snippets_enabled,omitempty"`
	ContainerRegistryEnabled                  bool                                  `json:"container_registry_enabled,omitempty"`
	ServiceDeskEnabled                        bool                                  `json:"service_desk_enabled,omitempty"`
	ServiceDeskAddress                        string                                `json:"service_desk_address,omitempty"`
	CanCreateMergeRequestIn                   bool                                  `json:"can_create_merge_request_in,omitempty"`
	IssuesAccessLevel                         string                                `json:"issues_access_level,omitempty"`
	RepositoryAccessLevel                     string                                `json:"repository_access_level,omitempty"`
	MergeRequestsAccessLevel                  string                                `json:"merge_requests_access_level,omitempty"`
	ForkingAccessLevel                        string                                `json:"forking_access_level,omitempty"`
	WikiAccessLevel                           string                                `json:"wiki_access_level,omitempty"`
	BuildsAccessLevel                         string                                `json:"builds_access_level,omitempty"`
	SnippetsAccessLevel                       string                                `json:"snippets_access_level,omitempty"`
	PagesAccessLevel                          string                                `json:"pages_access_level,omitempty"`
	AnalyticsAccessLevel                      string                                `json:"analytics_access_level,omitempty"`
	ContainerRegistryAccessLevel              string                                `json:"container_registry_access_level,omitempty"`
	SecurityAndComplianceAccessLevel          string                                `json:"security_and_compliance_access_level,omitempty"`
	ReleasesAccessLevel                       string                                `json:"releases_access_level,omitempty"`
	EnvironmentsAccessLevel                   string                                `json:"environments_access_level,omitempty"`
	FeatureFlagsAccessLevel                   string                                `json:"feature_flags_access_level,omitempty"`
	InfrastructureAccessLevel                 string                                `json:"infrastructure_access_level,omitempty"`
	MonitorAccessLevel                        string                                `json:"monitor_access_level,omitempty"`
	ModelExperimentsAccessLevel               string                                `json:"model_experiments_access_level,omitempty"`
	ModelRegistryAccessLevel                  string                                `json:"model_registry_access_level,omitempty"`
	EmailsDisabled                            bool                                  `json:"emails_disabled,omitempty"`
	EmailsEnabled                             bool                                  `json:"emails_enabled,omitempty"`
	SharedRunnersEnabled                      bool                                  `json:"shared_runners_enabled,omitempty"`
	LfsEnabled                                bool                                  `json:"lfs_enabled,omitempty"`
	CreatorId                                 int32                                 `json:"creator_id,omitempty"`
	ForkedFromProject                         *ApiEntitiesBasicProjectDetails       `json:"forked_from_project,omitempty"`
	MrDefaultTargetSelf                       bool                                  `json:"mr_default_target_self,omitempty"`
	ImportUrl                                 string                                `json:"import_url,omitempty"`
	ImportType                                string                                `json:"import_type,omitempty"`
	ImportStatus                              string                                `json:"import_status,omitempty"`
	ImportError                               string                                `json:"import_error,omitempty"`
	OpenIssuesCount                           int32                                 `json:"open_issues_count,omitempty"`
	DescriptionHtml                           string                                `json:"description_html,omitempty"`
	UpdatedAt                                 time.Time                             `json:"updated_at,omitempty"`
	CiDefaultGitDepth                         int32                                 `json:"ci_default_git_depth,omitempty"`
	CiForwardDeploymentEnabled                bool                                  `json:"ci_forward_deployment_enabled,omitempty"`
	CiForwardDeploymentRollbackAllowed        bool                                  `json:"ci_forward_deployment_rollback_allowed,omitempty"`
	CiJobTokenScopeEnabled                    bool                                  `json:"ci_job_token_scope_enabled,omitempty"`
	CiSeparatedCaches                         bool                                  `json:"ci_separated_caches,omitempty"`
	CiAllowForkPipelinesToRunInParentProject  bool                                  `json:"ci_allow_fork_pipelines_to_run_in_parent_project,omitempty"`
	BuildGitStrategy                          string                                `json:"build_git_strategy,omitempty"`
	KeepLatestArtifact                        bool                                  `json:"keep_latest_artifact,omitempty"`
	RestrictUserDefinedVariables              bool                                  `json:"restrict_user_defined_variables,omitempty"`
	CiPipelineVariablesMinimumOverrideRole    string                                `json:"ci_pipeline_variables_minimum_override_role,omitempty"`
	RunnersToken                              string                                `json:"runners_token,omitempty"`
	RunnerTokenExpirationInterval             int32                                 `json:"runner_token_expiration_interval,omitempty"`
	GroupRunnersEnabled                       bool                                  `json:"group_runners_enabled,omitempty"`
	AutoCancelPendingPipelines                string                                `json:"auto_cancel_pending_pipelines,omitempty"`
	BuildTimeout                              int32                                 `json:"build_timeout,omitempty"`
	AutoDevopsEnabled                         bool                                  `json:"auto_devops_enabled,omitempty"`
	AutoDevopsDeployStrategy                  string                                `json:"auto_devops_deploy_strategy,omitempty"`
	CiConfigPath                              string                                `json:"ci_config_path,omitempty"`
	PublicJobs                                bool                                  `json:"public_jobs,omitempty"`
	SharedWithGroups                          []string                              `json:"shared_with_groups,omitempty"`
	OnlyAllowMergeIfPipelineSucceeds          bool                                  `json:"only_allow_merge_if_pipeline_succeeds,omitempty"`
	AllowMergeOnSkippedPipeline               bool                                  `json:"allow_merge_on_skipped_pipeline,omitempty"`
	RequestAccessEnabled                      bool                                  `json:"request_access_enabled,omitempty"`
	OnlyAllowMergeIfAllDiscussionsAreResolved bool                                  `json:"only_allow_merge_if_all_discussions_are_resolved,omitempty"`
	RemoveSourceBranchAfterMerge              bool                                  `json:"remove_source_branch_after_merge,omitempty"`
	PrintingMergeRequestLinkEnabled           bool                                  `json:"printing_merge_request_link_enabled,omitempty"`
	MergeMethod                               string                                `json:"merge_method,omitempty"`
	SquashOption                              string                                `json:"squash_option,omitempty"`
	EnforceAuthChecksOnUploads                bool                                  `json:"enforce_auth_checks_on_uploads,omitempty"`
	SuggestionCommitMessage                   string                                `json:"suggestion_commit_message,omitempty"`
	MergeCommitTemplate                       string                                `json:"merge_commit_template,omitempty"`
	SquashCommitTemplate                      string                                `json:"squash_commit_template,omitempty"`
	IssueBranchTemplate                       string                                `json:"issue_branch_template,omitempty"`
	Statistics                                *ApiEntitiesProjectStatistics         `json:"statistics,omitempty"`
	WarnAboutPotentiallyUnwantedCharacters    bool                                  `json:"warn_about_potentially_unwanted_characters,omitempty"`
	AutocloseReferencedIssues                 bool                                  `json:"autoclose_referenced_issues,omitempty"`
	ApprovalsBeforeMerge                      string                                `json:"approvals_before_merge,omitempty"`
	Mirror                                    string                                `json:"mirror,omitempty"`
	MirrorUserId                              string                                `json:"mirror_user_id,omitempty"`
	MirrorTriggerBuilds                       string                                `json:"mirror_trigger_builds,omitempty"`
	OnlyMirrorProtectedBranches               string                                `json:"only_mirror_protected_branches,omitempty"`
	MirrorOverwritesDivergedBranches          string                                `json:"mirror_overwrites_diverged_branches,omitempty"`
	ExternalAuthorizationClassificationLabel  string                                `json:"external_authorization_classification_label,omitempty"`
	MarkedForDeletionAt                       string                                `json:"marked_for_deletion_at,omitempty"`
	MarkedForDeletionOn                       string                                `json:"marked_for_deletion_on,omitempty"`
	RequirementsEnabled                       string                                `json:"requirements_enabled,omitempty"`
	RequirementsAccessLevel                   string                                `json:"requirements_access_level,omitempty"`
	SecurityAndComplianceEnabled              string                                `json:"security_and_compliance_enabled,omitempty"`
	ComplianceFrameworks                      string                                `json:"compliance_frameworks,omitempty"`
	IssuesTemplate                            string                                `json:"issues_template,omitempty"`
	MergeRequestsTemplate                     string                                `json:"merge_requests_template,omitempty"`
	CiRestrictPipelineCancellationRole        string                                `json:"ci_restrict_pipeline_cancellation_role,omitempty"`
	MergePipelinesEnabled                     string                                `json:"merge_pipelines_enabled,omitempty"`
	MergeTrainsEnabled                        string                                `json:"merge_trains_enabled,omitempty"`
	MergeTrainsSkipTrainAllowed               string                                `json:"merge_trains_skip_train_allowed,omitempty"`
	OnlyAllowMergeIfAllStatusChecksPassed     string                                `json:"only_allow_merge_if_all_status_checks_passed,omitempty"`
	AllowPipelineTriggerApproveDeployment     bool                                  `json:"allow_pipeline_trigger_approve_deployment,omitempty"`
	PreventMergeWithoutJiraIssue              string                                `json:"prevent_merge_without_jira_issue,omitempty"`
}