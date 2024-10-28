# RiskScoreMappingInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** |  | 
**Operator** | **string** |  | 
**RiskScore** | Pointer to **int32** | Risk score (0 to 100) | [optional] 
**Value** | **string** |  | 

## Methods

### NewRiskScoreMappingInner

`func NewRiskScoreMappingInner(field string, operator string, value string, ) *RiskScoreMappingInner`

NewRiskScoreMappingInner instantiates a new RiskScoreMappingInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskScoreMappingInnerWithDefaults

`func NewRiskScoreMappingInnerWithDefaults() *RiskScoreMappingInner`

NewRiskScoreMappingInnerWithDefaults instantiates a new RiskScoreMappingInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *RiskScoreMappingInner) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *RiskScoreMappingInner) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *RiskScoreMappingInner) SetField(v string)`

SetField sets Field field to given value.


### GetOperator

`func (o *RiskScoreMappingInner) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *RiskScoreMappingInner) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *RiskScoreMappingInner) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetRiskScore

`func (o *RiskScoreMappingInner) GetRiskScore() int32`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *RiskScoreMappingInner) GetRiskScoreOk() (*int32, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *RiskScoreMappingInner) SetRiskScore(v int32)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *RiskScoreMappingInner) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### GetValue

`func (o *RiskScoreMappingInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RiskScoreMappingInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RiskScoreMappingInner) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


