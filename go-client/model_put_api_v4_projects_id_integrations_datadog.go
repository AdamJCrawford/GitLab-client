/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Create/Edit Datadog integration
type PutApiV4ProjectsIdIntegrationsDatadog struct {
	// API key used for authentication with Datadog
	ApiKey string `json:"api_key"`
	// The Datadog site to send data to. To send data to the EU site, use datadoghq.eu
	DatadogSite string `json:"datadog_site,omitempty"`
	// (Advanced) The full URL for your Datadog site
	ApiUrl string `json:"api_url,omitempty"`
	// When enabled, job logs are collected by Datadog and displayed along with pipeline execution traces.
	ArchiveTraceEvents bool `json:"archive_trace_events,omitempty"`
	// Tag all data from this GitLab instance in Datadog. Useful when managing several self-managed deployments
	DatadogService string `json:"datadog_service,omitempty"`
	// For self-managed deployments, set the env tag for all the data sent to Datadog
	DatadogEnv string `json:"datadog_env,omitempty"`
	// Custom tags in Datadog. Specify one tag per line in the format: \"key:value\\nkey2:value2\"
	DatadogTags string `json:"datadog_tags,omitempty"`
	// Trigger event when a pipeline status changes.
	PipelineEvents bool `json:"pipeline_events,omitempty"`
	// Trigger event when a build is created.
	BuildEvents bool `json:"build_events,omitempty"`
}