# ThresholdRulePatchFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertSuppression** | Pointer to [**ThresholdAlertSuppression**](ThresholdAlertSuppression.md) |  | [optional] 
**DataViewId** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to **[]interface{}** |  | [optional] 
**Index** | Pointer to **[]string** |  | [optional] 
**SavedId** | Pointer to **string** |  | [optional] 
**Language** | Pointer to [**KqlQueryLanguage**](KqlQueryLanguage.md) |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**Threshold** | Pointer to [**Threshold**](Threshold.md) |  | [optional] 
**Type** | Pointer to **string** | Rule type | [optional] 

## Methods

### NewThresholdRulePatchFields

`func NewThresholdRulePatchFields() *ThresholdRulePatchFields`

NewThresholdRulePatchFields instantiates a new ThresholdRulePatchFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThresholdRulePatchFieldsWithDefaults

`func NewThresholdRulePatchFieldsWithDefaults() *ThresholdRulePatchFields`

NewThresholdRulePatchFieldsWithDefaults instantiates a new ThresholdRulePatchFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertSuppression

`func (o *ThresholdRulePatchFields) GetAlertSuppression() ThresholdAlertSuppression`

GetAlertSuppression returns the AlertSuppression field if non-nil, zero value otherwise.

### GetAlertSuppressionOk

`func (o *ThresholdRulePatchFields) GetAlertSuppressionOk() (*ThresholdAlertSuppression, bool)`

GetAlertSuppressionOk returns a tuple with the AlertSuppression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertSuppression

`func (o *ThresholdRulePatchFields) SetAlertSuppression(v ThresholdAlertSuppression)`

SetAlertSuppression sets AlertSuppression field to given value.

### HasAlertSuppression

`func (o *ThresholdRulePatchFields) HasAlertSuppression() bool`

HasAlertSuppression returns a boolean if a field has been set.

### GetDataViewId

`func (o *ThresholdRulePatchFields) GetDataViewId() string`

GetDataViewId returns the DataViewId field if non-nil, zero value otherwise.

### GetDataViewIdOk

`func (o *ThresholdRulePatchFields) GetDataViewIdOk() (*string, bool)`

GetDataViewIdOk returns a tuple with the DataViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataViewId

`func (o *ThresholdRulePatchFields) SetDataViewId(v string)`

SetDataViewId sets DataViewId field to given value.

### HasDataViewId

`func (o *ThresholdRulePatchFields) HasDataViewId() bool`

HasDataViewId returns a boolean if a field has been set.

### GetFilters

`func (o *ThresholdRulePatchFields) GetFilters() []interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ThresholdRulePatchFields) GetFiltersOk() (*[]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ThresholdRulePatchFields) SetFilters(v []interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ThresholdRulePatchFields) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetIndex

`func (o *ThresholdRulePatchFields) GetIndex() []string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ThresholdRulePatchFields) GetIndexOk() (*[]string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ThresholdRulePatchFields) SetIndex(v []string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *ThresholdRulePatchFields) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetSavedId

`func (o *ThresholdRulePatchFields) GetSavedId() string`

GetSavedId returns the SavedId field if non-nil, zero value otherwise.

### GetSavedIdOk

`func (o *ThresholdRulePatchFields) GetSavedIdOk() (*string, bool)`

GetSavedIdOk returns a tuple with the SavedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedId

`func (o *ThresholdRulePatchFields) SetSavedId(v string)`

SetSavedId sets SavedId field to given value.

### HasSavedId

`func (o *ThresholdRulePatchFields) HasSavedId() bool`

HasSavedId returns a boolean if a field has been set.

### GetLanguage

`func (o *ThresholdRulePatchFields) GetLanguage() KqlQueryLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ThresholdRulePatchFields) GetLanguageOk() (*KqlQueryLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ThresholdRulePatchFields) SetLanguage(v KqlQueryLanguage)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ThresholdRulePatchFields) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetQuery

`func (o *ThresholdRulePatchFields) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ThresholdRulePatchFields) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ThresholdRulePatchFields) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *ThresholdRulePatchFields) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetThreshold

`func (o *ThresholdRulePatchFields) GetThreshold() Threshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *ThresholdRulePatchFields) GetThresholdOk() (*Threshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *ThresholdRulePatchFields) SetThreshold(v Threshold)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *ThresholdRulePatchFields) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetType

`func (o *ThresholdRulePatchFields) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ThresholdRulePatchFields) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ThresholdRulePatchFields) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ThresholdRulePatchFields) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


