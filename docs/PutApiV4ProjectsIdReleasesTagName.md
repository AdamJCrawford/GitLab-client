# PutApiV4ProjectsIdReleasesTagName

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The release name | [optional] [default to null]
**Description** | **string** | The description of the release. You can use Markdown | [optional] [default to null]
**ReleasedAt** | [**time.Time**](time.Time.md) | The date when the release is/was ready. Expected in ISO 8601 format (&#x60;2019-03-15T08:00:00Z&#x60;) | [optional] [default to null]
**Milestones** | **[]string** | The title of each milestone to associate with the release. GitLab Premium customers can specify group milestones. Cannot be combined with &#x60;milestone_ids&#x60; parameter. To remove all milestones from the release, specify &#x60;[]&#x60; | [optional] [default to null]
**MilestoneIds** | **string** | The ID of each milestone the release is associated with. GitLab Premium customers can specify group milestones. Cannot be combined with &#x60;milestones&#x60; parameter. To remove all milestones from the release, specify &#x60;[]&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


