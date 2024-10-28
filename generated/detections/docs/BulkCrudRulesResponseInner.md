# BulkCrudRulesResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | [**ErrorSchemaError**](ErrorSchemaError.md) |  | 
**Id** | Pointer to **string** |  | [optional] 
**ItemId** | Pointer to **string** |  | [optional] 
**ListId** | Pointer to **string** |  | [optional] 
**RuleId** | Pointer to **string** | Could be any string, not necessarily a UUID | [optional] 

## Methods

### NewBulkCrudRulesResponseInner

`func NewBulkCrudRulesResponseInner(error_ ErrorSchemaError, ) *BulkCrudRulesResponseInner`

NewBulkCrudRulesResponseInner instantiates a new BulkCrudRulesResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkCrudRulesResponseInnerWithDefaults

`func NewBulkCrudRulesResponseInnerWithDefaults() *BulkCrudRulesResponseInner`

NewBulkCrudRulesResponseInnerWithDefaults instantiates a new BulkCrudRulesResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *BulkCrudRulesResponseInner) GetError() ErrorSchemaError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BulkCrudRulesResponseInner) GetErrorOk() (*ErrorSchemaError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BulkCrudRulesResponseInner) SetError(v ErrorSchemaError)`

SetError sets Error field to given value.


### GetId

`func (o *BulkCrudRulesResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkCrudRulesResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkCrudRulesResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BulkCrudRulesResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItemId

`func (o *BulkCrudRulesResponseInner) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *BulkCrudRulesResponseInner) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *BulkCrudRulesResponseInner) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *BulkCrudRulesResponseInner) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetListId

`func (o *BulkCrudRulesResponseInner) GetListId() string`

GetListId returns the ListId field if non-nil, zero value otherwise.

### GetListIdOk

`func (o *BulkCrudRulesResponseInner) GetListIdOk() (*string, bool)`

GetListIdOk returns a tuple with the ListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListId

`func (o *BulkCrudRulesResponseInner) SetListId(v string)`

SetListId sets ListId field to given value.

### HasListId

`func (o *BulkCrudRulesResponseInner) HasListId() bool`

HasListId returns a boolean if a field has been set.

### GetRuleId

`func (o *BulkCrudRulesResponseInner) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *BulkCrudRulesResponseInner) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *BulkCrudRulesResponseInner) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *BulkCrudRulesResponseInner) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


