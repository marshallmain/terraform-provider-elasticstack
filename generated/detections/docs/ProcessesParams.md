# ProcessesParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | **string** |  | 
**Comment** | Pointer to **string** |  | [optional] 
**Config** | [**ProcessesParamsConfig**](ProcessesParamsConfig.md) |  | 

## Methods

### NewProcessesParams

`func NewProcessesParams(command string, config ProcessesParamsConfig, ) *ProcessesParams`

NewProcessesParams instantiates a new ProcessesParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessesParamsWithDefaults

`func NewProcessesParamsWithDefaults() *ProcessesParams`

NewProcessesParamsWithDefaults instantiates a new ProcessesParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *ProcessesParams) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *ProcessesParams) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *ProcessesParams) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetComment

`func (o *ProcessesParams) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ProcessesParams) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ProcessesParams) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ProcessesParams) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetConfig

`func (o *ProcessesParams) GetConfig() ProcessesParamsConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ProcessesParams) GetConfigOk() (*ProcessesParamsConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ProcessesParams) SetConfig(v ProcessesParamsConfig)`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


