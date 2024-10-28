# RuleExecutionMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutionGapDurationS** | Pointer to **int32** | Duration in seconds of execution gap | [optional] 
**TotalEnrichmentDurationMs** | Pointer to **int32** | Total time spent enriching documents during current rule execution cycle | [optional] 
**TotalIndexingDurationMs** | Pointer to **int32** | Total time spent indexing documents during current rule execution cycle | [optional] 
**TotalSearchDurationMs** | Pointer to **int32** | Total time spent performing ES searches as measured by Kibana; includes network latency and time spent serializing/deserializing request/response | [optional] 

## Methods

### NewRuleExecutionMetrics

`func NewRuleExecutionMetrics() *RuleExecutionMetrics`

NewRuleExecutionMetrics instantiates a new RuleExecutionMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleExecutionMetricsWithDefaults

`func NewRuleExecutionMetricsWithDefaults() *RuleExecutionMetrics`

NewRuleExecutionMetricsWithDefaults instantiates a new RuleExecutionMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutionGapDurationS

`func (o *RuleExecutionMetrics) GetExecutionGapDurationS() int32`

GetExecutionGapDurationS returns the ExecutionGapDurationS field if non-nil, zero value otherwise.

### GetExecutionGapDurationSOk

`func (o *RuleExecutionMetrics) GetExecutionGapDurationSOk() (*int32, bool)`

GetExecutionGapDurationSOk returns a tuple with the ExecutionGapDurationS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionGapDurationS

`func (o *RuleExecutionMetrics) SetExecutionGapDurationS(v int32)`

SetExecutionGapDurationS sets ExecutionGapDurationS field to given value.

### HasExecutionGapDurationS

`func (o *RuleExecutionMetrics) HasExecutionGapDurationS() bool`

HasExecutionGapDurationS returns a boolean if a field has been set.

### GetTotalEnrichmentDurationMs

`func (o *RuleExecutionMetrics) GetTotalEnrichmentDurationMs() int32`

GetTotalEnrichmentDurationMs returns the TotalEnrichmentDurationMs field if non-nil, zero value otherwise.

### GetTotalEnrichmentDurationMsOk

`func (o *RuleExecutionMetrics) GetTotalEnrichmentDurationMsOk() (*int32, bool)`

GetTotalEnrichmentDurationMsOk returns a tuple with the TotalEnrichmentDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEnrichmentDurationMs

`func (o *RuleExecutionMetrics) SetTotalEnrichmentDurationMs(v int32)`

SetTotalEnrichmentDurationMs sets TotalEnrichmentDurationMs field to given value.

### HasTotalEnrichmentDurationMs

`func (o *RuleExecutionMetrics) HasTotalEnrichmentDurationMs() bool`

HasTotalEnrichmentDurationMs returns a boolean if a field has been set.

### GetTotalIndexingDurationMs

`func (o *RuleExecutionMetrics) GetTotalIndexingDurationMs() int32`

GetTotalIndexingDurationMs returns the TotalIndexingDurationMs field if non-nil, zero value otherwise.

### GetTotalIndexingDurationMsOk

`func (o *RuleExecutionMetrics) GetTotalIndexingDurationMsOk() (*int32, bool)`

GetTotalIndexingDurationMsOk returns a tuple with the TotalIndexingDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIndexingDurationMs

`func (o *RuleExecutionMetrics) SetTotalIndexingDurationMs(v int32)`

SetTotalIndexingDurationMs sets TotalIndexingDurationMs field to given value.

### HasTotalIndexingDurationMs

`func (o *RuleExecutionMetrics) HasTotalIndexingDurationMs() bool`

HasTotalIndexingDurationMs returns a boolean if a field has been set.

### GetTotalSearchDurationMs

`func (o *RuleExecutionMetrics) GetTotalSearchDurationMs() int32`

GetTotalSearchDurationMs returns the TotalSearchDurationMs field if non-nil, zero value otherwise.

### GetTotalSearchDurationMsOk

`func (o *RuleExecutionMetrics) GetTotalSearchDurationMsOk() (*int32, bool)`

GetTotalSearchDurationMsOk returns a tuple with the TotalSearchDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSearchDurationMs

`func (o *RuleExecutionMetrics) SetTotalSearchDurationMs(v int32)`

SetTotalSearchDurationMs sets TotalSearchDurationMs field to given value.

### HasTotalSearchDurationMs

`func (o *RuleExecutionMetrics) HasTotalSearchDurationMs() bool`

HasTotalSearchDurationMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


