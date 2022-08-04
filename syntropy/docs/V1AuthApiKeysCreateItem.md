# V1AuthApiKeysCreateItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeyName** | **string** |  | 
**ApiKeyDescription** | Pointer to **string** |  | [optional] 
**ApiKeyValidUntil** | Pointer to **time.Time** |  | [optional] 
**ApiKeyAllowedTagNames** | Pointer to **[]string** |  | [optional] 
**ApiKeySecret** | **string** |  | 
**UserId** | **int32** |  | 
**ApiKeyId** | **int32** |  | 
**ApiKeyCreatedAt** | **time.Time** |  | 
**ApiKeyUpdatedAt** | **time.Time** |  | 

## Methods

### NewV1AuthApiKeysCreateItem

`func NewV1AuthApiKeysCreateItem(apiKeyName string, apiKeySecret string, userId int32, apiKeyId int32, apiKeyCreatedAt time.Time, apiKeyUpdatedAt time.Time, ) *V1AuthApiKeysCreateItem`

NewV1AuthApiKeysCreateItem instantiates a new V1AuthApiKeysCreateItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AuthApiKeysCreateItemWithDefaults

`func NewV1AuthApiKeysCreateItemWithDefaults() *V1AuthApiKeysCreateItem`

NewV1AuthApiKeysCreateItemWithDefaults instantiates a new V1AuthApiKeysCreateItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeyName

`func (o *V1AuthApiKeysCreateItem) GetApiKeyName() string`

GetApiKeyName returns the ApiKeyName field if non-nil, zero value otherwise.

### GetApiKeyNameOk

`func (o *V1AuthApiKeysCreateItem) GetApiKeyNameOk() (*string, bool)`

GetApiKeyNameOk returns a tuple with the ApiKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyName

`func (o *V1AuthApiKeysCreateItem) SetApiKeyName(v string)`

SetApiKeyName sets ApiKeyName field to given value.


### GetApiKeyDescription

`func (o *V1AuthApiKeysCreateItem) GetApiKeyDescription() string`

GetApiKeyDescription returns the ApiKeyDescription field if non-nil, zero value otherwise.

### GetApiKeyDescriptionOk

`func (o *V1AuthApiKeysCreateItem) GetApiKeyDescriptionOk() (*string, bool)`

GetApiKeyDescriptionOk returns a tuple with the ApiKeyDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyDescription

`func (o *V1AuthApiKeysCreateItem) SetApiKeyDescription(v string)`

SetApiKeyDescription sets ApiKeyDescription field to given value.

### HasApiKeyDescription

`func (o *V1AuthApiKeysCreateItem) HasApiKeyDescription() bool`

HasApiKeyDescription returns a boolean if a field has been set.

### GetApiKeyValidUntil

`func (o *V1AuthApiKeysCreateItem) GetApiKeyValidUntil() time.Time`

GetApiKeyValidUntil returns the ApiKeyValidUntil field if non-nil, zero value otherwise.

### GetApiKeyValidUntilOk

`func (o *V1AuthApiKeysCreateItem) GetApiKeyValidUntilOk() (*time.Time, bool)`

GetApiKeyValidUntilOk returns a tuple with the ApiKeyValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyValidUntil

`func (o *V1AuthApiKeysCreateItem) SetApiKeyValidUntil(v time.Time)`

SetApiKeyValidUntil sets ApiKeyValidUntil field to given value.

### HasApiKeyValidUntil

`func (o *V1AuthApiKeysCreateItem) HasApiKeyValidUntil() bool`

HasApiKeyValidUntil returns a boolean if a field has been set.

### GetApiKeyAllowedTagNames

`func (o *V1AuthApiKeysCreateItem) GetApiKeyAllowedTagNames() []string`

GetApiKeyAllowedTagNames returns the ApiKeyAllowedTagNames field if non-nil, zero value otherwise.

### GetApiKeyAllowedTagNamesOk

`func (o *V1AuthApiKeysCreateItem) GetApiKeyAllowedTagNamesOk() (*[]string, bool)`

GetApiKeyAllowedTagNamesOk returns a tuple with the ApiKeyAllowedTagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyAllowedTagNames

`func (o *V1AuthApiKeysCreateItem) SetApiKeyAllowedTagNames(v []string)`

SetApiKeyAllowedTagNames sets ApiKeyAllowedTagNames field to given value.

### HasApiKeyAllowedTagNames

`func (o *V1AuthApiKeysCreateItem) HasApiKeyAllowedTagNames() bool`

HasApiKeyAllowedTagNames returns a boolean if a field has been set.

### GetApiKeySecret

`func (o *V1AuthApiKeysCreateItem) GetApiKeySecret() string`

GetApiKeySecret returns the ApiKeySecret field if non-nil, zero value otherwise.

### GetApiKeySecretOk

`func (o *V1AuthApiKeysCreateItem) GetApiKeySecretOk() (*string, bool)`

GetApiKeySecretOk returns a tuple with the ApiKeySecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeySecret

`func (o *V1AuthApiKeysCreateItem) SetApiKeySecret(v string)`

SetApiKeySecret sets ApiKeySecret field to given value.


### GetUserId

`func (o *V1AuthApiKeysCreateItem) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *V1AuthApiKeysCreateItem) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *V1AuthApiKeysCreateItem) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetApiKeyId

`func (o *V1AuthApiKeysCreateItem) GetApiKeyId() int32`

GetApiKeyId returns the ApiKeyId field if non-nil, zero value otherwise.

### GetApiKeyIdOk

`func (o *V1AuthApiKeysCreateItem) GetApiKeyIdOk() (*int32, bool)`

GetApiKeyIdOk returns a tuple with the ApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyId

`func (o *V1AuthApiKeysCreateItem) SetApiKeyId(v int32)`

SetApiKeyId sets ApiKeyId field to given value.


### GetApiKeyCreatedAt

`func (o *V1AuthApiKeysCreateItem) GetApiKeyCreatedAt() time.Time`

GetApiKeyCreatedAt returns the ApiKeyCreatedAt field if non-nil, zero value otherwise.

### GetApiKeyCreatedAtOk

`func (o *V1AuthApiKeysCreateItem) GetApiKeyCreatedAtOk() (*time.Time, bool)`

GetApiKeyCreatedAtOk returns a tuple with the ApiKeyCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyCreatedAt

`func (o *V1AuthApiKeysCreateItem) SetApiKeyCreatedAt(v time.Time)`

SetApiKeyCreatedAt sets ApiKeyCreatedAt field to given value.


### GetApiKeyUpdatedAt

`func (o *V1AuthApiKeysCreateItem) GetApiKeyUpdatedAt() time.Time`

GetApiKeyUpdatedAt returns the ApiKeyUpdatedAt field if non-nil, zero value otherwise.

### GetApiKeyUpdatedAtOk

`func (o *V1AuthApiKeysCreateItem) GetApiKeyUpdatedAtOk() (*time.Time, bool)`

GetApiKeyUpdatedAtOk returns a tuple with the ApiKeyUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyUpdatedAt

`func (o *V1AuthApiKeysCreateItem) SetApiKeyUpdatedAt(v time.Time)`

SetApiKeyUpdatedAt sets ApiKeyUpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


