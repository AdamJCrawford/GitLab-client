# PutApiV4ProjectsIdMergeRequestsMergeRequestIidMerge

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MergeCommitMessage** | **string** | Custom merge commit message. | [optional] [default to null]
**SquashCommitMessage** | **string** | Custom squash commit message. | [optional] [default to null]
**ShouldRemoveSourceBranch** | **bool** | If &#x60;true&#x60;, removes the source branch. | [optional] [default to null]
**MergeWhenPipelineSucceeds** | **bool** | If &#x60;true&#x60;, the merge request is merged when the pipeline succeeds. | [optional] [default to null]
**Sha** | **string** | If present, then this SHA must match the HEAD of the source branch, otherwise the merge fails. | [optional] [default to null]
**Squash** | **bool** | If &#x60;true&#x60;, the commits are squashed into a single commit on merge. | [optional] [default to null]
**SkipMergeTrain** | **bool** | If &#x60;true&#x60; skips train restart when merging immediately in a merge train configured project. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


