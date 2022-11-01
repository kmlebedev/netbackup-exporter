# {{classname}}

All URIs are relative to *https://&lt;masterserver&gt;:1556/netbackup*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminJobsGet**](JobsApi.md#AdminJobsGet) | **Get** /admin/jobs | 
[**AdminJobsJobIdCancelPost**](JobsApi.md#AdminJobsJobIdCancelPost) | **Post** /admin/jobs/{jobId}/cancel | 
[**AdminJobsJobIdDelete**](JobsApi.md#AdminJobsJobIdDelete) | **Delete** /admin/jobs/{jobId} | 
[**AdminJobsJobIdFileListsGet**](JobsApi.md#AdminJobsJobIdFileListsGet) | **Get** /admin/jobs/{jobId}/file-lists | 
[**AdminJobsJobIdGet**](JobsApi.md#AdminJobsJobIdGet) | **Get** /admin/jobs/{jobId} | 
[**AdminJobsJobIdProgressLogsGet**](JobsApi.md#AdminJobsJobIdProgressLogsGet) | **Get** /admin/jobs/{jobId}/progress-logs | 
[**AdminJobsJobIdRestartPost**](JobsApi.md#AdminJobsJobIdRestartPost) | **Post** /admin/jobs/{jobId}/restart | 
[**AdminJobsJobIdResumePost**](JobsApi.md#AdminJobsJobIdResumePost) | **Post** /admin/jobs/{jobId}/resume | 
[**AdminJobsJobIdSuspendPost**](JobsApi.md#AdminJobsJobIdSuspendPost) | **Post** /admin/jobs/{jobId}/suspend | 
[**AdminJobsJobIdTryLogsAttemptGet**](JobsApi.md#AdminJobsJobIdTryLogsAttemptGet) | **Get** /admin/jobs/{jobId}/try-logs/{attempt} | 
[**AdminJobsJobIdTryLogsGet**](JobsApi.md#AdminJobsJobIdTryLogsGet) | **Get** /admin/jobs/{jobId}/try-logs | 
[**AdminManualBackupPost**](JobsApi.md#AdminManualBackupPost) | **Post** /admin/manual-backup | 

# **AdminJobsGet**
> InlineResponse2001 AdminJobsGet(ctx, optional)


Gets the list of jobs based on specified filters. If no filters are specified, information for 10 jobs sorted in descending order by job start time is retrieved.  To avoid sending large set of jobs data, API returns only subset of data (default size 10). To fetch additional results, take advantage of pagination using `page[Offset]` & `page[limit]`  which allows for subsequent requests to \"page\" through the rest of the results until the end is reached.  **Access Control**\\ Enforcement Type: [API-Level](http://sort.veritas.com/public/documents/nbu/8.3/windowsandunix/productguides/html/getting-started#user-content-object-level-enforcement)\\ Namespace: `|MANAGE|JOBS|`\\ Requires Operation: `|OPERATIONS|VIEW|` A given principal can also view the jobs that they initiate even if they are not granted the Jobs view operation. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***JobsApiAdminJobsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobsApiAdminJobsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| Filter criteria syntax is based on OData standards. Filtering is supported on the following attributes: |Name|Supported Operators|Description|Schema| |---|---|---|---| |jobId|eq ne lt gt le ge|The job identifier.|integer| |parentJobId|eq ne lt gt le ge|The parent job identifier.|integer| |clientName|eq ne contains, startswith, endswith|The name of the client to protect. If the client name contains a space or \&quot;%\&quot;, encode these characters as &#x27;%20&#x27; or &#x27;%25&#x27;, respectively.|string| |endTime|eq ne lt gt le ge|The time when the job finished.|date-time| |startTime|eq ne lt gt le ge|The time when the job started.|date-time| |state|eq ne|The state of the job. Possible values: QUEUED, ACTIVE, REQUEUED, DONE, SUSPENDED, INCOMPLETE|string| |status|eq ne lt gt le ge|The final job status code.|integer| |jobType|eq ne|The type of the job. Possible values: BACKUP, ARCHIVE, RESTORE, VERIFY, DUPLICATE, IMPORT, DBBACKUP, VAULT, LABEL, ERASE, TPREQ, TPCLEAN, TPFORMAT, VMPHYSINV, DQTS, DBRECOVER, MCONTENTS, IMAGEDELETE, GENERIC, REPLICATE, REPLICA_IMPORT, SNAPDUPE, SNAPREPLICATE, SNAPIMPORT, APP_STATE_CAPTURE, INDEXING, INDEXCLEANUP, SNAPINDEX, ACTIVATE_IR, DEACTIVATE_IR, REACTIVATE_IR, STOP_IR, INSTANT_RECOVERY, RMAN_CATALOG, LOCKVM, CLOUD_SNAPSHOT_REPLICATE|string| |policyName|contains, startswith, endswith|The name of the policy used by this job.|string| |policyType|eq ne|The type of the policy. Possible values: STANDARD, PROXY, NONSTANDARD, APOLLO_WBAK, OBACKUP, ANY, INFORMIX, SYBASE, SHAREPOINT, WINDOWS, NETWARE, BACKTRACK, AUSPEX_FASTBACK, WINDOWS_NT, OS2, SQL_SERVER, EXCHANGE, SAP, DB2, NDMP, FLASHBACKUP, SPLITMIRROR, AFS, DFS, DATASTORE, LOTUS_NOTES, NCR_TERADATA, VAX_VMS, HP3000_MPE, FBU_WINDOWS, VAULT, BE_SQL_SERVER, BE_EXCHANGE, MAC, DS, NBU_CATALOG, GENERIC, CMS_DB, PUREDISK_EXPORT, ENTERPRISE_VAULT, VMWARE, HYPERV, NBU_SEARCHSERVER, HYPERSCALE, BIGDATA, CLOUD, DEPLOYMENT, HYPERVISOR, UNIVERSAL_SHARE, NAS_DATA_PROTECTION|string| |restoreBackupIDs|eq, ne, contains, startswith, endswith|The backup images used for restore. If the client name contains a space or \&quot;%\&quot;, encode these characters as &#x27;%20&#x27; or &#x27;%25&#x27;, respectively.|string| |initiatorId|eq, ne, contains, startswith, endswith|The initiator of the job.|string| |assetID|eq|The asset identifier.|string| |destinationStorageUnitName|eq|The destination storage unit name.|string| |destinationMediaId|eq|The destination media id.|string|  | 
 **pageOffset** | **optional.Int32**| The job record number offset. | 
 **pageLimit** | **optional.Int32**| The number of records on one page after the offset. | 
 **sort** | **optional.String**| Sorts in ascending order on the specified property. Prefix with &#x27;-&#x27; for descending order. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[rbacSecurity](../README.md#rbacSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.netbackup+json;version=4.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminJobsJobIdCancelPost**
> AdminJobsJobIdCancelPost(ctx, jobId, optional)


Cancels the specified job.  **Access Control**\\ Enforcement Type: [API-Level](http://sort.veritas.com/public/documents/nbu/8.3/windowsandunix/productguides/html/getting-started#user-content-object-level-enforcement)\\ Namespace: `|MANAGE|JOBS|`\\ Requires Operation: `|OPERATIONS|UPDATE|` A given principal can also manage the jobs that they initiate even if they are not granted the Jobs update operation. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **int32**| The job identifier. | 
 **optional** | ***JobsApiAdminJobsJobIdCancelPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobsApiAdminJobsJobIdCancelPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xNetBackupAuditReason** | **optional.String**| Audit reason for the action being performed. This value must be URL-encoded. Maximum 512 bytes, in UTF-8 encoding. The remaining bytes are truncated.  | 

### Return type

 (empty response body)

### Authorization

[rbacSecurity](../README.md#rbacSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.netbackup+json;version=4.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminJobsJobIdDelete**
> AdminJobsJobIdDelete(ctx, jobId, optional)


Deletes the specified job.  **Access Control**\\ Enforcement Type: [API-Level](http://sort.veritas.com/public/documents/nbu/8.3/windowsandunix/productguides/html/getting-started#user-content-object-level-enforcement)\\ Namespace: `|MANAGE|JOBS|`\\ Requires Operation: `|OPERATIONS|DELETE|` A given principal can also manage the jobs that they initiate even if they are not granted the Jobs delete operation. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **int32**| The job identifier. | 
 **optional** | ***JobsApiAdminJobsJobIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobsApiAdminJobsJobIdDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xNetBackupAuditReason** | **optional.String**| Audit reason for the action being performed. This value must be URL-encoded. Maximum 512 bytes, in UTF-8 encoding. The remaining bytes are truncated.  | 

### Return type

 (empty response body)

### Authorization

[rbacSecurity](../README.md#rbacSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.netbackup+json;version=4.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminJobsJobIdFileListsGet**
> InlineResponse2006 AdminJobsJobIdFileListsGet(ctx, jobId)


Gets the file list for the specified job.  **Access Control**\\ Enforcement Type: [API-Level](http://sort.veritas.com/public/documents/nbu/8.3/windowsandunix/productguides/html/getting-started#user-content-object-level-enforcement)\\ Namespace: `|MANAGE|JOBS|`\\ Requires Operation: `|OPERATIONS|VIEW|` A given principal can also view the related file-lists for the jobs that they initiate even if they are not granted the Jobs view operation. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **int32**| The job identifier. | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[rbacSecurity](../README.md#rbacSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.netbackup+json;version=4.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminJobsJobIdGet**
> InlineResponse2002 AdminJobsJobIdGet(ctx, jobId)


Gets the job details for the specified job.  **Access Control**\\ Enforcement Type: [API-Level](http://sort.veritas.com/public/documents/nbu/8.3/windowsandunix/productguides/html/getting-started#user-content-object-level-enforcement)\\ Namespace: `|MANAGE|JOBS|`\\ Requires Operation: `|OPERATIONS|VIEW|` A given principal can also view the jobs that they initiate even if they are not granted the Jobs view operation. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **int32**| The job identifier. | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[rbacSecurity](../README.md#rbacSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.netbackup+json;version=4.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminJobsJobIdProgressLogsGet**
> InlineResponse2003 AdminJobsJobIdProgressLogsGet(ctx, jobId, optional)


Gets progress logs for the specified job.  **Access Control**\\ Enforcement Type: [API-Level](http://sort.veritas.com/public/documents/nbu/8.3/windowsandunix/productguides/html/getting-started#user-content-object-level-enforcement)\\ Namespace: `|MANAGE|JOBS|`\\ Requires Operation: `|OPERATIONS|VIEW|` A given principal can also view the related progress logs for the jobs that they initiate even if they are not granted the Jobs view operation. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **int32**| The job identifier. | 
 **optional** | ***JobsApiAdminJobsJobIdProgressLogsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobsApiAdminJobsJobIdProgressLogsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageLimit** | **optional.Int32**| The number of progress logs on one page. | [default to 10]
 **pageOffset** | **optional.Int32**| The progress log offset. | [default to 0]

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[rbacSecurity](../README.md#rbacSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.netbackup+json;version=4.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminJobsJobIdRestartPost**
> AdminJobsJobIdRestartPost(ctx, jobId, optional)


Restarts the specified job.  **Access Control**\\ Enforcement Type: [API-Level](http://sort.veritas.com/public/documents/nbu/8.3/windowsandunix/productguides/html/getting-started#user-content-object-level-enforcement)\\ Namespace: `|MANAGE|JOBS|`\\ Requires Operation: `|OPERATIONS|UPDATE|` A given principal can also manage the jobs that they initiate even if they are not granted the Jobs update operation. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **int32**| The job identifier. | 
 **optional** | ***JobsApiAdminJobsJobIdRestartPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobsApiAdminJobsJobIdRestartPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xNetBackupAuditReason** | **optional.String**| Audit reason for the action being performed. This value must be URL-encoded. Maximum 512 bytes, in UTF-8 encoding. The remaining bytes are truncated.  | 

### Return type

 (empty response body)

### Authorization

[rbacSecurity](../README.md#rbacSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.netbackup+json;version=4.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminJobsJobIdResumePost**
> AdminJobsJobIdResumePost(ctx, jobId, optional)


Resumes the specified job.  **Access Control**\\ Enforcement Type: [API-Level](http://sort.veritas.com/public/documents/nbu/8.3/windowsandunix/productguides/html/getting-started#user-content-object-level-enforcement)\\ Namespace: `|MANAGE|JOBS|`\\ Requires Operation: `|OPERATIONS|UPDATE|` A given principal can also manage the jobs that they initiate even if they are not granted the Jobs update operation. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **int32**| The job identifier. | 
 **optional** | ***JobsApiAdminJobsJobIdResumePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobsApiAdminJobsJobIdResumePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xNetBackupAuditReason** | **optional.String**| Audit reason for the action being performed. This value must be URL-encoded. Maximum 512 bytes, in UTF-8 encoding. The remaining bytes are truncated.  | 

### Return type

 (empty response body)

### Authorization

[rbacSecurity](../README.md#rbacSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.netbackup+json;version=4.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminJobsJobIdSuspendPost**
> AdminJobsJobIdSuspendPost(ctx, jobId, optional)


Suspends the specified job.  **Access Control**\\ Enforcement Type: [API-Level](http://sort.veritas.com/public/documents/nbu/8.3/windowsandunix/productguides/html/getting-started#user-content-object-level-enforcement)\\ Namespace: `|MANAGE|JOBS|`\\ Requires Operation: `|OPERATIONS|UPDATE|` A given principal can also manage the jobs that they initiate even if they are not granted the Jobs update operation. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **int32**| The job identifier. | 
 **optional** | ***JobsApiAdminJobsJobIdSuspendPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobsApiAdminJobsJobIdSuspendPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xNetBackupAuditReason** | **optional.String**| Audit reason for the action being performed. This value must be URL-encoded. Maximum 512 bytes, in UTF-8 encoding. The remaining bytes are truncated.  | 

### Return type

 (empty response body)

### Authorization

[rbacSecurity](../README.md#rbacSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.netbackup+json;version=4.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminJobsJobIdTryLogsAttemptGet**
> InlineResponse2005 AdminJobsJobIdTryLogsAttemptGet(ctx, jobId, attempt)


Gets the try-logs for the specified job ID and the specified try number.  **Access Control**\\ Enforcement Type: [API-Level](http://sort.veritas.com/public/documents/nbu/8.3/windowsandunix/productguides/html/getting-started#user-content-object-level-enforcement)\\ Namespace: `|MANAGE|JOBS|`\\ Requires Operation: `|OPERATIONS|VIEW|` A given principal can also view the related try-logs for the jobs that they initiate even if they are not granted the Jobs view operation. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **int32**| The job identifier. | 
  **attempt** | **int32**| The try attempt for the specified job. | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[rbacSecurity](../README.md#rbacSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.netbackup+json;version=4.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminJobsJobIdTryLogsGet**
> InlineResponse2004 AdminJobsJobIdTryLogsGet(ctx, jobId)


Gets logs for the specified job.  **Access Control**\\ Enforcement Type: [API-Level](http://sort.veritas.com/public/documents/nbu/8.3/windowsandunix/productguides/html/getting-started#user-content-object-level-enforcement)\\ Namespace: `|MANAGE|JOBS|`\\ Requires Operation: `|OPERATIONS|VIEW|` A given principal can also view the related try-logs for the jobs that they initiate even if they are not granted the Jobs view operation. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **int32**| The job identifier. | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[rbacSecurity](../README.md#rbacSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.netbackup+json;version=4.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminManualBackupPost**
> InlineResponse202 AdminManualBackupPost(ctx, body)


Start a manual backup job.  **Access Control**\\ Enforcement Type: [API-Level](http://sort.veritas.com/public/documents/nbu/8.3/windowsandunix/productguides/html/getting-started#user-content-api-level-enforcement)\\ Namespace: `|PROTECTION|POLICIES|`\\ Requires Operation: `|OPERATIONS|PROTECTION|POLICIES|MANUAL-BACKUP|` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AdminManualbackupBody**](AdminManualbackupBody.md)| The details of the manual backup to start. | 

### Return type

[**InlineResponse202**](inline_response_202.md)

### Authorization

[rbacSecurity](../README.md#rbacSecurity)

### HTTP request headers

 - **Content-Type**: application/vnd.netbackup+json;version=4.0
 - **Accept**: application/vnd.netbackup+json;version=4.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

