/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

import (
	"time"
)

// API_Entities_ResourceMilestoneEvent model
type ApiEntitiesResourceMilestoneEvent struct {
	Id           int32                 `json:"id,omitempty"`
	User         *ApiEntitiesUserBasic `json:"user,omitempty"`
	CreatedAt    time.Time             `json:"created_at,omitempty"`
	ResourceType string                `json:"resource_type,omitempty"`
	ResourceId   int32                 `json:"resource_id,omitempty"`
	Milestone    *ApiEntitiesMilestone `json:"milestone,omitempty"`
	Action       string                `json:"action,omitempty"`
	State        string                `json:"state,omitempty"`
}
