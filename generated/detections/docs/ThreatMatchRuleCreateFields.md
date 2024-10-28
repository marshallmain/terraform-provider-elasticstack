# ThreatMatchRuleCreateFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** |  | 
**ThreatIndex** | **[]string** |  | 
**ThreatMapping** | [**[]ThreatMappingInner**](ThreatMappingInner.md) |  | 
**ThreatQuery** | **string** | Query to run | 
**Type** | **string** | Rule type | 
**AlertSuppression** | Pointer to [**AlertSuppression**](AlertSuppression.md) |  | [optional] 
**ConcurrentSearches** | Pointer to **int32** |  | [optional] 
**DataViewId** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to **[]interface{}** |  | [optional] 
**Index** | Pointer to **[]string** |  | [optional] 
**ItemsPerSearch** | Pointer to **int32** |  | [optional] 
**SavedId** | Pointer to **string** |  | [optional] 
**ThreatFilters** | Pointer to **[]interface{}** |  | [optional] 
**ThreatIndicatorPath** | Pointer to **string** | Defines the path to the threat indicator in the indicator documents (optional) | [optional] 
**ThreatLanguage** | Pointer to [**KqlQueryLanguage**](KqlQueryLanguage.md) |  | [optional] 
**Language** | Pointer to [**KqlQueryLanguage**](KqlQueryLanguage.md) |  | [optional] 

## Methods

### NewThreatMatchRuleCreateFields

`func NewThreatMatchRuleCreateFields(query string, threatIndex []string, threatMapping []ThreatMappingInner, threatQuery string, type_ string, ) *ThreatMatchRuleCreateFields`

NewThreatMatchRuleCreateFields instantiates a new ThreatMatchRuleCreateFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreatMatchRuleCreateFieldsWithDefaults

`func NewThreatMatchRuleCreateFieldsWithDefaults() *ThreatMatchRuleCreateFields`

NewThreatMatchRuleCreateFieldsWithDefaults instantiates a new ThreatMatchRuleCreateFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *ThreatMatchRuleCreateFields) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ThreatMatchRuleCreateFields) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ThreatMatchRuleCreateFields) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetThreatIndex

`func (o *ThreatMatchRuleCreateFields) GetThreatIndex() []string`

GetThreatIndex returns the ThreatIndex field if non-nil, zero value otherwise.

### GetThreatIndexOk

`func (o *ThreatMatchRuleCreateFields) GetThreatIndexOk() (*[]string, bool)`

GetThreatIndexOk returns a tuple with the ThreatIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatIndex

`func (o *ThreatMatchRuleCreateFields) SetThreatIndex(v []string)`

SetThreatIndex sets ThreatIndex field to given value.


### GetThreatMapping

`func (o *ThreatMatchRuleCreateFields) GetThreatMapping() []ThreatMappingInner`

GetThreatMapping returns the ThreatMapping field if non-nil, zero value otherwise.

### GetThreatMappingOk

`func (o *ThreatMatchRuleCreateFields) GetThreatMappingOk() (*[]ThreatMappingInner, bool)`

GetThreatMappingOk returns a tuple with the ThreatMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatMapping

`func (o *ThreatMatchRuleCreateFields) SetThreatMapping(v []ThreatMappingInner)`

SetThreatMapping sets ThreatMapping field to given value.


### GetThreatQuery

`func (o *ThreatMatchRuleCreateFields) GetThreatQuery() string`

GetThreatQuery returns the ThreatQuery field if non-nil, zero value otherwise.

### GetThreatQueryOk

`func (o *ThreatMatchRuleCreateFields) GetThreatQueryOk() (*string, bool)`

GetThreatQueryOk returns a tuple with the ThreatQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatQuery

`func (o *ThreatMatchRuleCreateFields) SetThreatQuery(v string)`

SetThreatQuery sets ThreatQuery field to given value.


### GetType

`func (o *ThreatMatchRuleCreateFields) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ThreatMatchRuleCreateFields) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ThreatMatchRuleCreateFields) SetType(v string)`

SetType sets Type field to given value.


### GetAlertSuppression

`func (o *ThreatMatchRuleCreateFields) GetAlertSuppression() AlertSuppression`

GetAlertSuppression returns the AlertSuppression field if non-nil, zero value otherwise.

### GetAlertSuppressionOk

`func (o *ThreatMatchRuleCreateFields) GetAlertSuppressionOk() (*AlertSuppression, bool)`

GetAlertSuppressionOk returns a tuple with the AlertSuppression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertSuppression

`func (o *ThreatMatchRuleCreateFields) SetAlertSuppression(v AlertSuppression)`

SetAlertSuppression sets AlertSuppression field to given value.

### HasAlertSuppression

`func (o *ThreatMatchRuleCreateFields) HasAlertSuppression() bool`

HasAlertSuppression returns a boolean if a field has been set.

### GetConcurrentSearches

`func (o *ThreatMatchRuleCreateFields) GetConcurrentSearches() int32`

GetConcurrentSearches returns the ConcurrentSearches field if non-nil, zero value otherwise.

### GetConcurrentSearchesOk

`func (o *ThreatMatchRuleCreateFields) GetConcurrentSearchesOk() (*int32, bool)`

GetConcurrentSearchesOk returns a tuple with the ConcurrentSearches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrentSearches

`func (o *ThreatMatchRuleCreateFields) SetConcurrentSearches(v int32)`

SetConcurrentSearches sets ConcurrentSearches field to given value.

### HasConcurrentSearches

`func (o *ThreatMatchRuleCreateFields) HasConcurrentSearches() bool`

HasConcurrentSearches returns a boolean if a field has been set.

### GetDataViewId

`func (o *ThreatMatchRuleCreateFields) GetDataViewId() string`

GetDataViewId returns the DataViewId field if non-nil, zero value otherwise.

### GetDataViewIdOk

`func (o *ThreatMatchRuleCreateFields) GetDataViewIdOk() (*string, bool)`

GetDataViewIdOk returns a tuple with the DataViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataViewId

`func (o *ThreatMatchRuleCreateFields) SetDataViewId(v string)`

SetDataViewId sets DataViewId field to given value.

### HasDataViewId

`func (o *ThreatMatchRuleCreateFields) HasDataViewId() bool`

HasDataViewId returns a boolean if a field has been set.

### GetFilters

`func (o *ThreatMatchRuleCreateFields) GetFilters() []interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ThreatMatchRuleCreateFields) GetFiltersOk() (*[]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ThreatMatchRuleCreateFields) SetFilters(v []interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ThreatMatchRuleCreateFields) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetIndex

`func (o *ThreatMatchRuleCreateFields) GetIndex() []string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ThreatMatchRuleCreateFields) GetIndexOk() (*[]string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ThreatMatchRuleCreateFields) SetIndex(v []string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *ThreatMatchRuleCreateFields) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetItemsPerSearch

`func (o *ThreatMatchRuleCreateFields) GetItemsPerSearch() int32`

GetItemsPerSearch returns the ItemsPerSearch field if non-nil, zero value otherwise.

### GetItemsPerSearchOk

`func (o *ThreatMatchRuleCreateFields) GetItemsPerSearchOk() (*int32, bool)`

GetItemsPerSearchOk returns a tuple with the ItemsPerSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsPerSearch

`func (o *ThreatMatchRuleCreateFields) SetItemsPerSearch(v int32)`

SetItemsPerSearch sets ItemsPerSearch field to given value.

### HasItemsPerSearch

`func (o *ThreatMatchRuleCreateFields) HasItemsPerSearch() bool`

HasItemsPerSearch returns a boolean if a field has been set.

### GetSavedId

`func (o *ThreatMatchRuleCreateFields) GetSavedId() string`

GetSavedId returns the SavedId field if non-nil, zero value otherwise.

### GetSavedIdOk

`func (o *ThreatMatchRuleCreateFields) GetSavedIdOk() (*string, bool)`

GetSavedIdOk returns a tuple with the SavedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedId

`func (o *ThreatMatchRuleCreateFields) SetSavedId(v string)`

SetSavedId sets SavedId field to given value.

### HasSavedId

`func (o *ThreatMatchRuleCreateFields) HasSavedId() bool`

HasSavedId returns a boolean if a field has been set.

### GetThreatFilters

`func (o *ThreatMatchRuleCreateFields) GetThreatFilters() []interface{}`

GetThreatFilters returns the ThreatFilters field if non-nil, zero value otherwise.

### GetThreatFiltersOk

`func (o *ThreatMatchRuleCreateFields) GetThreatFiltersOk() (*[]interface{}, bool)`

GetThreatFiltersOk returns a tuple with the ThreatFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatFilters

`func (o *ThreatMatchRuleCreateFields) SetThreatFilters(v []interface{})`

SetThreatFilters sets ThreatFilters field to given value.

### HasThreatFilters

`func (o *ThreatMatchRuleCreateFields) HasThreatFilters() bool`

HasThreatFilters returns a boolean if a field has been set.

### GetThreatIndicatorPath

`func (o *ThreatMatchRuleCreateFields) GetThreatIndicatorPath() string`

GetThreatIndicatorPath returns the ThreatIndicatorPath field if non-nil, zero value otherwise.

### GetThreatIndicatorPathOk

`func (o *ThreatMatchRuleCreateFields) GetThreatIndicatorPathOk() (*string, bool)`

GetThreatIndicatorPathOk returns a tuple with the ThreatIndicatorPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatIndicatorPath

`func (o *ThreatMatchRuleCreateFields) SetThreatIndicatorPath(v string)`

SetThreatIndicatorPath sets ThreatIndicatorPath field to given value.

### HasThreatIndicatorPath

`func (o *ThreatMatchRuleCreateFields) HasThreatIndicatorPath() bool`

HasThreatIndicatorPath returns a boolean if a field has been set.

### GetThreatLanguage

`func (o *ThreatMatchRuleCreateFields) GetThreatLanguage() KqlQueryLanguage`

GetThreatLanguage returns the ThreatLanguage field if non-nil, zero value otherwise.

### GetThreatLanguageOk

`func (o *ThreatMatchRuleCreateFields) GetThreatLanguageOk() (*KqlQueryLanguage, bool)`

GetThreatLanguageOk returns a tuple with the ThreatLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatLanguage

`func (o *ThreatMatchRuleCreateFields) SetThreatLanguage(v KqlQueryLanguage)`

SetThreatLanguage sets ThreatLanguage field to given value.

### HasThreatLanguage

`func (o *ThreatMatchRuleCreateFields) HasThreatLanguage() bool`

HasThreatLanguage returns a boolean if a field has been set.

### GetLanguage

`func (o *ThreatMatchRuleCreateFields) GetLanguage() KqlQueryLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ThreatMatchRuleCreateFields) GetLanguageOk() (*KqlQueryLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ThreatMatchRuleCreateFields) SetLanguage(v KqlQueryLanguage)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ThreatMatchRuleCreateFields) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


