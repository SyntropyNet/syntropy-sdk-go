# V1AgentFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | Pointer to **[]int32** |  | [optional] 
**AgentProviderId** | Pointer to **[]int32** |  | [optional] 
**AgentTagId** | Pointer to **[]int32** |  | [optional] 
**AgentType** | Pointer to [**[]AgentType**](AgentType.md) |  | [optional] 
**AgentVersion** | Pointer to **[]string** |  | [optional] 
**AgentTagName** | Pointer to **[]string** |  | [optional] 
**AgentStatus** | Pointer to [**[]AgentFilterAgentStatus**](AgentFilterAgentStatus.md) |  | [optional] 
**AgentLocationCountry** | Pointer to **[]string** |  | [optional] 
**AgentModifiedAtFrom** | Pointer to **time.Time** |  | [optional] 
**AgentModifiedAtTo** | Pointer to **time.Time** |  | [optional] 
**AgentName** | Pointer to **string** |  | [optional] 

## Methods

### NewV1AgentFilter

`func NewV1AgentFilter() *V1AgentFilter`

NewV1AgentFilter instantiates a new V1AgentFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AgentFilterWithDefaults

`func NewV1AgentFilterWithDefaults() *V1AgentFilter`

NewV1AgentFilterWithDefaults instantiates a new V1AgentFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *V1AgentFilter) GetAgentId() []int32`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *V1AgentFilter) GetAgentIdOk() (*[]int32, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *V1AgentFilter) SetAgentId(v []int32)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *V1AgentFilter) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetAgentProviderId

`func (o *V1AgentFilter) GetAgentProviderId() []int32`

GetAgentProviderId returns the AgentProviderId field if non-nil, zero value otherwise.

### GetAgentProviderIdOk

`func (o *V1AgentFilter) GetAgentProviderIdOk() (*[]int32, bool)`

GetAgentProviderIdOk returns a tuple with the AgentProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentProviderId

`func (o *V1AgentFilter) SetAgentProviderId(v []int32)`

SetAgentProviderId sets AgentProviderId field to given value.

### HasAgentProviderId

`func (o *V1AgentFilter) HasAgentProviderId() bool`

HasAgentProviderId returns a boolean if a field has been set.

### GetAgentTagId

`func (o *V1AgentFilter) GetAgentTagId() []int32`

GetAgentTagId returns the AgentTagId field if non-nil, zero value otherwise.

### GetAgentTagIdOk

`func (o *V1AgentFilter) GetAgentTagIdOk() (*[]int32, bool)`

GetAgentTagIdOk returns a tuple with the AgentTagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentTagId

`func (o *V1AgentFilter) SetAgentTagId(v []int32)`

SetAgentTagId sets AgentTagId field to given value.

### HasAgentTagId

`func (o *V1AgentFilter) HasAgentTagId() bool`

HasAgentTagId returns a boolean if a field has been set.

### GetAgentType

`func (o *V1AgentFilter) GetAgentType() []AgentType`

GetAgentType returns the AgentType field if non-nil, zero value otherwise.

### GetAgentTypeOk

`func (o *V1AgentFilter) GetAgentTypeOk() (*[]AgentType, bool)`

GetAgentTypeOk returns a tuple with the AgentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentType

`func (o *V1AgentFilter) SetAgentType(v []AgentType)`

SetAgentType sets AgentType field to given value.

### HasAgentType

`func (o *V1AgentFilter) HasAgentType() bool`

HasAgentType returns a boolean if a field has been set.

### GetAgentVersion

`func (o *V1AgentFilter) GetAgentVersion() []string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *V1AgentFilter) GetAgentVersionOk() (*[]string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *V1AgentFilter) SetAgentVersion(v []string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *V1AgentFilter) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### GetAgentTagName

`func (o *V1AgentFilter) GetAgentTagName() []string`

GetAgentTagName returns the AgentTagName field if non-nil, zero value otherwise.

### GetAgentTagNameOk

`func (o *V1AgentFilter) GetAgentTagNameOk() (*[]string, bool)`

GetAgentTagNameOk returns a tuple with the AgentTagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentTagName

`func (o *V1AgentFilter) SetAgentTagName(v []string)`

SetAgentTagName sets AgentTagName field to given value.

### HasAgentTagName

`func (o *V1AgentFilter) HasAgentTagName() bool`

HasAgentTagName returns a boolean if a field has been set.

### GetAgentStatus

`func (o *V1AgentFilter) GetAgentStatus() []AgentFilterAgentStatus`

GetAgentStatus returns the AgentStatus field if non-nil, zero value otherwise.

### GetAgentStatusOk

`func (o *V1AgentFilter) GetAgentStatusOk() (*[]AgentFilterAgentStatus, bool)`

GetAgentStatusOk returns a tuple with the AgentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentStatus

`func (o *V1AgentFilter) SetAgentStatus(v []AgentFilterAgentStatus)`

SetAgentStatus sets AgentStatus field to given value.

### HasAgentStatus

`func (o *V1AgentFilter) HasAgentStatus() bool`

HasAgentStatus returns a boolean if a field has been set.

### GetAgentLocationCountry

`func (o *V1AgentFilter) GetAgentLocationCountry() []string`

GetAgentLocationCountry returns the AgentLocationCountry field if non-nil, zero value otherwise.

### GetAgentLocationCountryOk

`func (o *V1AgentFilter) GetAgentLocationCountryOk() (*[]string, bool)`

GetAgentLocationCountryOk returns a tuple with the AgentLocationCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentLocationCountry

`func (o *V1AgentFilter) SetAgentLocationCountry(v []string)`

SetAgentLocationCountry sets AgentLocationCountry field to given value.

### HasAgentLocationCountry

`func (o *V1AgentFilter) HasAgentLocationCountry() bool`

HasAgentLocationCountry returns a boolean if a field has been set.

### GetAgentModifiedAtFrom

`func (o *V1AgentFilter) GetAgentModifiedAtFrom() time.Time`

GetAgentModifiedAtFrom returns the AgentModifiedAtFrom field if non-nil, zero value otherwise.

### GetAgentModifiedAtFromOk

`func (o *V1AgentFilter) GetAgentModifiedAtFromOk() (*time.Time, bool)`

GetAgentModifiedAtFromOk returns a tuple with the AgentModifiedAtFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentModifiedAtFrom

`func (o *V1AgentFilter) SetAgentModifiedAtFrom(v time.Time)`

SetAgentModifiedAtFrom sets AgentModifiedAtFrom field to given value.

### HasAgentModifiedAtFrom

`func (o *V1AgentFilter) HasAgentModifiedAtFrom() bool`

HasAgentModifiedAtFrom returns a boolean if a field has been set.

### GetAgentModifiedAtTo

`func (o *V1AgentFilter) GetAgentModifiedAtTo() time.Time`

GetAgentModifiedAtTo returns the AgentModifiedAtTo field if non-nil, zero value otherwise.

### GetAgentModifiedAtToOk

`func (o *V1AgentFilter) GetAgentModifiedAtToOk() (*time.Time, bool)`

GetAgentModifiedAtToOk returns a tuple with the AgentModifiedAtTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentModifiedAtTo

`func (o *V1AgentFilter) SetAgentModifiedAtTo(v time.Time)`

SetAgentModifiedAtTo sets AgentModifiedAtTo field to given value.

### HasAgentModifiedAtTo

`func (o *V1AgentFilter) HasAgentModifiedAtTo() bool`

HasAgentModifiedAtTo returns a boolean if a field has been set.

### GetAgentName

`func (o *V1AgentFilter) GetAgentName() string`

GetAgentName returns the AgentName field if non-nil, zero value otherwise.

### GetAgentNameOk

`func (o *V1AgentFilter) GetAgentNameOk() (*string, bool)`

GetAgentNameOk returns a tuple with the AgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentName

`func (o *V1AgentFilter) SetAgentName(v string)`

SetAgentName sets AgentName field to given value.

### HasAgentName

`func (o *V1AgentFilter) HasAgentName() bool`

HasAgentName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


