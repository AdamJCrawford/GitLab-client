/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Set the status of the current user
type PatchApiV4UserStatus struct {
	// The emoji to set on the status
	Emoji string `json:"emoji,omitempty"`
	// The status message to set
	Message string `json:"message,omitempty"`
	// The availability of user to set
	Availability string `json:"availability,omitempty"`
	// Automatically clear emoji, message and availability fields after a certain time
	ClearStatusAfter string `json:"clear_status_after,omitempty"`
}