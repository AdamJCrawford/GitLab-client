/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Edit a pipeline schedule
type PutApiV4ProjectsIdPipelineSchedulesPipelineScheduleId struct {
	// The description of pipeline schedule
	Description string `json:"description,omitempty"`
	// The branch/tag name will be triggered
	Ref string `json:"ref,omitempty"`
	// The cron
	Cron string `json:"cron,omitempty"`
	// The timezone
	CronTimezone string `json:"cron_timezone,omitempty"`
	// The activation of pipeline schedule
	Active bool `json:"active,omitempty"`
}
