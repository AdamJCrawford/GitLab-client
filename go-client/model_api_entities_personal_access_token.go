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

// API_Entities_PersonalAccessToken model
type ApiEntitiesPersonalAccessToken struct {
	Id         int32     `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	Revoked    bool      `json:"revoked,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	Scopes     []string  `json:"scopes,omitempty"`
	UserId     int32     `json:"user_id,omitempty"`
	LastUsedAt time.Time `json:"last_used_at,omitempty"`
	Active     bool      `json:"active,omitempty"`
	ExpiresAt  time.Time `json:"expires_at,omitempty"`
}
