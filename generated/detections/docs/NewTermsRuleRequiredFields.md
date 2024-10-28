# NewTermsRuleRequiredFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HistoryWindowStart** | **string** | A string that is not empty and does not contain only whitespace | 
**NewTermsFields** | **[]string** |  | 
**Query** | **string** |  | 
**Type** | **string** | Rule type | 

## Methods

### NewNewTermsRuleRequiredFields

`func NewNewTermsRuleRequiredFields(historyWindowStart string, newTermsFields []string, query string, type_ string, ) *NewTermsRuleRequiredFields`

NewNewTermsRuleRequiredFields instantiates a new NewTermsRuleRequiredFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewTermsRuleRequiredFieldsWithDefaults

`func NewNewTermsRuleRequiredFieldsWithDefaults() *NewTermsRuleRequiredFields`

NewNewTermsRuleRequiredFieldsWithDefaults instantiates a new NewTermsRuleRequiredFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHistoryWindowStart

`func (o *NewTermsRuleRequiredFields) GetHistoryWindowStart() string`

GetHistoryWindowStart returns the HistoryWindowStart field if non-nil, zero value otherwise.

### GetHistoryWindowStartOk

`func (o *NewTermsRuleRequiredFields) GetHistoryWindowStartOk() (*string, bool)`

GetHistoryWindowStartOk returns a tuple with the HistoryWindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryWindowStart

`func (o *NewTermsRuleRequiredFields) SetHistoryWindowStart(v string)`

SetHistoryWindowStart sets HistoryWindowStart field to given value.


### GetNewTermsFields

`func (o *NewTermsRuleRequiredFields) GetNewTermsFields() []string`

GetNewTermsFields returns the NewTermsFields field if non-nil, zero value otherwise.

### GetNewTermsFieldsOk

`func (o *NewTermsRuleRequiredFields) GetNewTermsFieldsOk() (*[]string, bool)`

GetNewTermsFieldsOk returns a tuple with the NewTermsFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTermsFields

`func (o *NewTermsRuleRequiredFields) SetNewTermsFields(v []string)`

SetNewTermsFields sets NewTermsFields field to given value.


### GetQuery

`func (o *NewTermsRuleRequiredFields) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *NewTermsRuleRequiredFields) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *NewTermsRuleRequiredFields) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetType

`func (o *NewTermsRuleRequiredFields) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NewTermsRuleRequiredFields) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NewTermsRuleRequiredFields) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


