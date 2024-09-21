# \ProjectsApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4ProjectsId**](ProjectsApi.md#DeleteApiV4ProjectsId) | **Delete** /api/v4/projects/{id} | 
[**DeleteApiV4ProjectsIdArtifacts**](ProjectsApi.md#DeleteApiV4ProjectsIdArtifacts) | **Delete** /api/v4/projects/{id}/artifacts | 
[**DeleteApiV4ProjectsIdCustomAttributesKey**](ProjectsApi.md#DeleteApiV4ProjectsIdCustomAttributesKey) | **Delete** /api/v4/projects/{id}/custom_attributes/{key} | 
[**DeleteApiV4ProjectsIdFork**](ProjectsApi.md#DeleteApiV4ProjectsIdFork) | **Delete** /api/v4/projects/{id}/fork | 
[**DeleteApiV4ProjectsIdHooksHookIdUrlVariablesKey**](ProjectsApi.md#DeleteApiV4ProjectsIdHooksHookIdUrlVariablesKey) | **Delete** /api/v4/projects/{id}/hooks/{hook_id}/url_variables/{key} | 
[**DeleteApiV4ProjectsIdJobsJobIdArtifacts**](ProjectsApi.md#DeleteApiV4ProjectsIdJobsJobIdArtifacts) | **Delete** /api/v4/projects/{id}/jobs/{job_id}/artifacts | Delete the artifacts files from a job
[**DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteId**](ProjectsApi.md#DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteId) | **Delete** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/draft_notes/{draft_note_id} | 
[**DeleteApiV4ProjectsIdPagesDomainsDomain**](ProjectsApi.md#DeleteApiV4ProjectsIdPagesDomainsDomain) | **Delete** /api/v4/projects/{id}/pages/domains/{domain} | 
[**DeleteApiV4ProjectsIdPipelineSchedulesPipelineScheduleId**](ProjectsApi.md#DeleteApiV4ProjectsIdPipelineSchedulesPipelineScheduleId) | **Delete** /api/v4/projects/{id}/pipeline_schedules/{pipeline_schedule_id} | 
[**DeleteApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKey**](ProjectsApi.md#DeleteApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKey) | **Delete** /api/v4/projects/{id}/pipeline_schedules/{pipeline_schedule_id}/variables/{key} | 
[**DeleteApiV4ProjectsIdPipelinesPipelineId**](ProjectsApi.md#DeleteApiV4ProjectsIdPipelinesPipelineId) | **Delete** /api/v4/projects/{id}/pipelines/{pipeline_id} | Deletes a pipeline
[**DeleteApiV4ProjectsIdProtectedBranchesName**](ProjectsApi.md#DeleteApiV4ProjectsIdProtectedBranchesName) | **Delete** /api/v4/projects/{id}/protected_branches/{name} | 
[**DeleteApiV4ProjectsIdRepositoryFilesFilePath**](ProjectsApi.md#DeleteApiV4ProjectsIdRepositoryFilesFilePath) | **Delete** /api/v4/projects/{id}/repository/files/{file_path} | 
[**DeleteApiV4ProjectsIdRunnersRunnerId**](ProjectsApi.md#DeleteApiV4ProjectsIdRunnersRunnerId) | **Delete** /api/v4/projects/{id}/runners/{runner_id} | Disable a project runner from the project
[**DeleteApiV4ProjectsIdShareGroupId**](ProjectsApi.md#DeleteApiV4ProjectsIdShareGroupId) | **Delete** /api/v4/projects/{id}/share/{group_id} | 
[**DeleteApiV4ProjectsIdTriggersTriggerId**](ProjectsApi.md#DeleteApiV4ProjectsIdTriggersTriggerId) | **Delete** /api/v4/projects/{id}/triggers/{trigger_id} | 
[**GetApiV4Projects**](ProjectsApi.md#GetApiV4Projects) | **Get** /api/v4/projects | 
[**GetApiV4ProjectsId**](ProjectsApi.md#GetApiV4ProjectsId) | **Get** /api/v4/projects/{id} | 
[**GetApiV4ProjectsIdAuditEvents**](ProjectsApi.md#GetApiV4ProjectsIdAuditEvents) | **Get** /api/v4/projects/{id}/audit_events | 
[**GetApiV4ProjectsIdAuditEventsAuditEventId**](ProjectsApi.md#GetApiV4ProjectsIdAuditEventsAuditEventId) | **Get** /api/v4/projects/{id}/audit_events/{audit_event_id} | 
[**GetApiV4ProjectsIdCustomAttributes**](ProjectsApi.md#GetApiV4ProjectsIdCustomAttributes) | **Get** /api/v4/projects/{id}/custom_attributes | 
[**GetApiV4ProjectsIdCustomAttributesKey**](ProjectsApi.md#GetApiV4ProjectsIdCustomAttributesKey) | **Get** /api/v4/projects/{id}/custom_attributes/{key} | 
[**GetApiV4ProjectsIdEvents**](ProjectsApi.md#GetApiV4ProjectsIdEvents) | **Get** /api/v4/projects/{id}/events | 
[**GetApiV4ProjectsIdForks**](ProjectsApi.md#GetApiV4ProjectsIdForks) | **Get** /api/v4/projects/{id}/forks | 
[**GetApiV4ProjectsIdGroups**](ProjectsApi.md#GetApiV4ProjectsIdGroups) | **Get** /api/v4/projects/{id}/groups | 
[**GetApiV4ProjectsIdJobs**](ProjectsApi.md#GetApiV4ProjectsIdJobs) | **Get** /api/v4/projects/{id}/jobs | 
[**GetApiV4ProjectsIdJobsArtifactsRefNameDownload**](ProjectsApi.md#GetApiV4ProjectsIdJobsArtifactsRefNameDownload) | **Get** /api/v4/projects/{id}/jobs/artifacts/{ref_name}/download | Download the artifacts archive from a job
[**GetApiV4ProjectsIdJobsArtifactsRefNameRawartifactPath**](ProjectsApi.md#GetApiV4ProjectsIdJobsArtifactsRefNameRawartifactPath) | **Get** /api/v4/projects/{id}/jobs/artifacts/{ref_name}/raw/*artifact_path | Download a specific file from artifacts archive from a ref
[**GetApiV4ProjectsIdJobsJobId**](ProjectsApi.md#GetApiV4ProjectsIdJobsJobId) | **Get** /api/v4/projects/{id}/jobs/{job_id} | 
[**GetApiV4ProjectsIdJobsJobIdArtifacts**](ProjectsApi.md#GetApiV4ProjectsIdJobsJobIdArtifacts) | **Get** /api/v4/projects/{id}/jobs/{job_id}/artifacts | Download the artifacts archive from a job
[**GetApiV4ProjectsIdJobsJobIdArtifactsartifactPath**](ProjectsApi.md#GetApiV4ProjectsIdJobsJobIdArtifactsartifactPath) | **Get** /api/v4/projects/{id}/jobs/{job_id}/artifacts/*artifact_path | Download a specific file from artifacts archive
[**GetApiV4ProjectsIdJobsJobIdTrace**](ProjectsApi.md#GetApiV4ProjectsIdJobsJobIdTrace) | **Get** /api/v4/projects/{id}/jobs/{job_id}/trace | 
[**GetApiV4ProjectsIdLanguages**](ProjectsApi.md#GetApiV4ProjectsIdLanguages) | **Get** /api/v4/projects/{id}/languages | 
[**GetApiV4ProjectsIdMergeRequestsMergeRequestIidApprovalState**](ProjectsApi.md#GetApiV4ProjectsIdMergeRequestsMergeRequestIidApprovalState) | **Get** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/approval_state | 
[**GetApiV4ProjectsIdMergeRequestsMergeRequestIidApprovals**](ProjectsApi.md#GetApiV4ProjectsIdMergeRequestsMergeRequestIidApprovals) | **Get** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/approvals | 
[**GetApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotes**](ProjectsApi.md#GetApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotes) | **Get** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/draft_notes | 
[**GetApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteId**](ProjectsApi.md#GetApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteId) | **Get** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/draft_notes/{draft_note_id} | 
[**GetApiV4ProjectsIdPagesAccess**](ProjectsApi.md#GetApiV4ProjectsIdPagesAccess) | **Get** /api/v4/projects/{id}/pages_access | 
[**GetApiV4ProjectsIdPagesDomainsDomain**](ProjectsApi.md#GetApiV4ProjectsIdPagesDomainsDomain) | **Get** /api/v4/projects/{id}/pages/domains/{domain} | 
[**GetApiV4ProjectsIdPipelineSchedules**](ProjectsApi.md#GetApiV4ProjectsIdPipelineSchedules) | **Get** /api/v4/projects/{id}/pipeline_schedules | 
[**GetApiV4ProjectsIdPipelineSchedulesPipelineScheduleId**](ProjectsApi.md#GetApiV4ProjectsIdPipelineSchedulesPipelineScheduleId) | **Get** /api/v4/projects/{id}/pipeline_schedules/{pipeline_schedule_id} | 
[**GetApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdPipelines**](ProjectsApi.md#GetApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdPipelines) | **Get** /api/v4/projects/{id}/pipeline_schedules/{pipeline_schedule_id}/pipelines | 
[**GetApiV4ProjectsIdPipelines**](ProjectsApi.md#GetApiV4ProjectsIdPipelines) | **Get** /api/v4/projects/{id}/pipelines | Get all Pipelines of the project
[**GetApiV4ProjectsIdPipelinesLatest**](ProjectsApi.md#GetApiV4ProjectsIdPipelinesLatest) | **Get** /api/v4/projects/{id}/pipelines/latest | Gets the latest pipeline for the project branch
[**GetApiV4ProjectsIdPipelinesPipelineId**](ProjectsApi.md#GetApiV4ProjectsIdPipelinesPipelineId) | **Get** /api/v4/projects/{id}/pipelines/{pipeline_id} | Gets a specific pipeline for the project
[**GetApiV4ProjectsIdPipelinesPipelineIdBridges**](ProjectsApi.md#GetApiV4ProjectsIdPipelinesPipelineIdBridges) | **Get** /api/v4/projects/{id}/pipelines/{pipeline_id}/bridges | 
[**GetApiV4ProjectsIdPipelinesPipelineIdJobs**](ProjectsApi.md#GetApiV4ProjectsIdPipelinesPipelineIdJobs) | **Get** /api/v4/projects/{id}/pipelines/{pipeline_id}/jobs | 
[**GetApiV4ProjectsIdPipelinesPipelineIdTestReport**](ProjectsApi.md#GetApiV4ProjectsIdPipelinesPipelineIdTestReport) | **Get** /api/v4/projects/{id}/pipelines/{pipeline_id}/test_report | Gets the test report for a given pipeline
[**GetApiV4ProjectsIdPipelinesPipelineIdTestReportSummary**](ProjectsApi.md#GetApiV4ProjectsIdPipelinesPipelineIdTestReportSummary) | **Get** /api/v4/projects/{id}/pipelines/{pipeline_id}/test_report_summary | Gets the test report summary for a given pipeline
[**GetApiV4ProjectsIdPipelinesPipelineIdVariables**](ProjectsApi.md#GetApiV4ProjectsIdPipelinesPipelineIdVariables) | **Get** /api/v4/projects/{id}/pipelines/{pipeline_id}/variables | Gets the variables for a given pipeline
[**GetApiV4ProjectsIdProtectedBranches**](ProjectsApi.md#GetApiV4ProjectsIdProtectedBranches) | **Get** /api/v4/projects/{id}/protected_branches | 
[**GetApiV4ProjectsIdProtectedBranchesName**](ProjectsApi.md#GetApiV4ProjectsIdProtectedBranchesName) | **Get** /api/v4/projects/{id}/protected_branches/{name} | 
[**GetApiV4ProjectsIdRepositoryArchive**](ProjectsApi.md#GetApiV4ProjectsIdRepositoryArchive) | **Get** /api/v4/projects/{id}/repository/archive | 
[**GetApiV4ProjectsIdRepositoryBlobsSha**](ProjectsApi.md#GetApiV4ProjectsIdRepositoryBlobsSha) | **Get** /api/v4/projects/{id}/repository/blobs/{sha} | 
[**GetApiV4ProjectsIdRepositoryBlobsShaRaw**](ProjectsApi.md#GetApiV4ProjectsIdRepositoryBlobsShaRaw) | **Get** /api/v4/projects/{id}/repository/blobs/{sha}/raw | 
[**GetApiV4ProjectsIdRepositoryChangelog**](ProjectsApi.md#GetApiV4ProjectsIdRepositoryChangelog) | **Get** /api/v4/projects/{id}/repository/changelog | Generates a changelog section for a release and returns it
[**GetApiV4ProjectsIdRepositoryCommitsShaStatuses**](ProjectsApi.md#GetApiV4ProjectsIdRepositoryCommitsShaStatuses) | **Get** /api/v4/projects/{id}/repository/commits/{sha}/statuses | 
[**GetApiV4ProjectsIdRepositoryCompare**](ProjectsApi.md#GetApiV4ProjectsIdRepositoryCompare) | **Get** /api/v4/projects/{id}/repository/compare | 
[**GetApiV4ProjectsIdRepositoryContributors**](ProjectsApi.md#GetApiV4ProjectsIdRepositoryContributors) | **Get** /api/v4/projects/{id}/repository/contributors | 
[**GetApiV4ProjectsIdRepositoryFilesFilePath**](ProjectsApi.md#GetApiV4ProjectsIdRepositoryFilesFilePath) | **Get** /api/v4/projects/{id}/repository/files/{file_path} | 
[**GetApiV4ProjectsIdRepositoryFilesFilePathBlame**](ProjectsApi.md#GetApiV4ProjectsIdRepositoryFilesFilePathBlame) | **Get** /api/v4/projects/{id}/repository/files/{file_path}/blame | 
[**GetApiV4ProjectsIdRepositoryFilesFilePathRaw**](ProjectsApi.md#GetApiV4ProjectsIdRepositoryFilesFilePathRaw) | **Get** /api/v4/projects/{id}/repository/files/{file_path}/raw | 
[**GetApiV4ProjectsIdRepositoryMergeBase**](ProjectsApi.md#GetApiV4ProjectsIdRepositoryMergeBase) | **Get** /api/v4/projects/{id}/repository/merge_base | 
[**GetApiV4ProjectsIdRepositoryStorageMoves**](ProjectsApi.md#GetApiV4ProjectsIdRepositoryStorageMoves) | **Get** /api/v4/projects/{id}/repository_storage_moves | Get a list of all project repository storage moves
[**GetApiV4ProjectsIdRepositoryStorageMovesRepositoryStorageMoveId**](ProjectsApi.md#GetApiV4ProjectsIdRepositoryStorageMovesRepositoryStorageMoveId) | **Get** /api/v4/projects/{id}/repository_storage_moves/{repository_storage_move_id} | Get a project repository storage move
[**GetApiV4ProjectsIdRepositoryTree**](ProjectsApi.md#GetApiV4ProjectsIdRepositoryTree) | **Get** /api/v4/projects/{id}/repository/tree | 
[**GetApiV4ProjectsIdRunners**](ProjectsApi.md#GetApiV4ProjectsIdRunners) | **Get** /api/v4/projects/{id}/runners | List project&#39;s runners
[**GetApiV4ProjectsIdShareLocations**](ProjectsApi.md#GetApiV4ProjectsIdShareLocations) | **Get** /api/v4/projects/{id}/share_locations | 
[**GetApiV4ProjectsIdSnapshot**](ProjectsApi.md#GetApiV4ProjectsIdSnapshot) | **Get** /api/v4/projects/{id}/snapshot | Download a (possibly inconsistent) snapshot of a repository
[**GetApiV4ProjectsIdStarrers**](ProjectsApi.md#GetApiV4ProjectsIdStarrers) | **Get** /api/v4/projects/{id}/starrers | 
[**GetApiV4ProjectsIdStatistics**](ProjectsApi.md#GetApiV4ProjectsIdStatistics) | **Get** /api/v4/projects/{id}/statistics | 
[**GetApiV4ProjectsIdStorage**](ProjectsApi.md#GetApiV4ProjectsIdStorage) | **Get** /api/v4/projects/{id}/storage | 
[**GetApiV4ProjectsIdTemplatesType**](ProjectsApi.md#GetApiV4ProjectsIdTemplatesType) | **Get** /api/v4/projects/{id}/templates/{type} | Get a list of templates available to this project
[**GetApiV4ProjectsIdTemplatesTypeName**](ProjectsApi.md#GetApiV4ProjectsIdTemplatesTypeName) | **Get** /api/v4/projects/{id}/templates/{type}/{name} | Download a template available to this project
[**GetApiV4ProjectsIdTransferLocations**](ProjectsApi.md#GetApiV4ProjectsIdTransferLocations) | **Get** /api/v4/projects/{id}/transfer_locations | 
[**GetApiV4ProjectsIdTriggers**](ProjectsApi.md#GetApiV4ProjectsIdTriggers) | **Get** /api/v4/projects/{id}/triggers | 
[**GetApiV4ProjectsIdTriggersTriggerId**](ProjectsApi.md#GetApiV4ProjectsIdTriggersTriggerId) | **Get** /api/v4/projects/{id}/triggers/{trigger_id} | 
[**GetApiV4ProjectsIdUsers**](ProjectsApi.md#GetApiV4ProjectsIdUsers) | **Get** /api/v4/projects/{id}/users | 
[**GetApiV4UsersUserIdContributedProjects**](ProjectsApi.md#GetApiV4UsersUserIdContributedProjects) | **Get** /api/v4/users/{user_id}/contributed_projects | 
[**GetApiV4UsersUserIdProjects**](ProjectsApi.md#GetApiV4UsersUserIdProjects) | **Get** /api/v4/users/{user_id}/projects | 
[**GetApiV4UsersUserIdStarredProjects**](ProjectsApi.md#GetApiV4UsersUserIdStarredProjects) | **Get** /api/v4/users/{user_id}/starred_projects | 
[**HeadApiV4ProjectsIdRepositoryFilesFilePath**](ProjectsApi.md#HeadApiV4ProjectsIdRepositoryFilesFilePath) | **Head** /api/v4/projects/{id}/repository/files/{file_path} | 
[**HeadApiV4ProjectsIdRepositoryFilesFilePathBlame**](ProjectsApi.md#HeadApiV4ProjectsIdRepositoryFilesFilePathBlame) | **Head** /api/v4/projects/{id}/repository/files/{file_path}/blame | 
[**PatchApiV4ProjectsIdProtectedBranchesName**](ProjectsApi.md#PatchApiV4ProjectsIdProtectedBranchesName) | **Patch** /api/v4/projects/{id}/protected_branches/{name} | 
[**PostApiV4Projects**](ProjectsApi.md#PostApiV4Projects) | **Post** /api/v4/projects | 
[**PostApiV4ProjectsIdArchive**](ProjectsApi.md#PostApiV4ProjectsIdArchive) | **Post** /api/v4/projects/{id}/archive | 
[**PostApiV4ProjectsIdCreateCiConfig**](ProjectsApi.md#PostApiV4ProjectsIdCreateCiConfig) | **Post** /api/v4/projects/{id}/create_ci_config | 
[**PostApiV4ProjectsIdFork**](ProjectsApi.md#PostApiV4ProjectsIdFork) | **Post** /api/v4/projects/{id}/fork | 
[**PostApiV4ProjectsIdForkForkedFromId**](ProjectsApi.md#PostApiV4ProjectsIdForkForkedFromId) | **Post** /api/v4/projects/{id}/fork/{forked_from_id} | 
[**PostApiV4ProjectsIdHooksHookIdTestTrigger**](ProjectsApi.md#PostApiV4ProjectsIdHooksHookIdTestTrigger) | **Post** /api/v4/projects/{id}/hooks/{hook_id}/test/{trigger} | Triggers a hook test
[**PostApiV4ProjectsIdHousekeeping**](ProjectsApi.md#PostApiV4ProjectsIdHousekeeping) | **Post** /api/v4/projects/{id}/housekeeping | Start the housekeeping task for a project
[**PostApiV4ProjectsIdImportProjectMembersProjectId**](ProjectsApi.md#PostApiV4ProjectsIdImportProjectMembersProjectId) | **Post** /api/v4/projects/{id}/import_project_members/{project_id} | Import members from another project
[**PostApiV4ProjectsIdJobsJobIdArtifactsKeep**](ProjectsApi.md#PostApiV4ProjectsIdJobsJobIdArtifactsKeep) | **Post** /api/v4/projects/{id}/jobs/{job_id}/artifacts/keep | 
[**PostApiV4ProjectsIdJobsJobIdCancel**](ProjectsApi.md#PostApiV4ProjectsIdJobsJobIdCancel) | **Post** /api/v4/projects/{id}/jobs/{job_id}/cancel | 
[**PostApiV4ProjectsIdJobsJobIdErase**](ProjectsApi.md#PostApiV4ProjectsIdJobsJobIdErase) | **Post** /api/v4/projects/{id}/jobs/{job_id}/erase | 
[**PostApiV4ProjectsIdJobsJobIdPlay**](ProjectsApi.md#PostApiV4ProjectsIdJobsJobIdPlay) | **Post** /api/v4/projects/{id}/jobs/{job_id}/play | Trigger an actionable job (manual, delayed, etc)
[**PostApiV4ProjectsIdJobsJobIdRetry**](ProjectsApi.md#PostApiV4ProjectsIdJobsJobIdRetry) | **Post** /api/v4/projects/{id}/jobs/{job_id}/retry | 
[**PostApiV4ProjectsIdMergeRequestsMergeRequestIidApprovals**](ProjectsApi.md#PostApiV4ProjectsIdMergeRequestsMergeRequestIidApprovals) | **Post** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/approvals | Deprecated in 16.0: Use the merge request approvals API instead. Change approval-related configuration
[**PostApiV4ProjectsIdMergeRequestsMergeRequestIidApprove**](ProjectsApi.md#PostApiV4ProjectsIdMergeRequestsMergeRequestIidApprove) | **Post** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/approve | 
[**PostApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotes**](ProjectsApi.md#PostApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotes) | **Post** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/draft_notes | 
[**PostApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesBulkPublish**](ProjectsApi.md#PostApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesBulkPublish) | **Post** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/draft_notes/bulk_publish | 
[**PostApiV4ProjectsIdMergeRequestsMergeRequestIidUnapprove**](ProjectsApi.md#PostApiV4ProjectsIdMergeRequestsMergeRequestIidUnapprove) | **Post** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/unapprove | 
[**PostApiV4ProjectsIdPagesDomains**](ProjectsApi.md#PostApiV4ProjectsIdPagesDomains) | **Post** /api/v4/projects/{id}/pages/domains | 
[**PostApiV4ProjectsIdPipeline**](ProjectsApi.md#PostApiV4ProjectsIdPipeline) | **Post** /api/v4/projects/{id}/pipeline | Create a new pipeline
[**PostApiV4ProjectsIdPipelineSchedules**](ProjectsApi.md#PostApiV4ProjectsIdPipelineSchedules) | **Post** /api/v4/projects/{id}/pipeline_schedules | 
[**PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdPlay**](ProjectsApi.md#PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdPlay) | **Post** /api/v4/projects/{id}/pipeline_schedules/{pipeline_schedule_id}/play | Play a scheduled pipeline immediately
[**PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdTakeOwnership**](ProjectsApi.md#PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdTakeOwnership) | **Post** /api/v4/projects/{id}/pipeline_schedules/{pipeline_schedule_id}/take_ownership | 
[**PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariables**](ProjectsApi.md#PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariables) | **Post** /api/v4/projects/{id}/pipeline_schedules/{pipeline_schedule_id}/variables | 
[**PostApiV4ProjectsIdPipelinesPipelineIdCancel**](ProjectsApi.md#PostApiV4ProjectsIdPipelinesPipelineIdCancel) | **Post** /api/v4/projects/{id}/pipelines/{pipeline_id}/cancel | Cancel all builds in the pipeline
[**PostApiV4ProjectsIdPipelinesPipelineIdRetry**](ProjectsApi.md#PostApiV4ProjectsIdPipelinesPipelineIdRetry) | **Post** /api/v4/projects/{id}/pipelines/{pipeline_id}/retry | Retry builds in the pipeline
[**PostApiV4ProjectsIdProtectedBranches**](ProjectsApi.md#PostApiV4ProjectsIdProtectedBranches) | **Post** /api/v4/projects/{id}/protected_branches | 
[**PostApiV4ProjectsIdRefReftriggerPipeline**](ProjectsApi.md#PostApiV4ProjectsIdRefReftriggerPipeline) | **Post** /api/v4/projects/{id}/(ref/{ref}/)trigger/pipeline | 
[**PostApiV4ProjectsIdRepositoryChangelog**](ProjectsApi.md#PostApiV4ProjectsIdRepositoryChangelog) | **Post** /api/v4/projects/{id}/repository/changelog | Generates a changelog section for a release and commits it in a changelog file
[**PostApiV4ProjectsIdRepositoryFilesFilePath**](ProjectsApi.md#PostApiV4ProjectsIdRepositoryFilesFilePath) | **Post** /api/v4/projects/{id}/repository/files/{file_path} | 
[**PostApiV4ProjectsIdRepositorySize**](ProjectsApi.md#PostApiV4ProjectsIdRepositorySize) | **Post** /api/v4/projects/{id}/repository_size | Start a task to recalculate repository size for a project
[**PostApiV4ProjectsIdRepositoryStorageMoves**](ProjectsApi.md#PostApiV4ProjectsIdRepositoryStorageMoves) | **Post** /api/v4/projects/{id}/repository_storage_moves | Schedule a project repository storage move
[**PostApiV4ProjectsIdRestore**](ProjectsApi.md#PostApiV4ProjectsIdRestore) | **Post** /api/v4/projects/{id}/restore | 
[**PostApiV4ProjectsIdRunners**](ProjectsApi.md#PostApiV4ProjectsIdRunners) | **Post** /api/v4/projects/{id}/runners | Enable a runner in project
[**PostApiV4ProjectsIdRunnersResetRegistrationToken**](ProjectsApi.md#PostApiV4ProjectsIdRunnersResetRegistrationToken) | **Post** /api/v4/projects/{id}/runners/reset_registration_token | Reset the runner registration token for a project
[**PostApiV4ProjectsIdShare**](ProjectsApi.md#PostApiV4ProjectsIdShare) | **Post** /api/v4/projects/{id}/share | 
[**PostApiV4ProjectsIdStar**](ProjectsApi.md#PostApiV4ProjectsIdStar) | **Post** /api/v4/projects/{id}/star | 
[**PostApiV4ProjectsIdStatusesSha**](ProjectsApi.md#PostApiV4ProjectsIdStatusesSha) | **Post** /api/v4/projects/{id}/statuses/{sha} | 
[**PostApiV4ProjectsIdTriggers**](ProjectsApi.md#PostApiV4ProjectsIdTriggers) | **Post** /api/v4/projects/{id}/triggers | 
[**PostApiV4ProjectsIdUnarchive**](ProjectsApi.md#PostApiV4ProjectsIdUnarchive) | **Post** /api/v4/projects/{id}/unarchive | 
[**PostApiV4ProjectsIdUnstar**](ProjectsApi.md#PostApiV4ProjectsIdUnstar) | **Post** /api/v4/projects/{id}/unstar | 
[**PostApiV4ProjectsIdUploads**](ProjectsApi.md#PostApiV4ProjectsIdUploads) | **Post** /api/v4/projects/{id}/uploads | 
[**PostApiV4ProjectsIdUploadsAuthorize**](ProjectsApi.md#PostApiV4ProjectsIdUploadsAuthorize) | **Post** /api/v4/projects/{id}/uploads/authorize | Workhorse authorize the file upload
[**PostApiV4ProjectsUserUserId**](ProjectsApi.md#PostApiV4ProjectsUserUserId) | **Post** /api/v4/projects/user/{user_id} | 
[**PutApiV4ProjectsId**](ProjectsApi.md#PutApiV4ProjectsId) | **Put** /api/v4/projects/{id} | 
[**PutApiV4ProjectsIdCustomAttributesKey**](ProjectsApi.md#PutApiV4ProjectsIdCustomAttributesKey) | **Put** /api/v4/projects/{id}/custom_attributes/{key} | 
[**PutApiV4ProjectsIdHooksHookIdUrlVariablesKey**](ProjectsApi.md#PutApiV4ProjectsIdHooksHookIdUrlVariablesKey) | **Put** /api/v4/projects/{id}/hooks/{hook_id}/url_variables/{key} | 
[**PutApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteId**](ProjectsApi.md#PutApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteId) | **Put** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/draft_notes/{draft_note_id} | 
[**PutApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdPublish**](ProjectsApi.md#PutApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdPublish) | **Put** /api/v4/projects/{id}/merge_requests/{merge_request_iid}/draft_notes/{draft_note_id}/publish | 
[**PutApiV4ProjectsIdPagesDomainsDomain**](ProjectsApi.md#PutApiV4ProjectsIdPagesDomainsDomain) | **Put** /api/v4/projects/{id}/pages/domains/{domain} | 
[**PutApiV4ProjectsIdPipelineSchedulesPipelineScheduleId**](ProjectsApi.md#PutApiV4ProjectsIdPipelineSchedulesPipelineScheduleId) | **Put** /api/v4/projects/{id}/pipeline_schedules/{pipeline_schedule_id} | 
[**PutApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKey**](ProjectsApi.md#PutApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKey) | **Put** /api/v4/projects/{id}/pipeline_schedules/{pipeline_schedule_id}/variables/{key} | 
[**PutApiV4ProjectsIdPipelinesPipelineIdMetadata**](ProjectsApi.md#PutApiV4ProjectsIdPipelinesPipelineIdMetadata) | **Put** /api/v4/projects/{id}/pipelines/{pipeline_id}/metadata | Updates pipeline metadata
[**PutApiV4ProjectsIdRepositoryFilesFilePath**](ProjectsApi.md#PutApiV4ProjectsIdRepositoryFilesFilePath) | **Put** /api/v4/projects/{id}/repository/files/{file_path} | 
[**PutApiV4ProjectsIdRepositorySubmodulesSubmodule**](ProjectsApi.md#PutApiV4ProjectsIdRepositorySubmodulesSubmodule) | **Put** /api/v4/projects/{id}/repository/submodules/{submodule} | 
[**PutApiV4ProjectsIdTransfer**](ProjectsApi.md#PutApiV4ProjectsIdTransfer) | **Put** /api/v4/projects/{id}/transfer | 
[**PutApiV4ProjectsIdTriggersTriggerId**](ProjectsApi.md#PutApiV4ProjectsIdTriggersTriggerId) | **Put** /api/v4/projects/{id}/triggers/{trigger_id} | 


# **DeleteApiV4ProjectsId**
> DeleteApiV4ProjectsId(ctx, id)


Delete a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdArtifacts**
> DeleteApiV4ProjectsIdArtifacts(ctx, id)


Expire the artifacts files from a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdCustomAttributesKey**
> DeleteApiV4ProjectsIdCustomAttributesKey(ctx, key, id)


Delete a custom attribute on a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| The key of the custom attribute | 
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdFork**
> DeleteApiV4ProjectsIdFork(ctx, id)


Remove a forked_from relationship

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdHooksHookIdUrlVariablesKey**
> DeleteApiV4ProjectsIdHooksHookIdUrlVariablesKey(ctx, hookId, key, id)


Un-Set a url variable

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hookId** | **int32**| The ID of the hook | 
  **key** | **string**| The key of the variable | 
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdJobsJobIdArtifacts**
> DeleteApiV4ProjectsIdJobsJobIdArtifacts(ctx, id, jobId)
Delete the artifacts files from a job

This feature was introduced in GitLab 11.9

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **jobId** | **int32**| The ID of a job | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteId**
> ApiEntitiesDraftNote DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteId(ctx, id, mergeRequestIid, draftNoteId)


Delete a draft note

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a project | 
  **mergeRequestIid** | **int32**| The ID of a merge request | 
  **draftNoteId** | **int32**| The ID of a draft note | 

### Return type

[**ApiEntitiesDraftNote**](API_Entities_DraftNote.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdPagesDomainsDomain**
> DeleteApiV4ProjectsIdPagesDomainsDomain(ctx, id, domain)


Delete a pages domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **domain** | **string**| The domain | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdPipelineSchedulesPipelineScheduleId**
> DeleteApiV4ProjectsIdPipelineSchedulesPipelineScheduleId(ctx, id, pipelineScheduleId)


Delete a pipeline schedule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **pipelineScheduleId** | **int32**| The pipeline schedule id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKey**
> ApiEntitiesCiVariable DeleteApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKey(ctx, id, pipelineScheduleId, key)


Delete a pipeline schedule variable

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **pipelineScheduleId** | **int32**| The pipeline schedule id | 
  **key** | **string**| The key of the variable | 

### Return type

[**ApiEntitiesCiVariable**](API_Entities_Ci_Variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdPipelinesPipelineId**
> DeleteApiV4ProjectsIdPipelinesPipelineId(ctx, id, pipelineId)
Deletes a pipeline

This feature was introduced in GitLab 11.6

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID or URL-encoded path | 
  **pipelineId** | **int32**| The pipeline ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdProtectedBranchesName**
> DeleteApiV4ProjectsIdProtectedBranchesName(ctx, id, name)


Unprotect a single branch

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **name** | **string**| The name of the protected branch | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdRepositoryFilesFilePath**
> DeleteApiV4ProjectsIdRepositoryFilesFilePath(ctx, id, filePath, branch, commitMessage, optional)


Delete an existing file in repository

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID | 
  **filePath** | **string**| The url encoded path to the file. | 
  **branch** | **string**| Name of the branch to commit into. To create a new branch, also provide &#x60;start_branch&#x60;. | 
  **commitMessage** | **string**| Commit message | 
 **optional** | ***ProjectsApiDeleteApiV4ProjectsIdRepositoryFilesFilePathOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiDeleteApiV4ProjectsIdRepositoryFilesFilePathOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **startBranch** | **optional.String**| Name of the branch to start the new commit from | 
 **authorEmail** | **optional.String**| The email of the author | 
 **authorName** | **optional.String**| The name of the author | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdRunnersRunnerId**
> ApiEntitiesCiRunner DeleteApiV4ProjectsIdRunnersRunnerId(ctx, id, runnerId)
Disable a project runner from the project

It works only if the project isn't the only project associated with the specified runner. If so, an error is returned. Use the call to delete a runner instead.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **runnerId** | **int32**| The ID of a runner | 

### Return type

[**ApiEntitiesCiRunner**](API_Entities_Ci_Runner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdShareGroupId**
> DeleteApiV4ProjectsIdShareGroupId(ctx, id, groupId)


Remove a group share

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **groupId** | **int32**| The ID of the group | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4ProjectsIdTriggersTriggerId**
> DeleteApiV4ProjectsIdTriggersTriggerId(ctx, id, triggerId)


Delete a trigger token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **triggerId** | **int32**| The trigger token ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4Projects**
> []ApiEntitiesBasicProjectDetails GetApiV4Projects(ctx, optional)


Get a list of visible projects for authenticated user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProjectsApiGetApiV4ProjectsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4ProjectsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderBy** | **optional.String**| Return projects ordered by field. storage_size, repository_size, wiki_size, packages_size are only available to admins. Similarity is available when searching and is limited to projects the user has access to. | [default to created_at]
 **sort** | **optional.String**| Return projects sorted in ascending and descending order | [default to desc]
 **archived** | **optional.Bool**| Limit by archived status | 
 **visibility** | **optional.String**| Limit by visibility | 
 **search** | **optional.String**| Return list of projects matching the search criteria | 
 **searchNamespaces** | **optional.Bool**| Include ancestor namespaces when matching search criteria | 
 **owned** | **optional.Bool**| Limit by owned by authenticated user | [default to false]
 **starred** | **optional.Bool**| Limit by starred status | [default to false]
 **imported** | **optional.Bool**| Limit by imported by authenticated user | [default to false]
 **membership** | **optional.Bool**| Limit by projects that the current user is a member of | [default to false]
 **withIssuesEnabled** | **optional.Bool**| Limit by enabled issues feature | [default to false]
 **withMergeRequestsEnabled** | **optional.Bool**| Limit by enabled merge requests feature | [default to false]
 **withProgrammingLanguage** | **optional.String**| Limit to repositories which use the given programming language | 
 **minAccessLevel** | **optional.Int32**| Limit by minimum access level of authenticated user | 
 **idAfter** | **optional.Int32**| Limit results to projects with IDs greater than the specified ID | 
 **idBefore** | **optional.Int32**| Limit results to projects with IDs less than the specified ID | 
 **lastActivityAfter** | **optional.Time**| Limit results to projects with last_activity after specified time. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ | 
 **lastActivityBefore** | **optional.Time**| Limit results to projects with last_activity before specified time. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ | 
 **repositoryStorage** | **optional.String**| Which storage shard the repository is on. Available only to admins | 
 **topic** | [**optional.Interface of []string**](string.md)| Comma-separated list of topics. Limit results to projects having all topics | 
 **topicId** | **optional.Int32**| Limit results to projects with the assigned topic given by the topic ID | 
 **updatedBefore** | **optional.Time**| Return projects updated before the specified datetime. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ | 
 **updatedAfter** | **optional.Time**| Return projects updated after the specified datetime. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ | 
 **includePendingDelete** | **optional.Bool**| Include projects in pending delete state. Can only be set by admins | 
 **wikiChecksumFailed** | **optional.Bool**| Limit by projects where wiki checksum is failed | [default to false]
 **repositoryChecksumFailed** | **optional.Bool**| Limit by projects where repository checksum is failed | [default to false]
 **includeHidden** | **optional.Bool**| Include hidden projects. Can only be set by admins | [default to false]
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **simple** | **optional.Bool**| Return only the ID, URL, name, and path of each project | [default to false]
 **statistics** | **optional.Bool**| Include project statistics | [default to false]
 **withCustomAttributes** | **optional.Bool**| Include custom attributes in the response | [default to false]

### Return type

[**[]ApiEntitiesBasicProjectDetails**](API_Entities_BasicProjectDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsId**
> ApiEntitiesProjectWithAccess GetApiV4ProjectsId(ctx, id, optional)


Get a single project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
 **optional** | ***ProjectsApiGetApiV4ProjectsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4ProjectsIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **statistics** | **optional.Bool**| Include project statistics | [default to false]
 **withCustomAttributes** | **optional.Bool**| Include custom attributes in the response | [default to false]
 **license** | **optional.Bool**| Include project license data | [default to false]

### Return type

[**ApiEntitiesProjectWithAccess**](API_Entities_ProjectWithAccess.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdAuditEvents**
> []EeApiEntitiesAuditEvent GetApiV4ProjectsIdAuditEvents(ctx, id, optional)


Get a list of audit events in this project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***ProjectsApiGetApiV4ProjectsIdAuditEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4ProjectsIdAuditEventsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createdAfter** | **optional.Time**| Return audit events created after the specified time | 
 **createdBefore** | **optional.Time**| Return audit events created before the specified time | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]EeApiEntitiesAuditEvent**](EE_API_Entities_AuditEvent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdAuditEventsAuditEventId**
> EeApiEntitiesAuditEvent GetApiV4ProjectsIdAuditEventsAuditEventId(ctx, auditEventId, id)


Get a specific audit event in this project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **auditEventId** | **int32**| The ID of the audit event | 
  **id** | **int32**|  | 

### Return type

[**EeApiEntitiesAuditEvent**](EE_API_Entities_AuditEvent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdCustomAttributes**
> ApiEntitiesCustomAttribute GetApiV4ProjectsIdCustomAttributes(ctx, id)


Get all custom attributes on a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**ApiEntitiesCustomAttribute**](API_Entities_CustomAttribute.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdCustomAttributesKey**
> ApiEntitiesCustomAttribute GetApiV4ProjectsIdCustomAttributesKey(ctx, key, id)


Get a custom attribute on a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| The key of the custom attribute | 
  **id** | **int32**|  | 

### Return type

[**ApiEntitiesCustomAttribute**](API_Entities_CustomAttribute.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdEvents**
> ApiEntitiesEvent GetApiV4ProjectsIdEvents(ctx, id, optional)


List a project's visible events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
 **optional** | ***ProjectsApiGetApiV4ProjectsIdEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4ProjectsIdEventsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **action** | **optional.String**| Event action to filter on | 
 **targetType** | **optional.String**| Event target type to filter on | 
 **before** | **optional.String**| Include only events created before this date | 
 **after** | **optional.String**| Include only events created after this date | 
 **sort** | **optional.String**| Return events sorted in ascending and descending order | [default to desc]
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**ApiEntitiesEvent**](API_Entities_Event.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdForks**
> []ApiEntitiesProject GetApiV4ProjectsIdForks(ctx, id, optional)


List forks of this project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
 **optional** | ***ProjectsApiGetApiV4ProjectsIdForksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4ProjectsIdForksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderBy** | **optional.String**| Return projects ordered by field. storage_size, repository_size, wiki_size, packages_size are only available to admins. Similarity is available when searching and is limited to projects the user has access to. | [default to created_at]
 **sort** | **optional.String**| Return projects sorted in ascending and descending order | [default to desc]
 **archived** | **optional.Bool**| Limit by archived status | 
 **visibility** | **optional.String**| Limit by visibility | 
 **search** | **optional.String**| Return list of projects matching the search criteria | 
 **searchNamespaces** | **optional.Bool**| Include ancestor namespaces when matching search criteria | 
 **owned** | **optional.Bool**| Limit by owned by authenticated user | [default to false]
 **starred** | **optional.Bool**| Limit by starred status | [default to false]
 **imported** | **optional.Bool**| Limit by imported by authenticated user | [default to false]
 **membership** | **optional.Bool**| Limit by projects that the current user is a member of | [default to false]
 **withIssuesEnabled** | **optional.Bool**| Limit by enabled issues feature | [default to false]
 **withMergeRequestsEnabled** | **optional.Bool**| Limit by enabled merge requests feature | [default to false]
 **withProgrammingLanguage** | **optional.String**| Limit to repositories which use the given programming language | 
 **minAccessLevel** | **optional.Int32**| Limit by minimum access level of authenticated user | 
 **idAfter** | **optional.Int32**| Limit results to projects with IDs greater than the specified ID | 
 **idBefore** | **optional.Int32**| Limit results to projects with IDs less than the specified ID | 
 **lastActivityAfter** | **optional.Time**| Limit results to projects with last_activity after specified time. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ | 
 **lastActivityBefore** | **optional.Time**| Limit results to projects with last_activity before specified time. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ | 
 **repositoryStorage** | **optional.String**| Which storage shard the repository is on. Available only to admins | 
 **topic** | [**optional.Interface of []string**](string.md)| Comma-separated list of topics. Limit results to projects having all topics | 
 **topicId** | **optional.Int32**| Limit results to projects with the assigned topic given by the topic ID | 
 **updatedBefore** | **optional.Time**| Return projects updated before the specified datetime. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ | 
 **updatedAfter** | **optional.Time**| Return projects updated after the specified datetime. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ | 
 **includePendingDelete** | **optional.Bool**| Include projects in pending delete state. Can only be set by admins | 
 **wikiChecksumFailed** | **optional.Bool**| Limit by projects where wiki checksum is failed | [default to false]
 **repositoryChecksumFailed** | **optional.Bool**| Limit by projects where repository checksum is failed | [default to false]
 **includeHidden** | **optional.Bool**| Include hidden projects. Can only be set by admins | [default to false]
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **simple** | **optional.Bool**| Return only the ID, URL, name, and path of each project | [default to false]
 **withCustomAttributes** | **optional.Bool**| Include custom attributes in the response | [default to false]

### Return type

[**[]ApiEntitiesProject**](API_Entities_Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdGroups**
> []ApiEntitiesPublicGroupDetails GetApiV4ProjectsIdGroups(ctx, id, optional)


Get ancestor and shared groups for a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
 **optional** | ***ProjectsApiGetApiV4ProjectsIdGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4ProjectsIdGroupsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| Return list of groups matching the search criteria | 
 **skipGroups** | [**optional.Interface of []int32**](int32.md)| Array of group ids to exclude from list | 
 **withShared** | **optional.Bool**| Include shared groups | [default to false]
 **sharedVisibleOnly** | **optional.Bool**| Limit to shared groups user has access to | [default to false]
 **sharedMinAccessLevel** | **optional.Int32**| Limit returned shared groups by minimum access level to the project | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesPublicGroupDetails**](API_Entities_PublicGroupDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdJobs**
> []ApiEntitiesCiJob GetApiV4ProjectsIdJobs(ctx, id, optional)


Get a projects jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
 **optional** | ***ProjectsApiGetApiV4ProjectsIdJobsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4ProjectsIdJobsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scope** | [**optional.Interface of []string**](string.md)| The scope of builds to show | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesCiJob**](API_Entities_Ci_Job.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdJobsArtifactsRefNameDownload**
> GetApiV4ProjectsIdJobsArtifactsRefNameDownload(ctx, id, refName, job, optional)
Download the artifacts archive from a job

This feature was introduced in GitLab 8.10

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **refName** | **string**| Branch or tag name in repository. &#x60;HEAD&#x60; or &#x60;SHA&#x60; references are not supported. | 
  **job** | **string**| The name of the job. | 
 **optional** | ***ProjectsApiGetApiV4ProjectsIdJobsArtifactsRefNameDownloadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4ProjectsIdJobsArtifactsRefNameDownloadOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **jobToken** | **optional.String**| To be used with triggers for multi-project pipelines, available only on Premium and Ultimate tiers. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdJobsArtifactsRefNameRawartifactPath**
> GetApiV4ProjectsIdJobsArtifactsRefNameRawartifactPath(ctx, id, refName, job, artifactPath, optional)
Download a specific file from artifacts archive from a ref

This feature was introduced in GitLab 11.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **refName** | **string**| Branch or tag name in repository. &#x60;HEAD&#x60; or &#x60;SHA&#x60; references are not supported. | 
  **job** | **string**| The name of the job. | 
  **artifactPath** | **string**| Path to a file inside the artifacts archive. | 
 **optional** | ***ProjectsApiGetApiV4ProjectsIdJobsArtifactsRefNameRawartifactPathOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4ProjectsIdJobsArtifactsRefNameRawartifactPathOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **jobToken** | **optional.String**| To be used with triggers for multi-project pipelines, available only on Premium and Ultimate tiers. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdJobsJobId**
> ApiEntitiesCiJob GetApiV4ProjectsIdJobsJobId(ctx, jobId, id)


Get a specific job of a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **int32**| The ID of a job | 
  **id** | **int32**|  | 

### Return type

[**ApiEntitiesCiJob**](API_Entities_Ci_Job.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdJobsJobIdArtifacts**
> GetApiV4ProjectsIdJobsJobIdArtifacts(ctx, id, jobId, optional)
Download the artifacts archive from a job

This feature was introduced in GitLab 8.5

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **jobId** | **int32**| The ID of a job | 
 **optional** | ***ProjectsApiGetApiV4ProjectsIdJobsJobIdArtifactsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4ProjectsIdJobsJobIdArtifactsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **jobToken** | **optional.String**| To be used with triggers for multi-project pipelines, available only on Premium and Ultimate tiers. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdJobsJobIdArtifactsartifactPath**
> GetApiV4ProjectsIdJobsJobIdArtifactsartifactPath(ctx, id, jobId, artifactPath, optional)
Download a specific file from artifacts archive

This feature was introduced in GitLab 10.0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **jobId** | **int32**| The ID of a job | 
  **artifactPath** | **string**| Path to a file inside the artifacts archive. | 
 **optional** | ***ProjectsApiGetApiV4ProjectsIdJobsJobIdArtifactsartifactPathOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4ProjectsIdJobsJobIdArtifactsartifactPathOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **jobToken** | **optional.String**| To be used with triggers for multi-project pipelines, available only on Premium and Ultimate tiers. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdJobsJobIdTrace**
> ApiEntitiesCiJob GetApiV4ProjectsIdJobsJobIdTrace(ctx, jobId, id)


Get a trace of a specific job of a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **int32**| The ID of a job | 
  **id** | **int32**|  | 

### Return type

[**ApiEntitiesCiJob**](API_Entities_Ci_Job.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdLanguages**
> GetApiV4ProjectsIdLanguages(ctx, id)


Get languages in project repository

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdMergeRequestsMergeRequestIidApprovalState**
> EeApiEntitiesMergeRequestApprovalState GetApiV4ProjectsIdMergeRequestsMergeRequestIidApprovalState(ctx, id, mergeRequestIid)


Get approval state of merge request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **mergeRequestIid** | **int32**| The IID of a merge request | 

### Return type

[**EeApiEntitiesMergeRequestApprovalState**](EE_API_Entities_MergeRequestApprovalState.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdMergeRequestsMergeRequestIidApprovals**
> ApiEntitiesMergeRequestApprovals GetApiV4ProjectsIdMergeRequestsMergeRequestIidApprovals(ctx, id, mergeRequestIid)


List approvals for merge request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **mergeRequestIid** | **int32**|  | 

### Return type

[**ApiEntitiesMergeRequestApprovals**](API_Entities_MergeRequestApprovals.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotes**
> []ApiEntitiesDraftNote GetApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotes(ctx, id, mergeRequestIid)


Get a list of merge request draft notes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a project | 
  **mergeRequestIid** | **int32**| The ID of a merge request | 

### Return type

[**[]ApiEntitiesDraftNote**](API_Entities_DraftNote.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteId**
> ApiEntitiesDraftNote GetApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteId(ctx, id, mergeRequestIid, draftNoteId)


Get a single draft note

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a project | 
  **mergeRequestIid** | **int32**| The ID of a merge request | 
  **draftNoteId** | **int32**| The ID of a draft note | 

### Return type

[**ApiEntitiesDraftNote**](API_Entities_DraftNote.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPagesAccess**
> GetApiV4ProjectsIdPagesAccess(ctx, id)


Check pages access of this project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPagesDomainsDomain**
> ApiEntitiesPagesDomain GetApiV4ProjectsIdPagesDomainsDomain(ctx, id, domain)


Get a single pages domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **domain** | **string**| The domain | 

### Return type

[**ApiEntitiesPagesDomain**](API_Entities_PagesDomain.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPipelineSchedules**
> []ApiEntitiesCiPipelineSchedule GetApiV4ProjectsIdPipelineSchedules(ctx, id, optional)


Get all pipeline schedules

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
 **optional** | ***ProjectsApiGetApiV4ProjectsIdPipelineSchedulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4ProjectsIdPipelineSchedulesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **scope** | **optional.String**| The scope of pipeline schedules | 

### Return type

[**[]ApiEntitiesCiPipelineSchedule**](API_Entities_Ci_PipelineSchedule.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPipelineSchedulesPipelineScheduleId**
> ApiEntitiesCiPipelineScheduleDetails GetApiV4ProjectsIdPipelineSchedulesPipelineScheduleId(ctx, id, pipelineScheduleId)


Get a single pipeline schedule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **pipelineScheduleId** | **int32**| The pipeline schedule id | 

### Return type

[**ApiEntitiesCiPipelineScheduleDetails**](API_Entities_Ci_PipelineScheduleDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdPipelines**
> []ApiEntitiesCiPipelineBasic GetApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdPipelines(ctx, id, pipelineScheduleId)


Get all pipelines triggered from a pipeline schedule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **pipelineScheduleId** | **int32**| The pipeline schedule ID | 

### Return type

[**[]ApiEntitiesCiPipelineBasic**](API_Entities_Ci_PipelineBasic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPipelines**
> []ApiEntitiesCiPipelineBasic GetApiV4ProjectsIdPipelines(ctx, id, optional)
Get all Pipelines of the project

This feature was introduced in GitLab 8.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID or URL-encoded path | 
 **optional** | ***ProjectsApiGetApiV4ProjectsIdPipelinesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4ProjectsIdPipelinesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **scope** | **optional.String**| The scope of pipelines | 
 **status** | **optional.String**| The status of pipelines | 
 **ref** | **optional.String**| The ref of pipelines | 
 **sha** | **optional.String**| The sha of pipelines | 
 **yamlErrors** | **optional.Bool**| Returns pipelines with invalid configurations | 
 **username** | **optional.String**| The username of the user who triggered pipelines | 
 **updatedBefore** | **optional.Time**| Return pipelines updated before the specified datetime. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ | 
 **updatedAfter** | **optional.Time**| Return pipelines updated after the specified datetime. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ | 
 **orderBy** | **optional.String**| Order pipelines | [default to id]
 **sort** | **optional.String**| Sort pipelines | [default to desc]
 **source** | **optional.String**|  | 
 **name** | **optional.String**| Filter pipelines by name | 

### Return type

[**[]ApiEntitiesCiPipelineBasic**](API_Entities_Ci_PipelineBasic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPipelinesLatest**
> ApiEntitiesCiPipelineWithMetadata GetApiV4ProjectsIdPipelinesLatest(ctx, id, optional)
Gets the latest pipeline for the project branch

This feature was introduced in GitLab 12.3

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID or URL-encoded path | 
 **optional** | ***ProjectsApiGetApiV4ProjectsIdPipelinesLatestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4ProjectsIdPipelinesLatestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ref** | **optional.String**| Branch ref of pipeline. Uses project default branch if not specified. | 

### Return type

[**ApiEntitiesCiPipelineWithMetadata**](API_Entities_Ci_PipelineWithMetadata.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPipelinesPipelineId**
> ApiEntitiesCiPipelineWithMetadata GetApiV4ProjectsIdPipelinesPipelineId(ctx, id, pipelineId)
Gets a specific pipeline for the project

This feature was introduced in GitLab 8.11

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID or URL-encoded path | 
  **pipelineId** | **int32**| The pipeline ID | 

### Return type

[**ApiEntitiesCiPipelineWithMetadata**](API_Entities_Ci_PipelineWithMetadata.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPipelinesPipelineIdBridges**
> []ApiEntitiesCiBridge GetApiV4ProjectsIdPipelinesPipelineIdBridges(ctx, id, pipelineId, optional)


Get pipeline bridge jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID or URL-encoded path | 
  **pipelineId** | **int32**| The pipeline ID | 
 **optional** | ***ProjectsApiGetApiV4ProjectsIdPipelinesPipelineIdBridgesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4ProjectsIdPipelinesPipelineIdBridgesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **scope** | **optional.String**| The scope of builds to show | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesCiBridge**](API_Entities_Ci_Bridge.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPipelinesPipelineIdJobs**
> []ApiEntitiesCiJob GetApiV4ProjectsIdPipelinesPipelineIdJobs(ctx, id, pipelineId, optional)


Get pipeline jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID or URL-encoded path | 
  **pipelineId** | **int32**| The pipeline ID | 
 **optional** | ***ProjectsApiGetApiV4ProjectsIdPipelinesPipelineIdJobsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4ProjectsIdPipelinesPipelineIdJobsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeRetried** | **optional.Bool**| Includes retried jobs | [default to false]
 **scope** | **optional.String**| The scope of builds to show | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesCiJob**](API_Entities_Ci_Job.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPipelinesPipelineIdTestReport**
> TestReportEntity GetApiV4ProjectsIdPipelinesPipelineIdTestReport(ctx, id, pipelineId)
Gets the test report for a given pipeline

This feature was introduced in GitLab 13.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID or URL-encoded path | 
  **pipelineId** | **int32**| The pipeline ID | 

### Return type

[**TestReportEntity**](TestReportEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPipelinesPipelineIdTestReportSummary**
> TestReportSummaryEntity GetApiV4ProjectsIdPipelinesPipelineIdTestReportSummary(ctx, id, pipelineId)
Gets the test report summary for a given pipeline

This feature was introduced in GitLab 14.2

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID or URL-encoded path | 
  **pipelineId** | **int32**| The pipeline ID | 

### Return type

[**TestReportSummaryEntity**](TestReportSummaryEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdPipelinesPipelineIdVariables**
> []ApiEntitiesCiVariable GetApiV4ProjectsIdPipelinesPipelineIdVariables(ctx, id, pipelineId)
Gets the variables for a given pipeline

This feature was introduced in GitLab 11.11

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID or URL-encoded path | 
  **pipelineId** | **int32**| The pipeline ID | 

### Return type

[**[]ApiEntitiesCiVariable**](API_Entities_Ci_Variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdProtectedBranches**
> []ApiEntitiesProtectedBranch GetApiV4ProjectsIdProtectedBranches(ctx, id, optional)


Get a project's protected branches

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
 **optional** | ***ProjectsApiGetApiV4ProjectsIdProtectedBranchesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4ProjectsIdProtectedBranchesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **search** | **optional.String**| Search for a protected branch by name | 

### Return type

[**[]ApiEntitiesProtectedBranch**](API_Entities_ProtectedBranch.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdProtectedBranchesName**
> ApiEntitiesProtectedBranch GetApiV4ProjectsIdProtectedBranchesName(ctx, id, name)


Get a single protected branch

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **name** | **string**| The name of the branch or wildcard | 

### Return type

[**ApiEntitiesProtectedBranch**](API_Entities_ProtectedBranch.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdRepositoryArchive**
> GetApiV4ProjectsIdRepositoryArchive(ctx, id, optional)


Get an archive of the repository

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
 **optional** | ***ProjectsApiGetApiV4ProjectsIdRepositoryArchiveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4ProjectsIdRepositoryArchiveOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sha** | **optional.String**| The commit sha of the archive to be downloaded | 
 **format** | **optional.String**| The archive format | 
 **path** | **optional.String**| Subfolder of the repository to be downloaded | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdRepositoryBlobsSha**
> GetApiV4ProjectsIdRepositoryBlobsSha(ctx, id, sha)


Get a blob from the repository

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **sha** | **string**| The commit hash | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdRepositoryBlobsShaRaw**
> GetApiV4ProjectsIdRepositoryBlobsShaRaw(ctx, id, sha)


Get raw blob contents from the repository

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **sha** | **string**| The commit hash | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdRepositoryChangelog**
> ApiEntitiesChangelog GetApiV4ProjectsIdRepositoryChangelog(ctx, id, version, optional)
Generates a changelog section for a release and returns it

This feature was introduced in GitLab 14.6

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **version** | **string**| The version of the release, using the semantic versioning format | 
 **optional** | ***ProjectsApiGetApiV4ProjectsIdRepositoryChangelogOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4ProjectsIdRepositoryChangelogOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **optional.String**| The first commit in the range of commits to use for the changelog | 
 **to** | **optional.String**| The last commit in the range of commits to use for the changelog | 
 **date** | **optional.Time**| The date and time of the release | 
 **trailer** | **optional.String**| The Git trailer to use for determining if commits are to be included in the changelog | [default to Changelog]
 **configFile** | **optional.String**| The file path to the configuration file as stored in the project&#39;s Git repository. Defaults to &#39;.gitlab/changelog_config.yml&#39; | 

### Return type

[**ApiEntitiesChangelog**](API_Entities_Changelog.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdRepositoryCommitsShaStatuses**
> []ApiEntitiesCommitStatus GetApiV4ProjectsIdRepositoryCommitsShaStatuses(ctx, id, sha, optional)


Get a commit's statuses

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **sha** | **string**| The commit hash | 
 **optional** | ***ProjectsApiGetApiV4ProjectsIdRepositoryCommitsShaStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4ProjectsIdRepositoryCommitsShaStatusesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ref** | **optional.String**| The ref | 
 **stage** | **optional.String**| The stage | 
 **name** | **optional.String**| The name | 
 **all** | **optional.Bool**| Show all statuses | [default to false]
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesCommitStatus**](API_Entities_CommitStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdRepositoryCompare**
> ApiEntitiesCompare GetApiV4ProjectsIdRepositoryCompare(ctx, id, from, to, optional)


Compare two branches, tags, or commits

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **from** | **string**| The commit, branch name, or tag name to start comparison | 
  **to** | **string**| The commit, branch name, or tag name to stop comparison | 
 **optional** | ***ProjectsApiGetApiV4ProjectsIdRepositoryCompareOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4ProjectsIdRepositoryCompareOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fromProjectId** | **optional.Int32**| The project to compare from | 
 **straight** | **optional.Bool**| Comparison method, &#x60;true&#x60; for direct comparison between &#x60;from&#x60; and &#x60;to&#x60; (&#x60;from&#x60;..&#x60;to&#x60;), &#x60;false&#x60; to compare using merge base (&#x60;from&#x60;...&#x60;to&#x60;) | [default to false]
 **unidiff** | **optional.Bool**| A diff in a Unified diff format | [default to false]

### Return type

[**ApiEntitiesCompare**](API_Entities_Compare.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdRepositoryContributors**
> ApiEntitiesContributor GetApiV4ProjectsIdRepositoryContributors(ctx, id, optional)


Get repository contributors

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
 **optional** | ***ProjectsApiGetApiV4ProjectsIdRepositoryContributorsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4ProjectsIdRepositoryContributorsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **orderBy** | **optional.String**| Return contributors ordered by &#x60;name&#x60; or &#x60;email&#x60; or &#x60;commits&#x60; | [default to commits]
 **sort** | **optional.String**| Sort by asc (ascending) or desc (descending) | [default to asc]

### Return type

[**ApiEntitiesContributor**](API_Entities_Contributor.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdRepositoryFilesFilePath**
> GetApiV4ProjectsIdRepositoryFilesFilePath(ctx, id, filePath, ref)


Get a file from the repository

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID | 
  **filePath** | **string**| The url encoded path to the file. | 
  **ref** | **string**| The name of branch, tag or commit | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdRepositoryFilesFilePathBlame**
> GetApiV4ProjectsIdRepositoryFilesFilePathBlame(ctx, id, filePath, ref, rangeStart, rangeEnd)


Get blame file from the repository

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID | 
  **filePath** | **string**| The url encoded path to the file. | 
  **ref** | **string**| The name of branch, tag or commit | 
  **rangeStart** | **int32**| The first line of the range to blame | 
  **rangeEnd** | **int32**| The last line of the range to blame | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdRepositoryFilesFilePathRaw**
> *os.File GetApiV4ProjectsIdRepositoryFilesFilePathRaw(ctx, id, filePath, optional)


Get raw file contents from the repository

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID | 
  **filePath** | **string**| The url encoded path to the file. | 
 **optional** | ***ProjectsApiGetApiV4ProjectsIdRepositoryFilesFilePathRawOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4ProjectsIdRepositoryFilesFilePathRawOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ref** | **optional.String**| The name of branch, tag or commit | 
 **lfs** | **optional.Bool**| Retrieve binary data for a file that is an lfs pointer | [default to false]

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdRepositoryMergeBase**
> ApiEntitiesCommit GetApiV4ProjectsIdRepositoryMergeBase(ctx, id, refs)


Get the common ancestor between commits

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **refs** | [**[]string**](string.md)| The refs to find the common ancestor of, multiple refs can be passed | 

### Return type

[**ApiEntitiesCommit**](API_Entities_Commit.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdRepositoryStorageMoves**
> []ApiEntitiesProjectsRepositoryStorageMove GetApiV4ProjectsIdRepositoryStorageMoves(ctx, id, optional)
Get a list of all project repository storage moves

This feature was introduced in GitLab 13.1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
 **optional** | ***ProjectsApiGetApiV4ProjectsIdRepositoryStorageMovesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4ProjectsIdRepositoryStorageMovesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesProjectsRepositoryStorageMove**](API_Entities_Projects_RepositoryStorageMove.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdRepositoryStorageMovesRepositoryStorageMoveId**
> ApiEntitiesProjectsRepositoryStorageMove GetApiV4ProjectsIdRepositoryStorageMovesRepositoryStorageMoveId(ctx, id, repositoryStorageMoveId)
Get a project repository storage move

This feature was introduced in GitLab 13.1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **repositoryStorageMoveId** | **int32**| The ID of a project repository storage move | 

### Return type

[**ApiEntitiesProjectsRepositoryStorageMove**](API_Entities_Projects_RepositoryStorageMove.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdRepositoryTree**
> ApiEntitiesTreeObject GetApiV4ProjectsIdRepositoryTree(ctx, id, optional)


Get a project repository tree

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
 **optional** | ***ProjectsApiGetApiV4ProjectsIdRepositoryTreeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4ProjectsIdRepositoryTreeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ref** | **optional.String**| The name of a repository branch or tag, if not given the default branch is used | 
 **path** | **optional.String**| The path of the tree | 
 **recursive** | **optional.Bool**| Used to get a recursive tree | [default to false]
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **pagination** | **optional.String**| Specify the pagination method (\&quot;none\&quot; is only valid if \&quot;recursive\&quot; is true) | [default to legacy]
 **pageToken** | **optional.String**| Record from which to start the keyset pagination | 

### Return type

[**ApiEntitiesTreeObject**](API_Entities_TreeObject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdRunners**
> ApiEntitiesCiRunner GetApiV4ProjectsIdRunners(ctx, id, optional)
List project's runners

List all runners available in the project, including from ancestor groups and any allowed shared runners.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
 **optional** | ***ProjectsApiGetApiV4ProjectsIdRunnersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4ProjectsIdRunnersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scope** | **optional.String**| Deprecated: Use &#x60;type&#x60; or &#x60;status&#x60; instead. The scope of runners to return | 
 **type_** | **optional.String**| The type of runners to return | 
 **paused** | **optional.Bool**| Whether to include only runners that are accepting or ignoring new jobs | 
 **status** | **optional.String**| The status of runners to return | 
 **tagList** | [**optional.Interface of []string**](string.md)| A list of runner tags | 
 **versionPrefix** | **optional.String**| The version prefix of runners to return | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**ApiEntitiesCiRunner**](API_Entities_Ci_Runner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdShareLocations**
> ApiEntitiesGroup GetApiV4ProjectsIdShareLocations(ctx, id, optional)


Returns group that can be shared with the given project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The id of the project | 
 **optional** | ***ProjectsApiGetApiV4ProjectsIdShareLocationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4ProjectsIdShareLocationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| Return list of groups matching the search criteria | 

### Return type

[**ApiEntitiesGroup**](API_Entities_Group.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdSnapshot**
> *os.File GetApiV4ProjectsIdSnapshot(ctx, id, optional)
Download a (possibly inconsistent) snapshot of a repository

This feature was introduced in GitLab 10.7

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***ProjectsApiGetApiV4ProjectsIdSnapshotOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4ProjectsIdSnapshotOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wiki** | **optional.Bool**| Set to true to receive the wiki repository | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/x-tar

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdStarrers**
> []ApiEntitiesUserBasic GetApiV4ProjectsIdStarrers(ctx, id, optional)


Get the users who starred a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
 **optional** | ***ProjectsApiGetApiV4ProjectsIdStarrersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4ProjectsIdStarrersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| Return list of users matching the search criteria | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesUserBasic**](API_Entities_UserBasic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdStatistics**
> ApiEntitiesProjectDailyStatistics GetApiV4ProjectsIdStatistics(ctx, id)


Get the list of project fetch statistics for the last 30 days

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 

### Return type

[**ApiEntitiesProjectDailyStatistics**](API_Entities_ProjectDailyStatistics.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdStorage**
> ApiEntitiesProjectRepositoryStorage GetApiV4ProjectsIdStorage(ctx, id)


Show the storage information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of a project | 

### Return type

[**ApiEntitiesProjectRepositoryStorage**](API_Entities_ProjectRepositoryStorage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdTemplatesType**
> []ApiEntitiesTemplatesList GetApiV4ProjectsIdTemplatesType(ctx, id, type_, optional)
Get a list of templates available to this project

This endpoint was introduced in GitLab 11.4

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **type_** | **string**| The type (dockerfiles|gitignores|gitlab_ci_ymls|licenses|issues|merge_requests) of the template | 
 **optional** | ***ProjectsApiGetApiV4ProjectsIdTemplatesTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4ProjectsIdTemplatesTypeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesTemplatesList**](API_Entities_TemplatesList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdTemplatesTypeName**
> ApiEntitiesLicense GetApiV4ProjectsIdTemplatesTypeName(ctx, id, type_, name, optional)
Download a template available to this project

This endpoint was introduced in GitLab 11.4

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **type_** | **string**| The type (dockerfiles|gitignores|gitlab_ci_ymls|licenses|issues|merge_requests) of the template | 
  **name** | **string**| The key of the template, as obtained from the collection endpoint. | 
 **optional** | ***ProjectsApiGetApiV4ProjectsIdTemplatesTypeNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4ProjectsIdTemplatesTypeNameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sourceTemplateProjectId** | **optional.Int32**| The project id where a given template is being stored. This is useful when multiple templates from different projects have the same name | 
 **project** | **optional.String**| The project name to use when expanding placeholders in the template. Only affects licenses | 
 **fullname** | **optional.String**| The full name of the copyright holder to use when expanding placeholders in the template. Only affects licenses | 

### Return type

[**ApiEntitiesLicense**](API_Entities_License.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdTransferLocations**
> []ApiEntitiesPublicGroupDetails GetApiV4ProjectsIdTransferLocations(ctx, id, optional)


Get the namespaces to where the project can be transferred

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
 **optional** | ***ProjectsApiGetApiV4ProjectsIdTransferLocationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4ProjectsIdTransferLocationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| Return list of namespaces matching the search criteria | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesPublicGroupDetails**](API_Entities_PublicGroupDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdTriggers**
> []ApiEntitiesTrigger GetApiV4ProjectsIdTriggers(ctx, id, optional)


Get trigger tokens list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
 **optional** | ***ProjectsApiGetApiV4ProjectsIdTriggersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4ProjectsIdTriggersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesTrigger**](API_Entities_Trigger.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdTriggersTriggerId**
> ApiEntitiesTrigger GetApiV4ProjectsIdTriggersTriggerId(ctx, id, triggerId)


Get specific trigger token of a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **triggerId** | **int32**| The trigger token ID | 

### Return type

[**ApiEntitiesTrigger**](API_Entities_Trigger.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4ProjectsIdUsers**
> []ApiEntitiesUserBasic GetApiV4ProjectsIdUsers(ctx, id, optional)


Get the users list of a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
 **optional** | ***ProjectsApiGetApiV4ProjectsIdUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4ProjectsIdUsersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| Return list of users matching the search criteria | 
 **skipUsers** | [**optional.Interface of []int32**](int32.md)| Filter out users with the specified IDs | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesUserBasic**](API_Entities_UserBasic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4UsersUserIdContributedProjects**
> []ApiEntitiesBasicProjectDetails GetApiV4UsersUserIdContributedProjects(ctx, userId, optional)


Get projects that a user has contributed to

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| The ID or username of the user | 
 **optional** | ***ProjectsApiGetApiV4UsersUserIdContributedProjectsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4UsersUserIdContributedProjectsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderBy** | **optional.String**| Return projects ordered by field. storage_size, repository_size, wiki_size, packages_size are only available to admins. Similarity is available when searching and is limited to projects the user has access to. | [default to created_at]
 **sort** | **optional.String**| Return projects sorted in ascending and descending order | [default to desc]
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **simple** | **optional.Bool**| Return only the ID, URL, name, and path of each project | [default to false]

### Return type

[**[]ApiEntitiesBasicProjectDetails**](API_Entities_BasicProjectDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4UsersUserIdProjects**
> []ApiEntitiesBasicProjectDetails GetApiV4UsersUserIdProjects(ctx, userId, optional)


Get a user projects

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| The ID or username of the user | 
 **optional** | ***ProjectsApiGetApiV4UsersUserIdProjectsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4UsersUserIdProjectsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderBy** | **optional.String**| Return projects ordered by field. storage_size, repository_size, wiki_size, packages_size are only available to admins. Similarity is available when searching and is limited to projects the user has access to. | [default to created_at]
 **sort** | **optional.String**| Return projects sorted in ascending and descending order | [default to desc]
 **archived** | **optional.Bool**| Limit by archived status | 
 **visibility** | **optional.String**| Limit by visibility | 
 **search** | **optional.String**| Return list of projects matching the search criteria | 
 **searchNamespaces** | **optional.Bool**| Include ancestor namespaces when matching search criteria | 
 **owned** | **optional.Bool**| Limit by owned by authenticated user | [default to false]
 **starred** | **optional.Bool**| Limit by starred status | [default to false]
 **imported** | **optional.Bool**| Limit by imported by authenticated user | [default to false]
 **membership** | **optional.Bool**| Limit by projects that the current user is a member of | [default to false]
 **withIssuesEnabled** | **optional.Bool**| Limit by enabled issues feature | [default to false]
 **withMergeRequestsEnabled** | **optional.Bool**| Limit by enabled merge requests feature | [default to false]
 **withProgrammingLanguage** | **optional.String**| Limit to repositories which use the given programming language | 
 **minAccessLevel** | **optional.Int32**| Limit by minimum access level of authenticated user | 
 **idAfter** | **optional.Int32**| Limit results to projects with IDs greater than the specified ID | 
 **idBefore** | **optional.Int32**| Limit results to projects with IDs less than the specified ID | 
 **lastActivityAfter** | **optional.Time**| Limit results to projects with last_activity after specified time. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ | 
 **lastActivityBefore** | **optional.Time**| Limit results to projects with last_activity before specified time. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ | 
 **repositoryStorage** | **optional.String**| Which storage shard the repository is on. Available only to admins | 
 **topic** | [**optional.Interface of []string**](string.md)| Comma-separated list of topics. Limit results to projects having all topics | 
 **topicId** | **optional.Int32**| Limit results to projects with the assigned topic given by the topic ID | 
 **updatedBefore** | **optional.Time**| Return projects updated before the specified datetime. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ | 
 **updatedAfter** | **optional.Time**| Return projects updated after the specified datetime. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ | 
 **includePendingDelete** | **optional.Bool**| Include projects in pending delete state. Can only be set by admins | 
 **wikiChecksumFailed** | **optional.Bool**| Limit by projects where wiki checksum is failed | [default to false]
 **repositoryChecksumFailed** | **optional.Bool**| Limit by projects where repository checksum is failed | [default to false]
 **includeHidden** | **optional.Bool**| Include hidden projects. Can only be set by admins | [default to false]
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **simple** | **optional.Bool**| Return only the ID, URL, name, and path of each project | [default to false]
 **statistics** | **optional.Bool**| Include project statistics | [default to false]
 **withCustomAttributes** | **optional.Bool**| Include custom attributes in the response | [default to false]

### Return type

[**[]ApiEntitiesBasicProjectDetails**](API_Entities_BasicProjectDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4UsersUserIdStarredProjects**
> []ApiEntitiesBasicProjectDetails GetApiV4UsersUserIdStarredProjects(ctx, userId, optional)


Get projects starred by a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| The ID or username of the user | 
 **optional** | ***ProjectsApiGetApiV4UsersUserIdStarredProjectsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiGetApiV4UsersUserIdStarredProjectsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderBy** | **optional.String**| Return projects ordered by field. storage_size, repository_size, wiki_size, packages_size are only available to admins. Similarity is available when searching and is limited to projects the user has access to. | [default to created_at]
 **sort** | **optional.String**| Return projects sorted in ascending and descending order | [default to desc]
 **archived** | **optional.Bool**| Limit by archived status | 
 **visibility** | **optional.String**| Limit by visibility | 
 **search** | **optional.String**| Return list of projects matching the search criteria | 
 **searchNamespaces** | **optional.Bool**| Include ancestor namespaces when matching search criteria | 
 **owned** | **optional.Bool**| Limit by owned by authenticated user | [default to false]
 **starred** | **optional.Bool**| Limit by starred status | [default to false]
 **imported** | **optional.Bool**| Limit by imported by authenticated user | [default to false]
 **membership** | **optional.Bool**| Limit by projects that the current user is a member of | [default to false]
 **withIssuesEnabled** | **optional.Bool**| Limit by enabled issues feature | [default to false]
 **withMergeRequestsEnabled** | **optional.Bool**| Limit by enabled merge requests feature | [default to false]
 **withProgrammingLanguage** | **optional.String**| Limit to repositories which use the given programming language | 
 **minAccessLevel** | **optional.Int32**| Limit by minimum access level of authenticated user | 
 **idAfter** | **optional.Int32**| Limit results to projects with IDs greater than the specified ID | 
 **idBefore** | **optional.Int32**| Limit results to projects with IDs less than the specified ID | 
 **lastActivityAfter** | **optional.Time**| Limit results to projects with last_activity after specified time. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ | 
 **lastActivityBefore** | **optional.Time**| Limit results to projects with last_activity before specified time. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ | 
 **repositoryStorage** | **optional.String**| Which storage shard the repository is on. Available only to admins | 
 **topic** | [**optional.Interface of []string**](string.md)| Comma-separated list of topics. Limit results to projects having all topics | 
 **topicId** | **optional.Int32**| Limit results to projects with the assigned topic given by the topic ID | 
 **updatedBefore** | **optional.Time**| Return projects updated before the specified datetime. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ | 
 **updatedAfter** | **optional.Time**| Return projects updated after the specified datetime. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ | 
 **includePendingDelete** | **optional.Bool**| Include projects in pending delete state. Can only be set by admins | 
 **wikiChecksumFailed** | **optional.Bool**| Limit by projects where wiki checksum is failed | [default to false]
 **repositoryChecksumFailed** | **optional.Bool**| Limit by projects where repository checksum is failed | [default to false]
 **includeHidden** | **optional.Bool**| Include hidden projects. Can only be set by admins | [default to false]
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **simple** | **optional.Bool**| Return only the ID, URL, name, and path of each project | [default to false]
 **statistics** | **optional.Bool**| Include project statistics | [default to false]

### Return type

[**[]ApiEntitiesBasicProjectDetails**](API_Entities_BasicProjectDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HeadApiV4ProjectsIdRepositoryFilesFilePath**
> HeadApiV4ProjectsIdRepositoryFilesFilePath(ctx, id, filePath, ref)


Get file metadata from repository

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID | 
  **filePath** | **string**| The url encoded path to the file. | 
  **ref** | **string**| The name of branch, tag or commit | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HeadApiV4ProjectsIdRepositoryFilesFilePathBlame**
> HeadApiV4ProjectsIdRepositoryFilesFilePathBlame(ctx, id, filePath, ref)


Get blame file metadata from repository

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID | 
  **filePath** | **string**| The url encoded path to the file. | 
  **ref** | **string**| The name of branch, tag or commit | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchApiV4ProjectsIdProtectedBranchesName**
> ApiEntitiesProtectedBranch PatchApiV4ProjectsIdProtectedBranchesName(ctx, id, name, patchApiV4ProjectsIdProtectedBranchesName)


Update a protected branch

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **name** | **string**| The name of the branch | 
  **patchApiV4ProjectsIdProtectedBranchesName** | [**PatchApiV4ProjectsIdProtectedBranchesName**](PatchApiV4ProjectsIdProtectedBranchesName.md)|  | 

### Return type

[**ApiEntitiesProtectedBranch**](API_Entities_ProtectedBranch.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4Projects**
> ApiEntitiesProject PostApiV4Projects(ctx, postApiV4Projects)


Create new project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4Projects** | [**PostApiV4Projects**](PostApiV4Projects.md)|  | 

### Return type

[**ApiEntitiesProject**](API_Entities_Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdArchive**
> ApiEntitiesProject PostApiV4ProjectsIdArchive(ctx, id)


Archive a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 

### Return type

[**ApiEntitiesProject**](API_Entities_Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdCreateCiConfig**
> PostApiV4ProjectsIdCreateCiConfig(ctx, id)


Creates merge request for missing ci config in project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdFork**
> ApiEntitiesProject PostApiV4ProjectsIdFork(ctx, id, postApiV4ProjectsIdFork)


Fork new project for the current user or provided namespace.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **postApiV4ProjectsIdFork** | [**PostApiV4ProjectsIdFork**](PostApiV4ProjectsIdFork.md)|  | 

### Return type

[**ApiEntitiesProject**](API_Entities_Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdForkForkedFromId**
> ApiEntitiesProject PostApiV4ProjectsIdForkForkedFromId(ctx, id, forkedFromId)


Mark this project as forked from another

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **forkedFromId** | **string**| The ID of the project it was forked from | 

### Return type

[**ApiEntitiesProject**](API_Entities_Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdHooksHookIdTestTrigger**
> PostApiV4ProjectsIdHooksHookIdTestTrigger(ctx, hookId, trigger, id)
Triggers a hook test

Triggers a hook test

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hookId** | **int32**| The ID of the hook | 
  **trigger** | **string**| The type of trigger hook | 
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdHousekeeping**
> PostApiV4ProjectsIdHousekeeping(ctx, id, postApiV4ProjectsIdHousekeeping)
Start the housekeeping task for a project

This feature was introduced in GitLab 9.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **postApiV4ProjectsIdHousekeeping** | [**PostApiV4ProjectsIdHousekeeping**](PostApiV4ProjectsIdHousekeeping.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdImportProjectMembersProjectId**
> PostApiV4ProjectsIdImportProjectMembersProjectId(ctx, id, projectId)
Import members from another project

This feature was introduced in GitLab 14.2

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **projectId** | **int32**| The ID of the source project to import the members from. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdJobsJobIdArtifactsKeep**
> ApiEntitiesCiJob PostApiV4ProjectsIdJobsJobIdArtifactsKeep(ctx, id, jobId)


Keep the artifacts to prevent them from being deleted

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **jobId** | **int32**| The ID of a job | 

### Return type

[**ApiEntitiesCiJob**](API_Entities_Ci_Job.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdJobsJobIdCancel**
> ApiEntitiesCiJob PostApiV4ProjectsIdJobsJobIdCancel(ctx, jobId, id)


Cancel a specific job of a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **int32**| The ID of a job | 
  **id** | **int32**|  | 

### Return type

[**ApiEntitiesCiJob**](API_Entities_Ci_Job.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdJobsJobIdErase**
> ApiEntitiesCiJob PostApiV4ProjectsIdJobsJobIdErase(ctx, jobId, id)


Erase job (remove artifacts and the trace)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **int32**| The ID of a build | 
  **id** | **int32**|  | 

### Return type

[**ApiEntitiesCiJob**](API_Entities_Ci_Job.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdJobsJobIdPlay**
> ApiEntitiesCiJobBasic PostApiV4ProjectsIdJobsJobIdPlay(ctx, jobId, id, postApiV4ProjectsIdJobsJobIdPlay)
Trigger an actionable job (manual, delayed, etc)

This feature was added in GitLab 8.11

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **int32**| The ID of a Job | 
  **id** | **int32**|  | 
  **postApiV4ProjectsIdJobsJobIdPlay** | [**PostApiV4ProjectsIdJobsJobIdPlay**](PostApiV4ProjectsIdJobsJobIdPlay.md)|  | 

### Return type

[**ApiEntitiesCiJobBasic**](API_Entities_Ci_JobBasic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdJobsJobIdRetry**
> ApiEntitiesCiJob PostApiV4ProjectsIdJobsJobIdRetry(ctx, jobId, id)


Retry a specific job of a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **int32**| The ID of a job | 
  **id** | **int32**|  | 

### Return type

[**ApiEntitiesCiJob**](API_Entities_Ci_Job.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdMergeRequestsMergeRequestIidApprovals**
> EeApiEntitiesApprovalState PostApiV4ProjectsIdMergeRequestsMergeRequestIidApprovals(ctx, id, mergeRequestIid, postApiV4ProjectsIdMergeRequestsMergeRequestIidApprovals)
Deprecated in 16.0: Use the merge request approvals API instead. Change approval-related configuration

This feature was introduced in 10.6 and deprecated in 16.0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **mergeRequestIid** | **int32**| The IID of a merge request | 
  **postApiV4ProjectsIdMergeRequestsMergeRequestIidApprovals** | [**PostApiV4ProjectsIdMergeRequestsMergeRequestIidApprovals**](PostApiV4ProjectsIdMergeRequestsMergeRequestIidApprovals.md)|  | 

### Return type

[**EeApiEntitiesApprovalState**](EE_API_Entities_ApprovalState.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdMergeRequestsMergeRequestIidApprove**
> ApiEntitiesMergeRequestApprovals PostApiV4ProjectsIdMergeRequestsMergeRequestIidApprove(ctx, id, mergeRequestIid, postApiV4ProjectsIdMergeRequestsMergeRequestIidApprove)


Approve a merge request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **mergeRequestIid** | **int32**|  | 
  **postApiV4ProjectsIdMergeRequestsMergeRequestIidApprove** | [**PostApiV4ProjectsIdMergeRequestsMergeRequestIidApprove**](PostApiV4ProjectsIdMergeRequestsMergeRequestIidApprove.md)|  | 

### Return type

[**ApiEntitiesMergeRequestApprovals**](API_Entities_MergeRequestApprovals.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotes**
> ApiEntitiesDraftNote PostApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotes(ctx, id, mergeRequestIid, postApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotes)


Create a new draft note

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a project. | 
  **mergeRequestIid** | **int32**| The ID of a merge request. | 
  **postApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotes** | [**PostApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotes**](PostApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotes.md)|  | 

### Return type

[**ApiEntitiesDraftNote**](API_Entities_DraftNote.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesBulkPublish**
> PostApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesBulkPublish(ctx, id, mergeRequestIid)


Bulk publish all pending draft notes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a project | 
  **mergeRequestIid** | **int32**| The ID of a merge request | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdMergeRequestsMergeRequestIidUnapprove**
> ApiEntitiesMergeRequestApprovals PostApiV4ProjectsIdMergeRequestsMergeRequestIidUnapprove(ctx, id, mergeRequestIid)


Remove an approval from a merge request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **mergeRequestIid** | **int32**|  | 

### Return type

[**ApiEntitiesMergeRequestApprovals**](API_Entities_MergeRequestApprovals.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdPagesDomains**
> ApiEntitiesPagesDomain PostApiV4ProjectsIdPagesDomains(ctx, id, postApiV4ProjectsIdPagesDomains)


Create a new pages domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **postApiV4ProjectsIdPagesDomains** | [**PostApiV4ProjectsIdPagesDomains**](PostApiV4ProjectsIdPagesDomains.md)|  | 

### Return type

[**ApiEntitiesPagesDomain**](API_Entities_PagesDomain.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdPipeline**
> ApiEntitiesCiPipeline PostApiV4ProjectsIdPipeline(ctx, id, postApiV4ProjectsIdPipeline)
Create a new pipeline

This feature was introduced in GitLab 8.14

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID or URL-encoded path | 
  **postApiV4ProjectsIdPipeline** | [**PostApiV4ProjectsIdPipeline**](PostApiV4ProjectsIdPipeline.md)|  | 

### Return type

[**ApiEntitiesCiPipeline**](API_Entities_Ci_Pipeline.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdPipelineSchedules**
> ApiEntitiesCiPipelineScheduleDetails PostApiV4ProjectsIdPipelineSchedules(ctx, id, postApiV4ProjectsIdPipelineSchedules)


Create a new pipeline schedule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **postApiV4ProjectsIdPipelineSchedules** | [**PostApiV4ProjectsIdPipelineSchedules**](PostApiV4ProjectsIdPipelineSchedules.md)|  | 

### Return type

[**ApiEntitiesCiPipelineScheduleDetails**](API_Entities_Ci_PipelineScheduleDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdPlay**
> PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdPlay(ctx, id, pipelineScheduleId)
Play a scheduled pipeline immediately

This feature was added in GitLab 12.8

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **pipelineScheduleId** | **int32**| The pipeline schedule id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdTakeOwnership**
> ApiEntitiesCiPipelineScheduleDetails PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdTakeOwnership(ctx, id, pipelineScheduleId)


Take ownership of a pipeline schedule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **pipelineScheduleId** | **int32**| The pipeline schedule id | 

### Return type

[**ApiEntitiesCiPipelineScheduleDetails**](API_Entities_Ci_PipelineScheduleDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariables**
> ApiEntitiesCiVariable PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariables(ctx, id, pipelineScheduleId, postApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariables)


Create a new pipeline schedule variable

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **pipelineScheduleId** | **int32**| The pipeline schedule id | 
  **postApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariables** | [**PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariables**](PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariables.md)|  | 

### Return type

[**ApiEntitiesCiVariable**](API_Entities_Ci_Variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdPipelinesPipelineIdCancel**
> ApiEntitiesCiPipeline PostApiV4ProjectsIdPipelinesPipelineIdCancel(ctx, id, pipelineId)
Cancel all builds in the pipeline

This feature was introduced in GitLab 8.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID or URL-encoded path | 
  **pipelineId** | **int32**| The pipeline ID | 

### Return type

[**ApiEntitiesCiPipeline**](API_Entities_Ci_Pipeline.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdPipelinesPipelineIdRetry**
> ApiEntitiesCiPipeline PostApiV4ProjectsIdPipelinesPipelineIdRetry(ctx, id, pipelineId)
Retry builds in the pipeline

This feature was introduced in GitLab 8.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID or URL-encoded path | 
  **pipelineId** | **int32**| The pipeline ID | 

### Return type

[**ApiEntitiesCiPipeline**](API_Entities_Ci_Pipeline.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdProtectedBranches**
> ApiEntitiesProtectedBranch PostApiV4ProjectsIdProtectedBranches(ctx, id, postApiV4ProjectsIdProtectedBranches)


Protect a single branch

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **postApiV4ProjectsIdProtectedBranches** | [**PostApiV4ProjectsIdProtectedBranches**](PostApiV4ProjectsIdProtectedBranches.md)|  | 

### Return type

[**ApiEntitiesProtectedBranch**](API_Entities_ProtectedBranch.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdRefReftriggerPipeline**
> ApiEntitiesCiPipeline PostApiV4ProjectsIdRefReftriggerPipeline(ctx, id, ref, postApiV4ProjectsIdRefReftriggerPipeline)


Trigger a GitLab project pipeline

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **ref** | **string**| The commit sha or name of a branch or tag | 
  **postApiV4ProjectsIdRefReftriggerPipeline** | [**PostApiV4ProjectsIdRefReftriggerPipeline**](PostApiV4ProjectsIdRefReftriggerPipeline.md)|  | 

### Return type

[**ApiEntitiesCiPipeline**](API_Entities_Ci_Pipeline.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdRepositoryChangelog**
> PostApiV4ProjectsIdRepositoryChangelog(ctx, id, postApiV4ProjectsIdRepositoryChangelog)
Generates a changelog section for a release and commits it in a changelog file

This feature was introduced in GitLab 13.9

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **postApiV4ProjectsIdRepositoryChangelog** | [**PostApiV4ProjectsIdRepositoryChangelog**](PostApiV4ProjectsIdRepositoryChangelog.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdRepositoryFilesFilePath**
> PostApiV4ProjectsIdRepositoryFilesFilePath(ctx, id, filePath, postApiV4ProjectsIdRepositoryFilesFilePath)


Create new file in repository

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID | 
  **filePath** | **string**| The url encoded path to the file. | 
  **postApiV4ProjectsIdRepositoryFilesFilePath** | [**PostApiV4ProjectsIdRepositoryFilesFilePath**](PostApiV4ProjectsIdRepositoryFilesFilePath.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdRepositorySize**
> PostApiV4ProjectsIdRepositorySize(ctx, id)
Start a task to recalculate repository size for a project

This feature was introduced in GitLab 15.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdRepositoryStorageMoves**
> ApiEntitiesProjectsRepositoryStorageMove PostApiV4ProjectsIdRepositoryStorageMoves(ctx, id, postApiV4ProjectsIdRepositoryStorageMoves)
Schedule a project repository storage move

This feature was introduced in GitLab 13.1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **postApiV4ProjectsIdRepositoryStorageMoves** | [**PostApiV4ProjectsIdRepositoryStorageMoves**](PostApiV4ProjectsIdRepositoryStorageMoves.md)|  | 

### Return type

[**ApiEntitiesProjectsRepositoryStorageMove**](API_Entities_Projects_RepositoryStorageMove.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdRestore**
> ApiEntitiesProject PostApiV4ProjectsIdRestore(ctx, id)


Restore a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**ApiEntitiesProject**](API_Entities_Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdRunners**
> ApiEntitiesCiRunner PostApiV4ProjectsIdRunners(ctx, id, postApiV4ProjectsIdRunners)
Enable a runner in project

Enable an available project runner in the project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **postApiV4ProjectsIdRunners** | [**PostApiV4ProjectsIdRunners**](PostApiV4ProjectsIdRunners.md)|  | 

### Return type

[**ApiEntitiesCiRunner**](API_Entities_Ci_Runner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdRunnersResetRegistrationToken**
> ApiEntitiesCiResetTokenResult PostApiV4ProjectsIdRunnersResetRegistrationToken(ctx, id)
Reset the runner registration token for a project

Reset runner registration token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a project | 

### Return type

[**ApiEntitiesCiResetTokenResult**](API_Entities_Ci_ResetTokenResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdShare**
> ApiEntitiesProjectGroupLink PostApiV4ProjectsIdShare(ctx, id, postApiV4ProjectsIdShare)


Share the project with a group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **postApiV4ProjectsIdShare** | [**PostApiV4ProjectsIdShare**](PostApiV4ProjectsIdShare.md)|  | 

### Return type

[**ApiEntitiesProjectGroupLink**](API_Entities_ProjectGroupLink.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdStar**
> ApiEntitiesProject PostApiV4ProjectsIdStar(ctx, id)


Star a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 

### Return type

[**ApiEntitiesProject**](API_Entities_Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdStatusesSha**
> ApiEntitiesCommitStatus PostApiV4ProjectsIdStatusesSha(ctx, id, sha, postApiV4ProjectsIdStatusesSha)


Post status to a commit

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **sha** | **string**| The commit hash | 
  **postApiV4ProjectsIdStatusesSha** | [**PostApiV4ProjectsIdStatusesSha**](PostApiV4ProjectsIdStatusesSha.md)|  | 

### Return type

[**ApiEntitiesCommitStatus**](API_Entities_CommitStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdTriggers**
> ApiEntitiesTrigger PostApiV4ProjectsIdTriggers(ctx, id, postApiV4ProjectsIdTriggers)


Create a trigger token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **postApiV4ProjectsIdTriggers** | [**PostApiV4ProjectsIdTriggers**](PostApiV4ProjectsIdTriggers.md)|  | 

### Return type

[**ApiEntitiesTrigger**](API_Entities_Trigger.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdUnarchive**
> ApiEntitiesProject PostApiV4ProjectsIdUnarchive(ctx, id)


Unarchive a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 

### Return type

[**ApiEntitiesProject**](API_Entities_Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdUnstar**
> ApiEntitiesProject PostApiV4ProjectsIdUnstar(ctx, id)


Unstar a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 

### Return type

[**ApiEntitiesProject**](API_Entities_Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdUploads**
> ApiEntitiesProjectUpload PostApiV4ProjectsIdUploads(ctx, id, postApiV4ProjectsIdUploads)


Upload a file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **postApiV4ProjectsIdUploads** | [**PostApiV4ProjectsIdUploads**](PostApiV4ProjectsIdUploads.md)|  | 

### Return type

[**ApiEntitiesProjectUpload**](API_Entities_ProjectUpload.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsIdUploadsAuthorize**
> PostApiV4ProjectsIdUploadsAuthorize(ctx, id)
Workhorse authorize the file upload

This feature was introduced in GitLab 13.11

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4ProjectsUserUserId**
> ApiEntitiesProject PostApiV4ProjectsUserUserId(ctx, userId, postApiV4ProjectsUserUserId)


Create new project for a specified user. Only available to admin users.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int32**| The ID of a user | 
  **postApiV4ProjectsUserUserId** | [**PostApiV4ProjectsUserUserId**](PostApiV4ProjectsUserUserId.md)|  | 

### Return type

[**ApiEntitiesProject**](API_Entities_Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsId**
> ApiEntitiesProject PutApiV4ProjectsId(ctx, id, putApiV4ProjectsId)


Update an existing project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **putApiV4ProjectsId** | [**PutApiV4ProjectsId**](PutApiV4ProjectsId.md)|  | 

### Return type

[**ApiEntitiesProject**](API_Entities_Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdCustomAttributesKey**
> PutApiV4ProjectsIdCustomAttributesKey(ctx, key, id, putApiV4ProjectsIdCustomAttributesKey)


Set a custom attribute on a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| The key of the custom attribute | 
  **id** | **int32**|  | 
  **putApiV4ProjectsIdCustomAttributesKey** | [**PutApiV4ProjectsIdCustomAttributesKey**](PutApiV4ProjectsIdCustomAttributesKey.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdHooksHookIdUrlVariablesKey**
> PutApiV4ProjectsIdHooksHookIdUrlVariablesKey(ctx, hookId, key, id, putApiV4ProjectsIdHooksHookIdUrlVariablesKey)


Set a url variable

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hookId** | **int32**| The ID of the hook | 
  **key** | **string**| The key of the variable | 
  **id** | **int32**|  | 
  **putApiV4ProjectsIdHooksHookIdUrlVariablesKey** | [**PutApiV4ProjectsIdHooksHookIdUrlVariablesKey**](PutApiV4ProjectsIdHooksHookIdUrlVariablesKey.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteId**
> ApiEntitiesDraftNote PutApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteId(ctx, id, mergeRequestIid, draftNoteId, putApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteId)


Modify an existing draft note

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a project. | 
  **mergeRequestIid** | **int32**| The ID of a merge request. | 
  **draftNoteId** | **int32**| The ID of a draft note | 
  **putApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteId** | [**PutApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteId**](PutApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteId.md)|  | 

### Return type

[**ApiEntitiesDraftNote**](API_Entities_DraftNote.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdPublish**
> PutApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdPublish(ctx, id, mergeRequestIid, draftNoteId)


Publish a pending draft note

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a project | 
  **mergeRequestIid** | **int32**| The ID of a merge request | 
  **draftNoteId** | **int32**| The ID of a draft note | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdPagesDomainsDomain**
> PutApiV4ProjectsIdPagesDomainsDomain(ctx, id, domain, putApiV4ProjectsIdPagesDomainsDomain)


Updates a pages domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project owned by the authenticated user | 
  **domain** | **string**| The domain | 
  **putApiV4ProjectsIdPagesDomainsDomain** | [**PutApiV4ProjectsIdPagesDomainsDomain**](PutApiV4ProjectsIdPagesDomainsDomain.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdPipelineSchedulesPipelineScheduleId**
> ApiEntitiesCiPipelineScheduleDetails PutApiV4ProjectsIdPipelineSchedulesPipelineScheduleId(ctx, id, pipelineScheduleId, putApiV4ProjectsIdPipelineSchedulesPipelineScheduleId)


Edit a pipeline schedule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **pipelineScheduleId** | **int32**| The pipeline schedule id | 
  **putApiV4ProjectsIdPipelineSchedulesPipelineScheduleId** | [**PutApiV4ProjectsIdPipelineSchedulesPipelineScheduleId**](PutApiV4ProjectsIdPipelineSchedulesPipelineScheduleId.md)|  | 

### Return type

[**ApiEntitiesCiPipelineScheduleDetails**](API_Entities_Ci_PipelineScheduleDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKey**
> ApiEntitiesCiVariable PutApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKey(ctx, id, pipelineScheduleId, key, putApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKey)


Edit a pipeline schedule variable

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **pipelineScheduleId** | **int32**| The pipeline schedule id | 
  **key** | **string**| The key of the variable | 
  **putApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKey** | [**PutApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKey**](PutApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKey.md)|  | 

### Return type

[**ApiEntitiesCiVariable**](API_Entities_Ci_Variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdPipelinesPipelineIdMetadata**
> ApiEntitiesCiPipelineWithMetadata PutApiV4ProjectsIdPipelinesPipelineIdMetadata(ctx, id, pipelineId, putApiV4ProjectsIdPipelinesPipelineIdMetadata)
Updates pipeline metadata

This feature was introduced in GitLab 16.6

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID or URL-encoded path | 
  **pipelineId** | **int32**| The pipeline ID | 
  **putApiV4ProjectsIdPipelinesPipelineIdMetadata** | [**PutApiV4ProjectsIdPipelinesPipelineIdMetadata**](PutApiV4ProjectsIdPipelinesPipelineIdMetadata.md)|  | 

### Return type

[**ApiEntitiesCiPipelineWithMetadata**](API_Entities_Ci_PipelineWithMetadata.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdRepositoryFilesFilePath**
> PutApiV4ProjectsIdRepositoryFilesFilePath(ctx, id, filePath, putApiV4ProjectsIdRepositoryFilesFilePath)


Update existing file in repository

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The project ID | 
  **filePath** | **string**| The url encoded path to the file. | 
  **putApiV4ProjectsIdRepositoryFilesFilePath** | [**PutApiV4ProjectsIdRepositoryFilesFilePath**](PutApiV4ProjectsIdRepositoryFilesFilePath.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdRepositorySubmodulesSubmodule**
> ApiEntitiesCommitDetail PutApiV4ProjectsIdRepositorySubmodulesSubmodule(ctx, id, submodule, putApiV4ProjectsIdRepositorySubmodulesSubmodule)


Update existing submodule reference in repository

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of a project | 
  **submodule** | **string**| Url encoded full path to submodule. | 
  **putApiV4ProjectsIdRepositorySubmodulesSubmodule** | [**PutApiV4ProjectsIdRepositorySubmodulesSubmodule**](PutApiV4ProjectsIdRepositorySubmodulesSubmodule.md)|  | 

### Return type

[**ApiEntitiesCommitDetail**](API_Entities_CommitDetail.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdTransfer**
> ApiEntitiesProject PutApiV4ProjectsIdTransfer(ctx, id, putApiV4ProjectsIdTransfer)


Transfer a project to a new namespace

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **putApiV4ProjectsIdTransfer** | [**PutApiV4ProjectsIdTransfer**](PutApiV4ProjectsIdTransfer.md)|  | 

### Return type

[**ApiEntitiesProject**](API_Entities_Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4ProjectsIdTriggersTriggerId**
> ApiEntitiesTrigger PutApiV4ProjectsIdTriggersTriggerId(ctx, id, triggerId, putApiV4ProjectsIdTriggersTriggerId)


Update a trigger token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID or URL-encoded path of the project | 
  **triggerId** | **int32**| The trigger token ID | 
  **putApiV4ProjectsIdTriggersTriggerId** | [**PutApiV4ProjectsIdTriggersTriggerId**](PutApiV4ProjectsIdTriggersTriggerId.md)|  | 

### Return type

[**ApiEntitiesTrigger**](API_Entities_Trigger.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

