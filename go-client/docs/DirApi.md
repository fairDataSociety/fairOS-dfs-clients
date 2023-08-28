# \DirApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DirectoryLsHandler**](DirApi.md#DirectoryLsHandler) | **Get** /v1/dir/ls | List directory
[**DirectoryMkdirHandler**](DirApi.md#DirectoryMkdirHandler) | **Post** /v1/dir/mkdir | Create directory
[**DirectoryModeHandler**](DirApi.md#DirectoryModeHandler) | **Post** /v1/dir/chmod | change mode of a directory
[**DirectoryPresentHandler**](DirApi.md#DirectoryPresentHandler) | **Get** /v1/dir/present | Is directory present
[**DirectoryRenameHandler**](DirApi.md#DirectoryRenameHandler) | **Post** /v1/dir/rename | Rename directory
[**DirectoryRmdirHandler**](DirApi.md#DirectoryRmdirHandler) | **Delete** /v1/dir/rmdir | Remove directory
[**DirectoryStatHandler**](DirApi.md#DirectoryStatHandler) | **Get** /v1/dir/stat | Directory stat


# **DirectoryLsHandler**
> ApiListFileResponse DirectoryLsHandler(ctx, podName, dirPath, cookie)
List directory

DirectoryLsHandler is the api handler for listing the contents of a directory.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **podName** | **string**| pod name | 
  **dirPath** | **string**| dir path | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiListFileResponse**](api.ListFileResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DirectoryMkdirHandler**
> ApiResponse DirectoryMkdirHandler(ctx, dirRequest, cookie)
Create directory

DirectoryMkdirHandler is the api handler to create a new directory.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dirRequest** | [**ApiDirRequest**](ApiDirRequest.md)| pod name and dir path | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DirectoryModeHandler**
> ApiResponse DirectoryModeHandler(ctx, dirRequest, cookie)
change mode of a directory

DirectoryModeHandler is the api handler to change mode of a directory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dirRequest** | [**ApiDirModeRequest**](ApiDirModeRequest.md)| dir mode request | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DirectoryPresentHandler**
> ApiDirPresentResponse DirectoryPresentHandler(ctx, podName, dirPath, cookie)
Is directory present

DirectoryPresentHandler is the api handler which says if a directory is present or not

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **podName** | **string**| pod name | 
  **dirPath** | **string**| dir path | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiDirPresentResponse**](api.DirPresentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DirectoryRenameHandler**
> ApiResponse DirectoryRenameHandler(ctx, dirRequest, cookie)
Rename directory

DirectoryRenameHandler is the api handler to rename a directory.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dirRequest** | [**CommonRenameRequest**](CommonRenameRequest.md)| old name and new path | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DirectoryRmdirHandler**
> ApiResponse DirectoryRmdirHandler(ctx, dirRequest, cookie)
Remove directory

DirectoryRmdirHandler is the api handler to remove a directory.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dirRequest** | [**ApiDirRequest**](ApiDirRequest.md)| pod name and dir path | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DirectoryStatHandler**
> DirStats DirectoryStatHandler(ctx, podName, dirPath, cookie)
Directory stat

DirectoryStatHandler is the api handler which gives the information about a directory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **podName** | **string**| pod name | 
  **dirPath** | **string**| dir path | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**DirStats**](dir.Stats.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

