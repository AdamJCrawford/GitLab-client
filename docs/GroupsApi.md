# \GroupsApi

All URIs are relative to *https://gitlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV4GroupsId**](GroupsApi.md#DeleteApiV4GroupsId) | **Delete** /api/v4/groups/{id} | 
[**DeleteApiV4GroupsIdBillableMembersUserId**](GroupsApi.md#DeleteApiV4GroupsIdBillableMembersUserId) | **Delete** /api/v4/groups/{id}/billable_members/{user_id} | 
[**DeleteApiV4GroupsIdCustomAttributesKey**](GroupsApi.md#DeleteApiV4GroupsIdCustomAttributesKey) | **Delete** /api/v4/groups/{id}/custom_attributes/{key} | 
[**DeleteApiV4GroupsIdMembersUserIdOverride**](GroupsApi.md#DeleteApiV4GroupsIdMembersUserIdOverride) | **Delete** /api/v4/groups/{id}/members/{user_id}/override | 
[**DeleteApiV4GroupsIdShareGroupId**](GroupsApi.md#DeleteApiV4GroupsIdShareGroupId) | **Delete** /api/v4/groups/{id}/share/{group_id} | 
[**DeleteApiV4GroupsIdSshCertificatesSshCertificatesId**](GroupsApi.md#DeleteApiV4GroupsIdSshCertificatesSshCertificatesId) | **Delete** /api/v4/groups/{id}/ssh_certificates/{ssh_certificates_id} | Removes an ssh certificate from a group.
[**GetApiV4Groups**](GroupsApi.md#GetApiV4Groups) | **Get** /api/v4/groups | 
[**GetApiV4GroupsId**](GroupsApi.md#GetApiV4GroupsId) | **Get** /api/v4/groups/{id} | 
[**GetApiV4GroupsIdAuditEvents**](GroupsApi.md#GetApiV4GroupsIdAuditEvents) | **Get** /api/v4/groups/{id}/audit_events | 
[**GetApiV4GroupsIdAuditEventsAuditEventId**](GroupsApi.md#GetApiV4GroupsIdAuditEventsAuditEventId) | **Get** /api/v4/groups/{id}/audit_events/{audit_event_id} | 
[**GetApiV4GroupsIdBillableMembers**](GroupsApi.md#GetApiV4GroupsIdBillableMembers) | **Get** /api/v4/groups/{id}/billable_members | 
[**GetApiV4GroupsIdBillableMembersUserIdIndirect**](GroupsApi.md#GetApiV4GroupsIdBillableMembersUserIdIndirect) | **Get** /api/v4/groups/{id}/billable_members/{user_id}/indirect | 
[**GetApiV4GroupsIdBillableMembersUserIdMemberships**](GroupsApi.md#GetApiV4GroupsIdBillableMembersUserIdMemberships) | **Get** /api/v4/groups/{id}/billable_members/{user_id}/memberships | 
[**GetApiV4GroupsIdCustomAttributes**](GroupsApi.md#GetApiV4GroupsIdCustomAttributes) | **Get** /api/v4/groups/{id}/custom_attributes | 
[**GetApiV4GroupsIdCustomAttributesKey**](GroupsApi.md#GetApiV4GroupsIdCustomAttributesKey) | **Get** /api/v4/groups/{id}/custom_attributes/{key} | 
[**GetApiV4GroupsIdDescendantGroups**](GroupsApi.md#GetApiV4GroupsIdDescendantGroups) | **Get** /api/v4/groups/{id}/descendant_groups | 
[**GetApiV4GroupsIdPendingMembers**](GroupsApi.md#GetApiV4GroupsIdPendingMembers) | **Get** /api/v4/groups/{id}/pending_members | 
[**GetApiV4GroupsIdProjects**](GroupsApi.md#GetApiV4GroupsIdProjects) | **Get** /api/v4/groups/{id}/projects | 
[**GetApiV4GroupsIdProjectsShared**](GroupsApi.md#GetApiV4GroupsIdProjectsShared) | **Get** /api/v4/groups/{id}/projects/shared | 
[**GetApiV4GroupsIdProvisionedUsers**](GroupsApi.md#GetApiV4GroupsIdProvisionedUsers) | **Get** /api/v4/groups/{id}/provisioned_users | 
[**GetApiV4GroupsIdRunners**](GroupsApi.md#GetApiV4GroupsIdRunners) | **Get** /api/v4/groups/{id}/runners | List group&#39;s runners
[**GetApiV4GroupsIdSshCertificates**](GroupsApi.md#GetApiV4GroupsIdSshCertificates) | **Get** /api/v4/groups/{id}/ssh_certificates | Get a list of Groups::SshCertificate for a Group.
[**GetApiV4GroupsIdSubgroups**](GroupsApi.md#GetApiV4GroupsIdSubgroups) | **Get** /api/v4/groups/{id}/subgroups | 
[**GetApiV4GroupsIdTransferLocations**](GroupsApi.md#GetApiV4GroupsIdTransferLocations) | **Get** /api/v4/groups/{id}/transfer_locations | 
[**GetApiV4GroupsIdUsers**](GroupsApi.md#GetApiV4GroupsIdUsers) | **Get** /api/v4/groups/{id}/users | 
[**PostApiV4Groups**](GroupsApi.md#PostApiV4Groups) | **Post** /api/v4/groups | 
[**PostApiV4GroupsIdLdapSync**](GroupsApi.md#PostApiV4GroupsIdLdapSync) | **Post** /api/v4/groups/{id}/ldap_sync | 
[**PostApiV4GroupsIdMembersApproveAll**](GroupsApi.md#PostApiV4GroupsIdMembersApproveAll) | **Post** /api/v4/groups/{id}/members/approve_all | 
[**PostApiV4GroupsIdMembersUserIdOverride**](GroupsApi.md#PostApiV4GroupsIdMembersUserIdOverride) | **Post** /api/v4/groups/{id}/members/{user_id}/override | 
[**PostApiV4GroupsIdProjectsProjectId**](GroupsApi.md#PostApiV4GroupsIdProjectsProjectId) | **Post** /api/v4/groups/{id}/projects/{project_id} | 
[**PostApiV4GroupsIdRestore**](GroupsApi.md#PostApiV4GroupsIdRestore) | **Post** /api/v4/groups/{id}/restore | 
[**PostApiV4GroupsIdRunnersResetRegistrationToken**](GroupsApi.md#PostApiV4GroupsIdRunnersResetRegistrationToken) | **Post** /api/v4/groups/{id}/runners/reset_registration_token | Reset the runner registration token for a group
[**PostApiV4GroupsIdShare**](GroupsApi.md#PostApiV4GroupsIdShare) | **Post** /api/v4/groups/{id}/share | 
[**PostApiV4GroupsIdSshCertificates**](GroupsApi.md#PostApiV4GroupsIdSshCertificates) | **Post** /api/v4/groups/{id}/ssh_certificates | Add a Groups::SshCertificate.
[**PostApiV4GroupsIdTransfer**](GroupsApi.md#PostApiV4GroupsIdTransfer) | **Post** /api/v4/groups/{id}/transfer | 
[**PostApiV4RunnersResetRegistrationToken**](GroupsApi.md#PostApiV4RunnersResetRegistrationToken) | **Post** /api/v4/runners/reset_registration_token | Reset instance&#39;s runner registration token
[**PutApiV4GroupsId**](GroupsApi.md#PutApiV4GroupsId) | **Put** /api/v4/groups/{id} | 
[**PutApiV4GroupsIdCustomAttributesKey**](GroupsApi.md#PutApiV4GroupsIdCustomAttributesKey) | **Put** /api/v4/groups/{id}/custom_attributes/{key} | 
[**PutApiV4GroupsIdMembersMemberIdApprove**](GroupsApi.md#PutApiV4GroupsIdMembersMemberIdApprove) | **Put** /api/v4/groups/{id}/members/{member_id}/approve | 
[**PutApiV4GroupsIdMembersUserIdState**](GroupsApi.md#PutApiV4GroupsIdMembersUserIdState) | **Put** /api/v4/groups/{id}/members/{user_id}/state | 


# **DeleteApiV4GroupsId**
> DeleteApiV4GroupsId(ctx, id)


Remove a group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4GroupsIdBillableMembersUserId**
> DeleteApiV4GroupsIdBillableMembersUserId(ctx, id, userId)


Removes a billable member from a group or project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group | 
  **userId** | **int32**| The user ID of the member | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4GroupsIdCustomAttributesKey**
> DeleteApiV4GroupsIdCustomAttributesKey(ctx, key, id)


Delete a custom attribute on a group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| The key of the custom attribute | 
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4GroupsIdMembersUserIdOverride**
> ApiEntitiesMember DeleteApiV4GroupsIdMembersUserIdOverride(ctx, id, userId)


Remove an LDAP group member access level override.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group | 
  **userId** | **int32**| The user ID of the member | 

### Return type

[**ApiEntitiesMember**](API_Entities_Member.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4GroupsIdShareGroupId**
> DeleteApiV4GroupsIdShareGroupId(ctx, id, groupId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group | 
  **groupId** | **int32**| The ID of the shared group | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiV4GroupsIdSshCertificatesSshCertificatesId**
> DeleteApiV4GroupsIdSshCertificatesSshCertificatesId(ctx, id, sshCertificatesId)
Removes an ssh certificate from a group.

Removes a Groups::SshCertificate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **sshCertificatesId** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4Groups**
> []ApiEntitiesGroup GetApiV4Groups(ctx, optional)


Get a groups list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsApiGetApiV4GroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGetApiV4GroupsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **statistics** | **optional.Bool**| Include project statistics | [default to false]
 **skipGroups** | [**optional.Interface of []int32**](int32.md)| Array of group ids to exclude from list | 
 **allAvailable** | **optional.Bool**| Show all group that you have access to | 
 **visibility** | **optional.String**| Limit by visibility | 
 **search** | **optional.String**| Search for a specific group | 
 **owned** | **optional.Bool**| Limit by owned by authenticated user | [default to false]
 **orderBy** | **optional.String**| Order by name, path, id or similarity if searching | [default to name]
 **sort** | **optional.String**| Sort by asc (ascending) or desc (descending) | [default to asc]
 **minAccessLevel** | **optional.Int32**| Minimum access level of authenticated user | 
 **topLevelOnly** | **optional.Bool**| Only include top level groups | 
 **repositoryStorage** | **optional.String**| Filter by repository storage used by the group | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **withCustomAttributes** | **optional.Bool**| Include custom attributes in the response | [default to false]

### Return type

[**[]ApiEntitiesGroup**](API_Entities_Group.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsId**
> ApiEntitiesGroupDetail GetApiV4GroupsId(ctx, id, optional)


Get a single group, with containing projects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group | 
 **optional** | ***GroupsApiGetApiV4GroupsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGetApiV4GroupsIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withCustomAttributes** | **optional.Bool**| Include custom attributes in the response | [default to false]
 **withProjects** | **optional.Bool**| Omit project details | [default to true]

### Return type

[**ApiEntitiesGroupDetail**](API_Entities_GroupDetail.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdAuditEvents**
> []EeApiEntitiesAuditEvent GetApiV4GroupsIdAuditEvents(ctx, id, optional)


Get a list of audit events in this group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GroupsApiGetApiV4GroupsIdAuditEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGetApiV4GroupsIdAuditEventsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createdAfter** | **optional.Time**| Return audit events created after the specified time | 
 **createdBefore** | **optional.Time**| Return audit events created before the specified time | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]EeApiEntitiesAuditEvent**](EE_API_Entities_AuditEvent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdAuditEventsAuditEventId**
> EeApiEntitiesAuditEvent GetApiV4GroupsIdAuditEventsAuditEventId(ctx, auditEventId, id)


Get a specific audit event in this group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **auditEventId** | **int32**| The ID of the audit event | 
  **id** | **int32**|  | 

### Return type

[**EeApiEntitiesAuditEvent**](EE_API_Entities_AuditEvent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdBillableMembers**
> ApiEntitiesMember GetApiV4GroupsIdBillableMembers(ctx, id, optional)


Gets a list of billable users of root group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group | 
 **optional** | ***GroupsApiGetApiV4GroupsIdBillableMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGetApiV4GroupsIdBillableMembersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **search** | **optional.String**| The exact name of the subscribed member | 
 **sort** | **optional.String**| The sorting option | 

### Return type

[**ApiEntitiesMember**](API_Entities_Member.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdBillableMembersUserIdIndirect**
> EeApiEntitiesBillableMembership GetApiV4GroupsIdBillableMembersUserIdIndirect(ctx, id, userId, optional)


Get the indirect memberships of a billable user of a root group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group | 
  **userId** | **int32**| The user ID of the member | 
 **optional** | ***GroupsApiGetApiV4GroupsIdBillableMembersUserIdIndirectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGetApiV4GroupsIdBillableMembersUserIdIndirectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**EeApiEntitiesBillableMembership**](EE_API_Entities_BillableMembership.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdBillableMembersUserIdMemberships**
> EeApiEntitiesBillableMembership GetApiV4GroupsIdBillableMembersUserIdMemberships(ctx, id, userId, optional)


Get the direct memberships of a billable user of a root group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group | 
  **userId** | **int32**| The user ID of the member | 
 **optional** | ***GroupsApiGetApiV4GroupsIdBillableMembersUserIdMembershipsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGetApiV4GroupsIdBillableMembersUserIdMembershipsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**EeApiEntitiesBillableMembership**](EE_API_Entities_BillableMembership.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdCustomAttributes**
> ApiEntitiesCustomAttribute GetApiV4GroupsIdCustomAttributes(ctx, id)


Get all custom attributes on a group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**ApiEntitiesCustomAttribute**](API_Entities_CustomAttribute.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdCustomAttributesKey**
> ApiEntitiesCustomAttribute GetApiV4GroupsIdCustomAttributesKey(ctx, key, id)


Get a custom attribute on a group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| The key of the custom attribute | 
  **id** | **int32**|  | 

### Return type

[**ApiEntitiesCustomAttribute**](API_Entities_CustomAttribute.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdDescendantGroups**
> []ApiEntitiesGroup GetApiV4GroupsIdDescendantGroups(ctx, id, optional)


Get a list of descendant groups of this group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group | 
 **optional** | ***GroupsApiGetApiV4GroupsIdDescendantGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGetApiV4GroupsIdDescendantGroupsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **statistics** | **optional.Bool**| Include project statistics | [default to false]
 **skipGroups** | [**optional.Interface of []int32**](int32.md)| Array of group ids to exclude from list | 
 **allAvailable** | **optional.Bool**| Show all group that you have access to | 
 **visibility** | **optional.String**| Limit by visibility | 
 **search** | **optional.String**| Search for a specific group | 
 **owned** | **optional.Bool**| Limit by owned by authenticated user | [default to false]
 **orderBy** | **optional.String**| Order by name, path, id or similarity if searching | [default to name]
 **sort** | **optional.String**| Sort by asc (ascending) or desc (descending) | [default to asc]
 **minAccessLevel** | **optional.Int32**| Minimum access level of authenticated user | 
 **topLevelOnly** | **optional.Bool**| Only include top level groups | 
 **repositoryStorage** | **optional.String**| Filter by repository storage used by the group | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **withCustomAttributes** | **optional.Bool**| Include custom attributes in the response | [default to false]

### Return type

[**[]ApiEntitiesGroup**](API_Entities_Group.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdPendingMembers**
> GetApiV4GroupsIdPendingMembers(ctx, id, optional)


Lists all pending members for a group including invited users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group | 
 **optional** | ***GroupsApiGetApiV4GroupsIdPendingMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGetApiV4GroupsIdPendingMembersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdProjects**
> []ApiEntitiesProject GetApiV4GroupsIdProjects(ctx, id, optional)


Get a list of projects in this group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group | 
 **optional** | ***GroupsApiGetApiV4GroupsIdProjectsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGetApiV4GroupsIdProjectsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **optional.Bool**| Limit by archived status | 
 **visibility** | **optional.String**| Limit by visibility | 
 **search** | **optional.String**| Return list of authorized projects matching the search criteria | 
 **orderBy** | **optional.String**| Return projects ordered by field | [default to created_at]
 **sort** | **optional.String**| Return projects sorted in ascending and descending order | [default to desc]
 **simple** | **optional.Bool**| Return only the ID, URL, name, and path of each project | [default to false]
 **owned** | **optional.Bool**| Limit by owned by authenticated user | [default to false]
 **starred** | **optional.Bool**| Limit by starred status | [default to false]
 **withIssuesEnabled** | **optional.Bool**| Limit by enabled issues feature | [default to false]
 **withMergeRequestsEnabled** | **optional.Bool**| Limit by enabled merge requests feature | [default to false]
 **withShared** | **optional.Bool**| Include projects shared to this group | [default to true]
 **includeSubgroups** | **optional.Bool**| Includes projects in subgroups of this group | [default to false]
 **includeAncestorGroups** | **optional.Bool**| Includes projects in ancestors of this group | [default to false]
 **minAccessLevel** | **optional.Int32**| Limit by minimum access level of authenticated user on projects | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **withCustomAttributes** | **optional.Bool**| Include custom attributes in the response | [default to false]
 **withSecurityReports** | **optional.Bool**| Return only projects having security report artifacts present | [default to false]

### Return type

[**[]ApiEntitiesProject**](API_Entities_Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdProjectsShared**
> []ApiEntitiesProject GetApiV4GroupsIdProjectsShared(ctx, id, optional)


Get a list of shared projects in this group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group | 
 **optional** | ***GroupsApiGetApiV4GroupsIdProjectsSharedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGetApiV4GroupsIdProjectsSharedOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **optional.Bool**| Limit by archived status | 
 **visibility** | **optional.String**| Limit by visibility | 
 **search** | **optional.String**| Return list of authorized projects matching the search criteria | 
 **orderBy** | **optional.String**| Return projects ordered by field | [default to created_at]
 **sort** | **optional.String**| Return projects sorted in ascending and descending order | [default to desc]
 **simple** | **optional.Bool**| Return only the ID, URL, name, and path of each project | [default to false]
 **starred** | **optional.Bool**| Limit by starred status | [default to false]
 **withIssuesEnabled** | **optional.Bool**| Limit by enabled issues feature | [default to false]
 **withMergeRequestsEnabled** | **optional.Bool**| Limit by enabled merge requests feature | [default to false]
 **minAccessLevel** | **optional.Int32**| Limit by minimum access level of authenticated user on projects | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **withCustomAttributes** | **optional.Bool**| Include custom attributes in the response | [default to false]

### Return type

[**[]ApiEntitiesProject**](API_Entities_Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdProvisionedUsers**
> ApiEntitiesUserPublic GetApiV4GroupsIdProvisionedUsers(ctx, id, optional)


Get a list of users provisioned by the group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GroupsApiGetApiV4GroupsIdProvisionedUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGetApiV4GroupsIdProvisionedUsersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **optional.String**| Return a single user with a specific username | 
 **search** | **optional.String**| Search users by name, email or username | 
 **active** | **optional.Bool**| Return only active users | [default to false]
 **blocked** | **optional.Bool**| Return only blocked users | [default to false]
 **createdAfter** | **optional.Time**| Return users created after the specified time | 
 **createdBefore** | **optional.Time**| Return users created before the specified time | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**ApiEntitiesUserPublic**](API_Entities_UserPublic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdRunners**
> ApiEntitiesCiRunner GetApiV4GroupsIdRunners(ctx, id, optional)
List group's runners

List all runners available in the group as well as its ancestor groups, including any allowed shared runners.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group | 
 **optional** | ***GroupsApiGetApiV4GroupsIdRunnersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGetApiV4GroupsIdRunnersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**| The type of runners to return | 
 **paused** | **optional.Bool**| Whether to include only runners that are accepting or ignoring new jobs | 
 **status** | **optional.String**| The status of runners to return | 
 **tagList** | [**optional.Interface of []string**](string.md)| A list of runner tags | 
 **versionPrefix** | **optional.String**| The version prefix of runners to return | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**ApiEntitiesCiRunner**](API_Entities_Ci_Runner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdSshCertificates**
> []EeApiEntitiesSshCertificate GetApiV4GroupsIdSshCertificates(ctx, id, optional)
Get a list of Groups::SshCertificate for a Group.

Get a list of ssh certificates created for a group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GroupsApiGetApiV4GroupsIdSshCertificatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGetApiV4GroupsIdSshCertificatesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]EeApiEntitiesSshCertificate**](EE_API_Entities_SshCertificate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdSubgroups**
> []ApiEntitiesGroup GetApiV4GroupsIdSubgroups(ctx, id, optional)


Get a list of subgroups in this group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group | 
 **optional** | ***GroupsApiGetApiV4GroupsIdSubgroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGetApiV4GroupsIdSubgroupsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **statistics** | **optional.Bool**| Include project statistics | [default to false]
 **skipGroups** | [**optional.Interface of []int32**](int32.md)| Array of group ids to exclude from list | 
 **allAvailable** | **optional.Bool**| Show all group that you have access to | 
 **visibility** | **optional.String**| Limit by visibility | 
 **search** | **optional.String**| Search for a specific group | 
 **owned** | **optional.Bool**| Limit by owned by authenticated user | [default to false]
 **orderBy** | **optional.String**| Order by name, path, id or similarity if searching | [default to name]
 **sort** | **optional.String**| Sort by asc (ascending) or desc (descending) | [default to asc]
 **minAccessLevel** | **optional.Int32**| Minimum access level of authenticated user | 
 **topLevelOnly** | **optional.Bool**| Only include top level groups | 
 **repositoryStorage** | **optional.String**| Filter by repository storage used by the group | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]
 **withCustomAttributes** | **optional.Bool**| Include custom attributes in the response | [default to false]

### Return type

[**[]ApiEntitiesGroup**](API_Entities_Group.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdTransferLocations**
> []ApiEntitiesGroup GetApiV4GroupsIdTransferLocations(ctx, id, optional)


Get the groups to where the current group can be transferred to

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group | 
 **optional** | ***GroupsApiGetApiV4GroupsIdTransferLocationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGetApiV4GroupsIdTransferLocationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| Return list of namespaces matching the search criteria | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**[]ApiEntitiesGroup**](API_Entities_Group.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiV4GroupsIdUsers**
> ApiEntitiesUserPublic GetApiV4GroupsIdUsers(ctx, id, optional)


Get a list of users for the group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GroupsApiGetApiV4GroupsIdUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGetApiV4GroupsIdUsersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| Search users by name, email or username | 
 **active** | **optional.Bool**| Filters only active users | [default to false]
 **includeSamlUsers** | **optional.Bool**| Return users with a SAML identity in this group | 
 **includeServiceAccounts** | **optional.Bool**| Return service accounts owned by this group | 
 **page** | **optional.Int32**| Current page number | [default to 1]
 **perPage** | **optional.Int32**| Number of items per page | [default to 20]

### Return type

[**ApiEntitiesUserPublic**](API_Entities_UserPublic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4Groups**
> ApiEntitiesGroup PostApiV4Groups(ctx, postApiV4Groups)


Create a group. Available only for users who can create groups.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postApiV4Groups** | [**PostApiV4Groups**](PostApiV4Groups.md)|  | 

### Return type

[**ApiEntitiesGroup**](API_Entities_Group.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4GroupsIdLdapSync**
> PostApiV4GroupsIdLdapSync(ctx, id)


Sync a group with LDAP.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4GroupsIdMembersApproveAll**
> PostApiV4GroupsIdMembersApproveAll(ctx, id)


Approves all pending members

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4GroupsIdMembersUserIdOverride**
> ApiEntitiesMember PostApiV4GroupsIdMembersUserIdOverride(ctx, id, userId)


Overrides the access level of an LDAP group member.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group | 
  **userId** | **int32**| The user ID of the member | 

### Return type

[**ApiEntitiesMember**](API_Entities_Member.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4GroupsIdProjectsProjectId**
> ApiEntitiesGroupDetail PostApiV4GroupsIdProjectsProjectId(ctx, id, projectId)


Transfer a project to the group namespace. Available only for admin.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group | 
  **projectId** | **string**| The ID or path of the project | 

### Return type

[**ApiEntitiesGroupDetail**](API_Entities_GroupDetail.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4GroupsIdRestore**
> PostApiV4GroupsIdRestore(ctx, id)


Restore a group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4GroupsIdRunnersResetRegistrationToken**
> ApiEntitiesCiResetTokenResult PostApiV4GroupsIdRunnersResetRegistrationToken(ctx, id)
Reset the runner registration token for a group

Reset runner registration token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group | 

### Return type

[**ApiEntitiesCiResetTokenResult**](API_Entities_Ci_ResetTokenResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4GroupsIdShare**
> ApiEntitiesGroupDetail PostApiV4GroupsIdShare(ctx, id, postApiV4GroupsIdShare)


Share a group with a group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group | 
  **postApiV4GroupsIdShare** | [**PostApiV4GroupsIdShare**](PostApiV4GroupsIdShare.md)|  | 

### Return type

[**ApiEntitiesGroupDetail**](API_Entities_GroupDetail.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4GroupsIdSshCertificates**
> EeApiEntitiesSshCertificate PostApiV4GroupsIdSshCertificates(ctx, id, postApiV4GroupsIdSshCertificates)
Add a Groups::SshCertificate.

Create a ssh certificate for a group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **postApiV4GroupsIdSshCertificates** | [**PostApiV4GroupsIdSshCertificates**](PostApiV4GroupsIdSshCertificates.md)|  | 

### Return type

[**EeApiEntitiesSshCertificate**](EE_API_Entities_SshCertificate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4GroupsIdTransfer**
> PostApiV4GroupsIdTransfer(ctx, id, postApiV4GroupsIdTransfer)


Transfer a group to a new parent group or promote a subgroup to a root group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group | 
  **postApiV4GroupsIdTransfer** | [**PostApiV4GroupsIdTransfer**](PostApiV4GroupsIdTransfer.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApiV4RunnersResetRegistrationToken**
> ApiEntitiesCiResetTokenResult PostApiV4RunnersResetRegistrationToken(ctx, )
Reset instance's runner registration token

Reset runner registration token

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiEntitiesCiResetTokenResult**](API_Entities_Ci_ResetTokenResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4GroupsId**
> ApiEntitiesGroup PutApiV4GroupsId(ctx, id, putApiV4GroupsId)


Update a group. Available only for users who can administrate groups.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group | 
  **putApiV4GroupsId** | [**PutApiV4GroupsId**](PutApiV4GroupsId.md)|  | 

### Return type

[**ApiEntitiesGroup**](API_Entities_Group.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4GroupsIdCustomAttributesKey**
> PutApiV4GroupsIdCustomAttributesKey(ctx, key, id, putApiV4GroupsIdCustomAttributesKey)


Set a custom attribute on a group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| The key of the custom attribute | 
  **id** | **int32**|  | 
  **putApiV4GroupsIdCustomAttributesKey** | [**PutApiV4GroupsIdCustomAttributesKey**](PutApiV4GroupsIdCustomAttributesKey.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4GroupsIdMembersMemberIdApprove**
> PutApiV4GroupsIdMembersMemberIdApprove(ctx, id, memberId)


Approves a pending member

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group | 
  **memberId** | **int32**| The ID of the member requiring approval | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutApiV4GroupsIdMembersUserIdState**
> PutApiV4GroupsIdMembersUserIdState(ctx, id, userId, putApiV4GroupsIdMembersUserIdState)


Changes the state of the memberships of a user in the group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of a group | 
  **userId** | **int32**| The user ID of the user | 
  **putApiV4GroupsIdMembersUserIdState** | [**PutApiV4GroupsIdMembersUserIdState**](PutApiV4GroupsIdMembersUserIdState.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

