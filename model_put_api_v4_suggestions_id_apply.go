/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Apply suggestion patch in the Merge Request it was created
type PutApiV4SuggestionsIdApply struct {
	// A custom commit message to use instead of the default generated message or the project's default message
	CommitMessage string `json:"commit_message,omitempty"`
}