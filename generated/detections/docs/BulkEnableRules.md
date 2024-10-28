# BulkEnableRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**Ids** | Pointer to **[]string** | Array of rule IDs | [optional] 
**Query** | Pointer to **string** | Query to filter rules | [optional] 

## Methods

### NewBulkEnableRules

`func NewBulkEnableRules(action string, ) *BulkEnableRules`

NewBulkEnableRules instantiates a new BulkEnableRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkEnableRulesWithDefaults

`func NewBulkEnableRulesWithDefaults() *BulkEnableRules`

NewBulkEnableRulesWithDefaults instantiates a new BulkEnableRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *BulkEnableRules) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *BulkEnableRules) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *BulkEnableRules) SetAction(v string)`

SetAction sets Action field to given value.


### GetIds

`func (o *BulkEnableRules) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *BulkEnableRules) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *BulkEnableRules) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *BulkEnableRules) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetQuery

`func (o *BulkEnableRules) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *BulkEnableRules) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *BulkEnableRules) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *BulkEnableRules) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


