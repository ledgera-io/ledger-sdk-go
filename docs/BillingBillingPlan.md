# BillingBillingPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Expiration** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to [**BillingBillingPlanMetadata**](BillingBillingPlanMetadata.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**BillingBillingPlanStatus**](BillingBillingPlanStatus.md) |  | [optional] 

## Methods

### NewBillingBillingPlan

`func NewBillingBillingPlan() *BillingBillingPlan`

NewBillingBillingPlan instantiates a new BillingBillingPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingBillingPlanWithDefaults

`func NewBillingBillingPlanWithDefaults() *BillingBillingPlan`

NewBillingBillingPlanWithDefaults instantiates a new BillingBillingPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *BillingBillingPlan) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BillingBillingPlan) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BillingBillingPlan) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BillingBillingPlan) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *BillingBillingPlan) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *BillingBillingPlan) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *BillingBillingPlan) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *BillingBillingPlan) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDescription

`func (o *BillingBillingPlan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillingBillingPlan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillingBillingPlan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillingBillingPlan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpiration

`func (o *BillingBillingPlan) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *BillingBillingPlan) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *BillingBillingPlan) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *BillingBillingPlan) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetId

`func (o *BillingBillingPlan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingBillingPlan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingBillingPlan) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingBillingPlan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *BillingBillingPlan) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *BillingBillingPlan) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *BillingBillingPlan) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *BillingBillingPlan) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetMetadata

`func (o *BillingBillingPlan) GetMetadata() BillingBillingPlanMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BillingBillingPlan) GetMetadataOk() (*BillingBillingPlanMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BillingBillingPlan) SetMetadata(v BillingBillingPlanMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BillingBillingPlan) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *BillingBillingPlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingBillingPlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingBillingPlan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillingBillingPlan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *BillingBillingPlan) GetStatus() BillingBillingPlanStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BillingBillingPlan) GetStatusOk() (*BillingBillingPlanStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BillingBillingPlan) SetStatus(v BillingBillingPlanStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BillingBillingPlan) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


