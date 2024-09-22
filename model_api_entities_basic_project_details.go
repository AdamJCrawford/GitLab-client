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

// API_Entities_BasicProjectDetails model
type ApiEntitiesBasicProjectDetails struct {
	Id                int32                       `json:"id,omitempty"`
	Description       string                      `json:"description,omitempty"`
	Name              string                      `json:"name,omitempty"`
	NameWithNamespace string                      `json:"name_with_namespace,omitempty"`
	Path              string                      `json:"path,omitempty"`
	PathWithNamespace string                      `json:"path_with_namespace,omitempty"`
	CreatedAt         time.Time                   `json:"created_at,omitempty"`
	DefaultBranch     string                      `json:"default_branch,omitempty"`
	TagList           []string                    `json:"tag_list,omitempty"`
	Topics            []string                    `json:"topics,omitempty"`
	SshUrlToRepo      string                      `json:"ssh_url_to_repo,omitempty"`
	HttpUrlToRepo     string                      `json:"http_url_to_repo,omitempty"`
	WebUrl            string                      `json:"web_url,omitempty"`
	ReadmeUrl         string                      `json:"readme_url,omitempty"`
	ForksCount        int32                       `json:"forks_count,omitempty"`
	LicenseUrl        string                      `json:"license_url,omitempty"`
	License           *ApiEntitiesLicenseBasic    `json:"license,omitempty"`
	AvatarUrl         string                      `json:"avatar_url,omitempty"`
	StarCount         int32                       `json:"star_count,omitempty"`
	LastActivityAt    time.Time                   `json:"last_activity_at,omitempty"`
	Namespace         *ApiEntitiesNamespaceBasic  `json:"namespace,omitempty"`
	CustomAttributes  *ApiEntitiesCustomAttribute `json:"custom_attributes,omitempty"`
	RepositoryStorage string                      `json:"repository_storage,omitempty"`
}