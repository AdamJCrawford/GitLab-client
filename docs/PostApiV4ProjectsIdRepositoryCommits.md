# PostApiV4ProjectsIdRepositoryCommits

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | **string** | Name of the branch to commit into. To create a new branch, also provide either &#x60;start_branch&#x60; or &#x60;start_sha&#x60;, and optionally &#x60;start_project&#x60;. | [default to null]
**CommitMessage** | **string** | Commit message | [default to null]
**Actions** | [**[]PostApiV4ProjectsIdRepositoryCommitsActions**](postApiV4ProjectsIdRepositoryCommits_actions.md) | Actions to perform in commit | [default to null]
**StartBranch** | **string** | Name of the branch to start the new branch from | [optional] [default to null]
**StartSha** | **string** | SHA of the commit to start the new branch from | [optional] [default to null]
**StartProject** | **int32** | The ID or path of the project to start the new branch from | [optional] [default to null]
**AuthorEmail** | **string** | Author email for commit | [optional] [default to null]
**AuthorName** | **string** | Author name for commit | [optional] [default to null]
**Stats** | **bool** | Include commit stats | [optional] [default to null]
**Force** | **bool** | When &#x60;true&#x60; overwrites the target branch with a new commit based on the &#x60;start_branch&#x60; or &#x60;start_sha&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


