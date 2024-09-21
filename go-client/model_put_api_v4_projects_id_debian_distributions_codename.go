/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Update a Debian Distribution
type PutApiV4ProjectsIdDebianDistributionsCodename struct {
	// The Debian Suite
	Suite string `json:"suite,omitempty"`
	// The Debian Origin
	Origin string `json:"origin,omitempty"`
	// The Debian Label
	Label string `json:"label,omitempty"`
	// The Debian Version
	Version string `json:"version,omitempty"`
	// The Debian Description
	Description string `json:"description,omitempty"`
	// The duration before the Release file should be considered expired by the client
	ValidTimeDurationSeconds int32 `json:"valid_time_duration_seconds,omitempty"`
	// The list of Components
	Components []string `json:"components,omitempty"`
	// The list of Architectures
	Architectures []string `json:"architectures,omitempty"`
}
