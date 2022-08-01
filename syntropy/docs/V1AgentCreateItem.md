# V1AgentCreateItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentServicesDefaultStatus** | **bool** |  | 
**ApiKeyId** | **int32** |  | 
**AgentPublicIpv4** | **string** |  | 
**AgentLocationLat** | **NullableFloat32** |  | 
**AgentLocationLon** | **NullableFloat32** |  | 
**AgentLocationCity** | **NullableString** |  | 
**AgentLocationCountry** | **NullableString** |  | 
**AgentDeviceId** | **string** |  | 
**AgentName** | **string** |  | 
**AgentStatus** | [**NullableAgentStatus**](AgentStatus.md) |  | 
**AgentVersion** | **NullableString** |  | 
**AgentProviderId** | Pointer to **NullableInt32** |  | [optional] 
**AgentLockedFields** | [**AgentLockedFields**](AgentLockedFields.md) |  | 
**AgentModifiedAt** | **NullableTime** |  | 
**AgentIsVirtual** | **bool** |  | 
**AgentType** | [**AgentType**](AgentType.md) |  | 
**AgentInternalSubnet** | **string** |  | 
**AgentCreatedAt** | **time.Time** |  | 
**AgentUpdatedAt** | **time.Time** |  | 
**AgentId** | **int32** |  | 

## Methods

### NewV1AgentCreateItem

`func NewV1AgentCreateItem(agentServicesDefaultStatus bool, apiKeyId int32, agentPublicIpv4 string, agentLocationLat NullableFloat32, agentLocationLon NullableFloat32, agentLocationCity NullableString, agentLocationCountry NullableString, agentDeviceId string, agentName string, agentStatus NullableAgentStatus, agentVersion NullableString, agentLockedFields AgentLockedFields, agentModifiedAt NullableTime, agentIsVirtual bool, agentType AgentType, agentInternalSubnet string, agentCreatedAt time.Time, agentUpdatedAt time.Time, agentId int32, ) *V1AgentCreateItem`

NewV1AgentCreateItem instantiates a new V1AgentCreateItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AgentCreateItemWithDefaults

`func NewV1AgentCreateItemWithDefaults() *V1AgentCreateItem`

NewV1AgentCreateItemWithDefaults instantiates a new V1AgentCreateItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentServicesDefaultStatus

`func (o *V1AgentCreateItem) GetAgentServicesDefaultStatus() bool`

GetAgentServicesDefaultStatus returns the AgentServicesDefaultStatus field if non-nil, zero value otherwise.

### GetAgentServicesDefaultStatusOk

`func (o *V1AgentCreateItem) GetAgentServicesDefaultStatusOk() (*bool, bool)`

GetAgentServicesDefaultStatusOk returns a tuple with the AgentServicesDefaultStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentServicesDefaultStatus

`func (o *V1AgentCreateItem) SetAgentServicesDefaultStatus(v bool)`

SetAgentServicesDefaultStatus sets AgentServicesDefaultStatus field to given value.


### GetApiKeyId

`func (o *V1AgentCreateItem) GetApiKeyId() int32`

GetApiKeyId returns the ApiKeyId field if non-nil, zero value otherwise.

### GetApiKeyIdOk

`func (o *V1AgentCreateItem) GetApiKeyIdOk() (*int32, bool)`

GetApiKeyIdOk returns a tuple with the ApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyId

`func (o *V1AgentCreateItem) SetApiKeyId(v int32)`

SetApiKeyId sets ApiKeyId field to given value.


### GetAgentPublicIpv4

`func (o *V1AgentCreateItem) GetAgentPublicIpv4() string`

GetAgentPublicIpv4 returns the AgentPublicIpv4 field if non-nil, zero value otherwise.

### GetAgentPublicIpv4Ok

`func (o *V1AgentCreateItem) GetAgentPublicIpv4Ok() (*string, bool)`

GetAgentPublicIpv4Ok returns a tuple with the AgentPublicIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPublicIpv4

`func (o *V1AgentCreateItem) SetAgentPublicIpv4(v string)`

SetAgentPublicIpv4 sets AgentPublicIpv4 field to given value.


### GetAgentLocationLat

`func (o *V1AgentCreateItem) GetAgentLocationLat() float32`

GetAgentLocationLat returns the AgentLocationLat field if non-nil, zero value otherwise.

### GetAgentLocationLatOk

`func (o *V1AgentCreateItem) GetAgentLocationLatOk() (*float32, bool)`

GetAgentLocationLatOk returns a tuple with the AgentLocationLat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentLocationLat

`func (o *V1AgentCreateItem) SetAgentLocationLat(v float32)`

SetAgentLocationLat sets AgentLocationLat field to given value.


### SetAgentLocationLatNil

`func (o *V1AgentCreateItem) SetAgentLocationLatNil(b bool)`

 SetAgentLocationLatNil sets the value for AgentLocationLat to be an explicit nil

### UnsetAgentLocationLat
`func (o *V1AgentCreateItem) UnsetAgentLocationLat()`

UnsetAgentLocationLat ensures that no value is present for AgentLocationLat, not even an explicit nil
### GetAgentLocationLon

`func (o *V1AgentCreateItem) GetAgentLocationLon() float32`

GetAgentLocationLon returns the AgentLocationLon field if non-nil, zero value otherwise.

### GetAgentLocationLonOk

`func (o *V1AgentCreateItem) GetAgentLocationLonOk() (*float32, bool)`

GetAgentLocationLonOk returns a tuple with the AgentLocationLon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentLocationLon

`func (o *V1AgentCreateItem) SetAgentLocationLon(v float32)`

SetAgentLocationLon sets AgentLocationLon field to given value.


### SetAgentLocationLonNil

`func (o *V1AgentCreateItem) SetAgentLocationLonNil(b bool)`

 SetAgentLocationLonNil sets the value for AgentLocationLon to be an explicit nil

### UnsetAgentLocationLon
`func (o *V1AgentCreateItem) UnsetAgentLocationLon()`

UnsetAgentLocationLon ensures that no value is present for AgentLocationLon, not even an explicit nil
### GetAgentLocationCity

`func (o *V1AgentCreateItem) GetAgentLocationCity() string`

GetAgentLocationCity returns the AgentLocationCity field if non-nil, zero value otherwise.

### GetAgentLocationCityOk

`func (o *V1AgentCreateItem) GetAgentLocationCityOk() (*string, bool)`

GetAgentLocationCityOk returns a tuple with the AgentLocationCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentLocationCity

`func (o *V1AgentCreateItem) SetAgentLocationCity(v string)`

SetAgentLocationCity sets AgentLocationCity field to given value.


### SetAgentLocationCityNil

`func (o *V1AgentCreateItem) SetAgentLocationCityNil(b bool)`

 SetAgentLocationCityNil sets the value for AgentLocationCity to be an explicit nil

### UnsetAgentLocationCity
`func (o *V1AgentCreateItem) UnsetAgentLocationCity()`

UnsetAgentLocationCity ensures that no value is present for AgentLocationCity, not even an explicit nil
### GetAgentLocationCountry

`func (o *V1AgentCreateItem) GetAgentLocationCountry() string`

GetAgentLocationCountry returns the AgentLocationCountry field if non-nil, zero value otherwise.

### GetAgentLocationCountryOk

`func (o *V1AgentCreateItem) GetAgentLocationCountryOk() (*string, bool)`

GetAgentLocationCountryOk returns a tuple with the AgentLocationCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentLocationCountry

`func (o *V1AgentCreateItem) SetAgentLocationCountry(v string)`

SetAgentLocationCountry sets AgentLocationCountry field to given value.


### SetAgentLocationCountryNil

`func (o *V1AgentCreateItem) SetAgentLocationCountryNil(b bool)`

 SetAgentLocationCountryNil sets the value for AgentLocationCountry to be an explicit nil

### UnsetAgentLocationCountry
`func (o *V1AgentCreateItem) UnsetAgentLocationCountry()`

UnsetAgentLocationCountry ensures that no value is present for AgentLocationCountry, not even an explicit nil
### GetAgentDeviceId

`func (o *V1AgentCreateItem) GetAgentDeviceId() string`

GetAgentDeviceId returns the AgentDeviceId field if non-nil, zero value otherwise.

### GetAgentDeviceIdOk

`func (o *V1AgentCreateItem) GetAgentDeviceIdOk() (*string, bool)`

GetAgentDeviceIdOk returns a tuple with the AgentDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentDeviceId

`func (o *V1AgentCreateItem) SetAgentDeviceId(v string)`

SetAgentDeviceId sets AgentDeviceId field to given value.


### GetAgentName

`func (o *V1AgentCreateItem) GetAgentName() string`

GetAgentName returns the AgentName field if non-nil, zero value otherwise.

### GetAgentNameOk

`func (o *V1AgentCreateItem) GetAgentNameOk() (*string, bool)`

GetAgentNameOk returns a tuple with the AgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentName

`func (o *V1AgentCreateItem) SetAgentName(v string)`

SetAgentName sets AgentName field to given value.


### GetAgentStatus

`func (o *V1AgentCreateItem) GetAgentStatus() AgentStatus`

GetAgentStatus returns the AgentStatus field if non-nil, zero value otherwise.

### GetAgentStatusOk

`func (o *V1AgentCreateItem) GetAgentStatusOk() (*AgentStatus, bool)`

GetAgentStatusOk returns a tuple with the AgentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentStatus

`func (o *V1AgentCreateItem) SetAgentStatus(v AgentStatus)`

SetAgentStatus sets AgentStatus field to given value.


### SetAgentStatusNil

`func (o *V1AgentCreateItem) SetAgentStatusNil(b bool)`

 SetAgentStatusNil sets the value for AgentStatus to be an explicit nil

### UnsetAgentStatus
`func (o *V1AgentCreateItem) UnsetAgentStatus()`

UnsetAgentStatus ensures that no value is present for AgentStatus, not even an explicit nil
### GetAgentVersion

`func (o *V1AgentCreateItem) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *V1AgentCreateItem) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *V1AgentCreateItem) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.


### SetAgentVersionNil

`func (o *V1AgentCreateItem) SetAgentVersionNil(b bool)`

 SetAgentVersionNil sets the value for AgentVersion to be an explicit nil

### UnsetAgentVersion
`func (o *V1AgentCreateItem) UnsetAgentVersion()`

UnsetAgentVersion ensures that no value is present for AgentVersion, not even an explicit nil
### GetAgentProviderId

`func (o *V1AgentCreateItem) GetAgentProviderId() int32`

GetAgentProviderId returns the AgentProviderId field if non-nil, zero value otherwise.

### GetAgentProviderIdOk

`func (o *V1AgentCreateItem) GetAgentProviderIdOk() (*int32, bool)`

GetAgentProviderIdOk returns a tuple with the AgentProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentProviderId

`func (o *V1AgentCreateItem) SetAgentProviderId(v int32)`

SetAgentProviderId sets AgentProviderId field to given value.

### HasAgentProviderId

`func (o *V1AgentCreateItem) HasAgentProviderId() bool`

HasAgentProviderId returns a boolean if a field has been set.

### SetAgentProviderIdNil

`func (o *V1AgentCreateItem) SetAgentProviderIdNil(b bool)`

 SetAgentProviderIdNil sets the value for AgentProviderId to be an explicit nil

### UnsetAgentProviderId
`func (o *V1AgentCreateItem) UnsetAgentProviderId()`

UnsetAgentProviderId ensures that no value is present for AgentProviderId, not even an explicit nil
### GetAgentLockedFields

`func (o *V1AgentCreateItem) GetAgentLockedFields() AgentLockedFields`

GetAgentLockedFields returns the AgentLockedFields field if non-nil, zero value otherwise.

### GetAgentLockedFieldsOk

`func (o *V1AgentCreateItem) GetAgentLockedFieldsOk() (*AgentLockedFields, bool)`

GetAgentLockedFieldsOk returns a tuple with the AgentLockedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentLockedFields

`func (o *V1AgentCreateItem) SetAgentLockedFields(v AgentLockedFields)`

SetAgentLockedFields sets AgentLockedFields field to given value.


### GetAgentModifiedAt

`func (o *V1AgentCreateItem) GetAgentModifiedAt() time.Time`

GetAgentModifiedAt returns the AgentModifiedAt field if non-nil, zero value otherwise.

### GetAgentModifiedAtOk

`func (o *V1AgentCreateItem) GetAgentModifiedAtOk() (*time.Time, bool)`

GetAgentModifiedAtOk returns a tuple with the AgentModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentModifiedAt

`func (o *V1AgentCreateItem) SetAgentModifiedAt(v time.Time)`

SetAgentModifiedAt sets AgentModifiedAt field to given value.


### SetAgentModifiedAtNil

`func (o *V1AgentCreateItem) SetAgentModifiedAtNil(b bool)`

 SetAgentModifiedAtNil sets the value for AgentModifiedAt to be an explicit nil

### UnsetAgentModifiedAt
`func (o *V1AgentCreateItem) UnsetAgentModifiedAt()`

UnsetAgentModifiedAt ensures that no value is present for AgentModifiedAt, not even an explicit nil
### GetAgentIsVirtual

`func (o *V1AgentCreateItem) GetAgentIsVirtual() bool`

GetAgentIsVirtual returns the AgentIsVirtual field if non-nil, zero value otherwise.

### GetAgentIsVirtualOk

`func (o *V1AgentCreateItem) GetAgentIsVirtualOk() (*bool, bool)`

GetAgentIsVirtualOk returns a tuple with the AgentIsVirtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentIsVirtual

`func (o *V1AgentCreateItem) SetAgentIsVirtual(v bool)`

SetAgentIsVirtual sets AgentIsVirtual field to given value.


### GetAgentType

`func (o *V1AgentCreateItem) GetAgentType() AgentType`

GetAgentType returns the AgentType field if non-nil, zero value otherwise.

### GetAgentTypeOk

`func (o *V1AgentCreateItem) GetAgentTypeOk() (*AgentType, bool)`

GetAgentTypeOk returns a tuple with the AgentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentType

`func (o *V1AgentCreateItem) SetAgentType(v AgentType)`

SetAgentType sets AgentType field to given value.


### GetAgentInternalSubnet

`func (o *V1AgentCreateItem) GetAgentInternalSubnet() string`

GetAgentInternalSubnet returns the AgentInternalSubnet field if non-nil, zero value otherwise.

### GetAgentInternalSubnetOk

`func (o *V1AgentCreateItem) GetAgentInternalSubnetOk() (*string, bool)`

GetAgentInternalSubnetOk returns a tuple with the AgentInternalSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentInternalSubnet

`func (o *V1AgentCreateItem) SetAgentInternalSubnet(v string)`

SetAgentInternalSubnet sets AgentInternalSubnet field to given value.


### GetAgentCreatedAt

`func (o *V1AgentCreateItem) GetAgentCreatedAt() time.Time`

GetAgentCreatedAt returns the AgentCreatedAt field if non-nil, zero value otherwise.

### GetAgentCreatedAtOk

`func (o *V1AgentCreateItem) GetAgentCreatedAtOk() (*time.Time, bool)`

GetAgentCreatedAtOk returns a tuple with the AgentCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentCreatedAt

`func (o *V1AgentCreateItem) SetAgentCreatedAt(v time.Time)`

SetAgentCreatedAt sets AgentCreatedAt field to given value.


### GetAgentUpdatedAt

`func (o *V1AgentCreateItem) GetAgentUpdatedAt() time.Time`

GetAgentUpdatedAt returns the AgentUpdatedAt field if non-nil, zero value otherwise.

### GetAgentUpdatedAtOk

`func (o *V1AgentCreateItem) GetAgentUpdatedAtOk() (*time.Time, bool)`

GetAgentUpdatedAtOk returns a tuple with the AgentUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentUpdatedAt

`func (o *V1AgentCreateItem) SetAgentUpdatedAt(v time.Time)`

SetAgentUpdatedAt sets AgentUpdatedAt field to given value.


### GetAgentId

`func (o *V1AgentCreateItem) GetAgentId() int32`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *V1AgentCreateItem) GetAgentIdOk() (*int32, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *V1AgentCreateItem) SetAgentId(v int32)`

SetAgentId sets AgentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


