/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// API_Entities_Badge model
type ApiEntitiesBadge struct {
	Name             string `json:"name,omitempty"`
	LinkUrl          string `json:"link_url,omitempty"`
	ImageUrl         string `json:"image_url,omitempty"`
	RenderedLinkUrl  string `json:"rendered_link_url,omitempty"`
	RenderedImageUrl string `json:"rendered_image_url,omitempty"`
	Id               string `json:"id,omitempty"`
	Kind             string `json:"kind,omitempty"`
}