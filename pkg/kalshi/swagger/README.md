# Go API client for swagger

This documentation describes Kalshi's rest API for market makers  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *https://trading-api.kalshi.com/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountApi* | [**ChangeSubscription**](docs/AccountApi.md#changesubscription) | **Patch** /users/{user_id}/subscribe | ChangeSubscription
*AccountApi* | [**GetNotificationPreferences**](docs/AccountApi.md#getnotificationpreferences) | **Get** /users/{user_id}/notifications/preferences | GetNotificationPreferences
*AccountApi* | [**NotificationMarkRead**](docs/AccountApi.md#notificationmarkread) | **Put** /users/{user_id}/notifications/{notification_id}/read | NotificationMarkRead
*AccountApi* | [**UserGetAccountHistory**](docs/AccountApi.md#usergetaccounthistory) | **Get** /users/{user_id}/account/history | UserGetAccountHistory
*AccountApi* | [**UserGetNotifications**](docs/AccountApi.md#usergetnotifications) | **Get** /users/{user_id}/notifications | UserGetNotifications
*AccountApi* | [**UserGetProfitsAndLosses**](docs/AccountApi.md#usergetprofitsandlosses) | **Get** /users/{user_id}/account/pnl | UserGetProfitsAndLosses
*AuthApi* | [**Login**](docs/AuthApi.md#login) | **Post** /log_in | Login
*AuthApi* | [**Logout**](docs/AuthApi.md#logout) | **Post** /log_out | Logout
*AuthApi* | [**ResetPassword**](docs/AuthApi.md#resetpassword) | **Post** /passwords/reset | ResetPassword
*AuthApi* | [**ResetPasswordConfirm**](docs/AuthApi.md#resetpasswordconfirm) | **Put** /passwords/reset/{code}/confirm | ResetPasswordConfirm
*DefaultApi* | [**GetEventsCached**](docs/DefaultApi.md#geteventscached) | **Get** /events/ | GetEventsCached
*DefaultApi* | [**GetSeriesList**](docs/DefaultApi.md#getserieslist) | **Get** /series/ | GetSeriesList
*DefaultApi* | [**GetSeriesListCached**](docs/DefaultApi.md#getserieslistcached) | **Get** /cached/series | GetSeriesListCached
*DefaultApi* | [**GetTrades**](docs/DefaultApi.md#gettrades) | **Get** /trades | GetTrades
*EventsApi* | [**GetEventByTickerCached**](docs/EventsApi.md#geteventbytickercached) | **Get** /events/{ticker} | GetEventByTickerCached
*ExchangeApi* | [**GetExchangeStatus**](docs/ExchangeApi.md#getexchangestatus) | **Get** /exchange/status | 
*MarketApi* | [**GetActiveMarkets**](docs/MarketApi.md#getactivemarkets) | **Get** /active_markets | GetActiveMarkets
*MarketApi* | [**GetCandlestickMarketHistory**](docs/MarketApi.md#getcandlestickmarkethistory) | **Get** /markets/{market_id}/candlestick | GetCandlestickMarketHistory
*MarketApi* | [**GetCandlestickMarketHistoryCached**](docs/MarketApi.md#getcandlestickmarkethistorycached) | **Get** /cached/markets/{market_id}/candlestick | GetCandlestickMarketHistoryCached
*MarketApi* | [**GetMarket**](docs/MarketApi.md#getmarket) | **Get** /markets/{market_id} | GetMarket
*MarketApi* | [**GetMarketByTicker**](docs/MarketApi.md#getmarketbyticker) | **Get** /markets_by_ticker/{ticker_name} | GetMarketByTicker
*MarketApi* | [**GetMarketByTickerCached**](docs/MarketApi.md#getmarketbytickercached) | **Get** /cached/markets_by_ticker/{ticker_name} | GetMarketByTickerCached
*MarketApi* | [**GetMarketCached**](docs/MarketApi.md#getmarketcached) | **Get** /cached/markets/{market_id} | GetMarketCached
*MarketApi* | [**GetMarketHistory**](docs/MarketApi.md#getmarkethistory) | **Get** /markets/{market_id}/stats_history | GetMarketHistory
*MarketApi* | [**GetMarketHistoryCached**](docs/MarketApi.md#getmarkethistorycached) | **Get** /cached/markets/{market_id}/stats_history | GetMarketHistoryCached
*MarketApi* | [**GetMarketOrderBook**](docs/MarketApi.md#getmarketorderbook) | **Get** /markets/{market_id}/order_book | GetMarketOrderBook
*MarketApi* | [**GetMarketOrderBookCached**](docs/MarketApi.md#getmarketorderbookcached) | **Get** /cached/markets/{market_id}/order_book | GetMarketOrderBookCached
*MarketApi* | [**GetMarkets**](docs/MarketApi.md#getmarkets) | **Get** /markets | GetMarkets
*MarketApi* | [**GetMarketsCached**](docs/MarketApi.md#getmarketscached) | **Get** /cached/markets | GetMarketsCached
*PortfolioApi* | [**UserGetPortfolioHistory**](docs/PortfolioApi.md#usergetportfoliohistory) | **Get** /users/{user_id}/portfolio/history | UserGetPortfolioHistory
*RangedMarketApi* | [**GetRangedMarket**](docs/RangedMarketApi.md#getrangedmarket) | **Get** /ranged_markets/{ranged_market_id} | GetRangedMarket
*RangedMarketsApi* | [**GetRangedMarketByTicker**](docs/RangedMarketsApi.md#getrangedmarketbyticker) | **Get** /ranged_markets_by_ticker/{ticker} | GetRangedMarketByTicker
*RangedMarketsApi* | [**UserGetRangedMarketPosition**](docs/RangedMarketsApi.md#usergetrangedmarketposition) | **Get** /users/{user_id}/ranged_positions/{ranged_market_id} | UserGetRangedMarketPosition
*SeriesApi* | [**GetSeriesByTicker**](docs/SeriesApi.md#getseriesbyticker) | **Get** /series/{series_ticker} | GetSeriesByTicker
*UserApi* | [**UserAddWatchlist**](docs/UserApi.md#useraddwatchlist) | **Put** /users/{user_id}/watchlist/{market_id} | UserAddWatchlist
*UserApi* | [**UserBatchOrdersCancel**](docs/UserApi.md#userbatchorderscancel) | **Delete** /users/{user_id}/batch_orders | UserBatchOrdersCancel
*UserApi* | [**UserBatchOrdersCreate**](docs/UserApi.md#userbatchorderscreate) | **Post** /users/{user_id}/batch_orders | UserBatchOrdersCreate
*UserApi* | [**UserChangePassword**](docs/UserApi.md#userchangepassword) | **Put** /users/{user_id}/password | UserChangePassword
*UserApi* | [**UserDeactivate**](docs/UserApi.md#userdeactivate) | **Delete** /users | UserDeactivate
*UserApi* | [**UserGetBalance**](docs/UserApi.md#usergetbalance) | **Get** /users/{user_id}/balance | UserGetBalance
*UserApi* | [**UserGetMarketPosition**](docs/UserApi.md#usergetmarketposition) | **Get** /users/{user_id}/positions/{market_id} | UserGetMarketPosition
*UserApi* | [**UserGetMarketPositions**](docs/UserApi.md#usergetmarketpositions) | **Get** /users/{user_id}/positions | UserGetMarketPositions
*UserApi* | [**UserGetProfile**](docs/UserApi.md#usergetprofile) | **Get** /users/{user_id} | UserGetProfile
*UserApi* | [**UserGetReferralInfo**](docs/UserApi.md#usergetreferralinfo) | **Get** /users/{user_id}/referrals | UserGetReferralInfo
*UserApi* | [**UserGetWatchlist**](docs/UserApi.md#usergetwatchlist) | **Get** /users/{user_id}/watchlist | UserGetWatchlist
*UserApi* | [**UserOrderCancel**](docs/UserApi.md#userordercancel) | **Delete** /users/{user_id}/orders/{order_id} | UserOrderCancel
*UserApi* | [**UserOrderCreate**](docs/UserApi.md#userordercreate) | **Post** /users/{user_id}/orders | UserOrderCreate
*UserApi* | [**UserOrderDecrease**](docs/UserApi.md#userorderdecrease) | **Post** /users/{user_id}/orders/{order_id}/decrease | UserOrderDecrease
*UserApi* | [**UserOrdersGet**](docs/UserApi.md#userordersget) | **Get** /users/{user_id}/orders | UserOrdersGet
*UserApi* | [**UserRemoveWatchlist**](docs/UserApi.md#userremovewatchlist) | **Delete** /users/{user_id}/watchlist/{market_id} | UserRemoveWatchlist
*UserApi* | [**UserTradesGet**](docs/UserApi.md#usertradesget) | **Get** /users/{user_id}/trades | UserTradesGet
*UsersApi* | [**GetUserImmediateFunding**](docs/UsersApi.md#getuserimmediatefunding) | **Get** /users/{user_id}/immediate_funding | GetUserImmediateFunding
*UsersApi* | [**GetUserWithdrawalAvailableBalance**](docs/UsersApi.md#getuserwithdrawalavailablebalance) | **Get** /users/{user_id}/withdrawals/available | GetUserWithdrawalAvailableBalance

## Documentation For Models

 - [AccountHistoryEntry](docs/AccountHistoryEntry.md)
 - [AccountHistoryEntryData](docs/AccountHistoryEntryData.md)
 - [BankAccountDetails](docs/BankAccountDetails.md)
 - [ChangeSubscriptionRequest](docs/ChangeSubscriptionRequest.md)
 - [ChangeSubscriptionResponse](docs/ChangeSubscriptionResponse.md)
 - [ConfirmPasswordResetRequest](docs/ConfirmPasswordResetRequest.md)
 - [CreateUserRequest](docs/CreateUserRequest.md)
 - [CreateUserResponse](docs/CreateUserResponse.md)
 - [CreditHistory](docs/CreditHistory.md)
 - [Deposit](docs/Deposit.md)
 - [DepositHistory](docs/DepositHistory.md)
 - [EventChildMarket](docs/EventChildMarket.md)
 - [EventData](docs/EventData.md)
 - [GetActiveMarketsResponse](docs/GetActiveMarketsResponse.md)
 - [GetCandlestickMarketHistoryResponse](docs/GetCandlestickMarketHistoryResponse.md)
 - [GetEventByTickerResponse](docs/GetEventByTickerResponse.md)
 - [GetEventsResponse](docs/GetEventsResponse.md)
 - [GetMarketHistoryResponse](docs/GetMarketHistoryResponse.md)
 - [GetMarketOrderBookResponse](docs/GetMarketOrderBookResponse.md)
 - [GetNotificationPreferencesResponse](docs/GetNotificationPreferencesResponse.md)
 - [GetRangedMarketByTickerResponse](docs/GetRangedMarketByTickerResponse.md)
 - [GetRangedMarketResponse](docs/GetRangedMarketResponse.md)
 - [GetRangedMarketsResponse](docs/GetRangedMarketsResponse.md)
 - [GetSeriesByTickerResponse](docs/GetSeriesByTickerResponse.md)
 - [GetSeriesListResponse](docs/GetSeriesListResponse.md)
 - [GetUserDepositsResponse](docs/GetUserDepositsResponse.md)
 - [GetUserImmediateFundingResponse](docs/GetUserImmediateFundingResponse.md)
 - [GetUserWithdrawalsResponse](docs/GetUserWithdrawalsResponse.md)
 - [JsonError](docs/JsonError.md)
 - [LoginRequest](docs/LoginRequest.md)
 - [LoginResponse](docs/LoginResponse.md)
 - [Market](docs/Market.md)
 - [MarketActivity](docs/MarketActivity.md)
 - [MarketHistoryCandleStickInfo](docs/MarketHistoryCandleStickInfo.md)
 - [MarketPosition](docs/MarketPosition.md)
 - [MarketStatsPoint](docs/MarketStatsPoint.md)
 - [MarketTransaction](docs/MarketTransaction.md)
 - [Notification](docs/Notification.md)
 - [Olhcm](docs/Olhcm.md)
 - [Order](docs/Order.md)
 - [OrderBook](docs/OrderBook.md)
 - [OrderHistory](docs/OrderHistory.md)
 - [PeopleReferred](docs/PeopleReferred.md)
 - [PortfolioMeasurement](docs/PortfolioMeasurement.md)
 - [PublicTrade](docs/PublicTrade.md)
 - [RangedMarket](docs/RangedMarket.md)
 - [RangedMarketPosition](docs/RangedMarketPosition.md)
 - [ResetPasswordRequest](docs/ResetPasswordRequest.md)
 - [Series](docs/Series.md)
 - [SettlementHistory](docs/SettlementHistory.md)
 - [SettlementSource](docs/SettlementSource.md)
 - [SubscriptionPreference](docs/SubscriptionPreference.md)
 - [TradeHistory](docs/TradeHistory.md)
 - [TradesGetResponse](docs/TradesGetResponse.md)
 - [User](docs/User.md)
 - [UserBatchOrdersCancelRequest](docs/UserBatchOrdersCancelRequest.md)
 - [UserBatchOrdersCancelResponse](docs/UserBatchOrdersCancelResponse.md)
 - [UserBatchOrdersCancelSingleOrderResponse](docs/UserBatchOrdersCancelSingleOrderResponse.md)
 - [UserBatchOrdersCreateRequest](docs/UserBatchOrdersCreateRequest.md)
 - [UserBatchOrdersCreateResponse](docs/UserBatchOrdersCreateResponse.md)
 - [UserBatchOrdersCreateSingleOrderResponse](docs/UserBatchOrdersCreateSingleOrderResponse.md)
 - [UserChangePasswordRequest](docs/UserChangePasswordRequest.md)
 - [UserDepositRequest](docs/UserDepositRequest.md)
 - [UserDepositResponse](docs/UserDepositResponse.md)
 - [UserGetAccountHistoryResponse](docs/UserGetAccountHistoryResponse.md)
 - [UserGetBalanceResponse](docs/UserGetBalanceResponse.md)
 - [UserGetCurrentPortfolioValueResponse](docs/UserGetCurrentPortfolioValueResponse.md)
 - [UserGetMarketPositionResponse](docs/UserGetMarketPositionResponse.md)
 - [UserGetMarketPositionsResponse](docs/UserGetMarketPositionsResponse.md)
 - [UserGetMarketResponse](docs/UserGetMarketResponse.md)
 - [UserGetMarketsResponse](docs/UserGetMarketsResponse.md)
 - [UserGetNotificationsResponse](docs/UserGetNotificationsResponse.md)
 - [UserGetPortfolioHistoryRequest](docs/UserGetPortfolioHistoryRequest.md)
 - [UserGetPortfolioHistoryResponse](docs/UserGetPortfolioHistoryResponse.md)
 - [UserGetProfileResponse](docs/UserGetProfileResponse.md)
 - [UserGetRangedMarketPositionResponse](docs/UserGetRangedMarketPositionResponse.md)
 - [UserGetReferralInfoResponse](docs/UserGetReferralInfoResponse.md)
 - [UserGetWatchlistResponse](docs/UserGetWatchlistResponse.md)
 - [UserListLedgerxBankAccountsResponse](docs/UserListLedgerxBankAccountsResponse.md)
 - [UserOrderCreateRequest](docs/UserOrderCreateRequest.md)
 - [UserOrderCreateResponse](docs/UserOrderCreateResponse.md)
 - [UserOrderDecreaseRequest](docs/UserOrderDecreaseRequest.md)
 - [UserOrderDecreaseResponse](docs/UserOrderDecreaseResponse.md)
 - [UserOrdersGetResponse](docs/UserOrdersGetResponse.md)
 - [UserTrade](docs/UserTrade.md)
 - [UserTradesGetResponse](docs/UserTradesGetResponse.md)
 - [UserUpdateProfileRequest](docs/UserUpdateProfileRequest.md)
 - [UserWithdrawalAvailableBalanceResponse](docs/UserWithdrawalAvailableBalanceResponse.md)
 - [UserWithdrawalRequest](docs/UserWithdrawalRequest.md)
 - [UserWithdrawalResponse](docs/UserWithdrawalResponse.md)
 - [Watchlist](docs/Watchlist.md)
 - [Withdrawal](docs/Withdrawal.md)
 - [WithdrawalHistory](docs/WithdrawalHistory.md)

## Documentation For Authorization

## cookie
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author

