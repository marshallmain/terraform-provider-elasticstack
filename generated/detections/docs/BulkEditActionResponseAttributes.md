# BulkEditActionResponseAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]NormalizedRuleError**](NormalizedRuleError.md) |  | [optional] 
**Results** | [**BulkEditActionResults**](BulkEditActionResults.md) |  | 
**Summary** | [**BulkEditActionSummary**](BulkEditActionSummary.md) |  | 

## Methods

### NewBulkEditActionResponseAttributes

`func NewBulkEditActionResponseAttributes(results BulkEditActionResults, summary BulkEditActionSummary, ) *BulkEditActionResponseAttributes`

NewBulkEditActionResponseAttributes instantiates a new BulkEditActionResponseAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkEditActionResponseAttributesWithDefaults

`func NewBulkEditActionResponseAttributesWithDefaults() *BulkEditActionResponseAttributes`

NewBulkEditActionResponseAttributesWithDefaults instantiates a new BulkEditActionResponseAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *BulkEditActionResponseAttributes) GetErrors() []NormalizedRuleError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BulkEditActionResponseAttributes) GetErrorsOk() (*[]NormalizedRuleError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BulkEditActionResponseAttributes) SetErrors(v []NormalizedRuleError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BulkEditActionResponseAttributes) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetResults

`func (o *BulkEditActionResponseAttributes) GetResults() BulkEditActionResults`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BulkEditActionResponseAttributes) GetResultsOk() (*BulkEditActionResults, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BulkEditActionResponseAttributes) SetResults(v BulkEditActionResults)`

SetResults sets Results field to given value.


### GetSummary

`func (o *BulkEditActionResponseAttributes) GetSummary() BulkEditActionSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *BulkEditActionResponseAttributes) GetSummaryOk() (*BulkEditActionSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *BulkEditActionResponseAttributes) SetSummary(v BulkEditActionSummary)`

SetSummary sets Summary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


