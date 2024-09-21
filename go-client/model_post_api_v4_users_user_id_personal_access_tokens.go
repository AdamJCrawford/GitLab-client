/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Create a personal access token. Available only for admins.
type PostApiV4UsersUserIdPersonalAccessTokens struct {
	// The name of the personal access token
	Name string `json:"name"`
	// The array of scopes of the personal access token
	Scopes []string `json:"scopes"`
	// The expiration date in the format YEAR-MONTH-DAY of the personal access token
	ExpiresAt string `json:"expires_at,omitempty"`
}