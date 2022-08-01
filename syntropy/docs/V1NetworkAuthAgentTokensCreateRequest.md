# V1NetworkAuthAgentTokensCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentTokenName** | **string** |  | 
**AgentTokenDescription** | Pointer to **string** |  | [optional] 
**AgentTokenValidUntil** | Pointer to **time.Time** |  | [optional] 
**AgentTokenAllowedTagNames** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV1NetworkAuthAgentTokensCreateRequest

`func NewV1NetworkAuthAgentTokensCreateRequest(agentTokenName string, ) *V1NetworkAuthAgentTokensCreateRequest`

NewV1NetworkAuthAgentTokensCreateRequest instantiates a new V1NetworkAuthAgentTokensCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NetworkAuthAgentTokensCreateRequestWithDefaults

`func NewV1NetworkAuthAgentTokensCreateRequestWithDefaults() *V1NetworkAuthAgentTokensCreateRequest`

NewV1NetworkAuthAgentTokensCreateRequestWithDefaults instantiates a new V1NetworkAuthAgentTokensCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentTokenName

`func (o *V1NetworkAuthAgentTokensCreateRequest) GetAgentTokenName() string`

GetAgentTokenName returns the AgentTokenName field if non-nil, zero value otherwise.

### GetAgentTokenNameOk

`func (o *V1NetworkAuthAgentTokensCreateRequest) GetAgentTokenNameOk() (*string, bool)`

GetAgentTokenNameOk returns a tuple with the AgentTokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentTokenName

`func (o *V1NetworkAuthAgentTokensCreateRequest) SetAgentTokenName(v string)`

SetAgentTokenName sets AgentTokenName field to given value.


### GetAgentTokenDescription

`func (o *V1NetworkAuthAgentTokensCreateRequest) GetAgentTokenDescription() string`

GetAgentTokenDescription returns the AgentTokenDescription field if non-nil, zero value otherwise.

### GetAgentTokenDescriptionOk

`func (o *V1NetworkAuthAgentTokensCreateRequest) GetAgentTokenDescriptionOk() (*string, bool)`

GetAgentTokenDescriptionOk returns a tuple with the AgentTokenDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentTokenDescription

`func (o *V1NetworkAuthAgentTokensCreateRequest) SetAgentTokenDescription(v string)`

SetAgentTokenDescription sets AgentTokenDescription field to given value.

### HasAgentTokenDescription

`func (o *V1NetworkAuthAgentTokensCreateRequest) HasAgentTokenDescription() bool`

HasAgentTokenDescription returns a boolean if a field has been set.

### GetAgentTokenValidUntil

`func (o *V1NetworkAuthAgentTokensCreateRequest) GetAgentTokenValidUntil() time.Time`

GetAgentTokenValidUntil returns the AgentTokenValidUntil field if non-nil, zero value otherwise.

### GetAgentTokenValidUntilOk

`func (o *V1NetworkAuthAgentTokensCreateRequest) GetAgentTokenValidUntilOk() (*time.Time, bool)`

GetAgentTokenValidUntilOk returns a tuple with the AgentTokenValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentTokenValidUntil

`func (o *V1NetworkAuthAgentTokensCreateRequest) SetAgentTokenValidUntil(v time.Time)`

SetAgentTokenValidUntil sets AgentTokenValidUntil field to given value.

### HasAgentTokenValidUntil

`func (o *V1NetworkAuthAgentTokensCreateRequest) HasAgentTokenValidUntil() bool`

HasAgentTokenValidUntil returns a boolean if a field has been set.

### GetAgentTokenAllowedTagNames

`func (o *V1NetworkAuthAgentTokensCreateRequest) GetAgentTokenAllowedTagNames() []string`

GetAgentTokenAllowedTagNames returns the AgentTokenAllowedTagNames field if non-nil, zero value otherwise.

### GetAgentTokenAllowedTagNamesOk

`func (o *V1NetworkAuthAgentTokensCreateRequest) GetAgentTokenAllowedTagNamesOk() (*[]string, bool)`

GetAgentTokenAllowedTagNamesOk returns a tuple with the AgentTokenAllowedTagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentTokenAllowedTagNames

`func (o *V1NetworkAuthAgentTokensCreateRequest) SetAgentTokenAllowedTagNames(v []string)`

SetAgentTokenAllowedTagNames sets AgentTokenAllowedTagNames field to given value.

### HasAgentTokenAllowedTagNames

`func (o *V1NetworkAuthAgentTokensCreateRequest) HasAgentTokenAllowedTagNames() bool`

HasAgentTokenAllowedTagNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


