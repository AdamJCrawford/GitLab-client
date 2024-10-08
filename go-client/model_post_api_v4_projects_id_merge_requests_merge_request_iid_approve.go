/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Approve a merge request
type PostApiV4ProjectsIdMergeRequestsMergeRequestIidApprove struct {
	// When present, must have the HEAD SHA of the source branch
	Sha string `json:"sha,omitempty"`
	// Current user's password if project is set to require explicit auth on approval
	ApprovalPassword string `json:"approval_password,omitempty"`
}
