/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

import (
	"time"
)

// Create a release
type PostApiV4ProjectsIdReleases struct {
	// The tag where the release is created from
	TagName string `json:"tag_name"`
	// Message to use if creating a new annotated tag
	TagMessage string `json:"tag_message,omitempty"`
	// The release name
	Name string `json:"name,omitempty"`
	// The description of the release. You can use Markdown
	Description string `json:"description,omitempty"`
	// If a tag specified in `tag_name` doesn't exist, the release is created from `ref` and tagged with `tag_name`. It can be a commit SHA, another tag name, or a branch name.
	Ref    string                             `json:"ref,omitempty"`
	Assets *PostApiV4ProjectsIdReleasesAssets `json:"assets,omitempty"`
	// The title of each milestone the release is associated with. GitLab Premium customers can specify group milestones. Cannot be combined with `milestone_ids` parameter.
	Milestones []string `json:"milestones,omitempty"`
	// The ID of each milestone the release is associated with. GitLab Premium customers can specify group milestones. Cannot be combined with `milestones` parameter.
	MilestoneIds string `json:"milestone_ids,omitempty"`
	// Date and time for the release. Defaults to the current time. Expected in ISO 8601 format (`2019-03-15T08:00:00Z`). Only provide this field if creating an upcoming or historical release.
	ReleasedAt time.Time `json:"released_at,omitempty"`
	// If true, the release will be published to the CI catalog. This parameter is for internal use only and will be removed in a future release. If the feature flag ci_release_cli_catalog_publish_option is disabled, this parameter will be ignored and the release will published to the CI catalog as it was before this parameter was introduced.
	LegacyCatalogPublish bool `json:"legacy_catalog_publish,omitempty"`
}
