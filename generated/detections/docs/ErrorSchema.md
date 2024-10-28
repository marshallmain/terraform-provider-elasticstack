# ErrorSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | [**ErrorSchemaError**](ErrorSchemaError.md) |  | 
**Id** | Pointer to **string** |  | [optional] 
**ItemId** | Pointer to **string** |  | [optional] 
**ListId** | Pointer to **string** |  | [optional] 
**RuleId** | Pointer to **string** | Could be any string, not necessarily a UUID | [optional] 

## Methods

### NewErrorSchema

`func NewErrorSchema(error_ ErrorSchemaError, ) *ErrorSchema`

NewErrorSchema instantiates a new ErrorSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorSchemaWithDefaults

`func NewErrorSchemaWithDefaults() *ErrorSchema`

NewErrorSchemaWithDefaults instantiates a new ErrorSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ErrorSchema) GetError() ErrorSchemaError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ErrorSchema) GetErrorOk() (*ErrorSchemaError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ErrorSchema) SetError(v ErrorSchemaError)`

SetError sets Error field to given value.


### GetId

`func (o *ErrorSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ErrorSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ErrorSchema) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ErrorSchema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItemId

`func (o *ErrorSchema) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ErrorSchema) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ErrorSchema) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ErrorSchema) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetListId

`func (o *ErrorSchema) GetListId() string`

GetListId returns the ListId field if non-nil, zero value otherwise.

### GetListIdOk

`func (o *ErrorSchema) GetListIdOk() (*string, bool)`

GetListIdOk returns a tuple with the ListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListId

`func (o *ErrorSchema) SetListId(v string)`

SetListId sets ListId field to given value.

### HasListId

`func (o *ErrorSchema) HasListId() bool`

HasListId returns a boolean if a field has been set.

### GetRuleId

`func (o *ErrorSchema) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *ErrorSchema) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *ErrorSchema) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *ErrorSchema) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


