# Threat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Framework** | **string** | Relevant attack framework | 
**Tactic** | [**ThreatTactic**](ThreatTactic.md) |  | 
**Technique** | Pointer to [**[]ThreatTechnique**](ThreatTechnique.md) | Array containing information on the attack techniques (optional) | [optional] 

## Methods

### NewThreat

`func NewThreat(framework string, tactic ThreatTactic, ) *Threat`

NewThreat instantiates a new Threat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreatWithDefaults

`func NewThreatWithDefaults() *Threat`

NewThreatWithDefaults instantiates a new Threat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFramework

`func (o *Threat) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *Threat) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *Threat) SetFramework(v string)`

SetFramework sets Framework field to given value.


### GetTactic

`func (o *Threat) GetTactic() ThreatTactic`

GetTactic returns the Tactic field if non-nil, zero value otherwise.

### GetTacticOk

`func (o *Threat) GetTacticOk() (*ThreatTactic, bool)`

GetTacticOk returns a tuple with the Tactic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTactic

`func (o *Threat) SetTactic(v ThreatTactic)`

SetTactic sets Tactic field to given value.


### GetTechnique

`func (o *Threat) GetTechnique() []ThreatTechnique`

GetTechnique returns the Technique field if non-nil, zero value otherwise.

### GetTechniqueOk

`func (o *Threat) GetTechniqueOk() (*[]ThreatTechnique, bool)`

GetTechniqueOk returns a tuple with the Technique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnique

`func (o *Threat) SetTechnique(v []ThreatTechnique)`

SetTechnique sets Technique field to given value.

### HasTechnique

`func (o *Threat) HasTechnique() bool`

HasTechnique returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


