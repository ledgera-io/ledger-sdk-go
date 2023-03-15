# V1AccountIdJournalGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**Data** | Pointer to [**[]Statement**](Statement.md) |  | [optional] 

## Methods

### NewV1AccountIdJournalGet200Response

`func NewV1AccountIdJournalGet200Response() *V1AccountIdJournalGet200Response`

NewV1AccountIdJournalGet200Response instantiates a new V1AccountIdJournalGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AccountIdJournalGet200ResponseWithDefaults

`func NewV1AccountIdJournalGet200ResponseWithDefaults() *V1AccountIdJournalGet200Response`

NewV1AccountIdJournalGet200ResponseWithDefaults instantiates a new V1AccountIdJournalGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *V1AccountIdJournalGet200Response) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1AccountIdJournalGet200Response) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1AccountIdJournalGet200Response) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1AccountIdJournalGet200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetData

`func (o *V1AccountIdJournalGet200Response) GetData() []Statement`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V1AccountIdJournalGet200Response) GetDataOk() (*[]Statement, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V1AccountIdJournalGet200Response) SetData(v []Statement)`

SetData sets Data field to given value.

### HasData

`func (o *V1AccountIdJournalGet200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


