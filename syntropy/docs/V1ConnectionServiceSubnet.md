# V1ConnectionServiceSubnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentServiceSubnetId** | **int32** |  | 
**AgentConnectionSubnetId** | **int32** |  | 
**AgentConnectionSubnetIsEnabled** | **bool** |  | 
**AgentConnectionSubnetError** | **NullableString** |  | 
**AgentConnectionSubnetStatus** | [**AgentConnectionSubnetStatuses**](AgentConnectionSubnetStatuses.md) |  | 

## Methods

### NewV1ConnectionServiceSubnet

`func NewV1ConnectionServiceSubnet(agentServiceSubnetId int32, agentConnectionSubnetId int32, agentConnectionSubnetIsEnabled bool, agentConnectionSubnetError NullableString, agentConnectionSubnetStatus AgentConnectionSubnetStatuses, ) *V1ConnectionServiceSubnet`

NewV1ConnectionServiceSubnet instantiates a new V1ConnectionServiceSubnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ConnectionServiceSubnetWithDefaults

`func NewV1ConnectionServiceSubnetWithDefaults() *V1ConnectionServiceSubnet`

NewV1ConnectionServiceSubnetWithDefaults instantiates a new V1ConnectionServiceSubnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentServiceSubnetId

`func (o *V1ConnectionServiceSubnet) GetAgentServiceSubnetId() int32`

GetAgentServiceSubnetId returns the AgentServiceSubnetId field if non-nil, zero value otherwise.

### GetAgentServiceSubnetIdOk

`func (o *V1ConnectionServiceSubnet) GetAgentServiceSubnetIdOk() (*int32, bool)`

GetAgentServiceSubnetIdOk returns a tuple with the AgentServiceSubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentServiceSubnetId

`func (o *V1ConnectionServiceSubnet) SetAgentServiceSubnetId(v int32)`

SetAgentServiceSubnetId sets AgentServiceSubnetId field to given value.


### GetAgentConnectionSubnetId

`func (o *V1ConnectionServiceSubnet) GetAgentConnectionSubnetId() int32`

GetAgentConnectionSubnetId returns the AgentConnectionSubnetId field if non-nil, zero value otherwise.

### GetAgentConnectionSubnetIdOk

`func (o *V1ConnectionServiceSubnet) GetAgentConnectionSubnetIdOk() (*int32, bool)`

GetAgentConnectionSubnetIdOk returns a tuple with the AgentConnectionSubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentConnectionSubnetId

`func (o *V1ConnectionServiceSubnet) SetAgentConnectionSubnetId(v int32)`

SetAgentConnectionSubnetId sets AgentConnectionSubnetId field to given value.


### GetAgentConnectionSubnetIsEnabled

`func (o *V1ConnectionServiceSubnet) GetAgentConnectionSubnetIsEnabled() bool`

GetAgentConnectionSubnetIsEnabled returns the AgentConnectionSubnetIsEnabled field if non-nil, zero value otherwise.

### GetAgentConnectionSubnetIsEnabledOk

`func (o *V1ConnectionServiceSubnet) GetAgentConnectionSubnetIsEnabledOk() (*bool, bool)`

GetAgentConnectionSubnetIsEnabledOk returns a tuple with the AgentConnectionSubnetIsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentConnectionSubnetIsEnabled

`func (o *V1ConnectionServiceSubnet) SetAgentConnectionSubnetIsEnabled(v bool)`

SetAgentConnectionSubnetIsEnabled sets AgentConnectionSubnetIsEnabled field to given value.


### GetAgentConnectionSubnetError

`func (o *V1ConnectionServiceSubnet) GetAgentConnectionSubnetError() string`

GetAgentConnectionSubnetError returns the AgentConnectionSubnetError field if non-nil, zero value otherwise.

### GetAgentConnectionSubnetErrorOk

`func (o *V1ConnectionServiceSubnet) GetAgentConnectionSubnetErrorOk() (*string, bool)`

GetAgentConnectionSubnetErrorOk returns a tuple with the AgentConnectionSubnetError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentConnectionSubnetError

`func (o *V1ConnectionServiceSubnet) SetAgentConnectionSubnetError(v string)`

SetAgentConnectionSubnetError sets AgentConnectionSubnetError field to given value.


### SetAgentConnectionSubnetErrorNil

`func (o *V1ConnectionServiceSubnet) SetAgentConnectionSubnetErrorNil(b bool)`

 SetAgentConnectionSubnetErrorNil sets the value for AgentConnectionSubnetError to be an explicit nil

### UnsetAgentConnectionSubnetError
`func (o *V1ConnectionServiceSubnet) UnsetAgentConnectionSubnetError()`

UnsetAgentConnectionSubnetError ensures that no value is present for AgentConnectionSubnetError, not even an explicit nil
### GetAgentConnectionSubnetStatus

`func (o *V1ConnectionServiceSubnet) GetAgentConnectionSubnetStatus() AgentConnectionSubnetStatuses`

GetAgentConnectionSubnetStatus returns the AgentConnectionSubnetStatus field if non-nil, zero value otherwise.

### GetAgentConnectionSubnetStatusOk

`func (o *V1ConnectionServiceSubnet) GetAgentConnectionSubnetStatusOk() (*AgentConnectionSubnetStatuses, bool)`

GetAgentConnectionSubnetStatusOk returns a tuple with the AgentConnectionSubnetStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentConnectionSubnetStatus

`func (o *V1ConnectionServiceSubnet) SetAgentConnectionSubnetStatus(v AgentConnectionSubnetStatuses)`

SetAgentConnectionSubnetStatus sets AgentConnectionSubnetStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


