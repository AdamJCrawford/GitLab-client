/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type ApiEntitiesReleaseLinks struct {
	ClosedIssuesUrl        string `json:"closed_issues_url,omitempty"`
	ClosedMergeRequestsUrl string `json:"closed_merge_requests_url,omitempty"`
	EditUrl                string `json:"edit_url,omitempty"`
	MergedMergeRequestsUrl string `json:"merged_merge_requests_url,omitempty"`
	OpenedIssuesUrl        string `json:"opened_issues_url,omitempty"`
	OpenedMergeRequestsUrl string `json:"opened_merge_requests_url,omitempty"`
	Self                   string `json:"self,omitempty"`
}
