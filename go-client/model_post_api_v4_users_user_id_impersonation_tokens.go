/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Create a impersonation token. Available only for admins.
type PostApiV4UsersUserIdImpersonationTokens struct {
	// The name of the impersonation token
	Name string `json:"name"`
	// The expiration date in the format YEAR-MONTH-DAY of the impersonation token
	ExpiresAt string `json:"expires_at,omitempty"`
	// The array of scopes of the impersonation token
	Scopes []string `json:"scopes,omitempty"`
}