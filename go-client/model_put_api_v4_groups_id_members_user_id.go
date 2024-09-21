/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

// Updates a member of a group or project.
type PutApiV4GroupsIdMembersUserId struct {
	// A valid access level
	AccessLevel int32 `json:"access_level"`
	// Date string in the format YEAR-MONTH-DAY
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// The ID of the Member Role to be updated
	MemberRoleId int32 `json:"member_role_id,omitempty"`
}
