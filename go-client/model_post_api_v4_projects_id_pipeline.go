/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Create a new pipeline
type PostApiV4ProjectsIdPipeline struct {
	// Reference
	Ref string `json:"ref"`
	// Array of variables available in the pipeline
	Variables []PostApiV4ProjectsIdPipelineVariables `json:"variables,omitempty"`
}
