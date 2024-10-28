# RulePreviewLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int32** | Execution duration in milliseconds | 
**Errors** | **[]string** |  | 
**Requests** | Pointer to [**[]RulePreviewLoggedRequest**](RulePreviewLoggedRequest.md) |  | [optional] 
**StartedAt** | Pointer to **string** | A string that is not empty and does not contain only whitespace | [optional] 
**Warnings** | **[]string** |  | 

## Methods

### NewRulePreviewLogs

`func NewRulePreviewLogs(duration int32, errors []string, warnings []string, ) *RulePreviewLogs`

NewRulePreviewLogs instantiates a new RulePreviewLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRulePreviewLogsWithDefaults

`func NewRulePreviewLogsWithDefaults() *RulePreviewLogs`

NewRulePreviewLogsWithDefaults instantiates a new RulePreviewLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *RulePreviewLogs) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *RulePreviewLogs) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *RulePreviewLogs) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetErrors

`func (o *RulePreviewLogs) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *RulePreviewLogs) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *RulePreviewLogs) SetErrors(v []string)`

SetErrors sets Errors field to given value.


### GetRequests

`func (o *RulePreviewLogs) GetRequests() []RulePreviewLoggedRequest`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *RulePreviewLogs) GetRequestsOk() (*[]RulePreviewLoggedRequest, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *RulePreviewLogs) SetRequests(v []RulePreviewLoggedRequest)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *RulePreviewLogs) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetStartedAt

`func (o *RulePreviewLogs) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *RulePreviewLogs) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *RulePreviewLogs) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *RulePreviewLogs) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetWarnings

`func (o *RulePreviewLogs) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *RulePreviewLogs) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *RulePreviewLogs) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


