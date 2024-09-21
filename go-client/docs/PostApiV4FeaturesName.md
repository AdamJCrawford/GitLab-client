# PostApiV4FeaturesName

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** | &#x60;true&#x60; or &#x60;false&#x60; to enable/disable, or an integer for percentage of time | [default to null]
**Key** | **string** | &#x60;percentage_of_actors&#x60; or &#x60;percentage_of_time&#x60; (default) | [optional] [default to null]
**FeatureGroup** | **string** | A Feature group name | [optional] [default to null]
**User** | **string** | A GitLab username or comma-separated multiple usernames | [optional] [default to null]
**Group** | **string** | A GitLab group&#39;s path, for example &#x60;gitlab-org&#x60;, or comma-separated multiple group paths | [optional] [default to null]
**Namespace** | **string** | A GitLab group or user namespace&#39;s path, for example &#x60;john-doe&#x60;, or comma-separated multiple namespace paths. Introduced in GitLab 15.0. | [optional] [default to null]
**Project** | **string** | A projects path, for example &#x60;gitlab-org/gitlab-foss&#x60;, or comma-separated multiple project paths | [optional] [default to null]
**Repository** | **string** | A repository path, for example &#x60;gitlab-org/gitlab-test.git&#x60;, &#x60;gitlab-org/gitlab-test.wiki.git&#x60;, &#x60;snippets/21.git&#x60;, to name a few. Use comma to separate multiple repository paths | [optional] [default to null]
**Force** | **bool** | Skip feature flag validation checks, such as a YAML definition | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


