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

// API_Entities_Release model
type ApiEntitiesRelease struct {
	Name            string                         `json:"name,omitempty"`
	TagName         string                         `json:"tag_name,omitempty"`
	Description     string                         `json:"description,omitempty"`
	CreatedAt       time.Time                      `json:"created_at,omitempty"`
	ReleasedAt      time.Time                      `json:"released_at,omitempty"`
	UpcomingRelease bool                           `json:"upcoming_release,omitempty"`
	DescriptionHtml string                         `json:"description_html,omitempty"`
	Author          *ApiEntitiesUserBasic          `json:"author,omitempty"`
	Commit          *ApiEntitiesCommit             `json:"commit,omitempty"`
	Milestones      *ApiEntitiesMilestoneWithStats `json:"milestones,omitempty"`
	CommitPath      string                         `json:"commit_path,omitempty"`
	TagPath         string                         `json:"tag_path,omitempty"`
	Assets          *ApiEntitiesReleaseAssets      `json:"assets,omitempty"`
	Evidences       *ApiEntitiesReleasesEvidence   `json:"evidences,omitempty"`
	Links           *ApiEntitiesReleaseLinks       `json:"_links,omitempty"`
}
