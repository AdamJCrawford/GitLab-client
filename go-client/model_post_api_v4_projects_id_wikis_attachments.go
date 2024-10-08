/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

import (
	"os"
)

// Upload an attachment to the wiki repository
type PostApiV4ProjectsIdWikisAttachments struct {
	// The attachment file to be uploaded
	File **os.File `json:"file"`
	// The name of the branch
	Branch string `json:"branch,omitempty"`
}
