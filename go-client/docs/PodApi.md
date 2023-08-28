# \PodApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PodCloseHandler**](PodApi.md#PodCloseHandler) | **Post** /v1/pod/close | Close pod
[**PodCreateHandler**](PodApi.md#PodCreateHandler) | **Post** /v1/pod/new | Create pod
[**PodDeleteHandler**](PodApi.md#PodDeleteHandler) | **Delete** /v1/pod/delete | Delete pod
[**PodForkFromReferenceHandler**](PodApi.md#PodForkFromReferenceHandler) | **Post** /v1/pod/fork-from-reference | Fork a pod from sharing reference
[**PodForkHandler**](PodApi.md#PodForkHandler) | **Post** /v1/pod/fork | Fork a pod
[**PodListHandler**](PodApi.md#PodListHandler) | **Get** /v1/pod/ls | List pods
[**PodOpenAsyncHandler**](PodApi.md#PodOpenAsyncHandler) | **Post** /v1/pod/open-async | Open pod
[**PodOpenHandler**](PodApi.md#PodOpenHandler) | **Post** /v1/pod/open | Open pod
[**PodPresent**](PodApi.md#PodPresent) | **Get** /v1/pod/present | Is pod present
[**PodReceiveHandler**](PodApi.md#PodReceiveHandler) | **Get** /v1/pod/receive | Receive shared pod
[**PodReceiveInfoHandler**](PodApi.md#PodReceiveInfoHandler) | **Get** /v1/pod/receiveinfo | Receive shared pod info
[**PodShareHandler**](PodApi.md#PodShareHandler) | **Post** /v1/pod/share | Share pod
[**PodStatHandler**](PodApi.md#PodStatHandler) | **Get** /v1/pod/stat | Stats for pod
[**PodSyncAsyncHandler**](PodApi.md#PodSyncAsyncHandler) | **Post** /v1/pod/sync-async | Sync pod asynchronously
[**PodSyncHandler**](PodApi.md#PodSyncHandler) | **Post** /v1/pod/sync | Sync pod


# **PodCloseHandler**
> ApiResponse PodCloseHandler(ctx, podRequest, cookie)
Close pod

PodCloseHandler is the api handler to close an open pod

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **podRequest** | [**ApiPodNameRequest**](ApiPodNameRequest.md)| pod name | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PodCreateHandler**
> ApiResponse PodCreateHandler(ctx, podRequest, cookie)
Create pod

PodCreateHandler is the api handler to create a new pod

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **podRequest** | [**ApiPodNameRequest**](ApiPodNameRequest.md)| pod name and user password | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PodDeleteHandler**
> ApiResponse PodDeleteHandler(ctx, podRequest, cookie)
Delete pod

PodDeleteHandler is the api handler to delete a pod

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **podRequest** | [**ApiPodNameRequest**](ApiPodNameRequest.md)| pod name and user password | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PodForkFromReferenceHandler**
> ApiResponse PodForkFromReferenceHandler(ctx, podRequest, cookie)
Fork a pod from sharing reference

PodForkFromReferenceHandler is the api handler to fork a pod from a given sharing reference

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **podRequest** | [**ApiPodForkFromReferenceRequest**](ApiPodForkFromReferenceRequest.md)| pod name and user password | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PodForkHandler**
> ApiResponse PodForkHandler(ctx, podRequest, cookie)
Fork a pod

PodForkHandler is the api handler to fork a pod

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **podRequest** | [**ApiPodForkRequest**](ApiPodForkRequest.md)| pod name and user password | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PodListHandler**
> ApiPodListResponse PodListHandler(ctx, cookie)
List pods

PodListHandler is the api handler to list all pods

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiPodListResponse**](api.PodListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PodOpenAsyncHandler**
> ApiResponse PodOpenAsyncHandler(ctx, podRequest, cookie)
Open pod

PodOpenAsyncHandler is the api handler to open pod asynchronously

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **podRequest** | [**ApiPodNameRequest**](ApiPodNameRequest.md)| pod name and user password | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PodOpenHandler**
> ApiResponse PodOpenHandler(ctx, podRequest, cookie)
Open pod

PodOpenHandler is the api handler to open pod

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **podRequest** | [**ApiPodNameRequest**](ApiPodNameRequest.md)| pod name and user password | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PodPresent**
> ApiResponse PodPresent(ctx, podName, cookie)
Is pod present

PodPresentHandler is the api handler to check if a pod is present

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **podName** | **string**| pod name | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PodReceiveHandler**
> ApiResponse PodReceiveHandler(ctx, sharingRef, cookie, optional)
Receive shared pod

PodReceiveHandler is the api handler to receive shared pod from shared reference

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sharingRef** | **string**| pod sharing reference | 
  **cookie** | **string**| cookie parameter | 
 **optional** | ***PodApiPodReceiveHandlerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PodApiPodReceiveHandlerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sharedPodName** | **optional.String**| pod name to be saved as | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PodReceiveInfoHandler**
> PodShareInfo PodReceiveInfoHandler(ctx, sharingRef, cookie)
Receive shared pod info

PodReceiveInfoHandler is the api handler to receive shared pod info from shared reference

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sharingRef** | **string**| pod sharing reference | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**PodShareInfo**](pod.ShareInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PodShareHandler**
> ApiPodSharingReference PodShareHandler(ctx, podRequest, cookie)
Share pod

PodShareHandler is the api handler to share a pod to the public

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **podRequest** | [**CommonPodShareRequest**](CommonPodShareRequest.md)| pod name and user password | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiPodSharingReference**](api.PodSharingReference.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PodStatHandler**
> ApiPodStatResponse PodStatHandler(ctx, podName, cookie)
Stats for pod

PodStatHandler is the api handler get information about a pod

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **podName** | **string**| pod name | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiPodStatResponse**](api.PodStatResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PodSyncAsyncHandler**
> ApiResponse PodSyncAsyncHandler(ctx, podRequest, cookie)
Sync pod asynchronously

PodSyncAsyncHandler is the api handler to sync a pod's content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **podRequest** | [**ApiPodNameRequest**](ApiPodNameRequest.md)| pod name | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PodSyncHandler**
> ApiResponse PodSyncHandler(ctx, podRequest, cookie)
Sync pod

PodSyncHandler is the api handler to sync a pod's content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **podRequest** | [**ApiPodNameRequest**](ApiPodNameRequest.md)| pod name | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

