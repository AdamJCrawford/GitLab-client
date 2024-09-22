/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type PostApiV4ProjectsIdProtectedBranchesAllowedToPush struct {
	AccessLevel int32 `json:"access_level,omitempty"`
	UserId      int32 `json:"user_id,omitempty"`
	GroupId     int32 `json:"group_id,omitempty"`
	Id          int32 `json:"id,omitempty"`
	// Delete the object when true
	Destroy bool `json:"_destroy,omitempty"`
}