# ResponseAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionTypeId** | **string** |  | 
**Params** | [**EndpointResponseActionParams**](EndpointResponseActionParams.md) |  | 

## Methods

### NewResponseAction

`func NewResponseAction(actionTypeId string, params EndpointResponseActionParams, ) *ResponseAction`

NewResponseAction instantiates a new ResponseAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseActionWithDefaults

`func NewResponseActionWithDefaults() *ResponseAction`

NewResponseActionWithDefaults instantiates a new ResponseAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionTypeId

`func (o *ResponseAction) GetActionTypeId() string`

GetActionTypeId returns the ActionTypeId field if non-nil, zero value otherwise.

### GetActionTypeIdOk

`func (o *ResponseAction) GetActionTypeIdOk() (*string, bool)`

GetActionTypeIdOk returns a tuple with the ActionTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTypeId

`func (o *ResponseAction) SetActionTypeId(v string)`

SetActionTypeId sets ActionTypeId field to given value.


### GetParams

`func (o *ResponseAction) GetParams() EndpointResponseActionParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ResponseAction) GetParamsOk() (*EndpointResponseActionParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ResponseAction) SetParams(v EndpointResponseActionParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


