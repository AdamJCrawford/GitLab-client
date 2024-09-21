# EeApiEntitiesApprovalState

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [optional] [default to null]
**Iid** | **int32** |  | [optional] [default to null]
**ProjectId** | **int32** |  | [optional] [default to null]
**Title** | **string** |  | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**State** | **string** |  | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**MergeStatus** | **string** |  | [optional] [default to null]
**Approved** | **bool** |  | [optional] [default to null]
**ApprovalsRequired** | **int32** |  | [optional] [default to null]
**ApprovalsLeft** | **int32** |  | [optional] [default to null]
**RequirePasswordToApprove** | **bool** |  | [optional] [default to null]
**ApprovedBy** | [**[]ApiEntitiesApprovals**](API_Entities_Approvals.md) |  | [optional] [default to null]
**SuggestedApprovers** | [**[]ApiEntitiesUserBasic**](API_Entities_UserBasic.md) |  | [optional] [default to null]
**Approvers** | **string** |  | [optional] [default to null]
**ApproverGroups** | **string** |  | [optional] [default to null]
**UserHasApproved** | **bool** |  | [optional] [default to null]
**UserCanApprove** | **bool** |  | [optional] [default to null]
**ApprovalRulesLeft** | [**[]EeApiEntitiesApprovalRuleShort**](EE_API_Entities_ApprovalRuleShort.md) |  | [optional] [default to null]
**HasApprovalRules** | **bool** |  | [optional] [default to null]
**MergeRequestApproversAvailable** | **bool** |  | [optional] [default to null]
**MultipleApprovalRulesAvailable** | **bool** |  | [optional] [default to null]
**InvalidApproversRules** | [**[]EeApiEntitiesApprovalRuleShort**](EE_API_Entities_ApprovalRuleShort.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


