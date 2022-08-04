# V1Agent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | **int32** |  | 
**AgentPublicIpv4** | **string** |  | 
**AgentLocationCity** | **NullableString** |  | 
**AgentDeviceId** | **string** |  | 
**AgentName** | **string** |  | 
**AgentStatus** | [**NullableAgentStatus**](AgentStatus.md) |  | 
**AgentVersion** | **string** |  | 
**AgentLockedFields** | [**AgentLockedFields**](AgentLockedFields.md) |  | 
**AgentModifiedAt** | **time.Time** |  | 
**AgentIsVirtual** | **bool** |  | 
**AgentType** | [**AgentType**](AgentType.md) |  | 
**AgentProvider** | [**AgentProviderNameAndId**](AgentProviderNameAndId.md) |  | 
**AgentTags** | [**[]AgentTag**](AgentTag.md) |  | 
**AgentServicesSubnetsEnabledCount** | **int32** |  | 
**AgentServicesSubnetsCount** | **int32** |  | 
**AgentLocationCountry** | **NullableString** |  | 
**AgentIsOnline** | **bool** |  | 

## Methods

### NewV1Agent

`func NewV1Agent(agentId int32, agentPublicIpv4 string, agentLocationCity NullableString, agentDeviceId string, agentName string, agentStatus NullableAgentStatus, agentVersion string, agentLockedFields AgentLockedFields, agentModifiedAt time.Time, agentIsVirtual bool, agentType AgentType, agentProvider AgentProviderNameAndId, agentTags []AgentTag, agentServicesSubnetsEnabledCount int32, agentServicesSubnetsCount int32, agentLocationCountry NullableString, agentIsOnline bool, ) *V1Agent`

NewV1Agent instantiates a new V1Agent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AgentWithDefaults

`func NewV1AgentWithDefaults() *V1Agent`

NewV1AgentWithDefaults instantiates a new V1Agent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *V1Agent) GetAgentId() int32`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *V1Agent) GetAgentIdOk() (*int32, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *V1Agent) SetAgentId(v int32)`

SetAgentId sets AgentId field to given value.


### GetAgentPublicIpv4

`func (o *V1Agent) GetAgentPublicIpv4() string`

GetAgentPublicIpv4 returns the AgentPublicIpv4 field if non-nil, zero value otherwise.

### GetAgentPublicIpv4Ok

`func (o *V1Agent) GetAgentPublicIpv4Ok() (*string, bool)`

GetAgentPublicIpv4Ok returns a tuple with the AgentPublicIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPublicIpv4

`func (o *V1Agent) SetAgentPublicIpv4(v string)`

SetAgentPublicIpv4 sets AgentPublicIpv4 field to given value.


### GetAgentLocationCity

`func (o *V1Agent) GetAgentLocationCity() string`

GetAgentLocationCity returns the AgentLocationCity field if non-nil, zero value otherwise.

### GetAgentLocationCityOk

`func (o *V1Agent) GetAgentLocationCityOk() (*string, bool)`

GetAgentLocationCityOk returns a tuple with the AgentLocationCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentLocationCity

`func (o *V1Agent) SetAgentLocationCity(v string)`

SetAgentLocationCity sets AgentLocationCity field to given value.


### SetAgentLocationCityNil

`func (o *V1Agent) SetAgentLocationCityNil(b bool)`

 SetAgentLocationCityNil sets the value for AgentLocationCity to be an explicit nil

### UnsetAgentLocationCity
`func (o *V1Agent) UnsetAgentLocationCity()`

UnsetAgentLocationCity ensures that no value is present for AgentLocationCity, not even an explicit nil
### GetAgentDeviceId

`func (o *V1Agent) GetAgentDeviceId() string`

GetAgentDeviceId returns the AgentDeviceId field if non-nil, zero value otherwise.

### GetAgentDeviceIdOk

`func (o *V1Agent) GetAgentDeviceIdOk() (*string, bool)`

GetAgentDeviceIdOk returns a tuple with the AgentDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentDeviceId

`func (o *V1Agent) SetAgentDeviceId(v string)`

SetAgentDeviceId sets AgentDeviceId field to given value.


### GetAgentName

`func (o *V1Agent) GetAgentName() string`

GetAgentName returns the AgentName field if non-nil, zero value otherwise.

### GetAgentNameOk

`func (o *V1Agent) GetAgentNameOk() (*string, bool)`

GetAgentNameOk returns a tuple with the AgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentName

`func (o *V1Agent) SetAgentName(v string)`

SetAgentName sets AgentName field to given value.


### GetAgentStatus

`func (o *V1Agent) GetAgentStatus() AgentStatus`

GetAgentStatus returns the AgentStatus field if non-nil, zero value otherwise.

### GetAgentStatusOk

`func (o *V1Agent) GetAgentStatusOk() (*AgentStatus, bool)`

GetAgentStatusOk returns a tuple with the AgentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentStatus

`func (o *V1Agent) SetAgentStatus(v AgentStatus)`

SetAgentStatus sets AgentStatus field to given value.


### SetAgentStatusNil

`func (o *V1Agent) SetAgentStatusNil(b bool)`

 SetAgentStatusNil sets the value for AgentStatus to be an explicit nil

### UnsetAgentStatus
`func (o *V1Agent) UnsetAgentStatus()`

UnsetAgentStatus ensures that no value is present for AgentStatus, not even an explicit nil
### GetAgentVersion

`func (o *V1Agent) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *V1Agent) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *V1Agent) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.


### GetAgentLockedFields

`func (o *V1Agent) GetAgentLockedFields() AgentLockedFields`

GetAgentLockedFields returns the AgentLockedFields field if non-nil, zero value otherwise.

### GetAgentLockedFieldsOk

`func (o *V1Agent) GetAgentLockedFieldsOk() (*AgentLockedFields, bool)`

GetAgentLockedFieldsOk returns a tuple with the AgentLockedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentLockedFields

`func (o *V1Agent) SetAgentLockedFields(v AgentLockedFields)`

SetAgentLockedFields sets AgentLockedFields field to given value.


### GetAgentModifiedAt

`func (o *V1Agent) GetAgentModifiedAt() time.Time`

GetAgentModifiedAt returns the AgentModifiedAt field if non-nil, zero value otherwise.

### GetAgentModifiedAtOk

`func (o *V1Agent) GetAgentModifiedAtOk() (*time.Time, bool)`

GetAgentModifiedAtOk returns a tuple with the AgentModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentModifiedAt

`func (o *V1Agent) SetAgentModifiedAt(v time.Time)`

SetAgentModifiedAt sets AgentModifiedAt field to given value.


### GetAgentIsVirtual

`func (o *V1Agent) GetAgentIsVirtual() bool`

GetAgentIsVirtual returns the AgentIsVirtual field if non-nil, zero value otherwise.

### GetAgentIsVirtualOk

`func (o *V1Agent) GetAgentIsVirtualOk() (*bool, bool)`

GetAgentIsVirtualOk returns a tuple with the AgentIsVirtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentIsVirtual

`func (o *V1Agent) SetAgentIsVirtual(v bool)`

SetAgentIsVirtual sets AgentIsVirtual field to given value.


### GetAgentType

`func (o *V1Agent) GetAgentType() AgentType`

GetAgentType returns the AgentType field if non-nil, zero value otherwise.

### GetAgentTypeOk

`func (o *V1Agent) GetAgentTypeOk() (*AgentType, bool)`

GetAgentTypeOk returns a tuple with the AgentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentType

`func (o *V1Agent) SetAgentType(v AgentType)`

SetAgentType sets AgentType field to given value.


### GetAgentProvider

`func (o *V1Agent) GetAgentProvider() AgentProviderNameAndId`

GetAgentProvider returns the AgentProvider field if non-nil, zero value otherwise.

### GetAgentProviderOk

`func (o *V1Agent) GetAgentProviderOk() (*AgentProviderNameAndId, bool)`

GetAgentProviderOk returns a tuple with the AgentProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentProvider

`func (o *V1Agent) SetAgentProvider(v AgentProviderNameAndId)`

SetAgentProvider sets AgentProvider field to given value.


### GetAgentTags

`func (o *V1Agent) GetAgentTags() []AgentTag`

GetAgentTags returns the AgentTags field if non-nil, zero value otherwise.

### GetAgentTagsOk

`func (o *V1Agent) GetAgentTagsOk() (*[]AgentTag, bool)`

GetAgentTagsOk returns a tuple with the AgentTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentTags

`func (o *V1Agent) SetAgentTags(v []AgentTag)`

SetAgentTags sets AgentTags field to given value.


### GetAgentServicesSubnetsEnabledCount

`func (o *V1Agent) GetAgentServicesSubnetsEnabledCount() int32`

GetAgentServicesSubnetsEnabledCount returns the AgentServicesSubnetsEnabledCount field if non-nil, zero value otherwise.

### GetAgentServicesSubnetsEnabledCountOk

`func (o *V1Agent) GetAgentServicesSubnetsEnabledCountOk() (*int32, bool)`

GetAgentServicesSubnetsEnabledCountOk returns a tuple with the AgentServicesSubnetsEnabledCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentServicesSubnetsEnabledCount

`func (o *V1Agent) SetAgentServicesSubnetsEnabledCount(v int32)`

SetAgentServicesSubnetsEnabledCount sets AgentServicesSubnetsEnabledCount field to given value.


### GetAgentServicesSubnetsCount

`func (o *V1Agent) GetAgentServicesSubnetsCount() int32`

GetAgentServicesSubnetsCount returns the AgentServicesSubnetsCount field if non-nil, zero value otherwise.

### GetAgentServicesSubnetsCountOk

`func (o *V1Agent) GetAgentServicesSubnetsCountOk() (*int32, bool)`

GetAgentServicesSubnetsCountOk returns a tuple with the AgentServicesSubnetsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentServicesSubnetsCount

`func (o *V1Agent) SetAgentServicesSubnetsCount(v int32)`

SetAgentServicesSubnetsCount sets AgentServicesSubnetsCount field to given value.


### GetAgentLocationCountry

`func (o *V1Agent) GetAgentLocationCountry() string`

GetAgentLocationCountry returns the AgentLocationCountry field if non-nil, zero value otherwise.

### GetAgentLocationCountryOk

`func (o *V1Agent) GetAgentLocationCountryOk() (*string, bool)`

GetAgentLocationCountryOk returns a tuple with the AgentLocationCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentLocationCountry

`func (o *V1Agent) SetAgentLocationCountry(v string)`

SetAgentLocationCountry sets AgentLocationCountry field to given value.


### SetAgentLocationCountryNil

`func (o *V1Agent) SetAgentLocationCountryNil(b bool)`

 SetAgentLocationCountryNil sets the value for AgentLocationCountry to be an explicit nil

### UnsetAgentLocationCountry
`func (o *V1Agent) UnsetAgentLocationCountry()`

UnsetAgentLocationCountry ensures that no value is present for AgentLocationCountry, not even an explicit nil
### GetAgentIsOnline

`func (o *V1Agent) GetAgentIsOnline() bool`

GetAgentIsOnline returns the AgentIsOnline field if non-nil, zero value otherwise.

### GetAgentIsOnlineOk

`func (o *V1Agent) GetAgentIsOnlineOk() (*bool, bool)`

GetAgentIsOnlineOk returns a tuple with the AgentIsOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentIsOnline

`func (o *V1Agent) SetAgentIsOnline(v bool)`

SetAgentIsOnline sets AgentIsOnline field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


