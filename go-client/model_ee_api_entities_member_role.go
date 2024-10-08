/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type EeApiEntitiesMemberRole struct {
	Id                         int32 `json:"id,omitempty"`
	GroupId                    int32 `json:"group_id,omitempty"`
	BaseAccessLevel            int32 `json:"base_access_level,omitempty"`
	AdminCicdVariables         bool  `json:"admin_cicd_variables,omitempty"`
	AdminComplianceFramework   bool  `json:"admin_compliance_framework,omitempty"`
	AdminGroupMember           bool  `json:"admin_group_member,omitempty"`
	AdminMergeRequest          bool  `json:"admin_merge_request,omitempty"`
	AdminPushRules             bool  `json:"admin_push_rules,omitempty"`
	AdminTerraformState        bool  `json:"admin_terraform_state,omitempty"`
	AdminVulnerability         bool  `json:"admin_vulnerability,omitempty"`
	AdminWebHook               bool  `json:"admin_web_hook,omitempty"`
	ArchiveProject             bool  `json:"archive_project,omitempty"`
	ManageDeployTokens         bool  `json:"manage_deploy_tokens,omitempty"`
	ManageGroupAccessTokens    bool  `json:"manage_group_access_tokens,omitempty"`
	ManageMergeRequestSettings bool  `json:"manage_merge_request_settings,omitempty"`
	ManageProjectAccessTokens  bool  `json:"manage_project_access_tokens,omitempty"`
	ManageSecurityPolicyLink   bool  `json:"manage_security_policy_link,omitempty"`
	ReadCode                   bool  `json:"read_code,omitempty"`
	ReadDependency             bool  `json:"read_dependency,omitempty"`
	ReadVulnerability          bool  `json:"read_vulnerability,omitempty"`
	RemoveGroup                bool  `json:"remove_group,omitempty"`
	RemoveProject              bool  `json:"remove_project,omitempty"`
}
