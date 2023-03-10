# AccountPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**AccountType**](AccountType.md) |  | 
**Label** | **string** |  | 

## Methods

### NewAccountPostRequest

`func NewAccountPostRequest(type_ AccountType, label string, ) *AccountPostRequest`

NewAccountPostRequest instantiates a new AccountPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountPostRequestWithDefaults

`func NewAccountPostRequestWithDefaults() *AccountPostRequest`

NewAccountPostRequestWithDefaults instantiates a new AccountPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AccountPostRequest) GetType() AccountType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountPostRequest) GetTypeOk() (*AccountType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountPostRequest) SetType(v AccountType)`

SetType sets Type field to given value.


### GetLabel

`func (o *AccountPostRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AccountPostRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AccountPostRequest) SetLabel(v string)`

SetLabel sets Label field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


