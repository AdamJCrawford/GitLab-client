/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Create a new variable in a project
type PostApiV4ProjectsIdVariables struct {
	// The key of a variable
	Key string `json:"key"`
	// The value of a variable
	Value string `json:"value"`
	// Whether the variable is protected
	Protected bool `json:"protected,omitempty"`
	// Whether the variable is masked
	Masked bool `json:"masked,omitempty"`
	// Whether the variable is masked and hidden
	MaskedAndHidden bool `json:"masked_and_hidden,omitempty"`
	// Whether the variable will be expanded
	Raw bool `json:"raw,omitempty"`
	// The type of the variable. Default: env_var
	VariableType string `json:"variable_type,omitempty"`
	// The environment_scope of the variable
	EnvironmentScope string `json:"environment_scope,omitempty"`
	// The description of the variable
	Description string `json:"description,omitempty"`
}
