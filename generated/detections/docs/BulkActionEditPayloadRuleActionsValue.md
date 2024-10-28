# BulkActionEditPayloadRuleActionsValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | [**[]NormalizedRuleAction**](NormalizedRuleAction.md) |  | 
**Throttle** | Pointer to [**ThrottleForBulkActions**](ThrottleForBulkActions.md) |  | [optional] 

## Methods

### NewBulkActionEditPayloadRuleActionsValue

`func NewBulkActionEditPayloadRuleActionsValue(actions []NormalizedRuleAction, ) *BulkActionEditPayloadRuleActionsValue`

NewBulkActionEditPayloadRuleActionsValue instantiates a new BulkActionEditPayloadRuleActionsValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkActionEditPayloadRuleActionsValueWithDefaults

`func NewBulkActionEditPayloadRuleActionsValueWithDefaults() *BulkActionEditPayloadRuleActionsValue`

NewBulkActionEditPayloadRuleActionsValueWithDefaults instantiates a new BulkActionEditPayloadRuleActionsValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *BulkActionEditPayloadRuleActionsValue) GetActions() []NormalizedRuleAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *BulkActionEditPayloadRuleActionsValue) GetActionsOk() (*[]NormalizedRuleAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *BulkActionEditPayloadRuleActionsValue) SetActions(v []NormalizedRuleAction)`

SetActions sets Actions field to given value.


### GetThrottle

`func (o *BulkActionEditPayloadRuleActionsValue) GetThrottle() ThrottleForBulkActions`

GetThrottle returns the Throttle field if non-nil, zero value otherwise.

### GetThrottleOk

`func (o *BulkActionEditPayloadRuleActionsValue) GetThrottleOk() (*ThrottleForBulkActions, bool)`

GetThrottleOk returns a tuple with the Throttle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottle

`func (o *BulkActionEditPayloadRuleActionsValue) SetThrottle(v ThrottleForBulkActions)`

SetThrottle sets Throttle field to given value.

### HasThrottle

`func (o *BulkActionEditPayloadRuleActionsValue) HasThrottle() bool`

HasThrottle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


