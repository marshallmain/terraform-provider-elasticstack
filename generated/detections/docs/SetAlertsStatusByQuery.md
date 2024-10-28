# SetAlertsStatusByQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conflicts** | Pointer to **string** |  | [optional] [default to "abort"]
**Query** | **map[string]interface{}** |  | 
**Status** | [**AlertStatus**](AlertStatus.md) |  | 

## Methods

### NewSetAlertsStatusByQuery

`func NewSetAlertsStatusByQuery(query map[string]interface{}, status AlertStatus, ) *SetAlertsStatusByQuery`

NewSetAlertsStatusByQuery instantiates a new SetAlertsStatusByQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetAlertsStatusByQueryWithDefaults

`func NewSetAlertsStatusByQueryWithDefaults() *SetAlertsStatusByQuery`

NewSetAlertsStatusByQueryWithDefaults instantiates a new SetAlertsStatusByQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConflicts

`func (o *SetAlertsStatusByQuery) GetConflicts() string`

GetConflicts returns the Conflicts field if non-nil, zero value otherwise.

### GetConflictsOk

`func (o *SetAlertsStatusByQuery) GetConflictsOk() (*string, bool)`

GetConflictsOk returns a tuple with the Conflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflicts

`func (o *SetAlertsStatusByQuery) SetConflicts(v string)`

SetConflicts sets Conflicts field to given value.

### HasConflicts

`func (o *SetAlertsStatusByQuery) HasConflicts() bool`

HasConflicts returns a boolean if a field has been set.

### GetQuery

`func (o *SetAlertsStatusByQuery) GetQuery() map[string]interface{}`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SetAlertsStatusByQuery) GetQueryOk() (*map[string]interface{}, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SetAlertsStatusByQuery) SetQuery(v map[string]interface{})`

SetQuery sets Query field to given value.


### GetStatus

`func (o *SetAlertsStatusByQuery) GetStatus() AlertStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SetAlertsStatusByQuery) GetStatusOk() (*AlertStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SetAlertsStatusByQuery) SetStatus(v AlertStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


