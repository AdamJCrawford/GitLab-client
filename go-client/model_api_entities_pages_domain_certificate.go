/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type ApiEntitiesPagesDomainCertificate struct {
	Subject         string `json:"subject,omitempty"`
	Expired         string `json:"expired,omitempty"`
	Certificate     string `json:"certificate,omitempty"`
	CertificateText string `json:"certificate_text,omitempty"`
}
