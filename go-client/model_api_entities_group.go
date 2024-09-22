/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// API_Entities_Group model
type ApiEntitiesGroup struct {
	Id                              string                      `json:"id,omitempty"`
	WebUrl                          string                      `json:"web_url,omitempty"`
	Name                            string                      `json:"name,omitempty"`
	Path                            string                      `json:"path,omitempty"`
	Description                     string                      `json:"description,omitempty"`
	Visibility                      string                      `json:"visibility,omitempty"`
	ShareWithGroupLock              string                      `json:"share_with_group_lock,omitempty"`
	RequireTwoFactorAuthentication  string                      `json:"require_two_factor_authentication,omitempty"`
	TwoFactorGracePeriod            string                      `json:"two_factor_grace_period,omitempty"`
	ProjectCreationLevel            string                      `json:"project_creation_level,omitempty"`
	AutoDevopsEnabled               string                      `json:"auto_devops_enabled,omitempty"`
	SubgroupCreationLevel           string                      `json:"subgroup_creation_level,omitempty"`
	EmailsDisabled                  bool                        `json:"emails_disabled,omitempty"`
	EmailsEnabled                   bool                        `json:"emails_enabled,omitempty"`
	MentionsDisabled                string                      `json:"mentions_disabled,omitempty"`
	LfsEnabled                      string                      `json:"lfs_enabled,omitempty"`
	MathRenderingLimitsEnabled      bool                        `json:"math_rendering_limits_enabled,omitempty"`
	LockMathRenderingLimitsEnabled  bool                        `json:"lock_math_rendering_limits_enabled,omitempty"`
	DefaultBranch                   string                      `json:"default_branch,omitempty"`
	DefaultBranchProtection         string                      `json:"default_branch_protection,omitempty"`
	DefaultBranchProtectionDefaults string                      `json:"default_branch_protection_defaults,omitempty"`
	AvatarUrl                       string                      `json:"avatar_url,omitempty"`
	RequestAccessEnabled            string                      `json:"request_access_enabled,omitempty"`
	FullName                        string                      `json:"full_name,omitempty"`
	FullPath                        string                      `json:"full_path,omitempty"`
	CreatedAt                       string                      `json:"created_at,omitempty"`
	ParentId                        string                      `json:"parent_id,omitempty"`
	OrganizationId                  string                      `json:"organization_id,omitempty"`
	SharedRunnersSetting            string                      `json:"shared_runners_setting,omitempty"`
	CustomAttributes                *ApiEntitiesCustomAttribute `json:"custom_attributes,omitempty"`
	Statistics                      *ApiEntitiesGroupStatistics `json:"statistics,omitempty"`
	LdapCn                          string                      `json:"ldap_cn,omitempty"`
	LdapAccess                      string                      `json:"ldap_access,omitempty"`
	LdapGroupLinks                  *EeApiEntitiesLdapGroupLink `json:"ldap_group_links,omitempty"`
	SamlGroupLinks                  *EeApiEntitiesSamlGroupLink `json:"saml_group_links,omitempty"`
	FileTemplateProjectId           string                      `json:"file_template_project_id,omitempty"`
	MarkedForDeletionOn             string                      `json:"marked_for_deletion_on,omitempty"`
	WikiAccessLevel                 string                      `json:"wiki_access_level,omitempty"`
	RepositoryStorage               string                      `json:"repository_storage,omitempty"`
	DuoFeaturesEnabled              string                      `json:"duo_features_enabled,omitempty"`
	LockDuoFeaturesEnabled          string                      `json:"lock_duo_features_enabled,omitempty"`
}