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

// Updates a group or project invitation.
type PutApiV4GroupsIdInvitationsEmail struct {
	// A valid access level (defaults: `30`, developer access level)
	AccessLevel int32 `json:"access_level,omitempty"`
	// Date string in ISO 8601 format (`YYYY-MM-DDTHH:MM:SSZ`)
	ExpiresAt time.Time `json:"expires_at,omitempty"`
}
