# ProcessListResponseAttributes

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessName** | **string** | The name of the process. | [default to null]
**Ppid** | **int32** | The process ID of the parent process. | [optional] [default to null]
**Priority** | **int32** | The priority of the process. | [default to null]
**MemoryUsageMB** | **float64** | The memory usage of the process in MB. | [default to null]
**StartTime** | [**time.Time**](time.Time.md) | The time the process started in ISO 8601 time format. | [default to null]
**ElapsedTime** | **string** | The duration of the process since startTime, in ISO 8601 time format. | [default to null]
**Username** | **string** | The username that started the process. | [optional] [default to null]
**Arguments** | **string** | The arguments provided to the process. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

