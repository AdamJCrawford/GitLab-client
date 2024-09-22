/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Change plan limits
type PutApiV4ApplicationPlanLimits struct {
	// Name of the plan to update
	PlanName string `json:"plan_name"`
	// Maximum number of Instance-level CI/CD variables that can be defined
	CiInstanceLevelVariables int32 `json:"ci_instance_level_variables,omitempty"`
	// Maximum number of jobs in a single pipeline
	CiPipelineSize int32 `json:"ci_pipeline_size,omitempty"`
	// Total number of jobs in currently active pipelines
	CiActiveJobs int32 `json:"ci_active_jobs,omitempty"`
	// Maximum number of pipeline subscriptions to and from a project
	CiProjectSubscriptions int32 `json:"ci_project_subscriptions,omitempty"`
	// Maximum number of pipeline schedules
	CiPipelineSchedules int32 `json:"ci_pipeline_schedules,omitempty"`
	// Maximum number of needs dependencies that a job can have
	CiNeedsSizeLimit int32 `json:"ci_needs_size_limit,omitempty"`
	// Maximum number of runners registered per group
	CiRegisteredGroupRunners int32 `json:"ci_registered_group_runners,omitempty"`
	// Maximum number of runners registered per project
	CiRegisteredProjectRunners int32 `json:"ci_registered_project_runners,omitempty"`
	// Maximum Conan package file size in bytes
	ConanMaxFileSize int32 `json:"conan_max_file_size,omitempty"`
	// Maximum storage size for the root namespace enforcement in MiB
	EnforcementLimit int32 `json:"enforcement_limit,omitempty"`
	// Maximum generic package file size in bytes
	GenericPackagesMaxFileSize int32 `json:"generic_packages_max_file_size,omitempty"`
	// Maximum Helm chart file size in bytes
	HelmMaxFileSize int32 `json:"helm_max_file_size,omitempty"`
	// Maximum Maven package file size in bytes
	MavenMaxFileSize int32 `json:"maven_max_file_size,omitempty"`
	// Maximum storage size for the root namespace notifications in MiB
	NotificationLimit int32 `json:"notification_limit,omitempty"`
	// Maximum NPM package file size in bytes
	NpmMaxFileSize int32 `json:"npm_max_file_size,omitempty"`
	// Maximum NuGet package file size in bytes
	NugetMaxFileSize int32 `json:"nuget_max_file_size,omitempty"`
	// Maximum PyPI package file size in bytes
	PypiMaxFileSize int32 `json:"pypi_max_file_size,omitempty"`
	// Maximum Terraform Module package file size in bytes
	TerraformModuleMaxFileSize int32 `json:"terraform_module_max_file_size,omitempty"`
	// Maximum storage size for the root namespace in MiB
	StorageSizeLimit int32 `json:"storage_size_limit,omitempty"`
	// Maximum number of downstream pipelines in a pipeline's hierarchy tree
	PipelineHierarchySize int32 `json:"pipeline_hierarchy_size,omitempty"`
}