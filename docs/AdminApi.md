# \AdminApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV4AdminDatabasesDatabaseNameDictionaryTablesTableName**](AdminApi.md#GetApiV4AdminDatabasesDatabaseNameDictionaryTablesTableName) | **Get** /api/v4/admin/databases/{database_name}/dictionary/tables/{table_name} | 


# **GetApiV4AdminDatabasesDatabaseNameDictionaryTablesTableName**
> ApiEntitiesDictionaryTable GetApiV4AdminDatabasesDatabaseNameDictionaryTablesTableName(ctx, databaseName, tableName)


Retrieve dictionary details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **databaseName** | **string**| The database name | 
  **tableName** | **string**| The table name | 

### Return type

[**ApiEntitiesDictionaryTable**](API_Entities_Dictionary_Table.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

