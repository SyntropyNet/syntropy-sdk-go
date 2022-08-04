# V1NetworkAuthAccessTokensCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | **[]string** |  | 
**AccessTokenExpiration** | **time.Time** |  | 
**AccessTokenName** | **string** |  | 
**AccessTokenDescription** | Pointer to **string** |  | [optional] 

## Methods

### NewV1NetworkAuthAccessTokensCreateRequest

`func NewV1NetworkAuthAccessTokensCreateRequest(permissions []string, accessTokenExpiration time.Time, accessTokenName string, ) *V1NetworkAuthAccessTokensCreateRequest`

NewV1NetworkAuthAccessTokensCreateRequest instantiates a new V1NetworkAuthAccessTokensCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NetworkAuthAccessTokensCreateRequestWithDefaults

`func NewV1NetworkAuthAccessTokensCreateRequestWithDefaults() *V1NetworkAuthAccessTokensCreateRequest`

NewV1NetworkAuthAccessTokensCreateRequestWithDefaults instantiates a new V1NetworkAuthAccessTokensCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *V1NetworkAuthAccessTokensCreateRequest) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *V1NetworkAuthAccessTokensCreateRequest) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *V1NetworkAuthAccessTokensCreateRequest) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.


### GetAccessTokenExpiration

`func (o *V1NetworkAuthAccessTokensCreateRequest) GetAccessTokenExpiration() time.Time`

GetAccessTokenExpiration returns the AccessTokenExpiration field if non-nil, zero value otherwise.

### GetAccessTokenExpirationOk

`func (o *V1NetworkAuthAccessTokensCreateRequest) GetAccessTokenExpirationOk() (*time.Time, bool)`

GetAccessTokenExpirationOk returns a tuple with the AccessTokenExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenExpiration

`func (o *V1NetworkAuthAccessTokensCreateRequest) SetAccessTokenExpiration(v time.Time)`

SetAccessTokenExpiration sets AccessTokenExpiration field to given value.


### GetAccessTokenName

`func (o *V1NetworkAuthAccessTokensCreateRequest) GetAccessTokenName() string`

GetAccessTokenName returns the AccessTokenName field if non-nil, zero value otherwise.

### GetAccessTokenNameOk

`func (o *V1NetworkAuthAccessTokensCreateRequest) GetAccessTokenNameOk() (*string, bool)`

GetAccessTokenNameOk returns a tuple with the AccessTokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenName

`func (o *V1NetworkAuthAccessTokensCreateRequest) SetAccessTokenName(v string)`

SetAccessTokenName sets AccessTokenName field to given value.


### GetAccessTokenDescription

`func (o *V1NetworkAuthAccessTokensCreateRequest) GetAccessTokenDescription() string`

GetAccessTokenDescription returns the AccessTokenDescription field if non-nil, zero value otherwise.

### GetAccessTokenDescriptionOk

`func (o *V1NetworkAuthAccessTokensCreateRequest) GetAccessTokenDescriptionOk() (*string, bool)`

GetAccessTokenDescriptionOk returns a tuple with the AccessTokenDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenDescription

`func (o *V1NetworkAuthAccessTokensCreateRequest) SetAccessTokenDescription(v string)`

SetAccessTokenDescription sets AccessTokenDescription field to given value.

### HasAccessTokenDescription

`func (o *V1NetworkAuthAccessTokensCreateRequest) HasAccessTokenDescription() bool`

HasAccessTokenDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


