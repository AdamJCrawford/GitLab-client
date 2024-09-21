/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Create/Edit Campfire integration
type PutApiV4ProjectsIdIntegrationsCampfire struct {
	// API authentication token from Campfire. To get the token, sign in to Campfire and select **My info**.
	Token string `json:"token"`
	// `.campfirenow.com` subdomain when you're signed in.
	Subdomain string `json:"subdomain,omitempty"`
	// ID portion of the Campfire room URL.
	Room string `json:"room,omitempty"`
	// Trigger event for pushes to the repository.
	PushEvents bool `json:"push_events,omitempty"`
}