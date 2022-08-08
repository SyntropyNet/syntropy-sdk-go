# V1ConnectionFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentConnectionGroupId** | Pointer to **[]int32** |  | [optional] 
**AgentId** | Pointer to **[]int32** |  | [optional] 
**AgentPair** | Pointer to [**[]V1AgentPairFilter**](V1AgentPairFilter.md) |  | [optional] 

## Methods

### NewV1ConnectionFilter

`func NewV1ConnectionFilter() *V1ConnectionFilter`

NewV1ConnectionFilter instantiates a new V1ConnectionFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ConnectionFilterWithDefaults

`func NewV1ConnectionFilterWithDefaults() *V1ConnectionFilter`

NewV1ConnectionFilterWithDefaults instantiates a new V1ConnectionFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentConnectionGroupId

`func (o *V1ConnectionFilter) GetAgentConnectionGroupId() []int32`

GetAgentConnectionGroupId returns the AgentConnectionGroupId field if non-nil, zero value otherwise.

### GetAgentConnectionGroupIdOk

`func (o *V1ConnectionFilter) GetAgentConnectionGroupIdOk() (*[]int32, bool)`

GetAgentConnectionGroupIdOk returns a tuple with the AgentConnectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentConnectionGroupId

`func (o *V1ConnectionFilter) SetAgentConnectionGroupId(v []int32)`

SetAgentConnectionGroupId sets AgentConnectionGroupId field to given value.

### HasAgentConnectionGroupId

`func (o *V1ConnectionFilter) HasAgentConnectionGroupId() bool`

HasAgentConnectionGroupId returns a boolean if a field has been set.

### GetAgentId

`func (o *V1ConnectionFilter) GetAgentId() []int32`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *V1ConnectionFilter) GetAgentIdOk() (*[]int32, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *V1ConnectionFilter) SetAgentId(v []int32)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *V1ConnectionFilter) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetAgentPair

`func (o *V1ConnectionFilter) GetAgentPair() []V1AgentPairFilter`

GetAgentPair returns the AgentPair field if non-nil, zero value otherwise.

### GetAgentPairOk

`func (o *V1ConnectionFilter) GetAgentPairOk() (*[]V1AgentPairFilter, bool)`

GetAgentPairOk returns a tuple with the AgentPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPair

`func (o *V1ConnectionFilter) SetAgentPair(v []V1AgentPairFilter)`

SetAgentPair sets AgentPair field to given value.

### HasAgentPair

`func (o *V1ConnectionFilter) HasAgentPair() bool`

HasAgentPair returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


