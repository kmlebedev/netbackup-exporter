# {{classname}}

All URIs are relative to *https://&lt;masterserver&gt;:1556/netbackup*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminDiscoveryWorkloadsWorkloadTypeStartPost**](DiscoveryApi.md#AdminDiscoveryWorkloadsWorkloadTypeStartPost) | **Post** /admin/discovery/workloads/{workloadType}/start | 
[**AdminDiscoveryWorkloadsWorkloadTypeStatusGet**](DiscoveryApi.md#AdminDiscoveryWorkloadsWorkloadTypeStatusGet) | **Get** /admin/discovery/workloads/{workloadType}/status | 

# **AdminDiscoveryWorkloadsWorkloadTypeStartPost**
> InlineResponse2021 AdminDiscoveryWorkloadsWorkloadTypeStartPost(ctx, body, workloadType)


Starts discovery to identify assets and topology for the specified workload type.  **Access Control**\\ Enforcement Type: [API-Level](http://sort.veritas.com/public/documents/nbu/8.3/windowsandunix/productguides/html/getting-started#user-content-api-level-enforcement)\\ Namespace: `|ASSETS|RHV|RHVMGR|`\\ OR Namespace: `|ASSETS|VMWARE|ESX|`\\ OR Namespace: `|ASSETS|VMWARE|VCENTER|`\\ OR Namespace: `|ASSETS|VMWARE|RESTORE_ESX|`\\ - only discovers topology Requires Operation: `|OPERATIONS|UPDATE|`\\ \\ For Cloud Workload:\\ Namespace: `|MANAGE|SNAPSHOT-MGMT-SERVER|`\\ Requires Operation: `|OPERATIONS|MANAGE|SNAPSHOT-MGMT-SERVER|DISCOVER|` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WorkloadTypeStartBody**](WorkloadTypeStartBody.md)| The request body of the discovery request. | 
  **workloadType** | **string**| The workload type. | 

### Return type

[**InlineResponse2021**](inline_response_202_1.md)

### Authorization

[rbacSecurity](../README.md#rbacSecurity)

### HTTP request headers

 - **Content-Type**: application/vnd.netbackup+json;version=4.0
 - **Accept**: application/vnd.netbackup+json;version=4.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminDiscoveryWorkloadsWorkloadTypeStatusGet**
> InlineResponse2009 AdminDiscoveryWorkloadsWorkloadTypeStatusGet(ctx, workloadType, optional)


Gets the discovery status for the specified workload type.  **Access Control**\\ Enforcement Type: [API-Level](http://sort.veritas.com/public/documents/nbu/8.3/windowsandunix/productguides/html/getting-started#user-content-api-level-enforcement)\\ Namespace: `|MANAGE|SNAPSHOT-MGMT-SERVER|`\\ OR Namespace: `|ASSETS|RHV|RHVMGR|`\\ OR Namespace: `|ASSETS|VMWARE|ESX|`\\ OR Namespace: `|ASSETS|VMWARE|VCENTER|`\\ OR Namespace: `|ASSETS|VMWARE|RESTORE_ESX|`\\ - only discovers topology Requires Operation: `|OPERATIONS|VIEW|` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workloadType** | **string**| The workload type. | 
 **optional** | ***DiscoveryApiAdminDiscoveryWorkloadsWorkloadTypeStatusGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DiscoveryApiAdminDiscoveryWorkloadsWorkloadTypeStatusGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| Filter criteria syntax is based on OData standards. Filtering is supported on the following attributes: |Name|Supported Operators|Description|Schema| |---|---|---|---| |serverName|eq ne contains, startswith, endswith|The name of the server.|string| |workloadSubType|eq ne contains, startswith, endswith|The workload subtype.|string| |discoveryStatus|eq ne|The status of the discovery process.|string| |discoveryStartTime|eq ne lt gt le ge|The UTC timestamp of when the last attempted discovery started.|date-time| |discoveryFinishTime|eq ne lt gt le ge|The UTC timestamp of when the last attempted discovery finished.|date-time| |updateTime|eq ne lt gt le ge|The UTC timestamp of when the entry was updated.|date-time|  | 
 **pageOffset** | **optional.Int32**| The discovery status record number offset. | 
 **pageLimit** | **optional.Int32**| The number of records on one page after the offset. | 
 **sort** | **optional.String**| Sorts in ascending order on the specified property. Prefix with &#x27;-&#x27; for descending order. | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[rbacSecurity](../README.md#rbacSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.netbackup+json;version=4.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

