# WarningSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionPath** | **string** |  | 
**ButtonLabel** | Pointer to **string** |  | [optional] 
**Message** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewWarningSchema

`func NewWarningSchema(actionPath string, message string, type_ string, ) *WarningSchema`

NewWarningSchema instantiates a new WarningSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarningSchemaWithDefaults

`func NewWarningSchemaWithDefaults() *WarningSchema`

NewWarningSchemaWithDefaults instantiates a new WarningSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionPath

`func (o *WarningSchema) GetActionPath() string`

GetActionPath returns the ActionPath field if non-nil, zero value otherwise.

### GetActionPathOk

`func (o *WarningSchema) GetActionPathOk() (*string, bool)`

GetActionPathOk returns a tuple with the ActionPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionPath

`func (o *WarningSchema) SetActionPath(v string)`

SetActionPath sets ActionPath field to given value.


### GetButtonLabel

`func (o *WarningSchema) GetButtonLabel() string`

GetButtonLabel returns the ButtonLabel field if non-nil, zero value otherwise.

### GetButtonLabelOk

`func (o *WarningSchema) GetButtonLabelOk() (*string, bool)`

GetButtonLabelOk returns a tuple with the ButtonLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonLabel

`func (o *WarningSchema) SetButtonLabel(v string)`

SetButtonLabel sets ButtonLabel field to given value.

### HasButtonLabel

`func (o *WarningSchema) HasButtonLabel() bool`

HasButtonLabel returns a boolean if a field has been set.

### GetMessage

`func (o *WarningSchema) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WarningSchema) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WarningSchema) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetType

`func (o *WarningSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WarningSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WarningSchema) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


