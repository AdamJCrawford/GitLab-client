# PostApiV4Groups

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the group | [default to null]
**Path** | **string** | The path of the group | [default to null]
**ParentId** | **int32** | The parent group id for creating nested group | [optional] [default to null]
**OrganizationId** | **int32** | The organization id for the group | [optional] [default to null]
**Description** | **string** | The description of the group | [optional] [default to null]
**Visibility** | **string** | The visibility of the group | [optional] [default to null]
**Avatar** | [****os.File**](*os.File.md) | Avatar image for the group | [optional] [default to null]
**ShareWithGroupLock** | **bool** | Prevent sharing a project with another group within this group | [optional] [default to null]
**RequireTwoFactorAuthentication** | **bool** | Require all users in this group to setup Two-factor authentication | [optional] [default to null]
**TwoFactorGracePeriod** | **int32** | Time before Two-factor authentication is enforced | [optional] [default to null]
**ProjectCreationLevel** | **string** | Determine if developers can create projects in the group | [optional] [default to null]
**AutoDevopsEnabled** | **bool** | Default to Auto DevOps pipeline for all projects within this group | [optional] [default to null]
**SubgroupCreationLevel** | **string** | Allowed to create subgroups | [optional] [default to null]
**EmailsDisabled** | **bool** | _(Deprecated)_ Disable email notifications. Use: emails_enabled | [optional] [default to null]
**EmailsEnabled** | **bool** | Enable email notifications | [optional] [default to null]
**ShowDiffPreviewInEmail** | **bool** | Include the code diff preview in merge request notification emails | [optional] [default to null]
**MentionsDisabled** | **bool** | Disable a group from getting mentioned | [optional] [default to null]
**LfsEnabled** | **bool** | Enable/disable LFS for the projects in this group | [optional] [default to null]
**RequestAccessEnabled** | **bool** | Allow users to request member access | [optional] [default to null]
**DefaultBranch** | **string** | The default branch of group&#39;s projects | [optional] [default to null]
**DefaultBranchProtection** | **int32** | Determine if developers can push to default branch | [optional] [default to null]
**DefaultBranchProtectionDefaults** | [***PostApiV4GroupsDefaultBranchProtectionDefaults**](postApiV4Groups_default_branch_protection_defaults.md) |  | [optional] [default to null]
**SharedRunnersSetting** | **string** | Enable/disable shared runners for the group and its subgroups and projects | [optional] [default to null]
**EnabledGitAccessProtocol** | **string** | Allow only the selected protocols to be used for Git access. | [optional] [default to null]
**MembershipLock** | **bool** | Prevent adding new members to projects within this group | [optional] [default to null]
**LdapCn** | **string** | LDAP Common Name | [optional] [default to null]
**LdapAccess** | **int32** | A valid access level | [optional] [default to null]
**SharedRunnersMinutesLimit** | **int32** | (admin-only) Compute minutes quota for this group | [optional] [default to null]
**ExtraSharedRunnersMinutesLimit** | **int32** | (admin-only) Extra compute minutes quota for this group | [optional] [default to null]
**WikiAccessLevel** | **string** | Wiki access level. One of &#x60;disabled&#x60;, &#x60;private&#x60; or &#x60;enabled&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


