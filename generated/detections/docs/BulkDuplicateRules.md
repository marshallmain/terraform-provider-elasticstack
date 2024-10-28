# BulkDuplicateRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**Duplicate** | Pointer to [**BulkDuplicateRulesDuplicate**](BulkDuplicateRulesDuplicate.md) |  | [optional] 
**Ids** | Pointer to **[]string** | Array of rule IDs | [optional] 
**Query** | Pointer to **string** | Query to filter rules | [optional] 

## Methods

### NewBulkDuplicateRules

`func NewBulkDuplicateRules(action string, ) *BulkDuplicateRules`

NewBulkDuplicateRules instantiates a new BulkDuplicateRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkDuplicateRulesWithDefaults

`func NewBulkDuplicateRulesWithDefaults() *BulkDuplicateRules`

NewBulkDuplicateRulesWithDefaults instantiates a new BulkDuplicateRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *BulkDuplicateRules) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *BulkDuplicateRules) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *BulkDuplicateRules) SetAction(v string)`

SetAction sets Action field to given value.


### GetDuplicate

`func (o *BulkDuplicateRules) GetDuplicate() BulkDuplicateRulesDuplicate`

GetDuplicate returns the Duplicate field if non-nil, zero value otherwise.

### GetDuplicateOk

`func (o *BulkDuplicateRules) GetDuplicateOk() (*BulkDuplicateRulesDuplicate, bool)`

GetDuplicateOk returns a tuple with the Duplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicate

`func (o *BulkDuplicateRules) SetDuplicate(v BulkDuplicateRulesDuplicate)`

SetDuplicate sets Duplicate field to given value.

### HasDuplicate

`func (o *BulkDuplicateRules) HasDuplicate() bool`

HasDuplicate returns a boolean if a field has been set.

### GetIds

`func (o *BulkDuplicateRules) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *BulkDuplicateRules) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *BulkDuplicateRules) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *BulkDuplicateRules) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetQuery

`func (o *BulkDuplicateRules) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *BulkDuplicateRules) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *BulkDuplicateRules) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *BulkDuplicateRules) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


