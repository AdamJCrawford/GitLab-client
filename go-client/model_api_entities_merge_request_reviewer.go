/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// API_Entities_MergeRequestReviewer model
type ApiEntitiesMergeRequestReviewer struct {
	User      *ApiEntitiesUserBasic `json:"user,omitempty"`
	State     string                `json:"state,omitempty"`
	CreatedAt string                `json:"created_at,omitempty"`
}
