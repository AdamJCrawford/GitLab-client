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
	"time"
)

// API_Entities_PackageFile model
type ApiEntitiesPackageFile struct {
	Id         int32                       `json:"id,omitempty"`
	PackageId  int32                       `json:"package_id,omitempty"`
	CreatedAt  time.Time                   `json:"created_at,omitempty"`
	FileName   string                      `json:"file_name,omitempty"`
	Size       int32                       `json:"size,omitempty"`
	FileMd5    string                      `json:"file_md5,omitempty"`
	FileSha1   string                      `json:"file_sha1,omitempty"`
	FileSha256 string                      `json:"file_sha256,omitempty"`
	Pipelines  *ApiEntitiesPackagePipeline `json:"pipelines,omitempty"`
}
