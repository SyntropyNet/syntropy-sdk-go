# V1NetworkAgentsSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**V1AgentFilter**](V1AgentFilter.md) |  | [optional] 
**Order** | Pointer to [**[]AnyOfobjectobjectobjectobjectobject**](AnyOfobjectobjectobjectobjectobject.md) |  | [optional] 
**Skip** | Pointer to **int32** | Skip number of items. | [optional] [default to 0]
**Take** | Pointer to **int32** | Limit returned values count. | [optional] [default to 20]
**Search** | Pointer to **string** |  | [optional] 

## Methods

### NewV1NetworkAgentsSearchRequest

`func NewV1NetworkAgentsSearchRequest() *V1NetworkAgentsSearchRequest`

NewV1NetworkAgentsSearchRequest instantiates a new V1NetworkAgentsSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NetworkAgentsSearchRequestWithDefaults

`func NewV1NetworkAgentsSearchRequestWithDefaults() *V1NetworkAgentsSearchRequest`

NewV1NetworkAgentsSearchRequestWithDefaults instantiates a new V1NetworkAgentsSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *V1NetworkAgentsSearchRequest) GetFilter() V1AgentFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *V1NetworkAgentsSearchRequest) GetFilterOk() (*V1AgentFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *V1NetworkAgentsSearchRequest) SetFilter(v V1AgentFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *V1NetworkAgentsSearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetOrder

`func (o *V1NetworkAgentsSearchRequest) GetOrder() []AnyOfobjectobjectobjectobjectobject`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *V1NetworkAgentsSearchRequest) GetOrderOk() (*[]AnyOfobjectobjectobjectobjectobject, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *V1NetworkAgentsSearchRequest) SetOrder(v []AnyOfobjectobjectobjectobjectobject)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *V1NetworkAgentsSearchRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetSkip

`func (o *V1NetworkAgentsSearchRequest) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *V1NetworkAgentsSearchRequest) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *V1NetworkAgentsSearchRequest) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *V1NetworkAgentsSearchRequest) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetTake

`func (o *V1NetworkAgentsSearchRequest) GetTake() int32`

GetTake returns the Take field if non-nil, zero value otherwise.

### GetTakeOk

`func (o *V1NetworkAgentsSearchRequest) GetTakeOk() (*int32, bool)`

GetTakeOk returns a tuple with the Take field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTake

`func (o *V1NetworkAgentsSearchRequest) SetTake(v int32)`

SetTake sets Take field to given value.

### HasTake

`func (o *V1NetworkAgentsSearchRequest) HasTake() bool`

HasTake returns a boolean if a field has been set.

### GetSearch

`func (o *V1NetworkAgentsSearchRequest) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *V1NetworkAgentsSearchRequest) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *V1NetworkAgentsSearchRequest) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *V1NetworkAgentsSearchRequest) HasSearch() bool`

HasSearch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


