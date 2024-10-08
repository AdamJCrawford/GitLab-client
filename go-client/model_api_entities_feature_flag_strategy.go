/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type ApiEntitiesFeatureFlagStrategy struct {
	Id         int32                                `json:"id,omitempty"`
	Name       string                               `json:"name,omitempty"`
	Parameters string                               `json:"parameters,omitempty"`
	Scopes     *ApiEntitiesFeatureFlagScope         `json:"scopes,omitempty"`
	UserList   *ApiEntitiesFeatureFlagBasicUserList `json:"user_list,omitempty"`
}
