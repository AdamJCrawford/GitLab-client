/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Trigger a global slack command
type PostApiV4SlackTrigger struct {
	// Text of the slack command
	Text string `json:"text"`
}
