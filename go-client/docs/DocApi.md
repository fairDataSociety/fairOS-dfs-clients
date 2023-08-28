# \DocApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DocCount**](DocApi.md#DocCount) | **Post** /v1/doc/count | Count number of document in a table
[**DocCreate**](DocApi.md#DocCreate) | **Post** /v1/doc/new | Create in doc table
[**DocDelete**](DocApi.md#DocDelete) | **Delete** /v1/doc/delete | Delete a doc table
[**DocEntryDel**](DocApi.md#DocEntryDel) | **Delete** /v1/doc/entry/del | Delete a document from a document datastore
[**DocEntryGet**](DocApi.md#DocEntryGet) | **Get** /v1/doc/entry/get | Get a document from a document datastore
[**DocEntryPut**](DocApi.md#DocEntryPut) | **Post** /v1/doc/entry/put | Add a record in document datastore
[**DocFind**](DocApi.md#DocFind) | **Get** /v1/doc/find | Get rows from a given doc datastore
[**DocIndexjson**](DocApi.md#DocIndexjson) | **Post** /v1/doc/indexjson | Index a json file that is present in a pod, in to the given document database
[**DocLoadjson**](DocApi.md#DocLoadjson) | **Post** /v1/doc/loadjson | Load json file from local file system
[**DocLs**](DocApi.md#DocLs) | **Get** /v1/doc/ls | List all doc table
[**DocOpen**](DocApi.md#DocOpen) | **Post** /v1/doc/open | Open a doc table


# **DocCount**
> CollectionTableKeyCount DocCount(ctx, docRequest, cookie)
Count number of document in a table

DocCountHandler is the api handler to count the number of documents in a given document database

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **docRequest** | [**ApiDocCountRequest**](ApiDocCountRequest.md)| doc table info | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**CollectionTableKeyCount**](collection.TableKeyCount.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DocCreate**
> ApiResponse DocCreate(ctx, docRequest, cookie)
Create in doc table

DocCreateHandler is the api handler to create a new document database

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **docRequest** | [**ApiDocRequest**](ApiDocRequest.md)| doc table info. si or simple index is a comma separated list of keys and their types. eg: &#39;first_name&#x3D;string,age&#x3D;number&#39;. valid index types can be &#39;string&#39;, &#39;number&#39;, &#39;map&#39;, &#39;list&#39;. default index is &#39;id&#39; and it should be of type string | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DocDelete**
> ApiResponse DocDelete(ctx, docRequest, cookie)
Delete a doc table

DocDeleteHandler is the api handler to delete the given document database

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **docRequest** | [**ApiSimpleDocRequest**](ApiSimpleDocRequest.md)| doc table info | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DocEntryDel**
> ApiResponse DocEntryDel(ctx, cookie, optional)
Delete a document from a document datastore

DocEntryDelHandler is the api handler to delete a document from a document datastore

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cookie** | **string**| cookie parameter | 
 **optional** | ***DocApiDocEntryDelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DocApiDocEntryDelOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **optional.String**|  | 
 **podName** | **optional.String**|  | 
 **tableName** | **optional.String**|  | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DocEntryGet**
> ApiDocGet DocEntryGet(ctx, podName, tableName, id, cookie)
Get a document from a document datastore

DocEntryGetHandler is the api handler to get a document from a document datastore

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **podName** | **string**| pod name | 
  **tableName** | **string**| table name | 
  **id** | **string**| id to search for | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiDocGet**](api.DocGet.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DocEntryPut**
> ApiResponse DocEntryPut(ctx, cookie, optional)
Add a record in document datastore

DocEntryPutHandler is the api handler add a document in to a document datastore

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cookie** | **string**| cookie parameter | 
 **optional** | ***DocApiDocEntryPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DocApiDocEntryPutOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **doc** | **optional.String**|  | 
 **podName** | **optional.String**|  | 
 **tableName** | **optional.String**|  | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DocFind**
> ApiDocFind DocFind(ctx, podName, tableName, expr, cookie, optional)
Get rows from a given doc datastore

DocFindHandler is the api handler to select rows from a given document datastore

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **podName** | **string**| pod name | 
  **tableName** | **string**| table name | 
  **expr** | **string**| expression to search for. allowed operators in expr are &#x3D;, &gt;, &#x3D;&gt;, &lt;&#x3D;, &lt;. eg: &#39;first_name&#x3D;&gt;John&#39;, &#39;first_name&#x3D;&gt;J.&#39;, &#39;first_name&#x3D;&gt;.&#39;, &#39;age&#x3D;&gt;30&#39;, &#39;age&lt;&#x3D;30&#39;. if index is string, expr supports regex. we do not have support for multiple conditions in expr yet | 
  **cookie** | **string**| cookie parameter | 
 **optional** | ***DocApiDocFindOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DocApiDocFindOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **limit** | **optional.String**| number od documents | 

### Return type

[**ApiDocFind**](api.DocFind.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DocIndexjson**
> ApiResponse DocIndexjson(ctx, indexRequest, cookie)
Index a json file that is present in a pod, in to the given document database

DocIndexJsonHandler is the api handler to index a json file that is present in a pod, in to the given document database

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **indexRequest** | [**ApiDocIndexRequest**](ApiDocIndexRequest.md)| index request | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DocLoadjson**
> ApiResponse DocLoadjson(ctx, podName, tableName, json, cookie)
Load json file from local file system

DocLoadJsonHandler is the api handler that indexes a json file that is present in the local file system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **podName** | **string**| pod name | 
  **tableName** | **string**| table name | 
  **json** | ***os.File**| json to index | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DocLs**
> ApiDocumentDbs DocLs(ctx, podName, cookie)
List all doc table

DocListHandler is the api handler which lists all the document database in a pod

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **podName** | **string**| pod name | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiDocumentDbs**](api.DocumentDBs.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DocOpen**
> ApiDocumentDbs DocOpen(ctx, docRequest, cookie)
Open a doc table

DocOpenHandler is the api handler to open a document database

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **docRequest** | [**ApiDocRequest**](ApiDocRequest.md)| doc table info | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiDocumentDbs**](api.DocumentDBs.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

