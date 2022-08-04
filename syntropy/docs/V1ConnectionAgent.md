# V1ConnectionAgent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentIsOnline** | **bool** |  | 
**AgentStatus** | [**NullableAgentStatus**](AgentStatus.md) |  | 
**AgentSubnetsCount** | **int32** |  | 
**AgentIsVirtual** | **bool** |  | 
**AgentName** | **string** |  | 
**AgentPublicIpv4** | **string** |  | 
**AgentId** | **int32** |  | 
**AgentProviderId** | **NullableInt32** |  | 

## Methods

### NewV1ConnectionAgent

`func NewV1ConnectionAgent(agentIsOnline bool, agentStatus NullableAgentStatus, agentSubnetsCount int32, agentIsVirtual bool, agentName string, agentPublicIpv4 string, agentId int32, agentProviderId NullableInt32, ) *V1ConnectionAgent`

NewV1ConnectionAgent instantiates a new V1ConnectionAgent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ConnectionAgentWithDefaults

`func NewV1ConnectionAgentWithDefaults() *V1ConnectionAgent`

NewV1ConnectionAgentWithDefaults instantiates a new V1ConnectionAgent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentIsOnline

`func (o *V1ConnectionAgent) GetAgentIsOnline() bool`

GetAgentIsOnline returns the AgentIsOnline field if non-nil, zero value otherwise.

### GetAgentIsOnlineOk

`func (o *V1ConnectionAgent) GetAgentIsOnlineOk() (*bool, bool)`

GetAgentIsOnlineOk returns a tuple with the AgentIsOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentIsOnline

`func (o *V1ConnectionAgent) SetAgentIsOnline(v bool)`

SetAgentIsOnline sets AgentIsOnline field to given value.


### GetAgentStatus

`func (o *V1ConnectionAgent) GetAgentStatus() AgentStatus`

GetAgentStatus returns the AgentStatus field if non-nil, zero value otherwise.

### GetAgentStatusOk

`func (o *V1ConnectionAgent) GetAgentStatusOk() (*AgentStatus, bool)`

GetAgentStatusOk returns a tuple with the AgentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentStatus

`func (o *V1ConnectionAgent) SetAgentStatus(v AgentStatus)`

SetAgentStatus sets AgentStatus field to given value.


### SetAgentStatusNil

`func (o *V1ConnectionAgent) SetAgentStatusNil(b bool)`

 SetAgentStatusNil sets the value for AgentStatus to be an explicit nil

### UnsetAgentStatus
`func (o *V1ConnectionAgent) UnsetAgentStatus()`

UnsetAgentStatus ensures that no value is present for AgentStatus, not even an explicit nil
### GetAgentSubnetsCount

`func (o *V1ConnectionAgent) GetAgentSubnetsCount() int32`

GetAgentSubnetsCount returns the AgentSubnetsCount field if non-nil, zero value otherwise.

### GetAgentSubnetsCountOk

`func (o *V1ConnectionAgent) GetAgentSubnetsCountOk() (*int32, bool)`

GetAgentSubnetsCountOk returns a tuple with the AgentSubnetsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentSubnetsCount

`func (o *V1ConnectionAgent) SetAgentSubnetsCount(v int32)`

SetAgentSubnetsCount sets AgentSubnetsCount field to given value.


### GetAgentIsVirtual

`func (o *V1ConnectionAgent) GetAgentIsVirtual() bool`

GetAgentIsVirtual returns the AgentIsVirtual field if non-nil, zero value otherwise.

### GetAgentIsVirtualOk

`func (o *V1ConnectionAgent) GetAgentIsVirtualOk() (*bool, bool)`

GetAgentIsVirtualOk returns a tuple with the AgentIsVirtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentIsVirtual

`func (o *V1ConnectionAgent) SetAgentIsVirtual(v bool)`

SetAgentIsVirtual sets AgentIsVirtual field to given value.


### GetAgentName

`func (o *V1ConnectionAgent) GetAgentName() string`

GetAgentName returns the AgentName field if non-nil, zero value otherwise.

### GetAgentNameOk

`func (o *V1ConnectionAgent) GetAgentNameOk() (*string, bool)`

GetAgentNameOk returns a tuple with the AgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentName

`func (o *V1ConnectionAgent) SetAgentName(v string)`

SetAgentName sets AgentName field to given value.


### GetAgentPublicIpv4

`func (o *V1ConnectionAgent) GetAgentPublicIpv4() string`

GetAgentPublicIpv4 returns the AgentPublicIpv4 field if non-nil, zero value otherwise.

### GetAgentPublicIpv4Ok

`func (o *V1ConnectionAgent) GetAgentPublicIpv4Ok() (*string, bool)`

GetAgentPublicIpv4Ok returns a tuple with the AgentPublicIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPublicIpv4

`func (o *V1ConnectionAgent) SetAgentPublicIpv4(v string)`

SetAgentPublicIpv4 sets AgentPublicIpv4 field to given value.


### GetAgentId

`func (o *V1ConnectionAgent) GetAgentId() int32`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *V1ConnectionAgent) GetAgentIdOk() (*int32, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *V1ConnectionAgent) SetAgentId(v int32)`

SetAgentId sets AgentId field to given value.


### GetAgentProviderId

`func (o *V1ConnectionAgent) GetAgentProviderId() int32`

GetAgentProviderId returns the AgentProviderId field if non-nil, zero value otherwise.

### GetAgentProviderIdOk

`func (o *V1ConnectionAgent) GetAgentProviderIdOk() (*int32, bool)`

GetAgentProviderIdOk returns a tuple with the AgentProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentProviderId

`func (o *V1ConnectionAgent) SetAgentProviderId(v int32)`

SetAgentProviderId sets AgentProviderId field to given value.


### SetAgentProviderIdNil

`func (o *V1ConnectionAgent) SetAgentProviderIdNil(b bool)`

 SetAgentProviderIdNil sets the value for AgentProviderId to be an explicit nil

### UnsetAgentProviderId
`func (o *V1ConnectionAgent) UnsetAgentProviderId()`

UnsetAgentProviderId ensures that no value is present for AgentProviderId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


