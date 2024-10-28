# ThresholdRuleResponseFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** |  | 
**Threshold** | [**Threshold**](Threshold.md) |  | 
**Type** | **string** | Rule type | 
**AlertSuppression** | Pointer to [**ThresholdAlertSuppression**](ThresholdAlertSuppression.md) |  | [optional] 
**DataViewId** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to **[]interface{}** |  | [optional] 
**Index** | Pointer to **[]string** |  | [optional] 
**SavedId** | Pointer to **string** |  | [optional] 
**Language** | [**KqlQueryLanguage**](KqlQueryLanguage.md) |  | 

## Methods

### NewThresholdRuleResponseFields

`func NewThresholdRuleResponseFields(query string, threshold Threshold, type_ string, language KqlQueryLanguage, ) *ThresholdRuleResponseFields`

NewThresholdRuleResponseFields instantiates a new ThresholdRuleResponseFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThresholdRuleResponseFieldsWithDefaults

`func NewThresholdRuleResponseFieldsWithDefaults() *ThresholdRuleResponseFields`

NewThresholdRuleResponseFieldsWithDefaults instantiates a new ThresholdRuleResponseFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *ThresholdRuleResponseFields) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ThresholdRuleResponseFields) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ThresholdRuleResponseFields) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetThreshold

`func (o *ThresholdRuleResponseFields) GetThreshold() Threshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *ThresholdRuleResponseFields) GetThresholdOk() (*Threshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *ThresholdRuleResponseFields) SetThreshold(v Threshold)`

SetThreshold sets Threshold field to given value.


### GetType

`func (o *ThresholdRuleResponseFields) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ThresholdRuleResponseFields) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ThresholdRuleResponseFields) SetType(v string)`

SetType sets Type field to given value.


### GetAlertSuppression

`func (o *ThresholdRuleResponseFields) GetAlertSuppression() ThresholdAlertSuppression`

GetAlertSuppression returns the AlertSuppression field if non-nil, zero value otherwise.

### GetAlertSuppressionOk

`func (o *ThresholdRuleResponseFields) GetAlertSuppressionOk() (*ThresholdAlertSuppression, bool)`

GetAlertSuppressionOk returns a tuple with the AlertSuppression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertSuppression

`func (o *ThresholdRuleResponseFields) SetAlertSuppression(v ThresholdAlertSuppression)`

SetAlertSuppression sets AlertSuppression field to given value.

### HasAlertSuppression

`func (o *ThresholdRuleResponseFields) HasAlertSuppression() bool`

HasAlertSuppression returns a boolean if a field has been set.

### GetDataViewId

`func (o *ThresholdRuleResponseFields) GetDataViewId() string`

GetDataViewId returns the DataViewId field if non-nil, zero value otherwise.

### GetDataViewIdOk

`func (o *ThresholdRuleResponseFields) GetDataViewIdOk() (*string, bool)`

GetDataViewIdOk returns a tuple with the DataViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataViewId

`func (o *ThresholdRuleResponseFields) SetDataViewId(v string)`

SetDataViewId sets DataViewId field to given value.

### HasDataViewId

`func (o *ThresholdRuleResponseFields) HasDataViewId() bool`

HasDataViewId returns a boolean if a field has been set.

### GetFilters

`func (o *ThresholdRuleResponseFields) GetFilters() []interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ThresholdRuleResponseFields) GetFiltersOk() (*[]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ThresholdRuleResponseFields) SetFilters(v []interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ThresholdRuleResponseFields) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetIndex

`func (o *ThresholdRuleResponseFields) GetIndex() []string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ThresholdRuleResponseFields) GetIndexOk() (*[]string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ThresholdRuleResponseFields) SetIndex(v []string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *ThresholdRuleResponseFields) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetSavedId

`func (o *ThresholdRuleResponseFields) GetSavedId() string`

GetSavedId returns the SavedId field if non-nil, zero value otherwise.

### GetSavedIdOk

`func (o *ThresholdRuleResponseFields) GetSavedIdOk() (*string, bool)`

GetSavedIdOk returns a tuple with the SavedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedId

`func (o *ThresholdRuleResponseFields) SetSavedId(v string)`

SetSavedId sets SavedId field to given value.

### HasSavedId

`func (o *ThresholdRuleResponseFields) HasSavedId() bool`

HasSavedId returns a boolean if a field has been set.

### GetLanguage

`func (o *ThresholdRuleResponseFields) GetLanguage() KqlQueryLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ThresholdRuleResponseFields) GetLanguageOk() (*KqlQueryLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ThresholdRuleResponseFields) SetLanguage(v KqlQueryLanguage)`

SetLanguage sets Language field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


