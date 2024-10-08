/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type ApiEntitiesPushEventPayload struct {
	CommitCount int32  `json:"commit_count,omitempty"`
	Action      string `json:"action,omitempty"`
	RefType     string `json:"ref_type,omitempty"`
	CommitFrom  string `json:"commit_from,omitempty"`
	CommitTo    string `json:"commit_to,omitempty"`
	Ref         string `json:"ref,omitempty"`
	CommitTitle string `json:"commit_title,omitempty"`
	RefCount    int32  `json:"ref_count,omitempty"`
}
