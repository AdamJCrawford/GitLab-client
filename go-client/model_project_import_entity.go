/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// ProjectImportEntity model
type ProjectImportEntity struct {
	Id                    int32  `json:"id,omitempty"`
	Name                  string `json:"name,omitempty"`
	FullPath              string `json:"full_path,omitempty"`
	FullName              string `json:"full_name,omitempty"`
	RefsUrl               string `json:"refs_url,omitempty"`
	ImportSource          string `json:"import_source,omitempty"`
	ImportStatus          string `json:"import_status,omitempty"`
	HumanImportStatusName string `json:"human_import_status_name,omitempty"`
	ProviderLink          string `json:"provider_link,omitempty"`
	ImportError           string `json:"import_error,omitempty"`
	ImportWarning         string `json:"import_warning,omitempty"`
	RelationType          string `json:"relation_type,omitempty"`
}
