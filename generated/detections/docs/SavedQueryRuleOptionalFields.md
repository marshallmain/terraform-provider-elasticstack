# SavedQueryRuleOptionalFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertSuppression** | Pointer to [**AlertSuppression**](AlertSuppression.md) |  | [optional] 
**DataViewId** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to **[]interface{}** |  | [optional] 
**Index** | Pointer to **[]string** |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**ResponseActions** | Pointer to [**[]ResponseAction**](ResponseAction.md) |  | [optional] 

## Methods

### NewSavedQueryRuleOptionalFields

`func NewSavedQueryRuleOptionalFields() *SavedQueryRuleOptionalFields`

NewSavedQueryRuleOptionalFields instantiates a new SavedQueryRuleOptionalFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavedQueryRuleOptionalFieldsWithDefaults

`func NewSavedQueryRuleOptionalFieldsWithDefaults() *SavedQueryRuleOptionalFields`

NewSavedQueryRuleOptionalFieldsWithDefaults instantiates a new SavedQueryRuleOptionalFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertSuppression

`func (o *SavedQueryRuleOptionalFields) GetAlertSuppression() AlertSuppression`

GetAlertSuppression returns the AlertSuppression field if non-nil, zero value otherwise.

### GetAlertSuppressionOk

`func (o *SavedQueryRuleOptionalFields) GetAlertSuppressionOk() (*AlertSuppression, bool)`

GetAlertSuppressionOk returns a tuple with the AlertSuppression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertSuppression

`func (o *SavedQueryRuleOptionalFields) SetAlertSuppression(v AlertSuppression)`

SetAlertSuppression sets AlertSuppression field to given value.

### HasAlertSuppression

`func (o *SavedQueryRuleOptionalFields) HasAlertSuppression() bool`

HasAlertSuppression returns a boolean if a field has been set.

### GetDataViewId

`func (o *SavedQueryRuleOptionalFields) GetDataViewId() string`

GetDataViewId returns the DataViewId field if non-nil, zero value otherwise.

### GetDataViewIdOk

`func (o *SavedQueryRuleOptionalFields) GetDataViewIdOk() (*string, bool)`

GetDataViewIdOk returns a tuple with the DataViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataViewId

`func (o *SavedQueryRuleOptionalFields) SetDataViewId(v string)`

SetDataViewId sets DataViewId field to given value.

### HasDataViewId

`func (o *SavedQueryRuleOptionalFields) HasDataViewId() bool`

HasDataViewId returns a boolean if a field has been set.

### GetFilters

`func (o *SavedQueryRuleOptionalFields) GetFilters() []interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *SavedQueryRuleOptionalFields) GetFiltersOk() (*[]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *SavedQueryRuleOptionalFields) SetFilters(v []interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *SavedQueryRuleOptionalFields) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetIndex

`func (o *SavedQueryRuleOptionalFields) GetIndex() []string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *SavedQueryRuleOptionalFields) GetIndexOk() (*[]string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *SavedQueryRuleOptionalFields) SetIndex(v []string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *SavedQueryRuleOptionalFields) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetQuery

`func (o *SavedQueryRuleOptionalFields) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SavedQueryRuleOptionalFields) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SavedQueryRuleOptionalFields) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *SavedQueryRuleOptionalFields) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetResponseActions

`func (o *SavedQueryRuleOptionalFields) GetResponseActions() []ResponseAction`

GetResponseActions returns the ResponseActions field if non-nil, zero value otherwise.

### GetResponseActionsOk

`func (o *SavedQueryRuleOptionalFields) GetResponseActionsOk() (*[]ResponseAction, bool)`

GetResponseActionsOk returns a tuple with the ResponseActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseActions

`func (o *SavedQueryRuleOptionalFields) SetResponseActions(v []ResponseAction)`

SetResponseActions sets ResponseActions field to given value.

### HasResponseActions

`func (o *SavedQueryRuleOptionalFields) HasResponseActions() bool`

HasResponseActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


