# \FileApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FileDeleteHandler**](FileApi.md#FileDeleteHandler) | **Delete** /v1/file/delete | Delete a file
[**FileDownloadHandler**](FileApi.md#FileDownloadHandler) | **Get** /v1/file/download | Download a file
[**FileDownloadHandlerPost**](FileApi.md#FileDownloadHandlerPost) | **Post** /v1/file/download | Download a file
[**FileModeHandler**](FileApi.md#FileModeHandler) | **Post** /v1/file/chmod | chmod a file
[**FileReceiveHandler**](FileApi.md#FileReceiveHandler) | **Get** /v1/file/receive | Receive a file
[**FileReceiveInfoHandler**](FileApi.md#FileReceiveInfoHandler) | **Get** /v1/file/receiveinfo | Receive a file info
[**FileRenameHandler**](FileApi.md#FileRenameHandler) | **Post** /v1/file/rename | Info of a file
[**FileShareHandler**](FileApi.md#FileShareHandler) | **Post** /v1/file/share | Share a file
[**FileStatHandler**](FileApi.md#FileStatHandler) | **Get** /v1/file/stat | Info of a file
[**FileStatusHandler**](FileApi.md#FileStatusHandler) | **Get** /v1/file/status | Sync status of a file
[**FileUpdateHandler**](FileApi.md#FileUpdateHandler) | **Post** /v1/file/update | Update a file
[**FileUploadHandler**](FileApi.md#FileUploadHandler) | **Post** /v1/file/upload | Upload a file


# **FileDeleteHandler**
> ApiResponse FileDeleteHandler(ctx, fileDeleteRequest, cookie)
Delete a file

FileReceiveHandler is the api handler to delete a file from a given pod

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileDeleteRequest** | [**ApiFileDeleteRequest**](ApiFileDeleteRequest.md)| pod name and file path | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileDownloadHandler**
> []int32 FileDownloadHandler(ctx, podName, filePath, cookie)
Download a file

FileDownloadHandlerGet is the api handler to download a file from a given pod

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **podName** | **string**| pod name | 
  **filePath** | **string**| file path | 
  **cookie** | **string**| cookie parameter | 

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileDownloadHandlerPost**
> []int32 FileDownloadHandlerPost(ctx, podName, filePath, cookie)
Download a file

FileDownloadHandlerPost is the api handler to download a file from a given pod

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **podName** | **string**| pod name | 
  **filePath** | **string**| file path | 
  **cookie** | **string**| cookie parameter | 

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileModeHandler**
> ApiResponse FileModeHandler(ctx, fileRequest, cookie)
chmod a file

FileModeHandler is the api handler to change mode of a file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileRequest** | [**ApiFileModeRequest**](ApiFileModeRequest.md)| file mode request | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileReceiveHandler**
> ApiFileSharingReference FileReceiveHandler(ctx, podName, sharingRef, dirPath, cookie)
Receive a file

FileReceiveHandler is the api handler to receive a file in a given pod

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **podName** | **string**| pod name | 
  **sharingRef** | **string**| sharing reference | 
  **dirPath** | **string**| file location | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiFileSharingReference**](api.FileSharingReference.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileReceiveInfoHandler**
> UserReceiveFileInfo FileReceiveInfoHandler(ctx, sharingRef, cookie)
Receive a file info

FileReceiveInfoHandler is the api handler to receive a file info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sharingRef** | **string**| sharing reference | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**UserReceiveFileInfo**](user.ReceiveFileInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileRenameHandler**
> ApiResponse FileRenameHandler(ctx, renameRequest, cookie)
Info of a file

FileRenameHandler is the api handler to get the information of a file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **renameRequest** | [**CommonRenameRequest**](CommonRenameRequest.md)| old name &amp; new name | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileShareHandler**
> ApiFileSharingReference FileShareHandler(ctx, fileShareRequest, cookie)
Share a file

FileShareHandler is the api handler to share a file from a given pod

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileShareRequest** | [**ApiFileShareRequest**](ApiFileShareRequest.md)| file share request params | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiFileSharingReference**](api.FileSharingReference.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileStatHandler**
> FileStats FileStatHandler(ctx, podName, filePath, cookie)
Info of a file

FileStatHandler is the api handler to get the information of a file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **podName** | **string**| pod name | 
  **filePath** | **string**| file path | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**FileStats**](file.Stats.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileStatusHandler**
> []ApiStatusResponse FileStatusHandler(ctx, podName, filePath, cookie)
Sync status of a file

FileStatusHandler is the api handler to check sync status of a file from a given pod

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **podName** | **string**| pod name | 
  **filePath** | **string**| file path | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**[]ApiStatusResponse**](api.StatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileUpdateHandler**
> ApiResponse FileUpdateHandler(ctx, podName, filePath, file, offset, cookie)
Update a file

FileUpdateHandler is the api handler to update a file from a given offset

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **podName** | **string**| pod name | 
  **filePath** | **string**| location | 
  **file** | ***os.File**| file content to update | 
  **offset** | **string**| file offset | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileUploadHandler**
> ApiResponse FileUploadHandler(ctx, podName, dirPath, blockSize, files, cookie, optional)
Upload a file

FileUploadHandler is the api handler to upload a file from a local file system to the dfs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **podName** | **string**| pod name | 
  **dirPath** | **string**| location | 
  **blockSize** | **string**| block size to break the file | 
  **files** | ***os.File**| file to upload | 
  **cookie** | **string**| cookie parameter | 
 **optional** | ***FileApiFileUploadHandlerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FileApiFileUploadHandlerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **fairOSDfsCompression** | **optional.String**| cookie parameter | 
 **overwrite** | **optional.String**| overwrite the file if already exists | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

