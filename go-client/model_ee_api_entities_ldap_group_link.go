/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type EeApiEntitiesLdapGroupLink struct {
	Cn          string `json:"cn,omitempty"`
	GroupAccess int32  `json:"group_access,omitempty"`
	Provider    string `json:"provider,omitempty"`
	Filter      string `json:"filter,omitempty"`
}
