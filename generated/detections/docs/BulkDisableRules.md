# BulkDisableRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**Ids** | Pointer to **[]string** | Array of rule IDs | [optional] 
**Query** | Pointer to **string** | Query to filter rules | [optional] 

## Methods

### NewBulkDisableRules

`func NewBulkDisableRules(action string, ) *BulkDisableRules`

NewBulkDisableRules instantiates a new BulkDisableRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkDisableRulesWithDefaults

`func NewBulkDisableRulesWithDefaults() *BulkDisableRules`

NewBulkDisableRulesWithDefaults instantiates a new BulkDisableRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *BulkDisableRules) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *BulkDisableRules) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *BulkDisableRules) SetAction(v string)`

SetAction sets Action field to given value.


### GetIds

`func (o *BulkDisableRules) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *BulkDisableRules) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *BulkDisableRules) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *BulkDisableRules) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetQuery

`func (o *BulkDisableRules) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *BulkDisableRules) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *BulkDisableRules) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *BulkDisableRules) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


