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
	"time"
)

// API_Entities_Ci_SecureFile model
type ApiEntitiesCiSecureFile struct {
	Id                int32       `json:"id,omitempty"`
	Name              string      `json:"name,omitempty"`
	Checksum          string      `json:"checksum,omitempty"`
	ChecksumAlgorithm string      `json:"checksum_algorithm,omitempty"`
	CreatedAt         time.Time   `json:"created_at,omitempty"`
	ExpiresAt         time.Time   `json:"expires_at,omitempty"`
	Metadata          interface{} `json:"metadata,omitempty"`
	FileExtension     string      `json:"file_extension,omitempty"`
}