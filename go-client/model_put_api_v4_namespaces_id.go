/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Update a namespace
type PutApiV4NamespacesId struct {
	// Compute minutes quota for this namespace
	SharedRunnersMinutesLimit int32 `json:"shared_runners_minutes_limit,omitempty"`
	// Extra compute minutes for this namespace
	ExtraSharedRunnersMinutesLimit int32 `json:"extra_shared_runners_minutes_limit,omitempty"`
	// Additional storage size for this namespace
	AdditionalPurchasedStorageSize int32 `json:"additional_purchased_storage_size,omitempty"`
	// End of subscription of the additional purchased storage
	AdditionalPurchasedStorageEndsOn string                                            `json:"additional_purchased_storage_ends_on,omitempty"`
	GitlabSubscriptionAttributes     *PutApiV4NamespacesIdGitlabSubscriptionAttributes `json:"gitlab_subscription_attributes,omitempty"`
}
