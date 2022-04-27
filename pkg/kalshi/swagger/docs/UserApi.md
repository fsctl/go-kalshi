# {{classname}}

All URIs are relative to *https://trading-api.kalshi.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserAddWatchlist**](UserApi.md#UserAddWatchlist) | **Put** /users/{user_id}/watchlist/{market_id} | UserAddWatchlist
[**UserBatchOrdersCancel**](UserApi.md#UserBatchOrdersCancel) | **Delete** /users/{user_id}/batch_orders | UserBatchOrdersCancel
[**UserBatchOrdersCreate**](UserApi.md#UserBatchOrdersCreate) | **Post** /users/{user_id}/batch_orders | UserBatchOrdersCreate
[**UserChangePassword**](UserApi.md#UserChangePassword) | **Put** /users/{user_id}/password | UserChangePassword
[**UserDeactivate**](UserApi.md#UserDeactivate) | **Delete** /users | UserDeactivate
[**UserGetBalance**](UserApi.md#UserGetBalance) | **Get** /users/{user_id}/balance | UserGetBalance
[**UserGetMarketPosition**](UserApi.md#UserGetMarketPosition) | **Get** /users/{user_id}/positions/{market_id} | UserGetMarketPosition
[**UserGetMarketPositions**](UserApi.md#UserGetMarketPositions) | **Get** /users/{user_id}/positions | UserGetMarketPositions
[**UserGetProfile**](UserApi.md#UserGetProfile) | **Get** /users/{user_id} | UserGetProfile
[**UserGetReferralInfo**](UserApi.md#UserGetReferralInfo) | **Get** /users/{user_id}/referrals | UserGetReferralInfo
[**UserGetWatchlist**](UserApi.md#UserGetWatchlist) | **Get** /users/{user_id}/watchlist | UserGetWatchlist
[**UserOrderCancel**](UserApi.md#UserOrderCancel) | **Delete** /users/{user_id}/orders/{order_id} | UserOrderCancel
[**UserOrderCreate**](UserApi.md#UserOrderCreate) | **Post** /users/{user_id}/orders | UserOrderCreate
[**UserOrderDecrease**](UserApi.md#UserOrderDecrease) | **Post** /users/{user_id}/orders/{order_id}/decrease | UserOrderDecrease
[**UserOrdersGet**](UserApi.md#UserOrdersGet) | **Get** /users/{user_id}/orders | UserOrdersGet
[**UserRemoveWatchlist**](UserApi.md#UserRemoveWatchlist) | **Delete** /users/{user_id}/watchlist/{market_id} | UserRemoveWatchlist
[**UserTradesGet**](UserApi.md#UserTradesGet) | **Get** /users/{user_id}/trades | UserTradesGet

# **UserAddWatchlist**
> UserAddWatchlist(ctx, userId, marketId)
UserAddWatchlist

End-point for adding a market to the logged in user's watchlist.  The value for the user_id path parameter should match the user_id value returned on the response for the last login request (POST /log_in).  The value for the market_id path parameter should match the id value of the market to be added.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| user_id should be filled with your user_id provided on log_in | 
  **marketId** | **string**| market_id should be filled with the id of the market to be added to the watchlist | 

### Return type

 (empty response body)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserBatchOrdersCancel**
> UserBatchOrdersCancelResponse UserBatchOrdersCancel(ctx, userId, optional)
UserBatchOrdersCancel

End-point for cancelling up to 40 orders at once. Available to members with advanced access only.  The value for the user_id path parameter should match the user_id value returned on the response for the last login request (POST /log_in).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| This parameter should be filled with your user_id provided on log_in | 
 **optional** | ***UserApiUserBatchOrdersCancelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiUserBatchOrdersCancelOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UserBatchOrdersCancelRequest**](UserBatchOrdersCancelRequest.md)| Orders cancel input data | 

### Return type

[**UserBatchOrdersCancelResponse**](UserBatchOrdersCancelResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserBatchOrdersCreate**
> UserBatchOrdersCreateResponse UserBatchOrdersCreate(ctx, userId, optional)
UserBatchOrdersCreate

End-point for submitting up to 20 orders in a batch. Available to members with advanced access only.  The value for the user_id path parameter should match the user_id value returned on the response for the last login request (POST /log_in).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| This parameter should be filled with your user_id provided on log_in | 
 **optional** | ***UserApiUserBatchOrdersCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiUserBatchOrdersCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UserBatchOrdersCreateRequest**](UserBatchOrdersCreateRequest.md)| Order create input data | 

### Return type

[**UserBatchOrdersCreateResponse**](UserBatchOrdersCreateResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserChangePassword**
> UserChangePassword(ctx, userId, optional)
UserChangePassword

End-point for updating logged-in user password.  The value for the user_id path parameter should match the user_id value returned on the response for the last login request (POST /log_in).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| This parameter should be filled with your user_id provided on log_in | 
 **optional** | ***UserApiUserChangePasswordOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiUserChangePasswordOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UserChangePasswordRequest**](UserChangePasswordRequest.md)| Change password input fields. | 

### Return type

 (empty response body)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserDeactivate**
> CreateUserResponse UserDeactivate(ctx, )
UserDeactivate

End-point for deactivating an user. A call to this end-point deactivates the current user and ends the current session.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CreateUserResponse**](CreateUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserGetBalance**
> UserGetBalanceResponse UserGetBalance(ctx, userId)
UserGetBalance

End-point for getting the balance of the logged in user.  The value for the user_id path parameter should match the user_id value returned on the response for the last login request (POST /log_in).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| Should be filled with your user_id provided on log_in | 

### Return type

[**UserGetBalanceResponse**](UserGetBalanceResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserGetMarketPosition**
> UserGetMarketPositionResponse UserGetMarketPosition(ctx, userId, marketId)
UserGetMarketPosition

End-point for getting the market positions for the logged in user, in a specific market.  The value for the user_id path parameter should match the user_id value returned on the response for the last login request (POST /log_in).  The value for the market_id path parameter should match the id value of the target market.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| Should be filled with your user_id provided on log_in | 
  **marketId** | **string**| Should be filled with the id of the target market | 

### Return type

[**UserGetMarketPositionResponse**](UserGetMarketPositionResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserGetMarketPositions**
> UserGetMarketPositionsResponse UserGetMarketPositions(ctx, userId)
UserGetMarketPositions

End-point for getting all market positions for the logged in user.  The value for the user_id path parameter should match the user_id value returned on the response for the last login request (POST /log_in).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| Should be filled with your user_id provided on log_in | 

### Return type

[**UserGetMarketPositionsResponse**](UserGetMarketPositionsResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserGetProfile**
> UserGetProfileResponse UserGetProfile(ctx, userId)
UserGetProfile

End-point for retrieving the logged in user's profile.  The value for the user_id path parameter should match the user_id value returned on the response for the last login request (POST /log_in).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| Should be filled with your user_id provided on log_in | 

### Return type

[**UserGetProfileResponse**](UserGetProfileResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserGetReferralInfo**
> UserGetReferralInfoResponse UserGetReferralInfo(ctx, userId)
UserGetReferralInfo

End-point for getting all information related to a member's referral status.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| Should be filled with your user_id provided on log_in | 

### Return type

[**UserGetReferralInfoResponse**](UserGetReferralInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserGetWatchlist**
> UserGetWatchlistResponse UserGetWatchlist(ctx, userId)
UserGetWatchlist

End-point for getting the market watchlist for the logged in user.  The value for the user_id path parameter should match the user_id value returned on the response for the last login request (POST /log_in).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| Should be filled with your user_id provided on log_in | 

### Return type

[**UserGetWatchlistResponse**](UserGetWatchlistResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserOrderCancel**
> UserOrderDecreaseResponse UserOrderCancel(ctx, userId, orderId)
UserOrderCancel

End-point for canceling orders.  The value for the user_id path parameter should match the user_id value returned on the response for the last login request (POST /log_in). The value for the order_id should match the id field of the order you want to decrease. Commonly delete end-points return 204 status with no body content on success. But we can't completely delete the order, as it may be partially filled already. So what the delete end-point does is just reducing the order completely zeroing the remaining resting contracts on it. The zeroed order is returned on the response payload, as a form of validation for the client.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| This parameter should be filled with your user_id provided on log_in | 
  **orderId** | **string**| This order_id should be filled with the id of the order to be decrease | 

### Return type

[**UserOrderDecreaseResponse**](UserOrderDecreaseResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserOrderCreate**
> UserOrderCreateResponse UserOrderCreate(ctx, userId, optional)
UserOrderCreate

End-point for submitting orders in a market.  The value for the user_id path parameter should match the user_id value returned on the response for the last login request (POST /log_in).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| This parameter should be filled with your user_id provided on log_in | 
 **optional** | ***UserApiUserOrderCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiUserOrderCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UserOrderCreateRequest**](UserOrderCreateRequest.md)| Order create input data | 

### Return type

[**UserOrderCreateResponse**](UserOrderCreateResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserOrderDecrease**
> UserOrderDecreaseResponse UserOrderDecrease(ctx, userId, orderId, optional)
UserOrderDecrease

End-point for decreasing the number of contracts on orders. This is the only kind of edit we support on orders.  The value for the user_id path parameter should match the user_id value returned on the response for the last login request (POST /log_in).  The value for the order_id should match the id field of the order you want to decrease.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| This parameter should be filled with your user_id provided on log_in | 
  **orderId** | **string**| This order_id should be filled with the id of the order to be decrease | 
 **optional** | ***UserApiUserOrderDecreaseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiUserOrderDecreaseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of UserOrderDecreaseRequest**](UserOrderDecreaseRequest.md)| Order data | 

### Return type

[**UserOrderDecreaseResponse**](UserOrderDecreaseResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserOrdersGet**
> UserOrdersGetResponse UserOrdersGet(ctx, userId, optional)
UserOrdersGet

End-point for getting all orders for the logged in user.  The value for the user_id path parameter should match the user_id value returned on the response for the last login request (POST /log_in).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| This parameter should be filled with your user_id provided on log_in | 
 **optional** | ***UserApiUserOrdersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiUserOrdersGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **marketId** | **optional.String**| Restricts the response to orders in a single market | 
 **isYes** | **optional.Bool**| Restricts the response to orders in a single direction (yes or no) | 
 **minPrice** | **optional.Int64**| Restricts the response to orders within a minimum price | 
 **maxPrice** | **optional.Int64**| Restricts the response to orders within a maximum price | 
 **minPlaceCount** | **optional.Int32**| Restricts the response to orders within a minimum place count | 
 **maxPlaceCount** | **optional.Int32**| Restricts the response to orders within a maximum place count | 
 **minInitialCount** | **optional.Int32**| Restricts the response to orders within a minimum initial count | 
 **maxInitialCount** | **optional.Int32**| Restricts the response to orders within a maximum initial count | 
 **minRemainingCount** | **optional.Int32**| Restricts the response to orders within a minimum remaining resting contracts count | 
 **maxRemainingCount** | **optional.Int32**| Restricts the response to orders within a maximum remaining resting contracts count | 
 **minDate** | **optional.Time**| Restricts the response to orders after a timestamp | 
 **maxDate** | **optional.Time**| Restricts the response to orders before a timestamp | 
 **getQueuePosition** | **optional.Bool**| If true, gets the queue placement for every resting order returned | 
 **status** | **optional.String**| Restricts the response to orders that have a certain status: resting, canceled, or executed | 
 **pageSize** | **optional.Int64**| Parameter to specify the number of results per page | 
 **pageNumber** | **optional.Int64**| Parameter to specify which page of the results should be retrieved | 

### Return type

[**UserOrdersGetResponse**](UserOrdersGetResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserRemoveWatchlist**
> UserRemoveWatchlist(ctx, userId, marketId)
UserRemoveWatchlist

End-point for removing a market from the logged in user's watchlist.  The value for the user_id path parameter should match the user_id value returned on the response for the last login request (POST /log_in).  The value for the market_id path parameter should match the id value of the market to be added.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| Should be filled with your user_id provided on log_in | 
  **marketId** | **string**| Should be filled with the id of the target market | 

### Return type

 (empty response body)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserTradesGet**
> UserTradesGetResponse UserTradesGet(ctx, userId, optional)
UserTradesGet

End-point for getting all trades for the logged in user.  The value for the user_id path parameter should match the user_id value returned on the response for the last login request (POST /log_in).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| This parameter should be filled with your user_id provided on log_in | 
 **optional** | ***UserApiUserTradesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiUserTradesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **marketId** | **optional.String**| Restricts the response to trades in a specific market. | 
 **orderId** | **optional.String**| Restricts the response to trades related to a specific order. | 
 **minPrice** | **optional.Int64**| Restricts the response to trades within a minimum price. | 
 **maxPrice** | **optional.Int64**| Restricts the response to trades within a maximum price. | 
 **minCount** | **optional.Int32**| Restricts the response to trades within a minimum contracts count. | 
 **maxCount** | **optional.Int32**| Restricts the response to trades within a maximum contracts count. | 
 **minDate** | **optional.Time**| Restricts the response to trades after a timestamp. | 
 **maxDate** | **optional.Time**| Restricts the response to trades before a timestamp. | 
 **pageSize** | **optional.Int64**| Parameter to specify the number of results per page | 
 **pageNumber** | **optional.Int64**| Parameter to specify which page of the results should be retrieved | 

### Return type

[**UserTradesGetResponse**](UserTradesGetResponse.md)

### Authorization

[cookie](../README.md#cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

