/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ApiEntitiesProjectDailyStatisticsFetches struct {
	Total int32                            `json:"total,omitempty"`
	Days  []ApiEntitiesProjectDailyFetches `json:"days,omitempty"`
}