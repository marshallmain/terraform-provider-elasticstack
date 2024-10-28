# RulePreviewParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvocationCount** | **int32** |  | 
**TimeframeEnd** | **time.Time** |  | 

## Methods

### NewRulePreviewParams

`func NewRulePreviewParams(invocationCount int32, timeframeEnd time.Time, ) *RulePreviewParams`

NewRulePreviewParams instantiates a new RulePreviewParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRulePreviewParamsWithDefaults

`func NewRulePreviewParamsWithDefaults() *RulePreviewParams`

NewRulePreviewParamsWithDefaults instantiates a new RulePreviewParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvocationCount

`func (o *RulePreviewParams) GetInvocationCount() int32`

GetInvocationCount returns the InvocationCount field if non-nil, zero value otherwise.

### GetInvocationCountOk

`func (o *RulePreviewParams) GetInvocationCountOk() (*int32, bool)`

GetInvocationCountOk returns a tuple with the InvocationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvocationCount

`func (o *RulePreviewParams) SetInvocationCount(v int32)`

SetInvocationCount sets InvocationCount field to given value.


### GetTimeframeEnd

`func (o *RulePreviewParams) GetTimeframeEnd() time.Time`

GetTimeframeEnd returns the TimeframeEnd field if non-nil, zero value otherwise.

### GetTimeframeEndOk

`func (o *RulePreviewParams) GetTimeframeEndOk() (*time.Time, bool)`

GetTimeframeEndOk returns a tuple with the TimeframeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframeEnd

`func (o *RulePreviewParams) SetTimeframeEnd(v time.Time)`

SetTimeframeEnd sets TimeframeEnd field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


