# PostApiV4GroupsIdInvitations

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessLevel** | **int32** | A valid access level (defaults: &#x60;30&#x60;, developer access level) | [default to null]
**Email** | **[]string** | The email address to invite, or multiple emails separated by comma | [optional] [default to null]
**UserId** | **[]string** | The user ID of the new member or multiple IDs separated by commas. | [optional] [default to null]
**ExpiresAt** | [**time.Time**](time.Time.md) | Date string in the format YEAR-MONTH-DAY | [optional] [default to null]
**InviteSource** | **string** | Source that triggered the member creation process | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


