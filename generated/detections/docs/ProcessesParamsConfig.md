# ProcessesParamsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Field to use instead of process.pid | 
**Overwrite** | Pointer to **bool** | Whether to overwrite field with process.pid | [optional] [default to true]

## Methods

### NewProcessesParamsConfig

`func NewProcessesParamsConfig(field string, ) *ProcessesParamsConfig`

NewProcessesParamsConfig instantiates a new ProcessesParamsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessesParamsConfigWithDefaults

`func NewProcessesParamsConfigWithDefaults() *ProcessesParamsConfig`

NewProcessesParamsConfigWithDefaults instantiates a new ProcessesParamsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *ProcessesParamsConfig) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *ProcessesParamsConfig) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *ProcessesParamsConfig) SetField(v string)`

SetField sets Field field to given value.


### GetOverwrite

`func (o *ProcessesParamsConfig) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *ProcessesParamsConfig) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *ProcessesParamsConfig) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *ProcessesParamsConfig) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


