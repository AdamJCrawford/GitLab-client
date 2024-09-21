# PutApiV4ProjectsIdIntegrationsDatadog

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **string** | API key used for authentication with Datadog | [default to null]
**DatadogSite** | **string** | The Datadog site to send data to. To send data to the EU site, use datadoghq.eu | [optional] [default to null]
**ApiUrl** | **string** | (Advanced) The full URL for your Datadog site | [optional] [default to null]
**ArchiveTraceEvents** | **bool** | When enabled, job logs are collected by Datadog and displayed along with pipeline execution traces. | [optional] [default to null]
**DatadogService** | **string** | Tag all data from this GitLab instance in Datadog. Useful when managing several self-managed deployments | [optional] [default to null]
**DatadogEnv** | **string** | For self-managed deployments, set the env tag for all the data sent to Datadog | [optional] [default to null]
**DatadogTags** | **string** | Custom tags in Datadog. Specify one tag per line in the format: \&quot;key:value\\nkey2:value2\&quot; | [optional] [default to null]
**PipelineEvents** | **bool** | Trigger event when a pipeline status changes. | [optional] [default to null]
**BuildEvents** | **bool** | Trigger event when a build is created. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


