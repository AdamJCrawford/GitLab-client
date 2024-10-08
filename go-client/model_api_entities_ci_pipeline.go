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

// API_Entities_Ci_Pipeline model
type ApiEntitiesCiPipeline struct {
	Id          int32                 `json:"id,omitempty"`
	Iid         int32                 `json:"iid,omitempty"`
	ProjectId   int32                 `json:"project_id,omitempty"`
	Sha         string                `json:"sha,omitempty"`
	Ref         string                `json:"ref,omitempty"`
	Status      string                `json:"status,omitempty"`
	Source      string                `json:"source,omitempty"`
	CreatedAt   time.Time             `json:"created_at,omitempty"`
	UpdatedAt   time.Time             `json:"updated_at,omitempty"`
	WebUrl      string                `json:"web_url,omitempty"`
	BeforeSha   string                `json:"before_sha,omitempty"`
	Tag         bool                  `json:"tag,omitempty"`
	YamlErrors  string                `json:"yaml_errors,omitempty"`
	User        *ApiEntitiesUserBasic `json:"user,omitempty"`
	StartedAt   time.Time             `json:"started_at,omitempty"`
	FinishedAt  time.Time             `json:"finished_at,omitempty"`
	CommittedAt time.Time             `json:"committed_at,omitempty"`
	// Time spent running in seconds
	Duration int32 `json:"duration,omitempty"`
	// Time spent enqueued in seconds
	QueuedDuration int32                 `json:"queued_duration,omitempty"`
	Coverage       float32               `json:"coverage,omitempty"`
	DetailedStatus *DetailedStatusEntity `json:"detailed_status,omitempty"`
}
