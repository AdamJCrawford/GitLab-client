/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Create/Edit Pushover integration
type PutApiV4ProjectsIdIntegrationsPushover struct {
	// The application key
	ApiKey string `json:"api_key"`
	// The user key
	UserKey string `json:"user_key"`
	// The priority
	Priority string `json:"priority"`
	// Leave blank for all active devices
	Device string `json:"device"`
	// The sound of the notification
	Sound string `json:"sound"`
	// Trigger event for pushes to the repository.
	PushEvents bool `json:"push_events,omitempty"`
}
