# V1AuthAgentToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentTokenName** | Pointer to **string** |  | [optional] 
**AgentTokenDescription** | Pointer to **NullableString** |  | [optional] 
**AgentTokenValidUntil** | Pointer to **time.Time** |  | [optional] 
**AgentTokenStatus** | Pointer to **bool** |  | [optional] 
**AgentTokenId** | Pointer to **int32** |  | [optional] 
**AgentTokenCreatedAt** | Pointer to **time.Time** |  | [optional] 
**AgentTokenUpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewV1AuthAgentToken

`func NewV1AuthAgentToken() *V1AuthAgentToken`

NewV1AuthAgentToken instantiates a new V1AuthAgentToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AuthAgentTokenWithDefaults

`func NewV1AuthAgentTokenWithDefaults() *V1AuthAgentToken`

NewV1AuthAgentTokenWithDefaults instantiates a new V1AuthAgentToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentTokenName

`func (o *V1AuthAgentToken) GetAgentTokenName() string`

GetAgentTokenName returns the AgentTokenName field if non-nil, zero value otherwise.

### GetAgentTokenNameOk

`func (o *V1AuthAgentToken) GetAgentTokenNameOk() (*string, bool)`

GetAgentTokenNameOk returns a tuple with the AgentTokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentTokenName

`func (o *V1AuthAgentToken) SetAgentTokenName(v string)`

SetAgentTokenName sets AgentTokenName field to given value.

### HasAgentTokenName

`func (o *V1AuthAgentToken) HasAgentTokenName() bool`

HasAgentTokenName returns a boolean if a field has been set.

### GetAgentTokenDescription

`func (o *V1AuthAgentToken) GetAgentTokenDescription() string`

GetAgentTokenDescription returns the AgentTokenDescription field if non-nil, zero value otherwise.

### GetAgentTokenDescriptionOk

`func (o *V1AuthAgentToken) GetAgentTokenDescriptionOk() (*string, bool)`

GetAgentTokenDescriptionOk returns a tuple with the AgentTokenDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentTokenDescription

`func (o *V1AuthAgentToken) SetAgentTokenDescription(v string)`

SetAgentTokenDescription sets AgentTokenDescription field to given value.

### HasAgentTokenDescription

`func (o *V1AuthAgentToken) HasAgentTokenDescription() bool`

HasAgentTokenDescription returns a boolean if a field has been set.

### SetAgentTokenDescriptionNil

`func (o *V1AuthAgentToken) SetAgentTokenDescriptionNil(b bool)`

 SetAgentTokenDescriptionNil sets the value for AgentTokenDescription to be an explicit nil

### UnsetAgentTokenDescription
`func (o *V1AuthAgentToken) UnsetAgentTokenDescription()`

UnsetAgentTokenDescription ensures that no value is present for AgentTokenDescription, not even an explicit nil
### GetAgentTokenValidUntil

`func (o *V1AuthAgentToken) GetAgentTokenValidUntil() time.Time`

GetAgentTokenValidUntil returns the AgentTokenValidUntil field if non-nil, zero value otherwise.

### GetAgentTokenValidUntilOk

`func (o *V1AuthAgentToken) GetAgentTokenValidUntilOk() (*time.Time, bool)`

GetAgentTokenValidUntilOk returns a tuple with the AgentTokenValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentTokenValidUntil

`func (o *V1AuthAgentToken) SetAgentTokenValidUntil(v time.Time)`

SetAgentTokenValidUntil sets AgentTokenValidUntil field to given value.

### HasAgentTokenValidUntil

`func (o *V1AuthAgentToken) HasAgentTokenValidUntil() bool`

HasAgentTokenValidUntil returns a boolean if a field has been set.

### GetAgentTokenStatus

`func (o *V1AuthAgentToken) GetAgentTokenStatus() bool`

GetAgentTokenStatus returns the AgentTokenStatus field if non-nil, zero value otherwise.

### GetAgentTokenStatusOk

`func (o *V1AuthAgentToken) GetAgentTokenStatusOk() (*bool, bool)`

GetAgentTokenStatusOk returns a tuple with the AgentTokenStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentTokenStatus

`func (o *V1AuthAgentToken) SetAgentTokenStatus(v bool)`

SetAgentTokenStatus sets AgentTokenStatus field to given value.

### HasAgentTokenStatus

`func (o *V1AuthAgentToken) HasAgentTokenStatus() bool`

HasAgentTokenStatus returns a boolean if a field has been set.

### GetAgentTokenId

`func (o *V1AuthAgentToken) GetAgentTokenId() int32`

GetAgentTokenId returns the AgentTokenId field if non-nil, zero value otherwise.

### GetAgentTokenIdOk

`func (o *V1AuthAgentToken) GetAgentTokenIdOk() (*int32, bool)`

GetAgentTokenIdOk returns a tuple with the AgentTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentTokenId

`func (o *V1AuthAgentToken) SetAgentTokenId(v int32)`

SetAgentTokenId sets AgentTokenId field to given value.

### HasAgentTokenId

`func (o *V1AuthAgentToken) HasAgentTokenId() bool`

HasAgentTokenId returns a boolean if a field has been set.

### GetAgentTokenCreatedAt

`func (o *V1AuthAgentToken) GetAgentTokenCreatedAt() time.Time`

GetAgentTokenCreatedAt returns the AgentTokenCreatedAt field if non-nil, zero value otherwise.

### GetAgentTokenCreatedAtOk

`func (o *V1AuthAgentToken) GetAgentTokenCreatedAtOk() (*time.Time, bool)`

GetAgentTokenCreatedAtOk returns a tuple with the AgentTokenCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentTokenCreatedAt

`func (o *V1AuthAgentToken) SetAgentTokenCreatedAt(v time.Time)`

SetAgentTokenCreatedAt sets AgentTokenCreatedAt field to given value.

### HasAgentTokenCreatedAt

`func (o *V1AuthAgentToken) HasAgentTokenCreatedAt() bool`

HasAgentTokenCreatedAt returns a boolean if a field has been set.

### GetAgentTokenUpdatedAt

`func (o *V1AuthAgentToken) GetAgentTokenUpdatedAt() time.Time`

GetAgentTokenUpdatedAt returns the AgentTokenUpdatedAt field if non-nil, zero value otherwise.

### GetAgentTokenUpdatedAtOk

`func (o *V1AuthAgentToken) GetAgentTokenUpdatedAtOk() (*time.Time, bool)`

GetAgentTokenUpdatedAtOk returns a tuple with the AgentTokenUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentTokenUpdatedAt

`func (o *V1AuthAgentToken) SetAgentTokenUpdatedAt(v time.Time)`

SetAgentTokenUpdatedAt sets AgentTokenUpdatedAt field to given value.

### HasAgentTokenUpdatedAt

`func (o *V1AuthAgentToken) HasAgentTokenUpdatedAt() bool`

HasAgentTokenUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


