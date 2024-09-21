/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Approves an access request for the given user.
type PutApiV4ProjectsIdAccessRequestsUserIdApprove struct {
	// A valid access level (defaults: `30`, the Developer role)
	AccessLevel int32 `json:"access_level,omitempty"`
}
