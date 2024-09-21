# PostApiV4ProjectsIdMergeRequests

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | The title of the merge request. | [default to null]
**SourceBranch** | **string** | The source branch. | [default to null]
**TargetBranch** | **string** | The target branch. | [default to null]
**TargetProjectId** | **int32** | The target project of the merge request defaults to the :id of the project. | [optional] [default to null]
**AssigneeId** | **int32** | Assignee user ID. | [optional] [default to null]
**AssigneeIds** | **[]int32** | The IDs of the users to assign the merge request to, as a comma-separated list. Set to 0 or provide an empty value to unassign all assignees. | [optional] [default to null]
**ReviewerIds** | **[]int32** | The IDs of the users to review the merge request, as a comma-separated list. Set to 0 or provide an empty value to unassign all reviewers. | [optional] [default to null]
**Description** | **string** | Description of the merge request. Limited to 1,048,576 characters. | [optional] [default to null]
**Labels** | **[]string** | Comma-separated label names for a merge request. Set to an empty string to unassign all labels. | [optional] [default to null]
**AddLabels** | **[]string** | Comma-separated label names to add to a merge request. | [optional] [default to null]
**RemoveLabels** | **[]string** | Comma-separated label names to remove from a merge request. | [optional] [default to null]
**MilestoneId** | **int32** | The global ID of a milestone to assign the merge reques to. | [optional] [default to null]
**RemoveSourceBranch** | **bool** | Flag indicating if a merge request should remove the source branch when merging. | [optional] [default to null]
**AllowCollaboration** | **bool** | Allow commits from members who can merge to the target branch. | [optional] [default to null]
**AllowMaintainerToPush** | **bool** | [deprecated] See allow_collaboration | [optional] [default to null]
**Squash** | **bool** | Squash commits into a single commit when merging. | [optional] [default to null]
**ApprovalsBeforeMerge** | **int32** | Number of approvals required before this can be merged | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


