# BulkActionEditPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Value** | [**BulkActionEditPayloadScheduleValue**](BulkActionEditPayloadScheduleValue.md) |  | 
**OverwriteDataViews** | Pointer to **bool** |  | [optional] 

## Methods

### NewBulkActionEditPayload

`func NewBulkActionEditPayload(type_ string, value BulkActionEditPayloadScheduleValue, ) *BulkActionEditPayload`

NewBulkActionEditPayload instantiates a new BulkActionEditPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkActionEditPayloadWithDefaults

`func NewBulkActionEditPayloadWithDefaults() *BulkActionEditPayload`

NewBulkActionEditPayloadWithDefaults instantiates a new BulkActionEditPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BulkActionEditPayload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BulkActionEditPayload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BulkActionEditPayload) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *BulkActionEditPayload) GetValue() BulkActionEditPayloadScheduleValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BulkActionEditPayload) GetValueOk() (*BulkActionEditPayloadScheduleValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BulkActionEditPayload) SetValue(v BulkActionEditPayloadScheduleValue)`

SetValue sets Value field to given value.


### GetOverwriteDataViews

`func (o *BulkActionEditPayload) GetOverwriteDataViews() bool`

GetOverwriteDataViews returns the OverwriteDataViews field if non-nil, zero value otherwise.

### GetOverwriteDataViewsOk

`func (o *BulkActionEditPayload) GetOverwriteDataViewsOk() (*bool, bool)`

GetOverwriteDataViewsOk returns a tuple with the OverwriteDataViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteDataViews

`func (o *BulkActionEditPayload) SetOverwriteDataViews(v bool)`

SetOverwriteDataViews sets OverwriteDataViews field to given value.

### HasOverwriteDataViews

`func (o *BulkActionEditPayload) HasOverwriteDataViews() bool`

HasOverwriteDataViews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


