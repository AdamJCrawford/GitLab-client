/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

import (
	"time"
)

// API_Entities_Snippets_RepositoryStorageMove model
type ApiEntitiesSnippetsRepositoryStorageMove struct {
	Id                     int32                    `json:"id,omitempty"`
	CreatedAt              time.Time                `json:"created_at,omitempty"`
	State                  string                   `json:"state,omitempty"`
	SourceStorageName      string                   `json:"source_storage_name,omitempty"`
	DestinationStorageName string                   `json:"destination_storage_name,omitempty"`
	ErrorMessage           string                   `json:"error_message,omitempty"`
	Snippet                *ApiEntitiesBasicSnippet `json:"snippet,omitempty"`
}
