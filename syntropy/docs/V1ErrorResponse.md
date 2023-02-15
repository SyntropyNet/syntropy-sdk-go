# V1ErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]V1ErrorResponseErrorsInner**](V1ErrorResponseErrorsInner.md) |  | [optional] 

## Methods

### NewV1ErrorResponse

`func NewV1ErrorResponse() *V1ErrorResponse`

NewV1ErrorResponse instantiates a new V1ErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ErrorResponseWithDefaults

`func NewV1ErrorResponseWithDefaults() *V1ErrorResponse`

NewV1ErrorResponseWithDefaults instantiates a new V1ErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *V1ErrorResponse) GetErrors() []V1ErrorResponseErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V1ErrorResponse) GetErrorsOk() (*[]V1ErrorResponseErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V1ErrorResponse) SetErrors(v []V1ErrorResponseErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V1ErrorResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


