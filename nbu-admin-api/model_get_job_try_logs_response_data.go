/*
 * NetBackup Administration API
 *
 * The NetBackup Administration API provides access to administrative operations in NetBackup.
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type GetJobTryLogsResponseData struct {
	Type_      string                           `json:"type,omitempty"`
	Id         string                           `json:"id,omitempty"`
	Links      *GetJobTryLogsResponseLinks      `json:"links,omitempty"`
	Attributes *GetJobTryLogsResponseAttributes `json:"attributes,omitempty"`
}
