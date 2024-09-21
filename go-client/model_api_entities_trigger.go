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

// API_Entities_Trigger model
type ApiEntitiesTrigger struct {
	Id          int32                 `json:"id,omitempty"`
	Token       string                `json:"token,omitempty"`
	Description string                `json:"description,omitempty"`
	CreatedAt   time.Time             `json:"created_at,omitempty"`
	UpdatedAt   time.Time             `json:"updated_at,omitempty"`
	LastUsed    time.Time             `json:"last_used,omitempty"`
	Owner       *ApiEntitiesUserBasic `json:"owner,omitempty"`
}
