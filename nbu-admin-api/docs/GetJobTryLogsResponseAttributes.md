# GetJobTryLogsResponseAttributes

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TryJobId** | **int32** | The job ID for this try attempt. | [default to null]
**TryNumber** | **int32** | The try attempt number for which the following logs were generated. | [default to null]
**TryStartTime** | [**time.Time**](time.Time.md) | The UTC timestamp of when this try started. | [default to null]
**TryEndTime** | [**time.Time**](time.Time.md) | The UTC timestamp of when this try finished. | [default to null]
**TryStatusCode** | **int32** | The status code associated with this try attempt. | [default to null]
**TryStatusMessage** | **string** | The status message associated with the status of this try attempt. | [default to null]
**TryProcessId** | **int32** | The process ID for this try attempt. | [optional] [default to null]
**TryKbps** | **float64** | The rate of data transfer for this try attempt. | [optional] [default to null]
**TryKbWritten** | **float64** | The total data transferred (in KB) for this try attempt. | [optional] [default to null]
**TryFilesWritten** | **int32** | The files written in this try attempt. | [optional] [default to null]
**TryMediaServer** | **string** | The media server that was used for this try attempt. | [optional] [default to null]
**TrySourceStorageUnit** | **string** | The source storage unit for this try attempt. | [optional] [default to null]
**TryDestinationStorageUnit** | **string** | The destination storage unit for this try attempt. | [optional] [default to null]
**TrySessionId** | **int32** | The session ID for this try attempt. | [optional] [default to null]
**TryTapesToEject** | **int32** | The number of tapes to eject for this try attempt. | [optional] [default to null]
**TryTransportMode** | **string** | The transport mode that was used for this try attempt. | [optional] [default to null]
**TryLog** | [**[]GetJobTryLogsResponseAttributesTryLog**](getJobTryLogsResponse_attributes_tryLog.md) | The list of try-logs, parsed as JSON objects and sorted. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

