/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Protect a single branch
type PostApiV4ProjectsIdProtectedBranches struct {
	// The name of the protected branch
	Name string `json:"name"`
	// Access levels allowed to push (defaults: `40`, maintainer access level)
	PushAccessLevel int32 `json:"push_access_level,omitempty"`
	// Access levels allowed to merge (defaults: `40`, maintainer access level)
	MergeAccessLevel int32 `json:"merge_access_level,omitempty"`
	// Allow force push for all users with push access.
	AllowForcePush bool `json:"allow_force_push,omitempty"`
	// Access levels allowed to unprotect (defaults: `40`, maintainer access level)
	UnprotectAccessLevel int32 `json:"unprotect_access_level,omitempty"`
	// An array of users/groups allowed to push
	AllowedToPush []PostApiV4ProjectsIdProtectedBranchesAllowedToPush `json:"allowed_to_push,omitempty"`
	// An array of users/groups allowed to merge
	AllowedToMerge []PostApiV4ProjectsIdProtectedBranchesAllowedToPush `json:"allowed_to_merge,omitempty"`
	// An array of users/groups allowed to unprotect
	AllowedToUnprotect []PostApiV4ProjectsIdProtectedBranchesAllowedToUnprotect `json:"allowed_to_unprotect,omitempty"`
	// Prevent pushes to this branch if it matches an item in CODEOWNERS
	CodeOwnerApprovalRequired bool `json:"code_owner_approval_required,omitempty"`
}
