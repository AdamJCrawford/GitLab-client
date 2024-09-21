/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Edit group cluster
type PutApiV4GroupsIdClustersClusterId struct {
	// Cluster name
	Name string `json:"name,omitempty"`
	// Determines if cluster is active or not
	Enabled bool `json:"enabled,omitempty"`
	// Cluster base domain
	Domain string `json:"domain,omitempty"`
	// The associated environment to the cluster
	EnvironmentScope string `json:"environment_scope,omitempty"`
	// Deploy each environment to a separate Kubernetes namespace
	NamespacePerEnvironment bool `json:"namespace_per_environment,omitempty"`
	// The ID of the management project
	ManagementProjectId int32 `json:"management_project_id,omitempty"`
	// Determines if GitLab will manage namespaces and service accounts for this cluster
	Managed                      bool                                                           `json:"managed,omitempty"`
	PlatformKubernetesAttributes *PutApiV4GroupsIdClustersClusterIdPlatformKubernetesAttributes `json:"platform_kubernetes_attributes,omitempty"`
}