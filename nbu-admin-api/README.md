# Go API client for swagger

The NetBackup Administration API provides access to administrative operations in NetBackup. 

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 4.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *https://&lt;masterserver&gt;:1556/netbackup*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DiscoveryApi* | [**AdminDiscoveryWorkloadsWorkloadTypeStartPost**](docs/DiscoveryApi.md#admindiscoveryworkloadsworkloadtypestartpost) | **Post** /admin/discovery/workloads/{workloadType}/start | 
*DiscoveryApi* | [**AdminDiscoveryWorkloadsWorkloadTypeStatusGet**](docs/DiscoveryApi.md#admindiscoveryworkloadsworkloadtypestatusget) | **Get** /admin/discovery/workloads/{workloadType}/status | 
*HostsApi* | [**AdminHostsHostIdProcessesGet**](docs/HostsApi.md#adminhostshostidprocessesget) | **Get** /admin/hosts/{hostId}/processes | 
*HostsApi* | [**AdminHostsHostIdServicesGet**](docs/HostsApi.md#adminhostshostidservicesget) | **Get** /admin/hosts/{hostId}/services | 
*HostsApi* | [**AdminHostsHostIdServicesServiceNameGet**](docs/HostsApi.md#adminhostshostidservicesservicenameget) | **Get** /admin/hosts/{hostId}/services/{serviceName} | 
*JobsApi* | [**AdminJobsGet**](docs/JobsApi.md#adminjobsget) | **Get** /admin/jobs | 
*JobsApi* | [**AdminJobsJobIdCancelPost**](docs/JobsApi.md#adminjobsjobidcancelpost) | **Post** /admin/jobs/{jobId}/cancel | 
*JobsApi* | [**AdminJobsJobIdDelete**](docs/JobsApi.md#adminjobsjobiddelete) | **Delete** /admin/jobs/{jobId} | 
*JobsApi* | [**AdminJobsJobIdFileListsGet**](docs/JobsApi.md#adminjobsjobidfilelistsget) | **Get** /admin/jobs/{jobId}/file-lists | 
*JobsApi* | [**AdminJobsJobIdGet**](docs/JobsApi.md#adminjobsjobidget) | **Get** /admin/jobs/{jobId} | 
*JobsApi* | [**AdminJobsJobIdProgressLogsGet**](docs/JobsApi.md#adminjobsjobidprogresslogsget) | **Get** /admin/jobs/{jobId}/progress-logs | 
*JobsApi* | [**AdminJobsJobIdRestartPost**](docs/JobsApi.md#adminjobsjobidrestartpost) | **Post** /admin/jobs/{jobId}/restart | 
*JobsApi* | [**AdminJobsJobIdResumePost**](docs/JobsApi.md#adminjobsjobidresumepost) | **Post** /admin/jobs/{jobId}/resume | 
*JobsApi* | [**AdminJobsJobIdSuspendPost**](docs/JobsApi.md#adminjobsjobidsuspendpost) | **Post** /admin/jobs/{jobId}/suspend | 
*JobsApi* | [**AdminJobsJobIdTryLogsAttemptGet**](docs/JobsApi.md#adminjobsjobidtrylogsattemptget) | **Get** /admin/jobs/{jobId}/try-logs/{attempt} | 
*JobsApi* | [**AdminJobsJobIdTryLogsGet**](docs/JobsApi.md#adminjobsjobidtrylogsget) | **Get** /admin/jobs/{jobId}/try-logs | 
*JobsApi* | [**AdminManualBackupPost**](docs/JobsApi.md#adminmanualbackuppost) | **Post** /admin/manual-backup | 

## Documentation For Models

 - [AdminManualbackupBody](docs/AdminManualbackupBody.md)
 - [AdmindiscoveryworkloadsworkloadTypestartData](docs/AdmindiscoveryworkloadsworkloadTypestartData.md)
 - [AdmindiscoveryworkloadsworkloadTypestartDataAttributes](docs/AdmindiscoveryworkloadsworkloadTypestartDataAttributes.md)
 - [AdminmanualbackupData](docs/AdminmanualbackupData.md)
 - [AdminmanualbackupDataAttributes](docs/AdminmanualbackupDataAttributes.md)
 - [DiscoveryRequest](docs/DiscoveryRequest.md)
 - [DiscoveryResponse](docs/DiscoveryResponse.md)
 - [DiscoveryResponseAttributes](docs/DiscoveryResponseAttributes.md)
 - [DiscoveryResponseData](docs/DiscoveryResponseData.md)
 - [DiscoveryResponseDataAttributes](docs/DiscoveryResponseDataAttributes.md)
 - [DiscoveryResponseTypes](docs/DiscoveryResponseTypes.md)
 - [DiscoveryStatus](docs/DiscoveryStatus.md)
 - [DiscoveryStatusPage](docs/DiscoveryStatusPage.md)
 - [DiscoveryStatusPageAttributes](docs/DiscoveryStatusPageAttributes.md)
 - [DiscoveryStatusPageAttributesAdditionalAttributes](docs/DiscoveryStatusPageAttributesAdditionalAttributes.md)
 - [DiscoveryStatusPageData](docs/DiscoveryStatusPageData.md)
 - [GetDiscoveryStatusResponse](docs/GetDiscoveryStatusResponse.md)
 - [GetIndividualTryResponse](docs/GetIndividualTryResponse.md)
 - [GetJobByIdResponse](docs/GetJobByIdResponse.md)
 - [GetJobDetailsResponse](docs/GetJobDetailsResponse.md)
 - [GetJobFileListResponse](docs/GetJobFileListResponse.md)
 - [GetJobFileListResponseData](docs/GetJobFileListResponseData.md)
 - [GetJobFileListResponseDataAttributes](docs/GetJobFileListResponseDataAttributes.md)
 - [GetJobFileListResponseDataLinks](docs/GetJobFileListResponseDataLinks.md)
 - [GetJobProgressEntryResponse](docs/GetJobProgressEntryResponse.md)
 - [GetJobProgressLogsResponse](docs/GetJobProgressLogsResponse.md)
 - [GetJobProgressLogsResponseAttributes](docs/GetJobProgressLogsResponseAttributes.md)
 - [GetJobProgressLogsResponseData](docs/GetJobProgressLogsResponseData.md)
 - [GetJobProgressLogsResponseMeta](docs/GetJobProgressLogsResponseMeta.md)
 - [GetJobProgressLogsResponseMetaPagination](docs/GetJobProgressLogsResponseMetaPagination.md)
 - [GetJobTryAttemptResponse](docs/GetJobTryAttemptResponse.md)
 - [GetJobTryLogsResponse](docs/GetJobTryLogsResponse.md)
 - [GetJobTryLogsResponseAttributes](docs/GetJobTryLogsResponseAttributes.md)
 - [GetJobTryLogsResponseAttributesTryLog](docs/GetJobTryLogsResponseAttributesTryLog.md)
 - [GetJobTryLogsResponseData](docs/GetJobTryLogsResponseData.md)
 - [GetJobTryLogsResponseLinks](docs/GetJobTryLogsResponseLinks.md)
 - [GetJobsResponse](docs/GetJobsResponse.md)
 - [GetJobsResponseLinks](docs/GetJobsResponseLinks.md)
 - [GetJobsResponseMeta](docs/GetJobsResponseMeta.md)
 - [GetJobsResponseMetaPagination](docs/GetJobsResponseMetaPagination.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [InlineResponse2002](docs/InlineResponse2002.md)
 - [InlineResponse2003](docs/InlineResponse2003.md)
 - [InlineResponse2004](docs/InlineResponse2004.md)
 - [InlineResponse2005](docs/InlineResponse2005.md)
 - [InlineResponse2006](docs/InlineResponse2006.md)
 - [InlineResponse2007](docs/InlineResponse2007.md)
 - [InlineResponse2008](docs/InlineResponse2008.md)
 - [InlineResponse2009](docs/InlineResponse2009.md)
 - [InlineResponse202](docs/InlineResponse202.md)
 - [InlineResponse2021](docs/InlineResponse2021.md)
 - [InlineResponse401](docs/InlineResponse401.md)
 - [JobData](docs/JobData.md)
 - [JobPatchRequest](docs/JobPatchRequest.md)
 - [JobPatchRequestData](docs/JobPatchRequestData.md)
 - [JobPatchRequestDataAttributes](docs/JobPatchRequestDataAttributes.md)
 - [JobPostRequest](docs/JobPostRequest.md)
 - [JobPostRequestData](docs/JobPostRequestData.md)
 - [JobPostRequestDataAttributes](docs/JobPostRequestDataAttributes.md)
 - [JobPostRequestDataAttributesJobData](docs/JobPostRequestDataAttributesJobData.md)
 - [JobPostRequestDataAttributesJobUpdateData](docs/JobPostRequestDataAttributesJobUpdateData.md)
 - [JobPostRequestDataAttributesJobUpdateDataTryLogList](docs/JobPostRequestDataAttributesJobUpdateDataTryLogList.md)
 - [JobPostRequestDataAttributesJobUpdateDataTryLogListData](docs/JobPostRequestDataAttributesJobUpdateDataTryLogListData.md)
 - [JobUpdateData](docs/JobUpdateData.md)
 - [ManualBackupRequest](docs/ManualBackupRequest.md)
 - [ManualBackupResponse](docs/ManualBackupResponse.md)
 - [ManualBackupResponseAttributes](docs/ManualBackupResponseAttributes.md)
 - [ManualBackupResponseData](docs/ManualBackupResponseData.md)
 - [ManualBackupResponseLinks](docs/ManualBackupResponseLinks.md)
 - [ManualBackupResponseLinksSelf](docs/ManualBackupResponseLinksSelf.md)
 - [PaginationValues](docs/PaginationValues.md)
 - [ProcessListResponse](docs/ProcessListResponse.md)
 - [ProcessListResponseAttributes](docs/ProcessListResponseAttributes.md)
 - [ProcessListResponseData](docs/ProcessListResponseData.md)
 - [ServiceDetails](docs/ServiceDetails.md)
 - [ServiceListResponse](docs/ServiceListResponse.md)
 - [ServiceListResponseAttributes](docs/ServiceListResponseAttributes.md)
 - [ServiceListResponseData](docs/ServiceListResponseData.md)
 - [ServiceResponse](docs/ServiceResponse.md)
 - [TryLogList](docs/TryLogList.md)
 - [TryLogLogEntry](docs/TryLogLogEntry.md)
 - [WorkloadName](docs/WorkloadName.md)
 - [WorkloadTypeStartBody](docs/WorkloadTypeStartBody.md)

## Documentation For Authorization

## rbacSecurity
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author


