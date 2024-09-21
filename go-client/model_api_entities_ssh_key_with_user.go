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

// API_Entities_SSHKeyWithUser model
type ApiEntitiesSshKeyWithUser struct {
	Id        int32                  `json:"id,omitempty"`
	Title     string                 `json:"title,omitempty"`
	CreatedAt time.Time              `json:"created_at,omitempty"`
	ExpiresAt time.Time              `json:"expires_at,omitempty"`
	Key       string                 `json:"key,omitempty"`
	UsageType string                 `json:"usage_type,omitempty"`
	User      *ApiEntitiesUserPublic `json:"user,omitempty"`
}
