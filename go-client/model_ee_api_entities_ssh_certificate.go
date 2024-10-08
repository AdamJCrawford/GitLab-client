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

// EE_API_Entities_SshCertificate model
type EeApiEntitiesSshCertificate struct {
	Id        int32     `json:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Key       string    `json:"key,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
