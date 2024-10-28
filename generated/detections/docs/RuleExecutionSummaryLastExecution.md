# RuleExecutionSummaryLastExecution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **time.Time** | Date of the last execution | 
**Message** | **string** |  | 
**Metrics** | [**RuleExecutionMetrics**](RuleExecutionMetrics.md) |  | 
**Status** | [**RuleExecutionStatus**](RuleExecutionStatus.md) |  | 
**StatusOrder** | **int32** |  | 

## Methods

### NewRuleExecutionSummaryLastExecution

`func NewRuleExecutionSummaryLastExecution(date time.Time, message string, metrics RuleExecutionMetrics, status RuleExecutionStatus, statusOrder int32, ) *RuleExecutionSummaryLastExecution`

NewRuleExecutionSummaryLastExecution instantiates a new RuleExecutionSummaryLastExecution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleExecutionSummaryLastExecutionWithDefaults

`func NewRuleExecutionSummaryLastExecutionWithDefaults() *RuleExecutionSummaryLastExecution`

NewRuleExecutionSummaryLastExecutionWithDefaults instantiates a new RuleExecutionSummaryLastExecution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *RuleExecutionSummaryLastExecution) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *RuleExecutionSummaryLastExecution) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *RuleExecutionSummaryLastExecution) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetMessage

`func (o *RuleExecutionSummaryLastExecution) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RuleExecutionSummaryLastExecution) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RuleExecutionSummaryLastExecution) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetMetrics

`func (o *RuleExecutionSummaryLastExecution) GetMetrics() RuleExecutionMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *RuleExecutionSummaryLastExecution) GetMetricsOk() (*RuleExecutionMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *RuleExecutionSummaryLastExecution) SetMetrics(v RuleExecutionMetrics)`

SetMetrics sets Metrics field to given value.


### GetStatus

`func (o *RuleExecutionSummaryLastExecution) GetStatus() RuleExecutionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RuleExecutionSummaryLastExecution) GetStatusOk() (*RuleExecutionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RuleExecutionSummaryLastExecution) SetStatus(v RuleExecutionStatus)`

SetStatus sets Status field to given value.


### GetStatusOrder

`func (o *RuleExecutionSummaryLastExecution) GetStatusOrder() int32`

GetStatusOrder returns the StatusOrder field if non-nil, zero value otherwise.

### GetStatusOrderOk

`func (o *RuleExecutionSummaryLastExecution) GetStatusOrderOk() (*int32, bool)`

GetStatusOrderOk returns a tuple with the StatusOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusOrder

`func (o *RuleExecutionSummaryLastExecution) SetStatusOrder(v int32)`

SetStatusOrder sets StatusOrder field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


