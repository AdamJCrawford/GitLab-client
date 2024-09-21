/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

// API_Entities_CommitNote model
type ApiEntitiesCommitNote struct {
	Note      string                `json:"note,omitempty"`
	Path      string                `json:"path,omitempty"`
	Line      int32                 `json:"line,omitempty"`
	LineType  string                `json:"line_type,omitempty"`
	Author    *ApiEntitiesUserBasic `json:"author,omitempty"`
	CreatedAt time.Time             `json:"created_at,omitempty"`
}
