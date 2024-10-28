# AlertSuppression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to [**AlertSuppressionDuration**](AlertSuppressionDuration.md) |  | [optional] 
**GroupBy** | **[]string** |  | 
**MissingFieldsStrategy** | Pointer to [**AlertSuppressionMissingFieldsStrategy**](AlertSuppressionMissingFieldsStrategy.md) |  | [optional] 

## Methods

### NewAlertSuppression

`func NewAlertSuppression(groupBy []string, ) *AlertSuppression`

NewAlertSuppression instantiates a new AlertSuppression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertSuppressionWithDefaults

`func NewAlertSuppressionWithDefaults() *AlertSuppression`

NewAlertSuppressionWithDefaults instantiates a new AlertSuppression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *AlertSuppression) GetDuration() AlertSuppressionDuration`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AlertSuppression) GetDurationOk() (*AlertSuppressionDuration, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AlertSuppression) SetDuration(v AlertSuppressionDuration)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *AlertSuppression) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetGroupBy

`func (o *AlertSuppression) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *AlertSuppression) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *AlertSuppression) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.


### GetMissingFieldsStrategy

`func (o *AlertSuppression) GetMissingFieldsStrategy() AlertSuppressionMissingFieldsStrategy`

GetMissingFieldsStrategy returns the MissingFieldsStrategy field if non-nil, zero value otherwise.

### GetMissingFieldsStrategyOk

`func (o *AlertSuppression) GetMissingFieldsStrategyOk() (*AlertSuppressionMissingFieldsStrategy, bool)`

GetMissingFieldsStrategyOk returns a tuple with the MissingFieldsStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingFieldsStrategy

`func (o *AlertSuppression) SetMissingFieldsStrategy(v AlertSuppressionMissingFieldsStrategy)`

SetMissingFieldsStrategy sets MissingFieldsStrategy field to given value.

### HasMissingFieldsStrategy

`func (o *AlertSuppression) HasMissingFieldsStrategy() bool`

HasMissingFieldsStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


