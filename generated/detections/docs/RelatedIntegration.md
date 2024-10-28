# RelatedIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Integration** | Pointer to **string** | A string that is not empty and does not contain only whitespace | [optional] 
**Package** | **string** | A string that is not empty and does not contain only whitespace | 
**Version** | **string** | A string that is not empty and does not contain only whitespace | 

## Methods

### NewRelatedIntegration

`func NewRelatedIntegration(package_ string, version string, ) *RelatedIntegration`

NewRelatedIntegration instantiates a new RelatedIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelatedIntegrationWithDefaults

`func NewRelatedIntegrationWithDefaults() *RelatedIntegration`

NewRelatedIntegrationWithDefaults instantiates a new RelatedIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegration

`func (o *RelatedIntegration) GetIntegration() string`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *RelatedIntegration) GetIntegrationOk() (*string, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *RelatedIntegration) SetIntegration(v string)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *RelatedIntegration) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetPackage

`func (o *RelatedIntegration) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *RelatedIntegration) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *RelatedIntegration) SetPackage(v string)`

SetPackage sets Package field to given value.


### GetVersion

`func (o *RelatedIntegration) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RelatedIntegration) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RelatedIntegration) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


