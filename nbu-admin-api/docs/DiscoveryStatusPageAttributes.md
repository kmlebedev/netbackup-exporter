# DiscoveryStatusPageAttributes

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerName** | **string** | The name of the server. | [default to null]
**WorkloadType** | **string** |  | [default to null]
**WorkloadSubType** | **string** | The subtype as defined by the workload. | [default to null]
**DiscoveryStatus** | **string** | The status of the discovery process. | [default to null]
**DiscoveryStartTime** | [**time.Time**](time.Time.md) | The UTC timestamp of when the last attempted discovery started. | [default to null]
**DiscoveryFinishTime** | [**time.Time**](time.Time.md) | The UTC timestamp of when the last attempted discovery finished. | [default to null]
**UpdateTime** | [**time.Time**](time.Time.md) | The UTC timestamp of when the entry was updated. | [default to null]
**AdditionalAttributes** | [***DiscoveryStatusPageAttributesAdditionalAttributes**](discoveryStatusPage_attributes_additionalAttributes.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

