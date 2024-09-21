/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Create a personal access token with limited scopes for the currently authenticated user
type PostApiV4UserPersonalAccessTokens struct {
	// The name of the personal access token
	Name string `json:"name"`
	// The array of scopes of the personal access token
	Scopes []string `json:"scopes"`
	// The expiration date in the format YEAR-MONTH-DAY of the personal access token
	ExpiresAt string `json:"expires_at,omitempty"`
}
