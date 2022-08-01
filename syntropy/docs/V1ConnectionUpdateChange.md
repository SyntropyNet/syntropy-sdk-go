# V1ConnectionUpdateChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionGroupId** | **int32** |  | 
**IsSdnEnabled** | **bool** |  | 

## Methods

### NewV1ConnectionUpdateChange

`func NewV1ConnectionUpdateChange(connectionGroupId int32, isSdnEnabled bool, ) *V1ConnectionUpdateChange`

NewV1ConnectionUpdateChange instantiates a new V1ConnectionUpdateChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ConnectionUpdateChangeWithDefaults

`func NewV1ConnectionUpdateChangeWithDefaults() *V1ConnectionUpdateChange`

NewV1ConnectionUpdateChangeWithDefaults instantiates a new V1ConnectionUpdateChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionGroupId

`func (o *V1ConnectionUpdateChange) GetConnectionGroupId() int32`

GetConnectionGroupId returns the ConnectionGroupId field if non-nil, zero value otherwise.

### GetConnectionGroupIdOk

`func (o *V1ConnectionUpdateChange) GetConnectionGroupIdOk() (*int32, bool)`

GetConnectionGroupIdOk returns a tuple with the ConnectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionGroupId

`func (o *V1ConnectionUpdateChange) SetConnectionGroupId(v int32)`

SetConnectionGroupId sets ConnectionGroupId field to given value.


### GetIsSdnEnabled

`func (o *V1ConnectionUpdateChange) GetIsSdnEnabled() bool`

GetIsSdnEnabled returns the IsSdnEnabled field if non-nil, zero value otherwise.

### GetIsSdnEnabledOk

`func (o *V1ConnectionUpdateChange) GetIsSdnEnabledOk() (*bool, bool)`

GetIsSdnEnabledOk returns a tuple with the IsSdnEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSdnEnabled

`func (o *V1ConnectionUpdateChange) SetIsSdnEnabled(v bool)`

SetIsSdnEnabled sets IsSdnEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


