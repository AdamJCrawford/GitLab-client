/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Creates a storage limit exclusion for a Namespace
type PostApiV4NamespacesIdStorageLimitExclusion struct {
	// The reason the Namespace is being excluded
	Reason string `json:"reason"`
}
