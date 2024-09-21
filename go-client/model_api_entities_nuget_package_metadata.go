/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// API_Entities_Nuget_PackageMetadata model
type ApiEntitiesNugetPackageMetadata struct {
	Id             string                                       `json:"@id,omitempty"`
	PackageContent string                                       `json:"packageContent,omitempty"`
	CatalogEntry   *ApiEntitiesNugetPackageMetadataCatalogEntry `json:"catalogEntry,omitempty"`
}
