# {{classname}}

All URIs are relative to *https://trading-api.kalshi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSeriesByTicker**](SeriesApi.md#GetSeriesByTicker) | **Get** /series/{series_ticker} | GetSeriesByTicker

# **GetSeriesByTicker**
> GetSeriesByTickerResponse GetSeriesByTicker(ctx, seriesTicker)
GetSeriesByTicker

End-point for getting data about an event by its ticker

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **seriesTicker** | **string**| Should be the ticker of the series | 

### Return type

[**GetSeriesByTickerResponse**](GetSeriesByTickerResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

