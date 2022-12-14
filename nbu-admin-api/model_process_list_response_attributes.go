/*
 * NetBackup Administration API
 *
 * The NetBackup Administration API provides access to administrative operations in NetBackup.
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"time"
)

type ProcessListResponseAttributes struct {
	// The name of the process.
	ProcessName string `json:"processName"`
	// The process ID of the parent process.
	Ppid int32 `json:"ppid,omitempty"`
	// The priority of the process.
	Priority int32 `json:"priority"`
	// The memory usage of the process in MB.
	MemoryUsageMB float64 `json:"memoryUsageMB"`
	// The time the process started in ISO 8601 time format.
	StartTime time.Time `json:"startTime"`
	// The duration of the process since startTime, in ISO 8601 time format.
	ElapsedTime string `json:"elapsedTime"`
	// The username that started the process.
	Username string `json:"username,omitempty"`
	// The arguments provided to the process.
	Arguments string `json:"arguments,omitempty"`
}
