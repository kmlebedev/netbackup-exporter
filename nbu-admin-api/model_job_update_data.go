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

type JobUpdateData struct {
	// The identifier for the backup image that was used for the restore.
	RestoreBackupIDs []string `json:"restoreBackupIDs,omitempty"`
	// The UTC timestamp for when the job was started.
	StartTime time.Time `json:"startTime,omitempty"`
	// The UTC timestamp for when the job ended.
	EndTime time.Time `json:"endTime,omitempty"`
	// A UTC timestamp for when the current try started.
	ActiveTryStartTime string `json:"activeTryStartTime,omitempty"`
	// The UTC timestamp for when the job was modified.
	LastUpdateTime time.Time `json:"lastUpdateTime,omitempty"`
	// The current state of the job.
	State string `json:"state,omitempty"`
	// The final job status code.
	Status int32 `json:"status,omitempty"`
	// Indicates if a new try log should be started.
	NewTry bool `json:"newTry,omitempty"`
	// Indicates if the current try log should end.
	EndTry bool `json:"endTry,omitempty"`
	// If set to true bpjobd will track the job.
	StartTrack bool `json:"startTrack,omitempty"`
	// Indicates if bpjobd will stop tracking the job.
	StopTrack bool `json:"stopTrack,omitempty"`
	// Indicates if bpjobd will track the job.
	Done       bool                                                 `json:"done,omitempty"`
	TryLogList *JobPostRequestDataAttributesJobUpdateDataTryLogList `json:"tryLogList,omitempty"`
	FileList   []string                                             `json:"fileList,omitempty"`
}
