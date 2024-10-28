# BulkActionEditPayloadTags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Value** | **[]string** | String array containing words and phrases to help categorize, filter, and search rules. Defaults to an empty array. | 

## Methods

### NewBulkActionEditPayloadTags

`func NewBulkActionEditPayloadTags(type_ string, value []string, ) *BulkActionEditPayloadTags`

NewBulkActionEditPayloadTags instantiates a new BulkActionEditPayloadTags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkActionEditPayloadTagsWithDefaults

`func NewBulkActionEditPayloadTagsWithDefaults() *BulkActionEditPayloadTags`

NewBulkActionEditPayloadTagsWithDefaults instantiates a new BulkActionEditPayloadTags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BulkActionEditPayloadTags) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BulkActionEditPayloadTags) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BulkActionEditPayloadTags) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *BulkActionEditPayloadTags) GetValue() []string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BulkActionEditPayloadTags) GetValueOk() (*[]string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BulkActionEditPayloadTags) SetValue(v []string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


