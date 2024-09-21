# PostApiV4GroupsIdClustersUser

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Cluster name | [default to null]
**Enabled** | **bool** | Determines if cluster is active or not, defaults to true | [optional] [default to null]
**EnvironmentScope** | **string** | The associated environment to the cluster | [optional] [default to null]
**NamespacePerEnvironment** | **bool** | Deploy each environment to a separate Kubernetes namespace | [optional] [default to null]
**Domain** | **string** | Cluster base domain | [optional] [default to null]
**ManagementProjectId** | **int32** | The ID of the management project | [optional] [default to null]
**Managed** | **bool** | Determines if GitLab will manage namespaces and service accounts for this cluster, defaults to true | [optional] [default to null]
**PlatformKubernetesAttributes** | [***PostApiV4GroupsIdClustersUserPlatformKubernetesAttributes**](postApiV4GroupsIdClustersUser_platform_kubernetes_attributes.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

