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

type GetJobTryLogsResponseAttributes struct {
	// The job ID for this try attempt.
	TryJobId int32 `json:"tryJobId"`
	// The try attempt number for which the following logs were generated.
	TryNumber int32 `json:"tryNumber"`
	// The UTC timestamp of when this try started.
	TryStartTime time.Time `json:"tryStartTime"`
	// The UTC timestamp of when this try finished.
	TryEndTime time.Time `json:"tryEndTime"`
	// The status code associated with this try attempt.
	TryStatusCode int32 `json:"tryStatusCode"`
	// The status message associated with the status of this try attempt.
	TryStatusMessage string `json:"tryStatusMessage"`
	// The process ID for this try attempt.
	TryProcessId int32 `json:"tryProcessId,omitempty"`
	// The rate of data transfer for this try attempt.
	TryKbps float64 `json:"tryKbps,omitempty"`
	// The total data transferred (in KB) for this try attempt.
	TryKbWritten float64 `json:"tryKbWritten,omitempty"`
	// The files written in this try attempt.
	TryFilesWritten int32 `json:"tryFilesWritten,omitempty"`
	// The media server that was used for this try attempt.
	TryMediaServer string `json:"tryMediaServer,omitempty"`
	// The source storage unit for this try attempt.
	TrySourceStorageUnit string `json:"trySourceStorageUnit,omitempty"`
	// The destination storage unit for this try attempt.
	TryDestinationStorageUnit string `json:"tryDestinationStorageUnit,omitempty"`
	// The session ID for this try attempt.
	TrySessionId int32 `json:"trySessionId,omitempty"`
	// The number of tapes to eject for this try attempt.
	TryTapesToEject int32 `json:"tryTapesToEject,omitempty"`
	// The transport mode that was used for this try attempt.
	TryTransportMode string `json:"tryTransportMode,omitempty"`
	// The list of try-logs, parsed as JSON objects and sorted.
	TryLog []GetJobTryLogsResponseAttributesTryLog `json:"tryLog"`
}