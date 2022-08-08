# V1ConnectionServiceAgent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentServices** | [**[]V1ConnectionServiceAgentService**](V1ConnectionServiceAgentService.md) |  | 
**AgentType** | [**AgentType**](AgentType.md) |  | 
**AgentId** | **int32** |  | 

## Methods

### NewV1ConnectionServiceAgent

`func NewV1ConnectionServiceAgent(agentServices []V1ConnectionServiceAgentService, agentType AgentType, agentId int32, ) *V1ConnectionServiceAgent`

NewV1ConnectionServiceAgent instantiates a new V1ConnectionServiceAgent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ConnectionServiceAgentWithDefaults

`func NewV1ConnectionServiceAgentWithDefaults() *V1ConnectionServiceAgent`

NewV1ConnectionServiceAgentWithDefaults instantiates a new V1ConnectionServiceAgent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentServices

`func (o *V1ConnectionServiceAgent) GetAgentServices() []V1ConnectionServiceAgentService`

GetAgentServices returns the AgentServices field if non-nil, zero value otherwise.

### GetAgentServicesOk

`func (o *V1ConnectionServiceAgent) GetAgentServicesOk() (*[]V1ConnectionServiceAgentService, bool)`

GetAgentServicesOk returns a tuple with the AgentServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentServices

`func (o *V1ConnectionServiceAgent) SetAgentServices(v []V1ConnectionServiceAgentService)`

SetAgentServices sets AgentServices field to given value.


### GetAgentType

`func (o *V1ConnectionServiceAgent) GetAgentType() AgentType`

GetAgentType returns the AgentType field if non-nil, zero value otherwise.

### GetAgentTypeOk

`func (o *V1ConnectionServiceAgent) GetAgentTypeOk() (*AgentType, bool)`

GetAgentTypeOk returns a tuple with the AgentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentType

`func (o *V1ConnectionServiceAgent) SetAgentType(v AgentType)`

SetAgentType sets AgentType field to given value.


### GetAgentId

`func (o *V1ConnectionServiceAgent) GetAgentId() int32`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *V1ConnectionServiceAgent) GetAgentIdOk() (*int32, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *V1ConnectionServiceAgent) SetAgentId(v int32)`

SetAgentId sets AgentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


