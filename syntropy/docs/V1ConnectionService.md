# V1ConnectionService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentConnectionGroupId** | **int32** |  | 
**AgentConnectionSubnets** | [**[]V1ConnectionServiceSubnet**](V1ConnectionServiceSubnet.md) |  | 
**Agent2** | [**V1ConnectionServiceAgent**](V1ConnectionServiceAgent.md) |  | 
**Agent1** | [**V1ConnectionServiceAgent**](V1ConnectionServiceAgent.md) |  | 

## Methods

### NewV1ConnectionService

`func NewV1ConnectionService(agentConnectionGroupId int32, agentConnectionSubnets []V1ConnectionServiceSubnet, agent2 V1ConnectionServiceAgent, agent1 V1ConnectionServiceAgent, ) *V1ConnectionService`

NewV1ConnectionService instantiates a new V1ConnectionService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ConnectionServiceWithDefaults

`func NewV1ConnectionServiceWithDefaults() *V1ConnectionService`

NewV1ConnectionServiceWithDefaults instantiates a new V1ConnectionService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentConnectionGroupId

`func (o *V1ConnectionService) GetAgentConnectionGroupId() int32`

GetAgentConnectionGroupId returns the AgentConnectionGroupId field if non-nil, zero value otherwise.

### GetAgentConnectionGroupIdOk

`func (o *V1ConnectionService) GetAgentConnectionGroupIdOk() (*int32, bool)`

GetAgentConnectionGroupIdOk returns a tuple with the AgentConnectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentConnectionGroupId

`func (o *V1ConnectionService) SetAgentConnectionGroupId(v int32)`

SetAgentConnectionGroupId sets AgentConnectionGroupId field to given value.


### GetAgentConnectionSubnets

`func (o *V1ConnectionService) GetAgentConnectionSubnets() []V1ConnectionServiceSubnet`

GetAgentConnectionSubnets returns the AgentConnectionSubnets field if non-nil, zero value otherwise.

### GetAgentConnectionSubnetsOk

`func (o *V1ConnectionService) GetAgentConnectionSubnetsOk() (*[]V1ConnectionServiceSubnet, bool)`

GetAgentConnectionSubnetsOk returns a tuple with the AgentConnectionSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentConnectionSubnets

`func (o *V1ConnectionService) SetAgentConnectionSubnets(v []V1ConnectionServiceSubnet)`

SetAgentConnectionSubnets sets AgentConnectionSubnets field to given value.


### GetAgent2

`func (o *V1ConnectionService) GetAgent2() V1ConnectionServiceAgent`

GetAgent2 returns the Agent2 field if non-nil, zero value otherwise.

### GetAgent2Ok

`func (o *V1ConnectionService) GetAgent2Ok() (*V1ConnectionServiceAgent, bool)`

GetAgent2Ok returns a tuple with the Agent2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent2

`func (o *V1ConnectionService) SetAgent2(v V1ConnectionServiceAgent)`

SetAgent2 sets Agent2 field to given value.


### GetAgent1

`func (o *V1ConnectionService) GetAgent1() V1ConnectionServiceAgent`

GetAgent1 returns the Agent1 field if non-nil, zero value otherwise.

### GetAgent1Ok

`func (o *V1ConnectionService) GetAgent1Ok() (*V1ConnectionServiceAgent, bool)`

GetAgent1Ok returns a tuple with the Agent1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent1

`func (o *V1ConnectionService) SetAgent1(v V1ConnectionServiceAgent)`

SetAgent1 sets Agent1 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


