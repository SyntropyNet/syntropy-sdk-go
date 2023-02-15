# V1NetworkConnectionsCreateP2PRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentPairs** | [**[]V1NetworkConnectionsCreateP2PRequestAgentPairsInner**](V1NetworkConnectionsCreateP2PRequestAgentPairsInner.md) |  | 
**SdnEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewV1NetworkConnectionsCreateP2PRequest

`func NewV1NetworkConnectionsCreateP2PRequest(agentPairs []V1NetworkConnectionsCreateP2PRequestAgentPairsInner, ) *V1NetworkConnectionsCreateP2PRequest`

NewV1NetworkConnectionsCreateP2PRequest instantiates a new V1NetworkConnectionsCreateP2PRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NetworkConnectionsCreateP2PRequestWithDefaults

`func NewV1NetworkConnectionsCreateP2PRequestWithDefaults() *V1NetworkConnectionsCreateP2PRequest`

NewV1NetworkConnectionsCreateP2PRequestWithDefaults instantiates a new V1NetworkConnectionsCreateP2PRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentPairs

`func (o *V1NetworkConnectionsCreateP2PRequest) GetAgentPairs() []V1NetworkConnectionsCreateP2PRequestAgentPairsInner`

GetAgentPairs returns the AgentPairs field if non-nil, zero value otherwise.

### GetAgentPairsOk

`func (o *V1NetworkConnectionsCreateP2PRequest) GetAgentPairsOk() (*[]V1NetworkConnectionsCreateP2PRequestAgentPairsInner, bool)`

GetAgentPairsOk returns a tuple with the AgentPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPairs

`func (o *V1NetworkConnectionsCreateP2PRequest) SetAgentPairs(v []V1NetworkConnectionsCreateP2PRequestAgentPairsInner)`

SetAgentPairs sets AgentPairs field to given value.


### GetSdnEnabled

`func (o *V1NetworkConnectionsCreateP2PRequest) GetSdnEnabled() bool`

GetSdnEnabled returns the SdnEnabled field if non-nil, zero value otherwise.

### GetSdnEnabledOk

`func (o *V1NetworkConnectionsCreateP2PRequest) GetSdnEnabledOk() (*bool, bool)`

GetSdnEnabledOk returns a tuple with the SdnEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdnEnabled

`func (o *V1NetworkConnectionsCreateP2PRequest) SetSdnEnabled(v bool)`

SetSdnEnabled sets SdnEnabled field to given value.

### HasSdnEnabled

`func (o *V1NetworkConnectionsCreateP2PRequest) HasSdnEnabled() bool`

HasSdnEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


