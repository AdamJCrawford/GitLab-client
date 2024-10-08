/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// API_Entities_Compare model
type ApiEntitiesCompare struct {
	Commit         *ApiEntitiesCommit  `json:"commit,omitempty"`
	Commits        []ApiEntitiesCommit `json:"commits,omitempty"`
	Diffs          []ApiEntitiesDiff   `json:"diffs,omitempty"`
	CompareTimeout bool                `json:"compare_timeout,omitempty"`
	CompareSameRef bool                `json:"compare_same_ref,omitempty"`
	WebUrl         string              `json:"web_url,omitempty"`
}
