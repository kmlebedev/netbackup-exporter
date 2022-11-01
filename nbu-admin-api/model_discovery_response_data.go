/*
 * NetBackup Administration API
 *
 * The NetBackup Administration API provides access to administrative operations in NetBackup.
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DiscoveryResponseData struct {
	// The response type that is associated with discovery request.
	Type_ string `json:"type"`
	// The unique identifier of the discovery request.
	Id         string                           `json:"id"`
	Attributes *DiscoveryResponseDataAttributes `json:"attributes"`
}
