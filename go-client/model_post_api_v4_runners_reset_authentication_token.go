/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Reset runner authentication token with current token
type PostApiV4RunnersResetAuthenticationToken struct {
	// The current authentication token of the runner
	Token string `json:"token"`
}
