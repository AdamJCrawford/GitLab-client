/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// API_Entities_System_BroadcastMessage model
type ApiEntitiesSystemBroadcastMessage struct {
	Id                 string `json:"id,omitempty"`
	Message            string `json:"message,omitempty"`
	StartsAt           string `json:"starts_at,omitempty"`
	EndsAt             string `json:"ends_at,omitempty"`
	Color              string `json:"color,omitempty"`
	Font               string `json:"font,omitempty"`
	TargetAccessLevels string `json:"target_access_levels,omitempty"`
	TargetPath         string `json:"target_path,omitempty"`
	BroadcastType      string `json:"broadcast_type,omitempty"`
	Dismissable        string `json:"dismissable,omitempty"`
	Active             string `json:"active,omitempty"`
}
