/*
 * NetBackup Administration API
 *
 * The NetBackup Administration API provides access to administrative operations in NetBackup.
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type GetJobsResponse struct {
	Links *GetJobsResponseLinks      `json:"links,omitempty"`
	Data  []ManualBackupResponseData `json:"data,omitempty"`
	Meta  *GetJobsResponseMeta       `json:"meta,omitempty"`
}
