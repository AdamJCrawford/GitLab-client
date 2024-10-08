/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Update Error Tracking project settings. Available in GitLab 15.10 and later.
type PutApiV4ProjectsIdErrorTrackingSettings struct {
	// Pass true to enable the configured Error Tracking settings or false to disable it.
	Active bool `json:"active"`
	// Pass true to enable the integrated Error Tracking backend.
	Integrated bool `json:"integrated"`
}
