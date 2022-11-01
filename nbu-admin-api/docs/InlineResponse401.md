# InlineResponse401

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | **int64** | The NetBackup error code. | [default to null]
**ErrorMessage** | **string** | The NetBackup error code description. | [default to null]
**AttributeErrors** | [***interface{}**](interface{}.md) | Maps the attributes with the invalid values to an error that describes why the value is invalid, to provide more context to the nature of the error.  | [optional] [default to null]
**ErrorDetails** | **[]string** | A list that contains additional information about the error. | [optional] [default to null]
**FileUploadErrors** | [**[]interface{}**](interface{}.md) | If an API is supporting content-type multipart/form-data, then this holds a list of objects that contains names of failed files and associated error messages.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

