/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// API_Entities_PagesDomain model
type ApiEntitiesPagesDomain struct {
	Domain           string                             `json:"domain,omitempty"`
	Url              string                             `json:"url,omitempty"`
	Verified         string                             `json:"verified,omitempty"`
	VerificationCode string                             `json:"verification_code,omitempty"`
	EnabledUntil     string                             `json:"enabled_until,omitempty"`
	AutoSslEnabled   string                             `json:"auto_ssl_enabled,omitempty"`
	Certificate      *ApiEntitiesPagesDomainCertificate `json:"certificate,omitempty"`
}
