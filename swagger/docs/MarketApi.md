# {{classname}}

All URIs are relative to *https://trading-api.kalshi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetActiveMarkets**](MarketApi.md#GetActiveMarkets) | **Get** /active_markets | GetActiveMarkets
[**GetCandlestickMarketHistory**](MarketApi.md#GetCandlestickMarketHistory) | **Get** /markets/{market_id}/candlestick | GetCandlestickMarketHistory
[**GetCandlestickMarketHistoryCached**](MarketApi.md#GetCandlestickMarketHistoryCached) | **Get** /cached/markets/{market_id}/candlestick | GetCandlestickMarketHistoryCached
[**GetMarket**](MarketApi.md#GetMarket) | **Get** /markets/{market_id} | GetMarket
[**GetMarketByTicker**](MarketApi.md#GetMarketByTicker) | **Get** /markets_by_ticker/{ticker_name} | GetMarketByTicker
[**GetMarketByTickerCached**](MarketApi.md#GetMarketByTickerCached) | **Get** /cached/markets_by_ticker/{ticker_name} | GetMarketByTickerCached
[**GetMarketCached**](MarketApi.md#GetMarketCached) | **Get** /cached/markets/{market_id} | GetMarketCached
[**GetMarketHistory**](MarketApi.md#GetMarketHistory) | **Get** /markets/{market_id}/stats_history | GetMarketHistory
[**GetMarketHistoryCached**](MarketApi.md#GetMarketHistoryCached) | **Get** /cached/markets/{market_id}/stats_history | GetMarketHistoryCached
[**GetMarketOrderBook**](MarketApi.md#GetMarketOrderBook) | **Get** /markets/{market_id}/order_book | GetMarketOrderBook
[**GetMarketOrderBookCached**](MarketApi.md#GetMarketOrderBookCached) | **Get** /cached/markets/{market_id}/order_book | GetMarketOrderBookCached
[**GetMarkets**](MarketApi.md#GetMarkets) | **Get** /markets | GetMarkets
[**GetMarketsCached**](MarketApi.md#GetMarketsCached) | **Get** /cached/markets | GetMarketsCached

# **GetActiveMarkets**
> GetActiveMarketsResponse GetActiveMarkets(ctx, optional)
GetActiveMarkets

End-point for getting highly active markets on the exchange. Currently, gets 3 kinds of activity: Markets opening within the time window provided, markets closing within the time window provided, and markets that have had large price movements or large volumes within the time window provided.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MarketApiGetActiveMarketsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MarketApiGetActiveMarketsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int64**| The maximum number of markets returned, this is capped at 20 | 
 **minDate** | **optional.Int64**| The lower bound on the date searched through when looking for activity | 
 **maxDate** | **optional.Int64**| The upper bound on the date searched through when looking for activity | 

### Return type

[**GetActiveMarketsResponse**](GetActiveMarketsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCandlestickMarketHistory**
> GetCandlestickMarketHistoryResponse GetCandlestickMarketHistory(ctx, marketId, optional)
GetCandlestickMarketHistory

End-point for getting open, high, low, close (OHLC) and other data for candlestick plots. See the response body for full information on what is returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **marketId** | **string**| Should be filled with the id of the target market | 
 **optional** | ***MarketApiGetCandlestickMarketHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MarketApiGetCandlestickMarketHistoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lastSeenTs** | **optional.Int64**| If provided, restricts history to trades starting from lastSeenTs | 
 **numBuckets** | **optional.Int32**| If provided, this field represents the number of buckets used to divide the market history data. Please provide an integer between 1 and 7,200 (inclusive). Defaults to 1,400. We aggregate data in the buckets. | 

### Return type

[**GetCandlestickMarketHistoryResponse**](GetCandlestickMarketHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCandlestickMarketHistoryCached**
> GetCandlestickMarketHistoryResponse GetCandlestickMarketHistoryCached(ctx, marketId, optional)
GetCandlestickMarketHistoryCached

End-point for getting open, high, low, close (OHLC) and other data for candlestick plots. See the response body for full information on what is returned. Data is cached and so is slightly lagged.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **marketId** | **string**| Should be filled with the id of the target market | 
 **optional** | ***MarketApiGetCandlestickMarketHistoryCachedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MarketApiGetCandlestickMarketHistoryCachedOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lastSeenTs** | **optional.Int64**| If provided, restricts history to trades starting from lastSeenTs | 
 **numBuckets** | **optional.Int32**| If provided, this field represents the number of buckets used to divide the market history data. Please provide an integer between 1 and 7,200 (inclusive). Defaults to 1,400. We aggregate data in the buckets. | 

### Return type

[**GetCandlestickMarketHistoryResponse**](GetCandlestickMarketHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMarket**
> UserGetMarketResponse GetMarket(ctx, marketId)
GetMarket

End-point for getting data about a specific market.  The value for the market_id path parameter should match the id value of the target market.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **marketId** | **string**| Should be filled with the id of the target market | 

### Return type

[**UserGetMarketResponse**](UserGetMarketResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMarketByTicker**
> UserGetMarketResponse GetMarketByTicker(ctx, tickerName)
GetMarketByTicker

End-point for getting data about a specific market based on its ticker.  The value for the ticker_name path parameter should match the ticker_name value of the target market.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tickerName** | **string**| Should be filled with the ticker name of the target market | 

### Return type

[**UserGetMarketResponse**](UserGetMarketResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMarketByTickerCached**
> UserGetMarketResponse GetMarketByTickerCached(ctx, tickerName)
GetMarketByTickerCached

End-point for getting data about a specific market with data that is cached and so slightly lagged.  The value for the ticker_name path parameter should match the ticker_name value of the target market.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tickerName** | **string**| Should be filled with the ticker name of the target market | 

### Return type

[**UserGetMarketResponse**](UserGetMarketResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMarketCached**
> UserGetMarketResponse GetMarketCached(ctx, marketId)
GetMarketCached

End-point for getting data about a specific market with data that is cached and so slightly lagged.  The value for the market_id path parameter should match the id value of the target market.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **marketId** | **string**| Should be filled with the id of the target market | 

### Return type

[**UserGetMarketResponse**](UserGetMarketResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMarketHistory**
> GetMarketHistoryResponse GetMarketHistory(ctx, marketId, optional)
GetMarketHistory

End-point for getting the statistics history for a market.  The value for the market_id path parameter should match the id value of the target market. The last_seen_ts parameter is optional, and will restrict statistics to those after provided timestamp. The last_seen_ts is inclusive, which means a market history point at last_seen_ts will be returned

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **marketId** | **string**| Should be filled with the id of the target market | 
 **optional** | ***MarketApiGetMarketHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MarketApiGetMarketHistoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lastSeenTs** | **optional.Int64**| If provided, restricts history to trades starting from lastSeenTs | 
 **numBuckets** | **optional.Int32**| If provided, this field represents the number of buckets used to group the market price history data. Please provide an integer between 1 and 7,200 (inclusive). The higher this value, the more points will be returned from this endpoint. Note that the number of points returned may not be equal to the number of buckets because some buckets may not contain any points. Defaults to 1,400. We pick one representative point from each bucket (namely the last point in the bucket) | 

### Return type

[**GetMarketHistoryResponse**](GetMarketHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMarketHistoryCached**
> GetMarketHistoryResponse GetMarketHistoryCached(ctx, marketId, optional)
GetMarketHistoryCached

End-point for getting the statistics history for a market with data that is cached and so slightly lagged.  The value for the market_id path parameter should match the id value of the target market. The last_seen_ts parameter is optional, and will restrict statistics to those after provided timestamp. The last_seen_ts is inclusive, which means a market history point at last_seen_ts will be returned

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **marketId** | **string**| Should be filled with the id of the target market | 
 **optional** | ***MarketApiGetMarketHistoryCachedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MarketApiGetMarketHistoryCachedOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lastSeenTs** | **optional.Int64**| If provided, restricts history to trades starting from lastSeenTs | 
 **numBuckets** | **optional.Int32**| If provided, this field represents the number of buckets used to group the market price history data. Please provide an integer between 1 and 7,200 (inclusive). The higher this value, the more points will be returned from this endpoint. Note that the number of points returned may not be equal to the number of buckets because some buckets may not contain any points. Defaults to 1,400. We pick one representative point from each bucket (namely the last point in the bucket) | 

### Return type

[**GetMarketHistoryResponse**](GetMarketHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMarketOrderBook**
> GetMarketOrderBookResponse GetMarketOrderBook(ctx, marketId)
GetMarketOrderBook

End-point for getting the orderbook for a market.  The value for the market_id path parameter should match the id value of the target market.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **marketId** | **string**| Should be filled with the id of the target market | 

### Return type

[**GetMarketOrderBookResponse**](GetMarketOrderBookResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMarketOrderBookCached**
> GetMarketOrderBookResponse GetMarketOrderBookCached(ctx, marketId)
GetMarketOrderBookCached

End-point for getting the orderbook for a market with data that is cached and so slightly lagged.  The value for the market_id path parameter should match the id value of the target market.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **marketId** | **string**| Should be filled with the id of the target market | 

### Return type

[**GetMarketOrderBookResponse**](GetMarketOrderBookResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMarkets**
> UserGetMarketsResponse GetMarkets(ctx, )
GetMarkets

End-point for listing / discovering markets on Kalshi.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UserGetMarketsResponse**](UserGetMarketsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMarketsCached**
> UserGetMarketsResponse GetMarketsCached(ctx, )
GetMarketsCached

End-point for listing / discovering markets on Kalshi with data that is cached and so slightly lagged.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UserGetMarketsResponse**](UserGetMarketsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

