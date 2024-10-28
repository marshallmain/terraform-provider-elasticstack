# EsqlRuleRequiredFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Language** | [**EsqlQueryLanguage**](EsqlQueryLanguage.md) |  | 
**Query** | **string** |  | 
**Type** | **string** | Rule type | 

## Methods

### NewEsqlRuleRequiredFields

`func NewEsqlRuleRequiredFields(language EsqlQueryLanguage, query string, type_ string, ) *EsqlRuleRequiredFields`

NewEsqlRuleRequiredFields instantiates a new EsqlRuleRequiredFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsqlRuleRequiredFieldsWithDefaults

`func NewEsqlRuleRequiredFieldsWithDefaults() *EsqlRuleRequiredFields`

NewEsqlRuleRequiredFieldsWithDefaults instantiates a new EsqlRuleRequiredFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguage

`func (o *EsqlRuleRequiredFields) GetLanguage() EsqlQueryLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *EsqlRuleRequiredFields) GetLanguageOk() (*EsqlQueryLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *EsqlRuleRequiredFields) SetLanguage(v EsqlQueryLanguage)`

SetLanguage sets Language field to given value.


### GetQuery

`func (o *EsqlRuleRequiredFields) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *EsqlRuleRequiredFields) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *EsqlRuleRequiredFields) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetType

`func (o *EsqlRuleRequiredFields) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EsqlRuleRequiredFields) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EsqlRuleRequiredFields) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


