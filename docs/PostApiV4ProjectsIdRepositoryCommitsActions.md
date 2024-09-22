# PostApiV4ProjectsIdRepositoryCommitsActions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | The action to perform, &#x60;create&#x60;, &#x60;delete&#x60;, &#x60;move&#x60;, &#x60;update&#x60;, &#x60;chmod&#x60; | [default to null]
**FilePath** | **string** | Full path to the file. | [default to null]
**PreviousPath** | **string** | Original full path to the file being moved. | [default to null]
**Content** | **string** | File content | [default to null]
**Encoding** | **string** | &#x60;text&#x60; or &#x60;base64&#x60; | [optional] [default to null]
**LastCommitId** | **string** | Last known file commit id | [optional] [default to null]
**ExecuteFilemode** | **bool** | When &#x60;true/false&#x60; enables/disables the execute flag on the file. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


