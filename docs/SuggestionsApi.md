# \SuggestionsApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutApiV4SuggestionsBatchApply**](SuggestionsApi.md#PutApiV4SuggestionsBatchApply) | **Put** /api/v4/suggestions/batch_apply | 
[**PutApiV4SuggestionsIdApply**](SuggestionsApi.md#PutApiV4SuggestionsIdApply) | **Put** /api/v4/suggestions/{id}/apply | 


# **PutApiV4SuggestionsBatchApply**
> ApiEntitiesSuggestion PutApiV4SuggestionsBatchApply(ctx, putApiV4SuggestionsBatchApply)


Apply multiple suggestion patches in the Merge Request where they were created

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **putApiV4SuggestionsBatchApply** | [**PutApiV4SuggestionsBatchApply**](PutApiV4SuggestionsBatchApply.md)|  | 

### Return type

[**ApiEntitiesSuggestion**](API_Entities_Suggestion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4SuggestionsIdApply**
> ApiEntitiesSuggestion PutApiV4SuggestionsIdApply(ctx, id, putApiV4SuggestionsIdApply)


Apply suggestion patch in the Merge Request it was created

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the suggestion | 
  **putApiV4SuggestionsIdApply** | [**PutApiV4SuggestionsIdApply**](PutApiV4SuggestionsIdApply.md)|  | 

### Return type

[**ApiEntitiesSuggestion**](API_Entities_Suggestion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

