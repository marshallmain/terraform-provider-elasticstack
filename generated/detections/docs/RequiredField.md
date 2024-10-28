# RequiredField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ecs** | **bool** | Whether the field is an ECS field | 
**Name** | **string** | A string that is not empty and does not contain only whitespace | 
**Type** | **string** | A string that is not empty and does not contain only whitespace | 

## Methods

### NewRequiredField

`func NewRequiredField(ecs bool, name string, type_ string, ) *RequiredField`

NewRequiredField instantiates a new RequiredField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequiredFieldWithDefaults

`func NewRequiredFieldWithDefaults() *RequiredField`

NewRequiredFieldWithDefaults instantiates a new RequiredField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEcs

`func (o *RequiredField) GetEcs() bool`

GetEcs returns the Ecs field if non-nil, zero value otherwise.

### GetEcsOk

`func (o *RequiredField) GetEcsOk() (*bool, bool)`

GetEcsOk returns a tuple with the Ecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcs

`func (o *RequiredField) SetEcs(v bool)`

SetEcs sets Ecs field to given value.


### GetName

`func (o *RequiredField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequiredField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequiredField) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *RequiredField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequiredField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequiredField) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


