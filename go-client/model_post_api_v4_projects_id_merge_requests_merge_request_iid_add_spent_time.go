/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Add spent time for a merge_request
type PostApiV4ProjectsIdMergeRequestsMergeRequestIidAddSpentTime struct {
	// The duration in human format.
	Duration string `json:"duration"`
}
