# {{classname}}

All URIs are relative to *https://trading-api.kalshi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserGetPortfolioHistory**](PortfolioApi.md#UserGetPortfolioHistory) | **Get** /users/{user_id}/portfolio/history | UserGetPortfolioHistory

# **UserGetPortfolioHistory**
> UserGetPortfolioHistoryResponse UserGetPortfolioHistory(ctx, userId, optional)
UserGetPortfolioHistory

End-point for getting the logged in user's portfolio historical track.  The value for the user_id path parameter should match the user_id value returned on the response for the last login request (POST /log_in).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| This parameter should be filled with your user_id provided on log_in | 
 **optional** | ***PortfolioApiUserGetPortfolioHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiUserGetPortfolioHistoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UserGetPortfolioHistoryRequest**](UserGetPortfolioHistoryRequest.md)| Order create input data | 

### Return type

[**UserGetPortfolioHistoryResponse**](UserGetPortfolioHistoryResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

