# PostApiV4ProjectsIdRepositoryChangelog

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | The version of the release, using the semantic versioning format | [default to null]
**From** | **string** | The first commit in the range of commits to use for the changelog | [optional] [default to null]
**To** | **string** | The last commit in the range of commits to use for the changelog | [optional] [default to null]
**Date** | [**time.Time**](time.Time.md) | The date and time of the release | [optional] [default to null]
**Trailer** | **string** | The Git trailer to use for determining if commits are to be included in the changelog | [optional] [default to null]
**Branch** | **string** | The branch to commit the changelog changes to | [optional] [default to null]
**ConfigFile** | **string** | The file path to the configuration file as stored in the project&#39;s Git repository. Defaults to &#39;.gitlab/changelog_config.yml&#39; | [optional] [default to null]
**File** | **string** | The file to commit the changelog changes to | [optional] [default to null]
**Message** | **string** | The commit message to use when committing the changelog | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


