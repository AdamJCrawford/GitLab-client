/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Authorize Debian package upload
type PutApiV4ProjectsIdPackagesDebianFileNameAuthorize struct {
	// The Debian Codename or Suite
	Distribution string `json:"distribution,omitempty"`
	// The Debian Component
	Component string `json:"component"`
}
