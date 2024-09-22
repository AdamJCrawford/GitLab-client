/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// EE_API_Entities_AuditEvent model
type EeApiEntitiesAuditEvent struct {
	Id         string `json:"id,omitempty"`
	AuthorId   string `json:"author_id,omitempty"`
	EntityId   string `json:"entity_id,omitempty"`
	EntityType string `json:"entity_type,omitempty"`
	Details    string `json:"details,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
}