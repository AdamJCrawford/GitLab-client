/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ApiEntitiesNugetPackagesMetadataItem struct {
	Id    string                            `json:"@id,omitempty"`
	Lower string                            `json:"lower,omitempty"`
	Upper string                            `json:"upper,omitempty"`
	Count int32                             `json:"count,omitempty"`
	Items []ApiEntitiesNugetPackageMetadata `json:"items,omitempty"`
}
