# ExternalRuleSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsCustomized** | **bool** | Determines whether an external/prebuilt rule has been customized by the user (i.e. any of its fields have been modified and diverged from the base value). | 
**Type** | **string** |  | 

## Methods

### NewExternalRuleSource

`func NewExternalRuleSource(isCustomized bool, type_ string, ) *ExternalRuleSource`

NewExternalRuleSource instantiates a new ExternalRuleSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalRuleSourceWithDefaults

`func NewExternalRuleSourceWithDefaults() *ExternalRuleSource`

NewExternalRuleSourceWithDefaults instantiates a new ExternalRuleSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsCustomized

`func (o *ExternalRuleSource) GetIsCustomized() bool`

GetIsCustomized returns the IsCustomized field if non-nil, zero value otherwise.

### GetIsCustomizedOk

`func (o *ExternalRuleSource) GetIsCustomizedOk() (*bool, bool)`

GetIsCustomizedOk returns a tuple with the IsCustomized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomized

`func (o *ExternalRuleSource) SetIsCustomized(v bool)`

SetIsCustomized sets IsCustomized field to given value.


### GetType

`func (o *ExternalRuleSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalRuleSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalRuleSource) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


