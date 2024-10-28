# ThreatTechnique

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Technique ID | 
**Name** | **string** | Technique name | 
**Reference** | **string** | Technique reference | 
**Subtechnique** | Pointer to [**[]ThreatSubtechnique**](ThreatSubtechnique.md) | Array containing more specific information on the attack technique | [optional] 

## Methods

### NewThreatTechnique

`func NewThreatTechnique(id string, name string, reference string, ) *ThreatTechnique`

NewThreatTechnique instantiates a new ThreatTechnique object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreatTechniqueWithDefaults

`func NewThreatTechniqueWithDefaults() *ThreatTechnique`

NewThreatTechniqueWithDefaults instantiates a new ThreatTechnique object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ThreatTechnique) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThreatTechnique) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThreatTechnique) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ThreatTechnique) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ThreatTechnique) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ThreatTechnique) SetName(v string)`

SetName sets Name field to given value.


### GetReference

`func (o *ThreatTechnique) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ThreatTechnique) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ThreatTechnique) SetReference(v string)`

SetReference sets Reference field to given value.


### GetSubtechnique

`func (o *ThreatTechnique) GetSubtechnique() []ThreatSubtechnique`

GetSubtechnique returns the Subtechnique field if non-nil, zero value otherwise.

### GetSubtechniqueOk

`func (o *ThreatTechnique) GetSubtechniqueOk() (*[]ThreatSubtechnique, bool)`

GetSubtechniqueOk returns a tuple with the Subtechnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtechnique

`func (o *ThreatTechnique) SetSubtechnique(v []ThreatSubtechnique)`

SetSubtechnique sets Subtechnique field to given value.

### HasSubtechnique

`func (o *ThreatTechnique) HasSubtechnique() bool`

HasSubtechnique returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


