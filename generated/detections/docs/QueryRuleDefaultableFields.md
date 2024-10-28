# QueryRuleDefaultableFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Language** | Pointer to [**KqlQueryLanguage**](KqlQueryLanguage.md) |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 

## Methods

### NewQueryRuleDefaultableFields

`func NewQueryRuleDefaultableFields() *QueryRuleDefaultableFields`

NewQueryRuleDefaultableFields instantiates a new QueryRuleDefaultableFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryRuleDefaultableFieldsWithDefaults

`func NewQueryRuleDefaultableFieldsWithDefaults() *QueryRuleDefaultableFields`

NewQueryRuleDefaultableFieldsWithDefaults instantiates a new QueryRuleDefaultableFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguage

`func (o *QueryRuleDefaultableFields) GetLanguage() KqlQueryLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *QueryRuleDefaultableFields) GetLanguageOk() (*KqlQueryLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *QueryRuleDefaultableFields) SetLanguage(v KqlQueryLanguage)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *QueryRuleDefaultableFields) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetQuery

`func (o *QueryRuleDefaultableFields) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *QueryRuleDefaultableFields) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *QueryRuleDefaultableFields) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *QueryRuleDefaultableFields) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


