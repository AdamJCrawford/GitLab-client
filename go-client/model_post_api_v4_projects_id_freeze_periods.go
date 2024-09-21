/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Create a freeze period
type PostApiV4ProjectsIdFreezePeriods struct {
	// Start of the freeze period in cron format.
	FreezeStart string `json:"freeze_start"`
	// End of the freeze period in cron format
	FreezeEnd string `json:"freeze_end"`
	// The time zone for the cron fields, defaults to UTC if not provided
	CronTimezone string `json:"cron_timezone,omitempty"`
}
