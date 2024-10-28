# NewTermsRulePatchFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertSuppression** | Pointer to [**AlertSuppression**](AlertSuppression.md) |  | [optional] 
**DataViewId** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to **[]interface{}** |  | [optional] 
**Index** | Pointer to **[]string** |  | [optional] 
**ResponseActions** | Pointer to [**[]ResponseAction**](ResponseAction.md) |  | [optional] 
**Language** | Pointer to [**KqlQueryLanguage**](KqlQueryLanguage.md) |  | [optional] 
**HistoryWindowStart** | Pointer to **string** | A string that is not empty and does not contain only whitespace | [optional] 
**NewTermsFields** | Pointer to **[]string** |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | Rule type | [optional] 

## Methods

### NewNewTermsRulePatchFields

`func NewNewTermsRulePatchFields() *NewTermsRulePatchFields`

NewNewTermsRulePatchFields instantiates a new NewTermsRulePatchFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewTermsRulePatchFieldsWithDefaults

`func NewNewTermsRulePatchFieldsWithDefaults() *NewTermsRulePatchFields`

NewNewTermsRulePatchFieldsWithDefaults instantiates a new NewTermsRulePatchFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertSuppression

`func (o *NewTermsRulePatchFields) GetAlertSuppression() AlertSuppression`

GetAlertSuppression returns the AlertSuppression field if non-nil, zero value otherwise.

### GetAlertSuppressionOk

`func (o *NewTermsRulePatchFields) GetAlertSuppressionOk() (*AlertSuppression, bool)`

GetAlertSuppressionOk returns a tuple with the AlertSuppression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertSuppression

`func (o *NewTermsRulePatchFields) SetAlertSuppression(v AlertSuppression)`

SetAlertSuppression sets AlertSuppression field to given value.

### HasAlertSuppression

`func (o *NewTermsRulePatchFields) HasAlertSuppression() bool`

HasAlertSuppression returns a boolean if a field has been set.

### GetDataViewId

`func (o *NewTermsRulePatchFields) GetDataViewId() string`

GetDataViewId returns the DataViewId field if non-nil, zero value otherwise.

### GetDataViewIdOk

`func (o *NewTermsRulePatchFields) GetDataViewIdOk() (*string, bool)`

GetDataViewIdOk returns a tuple with the DataViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataViewId

`func (o *NewTermsRulePatchFields) SetDataViewId(v string)`

SetDataViewId sets DataViewId field to given value.

### HasDataViewId

`func (o *NewTermsRulePatchFields) HasDataViewId() bool`

HasDataViewId returns a boolean if a field has been set.

### GetFilters

`func (o *NewTermsRulePatchFields) GetFilters() []interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *NewTermsRulePatchFields) GetFiltersOk() (*[]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *NewTermsRulePatchFields) SetFilters(v []interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *NewTermsRulePatchFields) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetIndex

`func (o *NewTermsRulePatchFields) GetIndex() []string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *NewTermsRulePatchFields) GetIndexOk() (*[]string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *NewTermsRulePatchFields) SetIndex(v []string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *NewTermsRulePatchFields) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetResponseActions

`func (o *NewTermsRulePatchFields) GetResponseActions() []ResponseAction`

GetResponseActions returns the ResponseActions field if non-nil, zero value otherwise.

### GetResponseActionsOk

`func (o *NewTermsRulePatchFields) GetResponseActionsOk() (*[]ResponseAction, bool)`

GetResponseActionsOk returns a tuple with the ResponseActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseActions

`func (o *NewTermsRulePatchFields) SetResponseActions(v []ResponseAction)`

SetResponseActions sets ResponseActions field to given value.

### HasResponseActions

`func (o *NewTermsRulePatchFields) HasResponseActions() bool`

HasResponseActions returns a boolean if a field has been set.

### GetLanguage

`func (o *NewTermsRulePatchFields) GetLanguage() KqlQueryLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *NewTermsRulePatchFields) GetLanguageOk() (*KqlQueryLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *NewTermsRulePatchFields) SetLanguage(v KqlQueryLanguage)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *NewTermsRulePatchFields) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetHistoryWindowStart

`func (o *NewTermsRulePatchFields) GetHistoryWindowStart() string`

GetHistoryWindowStart returns the HistoryWindowStart field if non-nil, zero value otherwise.

### GetHistoryWindowStartOk

`func (o *NewTermsRulePatchFields) GetHistoryWindowStartOk() (*string, bool)`

GetHistoryWindowStartOk returns a tuple with the HistoryWindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryWindowStart

`func (o *NewTermsRulePatchFields) SetHistoryWindowStart(v string)`

SetHistoryWindowStart sets HistoryWindowStart field to given value.

### HasHistoryWindowStart

`func (o *NewTermsRulePatchFields) HasHistoryWindowStart() bool`

HasHistoryWindowStart returns a boolean if a field has been set.

### GetNewTermsFields

`func (o *NewTermsRulePatchFields) GetNewTermsFields() []string`

GetNewTermsFields returns the NewTermsFields field if non-nil, zero value otherwise.

### GetNewTermsFieldsOk

`func (o *NewTermsRulePatchFields) GetNewTermsFieldsOk() (*[]string, bool)`

GetNewTermsFieldsOk returns a tuple with the NewTermsFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTermsFields

`func (o *NewTermsRulePatchFields) SetNewTermsFields(v []string)`

SetNewTermsFields sets NewTermsFields field to given value.

### HasNewTermsFields

`func (o *NewTermsRulePatchFields) HasNewTermsFields() bool`

HasNewTermsFields returns a boolean if a field has been set.

### GetQuery

`func (o *NewTermsRulePatchFields) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *NewTermsRulePatchFields) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *NewTermsRulePatchFields) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *NewTermsRulePatchFields) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetType

`func (o *NewTermsRulePatchFields) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NewTermsRulePatchFields) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NewTermsRulePatchFields) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NewTermsRulePatchFields) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


