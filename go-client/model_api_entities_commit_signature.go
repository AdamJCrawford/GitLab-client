/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// API_Entities_CommitSignature model
type ApiEntitiesCommitSignature struct {
	SignatureType string `json:"signature_type,omitempty"`
	Signature     string `json:"signature,omitempty"`
	CommitSource  string `json:"commit_source,omitempty"`
}
