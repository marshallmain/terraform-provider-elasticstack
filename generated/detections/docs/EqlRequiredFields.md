# EqlRequiredFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Language** | [**EqlQueryLanguage**](EqlQueryLanguage.md) |  | 
**Query** | **string** |  | 
**Type** | **string** | Rule type | 

## Methods

### NewEqlRequiredFields

`func NewEqlRequiredFields(language EqlQueryLanguage, query string, type_ string, ) *EqlRequiredFields`

NewEqlRequiredFields instantiates a new EqlRequiredFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEqlRequiredFieldsWithDefaults

`func NewEqlRequiredFieldsWithDefaults() *EqlRequiredFields`

NewEqlRequiredFieldsWithDefaults instantiates a new EqlRequiredFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguage

`func (o *EqlRequiredFields) GetLanguage() EqlQueryLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *EqlRequiredFields) GetLanguageOk() (*EqlQueryLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *EqlRequiredFields) SetLanguage(v EqlQueryLanguage)`

SetLanguage sets Language field to given value.


### GetQuery

`func (o *EqlRequiredFields) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *EqlRequiredFields) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *EqlRequiredFields) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetType

`func (o *EqlRequiredFields) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EqlRequiredFields) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EqlRequiredFields) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


