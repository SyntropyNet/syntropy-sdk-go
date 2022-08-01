# V1NetworkConnectionsCreateMeshRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentIds** | [**[]V1NetworkConnectionsCreateMeshRequestAgentIds**](V1NetworkConnectionsCreateMeshRequestAgentIds.md) |  | 
**SdnEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewV1NetworkConnectionsCreateMeshRequest

`func NewV1NetworkConnectionsCreateMeshRequest(agentIds []V1NetworkConnectionsCreateMeshRequestAgentIds, ) *V1NetworkConnectionsCreateMeshRequest`

NewV1NetworkConnectionsCreateMeshRequest instantiates a new V1NetworkConnectionsCreateMeshRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NetworkConnectionsCreateMeshRequestWithDefaults

`func NewV1NetworkConnectionsCreateMeshRequestWithDefaults() *V1NetworkConnectionsCreateMeshRequest`

NewV1NetworkConnectionsCreateMeshRequestWithDefaults instantiates a new V1NetworkConnectionsCreateMeshRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentIds

`func (o *V1NetworkConnectionsCreateMeshRequest) GetAgentIds() []V1NetworkConnectionsCreateMeshRequestAgentIds`

GetAgentIds returns the AgentIds field if non-nil, zero value otherwise.

### GetAgentIdsOk

`func (o *V1NetworkConnectionsCreateMeshRequest) GetAgentIdsOk() (*[]V1NetworkConnectionsCreateMeshRequestAgentIds, bool)`

GetAgentIdsOk returns a tuple with the AgentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentIds

`func (o *V1NetworkConnectionsCreateMeshRequest) SetAgentIds(v []V1NetworkConnectionsCreateMeshRequestAgentIds)`

SetAgentIds sets AgentIds field to given value.


### GetSdnEnabled

`func (o *V1NetworkConnectionsCreateMeshRequest) GetSdnEnabled() bool`

GetSdnEnabled returns the SdnEnabled field if non-nil, zero value otherwise.

### GetSdnEnabledOk

`func (o *V1NetworkConnectionsCreateMeshRequest) GetSdnEnabledOk() (*bool, bool)`

GetSdnEnabledOk returns a tuple with the SdnEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdnEnabled

`func (o *V1NetworkConnectionsCreateMeshRequest) SetSdnEnabled(v bool)`

SetSdnEnabled sets SdnEnabled field to given value.

### HasSdnEnabled

`func (o *V1NetworkConnectionsCreateMeshRequest) HasSdnEnabled() bool`

HasSdnEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


