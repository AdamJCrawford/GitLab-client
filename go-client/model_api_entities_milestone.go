/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ApiEntitiesMilestone struct {
	Id          string `json:"id,omitempty"`
	Iid         string `json:"iid,omitempty"`
	ProjectId   string `json:"project_id,omitempty"`
	GroupId     string `json:"group_id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	State       string `json:"state,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
	DueDate     string `json:"due_date,omitempty"`
	StartDate   string `json:"start_date,omitempty"`
	Expired     string `json:"expired,omitempty"`
	WebUrl      string `json:"web_url,omitempty"`
}
