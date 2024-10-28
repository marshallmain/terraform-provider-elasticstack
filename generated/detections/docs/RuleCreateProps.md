# RuleCreateProps

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

### NewRuleCreateProps

`func NewRuleCreateProps(description string, name string, riskScore int32, severity Severity, language EsqlQueryLanguage, query string, type_ string, savedId string, threshold Threshold, threatIndex []string, threatMapping []ThreatMappingInner, threatQuery string, anomalyThreshold int32, machineLearningJobId MachineLearningJobId, historyWindowStart string, newTermsFields []string, ) *RuleCreateProps`

NewRuleCreateProps instantiates a new RuleCreateProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleCreatePropsWithDefaults

`func NewRuleCreatePropsWithDefaults() *RuleCreateProps`

NewRuleCreatePropsWithDefaults instantiates a new RuleCreateProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *RuleCreateProps) GetActions() []RuleAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *RuleCreateProps) GetActionsOk() (*[]RuleAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *RuleCreateProps) SetActions(v []RuleAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *RuleCreateProps) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetAliasPurpose

`func (o *RuleCreateProps) GetAliasPurpose() SavedObjectResolveAliasPurpose`

GetAliasPurpose returns the AliasPurpose field if non-nil, zero value otherwise.

### GetAliasPurposeOk

`func (o *RuleCreateProps) GetAliasPurposeOk() (*SavedObjectResolveAliasPurpose, bool)`

GetAliasPurposeOk returns a tuple with the AliasPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasPurpose

`func (o *RuleCreateProps) SetAliasPurpose(v SavedObjectResolveAliasPurpose)`

SetAliasPurpose sets AliasPurpose field to given value.

### HasAliasPurpose

`func (o *RuleCreateProps) HasAliasPurpose() bool`

HasAliasPurpose returns a boolean if a field has been set.

### GetAliasTargetId

`func (o *RuleCreateProps) GetAliasTargetId() string`

GetAliasTargetId returns the AliasTargetId field if non-nil, zero value otherwise.

### GetAliasTargetIdOk

`func (o *RuleCreateProps) GetAliasTargetIdOk() (*string, bool)`

GetAliasTargetIdOk returns a tuple with the AliasTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasTargetId

`func (o *RuleCreateProps) SetAliasTargetId(v string)`

SetAliasTargetId sets AliasTargetId field to given value.

### HasAliasTargetId

`func (o *RuleCreateProps) HasAliasTargetId() bool`

HasAliasTargetId returns a boolean if a field has been set.

### GetAuthor

`func (o *RuleCreateProps) GetAuthor() []string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *RuleCreateProps) GetAuthorOk() (*[]string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *RuleCreateProps) SetAuthor(v []string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *RuleCreateProps) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetBuildingBlockType

`func (o *RuleCreateProps) GetBuildingBlockType() string`

GetBuildingBlockType returns the BuildingBlockType field if non-nil, zero value otherwise.

### GetBuildingBlockTypeOk

`func (o *RuleCreateProps) GetBuildingBlockTypeOk() (*string, bool)`

GetBuildingBlockTypeOk returns a tuple with the BuildingBlockType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildingBlockType

`func (o *RuleCreateProps) SetBuildingBlockType(v string)`

SetBuildingBlockType sets BuildingBlockType field to given value.

### HasBuildingBlockType

`func (o *RuleCreateProps) HasBuildingBlockType() bool`

HasBuildingBlockType returns a boolean if a field has been set.

### GetDescription

`func (o *RuleCreateProps) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RuleCreateProps) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RuleCreateProps) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEnabled

`func (o *RuleCreateProps) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RuleCreateProps) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RuleCreateProps) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RuleCreateProps) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExceptionsList

`func (o *RuleCreateProps) GetExceptionsList() []RuleExceptionList`

GetExceptionsList returns the ExceptionsList field if non-nil, zero value otherwise.

### GetExceptionsListOk

`func (o *RuleCreateProps) GetExceptionsListOk() (*[]RuleExceptionList, bool)`

GetExceptionsListOk returns a tuple with the ExceptionsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionsList

`func (o *RuleCreateProps) SetExceptionsList(v []RuleExceptionList)`

SetExceptionsList sets ExceptionsList field to given value.

### HasExceptionsList

`func (o *RuleCreateProps) HasExceptionsList() bool`

HasExceptionsList returns a boolean if a field has been set.

### GetFalsePositives

`func (o *RuleCreateProps) GetFalsePositives() []string`

GetFalsePositives returns the FalsePositives field if non-nil, zero value otherwise.

### GetFalsePositivesOk

`func (o *RuleCreateProps) GetFalsePositivesOk() (*[]string, bool)`

GetFalsePositivesOk returns a tuple with the FalsePositives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFalsePositives

`func (o *RuleCreateProps) SetFalsePositives(v []string)`

SetFalsePositives sets FalsePositives field to given value.

### HasFalsePositives

`func (o *RuleCreateProps) HasFalsePositives() bool`

HasFalsePositives returns a boolean if a field has been set.

### GetFrom

`func (o *RuleCreateProps) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *RuleCreateProps) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *RuleCreateProps) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *RuleCreateProps) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetInterval

`func (o *RuleCreateProps) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *RuleCreateProps) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *RuleCreateProps) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *RuleCreateProps) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetInvestigationFields

`func (o *RuleCreateProps) GetInvestigationFields() InvestigationFields`

GetInvestigationFields returns the InvestigationFields field if non-nil, zero value otherwise.

### GetInvestigationFieldsOk

`func (o *RuleCreateProps) GetInvestigationFieldsOk() (*InvestigationFields, bool)`

GetInvestigationFieldsOk returns a tuple with the InvestigationFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigationFields

`func (o *RuleCreateProps) SetInvestigationFields(v InvestigationFields)`

SetInvestigationFields sets InvestigationFields field to given value.

### HasInvestigationFields

`func (o *RuleCreateProps) HasInvestigationFields() bool`

HasInvestigationFields returns a boolean if a field has been set.

### GetLicense

`func (o *RuleCreateProps) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *RuleCreateProps) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *RuleCreateProps) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *RuleCreateProps) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetMaxSignals

`func (o *RuleCreateProps) GetMaxSignals() int32`

GetMaxSignals returns the MaxSignals field if non-nil, zero value otherwise.

### GetMaxSignalsOk

`func (o *RuleCreateProps) GetMaxSignalsOk() (*int32, bool)`

GetMaxSignalsOk returns a tuple with the MaxSignals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSignals

`func (o *RuleCreateProps) SetMaxSignals(v int32)`

SetMaxSignals sets MaxSignals field to given value.

### HasMaxSignals

`func (o *RuleCreateProps) HasMaxSignals() bool`

HasMaxSignals returns a boolean if a field has been set.

### GetMeta

`func (o *RuleCreateProps) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RuleCreateProps) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RuleCreateProps) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *RuleCreateProps) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetName

`func (o *RuleCreateProps) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleCreateProps) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleCreateProps) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *RuleCreateProps) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *RuleCreateProps) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *RuleCreateProps) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *RuleCreateProps) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNote

`func (o *RuleCreateProps) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *RuleCreateProps) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *RuleCreateProps) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *RuleCreateProps) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetOutcome

`func (o *RuleCreateProps) GetOutcome() SavedObjectResolveOutcome`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *RuleCreateProps) GetOutcomeOk() (*SavedObjectResolveOutcome, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *RuleCreateProps) SetOutcome(v SavedObjectResolveOutcome)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *RuleCreateProps) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetOutputIndex

`func (o *RuleCreateProps) GetOutputIndex() string`

GetOutputIndex returns the OutputIndex field if non-nil, zero value otherwise.

### GetOutputIndexOk

`func (o *RuleCreateProps) GetOutputIndexOk() (*string, bool)`

GetOutputIndexOk returns a tuple with the OutputIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputIndex

`func (o *RuleCreateProps) SetOutputIndex(v string)`

SetOutputIndex sets OutputIndex field to given value.

### HasOutputIndex

`func (o *RuleCreateProps) HasOutputIndex() bool`

HasOutputIndex returns a boolean if a field has been set.

### GetReferences

`func (o *RuleCreateProps) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *RuleCreateProps) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *RuleCreateProps) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *RuleCreateProps) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetRelatedIntegrations

`func (o *RuleCreateProps) GetRelatedIntegrations() []RelatedIntegration`

GetRelatedIntegrations returns the RelatedIntegrations field if non-nil, zero value otherwise.

### GetRelatedIntegrationsOk

`func (o *RuleCreateProps) GetRelatedIntegrationsOk() (*[]RelatedIntegration, bool)`

GetRelatedIntegrationsOk returns a tuple with the RelatedIntegrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedIntegrations

`func (o *RuleCreateProps) SetRelatedIntegrations(v []RelatedIntegration)`

SetRelatedIntegrations sets RelatedIntegrations field to given value.

### HasRelatedIntegrations

`func (o *RuleCreateProps) HasRelatedIntegrations() bool`

HasRelatedIntegrations returns a boolean if a field has been set.

### GetRequiredFields

`func (o *RuleCreateProps) GetRequiredFields() []RequiredFieldInput`

GetRequiredFields returns the RequiredFields field if non-nil, zero value otherwise.

### GetRequiredFieldsOk

`func (o *RuleCreateProps) GetRequiredFieldsOk() (*[]RequiredFieldInput, bool)`

GetRequiredFieldsOk returns a tuple with the RequiredFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredFields

`func (o *RuleCreateProps) SetRequiredFields(v []RequiredFieldInput)`

SetRequiredFields sets RequiredFields field to given value.

### HasRequiredFields

`func (o *RuleCreateProps) HasRequiredFields() bool`

HasRequiredFields returns a boolean if a field has been set.

### GetRiskScore

`func (o *RuleCreateProps) GetRiskScore() int32`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *RuleCreateProps) GetRiskScoreOk() (*int32, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *RuleCreateProps) SetRiskScore(v int32)`

SetRiskScore sets RiskScore field to given value.


### GetRiskScoreMapping

`func (o *RuleCreateProps) GetRiskScoreMapping() []RiskScoreMappingInner`

GetRiskScoreMapping returns the RiskScoreMapping field if non-nil, zero value otherwise.

### GetRiskScoreMappingOk

`func (o *RuleCreateProps) GetRiskScoreMappingOk() (*[]RiskScoreMappingInner, bool)`

GetRiskScoreMappingOk returns a tuple with the RiskScoreMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScoreMapping

`func (o *RuleCreateProps) SetRiskScoreMapping(v []RiskScoreMappingInner)`

SetRiskScoreMapping sets RiskScoreMapping field to given value.

### HasRiskScoreMapping

`func (o *RuleCreateProps) HasRiskScoreMapping() bool`

HasRiskScoreMapping returns a boolean if a field has been set.

### GetRuleId

`func (o *RuleCreateProps) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *RuleCreateProps) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *RuleCreateProps) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *RuleCreateProps) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetRuleNameOverride

`func (o *RuleCreateProps) GetRuleNameOverride() string`

GetRuleNameOverride returns the RuleNameOverride field if non-nil, zero value otherwise.

### GetRuleNameOverrideOk

`func (o *RuleCreateProps) GetRuleNameOverrideOk() (*string, bool)`

GetRuleNameOverrideOk returns a tuple with the RuleNameOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleNameOverride

`func (o *RuleCreateProps) SetRuleNameOverride(v string)`

SetRuleNameOverride sets RuleNameOverride field to given value.

### HasRuleNameOverride

`func (o *RuleCreateProps) HasRuleNameOverride() bool`

HasRuleNameOverride returns a boolean if a field has been set.

### GetSetup

`func (o *RuleCreateProps) GetSetup() string`

GetSetup returns the Setup field if non-nil, zero value otherwise.

### GetSetupOk

`func (o *RuleCreateProps) GetSetupOk() (*string, bool)`

GetSetupOk returns a tuple with the Setup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetup

`func (o *RuleCreateProps) SetSetup(v string)`

SetSetup sets Setup field to given value.

### HasSetup

`func (o *RuleCreateProps) HasSetup() bool`

HasSetup returns a boolean if a field has been set.

### GetSeverity

`func (o *RuleCreateProps) GetSeverity() Severity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *RuleCreateProps) GetSeverityOk() (*Severity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *RuleCreateProps) SetSeverity(v Severity)`

SetSeverity sets Severity field to given value.


### GetSeverityMapping

`func (o *RuleCreateProps) GetSeverityMapping() []SeverityMappingInner`

GetSeverityMapping returns the SeverityMapping field if non-nil, zero value otherwise.

### GetSeverityMappingOk

`func (o *RuleCreateProps) GetSeverityMappingOk() (*[]SeverityMappingInner, bool)`

GetSeverityMappingOk returns a tuple with the SeverityMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityMapping

`func (o *RuleCreateProps) SetSeverityMapping(v []SeverityMappingInner)`

SetSeverityMapping sets SeverityMapping field to given value.

### HasSeverityMapping

`func (o *RuleCreateProps) HasSeverityMapping() bool`

HasSeverityMapping returns a boolean if a field has been set.

### GetTags

`func (o *RuleCreateProps) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RuleCreateProps) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RuleCreateProps) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RuleCreateProps) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThreat

`func (o *RuleCreateProps) GetThreat() []Threat`

GetThreat returns the Threat field if non-nil, zero value otherwise.

### GetThreatOk

`func (o *RuleCreateProps) GetThreatOk() (*[]Threat, bool)`

GetThreatOk returns a tuple with the Threat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreat

`func (o *RuleCreateProps) SetThreat(v []Threat)`

SetThreat sets Threat field to given value.

### HasThreat

`func (o *RuleCreateProps) HasThreat() bool`

HasThreat returns a boolean if a field has been set.

### GetThrottle

`func (o *RuleCreateProps) GetThrottle() RuleActionThrottle`

GetThrottle returns the Throttle field if non-nil, zero value otherwise.

### GetThrottleOk

`func (o *RuleCreateProps) GetThrottleOk() (*RuleActionThrottle, bool)`

GetThrottleOk returns a tuple with the Throttle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottle

`func (o *RuleCreateProps) SetThrottle(v RuleActionThrottle)`

SetThrottle sets Throttle field to given value.

### HasThrottle

`func (o *RuleCreateProps) HasThrottle() bool`

HasThrottle returns a boolean if a field has been set.

### GetTimelineId

`func (o *RuleCreateProps) GetTimelineId() string`

GetTimelineId returns the TimelineId field if non-nil, zero value otherwise.

### GetTimelineIdOk

`func (o *RuleCreateProps) GetTimelineIdOk() (*string, bool)`

GetTimelineIdOk returns a tuple with the TimelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineId

`func (o *RuleCreateProps) SetTimelineId(v string)`

SetTimelineId sets TimelineId field to given value.

### HasTimelineId

`func (o *RuleCreateProps) HasTimelineId() bool`

HasTimelineId returns a boolean if a field has been set.

### GetTimelineTitle

`func (o *RuleCreateProps) GetTimelineTitle() string`

GetTimelineTitle returns the TimelineTitle field if non-nil, zero value otherwise.

### GetTimelineTitleOk

`func (o *RuleCreateProps) GetTimelineTitleOk() (*string, bool)`

GetTimelineTitleOk returns a tuple with the TimelineTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineTitle

`func (o *RuleCreateProps) SetTimelineTitle(v string)`

SetTimelineTitle sets TimelineTitle field to given value.

### HasTimelineTitle

`func (o *RuleCreateProps) HasTimelineTitle() bool`

HasTimelineTitle returns a boolean if a field has been set.

### GetTimestampOverride

`func (o *RuleCreateProps) GetTimestampOverride() string`

GetTimestampOverride returns the TimestampOverride field if non-nil, zero value otherwise.

### GetTimestampOverrideOk

`func (o *RuleCreateProps) GetTimestampOverrideOk() (*string, bool)`

GetTimestampOverrideOk returns a tuple with the TimestampOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampOverride

`func (o *RuleCreateProps) SetTimestampOverride(v string)`

SetTimestampOverride sets TimestampOverride field to given value.

### HasTimestampOverride

`func (o *RuleCreateProps) HasTimestampOverride() bool`

HasTimestampOverride returns a boolean if a field has been set.

### GetTimestampOverrideFallbackDisabled

`func (o *RuleCreateProps) GetTimestampOverrideFallbackDisabled() bool`

GetTimestampOverrideFallbackDisabled returns the TimestampOverrideFallbackDisabled field if non-nil, zero value otherwise.

### GetTimestampOverrideFallbackDisabledOk

`func (o *RuleCreateProps) GetTimestampOverrideFallbackDisabledOk() (*bool, bool)`

GetTimestampOverrideFallbackDisabledOk returns a tuple with the TimestampOverrideFallbackDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampOverrideFallbackDisabled

`func (o *RuleCreateProps) SetTimestampOverrideFallbackDisabled(v bool)`

SetTimestampOverrideFallbackDisabled sets TimestampOverrideFallbackDisabled field to given value.

### HasTimestampOverrideFallbackDisabled

`func (o *RuleCreateProps) HasTimestampOverrideFallbackDisabled() bool`

HasTimestampOverrideFallbackDisabled returns a boolean if a field has been set.

### GetTo

`func (o *RuleCreateProps) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *RuleCreateProps) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *RuleCreateProps) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *RuleCreateProps) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetVersion

`func (o *RuleCreateProps) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RuleCreateProps) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RuleCreateProps) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RuleCreateProps) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLanguage

`func (o *RuleCreateProps) GetLanguage() EsqlQueryLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *RuleCreateProps) GetLanguageOk() (*EsqlQueryLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *RuleCreateProps) SetLanguage(v EsqlQueryLanguage)`

SetLanguage sets Language field to given value.


### GetQuery

`func (o *RuleCreateProps) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *RuleCreateProps) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *RuleCreateProps) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetType

`func (o *RuleCreateProps) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RuleCreateProps) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RuleCreateProps) SetType(v string)`

SetType sets Type field to given value.


### GetAlertSuppression

`func (o *RuleCreateProps) GetAlertSuppression() AlertSuppression`

GetAlertSuppression returns the AlertSuppression field if non-nil, zero value otherwise.

### GetAlertSuppressionOk

`func (o *RuleCreateProps) GetAlertSuppressionOk() (*AlertSuppression, bool)`

GetAlertSuppressionOk returns a tuple with the AlertSuppression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertSuppression

`func (o *RuleCreateProps) SetAlertSuppression(v AlertSuppression)`

SetAlertSuppression sets AlertSuppression field to given value.

### HasAlertSuppression

`func (o *RuleCreateProps) HasAlertSuppression() bool`

HasAlertSuppression returns a boolean if a field has been set.

### GetDataViewId

`func (o *RuleCreateProps) GetDataViewId() string`

GetDataViewId returns the DataViewId field if non-nil, zero value otherwise.

### GetDataViewIdOk

`func (o *RuleCreateProps) GetDataViewIdOk() (*string, bool)`

GetDataViewIdOk returns a tuple with the DataViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataViewId

`func (o *RuleCreateProps) SetDataViewId(v string)`

SetDataViewId sets DataViewId field to given value.

### HasDataViewId

`func (o *RuleCreateProps) HasDataViewId() bool`

HasDataViewId returns a boolean if a field has been set.

### GetEventCategoryOverride

`func (o *RuleCreateProps) GetEventCategoryOverride() string`

GetEventCategoryOverride returns the EventCategoryOverride field if non-nil, zero value otherwise.

### GetEventCategoryOverrideOk

`func (o *RuleCreateProps) GetEventCategoryOverrideOk() (*string, bool)`

GetEventCategoryOverrideOk returns a tuple with the EventCategoryOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCategoryOverride

`func (o *RuleCreateProps) SetEventCategoryOverride(v string)`

SetEventCategoryOverride sets EventCategoryOverride field to given value.

### HasEventCategoryOverride

`func (o *RuleCreateProps) HasEventCategoryOverride() bool`

HasEventCategoryOverride returns a boolean if a field has been set.

### GetFilters

`func (o *RuleCreateProps) GetFilters() []interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *RuleCreateProps) GetFiltersOk() (*[]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *RuleCreateProps) SetFilters(v []interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *RuleCreateProps) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetIndex

`func (o *RuleCreateProps) GetIndex() []string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *RuleCreateProps) GetIndexOk() (*[]string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *RuleCreateProps) SetIndex(v []string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *RuleCreateProps) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetResponseActions

`func (o *RuleCreateProps) GetResponseActions() []ResponseAction`

GetResponseActions returns the ResponseActions field if non-nil, zero value otherwise.

### GetResponseActionsOk

`func (o *RuleCreateProps) GetResponseActionsOk() (*[]ResponseAction, bool)`

GetResponseActionsOk returns a tuple with the ResponseActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseActions

`func (o *RuleCreateProps) SetResponseActions(v []ResponseAction)`

SetResponseActions sets ResponseActions field to given value.

### HasResponseActions

`func (o *RuleCreateProps) HasResponseActions() bool`

HasResponseActions returns a boolean if a field has been set.

### GetTiebreakerField

`func (o *RuleCreateProps) GetTiebreakerField() string`

GetTiebreakerField returns the TiebreakerField field if non-nil, zero value otherwise.

### GetTiebreakerFieldOk

`func (o *RuleCreateProps) GetTiebreakerFieldOk() (*string, bool)`

GetTiebreakerFieldOk returns a tuple with the TiebreakerField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiebreakerField

`func (o *RuleCreateProps) SetTiebreakerField(v string)`

SetTiebreakerField sets TiebreakerField field to given value.

### HasTiebreakerField

`func (o *RuleCreateProps) HasTiebreakerField() bool`

HasTiebreakerField returns a boolean if a field has been set.

### GetTimestampField

`func (o *RuleCreateProps) GetTimestampField() string`

GetTimestampField returns the TimestampField field if non-nil, zero value otherwise.

### GetTimestampFieldOk

`func (o *RuleCreateProps) GetTimestampFieldOk() (*string, bool)`

GetTimestampFieldOk returns a tuple with the TimestampField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampField

`func (o *RuleCreateProps) SetTimestampField(v string)`

SetTimestampField sets TimestampField field to given value.

### HasTimestampField

`func (o *RuleCreateProps) HasTimestampField() bool`

HasTimestampField returns a boolean if a field has been set.

### GetSavedId

`func (o *RuleCreateProps) GetSavedId() string`

GetSavedId returns the SavedId field if non-nil, zero value otherwise.

### GetSavedIdOk

`func (o *RuleCreateProps) GetSavedIdOk() (*string, bool)`

GetSavedIdOk returns a tuple with the SavedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedId

`func (o *RuleCreateProps) SetSavedId(v string)`

SetSavedId sets SavedId field to given value.


### GetThreshold

`func (o *RuleCreateProps) GetThreshold() Threshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *RuleCreateProps) GetThresholdOk() (*Threshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *RuleCreateProps) SetThreshold(v Threshold)`

SetThreshold sets Threshold field to given value.


### GetThreatIndex

`func (o *RuleCreateProps) GetThreatIndex() []string`

GetThreatIndex returns the ThreatIndex field if non-nil, zero value otherwise.

### GetThreatIndexOk

`func (o *RuleCreateProps) GetThreatIndexOk() (*[]string, bool)`

GetThreatIndexOk returns a tuple with the ThreatIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatIndex

`func (o *RuleCreateProps) SetThreatIndex(v []string)`

SetThreatIndex sets ThreatIndex field to given value.


### GetThreatMapping

`func (o *RuleCreateProps) GetThreatMapping() []ThreatMappingInner`

GetThreatMapping returns the ThreatMapping field if non-nil, zero value otherwise.

### GetThreatMappingOk

`func (o *RuleCreateProps) GetThreatMappingOk() (*[]ThreatMappingInner, bool)`

GetThreatMappingOk returns a tuple with the ThreatMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatMapping

`func (o *RuleCreateProps) SetThreatMapping(v []ThreatMappingInner)`

SetThreatMapping sets ThreatMapping field to given value.


### GetThreatQuery

`func (o *RuleCreateProps) GetThreatQuery() string`

GetThreatQuery returns the ThreatQuery field if non-nil, zero value otherwise.

### GetThreatQueryOk

`func (o *RuleCreateProps) GetThreatQueryOk() (*string, bool)`

GetThreatQueryOk returns a tuple with the ThreatQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatQuery

`func (o *RuleCreateProps) SetThreatQuery(v string)`

SetThreatQuery sets ThreatQuery field to given value.


### GetConcurrentSearches

`func (o *RuleCreateProps) GetConcurrentSearches() int32`

GetConcurrentSearches returns the ConcurrentSearches field if non-nil, zero value otherwise.

### GetConcurrentSearchesOk

`func (o *RuleCreateProps) GetConcurrentSearchesOk() (*int32, bool)`

GetConcurrentSearchesOk returns a tuple with the ConcurrentSearches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrentSearches

`func (o *RuleCreateProps) SetConcurrentSearches(v int32)`

SetConcurrentSearches sets ConcurrentSearches field to given value.

### HasConcurrentSearches

`func (o *RuleCreateProps) HasConcurrentSearches() bool`

HasConcurrentSearches returns a boolean if a field has been set.

### GetItemsPerSearch

`func (o *RuleCreateProps) GetItemsPerSearch() int32`

GetItemsPerSearch returns the ItemsPerSearch field if non-nil, zero value otherwise.

### GetItemsPerSearchOk

`func (o *RuleCreateProps) GetItemsPerSearchOk() (*int32, bool)`

GetItemsPerSearchOk returns a tuple with the ItemsPerSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsPerSearch

`func (o *RuleCreateProps) SetItemsPerSearch(v int32)`

SetItemsPerSearch sets ItemsPerSearch field to given value.

### HasItemsPerSearch

`func (o *RuleCreateProps) HasItemsPerSearch() bool`

HasItemsPerSearch returns a boolean if a field has been set.

### GetThreatFilters

`func (o *RuleCreateProps) GetThreatFilters() []interface{}`

GetThreatFilters returns the ThreatFilters field if non-nil, zero value otherwise.

### GetThreatFiltersOk

`func (o *RuleCreateProps) GetThreatFiltersOk() (*[]interface{}, bool)`

GetThreatFiltersOk returns a tuple with the ThreatFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatFilters

`func (o *RuleCreateProps) SetThreatFilters(v []interface{})`

SetThreatFilters sets ThreatFilters field to given value.

### HasThreatFilters

`func (o *RuleCreateProps) HasThreatFilters() bool`

HasThreatFilters returns a boolean if a field has been set.

### GetThreatIndicatorPath

`func (o *RuleCreateProps) GetThreatIndicatorPath() string`

GetThreatIndicatorPath returns the ThreatIndicatorPath field if non-nil, zero value otherwise.

### GetThreatIndicatorPathOk

`func (o *RuleCreateProps) GetThreatIndicatorPathOk() (*string, bool)`

GetThreatIndicatorPathOk returns a tuple with the ThreatIndicatorPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatIndicatorPath

`func (o *RuleCreateProps) SetThreatIndicatorPath(v string)`

SetThreatIndicatorPath sets ThreatIndicatorPath field to given value.

### HasThreatIndicatorPath

`func (o *RuleCreateProps) HasThreatIndicatorPath() bool`

HasThreatIndicatorPath returns a boolean if a field has been set.

### GetThreatLanguage

`func (o *RuleCreateProps) GetThreatLanguage() KqlQueryLanguage`

GetThreatLanguage returns the ThreatLanguage field if non-nil, zero value otherwise.

### GetThreatLanguageOk

`func (o *RuleCreateProps) GetThreatLanguageOk() (*KqlQueryLanguage, bool)`

GetThreatLanguageOk returns a tuple with the ThreatLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatLanguage

`func (o *RuleCreateProps) SetThreatLanguage(v KqlQueryLanguage)`

SetThreatLanguage sets ThreatLanguage field to given value.

### HasThreatLanguage

`func (o *RuleCreateProps) HasThreatLanguage() bool`

HasThreatLanguage returns a boolean if a field has been set.

### GetAnomalyThreshold

`func (o *RuleCreateProps) GetAnomalyThreshold() int32`

GetAnomalyThreshold returns the AnomalyThreshold field if non-nil, zero value otherwise.

### GetAnomalyThresholdOk

`func (o *RuleCreateProps) GetAnomalyThresholdOk() (*int32, bool)`

GetAnomalyThresholdOk returns a tuple with the AnomalyThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomalyThreshold

`func (o *RuleCreateProps) SetAnomalyThreshold(v int32)`

SetAnomalyThreshold sets AnomalyThreshold field to given value.


### GetMachineLearningJobId

`func (o *RuleCreateProps) GetMachineLearningJobId() MachineLearningJobId`

GetMachineLearningJobId returns the MachineLearningJobId field if non-nil, zero value otherwise.

### GetMachineLearningJobIdOk

`func (o *RuleCreateProps) GetMachineLearningJobIdOk() (*MachineLearningJobId, bool)`

GetMachineLearningJobIdOk returns a tuple with the MachineLearningJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineLearningJobId

`func (o *RuleCreateProps) SetMachineLearningJobId(v MachineLearningJobId)`

SetMachineLearningJobId sets MachineLearningJobId field to given value.


### GetHistoryWindowStart

`func (o *RuleCreateProps) GetHistoryWindowStart() string`

GetHistoryWindowStart returns the HistoryWindowStart field if non-nil, zero value otherwise.

### GetHistoryWindowStartOk

`func (o *RuleCreateProps) GetHistoryWindowStartOk() (*string, bool)`

GetHistoryWindowStartOk returns a tuple with the HistoryWindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryWindowStart

`func (o *RuleCreateProps) SetHistoryWindowStart(v string)`

SetHistoryWindowStart sets HistoryWindowStart field to given value.


### GetNewTermsFields

`func (o *RuleCreateProps) GetNewTermsFields() []string`

GetNewTermsFields returns the NewTermsFields field if non-nil, zero value otherwise.

### GetNewTermsFieldsOk

`func (o *RuleCreateProps) GetNewTermsFieldsOk() (*[]string, bool)`

GetNewTermsFieldsOk returns a tuple with the NewTermsFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTermsFields

`func (o *RuleCreateProps) SetNewTermsFields(v []string)`

SetNewTermsFields sets NewTermsFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


