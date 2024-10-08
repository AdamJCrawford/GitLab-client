/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Update an instance-level variable
type PutApiV4AdminCiVariablesKey struct {
	// The description of the variable
	Description string `json:"description,omitempty"`
	// The value of a variable
	Value string `json:"value,omitempty"`
	// Whether the variable is protected
	Protected bool `json:"protected,omitempty"`
	// Whether the variable is masked
	Masked bool `json:"masked,omitempty"`
	// Whether the variable will be expanded
	Raw bool `json:"raw,omitempty"`
	// The type of a variable. Available types are: env_var (default) and file
	VariableType string `json:"variable_type,omitempty"`
}
