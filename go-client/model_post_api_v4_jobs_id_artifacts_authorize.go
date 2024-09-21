/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Authorize uploading job artifact
type PostApiV4JobsIdArtifactsAuthorize struct {
	// Job's authentication token
	Token string `json:"token,omitempty"`
	// Size of artifact file
	Filesize int32 `json:"filesize,omitempty"`
	// The type of artifact
	ArtifactType string `json:"artifact_type,omitempty"`
}
