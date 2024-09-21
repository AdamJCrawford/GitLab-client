/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// API_Entities_Ci_RunnerDetails model
type ApiEntitiesCiRunnerDetails struct {
	Id              int32                           `json:"id,omitempty"`
	Description     string                          `json:"description,omitempty"`
	IpAddress       string                          `json:"ip_address,omitempty"`
	Active          bool                            `json:"active,omitempty"`
	Paused          bool                            `json:"paused,omitempty"`
	IsShared        bool                            `json:"is_shared,omitempty"`
	RunnerType      string                          `json:"runner_type,omitempty"`
	Name            string                          `json:"name,omitempty"`
	Online          bool                            `json:"online,omitempty"`
	Status          string                          `json:"status,omitempty"`
	TagList         string                          `json:"tag_list,omitempty"`
	RunUntagged     string                          `json:"run_untagged,omitempty"`
	Locked          string                          `json:"locked,omitempty"`
	MaximumTimeout  string                          `json:"maximum_timeout,omitempty"`
	AccessLevel     string                          `json:"access_level,omitempty"`
	Version         string                          `json:"version,omitempty"`
	Revision        string                          `json:"revision,omitempty"`
	Platform        string                          `json:"platform,omitempty"`
	Architecture    string                          `json:"architecture,omitempty"`
	ContactedAt     string                          `json:"contacted_at,omitempty"`
	MaintenanceNote string                          `json:"maintenance_note,omitempty"`
	Projects        *ApiEntitiesBasicProjectDetails `json:"projects,omitempty"`
	Groups          *ApiEntitiesBasicGroupDetails   `json:"groups,omitempty"`
}
