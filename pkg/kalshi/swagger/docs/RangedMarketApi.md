# {{classname}}

All URIs are relative to *https://trading-api.kalshi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRangedMarket**](RangedMarketApi.md#GetRangedMarket) | **Get** /ranged_markets/{ranged_market_id} | GetRangedMarket

# **GetRangedMarket**
> GetRangedMarketResponse GetRangedMarket(ctx, rangedMarketId)
GetRangedMarket

End-point for getting data about a ranged market

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rangedMarketId** | **string**| Should be filled in with a ranged market id | 

### Return type

[**GetRangedMarketResponse**](GetRangedMarketResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

