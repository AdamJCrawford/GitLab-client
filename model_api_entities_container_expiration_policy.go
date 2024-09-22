/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ApiEntitiesContainerExpirationPolicy struct {
	Cadence       string `json:"cadence,omitempty"`
	Enabled       string `json:"enabled,omitempty"`
	KeepN         string `json:"keep_n,omitempty"`
	OlderThan     string `json:"older_than,omitempty"`
	NameRegex     string `json:"name_regex,omitempty"`
	NameRegexKeep string `json:"name_regex_keep,omitempty"`
	NextRunAt     string `json:"next_run_at,omitempty"`
}