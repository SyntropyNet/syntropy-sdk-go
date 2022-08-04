# V1NetworkConnectionsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Changes** | [**[]V1ConnectionUpdateChange**](V1ConnectionUpdateChange.md) |  | 

## Methods

### NewV1NetworkConnectionsUpdateRequest

`func NewV1NetworkConnectionsUpdateRequest(changes []V1ConnectionUpdateChange, ) *V1NetworkConnectionsUpdateRequest`

NewV1NetworkConnectionsUpdateRequest instantiates a new V1NetworkConnectionsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NetworkConnectionsUpdateRequestWithDefaults

`func NewV1NetworkConnectionsUpdateRequestWithDefaults() *V1NetworkConnectionsUpdateRequest`

NewV1NetworkConnectionsUpdateRequestWithDefaults instantiates a new V1NetworkConnectionsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChanges

`func (o *V1NetworkConnectionsUpdateRequest) GetChanges() []V1ConnectionUpdateChange`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *V1NetworkConnectionsUpdateRequest) GetChangesOk() (*[]V1ConnectionUpdateChange, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *V1NetworkConnectionsUpdateRequest) SetChanges(v []V1ConnectionUpdateChange)`

SetChanges sets Changes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


