# Threshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cardinality** | Pointer to [**[]ThresholdCardinalityInner**](ThresholdCardinalityInner.md) |  | [optional] 
**Field** | [**ThresholdField**](ThresholdField.md) |  | 
**Value** | **int32** | Threshold value | 

## Methods

### NewThreshold

`func NewThreshold(field ThresholdField, value int32, ) *Threshold`

NewThreshold instantiates a new Threshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThresholdWithDefaults

`func NewThresholdWithDefaults() *Threshold`

NewThresholdWithDefaults instantiates a new Threshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardinality

`func (o *Threshold) GetCardinality() []ThresholdCardinalityInner`

GetCardinality returns the Cardinality field if non-nil, zero value otherwise.

### GetCardinalityOk

`func (o *Threshold) GetCardinalityOk() (*[]ThresholdCardinalityInner, bool)`

GetCardinalityOk returns a tuple with the Cardinality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardinality

`func (o *Threshold) SetCardinality(v []ThresholdCardinalityInner)`

SetCardinality sets Cardinality field to given value.

### HasCardinality

`func (o *Threshold) HasCardinality() bool`

HasCardinality returns a boolean if a field has been set.

### GetField

`func (o *Threshold) GetField() ThresholdField`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *Threshold) GetFieldOk() (*ThresholdField, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *Threshold) SetField(v ThresholdField)`

SetField sets Field field to given value.


### GetValue

`func (o *Threshold) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Threshold) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Threshold) SetValue(v int32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


