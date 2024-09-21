/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Update a wiki page
type PutApiV4ProjectsIdWikisSlug struct {
	// Title of a wiki page
	Title       string                             `json:"title,omitempty"`
	FrontMatter *PostApiV4GroupsIdWikisFrontMatter `json:"front_matter,omitempty"`
	// Content of a wiki page
	Content string `json:"content,omitempty"`
	// Format of a wiki page. Available formats are markdown, rdoc, asciidoc and org
	Format string `json:"format,omitempty"`
}
