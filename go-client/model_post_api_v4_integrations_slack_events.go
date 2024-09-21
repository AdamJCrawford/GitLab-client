/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Receive Slack events
type PostApiV4IntegrationsSlackEvents struct {
	// (Deprecated by Slack) The request token, unused by GitLab
	Token string `json:"token,omitempty"`
	// The Slack workspace ID of where the event occurred
	TeamId string `json:"team_id,omitempty"`
	// The Slack app ID
	ApiAppId string `json:"api_app_id,omitempty"`
	// The event object with variable properties
	Event interface{} `json:"event,omitempty"`
	// The kind of event this is, usually `event_callback`
	Type_ string `json:"type,omitempty"`
	// A unique identifier for this specific event
	EventId string `json:"event_id,omitempty"`
	// The epoch timestamp in seconds when this event was dispatched
	EventTime int32 `json:"event_time,omitempty"`
	// (Deprecated by Slack) An array of Slack user IDs
	AuthedUsers []string `json:"authed_users,omitempty"`
}
