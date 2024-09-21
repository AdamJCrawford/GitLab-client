# PutApiV4AdminClustersClusterId

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Cluster name | [optional] [default to null]
**Enabled** | **bool** | Enable or disable Gitlab&#39;s connection to your Kubernetes cluster | [optional] [default to null]
**EnvironmentScope** | **string** | The associated environment to the cluster | [optional] [default to null]
**NamespacePerEnvironment** | **bool** | Deploy each environment to a separate Kubernetes namespace | [optional] [default to null]
**Domain** | **string** | Cluster base domain | [optional] [default to null]
**ManagementProjectId** | **int32** | The ID of the management project | [optional] [default to null]
**Managed** | **bool** | Determines if GitLab will manage namespaces and service accounts for this cluster | [optional] [default to null]
**PlatformKubernetesAttributes** | [***PutApiV4ProjectsIdClustersClusterIdPlatformKubernetesAttributes**](putApiV4ProjectsIdClustersClusterId_platform_kubernetes_attributes.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


