/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Edit a pipeline schedule variable
type PutApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKey struct {
	// The value of the variable
	Value string `json:"value,omitempty"`
	// The type of variable, must be one of env_var or file
	VariableType string `json:"variable_type,omitempty"`
}
