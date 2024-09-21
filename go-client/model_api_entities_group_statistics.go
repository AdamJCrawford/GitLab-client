/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ApiEntitiesGroupStatistics struct {
	StorageSize           string `json:"storage_size,omitempty"`
	RepositorySize        string `json:"repository_size,omitempty"`
	WikiSize              string `json:"wiki_size,omitempty"`
	LfsObjectsSize        string `json:"lfs_objects_size,omitempty"`
	JobArtifactsSize      string `json:"job_artifacts_size,omitempty"`
	PipelineArtifactsSize string `json:"pipeline_artifacts_size,omitempty"`
	PackagesSize          string `json:"packages_size,omitempty"`
	SnippetsSize          string `json:"snippets_size,omitempty"`
	UploadsSize           string `json:"uploads_size,omitempty"`
}