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

// Set the avatar of the current user
type PutApiV4UserAvatar struct {
	// The avatar file (generated by Multipart middleware)
	Avatar **os.File `json:"avatar"`
}