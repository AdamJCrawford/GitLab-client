/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Share the project with a group
type PostApiV4ProjectsIdShare struct {
	// The ID of a group
	GroupId int32 `json:"group_id"`
	// The group access level
	GroupAccess int32 `json:"group_access"`
	// Share expiration date
	ExpiresAt string `json:"expires_at,omitempty"`
}