/*
 * NetBackup Administration API
 *
 * The NetBackup Administration API provides access to administrative operations in NetBackup.
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type GetJobProgressLogsResponseData struct {
	// The offset of the log message.
	Id         int32                                 `json:"id"`
	Type_      string                                `json:"type"`
	Attributes *GetJobProgressLogsResponseAttributes `json:"attributes"`
}
