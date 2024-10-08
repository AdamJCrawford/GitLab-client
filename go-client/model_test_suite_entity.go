/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type TestSuiteEntity struct {
	Name         string           `json:"name,omitempty"`
	TotalTime    int32            `json:"total_time,omitempty"`
	TotalCount   int32            `json:"total_count,omitempty"`
	SuccessCount int32            `json:"success_count,omitempty"`
	FailedCount  int32            `json:"failed_count,omitempty"`
	SkippedCount int32            `json:"skipped_count,omitempty"`
	ErrorCount   int32            `json:"error_count,omitempty"`
	SuiteError   string           `json:"suite_error,omitempty"`
	TestCases    []TestCaseEntity `json:"test_cases,omitempty"`
}
