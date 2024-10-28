# EsqlRulePatchProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertSuppression** | Pointer to [**AlertSuppression**](AlertSuppression.md) |  | [optional] 
**ResponseActions** | Pointer to [**[]ResponseAction**](ResponseAction.md) |  | [optional] 
**Actions** | Pointer to [**[]RuleAction**](RuleAction.md) |  | [optional] 
**AliasPurpose** | Pointer to [**SavedObjectResolveAliasPurpose**](SavedObjectResolveAliasPurpose.md) |  | [optional] 
**AliasTargetId** | Pointer to **string** |  | [optional] 
**Author** | Pointer to **[]string** |  | [optional] 
**BuildingBlockType** | Pointer to **string** | Determines if the rule acts as a building block. By default, building-block alerts are not displayed in the UI. These rules are used as a foundation for other rules that do generate alerts. Its value must be default. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** | Determines whether the rule is enabled. | [optional] 
**ExceptionsList** | Pointer to [**[]RuleExceptionList**](RuleExceptionList.md) |  | [optional] 
**FalsePositives** | Pointer to **[]string** |  | [optional] 
**From** | Pointer to **string** | Time from which data is analyzed each time the rule runs, using a date math range. For example, now-4200s means the rule analyzes data from 70 minutes before its start time. Defaults to now-6m (analyzes data from 6 minutes before the start time). | [optional] 
**Id** | Pointer to **string** | A universally unique identifier | [optional] 
**Interval** | Pointer to **string** | Frequency of rule execution, using a date math range. For example, \&quot;1h\&quot; means the rule runs every hour. Defaults to 5m (5 minutes). | [optional] 
**InvestigationFields** | Pointer to [**InvestigationFields**](InvestigationFields.md) |  | [optional] 
**Language** | Pointer to [**EsqlQueryLanguage**](EsqlQueryLanguage.md) |  | [optional] 
**License** | Pointer to **string** | The rule&#39;s license. | [optional] 
**MaxSignals** | Pointer to **int32** |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** | Has no effect. | [optional] 
**Note** | Pointer to **string** | Notes to help investigate alerts produced by the rule. | [optional] 
**Outcome** | Pointer to [**SavedObjectResolveOutcome**](SavedObjectResolveOutcome.md) |  | [optional] 
**OutputIndex** | Pointer to **string** | (deprecated) Has no effect. | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**RelatedIntegrations** | Pointer to [**[]RelatedIntegration**](RelatedIntegration.md) |  | [optional] 
**RequiredFields** | Pointer to [**[]RequiredFieldInput**](RequiredFieldInput.md) |  | [optional] 
**RiskScore** | Pointer to **int32** | Risk score (0 to 100) | [optional] 
**RiskScoreMapping** | Pointer to [**[]RiskScoreMappingInner**](RiskScoreMappingInner.md) | Overrides generated alerts&#39; risk_score with a value from the source event | [optional] 
**RuleId** | Pointer to **string** | Could be any string, not necessarily a UUID | [optional] 
**RuleNameOverride** | Pointer to **string** | Sets the source field for the alert&#39;s signal.rule.name value | [optional] 
**Setup** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to [**Severity**](Severity.md) |  | [optional] 
**SeverityMapping** | Pointer to [**[]SeverityMappingInner**](SeverityMappingInner.md) | Overrides generated alerts&#39; severity with values from the source event | [optional] 
**Tags** | Pointer to **[]string** | String array containing words and phrases to help categorize, filter, and search rules. Defaults to an empty array. | [optional] 
**Threat** | Pointer to [**[]Threat**](Threat.md) |  | [optional] 
**Throttle** | Pointer to [**RuleActionThrottle**](RuleActionThrottle.md) |  | [optional] 
**TimelineId** | Pointer to **string** | Timeline template ID | [optional] 
**TimelineTitle** | Pointer to **string** | Timeline template title | [optional] 
**TimestampOverride** | Pointer to **string** | Sets the time field used to query indices | [optional] 
**TimestampOverrideFallbackDisabled** | Pointer to **bool** | Disables the fallback to the event&#39;s @timestamp field | [optional] 
**To** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | Rule type | [optional] 
**Version** | Pointer to **int32** | The rule&#39;s version number. | [optional] 

## Methods

### NewEsqlRulePatchProps

`func NewEsqlRulePatchProps() *EsqlRulePatchProps`

NewEsqlRulePatchProps instantiates a new EsqlRulePatchProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsqlRulePatchPropsWithDefaults

`func NewEsqlRulePatchPropsWithDefaults() *EsqlRulePatchProps`

NewEsqlRulePatchPropsWithDefaults instantiates a new EsqlRulePatchProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertSuppression

`func (o *EsqlRulePatchProps) GetAlertSuppression() AlertSuppression`

GetAlertSuppression returns the AlertSuppression field if non-nil, zero value otherwise.

### GetAlertSuppressionOk

`func (o *EsqlRulePatchProps) GetAlertSuppressionOk() (*AlertSuppression, bool)`

GetAlertSuppressionOk returns a tuple with the AlertSuppression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertSuppression

`func (o *EsqlRulePatchProps) SetAlertSuppression(v AlertSuppression)`

SetAlertSuppression sets AlertSuppression field to given value.

### HasAlertSuppression

`func (o *EsqlRulePatchProps) HasAlertSuppression() bool`

HasAlertSuppression returns a boolean if a field has been set.

### GetResponseActions

`func (o *EsqlRulePatchProps) GetResponseActions() []ResponseAction`

GetResponseActions returns the ResponseActions field if non-nil, zero value otherwise.

### GetResponseActionsOk

`func (o *EsqlRulePatchProps) GetResponseActionsOk() (*[]ResponseAction, bool)`

GetResponseActionsOk returns a tuple with the ResponseActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseActions

`func (o *EsqlRulePatchProps) SetResponseActions(v []ResponseAction)`

SetResponseActions sets ResponseActions field to given value.

### HasResponseActions

`func (o *EsqlRulePatchProps) HasResponseActions() bool`

HasResponseActions returns a boolean if a field has been set.

### GetActions

`func (o *EsqlRulePatchProps) GetActions() []RuleAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *EsqlRulePatchProps) GetActionsOk() (*[]RuleAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *EsqlRulePatchProps) SetActions(v []RuleAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *EsqlRulePatchProps) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetAliasPurpose

`func (o *EsqlRulePatchProps) GetAliasPurpose() SavedObjectResolveAliasPurpose`

GetAliasPurpose returns the AliasPurpose field if non-nil, zero value otherwise.

### GetAliasPurposeOk

`func (o *EsqlRulePatchProps) GetAliasPurposeOk() (*SavedObjectResolveAliasPurpose, bool)`

GetAliasPurposeOk returns a tuple with the AliasPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasPurpose

`func (o *EsqlRulePatchProps) SetAliasPurpose(v SavedObjectResolveAliasPurpose)`

SetAliasPurpose sets AliasPurpose field to given value.

### HasAliasPurpose

`func (o *EsqlRulePatchProps) HasAliasPurpose() bool`

HasAliasPurpose returns a boolean if a field has been set.

### GetAliasTargetId

`func (o *EsqlRulePatchProps) GetAliasTargetId() string`

GetAliasTargetId returns the AliasTargetId field if non-nil, zero value otherwise.

### GetAliasTargetIdOk

`func (o *EsqlRulePatchProps) GetAliasTargetIdOk() (*string, bool)`

GetAliasTargetIdOk returns a tuple with the AliasTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasTargetId

`func (o *EsqlRulePatchProps) SetAliasTargetId(v string)`

SetAliasTargetId sets AliasTargetId field to given value.

### HasAliasTargetId

`func (o *EsqlRulePatchProps) HasAliasTargetId() bool`

HasAliasTargetId returns a boolean if a field has been set.

### GetAuthor

`func (o *EsqlRulePatchProps) GetAuthor() []string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *EsqlRulePatchProps) GetAuthorOk() (*[]string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *EsqlRulePatchProps) SetAuthor(v []string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *EsqlRulePatchProps) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetBuildingBlockType

`func (o *EsqlRulePatchProps) GetBuildingBlockType() string`

GetBuildingBlockType returns the BuildingBlockType field if non-nil, zero value otherwise.

### GetBuildingBlockTypeOk

`func (o *EsqlRulePatchProps) GetBuildingBlockTypeOk() (*string, bool)`

GetBuildingBlockTypeOk returns a tuple with the BuildingBlockType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildingBlockType

`func (o *EsqlRulePatchProps) SetBuildingBlockType(v string)`

SetBuildingBlockType sets BuildingBlockType field to given value.

### HasBuildingBlockType

`func (o *EsqlRulePatchProps) HasBuildingBlockType() bool`

HasBuildingBlockType returns a boolean if a field has been set.

### GetDescription

`func (o *EsqlRulePatchProps) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EsqlRulePatchProps) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EsqlRulePatchProps) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EsqlRulePatchProps) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *EsqlRulePatchProps) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EsqlRulePatchProps) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EsqlRulePatchProps) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *EsqlRulePatchProps) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExceptionsList

`func (o *EsqlRulePatchProps) GetExceptionsList() []RuleExceptionList`

GetExceptionsList returns the ExceptionsList field if non-nil, zero value otherwise.

### GetExceptionsListOk

`func (o *EsqlRulePatchProps) GetExceptionsListOk() (*[]RuleExceptionList, bool)`

GetExceptionsListOk returns a tuple with the ExceptionsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionsList

`func (o *EsqlRulePatchProps) SetExceptionsList(v []RuleExceptionList)`

SetExceptionsList sets ExceptionsList field to given value.

### HasExceptionsList

`func (o *EsqlRulePatchProps) HasExceptionsList() bool`

HasExceptionsList returns a boolean if a field has been set.

### GetFalsePositives

`func (o *EsqlRulePatchProps) GetFalsePositives() []string`

GetFalsePositives returns the FalsePositives field if non-nil, zero value otherwise.

### GetFalsePositivesOk

`func (o *EsqlRulePatchProps) GetFalsePositivesOk() (*[]string, bool)`

GetFalsePositivesOk returns a tuple with the FalsePositives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFalsePositives

`func (o *EsqlRulePatchProps) SetFalsePositives(v []string)`

SetFalsePositives sets FalsePositives field to given value.

### HasFalsePositives

`func (o *EsqlRulePatchProps) HasFalsePositives() bool`

HasFalsePositives returns a boolean if a field has been set.

### GetFrom

`func (o *EsqlRulePatchProps) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *EsqlRulePatchProps) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *EsqlRulePatchProps) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *EsqlRulePatchProps) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetId

`func (o *EsqlRulePatchProps) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EsqlRulePatchProps) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EsqlRulePatchProps) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EsqlRulePatchProps) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterval

`func (o *EsqlRulePatchProps) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *EsqlRulePatchProps) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *EsqlRulePatchProps) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *EsqlRulePatchProps) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetInvestigationFields

`func (o *EsqlRulePatchProps) GetInvestigationFields() InvestigationFields`

GetInvestigationFields returns the InvestigationFields field if non-nil, zero value otherwise.

### GetInvestigationFieldsOk

`func (o *EsqlRulePatchProps) GetInvestigationFieldsOk() (*InvestigationFields, bool)`

GetInvestigationFieldsOk returns a tuple with the InvestigationFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigationFields

`func (o *EsqlRulePatchProps) SetInvestigationFields(v InvestigationFields)`

SetInvestigationFields sets InvestigationFields field to given value.

### HasInvestigationFields

`func (o *EsqlRulePatchProps) HasInvestigationFields() bool`

HasInvestigationFields returns a boolean if a field has been set.

### GetLanguage

`func (o *EsqlRulePatchProps) GetLanguage() EsqlQueryLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *EsqlRulePatchProps) GetLanguageOk() (*EsqlQueryLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *EsqlRulePatchProps) SetLanguage(v EsqlQueryLanguage)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *EsqlRulePatchProps) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetLicense

`func (o *EsqlRulePatchProps) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *EsqlRulePatchProps) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *EsqlRulePatchProps) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *EsqlRulePatchProps) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetMaxSignals

`func (o *EsqlRulePatchProps) GetMaxSignals() int32`

GetMaxSignals returns the MaxSignals field if non-nil, zero value otherwise.

### GetMaxSignalsOk

`func (o *EsqlRulePatchProps) GetMaxSignalsOk() (*int32, bool)`

GetMaxSignalsOk returns a tuple with the MaxSignals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSignals

`func (o *EsqlRulePatchProps) SetMaxSignals(v int32)`

SetMaxSignals sets MaxSignals field to given value.

### HasMaxSignals

`func (o *EsqlRulePatchProps) HasMaxSignals() bool`

HasMaxSignals returns a boolean if a field has been set.

### GetMeta

`func (o *EsqlRulePatchProps) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *EsqlRulePatchProps) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *EsqlRulePatchProps) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *EsqlRulePatchProps) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetName

`func (o *EsqlRulePatchProps) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EsqlRulePatchProps) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EsqlRulePatchProps) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EsqlRulePatchProps) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *EsqlRulePatchProps) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *EsqlRulePatchProps) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *EsqlRulePatchProps) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *EsqlRulePatchProps) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNote

`func (o *EsqlRulePatchProps) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *EsqlRulePatchProps) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *EsqlRulePatchProps) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *EsqlRulePatchProps) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetOutcome

`func (o *EsqlRulePatchProps) GetOutcome() SavedObjectResolveOutcome`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *EsqlRulePatchProps) GetOutcomeOk() (*SavedObjectResolveOutcome, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *EsqlRulePatchProps) SetOutcome(v SavedObjectResolveOutcome)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *EsqlRulePatchProps) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetOutputIndex

`func (o *EsqlRulePatchProps) GetOutputIndex() string`

GetOutputIndex returns the OutputIndex field if non-nil, zero value otherwise.

### GetOutputIndexOk

`func (o *EsqlRulePatchProps) GetOutputIndexOk() (*string, bool)`

GetOutputIndexOk returns a tuple with the OutputIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputIndex

`func (o *EsqlRulePatchProps) SetOutputIndex(v string)`

SetOutputIndex sets OutputIndex field to given value.

### HasOutputIndex

`func (o *EsqlRulePatchProps) HasOutputIndex() bool`

HasOutputIndex returns a boolean if a field has been set.

### GetQuery

`func (o *EsqlRulePatchProps) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *EsqlRulePatchProps) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *EsqlRulePatchProps) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *EsqlRulePatchProps) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetReferences

`func (o *EsqlRulePatchProps) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *EsqlRulePatchProps) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *EsqlRulePatchProps) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *EsqlRulePatchProps) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetRelatedIntegrations

`func (o *EsqlRulePatchProps) GetRelatedIntegrations() []RelatedIntegration`

GetRelatedIntegrations returns the RelatedIntegrations field if non-nil, zero value otherwise.

### GetRelatedIntegrationsOk

`func (o *EsqlRulePatchProps) GetRelatedIntegrationsOk() (*[]RelatedIntegration, bool)`

GetRelatedIntegrationsOk returns a tuple with the RelatedIntegrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedIntegrations

`func (o *EsqlRulePatchProps) SetRelatedIntegrations(v []RelatedIntegration)`

SetRelatedIntegrations sets RelatedIntegrations field to given value.

### HasRelatedIntegrations

`func (o *EsqlRulePatchProps) HasRelatedIntegrations() bool`

HasRelatedIntegrations returns a boolean if a field has been set.

### GetRequiredFields

`func (o *EsqlRulePatchProps) GetRequiredFields() []RequiredFieldInput`

GetRequiredFields returns the RequiredFields field if non-nil, zero value otherwise.

### GetRequiredFieldsOk

`func (o *EsqlRulePatchProps) GetRequiredFieldsOk() (*[]RequiredFieldInput, bool)`

GetRequiredFieldsOk returns a tuple with the RequiredFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredFields

`func (o *EsqlRulePatchProps) SetRequiredFields(v []RequiredFieldInput)`

SetRequiredFields sets RequiredFields field to given value.

### HasRequiredFields

`func (o *EsqlRulePatchProps) HasRequiredFields() bool`

HasRequiredFields returns a boolean if a field has been set.

### GetRiskScore

`func (o *EsqlRulePatchProps) GetRiskScore() int32`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *EsqlRulePatchProps) GetRiskScoreOk() (*int32, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *EsqlRulePatchProps) SetRiskScore(v int32)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *EsqlRulePatchProps) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### GetRiskScoreMapping

`func (o *EsqlRulePatchProps) GetRiskScoreMapping() []RiskScoreMappingInner`

GetRiskScoreMapping returns the RiskScoreMapping field if non-nil, zero value otherwise.

### GetRiskScoreMappingOk

`func (o *EsqlRulePatchProps) GetRiskScoreMappingOk() (*[]RiskScoreMappingInner, bool)`

GetRiskScoreMappingOk returns a tuple with the RiskScoreMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScoreMapping

`func (o *EsqlRulePatchProps) SetRiskScoreMapping(v []RiskScoreMappingInner)`

SetRiskScoreMapping sets RiskScoreMapping field to given value.

### HasRiskScoreMapping

`func (o *EsqlRulePatchProps) HasRiskScoreMapping() bool`

HasRiskScoreMapping returns a boolean if a field has been set.

### GetRuleId

`func (o *EsqlRulePatchProps) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *EsqlRulePatchProps) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *EsqlRulePatchProps) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *EsqlRulePatchProps) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetRuleNameOverride

`func (o *EsqlRulePatchProps) GetRuleNameOverride() string`

GetRuleNameOverride returns the RuleNameOverride field if non-nil, zero value otherwise.

### GetRuleNameOverrideOk

`func (o *EsqlRulePatchProps) GetRuleNameOverrideOk() (*string, bool)`

GetRuleNameOverrideOk returns a tuple with the RuleNameOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleNameOverride

`func (o *EsqlRulePatchProps) SetRuleNameOverride(v string)`

SetRuleNameOverride sets RuleNameOverride field to given value.

### HasRuleNameOverride

`func (o *EsqlRulePatchProps) HasRuleNameOverride() bool`

HasRuleNameOverride returns a boolean if a field has been set.

### GetSetup

`func (o *EsqlRulePatchProps) GetSetup() string`

GetSetup returns the Setup field if non-nil, zero value otherwise.

### GetSetupOk

`func (o *EsqlRulePatchProps) GetSetupOk() (*string, bool)`

GetSetupOk returns a tuple with the Setup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetup

`func (o *EsqlRulePatchProps) SetSetup(v string)`

SetSetup sets Setup field to given value.

### HasSetup

`func (o *EsqlRulePatchProps) HasSetup() bool`

HasSetup returns a boolean if a field has been set.

### GetSeverity

`func (o *EsqlRulePatchProps) GetSeverity() Severity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *EsqlRulePatchProps) GetSeverityOk() (*Severity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *EsqlRulePatchProps) SetSeverity(v Severity)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *EsqlRulePatchProps) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSeverityMapping

`func (o *EsqlRulePatchProps) GetSeverityMapping() []SeverityMappingInner`

GetSeverityMapping returns the SeverityMapping field if non-nil, zero value otherwise.

### GetSeverityMappingOk

`func (o *EsqlRulePatchProps) GetSeverityMappingOk() (*[]SeverityMappingInner, bool)`

GetSeverityMappingOk returns a tuple with the SeverityMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityMapping

`func (o *EsqlRulePatchProps) SetSeverityMapping(v []SeverityMappingInner)`

SetSeverityMapping sets SeverityMapping field to given value.

### HasSeverityMapping

`func (o *EsqlRulePatchProps) HasSeverityMapping() bool`

HasSeverityMapping returns a boolean if a field has been set.

### GetTags

`func (o *EsqlRulePatchProps) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EsqlRulePatchProps) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EsqlRulePatchProps) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EsqlRulePatchProps) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThreat

`func (o *EsqlRulePatchProps) GetThreat() []Threat`

GetThreat returns the Threat field if non-nil, zero value otherwise.

### GetThreatOk

`func (o *EsqlRulePatchProps) GetThreatOk() (*[]Threat, bool)`

GetThreatOk returns a tuple with the Threat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreat

`func (o *EsqlRulePatchProps) SetThreat(v []Threat)`

SetThreat sets Threat field to given value.

### HasThreat

`func (o *EsqlRulePatchProps) HasThreat() bool`

HasThreat returns a boolean if a field has been set.

### GetThrottle

`func (o *EsqlRulePatchProps) GetThrottle() RuleActionThrottle`

GetThrottle returns the Throttle field if non-nil, zero value otherwise.

### GetThrottleOk

`func (o *EsqlRulePatchProps) GetThrottleOk() (*RuleActionThrottle, bool)`

GetThrottleOk returns a tuple with the Throttle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottle

`func (o *EsqlRulePatchProps) SetThrottle(v RuleActionThrottle)`

SetThrottle sets Throttle field to given value.

### HasThrottle

`func (o *EsqlRulePatchProps) HasThrottle() bool`

HasThrottle returns a boolean if a field has been set.

### GetTimelineId

`func (o *EsqlRulePatchProps) GetTimelineId() string`

GetTimelineId returns the TimelineId field if non-nil, zero value otherwise.

### GetTimelineIdOk

`func (o *EsqlRulePatchProps) GetTimelineIdOk() (*string, bool)`

GetTimelineIdOk returns a tuple with the TimelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineId

`func (o *EsqlRulePatchProps) SetTimelineId(v string)`

SetTimelineId sets TimelineId field to given value.

### HasTimelineId

`func (o *EsqlRulePatchProps) HasTimelineId() bool`

HasTimelineId returns a boolean if a field has been set.

### GetTimelineTitle

`func (o *EsqlRulePatchProps) GetTimelineTitle() string`

GetTimelineTitle returns the TimelineTitle field if non-nil, zero value otherwise.

### GetTimelineTitleOk

`func (o *EsqlRulePatchProps) GetTimelineTitleOk() (*string, bool)`

GetTimelineTitleOk returns a tuple with the TimelineTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineTitle

`func (o *EsqlRulePatchProps) SetTimelineTitle(v string)`

SetTimelineTitle sets TimelineTitle field to given value.

### HasTimelineTitle

`func (o *EsqlRulePatchProps) HasTimelineTitle() bool`

HasTimelineTitle returns a boolean if a field has been set.

### GetTimestampOverride

`func (o *EsqlRulePatchProps) GetTimestampOverride() string`

GetTimestampOverride returns the TimestampOverride field if non-nil, zero value otherwise.

### GetTimestampOverrideOk

`func (o *EsqlRulePatchProps) GetTimestampOverrideOk() (*string, bool)`

GetTimestampOverrideOk returns a tuple with the TimestampOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampOverride

`func (o *EsqlRulePatchProps) SetTimestampOverride(v string)`

SetTimestampOverride sets TimestampOverride field to given value.

### HasTimestampOverride

`func (o *EsqlRulePatchProps) HasTimestampOverride() bool`

HasTimestampOverride returns a boolean if a field has been set.

### GetTimestampOverrideFallbackDisabled

`func (o *EsqlRulePatchProps) GetTimestampOverrideFallbackDisabled() bool`

GetTimestampOverrideFallbackDisabled returns the TimestampOverrideFallbackDisabled field if non-nil, zero value otherwise.

### GetTimestampOverrideFallbackDisabledOk

`func (o *EsqlRulePatchProps) GetTimestampOverrideFallbackDisabledOk() (*bool, bool)`

GetTimestampOverrideFallbackDisabledOk returns a tuple with the TimestampOverrideFallbackDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampOverrideFallbackDisabled

`func (o *EsqlRulePatchProps) SetTimestampOverrideFallbackDisabled(v bool)`

SetTimestampOverrideFallbackDisabled sets TimestampOverrideFallbackDisabled field to given value.

### HasTimestampOverrideFallbackDisabled

`func (o *EsqlRulePatchProps) HasTimestampOverrideFallbackDisabled() bool`

HasTimestampOverrideFallbackDisabled returns a boolean if a field has been set.

### GetTo

`func (o *EsqlRulePatchProps) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EsqlRulePatchProps) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EsqlRulePatchProps) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *EsqlRulePatchProps) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetType

`func (o *EsqlRulePatchProps) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EsqlRulePatchProps) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EsqlRulePatchProps) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EsqlRulePatchProps) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *EsqlRulePatchProps) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EsqlRulePatchProps) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EsqlRulePatchProps) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EsqlRulePatchProps) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


