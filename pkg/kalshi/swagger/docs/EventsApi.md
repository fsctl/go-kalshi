# {{classname}}

All URIs are relative to *https://trading-api.kalshi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEventByTickerCached**](EventsApi.md#GetEventByTickerCached) | **Get** /events/{ticker} | GetEventByTickerCached

# **GetEventByTickerCached**
> GetEventByTickerResponse GetEventByTickerCached(ctx, ticker)
GetEventByTickerCached

End-point for getting data about an event by its ticker with data that is cached and so slightly lagged.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ticker** | **string**| Should be the ticker of the event | 

### Return type

[**GetEventByTickerResponse**](GetEventByTickerResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

