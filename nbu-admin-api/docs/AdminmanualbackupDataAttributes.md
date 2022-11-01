# AdminmanualbackupDataAttributes

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyName** | **string** | The name of the policy used to start the manual backup. | [default to null]
**ScheduleName** | **string** | The name of the schedule used to start the manual backup. | [optional] [default to null]
**ClientName** | **string** | The name of the client. | [optional] [default to null]
**Keyword** | **string** | A keyword that is added to the created image. Optional. | [optional] [default to null]
**InstanceName** | **string** | The name of the Oracle or Microsoft SQL Server instance to back up. | [optional] [default to null]
**DatabaseName** | **string** | The name of the Microsoft SQL Server database to back up. | [optional] [default to null]
**DbUniqueName** | **string** | The Oracle RAC database name. The database name should be used along with the dbid. | [optional] [default to null]
**Dbid** | **string** | The Oracle RAC database ID. The database ID should be used along with the dbUniqueName. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

