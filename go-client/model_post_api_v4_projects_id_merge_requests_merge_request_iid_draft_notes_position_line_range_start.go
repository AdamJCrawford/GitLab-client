/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type PostApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesPositionLineRangeStart struct {
	// Start line code for multi-line note
	LineCode string `json:"line_code,omitempty"`
	// Start line type for multi-line note
	Type_ string `json:"type,omitempty"`
	// Start old_line line number
	OldLine string `json:"old_line,omitempty"`
	// Start new_line line number
	NewLine string `json:"new_line,omitempty"`
}
