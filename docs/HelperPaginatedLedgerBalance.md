# HelperPaginatedLedgerBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**LedgerBalance**](LedgerBalance.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewHelperPaginatedLedgerBalance

`func NewHelperPaginatedLedgerBalance() *HelperPaginatedLedgerBalance`

NewHelperPaginatedLedgerBalance instantiates a new HelperPaginatedLedgerBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelperPaginatedLedgerBalanceWithDefaults

`func NewHelperPaginatedLedgerBalanceWithDefaults() *HelperPaginatedLedgerBalance`

NewHelperPaginatedLedgerBalanceWithDefaults instantiates a new HelperPaginatedLedgerBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *HelperPaginatedLedgerBalance) GetData() LedgerBalance`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HelperPaginatedLedgerBalance) GetDataOk() (*LedgerBalance, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HelperPaginatedLedgerBalance) SetData(v LedgerBalance)`

SetData sets Data field to given value.

### HasData

`func (o *HelperPaginatedLedgerBalance) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMetadata

`func (o *HelperPaginatedLedgerBalance) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HelperPaginatedLedgerBalance) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HelperPaginatedLedgerBalance) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HelperPaginatedLedgerBalance) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


