# \UserApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserDeleteV2**](UserApi.md#UserDeleteV2) | **Delete** /v2/user/delete | Delete user for ENS based authentication
[**UserIsloggedin**](UserApi.md#UserIsloggedin) | **Get** /v1/user/isloggedin | Is user logged-in
[**UserLoginV2**](UserApi.md#UserLoginV2) | **Post** /v2/user/login | Login User
[**UserLogout**](UserApi.md#UserLogout) | **Post** /v1/user/logout | Logout
[**UserPresentV2**](UserApi.md#UserPresentV2) | **Get** /v2/user/present | Check if user is present
[**UserSignupV2**](UserApi.md#UserSignupV2) | **Post** /v2/user/signup | Register New User
[**UserStat**](UserApi.md#UserStat) | **Get** /v1/user/stat | User stat
[**V1UserDeletePost**](UserApi.md#V1UserDeletePost) | **Post** /v1/user/delete | 
[**V1UserExportPost**](UserApi.md#V1UserExportPost) | **Post** /v1/user/export | 
[**V1UserLoginPost**](UserApi.md#V1UserLoginPost) | **Post** /v1/user/login | 
[**V1UserPresentGet**](UserApi.md#V1UserPresentGet) | **Get** /v1/user/present | 
[**V1UserSignupPost**](UserApi.md#V1UserSignupPost) | **Post** /v1/user/signup | 


# **UserDeleteV2**
> ApiResponse UserDeleteV2(ctx, userDeleteRequest, cookie)
Delete user for ENS based authentication

deletes user info from swarm

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userDeleteRequest** | [**ApiUserDeleteRequest**](ApiUserDeleteRequest.md)| user delete request | 
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserIsloggedin**
> ApiLoginStatus UserIsloggedin(ctx, userName)
Is user logged-in

Check if the given user is logged-in

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userName** | **string**| username | 

### Return type

[**ApiLoginStatus**](api.LoginStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserLoginV2**
> ApiUserLoginResponse UserLoginV2(ctx, userRequest)
Login User

login user with the new ENS based authentication

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userRequest** | [**CommonUserLoginRequest**](CommonUserLoginRequest.md)| username | 

### Return type

[**ApiUserLoginResponse**](api.UserLoginResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserLogout**
> ApiResponse UserLogout(ctx, cookie)
Logout

logs-out user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cookie** | **string**| cookie parameter | 

### Return type

[**ApiResponse**](api.response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserPresentV2**
> ApiPresentResponse UserPresentV2(ctx, userName)
Check if user is present

checks if the new user is present in the new ENS based authentication

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userName** | **string**| username | 

### Return type

[**ApiPresentResponse**](api.PresentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserSignupV2**
> ApiUserSignupResponse UserSignupV2(ctx, userRequest)
Register New User

registers new user with the new ENS based authentication

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userRequest** | [**CommonUserSignupRequest**](CommonUserSignupRequest.md)| username | 

### Return type

[**ApiUserSignupResponse**](api.UserSignupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserStat**
> UserStat UserStat(ctx, cookie)
User stat

show user stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cookie** | **string**| cookie parameter | 

### Return type

[**UserStat**](user.Stat.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UserDeletePost**
> V1UserDeletePost(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UserExportPost**
> V1UserExportPost(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UserLoginPost**
> V1UserLoginPost(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UserPresentGet**
> V1UserPresentGet(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UserSignupPost**
> V1UserSignupPost(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

