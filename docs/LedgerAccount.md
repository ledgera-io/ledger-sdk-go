# LedgerAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balances** | Pointer to [**[]LedgerBalance**](LedgerBalance.md) |  | [optional] 
**BillingPlanId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**SubType** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewLedgerAccount

`func NewLedgerAccount() *LedgerAccount`

NewLedgerAccount instantiates a new LedgerAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLedgerAccountWithDefaults

`func NewLedgerAccountWithDefaults() *LedgerAccount`

NewLedgerAccountWithDefaults instantiates a new LedgerAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalances

`func (o *LedgerAccount) GetBalances() []LedgerBalance`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *LedgerAccount) GetBalancesOk() (*[]LedgerBalance, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *LedgerAccount) SetBalances(v []LedgerBalance)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *LedgerAccount) HasBalances() bool`

HasBalances returns a boolean if a field has been set.

### GetBillingPlanId

`func (o *LedgerAccount) GetBillingPlanId() string`

GetBillingPlanId returns the BillingPlanId field if non-nil, zero value otherwise.

### GetBillingPlanIdOk

`func (o *LedgerAccount) GetBillingPlanIdOk() (*string, bool)`

GetBillingPlanIdOk returns a tuple with the BillingPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPlanId

`func (o *LedgerAccount) SetBillingPlanId(v string)`

SetBillingPlanId sets BillingPlanId field to given value.

### HasBillingPlanId

`func (o *LedgerAccount) HasBillingPlanId() bool`

HasBillingPlanId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LedgerAccount) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LedgerAccount) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LedgerAccount) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LedgerAccount) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *LedgerAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LedgerAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LedgerAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LedgerAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSubType

`func (o *LedgerAccount) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *LedgerAccount) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *LedgerAccount) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *LedgerAccount) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetType

`func (o *LedgerAccount) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LedgerAccount) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LedgerAccount) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LedgerAccount) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


