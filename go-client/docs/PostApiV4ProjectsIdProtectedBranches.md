# PostApiV4ProjectsIdProtectedBranches

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the protected branch | [default to null]
**PushAccessLevel** | **int32** | Access levels allowed to push (defaults: &#x60;40&#x60;, maintainer access level) | [optional] [default to null]
**MergeAccessLevel** | **int32** | Access levels allowed to merge (defaults: &#x60;40&#x60;, maintainer access level) | [optional] [default to null]
**AllowForcePush** | **bool** | Allow force push for all users with push access. | [optional] [default to null]
**UnprotectAccessLevel** | **int32** | Access levels allowed to unprotect (defaults: &#x60;40&#x60;, maintainer access level) | [optional] [default to null]
**AllowedToPush** | [**[]PostApiV4ProjectsIdProtectedBranchesAllowedToPush**](postApiV4ProjectsIdProtectedBranches_allowed_to_push.md) | An array of users/groups allowed to push | [optional] [default to null]
**AllowedToMerge** | [**[]PostApiV4ProjectsIdProtectedBranchesAllowedToPush**](postApiV4ProjectsIdProtectedBranches_allowed_to_push.md) | An array of users/groups allowed to merge | [optional] [default to null]
**AllowedToUnprotect** | [**[]PostApiV4ProjectsIdProtectedBranchesAllowedToUnprotect**](postApiV4ProjectsIdProtectedBranches_allowed_to_unprotect.md) | An array of users/groups allowed to unprotect | [optional] [default to null]
**CodeOwnerApprovalRequired** | **bool** | Prevent pushes to this branch if it matches an item in CODEOWNERS | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


