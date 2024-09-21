/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Adds a badge to a group.
type PostApiV4GroupsIdBadges struct {
	// URL of the badge link
	LinkUrl string `json:"link_url"`
	// URL of the badge image
	ImageUrl string `json:"image_url"`
	// Name for the badge
	Name string `json:"name,omitempty"`
}
