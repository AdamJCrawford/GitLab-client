/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

import (
	"time"
)

// API_Entities_UserPublic model
type ApiEntitiesUserPublic struct {
	Id                             int32                        `json:"id,omitempty"`
	Username                       string                       `json:"username,omitempty"`
	Name                           string                       `json:"name,omitempty"`
	State                          string                       `json:"state,omitempty"`
	Locked                         bool                         `json:"locked,omitempty"`
	AvatarUrl                      string                       `json:"avatar_url,omitempty"`
	AvatarPath                     string                       `json:"avatar_path,omitempty"`
	CustomAttributes               []ApiEntitiesCustomAttribute `json:"custom_attributes,omitempty"`
	WebUrl                         string                       `json:"web_url,omitempty"`
	CreatedAt                      string                       `json:"created_at,omitempty"`
	Bio                            string                       `json:"bio,omitempty"`
	Location                       string                       `json:"location,omitempty"`
	PublicEmail                    string                       `json:"public_email,omitempty"`
	Skype                          string                       `json:"skype,omitempty"`
	Linkedin                       string                       `json:"linkedin,omitempty"`
	Twitter                        string                       `json:"twitter,omitempty"`
	Discord                        string                       `json:"discord,omitempty"`
	WebsiteUrl                     string                       `json:"website_url,omitempty"`
	Organization                   string                       `json:"organization,omitempty"`
	JobTitle                       string                       `json:"job_title,omitempty"`
	Pronouns                       string                       `json:"pronouns,omitempty"`
	Bot                            string                       `json:"bot,omitempty"`
	WorkInformation                string                       `json:"work_information,omitempty"`
	Followers                      string                       `json:"followers,omitempty"`
	Following                      string                       `json:"following,omitempty"`
	IsFollowed                     string                       `json:"is_followed,omitempty"`
	LocalTime                      string                       `json:"local_time,omitempty"`
	LastSignInAt                   time.Time                    `json:"last_sign_in_at,omitempty"`
	ConfirmedAt                    time.Time                    `json:"confirmed_at,omitempty"`
	LastActivityOn                 time.Time                    `json:"last_activity_on,omitempty"`
	Email                          string                       `json:"email,omitempty"`
	ThemeId                        int32                        `json:"theme_id,omitempty"`
	ColorSchemeId                  int32                        `json:"color_scheme_id,omitempty"`
	ProjectsLimit                  int32                        `json:"projects_limit,omitempty"`
	CurrentSignInAt                time.Time                    `json:"current_sign_in_at,omitempty"`
	Identities                     *ApiEntitiesIdentity         `json:"identities,omitempty"`
	CanCreateGroup                 bool                         `json:"can_create_group,omitempty"`
	CanCreateProject               bool                         `json:"can_create_project,omitempty"`
	TwoFactorEnabled               bool                         `json:"two_factor_enabled,omitempty"`
	External                       string                       `json:"external,omitempty"`
	PrivateProfile                 bool                         `json:"private_profile,omitempty"`
	CommitEmail                    string                       `json:"commit_email,omitempty"`
	SharedRunnersMinutesLimit      string                       `json:"shared_runners_minutes_limit,omitempty"`
	ExtraSharedRunnersMinutesLimit string                       `json:"extra_shared_runners_minutes_limit,omitempty"`
	ScimIdentities                 *ApiEntitiesScimIdentity     `json:"scim_identities,omitempty"`
}