# {{classname}}

All URIs are relative to *https://trading-api.kalshi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeSubscription**](AccountApi.md#ChangeSubscription) | **Patch** /users/{user_id}/subscribe | ChangeSubscription
[**GetNotificationPreferences**](AccountApi.md#GetNotificationPreferences) | **Get** /users/{user_id}/notifications/preferences | GetNotificationPreferences
[**NotificationMarkRead**](AccountApi.md#NotificationMarkRead) | **Put** /users/{user_id}/notifications/{notification_id}/read | NotificationMarkRead
[**UserGetAccountHistory**](AccountApi.md#UserGetAccountHistory) | **Get** /users/{user_id}/account/history | UserGetAccountHistory
[**UserGetNotifications**](AccountApi.md#UserGetNotifications) | **Get** /users/{user_id}/notifications | UserGetNotifications
[**UserGetProfitsAndLosses**](AccountApi.md#UserGetProfitsAndLosses) | **Get** /users/{user_id}/account/pnl | UserGetProfitsAndLosses

# **ChangeSubscription**
> ChangeSubscriptionResponse ChangeSubscription(ctx, userId, optional)
ChangeSubscription

End-point for changing e-mail subscription mode for the current user.  This end-point is very useful for users that have a large volume of orders and don't want to be email notified whenever an order is submitted / edited / canceled or matches.  This is specially useful for Market Makers.  The value for the user_id path parameter should match the user_id value returned on the response for the last login request (POST /log_in).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| Should be filled with your user_id provided on log_in | 
 **optional** | ***AccountApiChangeSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountApiChangeSubscriptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ChangeSubscriptionRequest**](ChangeSubscriptionRequest.md)| Change subscription data | 

### Return type

[**ChangeSubscriptionResponse**](ChangeSubscriptionResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNotificationPreferences**
> GetNotificationPreferencesResponse GetNotificationPreferences(ctx, userId)
GetNotificationPreferences

End-point for getting e-mail subscription mode for the current user.  The value for the user_id path parameter should match the user_id value returned on the response for the last login request (POST /log_in).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| This parameter should be filled with your user_id provided on log_in | 

### Return type

[**GetNotificationPreferencesResponse**](GetNotificationPreferencesResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NotificationMarkRead**
> NotificationMarkRead(ctx, userId, notificationId)
NotificationMarkRead

End-point for marking a notification as read.  The value for the user_id path parameter should match the user_id value returned on the response for the last login request (POST /log_in).  The value for the notification_id path parameter should match the notification_id value of the notification to be marked as read.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| user_id should be filled with your user_id provided on log_in | 
  **notificationId** | **string**| notification_id should be filled with the id of the notification to be mark as read | 

### Return type

 (empty response body)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserGetAccountHistory**
> UserGetAccountHistoryResponse UserGetAccountHistory(ctx, userId, optional)
UserGetAccountHistory

End-point for getting the logged in user's important past actions and events related to the user's positions.  This contains entries for user's explicit actions but also for market events.  There will be entries for:  submitting, editing / canceling orders requesting deposits and withdrawals trade execution (order matching) market settlements on markets where you have a position  The value for the user_id path parameter should match the user_id value returned on the response for the last login request (POST /log_in).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| This parameter should be filled with your user_id provided on log_in | 
 **optional** | ***AccountApiUserGetAccountHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountApiUserGetAccountHistoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deposits** | **optional.Bool**| If true the response should include deposit entries | 
 **withdrawals** | **optional.Bool**| If true the response should include withdrawal entries | 
 **orders** | **optional.Bool**| If true the response should include order entries | 
 **settlements** | **optional.Bool**| If true the response should include settlement entries | 
 **trades** | **optional.Bool**| If true the response should include trade entries | 
 **credits** | **optional.Bool**| If true the response should include credit entries | 
 **limit** | **optional.Int64**| Restricts the response to a return the first \&quot;limit\&quot; amount of acct history items. Note if you specify a limit, you cannot specify a PageSize or PageNumber | 
 **pageSize** | **optional.Int64**| Parameter to specify the number of results per page. | 
 **pageNumber** | **optional.Int64**| Parameter to specify which page of the results should be retrieved | 

### Return type

[**UserGetAccountHistoryResponse**](UserGetAccountHistoryResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserGetNotifications**
> UserGetNotificationsResponse UserGetNotifications(ctx, userId, pageSize, pageNumber)
UserGetNotifications

End-point for getting notifications for the current logged in user.  The value for the user_id path parameter should match the user_id value returned on the response for the last login request (POST /log_in).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| This parameter should be filled with your user_id provided on log_in | 
  **pageSize** | **int64**| Number of results per page | 
  **pageNumber** | **int64**| Page of the results | 

### Return type

[**UserGetNotificationsResponse**](UserGetNotificationsResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserGetProfitsAndLosses**
> int64 UserGetProfitsAndLosses(ctx, userId, optional)
UserGetProfitsAndLosses

This end point returns profits and losses between two dates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| This parameter should be filled with your user_id provided on log_in | 
 **optional** | ***AccountApiUserGetProfitsAndLossesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountApiUserGetProfitsAndLossesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTs** | **optional.Int64**| Start time of pnl period, represented as the number of seconds since Unix epoch | 
 **endTs** | **optional.Int64**| End time of pnl period, represented as the number of seconds since Unix epoch | 

### Return type

**int64**

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

