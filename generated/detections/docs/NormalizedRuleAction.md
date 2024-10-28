# NormalizedRuleAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertsFilter** | Pointer to **map[string]interface{}** |  | [optional] 
**Frequency** | Pointer to [**RuleActionFrequency**](RuleActionFrequency.md) |  | [optional] 
**Group** | Pointer to **string** | Optionally groups actions by use cases. Use &#x60;default&#x60; for alert notifications. | [optional] 
**Id** | **string** | The connector ID. | 
**Params** | **map[string]interface{}** | Object containing the allowed connector fields, which varies according to the connector type. | 

## Methods

### NewNormalizedRuleAction

`func NewNormalizedRuleAction(id string, params map[string]interface{}, ) *NormalizedRuleAction`

NewNormalizedRuleAction instantiates a new NormalizedRuleAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNormalizedRuleActionWithDefaults

`func NewNormalizedRuleActionWithDefaults() *NormalizedRuleAction`

NewNormalizedRuleActionWithDefaults instantiates a new NormalizedRuleAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertsFilter

`func (o *NormalizedRuleAction) GetAlertsFilter() map[string]interface{}`

GetAlertsFilter returns the AlertsFilter field if non-nil, zero value otherwise.

### GetAlertsFilterOk

`func (o *NormalizedRuleAction) GetAlertsFilterOk() (*map[string]interface{}, bool)`

GetAlertsFilterOk returns a tuple with the AlertsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertsFilter

`func (o *NormalizedRuleAction) SetAlertsFilter(v map[string]interface{})`

SetAlertsFilter sets AlertsFilter field to given value.

### HasAlertsFilter

`func (o *NormalizedRuleAction) HasAlertsFilter() bool`

HasAlertsFilter returns a boolean if a field has been set.

### GetFrequency

`func (o *NormalizedRuleAction) GetFrequency() RuleActionFrequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *NormalizedRuleAction) GetFrequencyOk() (*RuleActionFrequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *NormalizedRuleAction) SetFrequency(v RuleActionFrequency)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *NormalizedRuleAction) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetGroup

`func (o *NormalizedRuleAction) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *NormalizedRuleAction) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *NormalizedRuleAction) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *NormalizedRuleAction) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetId

`func (o *NormalizedRuleAction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NormalizedRuleAction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NormalizedRuleAction) SetId(v string)`

SetId sets Id field to given value.


### GetParams

`func (o *NormalizedRuleAction) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *NormalizedRuleAction) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *NormalizedRuleAction) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


