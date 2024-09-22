/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type PostApiV4GeoStatusDataStatus struct {
	// Projects count
	ProjectsCount int32 `json:"projects_count,omitempty"`
	// Container repositories replication enabled
	ContainerRepositoriesReplicationEnabled bool `json:"container_repositories_replication_enabled,omitempty"`
	// LFS objects count
	LfsObjectsCount int32 `json:"lfs_objects_count,omitempty"`
	// LFS objects checksum total count
	LfsObjectsChecksumTotalCount int32 `json:"lfs_objects_checksum_total_count,omitempty"`
	// LFS objects checksummed count
	LfsObjectsChecksummedCount int32 `json:"lfs_objects_checksummed_count,omitempty"`
	// LFS objects checksum failed count
	LfsObjectsChecksumFailedCount int32 `json:"lfs_objects_checksum_failed_count,omitempty"`
	// LFS objects synced count
	LfsObjectsSyncedCount int32 `json:"lfs_objects_synced_count,omitempty"`
	// LFS objects failed count
	LfsObjectsFailedCount int32 `json:"lfs_objects_failed_count,omitempty"`
	// LFS objects registry count
	LfsObjectsRegistryCount int32 `json:"lfs_objects_registry_count,omitempty"`
	// LFS objects verification total count
	LfsObjectsVerificationTotalCount int32 `json:"lfs_objects_verification_total_count,omitempty"`
	// LFS objects verified count
	LfsObjectsVerifiedCount int32 `json:"lfs_objects_verified_count,omitempty"`
	// LFS objects verification failed count
	LfsObjectsVerificationFailedCount int32 `json:"lfs_objects_verification_failed_count,omitempty"`
	// Merge request diffs count
	MergeRequestDiffsCount int32 `json:"merge_request_diffs_count,omitempty"`
	// Merge request diffs checksum total count
	MergeRequestDiffsChecksumTotalCount int32 `json:"merge_request_diffs_checksum_total_count,omitempty"`
	// Merge request diffs checksummed count
	MergeRequestDiffsChecksummedCount int32 `json:"merge_request_diffs_checksummed_count,omitempty"`
	// Merge request diffs checksum failed count
	MergeRequestDiffsChecksumFailedCount int32 `json:"merge_request_diffs_checksum_failed_count,omitempty"`
	// Merge request diffs synced count
	MergeRequestDiffsSyncedCount int32 `json:"merge_request_diffs_synced_count,omitempty"`
	// Merge request diffs failed count
	MergeRequestDiffsFailedCount int32 `json:"merge_request_diffs_failed_count,omitempty"`
	// Merge request diffs registry count
	MergeRequestDiffsRegistryCount int32 `json:"merge_request_diffs_registry_count,omitempty"`
	// Merge request diffs verification total count
	MergeRequestDiffsVerificationTotalCount int32 `json:"merge_request_diffs_verification_total_count,omitempty"`
	// Merge request diffs verified count
	MergeRequestDiffsVerifiedCount int32 `json:"merge_request_diffs_verified_count,omitempty"`
	// Merge request diffs verified count
	MergeRequestDiffsVerificationFailedCount int32 `json:"merge_request_diffs_verification_failed_count,omitempty"`
	// Packages files count
	PackageFilesCount int32 `json:"package_files_count,omitempty"`
	// Packages files checksum total count
	PackageFilesChecksumTotalCount int32 `json:"package_files_checksum_total_count,omitempty"`
	// Packages files checksummed count
	PackageFilesChecksummedCount int32 `json:"package_files_checksummed_count,omitempty"`
	// Packages files checksum failed count
	PackageFilesChecksumFailedCount int32 `json:"package_files_checksum_failed_count,omitempty"`
	// Packages files synced count
	PackageFilesSyncedCount int32 `json:"package_files_synced_count,omitempty"`
	// Packages files failed count
	PackageFilesFailedCount int32 `json:"package_files_failed_count,omitempty"`
	// Packages files registry count
	PackageFilesRegistryCount int32 `json:"package_files_registry_count,omitempty"`
	// Packages files verification total count
	PackageFilesVerificationTotalCount int32 `json:"package_files_verification_total_count,omitempty"`
	// Packages files verified count
	PackageFilesVerifiedCount int32 `json:"package_files_verified_count,omitempty"`
	// Packages files verification failed count
	PackageFilesVerificationFailedCount int32 `json:"package_files_verification_failed_count,omitempty"`
	// Terraform state versions count
	TerraformStateVersionsCount int32 `json:"terraform_state_versions_count,omitempty"`
	// Terraform state versions checksum total count
	TerraformStateVersionsChecksumTotalCount int32 `json:"terraform_state_versions_checksum_total_count,omitempty"`
	// Terraform state versions checksummed count
	TerraformStateVersionsChecksummedCount int32 `json:"terraform_state_versions_checksummed_count,omitempty"`
	// Terraform state versions checksum failed count
	TerraformStateVersionsChecksumFailedCount int32 `json:"terraform_state_versions_checksum_failed_count,omitempty"`
	// Terraform state versions synced count
	TerraformStateVersionsSyncedCount int32 `json:"terraform_state_versions_synced_count,omitempty"`
	// Terraform state versions failed count
	TerraformStateVersionsFailedCount int32 `json:"terraform_state_versions_failed_count,omitempty"`
	// Terraform state versions registry count
	TerraformStateVersionsRegistryCount int32 `json:"terraform_state_versions_registry_count,omitempty"`
	// Terraform state versions verification total count
	TerraformStateVersionsVerificationTotalCount int32 `json:"terraform_state_versions_verification_total_count,omitempty"`
	// Terraform state versions verified count
	TerraformStateVersionsVerifiedCount int32 `json:"terraform_state_versions_verified_count,omitempty"`
	// Terraform state versions verification failed count
	TerraformStateVersionsVerificationFailedCount int32 `json:"terraform_state_versions_verification_failed_count,omitempty"`
	// Snippet repositories count
	SnippetRepositoriesCount int32 `json:"snippet_repositories_count,omitempty"`
	// Snippet repositories checksum total count
	SnippetRepositoriesChecksumTotalCount int32 `json:"snippet_repositories_checksum_total_count,omitempty"`
	// Snippet repositories checksummed count
	SnippetRepositoriesChecksummedCount int32 `json:"snippet_repositories_checksummed_count,omitempty"`
	// Snippet repositories checksum failed count
	SnippetRepositoriesChecksumFailedCount int32 `json:"snippet_repositories_checksum_failed_count,omitempty"`
	// Snippet repositories synced count
	SnippetRepositoriesSyncedCount int32 `json:"snippet_repositories_synced_count,omitempty"`
	// Snippet repositories failed count
	SnippetRepositoriesFailedCount int32 `json:"snippet_repositories_failed_count,omitempty"`
	// Snippet repositories registry count
	SnippetRepositoriesRegistryCount int32 `json:"snippet_repositories_registry_count,omitempty"`
	// Snippet repositories verification total count
	SnippetRepositoriesVerificationTotalCount int32 `json:"snippet_repositories_verification_total_count,omitempty"`
	// Snippet repositories verified count
	SnippetRepositoriesVerifiedCount int32 `json:"snippet_repositories_verified_count,omitempty"`
	// Snippet repositories verification failed count
	SnippetRepositoriesVerificationFailedCount int32 `json:"snippet_repositories_verification_failed_count,omitempty"`
	// Group wiki repositories count
	GroupWikiRepositoriesCount int32 `json:"group_wiki_repositories_count,omitempty"`
	// Group wiki repositories checksum total count
	GroupWikiRepositoriesChecksumTotalCount int32 `json:"group_wiki_repositories_checksum_total_count,omitempty"`
	// Group wiki repositories checksummed count
	GroupWikiRepositoriesChecksummedCount int32 `json:"group_wiki_repositories_checksummed_count,omitempty"`
	// Group wiki repositories checksum failed count
	GroupWikiRepositoriesChecksumFailedCount int32 `json:"group_wiki_repositories_checksum_failed_count,omitempty"`
	// Group wiki repositories synced count
	GroupWikiRepositoriesSyncedCount int32 `json:"group_wiki_repositories_synced_count,omitempty"`
	// Group wiki repositories failed count
	GroupWikiRepositoriesFailedCount int32 `json:"group_wiki_repositories_failed_count,omitempty"`
	// Group wiki repositories registry count
	GroupWikiRepositoriesRegistryCount int32 `json:"group_wiki_repositories_registry_count,omitempty"`
	// Group wiki repositories verification total count
	GroupWikiRepositoriesVerificationTotalCount int32 `json:"group_wiki_repositories_verification_total_count,omitempty"`
	// Group wiki repositories verified count
	GroupWikiRepositoriesVerifiedCount int32 `json:"group_wiki_repositories_verified_count,omitempty"`
	// Group wiki repositories verification failed count
	GroupWikiRepositoriesVerificationFailedCount int32 `json:"group_wiki_repositories_verification_failed_count,omitempty"`
	// Pipeline artifacts count
	PipelineArtifactsCount int32 `json:"pipeline_artifacts_count,omitempty"`
	// Pipeline artifacts checksum total count
	PipelineArtifactsChecksumTotalCount int32 `json:"pipeline_artifacts_checksum_total_count,omitempty"`
	// Pipeline artifacts checksummed count
	PipelineArtifactsChecksummedCount int32 `json:"pipeline_artifacts_checksummed_count,omitempty"`
	// Pipeline artifacts checksum failed count
	PipelineArtifactsChecksumFailedCount int32 `json:"pipeline_artifacts_checksum_failed_count,omitempty"`
	// Pipeline artifacts synced count
	PipelineArtifactsSyncedCount int32 `json:"pipeline_artifacts_synced_count,omitempty"`
	// Pipeline artifacts failed count
	PipelineArtifactsFailedCount int32 `json:"pipeline_artifacts_failed_count,omitempty"`
	// Pipeline artifacts registry count
	PipelineArtifactsRegistryCount int32 `json:"pipeline_artifacts_registry_count,omitempty"`
	// Pipeline artifacts verification total count
	PipelineArtifactsVerificationTotalCount int32 `json:"pipeline_artifacts_verification_total_count,omitempty"`
	// Pipeline artifacts verified count
	PipelineArtifactsVerifiedCount int32 `json:"pipeline_artifacts_verified_count,omitempty"`
	// Pipeline artifacts verification failed count
	PipelineArtifactsVerificationFailedCount int32 `json:"pipeline_artifacts_verification_failed_count,omitempty"`
	// Pages deployments count
	PagesDeploymentsCount int32 `json:"pages_deployments_count,omitempty"`
	// Pages deployments checksum total count
	PagesDeploymentsChecksumTotalCount int32 `json:"pages_deployments_checksum_total_count,omitempty"`
	// Pages deployments checksummed count
	PagesDeploymentsChecksummedCount int32 `json:"pages_deployments_checksummed_count,omitempty"`
	// Pages deployments checksum failed count
	PagesDeploymentsChecksumFailedCount int32 `json:"pages_deployments_checksum_failed_count,omitempty"`
	// Pages deployments synced count
	PagesDeploymentsSyncedCount int32 `json:"pages_deployments_synced_count,omitempty"`
	// Pages deployments failed count
	PagesDeploymentsFailedCount int32 `json:"pages_deployments_failed_count,omitempty"`
	// Pages deployments registry count
	PagesDeploymentsRegistryCount int32 `json:"pages_deployments_registry_count,omitempty"`
	// Pages deployments verification total count
	PagesDeploymentsVerificationTotalCount int32 `json:"pages_deployments_verification_total_count,omitempty"`
	// Pages deployments verified count
	PagesDeploymentsVerifiedCount int32 `json:"pages_deployments_verified_count,omitempty"`
	// Pages deployments verification failed count
	PagesDeploymentsVerificationFailedCount int32 `json:"pages_deployments_verification_failed_count,omitempty"`
	// Uploads count
	UploadsCount int32 `json:"uploads_count,omitempty"`
	// Uploads checksum total count
	UploadsChecksumTotalCount int32 `json:"uploads_checksum_total_count,omitempty"`
	// Uploads checksummed count
	UploadsChecksummedCount int32 `json:"uploads_checksummed_count,omitempty"`
	// Uploads checksum failed count
	UploadsChecksumFailedCount int32 `json:"uploads_checksum_failed_count,omitempty"`
	// Uploads synced count
	UploadsSyncedCount int32 `json:"uploads_synced_count,omitempty"`
	// Uploads failed count
	UploadsFailedCount int32 `json:"uploads_failed_count,omitempty"`
	// Uploads registry count
	UploadsRegistryCount int32 `json:"uploads_registry_count,omitempty"`
	// Uploads verification total count
	UploadsVerificationTotalCount int32 `json:"uploads_verification_total_count,omitempty"`
	// Uploads verified count
	UploadsVerifiedCount int32 `json:"uploads_verified_count,omitempty"`
	// Uploads verification failed count
	UploadsVerificationFailedCount int32 `json:"uploads_verification_failed_count,omitempty"`
	// Job artifacts count
	JobArtifactsCount int32 `json:"job_artifacts_count,omitempty"`
	// Job artifacts checksum total count
	JobArtifactsChecksumTotalCount int32 `json:"job_artifacts_checksum_total_count,omitempty"`
	// Job artifacts checksummed count
	JobArtifactsChecksummedCount int32 `json:"job_artifacts_checksummed_count,omitempty"`
	// Job artifacts checksum failed count
	JobArtifactsChecksumFailedCount int32 `json:"job_artifacts_checksum_failed_count,omitempty"`
	// Job artifacts synced count
	JobArtifactsSyncedCount int32 `json:"job_artifacts_synced_count,omitempty"`
	// Job artifacts failed count
	JobArtifactsFailedCount int32 `json:"job_artifacts_failed_count,omitempty"`
	// Job artifacts registry count
	JobArtifactsRegistryCount int32 `json:"job_artifacts_registry_count,omitempty"`
	// Job artifacts verification total count
	JobArtifactsVerificationTotalCount int32 `json:"job_artifacts_verification_total_count,omitempty"`
	// Job artifacts verified count
	JobArtifactsVerifiedCount int32 `json:"job_artifacts_verified_count,omitempty"`
	// Job artifacts verification failed count
	JobArtifactsVerificationFailedCount int32 `json:"job_artifacts_verification_failed_count,omitempty"`
	// CI secure files count
	CiSecureFilesCount int32 `json:"ci_secure_files_count,omitempty"`
	// CI secure files checksum total count
	CiSecureFilesChecksumTotalCount int32 `json:"ci_secure_files_checksum_total_count,omitempty"`
	// CI secure files checksummed count
	CiSecureFilesChecksummedCount int32 `json:"ci_secure_files_checksummed_count,omitempty"`
	// CI secure files checksum failed count
	CiSecureFilesChecksumFailedCount int32 `json:"ci_secure_files_checksum_failed_count,omitempty"`
	// CI secure files synced count
	CiSecureFilesSyncedCount int32 `json:"ci_secure_files_synced_count,omitempty"`
	// CI secure files failed count
	CiSecureFilesFailedCount int32 `json:"ci_secure_files_failed_count,omitempty"`
	// CI secure files registry count
	CiSecureFilesRegistryCount int32 `json:"ci_secure_files_registry_count,omitempty"`
	// CI secure files verification total count
	CiSecureFilesVerificationTotalCount int32 `json:"ci_secure_files_verification_total_count,omitempty"`
	// CI secure files verified count
	CiSecureFilesVerifiedCount int32 `json:"ci_secure_files_verified_count,omitempty"`
	// CI secure files verification failed count
	CiSecureFilesVerificationFailedCount int32 `json:"ci_secure_files_verification_failed_count,omitempty"`
	// Container repositories count
	ContainerRepositoriesCount int32 `json:"container_repositories_count,omitempty"`
	// Container repositories checksum total count
	ContainerRepositoriesChecksumTotalCount int32 `json:"container_repositories_checksum_total_count,omitempty"`
	// Container repositories checksummed count
	ContainerRepositoriesChecksummedCount int32 `json:"container_repositories_checksummed_count,omitempty"`
	// Container repositories checksum failed count
	ContainerRepositoriesChecksumFailedCount int32 `json:"container_repositories_checksum_failed_count,omitempty"`
	// Container repositories synced count
	ContainerRepositoriesSyncedCount int32 `json:"container_repositories_synced_count,omitempty"`
	// Container repositories failed count
	ContainerRepositoriesFailedCount int32 `json:"container_repositories_failed_count,omitempty"`
	// Container repositories registry count
	ContainerRepositoriesRegistryCount int32 `json:"container_repositories_registry_count,omitempty"`
	// Container repositories verification total count
	ContainerRepositoriesVerificationTotalCount int32 `json:"container_repositories_verification_total_count,omitempty"`
	// Container repositories verified count
	ContainerRepositoriesVerifiedCount int32 `json:"container_repositories_verified_count,omitempty"`
	// Container repositories verification failed count
	ContainerRepositoriesVerificationFailedCount int32 `json:"container_repositories_verification_failed_count,omitempty"`
	// Git fetch event count weekly
	GitFetchEventCountWeekly int32 `json:"git_fetch_event_count_weekly,omitempty"`
	// Git push event count weekly
	GitPushEventCountWeekly int32 `json:"git_push_event_count_weekly,omitempty"`
	// Proxy remote requests event count weekly
	ProxyRemoteRequestsEventCountWeekly int32 `json:"proxy_remote_requests_event_count_weekly,omitempty"`
	// Proxy local requests event count weekly
	ProxyLocalRequestsEventCountWeekly int32 `json:"proxy_local_requests_event_count_weekly,omitempty"`
}