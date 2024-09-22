# PostApiV4ProjectsIdReleases

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagName** | **string** | The tag where the release is created from | [default to null]
**TagMessage** | **string** | Message to use if creating a new annotated tag | [optional] [default to null]
**Name** | **string** | The release name | [optional] [default to null]
**Description** | **string** | The description of the release. You can use Markdown | [optional] [default to null]
**Ref** | **string** | If a tag specified in &#x60;tag_name&#x60; doesn&#39;t exist, the release is created from &#x60;ref&#x60; and tagged with &#x60;tag_name&#x60;. It can be a commit SHA, another tag name, or a branch name. | [optional] [default to null]
**Assets** | [***PostApiV4ProjectsIdReleasesAssets**](postApiV4ProjectsIdReleases_assets.md) |  | [optional] [default to null]
**Milestones** | **[]string** | The title of each milestone the release is associated with. GitLab Premium customers can specify group milestones. Cannot be combined with &#x60;milestone_ids&#x60; parameter. | [optional] [default to null]
**MilestoneIds** | **string** | The ID of each milestone the release is associated with. GitLab Premium customers can specify group milestones. Cannot be combined with &#x60;milestones&#x60; parameter. | [optional] [default to null]
**ReleasedAt** | [**time.Time**](time.Time.md) | Date and time for the release. Defaults to the current time. Expected in ISO 8601 format (&#x60;2019-03-15T08:00:00Z&#x60;). Only provide this field if creating an upcoming or historical release. | [optional] [default to null]
**LegacyCatalogPublish** | **bool** | If true, the release will be published to the CI catalog. This parameter is for internal use only and will be removed in a future release. If the feature flag ci_release_cli_catalog_publish_option is disabled, this parameter will be ignored and the release will published to the CI catalog as it was before this parameter was introduced. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


