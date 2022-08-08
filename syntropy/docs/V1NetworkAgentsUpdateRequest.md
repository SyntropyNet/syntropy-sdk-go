# V1NetworkAgentsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentTags** | Pointer to **[]string** |  | [optional] 
**AgentName** | Pointer to **string** |  | [optional] 
**AgentProviderId** | Pointer to **int32** |  | [optional] 
**Network** | Pointer to [**AgentInterfacesMetadata**](AgentInterfacesMetadata.md) |  | [optional] 

## Methods

### NewV1NetworkAgentsUpdateRequest

`func NewV1NetworkAgentsUpdateRequest() *V1NetworkAgentsUpdateRequest`

NewV1NetworkAgentsUpdateRequest instantiates a new V1NetworkAgentsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NetworkAgentsUpdateRequestWithDefaults

`func NewV1NetworkAgentsUpdateRequestWithDefaults() *V1NetworkAgentsUpdateRequest`

NewV1NetworkAgentsUpdateRequestWithDefaults instantiates a new V1NetworkAgentsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentTags

`func (o *V1NetworkAgentsUpdateRequest) GetAgentTags() []string`

GetAgentTags returns the AgentTags field if non-nil, zero value otherwise.

### GetAgentTagsOk

`func (o *V1NetworkAgentsUpdateRequest) GetAgentTagsOk() (*[]string, bool)`

GetAgentTagsOk returns a tuple with the AgentTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentTags

`func (o *V1NetworkAgentsUpdateRequest) SetAgentTags(v []string)`

SetAgentTags sets AgentTags field to given value.

### HasAgentTags

`func (o *V1NetworkAgentsUpdateRequest) HasAgentTags() bool`

HasAgentTags returns a boolean if a field has been set.

### GetAgentName

`func (o *V1NetworkAgentsUpdateRequest) GetAgentName() string`

GetAgentName returns the AgentName field if non-nil, zero value otherwise.

### GetAgentNameOk

`func (o *V1NetworkAgentsUpdateRequest) GetAgentNameOk() (*string, bool)`

GetAgentNameOk returns a tuple with the AgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentName

`func (o *V1NetworkAgentsUpdateRequest) SetAgentName(v string)`

SetAgentName sets AgentName field to given value.

### HasAgentName

`func (o *V1NetworkAgentsUpdateRequest) HasAgentName() bool`

HasAgentName returns a boolean if a field has been set.

### GetAgentProviderId

`func (o *V1NetworkAgentsUpdateRequest) GetAgentProviderId() int32`

GetAgentProviderId returns the AgentProviderId field if non-nil, zero value otherwise.

### GetAgentProviderIdOk

`func (o *V1NetworkAgentsUpdateRequest) GetAgentProviderIdOk() (*int32, bool)`

GetAgentProviderIdOk returns a tuple with the AgentProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentProviderId

`func (o *V1NetworkAgentsUpdateRequest) SetAgentProviderId(v int32)`

SetAgentProviderId sets AgentProviderId field to given value.

### HasAgentProviderId

`func (o *V1NetworkAgentsUpdateRequest) HasAgentProviderId() bool`

HasAgentProviderId returns a boolean if a field has been set.

### GetNetwork

`func (o *V1NetworkAgentsUpdateRequest) GetNetwork() AgentInterfacesMetadata`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *V1NetworkAgentsUpdateRequest) GetNetworkOk() (*AgentInterfacesMetadata, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *V1NetworkAgentsUpdateRequest) SetNetwork(v AgentInterfacesMetadata)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *V1NetworkAgentsUpdateRequest) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


