# JobData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobType** | **string** | The job type. | [default to null]
**JobSubType** | **string** | The job subtype. | [default to null]
**PolicyType** | **string** | The type of backup policy that is associated with the job. | [optional] [default to null]
**JobGroup** | **string** | The name of the group that initiated the job. | [optional] [default to null]
**JobOwner** | **string** | The job owner or the user that initiated the job. | [optional] [default to null]
**AuditUserName** | **string** | The user name in the audit record. | [optional] [default to null]
**AuditDomainName** | **string** | The domain name in the audit record. | [optional] [default to null]
**AuditDomainType** | **int32** | The domain type in the audit record. | [optional] [default to null]
**Restartable** | **bool** | Indicates if the job can be restarted. | [optional] [default to false]
**Cancellable** | **bool** | Indicates if the job can be stopped. | [optional] [default to false]
**Resumable** | **bool** | Indicates if the job can be resumed. | [optional] [default to false]
**Suspendable** | **bool** | Indicates if the job can be suspended. | [optional] [default to false]
**ControlHost** | **string** | The host from which the job was initiated. | [optional] [default to null]
**SourceMediaServerName** | **string** | The name of the source media server. | [optional] [default to null]
**DestinationMediaServerName** | **string** | The name of the destination media server. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

