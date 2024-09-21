/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// API_Entities_ErrorTracking_ProjectSetting model
type ApiEntitiesErrorTrackingProjectSetting struct {
	Active            bool   `json:"active,omitempty"`
	ProjectName       string `json:"project_name,omitempty"`
	SentryExternalUrl string `json:"sentry_external_url,omitempty"`
	ApiUrl            string `json:"api_url,omitempty"`
	Integrated        bool   `json:"integrated,omitempty"`
}
