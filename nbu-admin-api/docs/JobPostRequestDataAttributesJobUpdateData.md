# JobPostRequestDataAttributesJobUpdateData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestoreBackupIDs** | **[]string** | The identifier for the backup image that was used for the restore. | [optional] [default to null]
**StartTime** | [**time.Time**](time.Time.md) | The UTC timestamp for when the job was started. | [optional] [default to null]
**EndTime** | [**time.Time**](time.Time.md) | The UTC timestamp for when the job ended. | [optional] [default to null]
**ActiveTryStartTime** | **string** | A UTC timestamp for when the current try started. | [optional] [default to null]
**LastUpdateTime** | [**time.Time**](time.Time.md) | The UTC timestamp for when the job was modified. | [optional] [default to null]
**State** | **string** | The current state of the job. | [optional] [default to null]
**Status** | **int32** | The final job status code. | [optional] [default to null]
**NewTry** | **bool** | Indicates if a new try log should be started. | [optional] [default to false]
**EndTry** | **bool** | Indicates if the current try log should end. | [optional] [default to false]
**StartTrack** | **bool** | If set to true bpjobd will track the job. | [optional] [default to true]
**StopTrack** | **bool** | Indicates if bpjobd will stop tracking the job. | [optional] [default to false]
**Done** | **bool** | Indicates if bpjobd will track the job. | [optional] [default to false]
**TryLogList** | [***JobPostRequestDataAttributesJobUpdateDataTryLogList**](jobPostRequest_data_attributes_jobUpdateData_tryLogList.md) |  | [optional] [default to null]
**FileList** | **[]string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

