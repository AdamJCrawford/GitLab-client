/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Rotate a resource access token
type PostApiV4ProjectsIdAccessTokensTokenIdRotate struct {
	// The expiration date of the token
	ExpiresAt string `json:"expires_at,omitempty"`
}
