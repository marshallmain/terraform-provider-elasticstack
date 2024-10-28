# RuleUpdateProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]RuleAction**](RuleAction.md) |  | [optional] 
**AliasPurpose** | Pointer to [**SavedObjectResolveAliasPurpose**](SavedObjectResolveAliasPurpose.md) |  | [optional] 
**AliasTargetId** | Pointer to **string** |  | [optional] 
**Author** | Pointer to **[]string** |  | [optional] 
**BuildingBlockType** | Pointer to **string** | Determines if the rule acts as a building block. By default, building-block alerts are not displayed in the UI. These rules are used as a foundation for other rules that do generate alerts. Its value must be default. | [optional] 
**Description** | **string** |  | 
**Enabled** | Pointer to **bool** | Determines whether the rule is enabled. | [optional] 
**ExceptionsList** | Pointer to [**[]RuleExceptionList**](RuleExceptionList.md) |  | [optional] 
**FalsePositives** | Pointer to **[]string** |  | [optional] 
**From** | Pointer to **string** | Time from which data is analyzed each time the rule runs, using a date math range. For example, now-4200s means the rule analyzes data from 70 minutes before its start time. Defaults to now-6m (analyzes data from 6 minutes before the start time). | [optional] 
**Id** | Pointer to **string** | A universally unique identifier | [optional] 
**Interval** | Pointer to **string** | Frequency of rule execution, using a date math range. For example, \&quot;1h\&quot; means the rule runs every hour. Defaults to 5m (5 minutes). | [optional] 
**InvestigationFields** | Pointer to [**InvestigationFields**](InvestigationFields.md) |  | [optional] 
**License** | Pointer to **string** | The rule&#39;s license. | [optional] 
**MaxSignals** | Pointer to **int32** |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | **string** |  | 
**Namespace** | Pointer to **string** | Has no effect. | [optional] 
**Note** | Pointer to **string** | Notes to help investigate alerts produced by the rule. | [optional] 
**Outcome** | Pointer to [**SavedObjectResolveOutcome**](SavedObjectResolveOutcome.md) |  | [optional] 
**OutputIndex** | Pointer to **string** | (deprecated) Has no effect. | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**RelatedIntegrations** | Pointer to [**[]RelatedIntegration**](RelatedIntegration.md) |  | [optional] 
**RequiredFields** | Pointer to [**[]RequiredFieldInput**](RequiredFieldInput.md) |  | [optional] 
**RiskScore** | **int32** | Risk score (0 to 100) | 
**RiskScoreMapping** | Pointer to [**[]RiskScoreMappingInner**](RiskScoreMappingInner.md) | Overrides generated alerts&#39; risk_score with a value from the source event | [optional] 
**RuleId** | Pointer to **string** | Could be any string, not necessarily a UUID | [optional] 
**RuleNameOverride** | Pointer to **string** | Sets the source field for the alert&#39;s signal.rule.name value | [optional] 
**Setup** | Pointer to **string** |  | [optional] 
**Severity** | [**Severity**](Severity.md) |  | 
**SeverityMapping** | Pointer to [**[]SeverityMappingInner**](SeverityMappingInner.md) | Overrides generated alerts&#39; severity with values from the source event | [optional] 
**Tags** | Pointer to **[]string** | String array containing words and phrases to help categorize, filter, and search rules. Defaults to an empty array. | [optional] 
**Threat** | Pointer to [**[]Threat**](Threat.md) |  | [optional] 
**Throttle** | Pointer to [**RuleActionThrottle**](RuleActionThrottle.md) |  | [optional] 
**TimelineId** | Pointer to **string** | Timeline template ID | [optional] 
**TimelineTitle** | Pointer to **string** | Timeline template title | [optional] 
**TimestampOverride** | Pointer to **string** | Sets the time field used to query indices | [optional] 
**TimestampOverrideFallbackDisabled** | Pointer to **bool** | Disables the fallback to the event&#39;s @timestamp field | [optional] 
**To** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int32** | The rule&#39;s version number. | [optional] 
**Language** | [**EsqlQueryLanguage**](EsqlQueryLanguage.md) |  | 
**Query** | **string** |  | 
**Type** | **string** | Rule type | 
**AlertSuppression** | Pointer to [**AlertSuppression**](AlertSuppression.md) |  | [optional] 
**DataViewId** | Pointer to **string** |  | [optional] 
**EventCategoryOverride** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to **[]interface{}** |  | [optional] 
**Index** | Pointer to **[]string** |  | [optional] 
**ResponseActions** | Pointer to [**[]ResponseAction**](ResponseAction.md) |  | [optional] 
**TiebreakerField** | Pointer to **string** | Sets a secondary field for sorting events | [optional] 
**TimestampField** | Pointer to **string** | Contains the event timestamp used for sorting a sequence of events | [optional] 
**SavedId** | **string** |  | 
**Threshold** | [**Threshold**](Threshold.md) |  | 
**ThreatIndex** | **[]string** |  | 
**ThreatMapping** | [**[]ThreatMappingInner**](ThreatMappingInner.md) |  | 
**ThreatQuery** | **string** | Query to run | 
**ConcurrentSearches** | Pointer to **int32** |  | [optional] 
**ItemsPerSearch** | Pointer to **int32** |  | [optional] 
**ThreatFilters** | Pointer to **[]interface{}** |  | [optional] 
**ThreatIndicatorPath** | Pointer to **string** | Defines the path to the threat indicator in the indicator documents (optional) | [optional] 
**ThreatLanguage** | Pointer to [**KqlQueryLanguage**](KqlQueryLanguage.md) |  | [optional] 
**AnomalyThreshold** | **int32** | Anomaly threshold | 
**MachineLearningJobId** | [**MachineLearningJobId**](MachineLearningJobId.md) |  | 
**HistoryWindowStart** | **string** | A string that is not empty and does not contain only whitespace | 
**NewTermsFields** | **[]string** |  | 

## Methods

### NewRuleUpdateProps

`func NewRuleUpdateProps(description string, name string, riskScore int32, severity Severity, language EsqlQueryLanguage, query string, type_ string, savedId string, threshold Threshold, threatIndex []string, threatMapping []ThreatMappingInner, threatQuery string, anomalyThreshold int32, machineLearningJobId MachineLearningJobId, historyWindowStart string, newTermsFields []string, ) *RuleUpdateProps`

NewRuleUpdateProps instantiates a new RuleUpdateProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleUpdatePropsWithDefaults

`func NewRuleUpdatePropsWithDefaults() *RuleUpdateProps`

NewRuleUpdatePropsWithDefaults instantiates a new RuleUpdateProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *RuleUpdateProps) GetActions() []RuleAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *RuleUpdateProps) GetActionsOk() (*[]RuleAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *RuleUpdateProps) SetActions(v []RuleAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *RuleUpdateProps) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetAliasPurpose

`func (o *RuleUpdateProps) GetAliasPurpose() SavedObjectResolveAliasPurpose`

GetAliasPurpose returns the AliasPurpose field if non-nil, zero value otherwise.

### GetAliasPurposeOk

`func (o *RuleUpdateProps) GetAliasPurposeOk() (*SavedObjectResolveAliasPurpose, bool)`

GetAliasPurposeOk returns a tuple with the AliasPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasPurpose

`func (o *RuleUpdateProps) SetAliasPurpose(v SavedObjectResolveAliasPurpose)`

SetAliasPurpose sets AliasPurpose field to given value.

### HasAliasPurpose

`func (o *RuleUpdateProps) HasAliasPurpose() bool`

HasAliasPurpose returns a boolean if a field has been set.

### GetAliasTargetId

`func (o *RuleUpdateProps) GetAliasTargetId() string`

GetAliasTargetId returns the AliasTargetId field if non-nil, zero value otherwise.

### GetAliasTargetIdOk

`func (o *RuleUpdateProps) GetAliasTargetIdOk() (*string, bool)`

GetAliasTargetIdOk returns a tuple with the AliasTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasTargetId

`func (o *RuleUpdateProps) SetAliasTargetId(v string)`

SetAliasTargetId sets AliasTargetId field to given value.

### HasAliasTargetId

`func (o *RuleUpdateProps) HasAliasTargetId() bool`

HasAliasTargetId returns a boolean if a field has been set.

### GetAuthor

`func (o *RuleUpdateProps) GetAuthor() []string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *RuleUpdateProps) GetAuthorOk() (*[]string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *RuleUpdateProps) SetAuthor(v []string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *RuleUpdateProps) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetBuildingBlockType

`func (o *RuleUpdateProps) GetBuildingBlockType() string`

GetBuildingBlockType returns the BuildingBlockType field if non-nil, zero value otherwise.

### GetBuildingBlockTypeOk

`func (o *RuleUpdateProps) GetBuildingBlockTypeOk() (*string, bool)`

GetBuildingBlockTypeOk returns a tuple with the BuildingBlockType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildingBlockType

`func (o *RuleUpdateProps) SetBuildingBlockType(v string)`

SetBuildingBlockType sets BuildingBlockType field to given value.

### HasBuildingBlockType

`func (o *RuleUpdateProps) HasBuildingBlockType() bool`

HasBuildingBlockType returns a boolean if a field has been set.

### GetDescription

`func (o *RuleUpdateProps) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RuleUpdateProps) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RuleUpdateProps) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEnabled

`func (o *RuleUpdateProps) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RuleUpdateProps) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RuleUpdateProps) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RuleUpdateProps) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExceptionsList

`func (o *RuleUpdateProps) GetExceptionsList() []RuleExceptionList`

GetExceptionsList returns the ExceptionsList field if non-nil, zero value otherwise.

### GetExceptionsListOk

`func (o *RuleUpdateProps) GetExceptionsListOk() (*[]RuleExceptionList, bool)`

GetExceptionsListOk returns a tuple with the ExceptionsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionsList

`func (o *RuleUpdateProps) SetExceptionsList(v []RuleExceptionList)`

SetExceptionsList sets ExceptionsList field to given value.

### HasExceptionsList

`func (o *RuleUpdateProps) HasExceptionsList() bool`

HasExceptionsList returns a boolean if a field has been set.

### GetFalsePositives

`func (o *RuleUpdateProps) GetFalsePositives() []string`

GetFalsePositives returns the FalsePositives field if non-nil, zero value otherwise.

### GetFalsePositivesOk

`func (o *RuleUpdateProps) GetFalsePositivesOk() (*[]string, bool)`

GetFalsePositivesOk returns a tuple with the FalsePositives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFalsePositives

`func (o *RuleUpdateProps) SetFalsePositives(v []string)`

SetFalsePositives sets FalsePositives field to given value.

### HasFalsePositives

`func (o *RuleUpdateProps) HasFalsePositives() bool`

HasFalsePositives returns a boolean if a field has been set.

### GetFrom

`func (o *RuleUpdateProps) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *RuleUpdateProps) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *RuleUpdateProps) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *RuleUpdateProps) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetId

`func (o *RuleUpdateProps) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuleUpdateProps) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuleUpdateProps) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RuleUpdateProps) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterval

`func (o *RuleUpdateProps) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *RuleUpdateProps) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *RuleUpdateProps) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *RuleUpdateProps) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetInvestigationFields

`func (o *RuleUpdateProps) GetInvestigationFields() InvestigationFields`

GetInvestigationFields returns the InvestigationFields field if non-nil, zero value otherwise.

### GetInvestigationFieldsOk

`func (o *RuleUpdateProps) GetInvestigationFieldsOk() (*InvestigationFields, bool)`

GetInvestigationFieldsOk returns a tuple with the InvestigationFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigationFields

`func (o *RuleUpdateProps) SetInvestigationFields(v InvestigationFields)`

SetInvestigationFields sets InvestigationFields field to given value.

### HasInvestigationFields

`func (o *RuleUpdateProps) HasInvestigationFields() bool`

HasInvestigationFields returns a boolean if a field has been set.

### GetLicense

`func (o *RuleUpdateProps) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *RuleUpdateProps) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *RuleUpdateProps) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *RuleUpdateProps) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetMaxSignals

`func (o *RuleUpdateProps) GetMaxSignals() int32`

GetMaxSignals returns the MaxSignals field if non-nil, zero value otherwise.

### GetMaxSignalsOk

`func (o *RuleUpdateProps) GetMaxSignalsOk() (*int32, bool)`

GetMaxSignalsOk returns a tuple with the MaxSignals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSignals

`func (o *RuleUpdateProps) SetMaxSignals(v int32)`

SetMaxSignals sets MaxSignals field to given value.

### HasMaxSignals

`func (o *RuleUpdateProps) HasMaxSignals() bool`

HasMaxSignals returns a boolean if a field has been set.

### GetMeta

`func (o *RuleUpdateProps) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RuleUpdateProps) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RuleUpdateProps) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *RuleUpdateProps) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetName

`func (o *RuleUpdateProps) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleUpdateProps) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleUpdateProps) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *RuleUpdateProps) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *RuleUpdateProps) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *RuleUpdateProps) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *RuleUpdateProps) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNote

`func (o *RuleUpdateProps) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *RuleUpdateProps) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *RuleUpdateProps) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *RuleUpdateProps) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetOutcome

`func (o *RuleUpdateProps) GetOutcome() SavedObjectResolveOutcome`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *RuleUpdateProps) GetOutcomeOk() (*SavedObjectResolveOutcome, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *RuleUpdateProps) SetOutcome(v SavedObjectResolveOutcome)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *RuleUpdateProps) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetOutputIndex

`func (o *RuleUpdateProps) GetOutputIndex() string`

GetOutputIndex returns the OutputIndex field if non-nil, zero value otherwise.

### GetOutputIndexOk

`func (o *RuleUpdateProps) GetOutputIndexOk() (*string, bool)`

GetOutputIndexOk returns a tuple with the OutputIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputIndex

`func (o *RuleUpdateProps) SetOutputIndex(v string)`

SetOutputIndex sets OutputIndex field to given value.

### HasOutputIndex

`func (o *RuleUpdateProps) HasOutputIndex() bool`

HasOutputIndex returns a boolean if a field has been set.

### GetReferences

`func (o *RuleUpdateProps) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *RuleUpdateProps) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *RuleUpdateProps) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *RuleUpdateProps) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetRelatedIntegrations

`func (o *RuleUpdateProps) GetRelatedIntegrations() []RelatedIntegration`

GetRelatedIntegrations returns the RelatedIntegrations field if non-nil, zero value otherwise.

### GetRelatedIntegrationsOk

`func (o *RuleUpdateProps) GetRelatedIntegrationsOk() (*[]RelatedIntegration, bool)`

GetRelatedIntegrationsOk returns a tuple with the RelatedIntegrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedIntegrations

`func (o *RuleUpdateProps) SetRelatedIntegrations(v []RelatedIntegration)`

SetRelatedIntegrations sets RelatedIntegrations field to given value.

### HasRelatedIntegrations

`func (o *RuleUpdateProps) HasRelatedIntegrations() bool`

HasRelatedIntegrations returns a boolean if a field has been set.

### GetRequiredFields

`func (o *RuleUpdateProps) GetRequiredFields() []RequiredFieldInput`

GetRequiredFields returns the RequiredFields field if non-nil, zero value otherwise.

### GetRequiredFieldsOk

`func (o *RuleUpdateProps) GetRequiredFieldsOk() (*[]RequiredFieldInput, bool)`

GetRequiredFieldsOk returns a tuple with the RequiredFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredFields

`func (o *RuleUpdateProps) SetRequiredFields(v []RequiredFieldInput)`

SetRequiredFields sets RequiredFields field to given value.

### HasRequiredFields

`func (o *RuleUpdateProps) HasRequiredFields() bool`

HasRequiredFields returns a boolean if a field has been set.

### GetRiskScore

`func (o *RuleUpdateProps) GetRiskScore() int32`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *RuleUpdateProps) GetRiskScoreOk() (*int32, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *RuleUpdateProps) SetRiskScore(v int32)`

SetRiskScore sets RiskScore field to given value.


### GetRiskScoreMapping

`func (o *RuleUpdateProps) GetRiskScoreMapping() []RiskScoreMappingInner`

GetRiskScoreMapping returns the RiskScoreMapping field if non-nil, zero value otherwise.

### GetRiskScoreMappingOk

`func (o *RuleUpdateProps) GetRiskScoreMappingOk() (*[]RiskScoreMappingInner, bool)`

GetRiskScoreMappingOk returns a tuple with the RiskScoreMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScoreMapping

`func (o *RuleUpdateProps) SetRiskScoreMapping(v []RiskScoreMappingInner)`

SetRiskScoreMapping sets RiskScoreMapping field to given value.

### HasRiskScoreMapping

`func (o *RuleUpdateProps) HasRiskScoreMapping() bool`

HasRiskScoreMapping returns a boolean if a field has been set.

### GetRuleId

`func (o *RuleUpdateProps) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *RuleUpdateProps) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *RuleUpdateProps) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *RuleUpdateProps) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetRuleNameOverride

`func (o *RuleUpdateProps) GetRuleNameOverride() string`

GetRuleNameOverride returns the RuleNameOverride field if non-nil, zero value otherwise.

### GetRuleNameOverrideOk

`func (o *RuleUpdateProps) GetRuleNameOverrideOk() (*string, bool)`

GetRuleNameOverrideOk returns a tuple with the RuleNameOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleNameOverride

`func (o *RuleUpdateProps) SetRuleNameOverride(v string)`

SetRuleNameOverride sets RuleNameOverride field to given value.

### HasRuleNameOverride

`func (o *RuleUpdateProps) HasRuleNameOverride() bool`

HasRuleNameOverride returns a boolean if a field has been set.

### GetSetup

`func (o *RuleUpdateProps) GetSetup() string`

GetSetup returns the Setup field if non-nil, zero value otherwise.

### GetSetupOk

`func (o *RuleUpdateProps) GetSetupOk() (*string, bool)`

GetSetupOk returns a tuple with the Setup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetup

`func (o *RuleUpdateProps) SetSetup(v string)`

SetSetup sets Setup field to given value.

### HasSetup

`func (o *RuleUpdateProps) HasSetup() bool`

HasSetup returns a boolean if a field has been set.

### GetSeverity

`func (o *RuleUpdateProps) GetSeverity() Severity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *RuleUpdateProps) GetSeverityOk() (*Severity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *RuleUpdateProps) SetSeverity(v Severity)`

SetSeverity sets Severity field to given value.


### GetSeverityMapping

`func (o *RuleUpdateProps) GetSeverityMapping() []SeverityMappingInner`

GetSeverityMapping returns the SeverityMapping field if non-nil, zero value otherwise.

### GetSeverityMappingOk

`func (o *RuleUpdateProps) GetSeverityMappingOk() (*[]SeverityMappingInner, bool)`

GetSeverityMappingOk returns a tuple with the SeverityMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityMapping

`func (o *RuleUpdateProps) SetSeverityMapping(v []SeverityMappingInner)`

SetSeverityMapping sets SeverityMapping field to given value.

### HasSeverityMapping

`func (o *RuleUpdateProps) HasSeverityMapping() bool`

HasSeverityMapping returns a boolean if a field has been set.

### GetTags

`func (o *RuleUpdateProps) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RuleUpdateProps) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RuleUpdateProps) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RuleUpdateProps) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThreat

`func (o *RuleUpdateProps) GetThreat() []Threat`

GetThreat returns the Threat field if non-nil, zero value otherwise.

### GetThreatOk

`func (o *RuleUpdateProps) GetThreatOk() (*[]Threat, bool)`

GetThreatOk returns a tuple with the Threat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreat

`func (o *RuleUpdateProps) SetThreat(v []Threat)`

SetThreat sets Threat field to given value.

### HasThreat

`func (o *RuleUpdateProps) HasThreat() bool`

HasThreat returns a boolean if a field has been set.

### GetThrottle

`func (o *RuleUpdateProps) GetThrottle() RuleActionThrottle`

GetThrottle returns the Throttle field if non-nil, zero value otherwise.

### GetThrottleOk

`func (o *RuleUpdateProps) GetThrottleOk() (*RuleActionThrottle, bool)`

GetThrottleOk returns a tuple with the Throttle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottle

`func (o *RuleUpdateProps) SetThrottle(v RuleActionThrottle)`

SetThrottle sets Throttle field to given value.

### HasThrottle

`func (o *RuleUpdateProps) HasThrottle() bool`

HasThrottle returns a boolean if a field has been set.

### GetTimelineId

`func (o *RuleUpdateProps) GetTimelineId() string`

GetTimelineId returns the TimelineId field if non-nil, zero value otherwise.

### GetTimelineIdOk

`func (o *RuleUpdateProps) GetTimelineIdOk() (*string, bool)`

GetTimelineIdOk returns a tuple with the TimelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineId

`func (o *RuleUpdateProps) SetTimelineId(v string)`

SetTimelineId sets TimelineId field to given value.

### HasTimelineId

`func (o *RuleUpdateProps) HasTimelineId() bool`

HasTimelineId returns a boolean if a field has been set.

### GetTimelineTitle

`func (o *RuleUpdateProps) GetTimelineTitle() string`

GetTimelineTitle returns the TimelineTitle field if non-nil, zero value otherwise.

### GetTimelineTitleOk

`func (o *RuleUpdateProps) GetTimelineTitleOk() (*string, bool)`

GetTimelineTitleOk returns a tuple with the TimelineTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineTitle

`func (o *RuleUpdateProps) SetTimelineTitle(v string)`

SetTimelineTitle sets TimelineTitle field to given value.

### HasTimelineTitle

`func (o *RuleUpdateProps) HasTimelineTitle() bool`

HasTimelineTitle returns a boolean if a field has been set.

### GetTimestampOverride

`func (o *RuleUpdateProps) GetTimestampOverride() string`

GetTimestampOverride returns the TimestampOverride field if non-nil, zero value otherwise.

### GetTimestampOverrideOk

`func (o *RuleUpdateProps) GetTimestampOverrideOk() (*string, bool)`

GetTimestampOverrideOk returns a tuple with the TimestampOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampOverride

`func (o *RuleUpdateProps) SetTimestampOverride(v string)`

SetTimestampOverride sets TimestampOverride field to given value.

### HasTimestampOverride

`func (o *RuleUpdateProps) HasTimestampOverride() bool`

HasTimestampOverride returns a boolean if a field has been set.

### GetTimestampOverrideFallbackDisabled

`func (o *RuleUpdateProps) GetTimestampOverrideFallbackDisabled() bool`

GetTimestampOverrideFallbackDisabled returns the TimestampOverrideFallbackDisabled field if non-nil, zero value otherwise.

### GetTimestampOverrideFallbackDisabledOk

`func (o *RuleUpdateProps) GetTimestampOverrideFallbackDisabledOk() (*bool, bool)`

GetTimestampOverrideFallbackDisabledOk returns a tuple with the TimestampOverrideFallbackDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampOverrideFallbackDisabled

`func (o *RuleUpdateProps) SetTimestampOverrideFallbackDisabled(v bool)`

SetTimestampOverrideFallbackDisabled sets TimestampOverrideFallbackDisabled field to given value.

### HasTimestampOverrideFallbackDisabled

`func (o *RuleUpdateProps) HasTimestampOverrideFallbackDisabled() bool`

HasTimestampOverrideFallbackDisabled returns a boolean if a field has been set.

### GetTo

`func (o *RuleUpdateProps) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *RuleUpdateProps) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *RuleUpdateProps) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *RuleUpdateProps) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetVersion

`func (o *RuleUpdateProps) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RuleUpdateProps) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RuleUpdateProps) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RuleUpdateProps) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLanguage

`func (o *RuleUpdateProps) GetLanguage() EsqlQueryLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *RuleUpdateProps) GetLanguageOk() (*EsqlQueryLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *RuleUpdateProps) SetLanguage(v EsqlQueryLanguage)`

SetLanguage sets Language field to given value.


### GetQuery

`func (o *RuleUpdateProps) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *RuleUpdateProps) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *RuleUpdateProps) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetType

`func (o *RuleUpdateProps) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RuleUpdateProps) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RuleUpdateProps) SetType(v string)`

SetType sets Type field to given value.


### GetAlertSuppression

`func (o *RuleUpdateProps) GetAlertSuppression() AlertSuppression`

GetAlertSuppression returns the AlertSuppression field if non-nil, zero value otherwise.

### GetAlertSuppressionOk

`func (o *RuleUpdateProps) GetAlertSuppressionOk() (*AlertSuppression, bool)`

GetAlertSuppressionOk returns a tuple with the AlertSuppression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertSuppression

`func (o *RuleUpdateProps) SetAlertSuppression(v AlertSuppression)`

SetAlertSuppression sets AlertSuppression field to given value.

### HasAlertSuppression

`func (o *RuleUpdateProps) HasAlertSuppression() bool`

HasAlertSuppression returns a boolean if a field has been set.

### GetDataViewId

`func (o *RuleUpdateProps) GetDataViewId() string`

GetDataViewId returns the DataViewId field if non-nil, zero value otherwise.

### GetDataViewIdOk

`func (o *RuleUpdateProps) GetDataViewIdOk() (*string, bool)`

GetDataViewIdOk returns a tuple with the DataViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataViewId

`func (o *RuleUpdateProps) SetDataViewId(v string)`

SetDataViewId sets DataViewId field to given value.

### HasDataViewId

`func (o *RuleUpdateProps) HasDataViewId() bool`

HasDataViewId returns a boolean if a field has been set.

### GetEventCategoryOverride

`func (o *RuleUpdateProps) GetEventCategoryOverride() string`

GetEventCategoryOverride returns the EventCategoryOverride field if non-nil, zero value otherwise.

### GetEventCategoryOverrideOk

`func (o *RuleUpdateProps) GetEventCategoryOverrideOk() (*string, bool)`

GetEventCategoryOverrideOk returns a tuple with the EventCategoryOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCategoryOverride

`func (o *RuleUpdateProps) SetEventCategoryOverride(v string)`

SetEventCategoryOverride sets EventCategoryOverride field to given value.

### HasEventCategoryOverride

`func (o *RuleUpdateProps) HasEventCategoryOverride() bool`

HasEventCategoryOverride returns a boolean if a field has been set.

### GetFilters

`func (o *RuleUpdateProps) GetFilters() []interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *RuleUpdateProps) GetFiltersOk() (*[]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *RuleUpdateProps) SetFilters(v []interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *RuleUpdateProps) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetIndex

`func (o *RuleUpdateProps) GetIndex() []string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *RuleUpdateProps) GetIndexOk() (*[]string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *RuleUpdateProps) SetIndex(v []string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *RuleUpdateProps) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetResponseActions

`func (o *RuleUpdateProps) GetResponseActions() []ResponseAction`

GetResponseActions returns the ResponseActions field if non-nil, zero value otherwise.

### GetResponseActionsOk

`func (o *RuleUpdateProps) GetResponseActionsOk() (*[]ResponseAction, bool)`

GetResponseActionsOk returns a tuple with the ResponseActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseActions

`func (o *RuleUpdateProps) SetResponseActions(v []ResponseAction)`

SetResponseActions sets ResponseActions field to given value.

### HasResponseActions

`func (o *RuleUpdateProps) HasResponseActions() bool`

HasResponseActions returns a boolean if a field has been set.

### GetTiebreakerField

`func (o *RuleUpdateProps) GetTiebreakerField() string`

GetTiebreakerField returns the TiebreakerField field if non-nil, zero value otherwise.

### GetTiebreakerFieldOk

`func (o *RuleUpdateProps) GetTiebreakerFieldOk() (*string, bool)`

GetTiebreakerFieldOk returns a tuple with the TiebreakerField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiebreakerField

`func (o *RuleUpdateProps) SetTiebreakerField(v string)`

SetTiebreakerField sets TiebreakerField field to given value.

### HasTiebreakerField

`func (o *RuleUpdateProps) HasTiebreakerField() bool`

HasTiebreakerField returns a boolean if a field has been set.

### GetTimestampField

`func (o *RuleUpdateProps) GetTimestampField() string`

GetTimestampField returns the TimestampField field if non-nil, zero value otherwise.

### GetTimestampFieldOk

`func (o *RuleUpdateProps) GetTimestampFieldOk() (*string, bool)`

GetTimestampFieldOk returns a tuple with the TimestampField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampField

`func (o *RuleUpdateProps) SetTimestampField(v string)`

SetTimestampField sets TimestampField field to given value.

### HasTimestampField

`func (o *RuleUpdateProps) HasTimestampField() bool`

HasTimestampField returns a boolean if a field has been set.

### GetSavedId

`func (o *RuleUpdateProps) GetSavedId() string`

GetSavedId returns the SavedId field if non-nil, zero value otherwise.

### GetSavedIdOk

`func (o *RuleUpdateProps) GetSavedIdOk() (*string, bool)`

GetSavedIdOk returns a tuple with the SavedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedId

`func (o *RuleUpdateProps) SetSavedId(v string)`

SetSavedId sets SavedId field to given value.


### GetThreshold

`func (o *RuleUpdateProps) GetThreshold() Threshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *RuleUpdateProps) GetThresholdOk() (*Threshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *RuleUpdateProps) SetThreshold(v Threshold)`

SetThreshold sets Threshold field to given value.


### GetThreatIndex

`func (o *RuleUpdateProps) GetThreatIndex() []string`

GetThreatIndex returns the ThreatIndex field if non-nil, zero value otherwise.

### GetThreatIndexOk

`func (o *RuleUpdateProps) GetThreatIndexOk() (*[]string, bool)`

GetThreatIndexOk returns a tuple with the ThreatIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatIndex

`func (o *RuleUpdateProps) SetThreatIndex(v []string)`

SetThreatIndex sets ThreatIndex field to given value.


### GetThreatMapping

`func (o *RuleUpdateProps) GetThreatMapping() []ThreatMappingInner`

GetThreatMapping returns the ThreatMapping field if non-nil, zero value otherwise.

### GetThreatMappingOk

`func (o *RuleUpdateProps) GetThreatMappingOk() (*[]ThreatMappingInner, bool)`

GetThreatMappingOk returns a tuple with the ThreatMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatMapping

`func (o *RuleUpdateProps) SetThreatMapping(v []ThreatMappingInner)`

SetThreatMapping sets ThreatMapping field to given value.


### GetThreatQuery

`func (o *RuleUpdateProps) GetThreatQuery() string`

GetThreatQuery returns the ThreatQuery field if non-nil, zero value otherwise.

### GetThreatQueryOk

`func (o *RuleUpdateProps) GetThreatQueryOk() (*string, bool)`

GetThreatQueryOk returns a tuple with the ThreatQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatQuery

`func (o *RuleUpdateProps) SetThreatQuery(v string)`

SetThreatQuery sets ThreatQuery field to given value.


### GetConcurrentSearches

`func (o *RuleUpdateProps) GetConcurrentSearches() int32`

GetConcurrentSearches returns the ConcurrentSearches field if non-nil, zero value otherwise.

### GetConcurrentSearchesOk

`func (o *RuleUpdateProps) GetConcurrentSearchesOk() (*int32, bool)`

GetConcurrentSearchesOk returns a tuple with the ConcurrentSearches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrentSearches

`func (o *RuleUpdateProps) SetConcurrentSearches(v int32)`

SetConcurrentSearches sets ConcurrentSearches field to given value.

### HasConcurrentSearches

`func (o *RuleUpdateProps) HasConcurrentSearches() bool`

HasConcurrentSearches returns a boolean if a field has been set.

### GetItemsPerSearch

`func (o *RuleUpdateProps) GetItemsPerSearch() int32`

GetItemsPerSearch returns the ItemsPerSearch field if non-nil, zero value otherwise.

### GetItemsPerSearchOk

`func (o *RuleUpdateProps) GetItemsPerSearchOk() (*int32, bool)`

GetItemsPerSearchOk returns a tuple with the ItemsPerSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsPerSearch

`func (o *RuleUpdateProps) SetItemsPerSearch(v int32)`

SetItemsPerSearch sets ItemsPerSearch field to given value.

### HasItemsPerSearch

`func (o *RuleUpdateProps) HasItemsPerSearch() bool`

HasItemsPerSearch returns a boolean if a field has been set.

### GetThreatFilters

`func (o *RuleUpdateProps) GetThreatFilters() []interface{}`

GetThreatFilters returns the ThreatFilters field if non-nil, zero value otherwise.

### GetThreatFiltersOk

`func (o *RuleUpdateProps) GetThreatFiltersOk() (*[]interface{}, bool)`

GetThreatFiltersOk returns a tuple with the ThreatFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatFilters

`func (o *RuleUpdateProps) SetThreatFilters(v []interface{})`

SetThreatFilters sets ThreatFilters field to given value.

### HasThreatFilters

`func (o *RuleUpdateProps) HasThreatFilters() bool`

HasThreatFilters returns a boolean if a field has been set.

### GetThreatIndicatorPath

`func (o *RuleUpdateProps) GetThreatIndicatorPath() string`

GetThreatIndicatorPath returns the ThreatIndicatorPath field if non-nil, zero value otherwise.

### GetThreatIndicatorPathOk

`func (o *RuleUpdateProps) GetThreatIndicatorPathOk() (*string, bool)`

GetThreatIndicatorPathOk returns a tuple with the ThreatIndicatorPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatIndicatorPath

`func (o *RuleUpdateProps) SetThreatIndicatorPath(v string)`

SetThreatIndicatorPath sets ThreatIndicatorPath field to given value.

### HasThreatIndicatorPath

`func (o *RuleUpdateProps) HasThreatIndicatorPath() bool`

HasThreatIndicatorPath returns a boolean if a field has been set.

### GetThreatLanguage

`func (o *RuleUpdateProps) GetThreatLanguage() KqlQueryLanguage`

GetThreatLanguage returns the ThreatLanguage field if non-nil, zero value otherwise.

### GetThreatLanguageOk

`func (o *RuleUpdateProps) GetThreatLanguageOk() (*KqlQueryLanguage, bool)`

GetThreatLanguageOk returns a tuple with the ThreatLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatLanguage

`func (o *RuleUpdateProps) SetThreatLanguage(v KqlQueryLanguage)`

SetThreatLanguage sets ThreatLanguage field to given value.

### HasThreatLanguage

`func (o *RuleUpdateProps) HasThreatLanguage() bool`

HasThreatLanguage returns a boolean if a field has been set.

### GetAnomalyThreshold

`func (o *RuleUpdateProps) GetAnomalyThreshold() int32`

GetAnomalyThreshold returns the AnomalyThreshold field if non-nil, zero value otherwise.

### GetAnomalyThresholdOk

`func (o *RuleUpdateProps) GetAnomalyThresholdOk() (*int32, bool)`

GetAnomalyThresholdOk returns a tuple with the AnomalyThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomalyThreshold

`func (o *RuleUpdateProps) SetAnomalyThreshold(v int32)`

SetAnomalyThreshold sets AnomalyThreshold field to given value.


### GetMachineLearningJobId

`func (o *RuleUpdateProps) GetMachineLearningJobId() MachineLearningJobId`

GetMachineLearningJobId returns the MachineLearningJobId field if non-nil, zero value otherwise.

### GetMachineLearningJobIdOk

`func (o *RuleUpdateProps) GetMachineLearningJobIdOk() (*MachineLearningJobId, bool)`

GetMachineLearningJobIdOk returns a tuple with the MachineLearningJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineLearningJobId

`func (o *RuleUpdateProps) SetMachineLearningJobId(v MachineLearningJobId)`

SetMachineLearningJobId sets MachineLearningJobId field to given value.


### GetHistoryWindowStart

`func (o *RuleUpdateProps) GetHistoryWindowStart() string`

GetHistoryWindowStart returns the HistoryWindowStart field if non-nil, zero value otherwise.

### GetHistoryWindowStartOk

`func (o *RuleUpdateProps) GetHistoryWindowStartOk() (*string, bool)`

GetHistoryWindowStartOk returns a tuple with the HistoryWindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryWindowStart

`func (o *RuleUpdateProps) SetHistoryWindowStart(v string)`

SetHistoryWindowStart sets HistoryWindowStart field to given value.


### GetNewTermsFields

`func (o *RuleUpdateProps) GetNewTermsFields() []string`

GetNewTermsFields returns the NewTermsFields field if non-nil, zero value otherwise.

### GetNewTermsFieldsOk

`func (o *RuleUpdateProps) GetNewTermsFieldsOk() (*[]string, bool)`

GetNewTermsFieldsOk returns a tuple with the NewTermsFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTermsFields

`func (o *RuleUpdateProps) SetNewTermsFields(v []string)`

SetNewTermsFields sets NewTermsFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


