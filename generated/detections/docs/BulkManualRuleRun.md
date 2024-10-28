# BulkManualRuleRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**Ids** | Pointer to **[]string** | Array of rule IDs | [optional] 
**Query** | Pointer to **string** | Query to filter rules | [optional] 
**Run** | [**BulkManualRuleRunRun**](BulkManualRuleRunRun.md) |  | 

## Methods

### NewBulkManualRuleRun

`func NewBulkManualRuleRun(action string, run BulkManualRuleRunRun, ) *BulkManualRuleRun`

NewBulkManualRuleRun instantiates a new BulkManualRuleRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkManualRuleRunWithDefaults

`func NewBulkManualRuleRunWithDefaults() *BulkManualRuleRun`

NewBulkManualRuleRunWithDefaults instantiates a new BulkManualRuleRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *BulkManualRuleRun) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *BulkManualRuleRun) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *BulkManualRuleRun) SetAction(v string)`

SetAction sets Action field to given value.


### GetIds

`func (o *BulkManualRuleRun) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *BulkManualRuleRun) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *BulkManualRuleRun) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *BulkManualRuleRun) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetQuery

`func (o *BulkManualRuleRun) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *BulkManualRuleRun) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *BulkManualRuleRun) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *BulkManualRuleRun) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetRun

`func (o *BulkManualRuleRun) GetRun() BulkManualRuleRunRun`

GetRun returns the Run field if non-nil, zero value otherwise.

### GetRunOk

`func (o *BulkManualRuleRun) GetRunOk() (*BulkManualRuleRunRun, bool)`

GetRunOk returns a tuple with the Run field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRun

`func (o *BulkManualRuleRun) SetRun(v BulkManualRuleRunRun)`

SetRun sets Run field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


