/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Add new email address to the currently authenticated user
type PostApiV4UserEmails struct {
	// The new email
	Email string `json:"email"`
}