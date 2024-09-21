/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Create/Edit Prometheus integration
type PutApiV4ProjectsIdIntegrationsPrometheus struct {
	// When enabled, the default settings will be overridden with your custom configuration
	ManualConfiguration bool `json:"manual_configuration,omitempty"`
	// Prometheus API Base URL, like http://prometheus.example.com/
	ApiUrl string `json:"api_url"`
	// Client ID of the IAP-secured resource (looks like IAP_CLIENT_ID.apps.googleusercontent.com)
	GoogleIapAudienceClientId string `json:"google_iap_audience_client_id"`
	// Contents of the credentials.json file of your service account, like: { \"type\": \"service_account\", \"project_id\": ... }
	GoogleIapServiceAccountJson string `json:"google_iap_service_account_json"`
}
