/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Create merge request context commits
type PostApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommits struct {
	// The context commits’ SHA.
	Commits []string `json:"commits"`
}