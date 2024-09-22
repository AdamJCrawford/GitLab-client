# PutApiV4ApplicationPlanLimits

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanName** | **string** | Name of the plan to update | [default to null]
**CiInstanceLevelVariables** | **int32** | Maximum number of Instance-level CI/CD variables that can be defined | [optional] [default to null]
**CiPipelineSize** | **int32** | Maximum number of jobs in a single pipeline | [optional] [default to null]
**CiActiveJobs** | **int32** | Total number of jobs in currently active pipelines | [optional] [default to null]
**CiProjectSubscriptions** | **int32** | Maximum number of pipeline subscriptions to and from a project | [optional] [default to null]
**CiPipelineSchedules** | **int32** | Maximum number of pipeline schedules | [optional] [default to null]
**CiNeedsSizeLimit** | **int32** | Maximum number of needs dependencies that a job can have | [optional] [default to null]
**CiRegisteredGroupRunners** | **int32** | Maximum number of runners registered per group | [optional] [default to null]
**CiRegisteredProjectRunners** | **int32** | Maximum number of runners registered per project | [optional] [default to null]
**ConanMaxFileSize** | **int32** | Maximum Conan package file size in bytes | [optional] [default to null]
**EnforcementLimit** | **int32** | Maximum storage size for the root namespace enforcement in MiB | [optional] [default to null]
**GenericPackagesMaxFileSize** | **int32** | Maximum generic package file size in bytes | [optional] [default to null]
**HelmMaxFileSize** | **int32** | Maximum Helm chart file size in bytes | [optional] [default to null]
**MavenMaxFileSize** | **int32** | Maximum Maven package file size in bytes | [optional] [default to null]
**NotificationLimit** | **int32** | Maximum storage size for the root namespace notifications in MiB | [optional] [default to null]
**NpmMaxFileSize** | **int32** | Maximum NPM package file size in bytes | [optional] [default to null]
**NugetMaxFileSize** | **int32** | Maximum NuGet package file size in bytes | [optional] [default to null]
**PypiMaxFileSize** | **int32** | Maximum PyPI package file size in bytes | [optional] [default to null]
**TerraformModuleMaxFileSize** | **int32** | Maximum Terraform Module package file size in bytes | [optional] [default to null]
**StorageSizeLimit** | **int32** | Maximum storage size for the root namespace in MiB | [optional] [default to null]
**PipelineHierarchySize** | **int32** | Maximum number of downstream pipelines in a pipeline&#39;s hierarchy tree | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


