/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Create/Edit Diffblue Cover integration
type PutApiV4ProjectsIdIntegrationsDiffblueCover struct {
	// Diffblue Cover license key.
	DiffblueLicenseKey string `json:"diffblue_license_key"`
	// Access token name used by Diffblue Cover in pipelines.
	DiffblueAccessTokenName string `json:"diffblue_access_token_name"`
	// Access token secret used by Diffblue Cover in pipelines.
	DiffblueAccessTokenSecret string `json:"diffblue_access_token_secret"`
}
