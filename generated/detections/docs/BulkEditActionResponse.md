# BulkEditActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**BulkEditActionResponseAttributes**](BulkEditActionResponseAttributes.md) |  | 
**Message** | Pointer to **string** |  | [optional] 
**RulesCount** | Pointer to **int32** |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewBulkEditActionResponse

`func NewBulkEditActionResponse(attributes BulkEditActionResponseAttributes, ) *BulkEditActionResponse`

NewBulkEditActionResponse instantiates a new BulkEditActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkEditActionResponseWithDefaults

`func NewBulkEditActionResponseWithDefaults() *BulkEditActionResponse`

NewBulkEditActionResponseWithDefaults instantiates a new BulkEditActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *BulkEditActionResponse) GetAttributes() BulkEditActionResponseAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BulkEditActionResponse) GetAttributesOk() (*BulkEditActionResponseAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BulkEditActionResponse) SetAttributes(v BulkEditActionResponseAttributes)`

SetAttributes sets Attributes field to given value.


### GetMessage

`func (o *BulkEditActionResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BulkEditActionResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BulkEditActionResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BulkEditActionResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRulesCount

`func (o *BulkEditActionResponse) GetRulesCount() int32`

GetRulesCount returns the RulesCount field if non-nil, zero value otherwise.

### GetRulesCountOk

`func (o *BulkEditActionResponse) GetRulesCountOk() (*int32, bool)`

GetRulesCountOk returns a tuple with the RulesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesCount

`func (o *BulkEditActionResponse) SetRulesCount(v int32)`

SetRulesCount sets RulesCount field to given value.

### HasRulesCount

`func (o *BulkEditActionResponse) HasRulesCount() bool`

HasRulesCount returns a boolean if a field has been set.

### GetStatusCode

`func (o *BulkEditActionResponse) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *BulkEditActionResponse) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *BulkEditActionResponse) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *BulkEditActionResponse) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetSuccess

`func (o *BulkEditActionResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *BulkEditActionResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *BulkEditActionResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *BulkEditActionResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


