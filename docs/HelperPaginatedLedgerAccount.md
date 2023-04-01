# HelperPaginatedLedgerAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**LedgerAccount**](LedgerAccount.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewHelperPaginatedLedgerAccount

`func NewHelperPaginatedLedgerAccount() *HelperPaginatedLedgerAccount`

NewHelperPaginatedLedgerAccount instantiates a new HelperPaginatedLedgerAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelperPaginatedLedgerAccountWithDefaults

`func NewHelperPaginatedLedgerAccountWithDefaults() *HelperPaginatedLedgerAccount`

NewHelperPaginatedLedgerAccountWithDefaults instantiates a new HelperPaginatedLedgerAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *HelperPaginatedLedgerAccount) GetData() LedgerAccount`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HelperPaginatedLedgerAccount) GetDataOk() (*LedgerAccount, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HelperPaginatedLedgerAccount) SetData(v LedgerAccount)`

SetData sets Data field to given value.

### HasData

`func (o *HelperPaginatedLedgerAccount) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMetadata

`func (o *HelperPaginatedLedgerAccount) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HelperPaginatedLedgerAccount) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HelperPaginatedLedgerAccount) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HelperPaginatedLedgerAccount) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


