# RuleSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsCustomized** | **bool** | Determines whether an external/prebuilt rule has been customized by the user (i.e. any of its fields have been modified and diverged from the base value). | 
**Type** | **string** |  | 

## Methods

### NewRuleSource

`func NewRuleSource(isCustomized bool, type_ string, ) *RuleSource`

NewRuleSource instantiates a new RuleSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleSourceWithDefaults

`func NewRuleSourceWithDefaults() *RuleSource`

NewRuleSourceWithDefaults instantiates a new RuleSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsCustomized

`func (o *RuleSource) GetIsCustomized() bool`

GetIsCustomized returns the IsCustomized field if non-nil, zero value otherwise.

### GetIsCustomizedOk

`func (o *RuleSource) GetIsCustomizedOk() (*bool, bool)`

GetIsCustomizedOk returns a tuple with the IsCustomized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomized

`func (o *RuleSource) SetIsCustomized(v bool)`

SetIsCustomized sets IsCustomized field to given value.


### GetType

`func (o *RuleSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RuleSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RuleSource) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


