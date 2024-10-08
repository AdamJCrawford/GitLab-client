/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Create/Edit Harbor integration
type PutApiV4ProjectsIdServicesHarbor struct {
	// The base URL to the Harbor instance linked to the GitLab project. For example, `https://demo.goharbor.io`.
	Url string `json:"url"`
	// The name of the project in the Harbor instance. For example, `testproject`.
	ProjectName string `json:"project_name"`
	// The username created in the Harbor interface.
	Username string `json:"username"`
	// The password of the user.
	Password string `json:"password"`
}
