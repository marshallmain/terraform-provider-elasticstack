# BulkEditActionResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | [**[]RuleResponse**](RuleResponse.md) |  | 
**Deleted** | [**[]RuleResponse**](RuleResponse.md) |  | 
**Skipped** | [**[]BulkActionSkipResult**](BulkActionSkipResult.md) |  | 
**Updated** | [**[]RuleResponse**](RuleResponse.md) |  | 

## Methods

### NewBulkEditActionResults

`func NewBulkEditActionResults(created []RuleResponse, deleted []RuleResponse, skipped []BulkActionSkipResult, updated []RuleResponse, ) *BulkEditActionResults`

NewBulkEditActionResults instantiates a new BulkEditActionResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkEditActionResultsWithDefaults

`func NewBulkEditActionResultsWithDefaults() *BulkEditActionResults`

NewBulkEditActionResultsWithDefaults instantiates a new BulkEditActionResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *BulkEditActionResults) GetCreated() []RuleResponse`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BulkEditActionResults) GetCreatedOk() (*[]RuleResponse, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BulkEditActionResults) SetCreated(v []RuleResponse)`

SetCreated sets Created field to given value.


### GetDeleted

`func (o *BulkEditActionResults) GetDeleted() []RuleResponse`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *BulkEditActionResults) GetDeletedOk() (*[]RuleResponse, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *BulkEditActionResults) SetDeleted(v []RuleResponse)`

SetDeleted sets Deleted field to given value.


### GetSkipped

`func (o *BulkEditActionResults) GetSkipped() []BulkActionSkipResult`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *BulkEditActionResults) GetSkippedOk() (*[]BulkActionSkipResult, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *BulkEditActionResults) SetSkipped(v []BulkActionSkipResult)`

SetSkipped sets Skipped field to given value.


### GetUpdated

`func (o *BulkEditActionResults) GetUpdated() []RuleResponse`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *BulkEditActionResults) GetUpdatedOk() (*[]RuleResponse, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *BulkEditActionResults) SetUpdated(v []RuleResponse)`

SetUpdated sets Updated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


