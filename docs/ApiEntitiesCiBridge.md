# ApiEntitiesCiBridge

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [optional] [default to null]
**Status** | **string** |  | [optional] [default to null]
**Stage** | **string** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**Ref** | **string** |  | [optional] [default to null]
**Tag** | **bool** |  | [optional] [default to null]
**Coverage** | **float32** |  | [optional] [default to null]
**AllowFailure** | **bool** |  | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**StartedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**FinishedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**ErasedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Duration** | **float32** | Time spent running | [optional] [default to null]
**QueuedDuration** | **float32** | Time spent enqueued | [optional] [default to null]
**User** | [***ApiEntitiesUser**](API_Entities_User.md) |  | [optional] [default to null]
**Commit** | [***ApiEntitiesCommit**](API_Entities_Commit.md) |  | [optional] [default to null]
**Pipeline** | [***ApiEntitiesCiPipelineBasic**](API_Entities_Ci_PipelineBasic.md) |  | [optional] [default to null]
**FailureReason** | **string** |  | [optional] [default to null]
**WebUrl** | **string** |  | [optional] [default to null]
**Project** | [***ApiEntitiesCiJobProject**](API_Entities_Ci_Job_project.md) |  | [optional] [default to null]
**DownstreamPipeline** | [***ApiEntitiesCiPipelineBasic**](API_Entities_Ci_PipelineBasic.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


