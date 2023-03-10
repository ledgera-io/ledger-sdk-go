# AccountIdStatementGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**Data** | Pointer to [**[]Statement**](Statement.md) |  | [optional] 

## Methods

### NewAccountIdStatementGet200Response

`func NewAccountIdStatementGet200Response() *AccountIdStatementGet200Response`

NewAccountIdStatementGet200Response instantiates a new AccountIdStatementGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountIdStatementGet200ResponseWithDefaults

`func NewAccountIdStatementGet200ResponseWithDefaults() *AccountIdStatementGet200Response`

NewAccountIdStatementGet200ResponseWithDefaults instantiates a new AccountIdStatementGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *AccountIdStatementGet200Response) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AccountIdStatementGet200Response) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AccountIdStatementGet200Response) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AccountIdStatementGet200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetData

`func (o *AccountIdStatementGet200Response) GetData() []Statement`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AccountIdStatementGet200Response) GetDataOk() (*[]Statement, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AccountIdStatementGet200Response) SetData(v []Statement)`

SetData sets Data field to given value.

### HasData

`func (o *AccountIdStatementGet200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


