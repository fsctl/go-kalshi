# {{classname}}

All URIs are relative to *https://trading-api.kalshi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEventsCached**](DefaultApi.md#GetEventsCached) | **Get** /events/ | GetEventsCached
[**GetSeriesList**](DefaultApi.md#GetSeriesList) | **Get** /series/ | GetSeriesList
[**GetSeriesListCached**](DefaultApi.md#GetSeriesListCached) | **Get** /cached/series | GetSeriesListCached
[**GetTrades**](DefaultApi.md#GetTrades) | **Get** /trades | GetTrades

# **GetEventsCached**
> GetEventsResponse GetEventsCached(ctx, )
GetEventsCached

End-point for getting data about all events with data that is cached and so slightly lagged.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetEventsResponse**](GetEventsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSeriesList**
> GetSeriesListResponse GetSeriesList(ctx, )
GetSeriesList

End-point for getting data about all series

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetSeriesListResponse**](GetSeriesListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSeriesListCached**
> GetSeriesListResponse GetSeriesListCached(ctx, )
GetSeriesListCached

End-point for getting data about all series. Endpoint is cached so it is slightly lagged.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetSeriesListResponse**](GetSeriesListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTrades**
> TradesGetResponse GetTrades(ctx, optional)
GetTrades

End-point for getting all trades for all markets.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiGetTradesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiGetTradesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tradesDate** | **optional.String**| Restricts the response to trades during a certain day (give trades_date in ET with format: YYYY-MM-DD). Dates returned will be UTC | 
 **pageSize** | **optional.Int64**| Parameter to specify the number of results per page | 
 **pageNumber** | **optional.Int64**| Parameter to specify which page of the results should be retrieved | 
 **marketId** | **optional.String**| Parameter to specify a specific marketId to get trades from | 

### Return type

[**TradesGetResponse**](TradesGetResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

