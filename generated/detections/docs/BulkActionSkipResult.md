# BulkActionSkipResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**SkipReason** | [**BulkEditSkipReason**](BulkEditSkipReason.md) |  | 

## Methods

### NewBulkActionSkipResult

`func NewBulkActionSkipResult(id string, skipReason BulkEditSkipReason, ) *BulkActionSkipResult`

NewBulkActionSkipResult instantiates a new BulkActionSkipResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkActionSkipResultWithDefaults

`func NewBulkActionSkipResultWithDefaults() *BulkActionSkipResult`

NewBulkActionSkipResultWithDefaults instantiates a new BulkActionSkipResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkActionSkipResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkActionSkipResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkActionSkipResult) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BulkActionSkipResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkActionSkipResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkActionSkipResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BulkActionSkipResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSkipReason

`func (o *BulkActionSkipResult) GetSkipReason() BulkEditSkipReason`

GetSkipReason returns the SkipReason field if non-nil, zero value otherwise.

### GetSkipReasonOk

`func (o *BulkActionSkipResult) GetSkipReasonOk() (*BulkEditSkipReason, bool)`

GetSkipReasonOk returns a tuple with the SkipReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipReason

`func (o *BulkActionSkipResult) SetSkipReason(v BulkEditSkipReason)`

SetSkipReason sets SkipReason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


