# SetAlertsStatusByIds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignalIds** | **[]string** |  | 
**Status** | [**AlertStatus**](AlertStatus.md) |  | 

## Methods

### NewSetAlertsStatusByIds

`func NewSetAlertsStatusByIds(signalIds []string, status AlertStatus, ) *SetAlertsStatusByIds`

NewSetAlertsStatusByIds instantiates a new SetAlertsStatusByIds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetAlertsStatusByIdsWithDefaults

`func NewSetAlertsStatusByIdsWithDefaults() *SetAlertsStatusByIds`

NewSetAlertsStatusByIdsWithDefaults instantiates a new SetAlertsStatusByIds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignalIds

`func (o *SetAlertsStatusByIds) GetSignalIds() []string`

GetSignalIds returns the SignalIds field if non-nil, zero value otherwise.

### GetSignalIdsOk

`func (o *SetAlertsStatusByIds) GetSignalIdsOk() (*[]string, bool)`

GetSignalIdsOk returns a tuple with the SignalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalIds

`func (o *SetAlertsStatusByIds) SetSignalIds(v []string)`

SetSignalIds sets SignalIds field to given value.


### GetStatus

`func (o *SetAlertsStatusByIds) GetStatus() AlertStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SetAlertsStatusByIds) GetStatusOk() (*AlertStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SetAlertsStatusByIds) SetStatus(v AlertStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


