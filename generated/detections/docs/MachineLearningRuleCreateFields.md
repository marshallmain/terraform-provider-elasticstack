# MachineLearningRuleCreateFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnomalyThreshold** | **int32** | Anomaly threshold | 
**MachineLearningJobId** | [**MachineLearningJobId**](MachineLearningJobId.md) |  | 
**Type** | **string** | Rule type | 
**AlertSuppression** | Pointer to [**AlertSuppression**](AlertSuppression.md) |  | [optional] 

## Methods

### NewMachineLearningRuleCreateFields

`func NewMachineLearningRuleCreateFields(anomalyThreshold int32, machineLearningJobId MachineLearningJobId, type_ string, ) *MachineLearningRuleCreateFields`

NewMachineLearningRuleCreateFields instantiates a new MachineLearningRuleCreateFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineLearningRuleCreateFieldsWithDefaults

`func NewMachineLearningRuleCreateFieldsWithDefaults() *MachineLearningRuleCreateFields`

NewMachineLearningRuleCreateFieldsWithDefaults instantiates a new MachineLearningRuleCreateFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnomalyThreshold

`func (o *MachineLearningRuleCreateFields) GetAnomalyThreshold() int32`

GetAnomalyThreshold returns the AnomalyThreshold field if non-nil, zero value otherwise.

### GetAnomalyThresholdOk

`func (o *MachineLearningRuleCreateFields) GetAnomalyThresholdOk() (*int32, bool)`

GetAnomalyThresholdOk returns a tuple with the AnomalyThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomalyThreshold

`func (o *MachineLearningRuleCreateFields) SetAnomalyThreshold(v int32)`

SetAnomalyThreshold sets AnomalyThreshold field to given value.


### GetMachineLearningJobId

`func (o *MachineLearningRuleCreateFields) GetMachineLearningJobId() MachineLearningJobId`

GetMachineLearningJobId returns the MachineLearningJobId field if non-nil, zero value otherwise.

### GetMachineLearningJobIdOk

`func (o *MachineLearningRuleCreateFields) GetMachineLearningJobIdOk() (*MachineLearningJobId, bool)`

GetMachineLearningJobIdOk returns a tuple with the MachineLearningJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineLearningJobId

`func (o *MachineLearningRuleCreateFields) SetMachineLearningJobId(v MachineLearningJobId)`

SetMachineLearningJobId sets MachineLearningJobId field to given value.


### GetType

`func (o *MachineLearningRuleCreateFields) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MachineLearningRuleCreateFields) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MachineLearningRuleCreateFields) SetType(v string)`

SetType sets Type field to given value.


### GetAlertSuppression

`func (o *MachineLearningRuleCreateFields) GetAlertSuppression() AlertSuppression`

GetAlertSuppression returns the AlertSuppression field if non-nil, zero value otherwise.

### GetAlertSuppressionOk

`func (o *MachineLearningRuleCreateFields) GetAlertSuppressionOk() (*AlertSuppression, bool)`

GetAlertSuppressionOk returns a tuple with the AlertSuppression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertSuppression

`func (o *MachineLearningRuleCreateFields) SetAlertSuppression(v AlertSuppression)`

SetAlertSuppression sets AlertSuppression field to given value.

### HasAlertSuppression

`func (o *MachineLearningRuleCreateFields) HasAlertSuppression() bool`

HasAlertSuppression returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


