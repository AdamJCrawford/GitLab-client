/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

// API_Entities_BulkImports_ExportStatus model
type ApiEntitiesBulkImportsExportStatus struct {
	Relation          string                                   `json:"relation,omitempty"`
	Status            string                                   `json:"status,omitempty"`
	Error_            string                                   `json:"error,omitempty"`
	UpdatedAt         time.Time                                `json:"updated_at,omitempty"`
	Batched           bool                                     `json:"batched,omitempty"`
	BatchesCount      int32                                    `json:"batches_count,omitempty"`
	TotalObjectsCount int32                                    `json:"total_objects_count,omitempty"`
	Batches           *ApiEntitiesBulkImportsExportBatchStatus `json:"batches,omitempty"`
}
