/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"os"
)

// Upload a file
type PostApiV4ProjectsIdUploads struct {
	// The attachment file to be uploaded
	File **os.File `json:"file"`
}
