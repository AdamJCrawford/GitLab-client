/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type PostApiV4ProjectsIdReleasesAssetsLinks struct {
	// The name of the link. Link names must be unique within the release
	Name string `json:"name"`
	// The URL of the link. Link URLs must be unique within the release
	Url string `json:"url"`
	// Optional path for a direct asset link
	DirectAssetPath string `json:"direct_asset_path,omitempty"`
	// Deprecated: optional path for a direct asset link
	Filepath string `json:"filepath,omitempty"`
	// The type of the link: `other`, `runbook`, `image`, `package`. Defaults to `other`
	LinkType string `json:"link_type,omitempty"`
}
