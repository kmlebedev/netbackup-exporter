/*
 * NetBackup Administration API
 *
 * The NetBackup Administration API provides access to administrative operations in NetBackup.
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The attributes of the discovery request.
type AdmindiscoveryworkloadsworkloadTypestartDataAttributes struct {
	// The name of the server to perform discovery on. Required for the following discovery workloads: vmware, rhv, cloud
	ServerName string `json:"serverName,omitempty"`
	// The NetBackup host to perform the discovery. Optional for the following workloads: vmware, rhv
	DiscoveryHost string `json:"discoveryHost,omitempty"`
}
