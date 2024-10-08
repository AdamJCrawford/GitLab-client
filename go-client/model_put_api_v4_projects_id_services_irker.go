/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Create/Edit Irker integration
type PutApiV4ProjectsIdServicesIrker struct {
	// Recipients/channels separated by whitespaces
	Recipients string `json:"recipients"`
	// Default: irc://irc.network.net:6697
	DefaultIrcUri string `json:"default_irc_uri,omitempty"`
	// Server host. Default localhost
	ServerHost string `json:"server_host,omitempty"`
	// Server port. Default 6659
	ServerPort int32 `json:"server_port,omitempty"`
	// Colorize messages
	ColorizeMessages bool `json:"colorize_messages,omitempty"`
	// Trigger event for pushes to the repository.
	PushEvents bool `json:"push_events,omitempty"`
}
