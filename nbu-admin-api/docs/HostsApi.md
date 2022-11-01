# {{classname}}

All URIs are relative to *https://&lt;masterserver&gt;:1556/netbackup*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminHostsHostIdProcessesGet**](HostsApi.md#AdminHostsHostIdProcessesGet) | **Get** /admin/hosts/{hostId}/processes | 
[**AdminHostsHostIdServicesGet**](HostsApi.md#AdminHostsHostIdServicesGet) | **Get** /admin/hosts/{hostId}/services | 
[**AdminHostsHostIdServicesServiceNameGet**](HostsApi.md#AdminHostsHostIdServicesServiceNameGet) | **Get** /admin/hosts/{hostId}/services/{serviceName} | 

# **AdminHostsHostIdProcessesGet**
> InlineResponse200 AdminHostsHostIdProcessesGet(ctx, hostId, optional)


Get the list of NetBackup processes running on the specified host.  **Access Control**\\ Enforcement Type: [API-Level](http://sort.veritas.com/public/documents/nbu/8.3/windowsandunix/productguides/html/getting-started#user-content-object-level-enforcement)\\ Namespace: `|MANAGE|HOSTS|PROCESSES|`\\ Requires Operation: `|OPERATIONS|VIEW|` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostId** | **string**| The unique ID (UUID) of the host. See the &#x27;uuid&#x27; attribute of the response of GET /config/hosts API. | 
 **optional** | ***HostsApiAdminHostsHostIdProcessesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HostsApiAdminHostsHostIdProcessesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| Filter criteria syntax is based on OData standards. Filtering is supported on the following attributes: |Name|Supported Operators|Description|Schema| |---|---|---|---| |processName|eq|The process name.|string|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[rbacSecurity](../README.md#rbacSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.netbackup+json;version=4.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminHostsHostIdServicesGet**
> InlineResponse2007 AdminHostsHostIdServicesGet(ctx, hostId)


Get the list of NetBackup services on the specified host.  **Access Control**\\ Enforcement Type: [API-Level](http://sort.veritas.com/public/documents/nbu/8.3/windowsandunix/productguides/html/getting-started#user-content-api-level-enforcement)\\ Namespace: `|MANAGE|HOSTS|PROCESSES|`\\ Requires Operation: `|OPERATIONS|VIEW|` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostId** | **string**| The unique ID (UUID) of the host. See the &#x27;uuid&#x27; attribute of the response of GET /config/hosts API. | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[rbacSecurity](../README.md#rbacSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.netbackup+json;version=4.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminHostsHostIdServicesServiceNameGet**
> InlineResponse2008 AdminHostsHostIdServicesServiceNameGet(ctx, hostId, serviceName)


Get the details of the given NetBackup service on the specified host.  **Access Control**\\ Enforcement Type: [API-Level](http://sort.veritas.com/public/documents/nbu/8.3/windowsandunix/productguides/html/getting-started#user-content-api-level-enforcement)\\ Namespace: `|MANAGE|HOSTS|PROCESSES|`\\ Requires Operation: `|OPERATIONS|VIEW|` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostId** | **string**| The unique ID (UUID) of the host. See the &#x27;uuid&#x27; attribute of the response of GET /config/hosts API. | 
  **serviceName** | **string**| The service (daemon) name. | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[rbacSecurity](../README.md#rbacSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.netbackup+json;version=4.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

