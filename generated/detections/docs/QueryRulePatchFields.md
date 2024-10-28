# QueryRulePatchFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertSuppression** | Pointer to [**AlertSuppression**](AlertSuppression.md) |  | [optional] 
**DataViewId** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to **[]interface{}** |  | [optional] 
**Index** | Pointer to **[]string** |  | [optional] 
**ResponseActions** | Pointer to [**[]ResponseAction**](ResponseAction.md) |  | [optional] 
**SavedId** | Pointer to **string** |  | [optional] 
**Language** | Pointer to [**KqlQueryLanguage**](KqlQueryLanguage.md) |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | Rule type | [optional] 

## Methods

### NewQueryRulePatchFields

`func NewQueryRulePatchFields() *QueryRulePatchFields`

NewQueryRulePatchFields instantiates a new QueryRulePatchFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryRulePatchFieldsWithDefaults

`func NewQueryRulePatchFieldsWithDefaults() *QueryRulePatchFields`

NewQueryRulePatchFieldsWithDefaults instantiates a new QueryRulePatchFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertSuppression

`func (o *QueryRulePatchFields) GetAlertSuppression() AlertSuppression`

GetAlertSuppression returns the AlertSuppression field if non-nil, zero value otherwise.

### GetAlertSuppressionOk

`func (o *QueryRulePatchFields) GetAlertSuppressionOk() (*AlertSuppression, bool)`

GetAlertSuppressionOk returns a tuple with the AlertSuppression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertSuppression

`func (o *QueryRulePatchFields) SetAlertSuppression(v AlertSuppression)`

SetAlertSuppression sets AlertSuppression field to given value.

### HasAlertSuppression

`func (o *QueryRulePatchFields) HasAlertSuppression() bool`

HasAlertSuppression returns a boolean if a field has been set.

### GetDataViewId

`func (o *QueryRulePatchFields) GetDataViewId() string`

GetDataViewId returns the DataViewId field if non-nil, zero value otherwise.

### GetDataViewIdOk

`func (o *QueryRulePatchFields) GetDataViewIdOk() (*string, bool)`

GetDataViewIdOk returns a tuple with the DataViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataViewId

`func (o *QueryRulePatchFields) SetDataViewId(v string)`

SetDataViewId sets DataViewId field to given value.

### HasDataViewId

`func (o *QueryRulePatchFields) HasDataViewId() bool`

HasDataViewId returns a boolean if a field has been set.

### GetFilters

`func (o *QueryRulePatchFields) GetFilters() []interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *QueryRulePatchFields) GetFiltersOk() (*[]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *QueryRulePatchFields) SetFilters(v []interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *QueryRulePatchFields) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetIndex

`func (o *QueryRulePatchFields) GetIndex() []string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *QueryRulePatchFields) GetIndexOk() (*[]string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *QueryRulePatchFields) SetIndex(v []string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *QueryRulePatchFields) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetResponseActions

`func (o *QueryRulePatchFields) GetResponseActions() []ResponseAction`

GetResponseActions returns the ResponseActions field if non-nil, zero value otherwise.

### GetResponseActionsOk

`func (o *QueryRulePatchFields) GetResponseActionsOk() (*[]ResponseAction, bool)`

GetResponseActionsOk returns a tuple with the ResponseActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseActions

`func (o *QueryRulePatchFields) SetResponseActions(v []ResponseAction)`

SetResponseActions sets ResponseActions field to given value.

### HasResponseActions

`func (o *QueryRulePatchFields) HasResponseActions() bool`

HasResponseActions returns a boolean if a field has been set.

### GetSavedId

`func (o *QueryRulePatchFields) GetSavedId() string`

GetSavedId returns the SavedId field if non-nil, zero value otherwise.

### GetSavedIdOk

`func (o *QueryRulePatchFields) GetSavedIdOk() (*string, bool)`

GetSavedIdOk returns a tuple with the SavedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedId

`func (o *QueryRulePatchFields) SetSavedId(v string)`

SetSavedId sets SavedId field to given value.

### HasSavedId

`func (o *QueryRulePatchFields) HasSavedId() bool`

HasSavedId returns a boolean if a field has been set.

### GetLanguage

`func (o *QueryRulePatchFields) GetLanguage() KqlQueryLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *QueryRulePatchFields) GetLanguageOk() (*KqlQueryLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *QueryRulePatchFields) SetLanguage(v KqlQueryLanguage)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *QueryRulePatchFields) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetQuery

`func (o *QueryRulePatchFields) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *QueryRulePatchFields) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *QueryRulePatchFields) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *QueryRulePatchFields) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetType

`func (o *QueryRulePatchFields) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QueryRulePatchFields) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QueryRulePatchFields) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *QueryRulePatchFields) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


