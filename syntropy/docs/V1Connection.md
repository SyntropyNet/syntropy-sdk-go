# V1Connection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentConnectionSubnetsEnabledCount** | **int32** |  | 
**AgentConnectionGroupUpdatedAt** | **time.Time** |  | 
**AgentConnectionGroupStatusReason** | Pointer to **string** |  | [optional] 
**AgentConnectionGroupStatus** | [**V1ConnectionStatus**](V1ConnectionStatus.md) |  | 
**AgentConnectionGroupSdnEnabled** | **bool** |  | 
**AgentConnectionGroupId** | **int32** |  | 
**AgentConnectionGroupCreatedBy** | [**V1ConnectionCreatedBy**](V1ConnectionCreatedBy.md) |  | 
**Agent2** | [**V1ConnectionAgent**](V1ConnectionAgent.md) |  | 
**Agent1** | [**V1ConnectionAgent**](V1ConnectionAgent.md) |  | 

## Methods

### NewV1Connection

`func NewV1Connection(agentConnectionSubnetsEnabledCount int32, agentConnectionGroupUpdatedAt time.Time, agentConnectionGroupStatus V1ConnectionStatus, agentConnectionGroupSdnEnabled bool, agentConnectionGroupId int32, agentConnectionGroupCreatedBy V1ConnectionCreatedBy, agent2 V1ConnectionAgent, agent1 V1ConnectionAgent, ) *V1Connection`

NewV1Connection instantiates a new V1Connection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ConnectionWithDefaults

`func NewV1ConnectionWithDefaults() *V1Connection`

NewV1ConnectionWithDefaults instantiates a new V1Connection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentConnectionSubnetsEnabledCount

`func (o *V1Connection) GetAgentConnectionSubnetsEnabledCount() int32`

GetAgentConnectionSubnetsEnabledCount returns the AgentConnectionSubnetsEnabledCount field if non-nil, zero value otherwise.

### GetAgentConnectionSubnetsEnabledCountOk

`func (o *V1Connection) GetAgentConnectionSubnetsEnabledCountOk() (*int32, bool)`

GetAgentConnectionSubnetsEnabledCountOk returns a tuple with the AgentConnectionSubnetsEnabledCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentConnectionSubnetsEnabledCount

`func (o *V1Connection) SetAgentConnectionSubnetsEnabledCount(v int32)`

SetAgentConnectionSubnetsEnabledCount sets AgentConnectionSubnetsEnabledCount field to given value.


### GetAgentConnectionGroupUpdatedAt

`func (o *V1Connection) GetAgentConnectionGroupUpdatedAt() time.Time`

GetAgentConnectionGroupUpdatedAt returns the AgentConnectionGroupUpdatedAt field if non-nil, zero value otherwise.

### GetAgentConnectionGroupUpdatedAtOk

`func (o *V1Connection) GetAgentConnectionGroupUpdatedAtOk() (*time.Time, bool)`

GetAgentConnectionGroupUpdatedAtOk returns a tuple with the AgentConnectionGroupUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentConnectionGroupUpdatedAt

`func (o *V1Connection) SetAgentConnectionGroupUpdatedAt(v time.Time)`

SetAgentConnectionGroupUpdatedAt sets AgentConnectionGroupUpdatedAt field to given value.


### GetAgentConnectionGroupStatusReason

`func (o *V1Connection) GetAgentConnectionGroupStatusReason() string`

GetAgentConnectionGroupStatusReason returns the AgentConnectionGroupStatusReason field if non-nil, zero value otherwise.

### GetAgentConnectionGroupStatusReasonOk

`func (o *V1Connection) GetAgentConnectionGroupStatusReasonOk() (*string, bool)`

GetAgentConnectionGroupStatusReasonOk returns a tuple with the AgentConnectionGroupStatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentConnectionGroupStatusReason

`func (o *V1Connection) SetAgentConnectionGroupStatusReason(v string)`

SetAgentConnectionGroupStatusReason sets AgentConnectionGroupStatusReason field to given value.

### HasAgentConnectionGroupStatusReason

`func (o *V1Connection) HasAgentConnectionGroupStatusReason() bool`

HasAgentConnectionGroupStatusReason returns a boolean if a field has been set.

### GetAgentConnectionGroupStatus

`func (o *V1Connection) GetAgentConnectionGroupStatus() V1ConnectionStatus`

GetAgentConnectionGroupStatus returns the AgentConnectionGroupStatus field if non-nil, zero value otherwise.

### GetAgentConnectionGroupStatusOk

`func (o *V1Connection) GetAgentConnectionGroupStatusOk() (*V1ConnectionStatus, bool)`

GetAgentConnectionGroupStatusOk returns a tuple with the AgentConnectionGroupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentConnectionGroupStatus

`func (o *V1Connection) SetAgentConnectionGroupStatus(v V1ConnectionStatus)`

SetAgentConnectionGroupStatus sets AgentConnectionGroupStatus field to given value.


### GetAgentConnectionGroupSdnEnabled

`func (o *V1Connection) GetAgentConnectionGroupSdnEnabled() bool`

GetAgentConnectionGroupSdnEnabled returns the AgentConnectionGroupSdnEnabled field if non-nil, zero value otherwise.

### GetAgentConnectionGroupSdnEnabledOk

`func (o *V1Connection) GetAgentConnectionGroupSdnEnabledOk() (*bool, bool)`

GetAgentConnectionGroupSdnEnabledOk returns a tuple with the AgentConnectionGroupSdnEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentConnectionGroupSdnEnabled

`func (o *V1Connection) SetAgentConnectionGroupSdnEnabled(v bool)`

SetAgentConnectionGroupSdnEnabled sets AgentConnectionGroupSdnEnabled field to given value.


### GetAgentConnectionGroupId

`func (o *V1Connection) GetAgentConnectionGroupId() int32`

GetAgentConnectionGroupId returns the AgentConnectionGroupId field if non-nil, zero value otherwise.

### GetAgentConnectionGroupIdOk

`func (o *V1Connection) GetAgentConnectionGroupIdOk() (*int32, bool)`

GetAgentConnectionGroupIdOk returns a tuple with the AgentConnectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentConnectionGroupId

`func (o *V1Connection) SetAgentConnectionGroupId(v int32)`

SetAgentConnectionGroupId sets AgentConnectionGroupId field to given value.


### GetAgentConnectionGroupCreatedBy

`func (o *V1Connection) GetAgentConnectionGroupCreatedBy() V1ConnectionCreatedBy`

GetAgentConnectionGroupCreatedBy returns the AgentConnectionGroupCreatedBy field if non-nil, zero value otherwise.

### GetAgentConnectionGroupCreatedByOk

`func (o *V1Connection) GetAgentConnectionGroupCreatedByOk() (*V1ConnectionCreatedBy, bool)`

GetAgentConnectionGroupCreatedByOk returns a tuple with the AgentConnectionGroupCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentConnectionGroupCreatedBy

`func (o *V1Connection) SetAgentConnectionGroupCreatedBy(v V1ConnectionCreatedBy)`

SetAgentConnectionGroupCreatedBy sets AgentConnectionGroupCreatedBy field to given value.


### GetAgent2

`func (o *V1Connection) GetAgent2() V1ConnectionAgent`

GetAgent2 returns the Agent2 field if non-nil, zero value otherwise.

### GetAgent2Ok

`func (o *V1Connection) GetAgent2Ok() (*V1ConnectionAgent, bool)`

GetAgent2Ok returns a tuple with the Agent2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent2

`func (o *V1Connection) SetAgent2(v V1ConnectionAgent)`

SetAgent2 sets Agent2 field to given value.


### GetAgent1

`func (o *V1Connection) GetAgent1() V1ConnectionAgent`

GetAgent1 returns the Agent1 field if non-nil, zero value otherwise.

### GetAgent1Ok

`func (o *V1Connection) GetAgent1Ok() (*V1ConnectionAgent, bool)`

GetAgent1Ok returns a tuple with the Agent1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent1

`func (o *V1Connection) SetAgent1(v V1ConnectionAgent)`

SetAgent1 sets Agent1 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


