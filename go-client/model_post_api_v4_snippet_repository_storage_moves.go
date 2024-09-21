/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Schedule bulk snippet repository storage moves
type PostApiV4SnippetRepositoryStorageMoves struct {
	// The source storage shard
	SourceStorageName string `json:"source_storage_name"`
	// The destination storage shard
	DestinationStorageName string `json:"destination_storage_name,omitempty"`
}
