# {{classname}}

All URIs are relative to *https://trading-api.kalshi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserImmediateFunding**](UsersApi.md#GetUserImmediateFunding) | **Get** /users/{user_id}/immediate_funding | GetUserImmediateFunding
[**GetUserWithdrawalAvailableBalance**](UsersApi.md#GetUserWithdrawalAvailableBalance) | **Get** /users/{user_id}/withdrawals/available | GetUserWithdrawalAvailableBalance

# **GetUserImmediateFunding**
> GetUserImmediateFundingResponse GetUserImmediateFunding(ctx, userId, optional)
GetUserImmediateFunding

End-point for getting immediate funding info for a member.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| This parameter should be filled with your user_id provided on log_in | 
 **optional** | ***UsersApiGetUserImmediateFundingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiGetUserImmediateFundingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **depositAmountCents** | **optional.Int64**| Pass this parameter if you&#x27;d like to see how much of a deposit will be funded by immediate funding. If you don&#x27;t need this information, pass 0 cents. | 

### Return type

[**GetUserImmediateFundingResponse**](GetUserImmediateFundingResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserWithdrawalAvailableBalance**
> UserWithdrawalAvailableBalanceResponse GetUserWithdrawalAvailableBalance(ctx, userId)
GetUserWithdrawalAvailableBalance

End-point for getting how much money a member is elgible to withdraw

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| This parameter should be filled with your user_id provided on log_in | 

### Return type

[**UserWithdrawalAvailableBalanceResponse**](UserWithdrawalAvailableBalanceResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

