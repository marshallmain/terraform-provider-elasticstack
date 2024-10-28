# BulkActionEditPayloadScheduleValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interval** | **string** | Interval in which the rule runs. For example, &#x60;\&quot;1h\&quot;&#x60; means the rule runs every hour. | 
**Lookback** | **string** | Lookback time for the rule | 

## Methods

### NewBulkActionEditPayloadScheduleValue

`func NewBulkActionEditPayloadScheduleValue(interval string, lookback string, ) *BulkActionEditPayloadScheduleValue`

NewBulkActionEditPayloadScheduleValue instantiates a new BulkActionEditPayloadScheduleValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkActionEditPayloadScheduleValueWithDefaults

`func NewBulkActionEditPayloadScheduleValueWithDefaults() *BulkActionEditPayloadScheduleValue`

NewBulkActionEditPayloadScheduleValueWithDefaults instantiates a new BulkActionEditPayloadScheduleValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterval

`func (o *BulkActionEditPayloadScheduleValue) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *BulkActionEditPayloadScheduleValue) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *BulkActionEditPayloadScheduleValue) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetLookback

`func (o *BulkActionEditPayloadScheduleValue) GetLookback() string`

GetLookback returns the Lookback field if non-nil, zero value otherwise.

### GetLookbackOk

`func (o *BulkActionEditPayloadScheduleValue) GetLookbackOk() (*string, bool)`

GetLookbackOk returns a tuple with the Lookback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookback

`func (o *BulkActionEditPayloadScheduleValue) SetLookback(v string)`

SetLookback sets Lookback field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


