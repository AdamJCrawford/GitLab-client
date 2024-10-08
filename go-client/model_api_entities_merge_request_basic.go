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

// API_Entities_MergeRequestBasic model
type ApiEntitiesMergeRequestBasic struct {
	Id                          int32                          `json:"id,omitempty"`
	Iid                         int32                          `json:"iid,omitempty"`
	ProjectId                   int32                          `json:"project_id,omitempty"`
	Title                       string                         `json:"title,omitempty"`
	Description                 string                         `json:"description,omitempty"`
	State                       string                         `json:"state,omitempty"`
	CreatedAt                   time.Time                      `json:"created_at,omitempty"`
	UpdatedAt                   time.Time                      `json:"updated_at,omitempty"`
	MergedBy                    *ApiEntitiesUserBasic          `json:"merged_by,omitempty"`
	MergeUser                   *ApiEntitiesUserBasic          `json:"merge_user,omitempty"`
	MergedAt                    string                         `json:"merged_at,omitempty"`
	ClosedBy                    *ApiEntitiesUserBasic          `json:"closed_by,omitempty"`
	ClosedAt                    string                         `json:"closed_at,omitempty"`
	TitleHtml                   string                         `json:"title_html,omitempty"`
	DescriptionHtml             string                         `json:"description_html,omitempty"`
	TargetBranch                string                         `json:"target_branch,omitempty"`
	SourceBranch                string                         `json:"source_branch,omitempty"`
	UserNotesCount              string                         `json:"user_notes_count,omitempty"`
	Upvotes                     string                         `json:"upvotes,omitempty"`
	Downvotes                   string                         `json:"downvotes,omitempty"`
	Author                      *ApiEntitiesUserBasic          `json:"author,omitempty"`
	Assignees                   *ApiEntitiesUserBasic          `json:"assignees,omitempty"`
	Assignee                    *ApiEntitiesUserBasic          `json:"assignee,omitempty"`
	Reviewers                   *ApiEntitiesUserBasic          `json:"reviewers,omitempty"`
	SourceProjectId             string                         `json:"source_project_id,omitempty"`
	TargetProjectId             string                         `json:"target_project_id,omitempty"`
	Labels                      string                         `json:"labels,omitempty"`
	Draft                       string                         `json:"draft,omitempty"`
	Imported                    string                         `json:"imported,omitempty"`
	ImportedFrom                string                         `json:"imported_from,omitempty"`
	WorkInProgress              string                         `json:"work_in_progress,omitempty"`
	Milestone                   *ApiEntitiesMilestone          `json:"milestone,omitempty"`
	MergeWhenPipelineSucceeds   string                         `json:"merge_when_pipeline_succeeds,omitempty"`
	MergeStatus                 string                         `json:"merge_status,omitempty"`
	DetailedMergeStatus         string                         `json:"detailed_merge_status,omitempty"`
	Sha                         string                         `json:"sha,omitempty"`
	MergeCommitSha              string                         `json:"merge_commit_sha,omitempty"`
	SquashCommitSha             string                         `json:"squash_commit_sha,omitempty"`
	DiscussionLocked            string                         `json:"discussion_locked,omitempty"`
	ShouldRemoveSourceBranch    string                         `json:"should_remove_source_branch,omitempty"`
	ForceRemoveSourceBranch     string                         `json:"force_remove_source_branch,omitempty"`
	PreparedAt                  string                         `json:"prepared_at,omitempty"`
	AllowCollaboration          string                         `json:"allow_collaboration,omitempty"`
	AllowMaintainerToPush       string                         `json:"allow_maintainer_to_push,omitempty"`
	Reference                   string                         `json:"reference,omitempty"`
	References                  *ApiEntitiesIssuableReferences `json:"references,omitempty"`
	WebUrl                      string                         `json:"web_url,omitempty"`
	TimeStats                   *ApiEntitiesIssuableTimeStats  `json:"time_stats,omitempty"`
	Squash                      string                         `json:"squash,omitempty"`
	SquashOnMerge               string                         `json:"squash_on_merge,omitempty"`
	TaskCompletionStatus        string                         `json:"task_completion_status,omitempty"`
	HasConflicts                string                         `json:"has_conflicts,omitempty"`
	BlockingDiscussionsResolved string                         `json:"blocking_discussions_resolved,omitempty"`
	ApprovalsBeforeMerge        string                         `json:"approvals_before_merge,omitempty"`
}
