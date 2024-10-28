# NormalizedRuleError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrCode** | Pointer to [**BulkActionsDryRunErrCode**](BulkActionsDryRunErrCode.md) |  | [optional] 
**Message** | **string** |  | 
**Rules** | [**[]RuleDetailsInError**](RuleDetailsInError.md) |  | 
**StatusCode** | **int32** |  | 

## Methods

### NewNormalizedRuleError

`func NewNormalizedRuleError(message string, rules []RuleDetailsInError, statusCode int32, ) *NormalizedRuleError`

NewNormalizedRuleError instantiates a new NormalizedRuleError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNormalizedRuleErrorWithDefaults

`func NewNormalizedRuleErrorWithDefaults() *NormalizedRuleError`

NewNormalizedRuleErrorWithDefaults instantiates a new NormalizedRuleError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrCode

`func (o *NormalizedRuleError) GetErrCode() BulkActionsDryRunErrCode`

GetErrCode returns the ErrCode field if non-nil, zero value otherwise.

### GetErrCodeOk

`func (o *NormalizedRuleError) GetErrCodeOk() (*BulkActionsDryRunErrCode, bool)`

GetErrCodeOk returns a tuple with the ErrCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrCode

`func (o *NormalizedRuleError) SetErrCode(v BulkActionsDryRunErrCode)`

SetErrCode sets ErrCode field to given value.

### HasErrCode

`func (o *NormalizedRuleError) HasErrCode() bool`

HasErrCode returns a boolean if a field has been set.

### GetMessage

`func (o *NormalizedRuleError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NormalizedRuleError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NormalizedRuleError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetRules

`func (o *NormalizedRuleError) GetRules() []RuleDetailsInError`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *NormalizedRuleError) GetRulesOk() (*[]RuleDetailsInError, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *NormalizedRuleError) SetRules(v []RuleDetailsInError)`

SetRules sets Rules field to given value.


### GetStatusCode

`func (o *NormalizedRuleError) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *NormalizedRuleError) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *NormalizedRuleError) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


