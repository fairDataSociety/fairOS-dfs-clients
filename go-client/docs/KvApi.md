# \KvApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KvCount**](KvApi.md#KvCount) | **Post** /v1/kv/count | Count rows in a key value table
[**KvCreateHandler**](KvApi.md#KvCreateHandler) | **Post** /v1/kv/new | Create a key value table
[**KvDel**](KvApi.md#KvDel) | **Delete** /v1/kv/entry/del | Delete key-value from the kv table
[**KvDelete**](KvApi.md#KvDelete) | **Delete** /v1/kv/delete | Delete a key value table
[**KvExport**](KvApi.md#KvExport) | **Post** /v1/kv/export | Export from a particular key with the given prefix
[**KvGet**](KvApi.md#KvGet) | **Get** /v1/kv/entry/get | get value from the kv table
[**KvGetData**](KvApi.md#KvGetData) | **Get** /v1/kv/entry/get-data | get value from the kv table
[**KvGetNext**](KvApi.md#KvGetNext) | **Post** /v1/kv/seek/next | Get next value from last seek in kv table
[**KvLoadcsv**](KvApi.md#KvLoadcsv) | **Post** /v1/kv/loadcsv | Upload a csv file in kv table
[**KvLs**](KvApi.md#KvLs) | **Get** /v1/kv/ls | List all key value tables
[**KvOpen**](KvApi.md#KvOpen) | **Post** /v1/kv/open | Open a key value table
[**KvPresentHandler**](KvApi.md#KvPresentHandler) | **Get** /v1/kv/entry/present | Check if a value exists in the kv table
[**KvPut**](KvApi.md#KvPut) | **Post** /v1/kv/entry/put | put key and value in the kv table
[**KvSeek**](KvApi.md#KvSeek) | **Post** /v1/kv/seek | Seek in kv table


# **KvCount**
> CollectionTableKeyCount KvCount(ctx, kvTableRequest, cookie)
Count rows in a key value table

KVCountHandler is the api handler to count the number of rows in a key value table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **kvTableRequest** | [**ApiKvTableRequest**](ApiKvTableRequest.md)| kv table request | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**CollectionTableKeyCount**](collection.TableKeyCount.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **KvCreateHandler**
> ApiResponse KvCreateHandler(ctx, kvTableRequest, cookie)
Create a key value table

KVCreateHandler is the api handler to create a key value table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **kvTableRequest** | [**ApiKvTableRequest**](ApiKvTableRequest.md)| kv table request | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **KvDel**
> ApiKvResponseRaw KvDel(ctx, deleteRequest, cookie)
Delete key-value from the kv table

KVDelHandler is the api handler to delete a key and value from the kv table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteRequest** | [**ApiKvEntryDeleteRequest**](ApiKvEntryDeleteRequest.md)| delete request | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiKvResponseRaw**](api.KVResponseRaw.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **KvDelete**
> ApiResponse KvDelete(ctx, kvTableRequest, cookie)
Delete a key value table

KVDeleteHandler is the api handler to delete a key value table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **kvTableRequest** | [**ApiKvTableRequest**](ApiKvTableRequest.md)| kv table request | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **KvExport**
> []map[string][]map[string]string KvExport(ctx, exportRequest, cookie)
Export from a particular key with the given prefix

KVExportHandler is the api handler to export from a particular key with the given prefix

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **exportRequest** | [**ApiKvExportRequest**](ApiKvExportRequest.md)| kv export info | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**[]map[string][]map[string]string**](map.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **KvGet**
> ApiKvResponse KvGet(ctx, podName, tableName, key, cookie)
get value from the kv table

KVGetHandler is the api handler to get a value from the kv table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **podName** | **string**| pod name | 
  **tableName** | **string**| table name | 
  **key** | **string**| key | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiKvResponse**](api.KVResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **KvGetData**
> ApiKvResponseRaw KvGetData(ctx, podName, tableName, key, cookie, optional)
get value from the kv table

KVGetDataHandler is the api handler to get raw value from the kv table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **podName** | **string**| pod name | 
  **tableName** | **string**| table name | 
  **key** | **string**| key | 
  **cookie** | **string**| cookie parameter | 
 **optional** | ***KvApiKvGetDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KvApiKvGetDataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **format** | **optional.String**| format of the value | 

### Return type

[**ApiKvResponseRaw**](api.KVResponseRaw.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **KvGetNext**
> ApiKvResponse KvGetNext(ctx, podName, tableName, cookie)
Get next value from last seek in kv table

KVGetNextHandler is the api handler to get the key and value from the current seek position

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **podName** | **string**| pod name | 
  **tableName** | **string**| table name | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiKvResponse**](api.KVResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **KvLoadcsv**
> ApiResponse KvLoadcsv(ctx, podName, tableName, csv, cookie, optional)
Upload a csv file in kv table

KVLoadCSVHandler is the api handler to load a csv file as key and value in a KV table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **podName** | **string**| pod name | 
  **tableName** | **string**| table name | 
  **csv** | ***os.File**| file to upload | 
  **cookie** | **string**| cookie parameter | 
 **optional** | ***KvApiKvLoadcsvOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KvApiKvLoadcsvOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **memory** | **optional.String**| keep in memory | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **KvLs**
> ApiCollections KvLs(ctx, podName, cookie)
List all key value tables

KVListHandler is the api handler to list all the key value tables in a pod

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **podName** | **string**| pod name | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiCollections**](api.Collections.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **KvOpen**
> ApiResponse KvOpen(ctx, kvTableRequest, cookie)
Open a key value table

KVOpenHandler is the api handler to open a key value table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **kvTableRequest** | [**ApiKvTableRequest**](ApiKvTableRequest.md)| kv table request | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **KvPresentHandler**
> ApiResponse KvPresentHandler(ctx, podName, tableName, key, cookie)
Check if a value exists in the kv table

KVPresentHandler is the api handler to check if a value exists in the kv table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **podName** | **string**| pod name | 
  **tableName** | **string**| table name | 
  **key** | **string**| key | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **KvPut**
> ApiResponse KvPut(ctx, kvEntry, cookie)
put key and value in the kv table

KVPutHandler is the api handler to put a key-value  in the kv table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **kvEntry** | [**ApiKvEntryRequest**](ApiKvEntryRequest.md)| kv entry | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **KvSeek**
> ApiResponse KvSeek(ctx, exportRequest, cookie)
Seek in kv table

KVSeekHandler is the api handler to seek to a particular key with the given prefix

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **exportRequest** | [**ApiKvExportRequest**](ApiKvExportRequest.md)| kv seek info | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

