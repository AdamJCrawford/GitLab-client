/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Mark the migration as successfully executed
type PostApiV4AdminMigrationsTimestampMark struct {
	// The name of the database
	Database string `json:"database,omitempty"`
}
