# NewTermsRuleResponseFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HistoryWindowStart** | **string** | A string that is not empty and does not contain only whitespace | 
**NewTermsFields** | **[]string** |  | 
**Query** | **string** |  | 
**Type** | **string** | Rule type | 
**AlertSuppression** | Pointer to [**AlertSuppression**](AlertSuppression.md) |  | [optional] 
**DataViewId** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to **[]interface{}** |  | [optional] 
**Index** | Pointer to **[]string** |  | [optional] 
**ResponseActions** | Pointer to [**[]ResponseAction**](ResponseAction.md) |  | [optional] 
**Language** | [**KqlQueryLanguage**](KqlQueryLanguage.md) |  | 

## Methods

### NewNewTermsRuleResponseFields

`func NewNewTermsRuleResponseFields(historyWindowStart string, newTermsFields []string, query string, type_ string, language KqlQueryLanguage, ) *NewTermsRuleResponseFields`

NewNewTermsRuleResponseFields instantiates a new NewTermsRuleResponseFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewTermsRuleResponseFieldsWithDefaults

`func NewNewTermsRuleResponseFieldsWithDefaults() *NewTermsRuleResponseFields`

NewNewTermsRuleResponseFieldsWithDefaults instantiates a new NewTermsRuleResponseFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHistoryWindowStart

`func (o *NewTermsRuleResponseFields) GetHistoryWindowStart() string`

GetHistoryWindowStart returns the HistoryWindowStart field if non-nil, zero value otherwise.

### GetHistoryWindowStartOk

`func (o *NewTermsRuleResponseFields) GetHistoryWindowStartOk() (*string, bool)`

GetHistoryWindowStartOk returns a tuple with the HistoryWindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryWindowStart

`func (o *NewTermsRuleResponseFields) SetHistoryWindowStart(v string)`

SetHistoryWindowStart sets HistoryWindowStart field to given value.


### GetNewTermsFields

`func (o *NewTermsRuleResponseFields) GetNewTermsFields() []string`

GetNewTermsFields returns the NewTermsFields field if non-nil, zero value otherwise.

### GetNewTermsFieldsOk

`func (o *NewTermsRuleResponseFields) GetNewTermsFieldsOk() (*[]string, bool)`

GetNewTermsFieldsOk returns a tuple with the NewTermsFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTermsFields

`func (o *NewTermsRuleResponseFields) SetNewTermsFields(v []string)`

SetNewTermsFields sets NewTermsFields field to given value.


### GetQuery

`func (o *NewTermsRuleResponseFields) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *NewTermsRuleResponseFields) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *NewTermsRuleResponseFields) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetType

`func (o *NewTermsRuleResponseFields) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NewTermsRuleResponseFields) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NewTermsRuleResponseFields) SetType(v string)`

SetType sets Type field to given value.


### GetAlertSuppression

`func (o *NewTermsRuleResponseFields) GetAlertSuppression() AlertSuppression`

GetAlertSuppression returns the AlertSuppression field if non-nil, zero value otherwise.

### GetAlertSuppressionOk

`func (o *NewTermsRuleResponseFields) GetAlertSuppressionOk() (*AlertSuppression, bool)`

GetAlertSuppressionOk returns a tuple with the AlertSuppression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertSuppression

`func (o *NewTermsRuleResponseFields) SetAlertSuppression(v AlertSuppression)`

SetAlertSuppression sets AlertSuppression field to given value.

### HasAlertSuppression

`func (o *NewTermsRuleResponseFields) HasAlertSuppression() bool`

HasAlertSuppression returns a boolean if a field has been set.

### GetDataViewId

`func (o *NewTermsRuleResponseFields) GetDataViewId() string`

GetDataViewId returns the DataViewId field if non-nil, zero value otherwise.

### GetDataViewIdOk

`func (o *NewTermsRuleResponseFields) GetDataViewIdOk() (*string, bool)`

GetDataViewIdOk returns a tuple with the DataViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataViewId

`func (o *NewTermsRuleResponseFields) SetDataViewId(v string)`

SetDataViewId sets DataViewId field to given value.

### HasDataViewId

`func (o *NewTermsRuleResponseFields) HasDataViewId() bool`

HasDataViewId returns a boolean if a field has been set.

### GetFilters

`func (o *NewTermsRuleResponseFields) GetFilters() []interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *NewTermsRuleResponseFields) GetFiltersOk() (*[]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *NewTermsRuleResponseFields) SetFilters(v []interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *NewTermsRuleResponseFields) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetIndex

`func (o *NewTermsRuleResponseFields) GetIndex() []string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *NewTermsRuleResponseFields) GetIndexOk() (*[]string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *NewTermsRuleResponseFields) SetIndex(v []string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *NewTermsRuleResponseFields) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetResponseActions

`func (o *NewTermsRuleResponseFields) GetResponseActions() []ResponseAction`

GetResponseActions returns the ResponseActions field if non-nil, zero value otherwise.

### GetResponseActionsOk

`func (o *NewTermsRuleResponseFields) GetResponseActionsOk() (*[]ResponseAction, bool)`

GetResponseActionsOk returns a tuple with the ResponseActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseActions

`func (o *NewTermsRuleResponseFields) SetResponseActions(v []ResponseAction)`

SetResponseActions sets ResponseActions field to given value.

### HasResponseActions

`func (o *NewTermsRuleResponseFields) HasResponseActions() bool`

HasResponseActions returns a boolean if a field has been set.

### GetLanguage

`func (o *NewTermsRuleResponseFields) GetLanguage() KqlQueryLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *NewTermsRuleResponseFields) GetLanguageOk() (*KqlQueryLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *NewTermsRuleResponseFields) SetLanguage(v KqlQueryLanguage)`

SetLanguage sets Language field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


