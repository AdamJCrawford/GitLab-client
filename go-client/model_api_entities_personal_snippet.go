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

// API_Entities_PersonalSnippet model
type ApiEntitiesPersonalSnippet struct {
	Id                int32                 `json:"id,omitempty"`
	Title             string                `json:"title,omitempty"`
	Description       string                `json:"description,omitempty"`
	Visibility        string                `json:"visibility,omitempty"`
	Author            *ApiEntitiesUserBasic `json:"author,omitempty"`
	CreatedAt         time.Time             `json:"created_at,omitempty"`
	UpdatedAt         time.Time             `json:"updated_at,omitempty"`
	ProjectId         int32                 `json:"project_id,omitempty"`
	WebUrl            string                `json:"web_url,omitempty"`
	RawUrl            string                `json:"raw_url,omitempty"`
	SshUrlToRepo      string                `json:"ssh_url_to_repo,omitempty"`
	HttpUrlToRepo     string                `json:"http_url_to_repo,omitempty"`
	FileName          string                `json:"file_name,omitempty"`
	Files             []string              `json:"files,omitempty"`
	Imported          bool                  `json:"imported,omitempty"`
	ImportedFrom      string                `json:"imported_from,omitempty"`
	RepositoryStorage string                `json:"repository_storage,omitempty"`
}
