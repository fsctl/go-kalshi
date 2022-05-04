# UserGetPortfolioHistoryRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxDate** | [**time.Time**](time.Time.md) | Restricts the response to orders before a timestamp in: query | [optional] [default to null]
**MaxTs** | **int64** | Restricts the response to orders before a timestamp in unix seconds, overrides max_date, defaults to now. in: query | [optional] [default to null]
**MinDate** | [**time.Time**](time.Time.md) | Restricts the response to orders after a timestamp in: query | [optional] [default to null]
**MinTs** | **int64** | Restricts the response to orders after a timestamp in unix seconds, overrides min_date, defaults to one hour before now. in: query | [optional] [default to null]
**NumBuckets** | **int32** | Determines the number of buckets to average over when performing subsampling, defaults to 1440. in: query | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

