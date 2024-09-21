/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Update a job
type PutApiV4JobsId struct {
	// Job token
	Token string `json:"token"`
	// Job's status: success, failed
	State string `json:"state,omitempty"`
	// Job's trace CRC32 checksum
	Checksum string `json:"checksum,omitempty"`
	// Job's failure_reason
	FailureReason string                `json:"failure_reason,omitempty"`
	Output        *PutApiV4JobsIdOutput `json:"output,omitempty"`
	// Job's exit code
	ExitCode int32 `json:"exit_code,omitempty"`
}
