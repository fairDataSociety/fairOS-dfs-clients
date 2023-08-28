# \PublicApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PublicDirGet**](PublicApi.md#PublicDirGet) | **Get** /public-dir | List directory content
[**PublicFileGet**](PublicApi.md#PublicFileGet) | **Get** /public-file | download file from a shared pod
[**PublicKvGet**](PublicApi.md#PublicKvGet) | **Get** /public-kv | get key from public pod
[**PublicRefFileGet**](PublicApi.md#PublicRefFileGet) | **Get** /public/{ref}/{file} | download file from a shared pod


# **PublicDirGet**
> ApiListFileResponse PublicDirGet(ctx, sharingRef, dirPath)
List directory content

PublicPodGetDirHandler is the api handler to list content of a directory from a public pod

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sharingRef** | **string**| pod sharing reference | 
  **dirPath** | **string**| dir location in the pod | 

### Return type

[**ApiListFileResponse**](api.ListFileResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublicFileGet**
> []int32 PublicFileGet(ctx, sharingRef, filePath)
download file from a shared pod

PodReceiveInfoHandler is the api handler to download file from a shared pod

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sharingRef** | **string**| pod sharing reference | 
  **filePath** | **string**| file location in the pod | 

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublicKvGet**
> ApiKvResponse PublicKvGet(ctx, sharingRef, tableName, key)
get key from public pod

PublicPodKVEntryGetHandler is the api handler to get key from key value store from a public pod

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sharingRef** | **string**| pod sharing reference | 
  **tableName** | **string**| table name | 
  **key** | **string**| key to look up | 

### Return type

[**ApiKvResponse**](api.KVResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublicRefFileGet**
> []int32 PublicRefFileGet(ctx, ref, file)
download file from a shared pod

PublicPodFilePathHandler is the api handler to download file from a shared pod

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ref** | **string**| pod sharing reference | 
  **file** | **string**| file location in the pod | 

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

