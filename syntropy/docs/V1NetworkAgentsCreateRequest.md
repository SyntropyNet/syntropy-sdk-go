# V1NetworkAgentsCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentTags** | Pointer to **[]string** |  | [optional] 
**AgentToken** | **string** |  | 
**AgentName** | **string** |  | 
**AgentProviderId** | Pointer to **int32** |  | [optional] 

## Methods

### NewV1NetworkAgentsCreateRequest

`func NewV1NetworkAgentsCreateRequest(agentToken string, agentName string, ) *V1NetworkAgentsCreateRequest`

NewV1NetworkAgentsCreateRequest instantiates a new V1NetworkAgentsCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NetworkAgentsCreateRequestWithDefaults

`func NewV1NetworkAgentsCreateRequestWithDefaults() *V1NetworkAgentsCreateRequest`

NewV1NetworkAgentsCreateRequestWithDefaults instantiates a new V1NetworkAgentsCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentTags

`func (o *V1NetworkAgentsCreateRequest) GetAgentTags() []string`

GetAgentTags returns the AgentTags field if non-nil, zero value otherwise.

### GetAgentTagsOk

`func (o *V1NetworkAgentsCreateRequest) GetAgentTagsOk() (*[]string, bool)`

GetAgentTagsOk returns a tuple with the AgentTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentTags

`func (o *V1NetworkAgentsCreateRequest) SetAgentTags(v []string)`

SetAgentTags sets AgentTags field to given value.

### HasAgentTags

`func (o *V1NetworkAgentsCreateRequest) HasAgentTags() bool`

HasAgentTags returns a boolean if a field has been set.

### GetAgentToken

`func (o *V1NetworkAgentsCreateRequest) GetAgentToken() string`

GetAgentToken returns the AgentToken field if non-nil, zero value otherwise.

### GetAgentTokenOk

`func (o *V1NetworkAgentsCreateRequest) GetAgentTokenOk() (*string, bool)`

GetAgentTokenOk returns a tuple with the AgentToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentToken

`func (o *V1NetworkAgentsCreateRequest) SetAgentToken(v string)`

SetAgentToken sets AgentToken field to given value.


### GetAgentName

`func (o *V1NetworkAgentsCreateRequest) GetAgentName() string`

GetAgentName returns the AgentName field if non-nil, zero value otherwise.

### GetAgentNameOk

`func (o *V1NetworkAgentsCreateRequest) GetAgentNameOk() (*string, bool)`

GetAgentNameOk returns a tuple with the AgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentName

`func (o *V1NetworkAgentsCreateRequest) SetAgentName(v string)`

SetAgentName sets AgentName field to given value.


### GetAgentProviderId

`func (o *V1NetworkAgentsCreateRequest) GetAgentProviderId() int32`

GetAgentProviderId returns the AgentProviderId field if non-nil, zero value otherwise.

### GetAgentProviderIdOk

`func (o *V1NetworkAgentsCreateRequest) GetAgentProviderIdOk() (*int32, bool)`

GetAgentProviderIdOk returns a tuple with the AgentProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentProviderId

`func (o *V1NetworkAgentsCreateRequest) SetAgentProviderId(v int32)`

SetAgentProviderId sets AgentProviderId field to given value.

### HasAgentProviderId

`func (o *V1NetworkAgentsCreateRequest) HasAgentProviderId() bool`

HasAgentProviderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


