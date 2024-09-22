# PostApiV4GeoStatusData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeoNodeId** | **int32** | Geo Node ID to look up its status | [default to null]
**DbReplicationLagSeconds** | **int32** | DB replication lag in seconds | [optional] [default to null]
**LastEventId** | **int32** | Last event ID | [optional] [default to null]
**LastEventDate** | [**time.Time**](time.Time.md) | Last event date | [optional] [default to null]
**CursorLastEventId** | **int32** | Cursor last event ID | [optional] [default to null]
**CursorLastEventDate** | [**time.Time**](time.Time.md) | Cursor last event date | [optional] [default to null]
**LastSuccessfulStatusCheckAt** | [**time.Time**](time.Time.md) | Last successful status check date | [optional] [default to null]
**StatusMessage** | **string** | Status message | [optional] [default to null]
**ReplicationSlotsCount** | **int32** | Replication slots count | [optional] [default to null]
**ReplicationSlotsUsedCount** | **int32** | Replication slots used count | [optional] [default to null]
**ReplicationSlotsMaxRetainedWalBytes** | **int32** | Maximum number of bytes retained in the WAL on the primary | [optional] [default to null]
**Version** | **string** | Gitlab version | [optional] [default to null]
**Revision** | **string** | Gitlab revision | [optional] [default to null]
**Status** | [***PostApiV4GeoStatusDataStatus**](postApiV4GeoStatus_data_status.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


