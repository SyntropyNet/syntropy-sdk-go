# V1NetworkConnectionsSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**V1ConnectionFilter**](V1ConnectionFilter.md) |  | [optional] 
**Order** | Pointer to [**[]V1ConnectionOrderInner**](V1ConnectionOrderInner.md) |  | [optional] 
**Skip** | Pointer to **int32** | Skip number of items. | [optional] [default to 0]
**Take** | Pointer to **int32** | Limit returned values count. | [optional] [default to 20]

## Methods

### NewV1NetworkConnectionsSearchRequest

`func NewV1NetworkConnectionsSearchRequest() *V1NetworkConnectionsSearchRequest`

NewV1NetworkConnectionsSearchRequest instantiates a new V1NetworkConnectionsSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NetworkConnectionsSearchRequestWithDefaults

`func NewV1NetworkConnectionsSearchRequestWithDefaults() *V1NetworkConnectionsSearchRequest`

NewV1NetworkConnectionsSearchRequestWithDefaults instantiates a new V1NetworkConnectionsSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *V1NetworkConnectionsSearchRequest) GetFilter() V1ConnectionFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *V1NetworkConnectionsSearchRequest) GetFilterOk() (*V1ConnectionFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *V1NetworkConnectionsSearchRequest) SetFilter(v V1ConnectionFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *V1NetworkConnectionsSearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetOrder

`func (o *V1NetworkConnectionsSearchRequest) GetOrder() []V1ConnectionOrderInner`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *V1NetworkConnectionsSearchRequest) GetOrderOk() (*[]V1ConnectionOrderInner, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *V1NetworkConnectionsSearchRequest) SetOrder(v []V1ConnectionOrderInner)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *V1NetworkConnectionsSearchRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetSkip

`func (o *V1NetworkConnectionsSearchRequest) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *V1NetworkConnectionsSearchRequest) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *V1NetworkConnectionsSearchRequest) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *V1NetworkConnectionsSearchRequest) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetTake

`func (o *V1NetworkConnectionsSearchRequest) GetTake() int32`

GetTake returns the Take field if non-nil, zero value otherwise.

### GetTakeOk

`func (o *V1NetworkConnectionsSearchRequest) GetTakeOk() (*int32, bool)`

GetTakeOk returns a tuple with the Take field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTake

`func (o *V1NetworkConnectionsSearchRequest) SetTake(v int32)`

SetTake sets Take field to given value.

### HasTake

`func (o *V1NetworkConnectionsSearchRequest) HasTake() bool`

HasTake returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


