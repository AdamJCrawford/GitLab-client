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

// Upload a gem
type PostApiV4ProjectsIdPackagesRubygemsApiV1Gems struct {
	// The package file to be published (generated by Multipart middleware)
	File **os.File `json:"file"`
}
