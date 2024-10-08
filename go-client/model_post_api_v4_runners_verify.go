/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Validate authentication credentials
type PostApiV4RunnersVerify struct {
	// The runner's authentication token
	Token string `json:"token"`
	// The runner's system identifier
	SystemId string `json:"system_id,omitempty"`
}
