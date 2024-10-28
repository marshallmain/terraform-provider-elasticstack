# RuleAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionTypeId** | **string** | The action type used for sending notifications. | 
**AlertsFilter** | Pointer to **map[string]interface{}** |  | [optional] 
**Frequency** | Pointer to [**RuleActionFrequency**](RuleActionFrequency.md) |  | [optional] 
**Group** | Pointer to **string** | Optionally groups actions by use cases. Use &#x60;default&#x60; for alert notifications. | [optional] 
**Id** | **string** | The connector ID. | 
**Params** | **map[string]interface{}** | Object containing the allowed connector fields, which varies according to the connector type. | 
**Uuid** | Pointer to **string** | A string that is not empty and does not contain only whitespace | [optional] 

## Methods

### NewRuleAction

`func NewRuleAction(actionTypeId string, id string, params map[string]interface{}, ) *RuleAction`

NewRuleAction instantiates a new RuleAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleActionWithDefaults

`func NewRuleActionWithDefaults() *RuleAction`

NewRuleActionWithDefaults instantiates a new RuleAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionTypeId

`func (o *RuleAction) GetActionTypeId() string`

GetActionTypeId returns the ActionTypeId field if non-nil, zero value otherwise.

### GetActionTypeIdOk

`func (o *RuleAction) GetActionTypeIdOk() (*string, bool)`

GetActionTypeIdOk returns a tuple with the ActionTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTypeId

`func (o *RuleAction) SetActionTypeId(v string)`

SetActionTypeId sets ActionTypeId field to given value.


### GetAlertsFilter

`func (o *RuleAction) GetAlertsFilter() map[string]interface{}`

GetAlertsFilter returns the AlertsFilter field if non-nil, zero value otherwise.

### GetAlertsFilterOk

`func (o *RuleAction) GetAlertsFilterOk() (*map[string]interface{}, bool)`

GetAlertsFilterOk returns a tuple with the AlertsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertsFilter

`func (o *RuleAction) SetAlertsFilter(v map[string]interface{})`

SetAlertsFilter sets AlertsFilter field to given value.

### HasAlertsFilter

`func (o *RuleAction) HasAlertsFilter() bool`

HasAlertsFilter returns a boolean if a field has been set.

### GetFrequency

`func (o *RuleAction) GetFrequency() RuleActionFrequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *RuleAction) GetFrequencyOk() (*RuleActionFrequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *RuleAction) SetFrequency(v RuleActionFrequency)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *RuleAction) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetGroup

`func (o *RuleAction) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *RuleAction) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *RuleAction) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *RuleAction) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetId

`func (o *RuleAction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuleAction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuleAction) SetId(v string)`

SetId sets Id field to given value.


### GetParams

`func (o *RuleAction) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *RuleAction) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *RuleAction) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.


### GetUuid

`func (o *RuleAction) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RuleAction) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RuleAction) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RuleAction) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


