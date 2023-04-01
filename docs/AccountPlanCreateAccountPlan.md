# AccountPlanCreateAccountPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceType** | [**AccountPlanBalanceType**](AccountPlanBalanceType.md) |  | 
**BalanceValidation** | **string** |  | 
**IsUnique** | Pointer to **bool** |  | [optional] 
**SubType** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewAccountPlanCreateAccountPlan

`func NewAccountPlanCreateAccountPlan(balanceType AccountPlanBalanceType, balanceValidation string, subType string, type_ string, ) *AccountPlanCreateAccountPlan`

NewAccountPlanCreateAccountPlan instantiates a new AccountPlanCreateAccountPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountPlanCreateAccountPlanWithDefaults

`func NewAccountPlanCreateAccountPlanWithDefaults() *AccountPlanCreateAccountPlan`

NewAccountPlanCreateAccountPlanWithDefaults instantiates a new AccountPlanCreateAccountPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceType

`func (o *AccountPlanCreateAccountPlan) GetBalanceType() AccountPlanBalanceType`

GetBalanceType returns the BalanceType field if non-nil, zero value otherwise.

### GetBalanceTypeOk

`func (o *AccountPlanCreateAccountPlan) GetBalanceTypeOk() (*AccountPlanBalanceType, bool)`

GetBalanceTypeOk returns a tuple with the BalanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceType

`func (o *AccountPlanCreateAccountPlan) SetBalanceType(v AccountPlanBalanceType)`

SetBalanceType sets BalanceType field to given value.


### GetBalanceValidation

`func (o *AccountPlanCreateAccountPlan) GetBalanceValidation() string`

GetBalanceValidation returns the BalanceValidation field if non-nil, zero value otherwise.

### GetBalanceValidationOk

`func (o *AccountPlanCreateAccountPlan) GetBalanceValidationOk() (*string, bool)`

GetBalanceValidationOk returns a tuple with the BalanceValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceValidation

`func (o *AccountPlanCreateAccountPlan) SetBalanceValidation(v string)`

SetBalanceValidation sets BalanceValidation field to given value.


### GetIsUnique

`func (o *AccountPlanCreateAccountPlan) GetIsUnique() bool`

GetIsUnique returns the IsUnique field if non-nil, zero value otherwise.

### GetIsUniqueOk

`func (o *AccountPlanCreateAccountPlan) GetIsUniqueOk() (*bool, bool)`

GetIsUniqueOk returns a tuple with the IsUnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnique

`func (o *AccountPlanCreateAccountPlan) SetIsUnique(v bool)`

SetIsUnique sets IsUnique field to given value.

### HasIsUnique

`func (o *AccountPlanCreateAccountPlan) HasIsUnique() bool`

HasIsUnique returns a boolean if a field has been set.

### GetSubType

`func (o *AccountPlanCreateAccountPlan) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *AccountPlanCreateAccountPlan) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *AccountPlanCreateAccountPlan) SetSubType(v string)`

SetSubType sets SubType field to given value.


### GetType

`func (o *AccountPlanCreateAccountPlan) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountPlanCreateAccountPlan) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountPlanCreateAccountPlan) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


