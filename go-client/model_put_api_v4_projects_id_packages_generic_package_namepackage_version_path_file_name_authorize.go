/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Workhorse authorize generic package file
type PutApiV4ProjectsIdPackagesGenericPackageNamepackageVersionPathFileNameAuthorize struct {
	// Package version
	PackageVersion string `json:"package_version"`
	// Package status
	Status string `json:"status,omitempty"`
	Path   int32  `json:"path"`
}
