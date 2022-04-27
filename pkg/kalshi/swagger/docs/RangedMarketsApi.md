# {{classname}}

All URIs are relative to *https://trading-api.kalshi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRangedMarketByTicker**](RangedMarketsApi.md#GetRangedMarketByTicker) | **Get** /ranged_markets_by_ticker/{ticker} | GetRangedMarketByTicker
[**UserGetRangedMarketPosition**](RangedMarketsApi.md#UserGetRangedMarketPosition) | **Get** /users/{user_id}/ranged_positions/{ranged_market_id} | UserGetRangedMarketPosition

# **GetRangedMarketByTicker**
> GetRangedMarketByTickerResponse GetRangedMarketByTicker(ctx, ticker)
GetRangedMarketByTicker

End-point for getting data about a ranged market by its ticker

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ticker** | **string**| Should be the ticker of the ranged market | 

### Return type

[**GetRangedMarketByTickerResponse**](GetRangedMarketByTickerResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserGetRangedMarketPosition**
> UserGetRangedMarketPositionResponse UserGetRangedMarketPosition(ctx, userId, rangedMarketId)
UserGetRangedMarketPosition

End-point for getting the market positions and additional data for the logged in user for all markets whose results linked by a single outcome. These markets share a parameter ranged_market_id, which indicates their relationship with each other.  The value for the user_id path parameter should match the user_id value returned on the response for the last login request (POST /log_in).  The value for the ranged_market_id path parameter should match the id value of the target market.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| Should be filled in with your user_id provided on log_in | 
  **rangedMarketId** | **string**| Should be filled with the id of the target ranged market | 

### Return type

[**UserGetRangedMarketPositionResponse**](UserGetRangedMarketPositionResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

