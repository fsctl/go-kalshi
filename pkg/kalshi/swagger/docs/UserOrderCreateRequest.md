# UserOrderCreateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | Specifies how many contracts should be bought | [default to null]
**ExpirationUnixTs** | **int64** | Specifies the expiration time of the order, in unix seconds.  If this is not supplied, or is 0, the order won&#x27;t expire. | [optional] [default to null]
**MarketId** | **string** | Specifies the id of the market for this order | [default to null]
**MaxCostCents** | **int64** |  | [optional] [default to null]
**Price** | **int64** |  | [default to null]
**SellPositionCapped** | **bool** | Specifies whether the order place count should be capped by the members current position. | [optional] [default to null]
**Side** | **string** | Specifies if this is a &#x27;yes&#x27; or &#x27;no&#x27; order | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

