/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// API_Entities_Metadata model
type ApiEntitiesMetadata struct {
	Version    string                  `json:"version,omitempty"`
	Revision   string                  `json:"revision,omitempty"`
	Kas        *ApiEntitiesMetadataKas `json:"kas,omitempty"`
	Enterprise bool                    `json:"enterprise,omitempty"`
}