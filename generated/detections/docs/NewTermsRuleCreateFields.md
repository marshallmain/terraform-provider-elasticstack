# NewTermsRuleCreateFields

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
**Language** | Pointer to [**KqlQueryLanguage**](KqlQueryLanguage.md) |  | [optional] 

## Methods

### NewNewTermsRuleCreateFields

`func NewNewTermsRuleCreateFields(historyWindowStart string, newTermsFields []string, query string, type_ string, ) *NewTermsRuleCreateFields`

NewNewTermsRuleCreateFields instantiates a new NewTermsRuleCreateFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewTermsRuleCreateFieldsWithDefaults

`func NewNewTermsRuleCreateFieldsWithDefaults() *NewTermsRuleCreateFields`

NewNewTermsRuleCreateFieldsWithDefaults instantiates a new NewTermsRuleCreateFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHistoryWindowStart

`func (o *NewTermsRuleCreateFields) GetHistoryWindowStart() string`

GetHistoryWindowStart returns the HistoryWindowStart field if non-nil, zero value otherwise.

### GetHistoryWindowStartOk

`func (o *NewTermsRuleCreateFields) GetHistoryWindowStartOk() (*string, bool)`

GetHistoryWindowStartOk returns a tuple with the HistoryWindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryWindowStart

`func (o *NewTermsRuleCreateFields) SetHistoryWindowStart(v string)`

SetHistoryWindowStart sets HistoryWindowStart field to given value.


### GetNewTermsFields

`func (o *NewTermsRuleCreateFields) GetNewTermsFields() []string`

GetNewTermsFields returns the NewTermsFields field if non-nil, zero value otherwise.

### GetNewTermsFieldsOk

`func (o *NewTermsRuleCreateFields) GetNewTermsFieldsOk() (*[]string, bool)`

GetNewTermsFieldsOk returns a tuple with the NewTermsFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTermsFields

`func (o *NewTermsRuleCreateFields) SetNewTermsFields(v []string)`

SetNewTermsFields sets NewTermsFields field to given value.


### GetQuery

`func (o *NewTermsRuleCreateFields) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *NewTermsRuleCreateFields) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *NewTermsRuleCreateFields) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetType

`func (o *NewTermsRuleCreateFields) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NewTermsRuleCreateFields) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NewTermsRuleCreateFields) SetType(v string)`

SetType sets Type field to given value.


### GetAlertSuppression

`func (o *NewTermsRuleCreateFields) GetAlertSuppression() AlertSuppression`

GetAlertSuppression returns the AlertSuppression field if non-nil, zero value otherwise.

### GetAlertSuppressionOk

`func (o *NewTermsRuleCreateFields) GetAlertSuppressionOk() (*AlertSuppression, bool)`

GetAlertSuppressionOk returns a tuple with the AlertSuppression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertSuppression

`func (o *NewTermsRuleCreateFields) SetAlertSuppression(v AlertSuppression)`

SetAlertSuppression sets AlertSuppression field to given value.

### HasAlertSuppression

`func (o *NewTermsRuleCreateFields) HasAlertSuppression() bool`

HasAlertSuppression returns a boolean if a field has been set.

### GetDataViewId

`func (o *NewTermsRuleCreateFields) GetDataViewId() string`

GetDataViewId returns the DataViewId field if non-nil, zero value otherwise.

### GetDataViewIdOk

`func (o *NewTermsRuleCreateFields) GetDataViewIdOk() (*string, bool)`

GetDataViewIdOk returns a tuple with the DataViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataViewId

`func (o *NewTermsRuleCreateFields) SetDataViewId(v string)`

SetDataViewId sets DataViewId field to given value.

### HasDataViewId

`func (o *NewTermsRuleCreateFields) HasDataViewId() bool`

HasDataViewId returns a boolean if a field has been set.

### GetFilters

`func (o *NewTermsRuleCreateFields) GetFilters() []interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *NewTermsRuleCreateFields) GetFiltersOk() (*[]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *NewTermsRuleCreateFields) SetFilters(v []interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *NewTermsRuleCreateFields) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetIndex

`func (o *NewTermsRuleCreateFields) GetIndex() []string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *NewTermsRuleCreateFields) GetIndexOk() (*[]string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *NewTermsRuleCreateFields) SetIndex(v []string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *NewTermsRuleCreateFields) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetResponseActions

`func (o *NewTermsRuleCreateFields) GetResponseActions() []ResponseAction`

GetResponseActions returns the ResponseActions field if non-nil, zero value otherwise.

### GetResponseActionsOk

`func (o *NewTermsRuleCreateFields) GetResponseActionsOk() (*[]ResponseAction, bool)`

GetResponseActionsOk returns a tuple with the ResponseActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseActions

`func (o *NewTermsRuleCreateFields) SetResponseActions(v []ResponseAction)`

SetResponseActions sets ResponseActions field to given value.

### HasResponseActions

`func (o *NewTermsRuleCreateFields) HasResponseActions() bool`

HasResponseActions returns a boolean if a field has been set.

### GetLanguage

`func (o *NewTermsRuleCreateFields) GetLanguage() KqlQueryLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *NewTermsRuleCreateFields) GetLanguageOk() (*KqlQueryLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *NewTermsRuleCreateFields) SetLanguage(v KqlQueryLanguage)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *NewTermsRuleCreateFields) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


