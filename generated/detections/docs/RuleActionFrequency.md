# RuleActionFrequency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifyWhen** | [**RuleActionNotifyWhen**](RuleActionNotifyWhen.md) |  | 
**Summary** | **bool** | Action summary indicates whether we will send a summary notification about all the generate alerts or notification per individual alert | 
**Throttle** | [**RuleActionThrottle**](RuleActionThrottle.md) |  | 

## Methods

### NewRuleActionFrequency

`func NewRuleActionFrequency(notifyWhen RuleActionNotifyWhen, summary bool, throttle RuleActionThrottle, ) *RuleActionFrequency`

NewRuleActionFrequency instantiates a new RuleActionFrequency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleActionFrequencyWithDefaults

`func NewRuleActionFrequencyWithDefaults() *RuleActionFrequency`

NewRuleActionFrequencyWithDefaults instantiates a new RuleActionFrequency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifyWhen

`func (o *RuleActionFrequency) GetNotifyWhen() RuleActionNotifyWhen`

GetNotifyWhen returns the NotifyWhen field if non-nil, zero value otherwise.

### GetNotifyWhenOk

`func (o *RuleActionFrequency) GetNotifyWhenOk() (*RuleActionNotifyWhen, bool)`

GetNotifyWhenOk returns a tuple with the NotifyWhen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyWhen

`func (o *RuleActionFrequency) SetNotifyWhen(v RuleActionNotifyWhen)`

SetNotifyWhen sets NotifyWhen field to given value.


### GetSummary

`func (o *RuleActionFrequency) GetSummary() bool`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *RuleActionFrequency) GetSummaryOk() (*bool, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *RuleActionFrequency) SetSummary(v bool)`

SetSummary sets Summary field to given value.


### GetThrottle

`func (o *RuleActionFrequency) GetThrottle() RuleActionThrottle`

GetThrottle returns the Throttle field if non-nil, zero value otherwise.

### GetThrottleOk

`func (o *RuleActionFrequency) GetThrottleOk() (*RuleActionThrottle, bool)`

GetThrottleOk returns a tuple with the Throttle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottle

`func (o *RuleActionFrequency) SetThrottle(v RuleActionThrottle)`

SetThrottle sets Throttle field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


