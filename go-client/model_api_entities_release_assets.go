/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ApiEntitiesReleaseAssets struct {
	Count   int32                      `json:"count,omitempty"`
	Sources *ApiEntitiesReleasesSource `json:"sources,omitempty"`
	Links   *ApiEntitiesReleasesLink   `json:"links,omitempty"`
}
