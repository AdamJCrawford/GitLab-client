/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Runner's metadata
type PostApiV4RunnersInfo struct {
	// Runner's name
	Name string `json:"name,omitempty"`
	// Runner's version
	Version string `json:"version,omitempty"`
	// Runner's revision
	Revision string `json:"revision,omitempty"`
	// Runner's platform
	Platform string `json:"platform,omitempty"`
	// Runner's architecture
	Architecture string `json:"architecture,omitempty"`
}
