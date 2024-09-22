# PostApiV4Users

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The email of the user | [default to null]
**Password** | **string** | The password of the new user | [optional] [default to null]
**ResetPassword** | **bool** | Flag indicating the user will be sent a password reset token | [optional] [default to null]
**SkipConfirmation** | **bool** | Flag indicating the account is confirmed | [optional] [default to null]
**ForceRandomPassword** | **bool** | Flag indicating a random password will be set | [optional] [default to null]
**Name** | **string** | The name of the user | [default to null]
**Username** | **string** | The username of the user | [default to null]
**Skype** | **string** | The Skype username | [optional] [default to null]
**Linkedin** | **string** | The LinkedIn username | [optional] [default to null]
**Twitter** | **string** | The Twitter username | [optional] [default to null]
**Discord** | **string** | The Discord user ID | [optional] [default to null]
**WebsiteUrl** | **string** | The website of the user | [optional] [default to null]
**Organization** | **string** | The organization of the user | [optional] [default to null]
**ProjectsLimit** | **int32** | The number of projects a user can create | [optional] [default to null]
**ExternUid** | **string** | The external authentication provider UID | [optional] [default to null]
**Provider** | **string** | The external provider | [optional] [default to null]
**Bio** | **string** | The biography of the user | [optional] [default to null]
**Location** | **string** | The location of the user | [optional] [default to null]
**Pronouns** | **string** | The pronouns of the user | [optional] [default to null]
**PublicEmail** | **string** | The public email of the user | [optional] [default to null]
**CommitEmail** | **string** | The commit email, _private for private commit email | [optional] [default to null]
**Admin** | **bool** | Flag indicating the user is an administrator | [optional] [default to null]
**CanCreateGroup** | **bool** | Flag indicating the user can create groups | [optional] [default to null]
**External** | **bool** | Flag indicating the user is an external user | [optional] [default to null]
**Avatar** | [****os.File**](*os.File.md) | Avatar image for user | [optional] [default to null]
**ThemeId** | **int32** | The GitLab theme for the user | [optional] [default to null]
**ColorSchemeId** | **int32** | The color scheme for the file viewer | [optional] [default to null]
**PrivateProfile** | **bool** | Flag indicating the user has a private profile | [optional] [default to null]
**Note** | **string** | Admin note for this user | [optional] [default to null]
**ViewDiffsFileByFile** | **bool** | Flag indicating the user sees only one file diff per page | [optional] [default to null]
**SharedRunnersMinutesLimit** | **int32** | Compute minutes quota for this user | [optional] [default to null]
**ExtraSharedRunnersMinutesLimit** | **int32** | (admin-only) Extra compute minutes quota for this user | [optional] [default to null]
**GroupIdForSaml** | **int32** | ID for group where SAML has been configured | [optional] [default to null]
**Auditor** | **bool** | Flag indicating auditor status of the user | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


