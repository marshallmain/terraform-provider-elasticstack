# RulePreviewLoggedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A string that is not empty and does not contain only whitespace | [optional] 
**Duration** | Pointer to **int32** |  | [optional] 
**Request** | **string** | A string that is not empty and does not contain only whitespace | 

## Methods

### NewRulePreviewLoggedRequest

`func NewRulePreviewLoggedRequest(request string, ) *RulePreviewLoggedRequest`

NewRulePreviewLoggedRequest instantiates a new RulePreviewLoggedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRulePreviewLoggedRequestWithDefaults

`func NewRulePreviewLoggedRequestWithDefaults() *RulePreviewLoggedRequest`

NewRulePreviewLoggedRequestWithDefaults instantiates a new RulePreviewLoggedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *RulePreviewLoggedRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RulePreviewLoggedRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RulePreviewLoggedRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RulePreviewLoggedRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDuration

`func (o *RulePreviewLoggedRequest) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *RulePreviewLoggedRequest) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *RulePreviewLoggedRequest) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *RulePreviewLoggedRequest) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetRequest

`func (o *RulePreviewLoggedRequest) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *RulePreviewLoggedRequest) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *RulePreviewLoggedRequest) SetRequest(v string)`

SetRequest sets Request field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


