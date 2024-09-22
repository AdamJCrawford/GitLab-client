# PutApiV4ProjectsIdIntegrationsJenkins

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JenkinsUrl** | **string** | Jenkins root URL like https://jenkins.example.com | [default to null]
**EnableSslVerification** | **bool** | Enable SSL verification | [optional] [default to null]
**ProjectName** | **string** | The URL-friendly project name. Example: my_project_name | [default to null]
**Username** | **string** | A user with access to the Jenkins server, if applicable | [optional] [default to null]
**Password** | **string** | The password of the user | [optional] [default to null]
**PushEvents** | **bool** | Trigger event for pushes to the repository. | [optional] [default to null]
**MergeRequestsEvents** | **bool** | Trigger event when a merge request is created, updated, or merged. | [optional] [default to null]
**TagPushEvents** | **bool** | Trigger event for new tags pushed to the repository. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


