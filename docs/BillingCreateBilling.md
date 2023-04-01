# BillingCreateBilling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Entries** | [**[]BillingTransactionEntry**](BillingTransactionEntry.md) |  | 
**Expiration** | **string** |  | 
**IsDefault** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewBillingCreateBilling

`func NewBillingCreateBilling(entries []BillingTransactionEntry, expiration string, name string, ) *BillingCreateBilling`

NewBillingCreateBilling instantiates a new BillingCreateBilling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingCreateBillingWithDefaults

`func NewBillingCreateBillingWithDefaults() *BillingCreateBilling`

NewBillingCreateBillingWithDefaults instantiates a new BillingCreateBilling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *BillingCreateBilling) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillingCreateBilling) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillingCreateBilling) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillingCreateBilling) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntries

`func (o *BillingCreateBilling) GetEntries() []BillingTransactionEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *BillingCreateBilling) GetEntriesOk() (*[]BillingTransactionEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *BillingCreateBilling) SetEntries(v []BillingTransactionEntry)`

SetEntries sets Entries field to given value.


### GetExpiration

`func (o *BillingCreateBilling) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *BillingCreateBilling) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *BillingCreateBilling) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.


### GetIsDefault

`func (o *BillingCreateBilling) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *BillingCreateBilling) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *BillingCreateBilling) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *BillingCreateBilling) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *BillingCreateBilling) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingCreateBilling) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingCreateBilling) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


