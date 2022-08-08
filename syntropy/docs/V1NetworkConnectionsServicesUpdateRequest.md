# V1NetworkConnectionsServicesUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentConnectionGroupId** | Pointer to **int32** |  | [optional] 
**Changes** | [**[]AgentServicesUpdateChanges**](AgentServicesUpdateChanges.md) |  | 

## Methods

### NewV1NetworkConnectionsServicesUpdateRequest

`func NewV1NetworkConnectionsServicesUpdateRequest(changes []AgentServicesUpdateChanges, ) *V1NetworkConnectionsServicesUpdateRequest`

NewV1NetworkConnectionsServicesUpdateRequest instantiates a new V1NetworkConnectionsServicesUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NetworkConnectionsServicesUpdateRequestWithDefaults

`func NewV1NetworkConnectionsServicesUpdateRequestWithDefaults() *V1NetworkConnectionsServicesUpdateRequest`

NewV1NetworkConnectionsServicesUpdateRequestWithDefaults instantiates a new V1NetworkConnectionsServicesUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentConnectionGroupId

`func (o *V1NetworkConnectionsServicesUpdateRequest) GetAgentConnectionGroupId() int32`

GetAgentConnectionGroupId returns the AgentConnectionGroupId field if non-nil, zero value otherwise.

### GetAgentConnectionGroupIdOk

`func (o *V1NetworkConnectionsServicesUpdateRequest) GetAgentConnectionGroupIdOk() (*int32, bool)`

GetAgentConnectionGroupIdOk returns a tuple with the AgentConnectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentConnectionGroupId

`func (o *V1NetworkConnectionsServicesUpdateRequest) SetAgentConnectionGroupId(v int32)`

SetAgentConnectionGroupId sets AgentConnectionGroupId field to given value.

### HasAgentConnectionGroupId

`func (o *V1NetworkConnectionsServicesUpdateRequest) HasAgentConnectionGroupId() bool`

HasAgentConnectionGroupId returns a boolean if a field has been set.

### GetChanges

`func (o *V1NetworkConnectionsServicesUpdateRequest) GetChanges() []AgentServicesUpdateChanges`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *V1NetworkConnectionsServicesUpdateRequest) GetChangesOk() (*[]AgentServicesUpdateChanges, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *V1NetworkConnectionsServicesUpdateRequest) SetChanges(v []AgentServicesUpdateChanges)`

SetChanges sets Changes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


