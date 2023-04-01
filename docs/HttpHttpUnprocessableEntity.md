# HttpHttpUnprocessableEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]ExceptionValidationDetails**](ExceptionValidationDetails.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewHttpHttpUnprocessableEntity

`func NewHttpHttpUnprocessableEntity() *HttpHttpUnprocessableEntity`

NewHttpHttpUnprocessableEntity instantiates a new HttpHttpUnprocessableEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpHttpUnprocessableEntityWithDefaults

`func NewHttpHttpUnprocessableEntityWithDefaults() *HttpHttpUnprocessableEntity`

NewHttpHttpUnprocessableEntityWithDefaults instantiates a new HttpHttpUnprocessableEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *HttpHttpUnprocessableEntity) GetErrors() []ExceptionValidationDetails`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *HttpHttpUnprocessableEntity) GetErrorsOk() (*[]ExceptionValidationDetails, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *HttpHttpUnprocessableEntity) SetErrors(v []ExceptionValidationDetails)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *HttpHttpUnprocessableEntity) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetMessage

`func (o *HttpHttpUnprocessableEntity) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HttpHttpUnprocessableEntity) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HttpHttpUnprocessableEntity) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *HttpHttpUnprocessableEntity) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


