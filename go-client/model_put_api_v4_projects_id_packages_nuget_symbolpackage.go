/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

import (
	"os"
)

// The NuGet Symbol Package Publish endpoint
type PutApiV4ProjectsIdPackagesNugetSymbolpackage struct {
	// The package file to be published (generated by Multipart middleware)
	Package_ **os.File `json:"package"`
}