# PostApiV4IntegrationsSlackEvents

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | (Deprecated by Slack) The request token, unused by GitLab | [optional] [default to null]
**TeamId** | **string** | The Slack workspace ID of where the event occurred | [optional] [default to null]
**ApiAppId** | **string** | The Slack app ID | [optional] [default to null]
**Event** | **interface{}** | The event object with variable properties | [optional] [default to null]
**Type_** | **string** | The kind of event this is, usually &#x60;event_callback&#x60; | [optional] [default to null]
**EventId** | **string** | A unique identifier for this specific event | [optional] [default to null]
**EventTime** | **int32** | The epoch timestamp in seconds when this event was dispatched | [optional] [default to null]
**AuthedUsers** | **[]string** | (Deprecated by Slack) An array of Slack user IDs | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


