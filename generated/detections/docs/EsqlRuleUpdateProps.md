# EsqlRuleUpdateProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertSuppression** | Pointer to [**AlertSuppression**](AlertSuppression.md) |  | [optional] 
**ResponseActions** | Pointer to [**[]ResponseAction**](ResponseAction.md) |  | [optional] 
**Language** | [**EsqlQueryLanguage**](EsqlQueryLanguage.md) |  | 
**Query** | **string** |  | 
**Type** | **string** | Rule type | 
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

## Methods

### NewEsqlRuleUpdateProps

`func NewEsqlRuleUpdateProps(language EsqlQueryLanguage, query string, type_ string, description string, name string, riskScore int32, severity Severity, ) *EsqlRuleUpdateProps`

NewEsqlRuleUpdateProps instantiates a new EsqlRuleUpdateProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsqlRuleUpdatePropsWithDefaults

`func NewEsqlRuleUpdatePropsWithDefaults() *EsqlRuleUpdateProps`

NewEsqlRuleUpdatePropsWithDefaults instantiates a new EsqlRuleUpdateProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertSuppression

`func (o *EsqlRuleUpdateProps) GetAlertSuppression() AlertSuppression`

GetAlertSuppression returns the AlertSuppression field if non-nil, zero value otherwise.

### GetAlertSuppressionOk

`func (o *EsqlRuleUpdateProps) GetAlertSuppressionOk() (*AlertSuppression, bool)`

GetAlertSuppressionOk returns a tuple with the AlertSuppression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertSuppression

`func (o *EsqlRuleUpdateProps) SetAlertSuppression(v AlertSuppression)`

SetAlertSuppression sets AlertSuppression field to given value.

### HasAlertSuppression

`func (o *EsqlRuleUpdateProps) HasAlertSuppression() bool`

HasAlertSuppression returns a boolean if a field has been set.

### GetResponseActions

`func (o *EsqlRuleUpdateProps) GetResponseActions() []ResponseAction`

GetResponseActions returns the ResponseActions field if non-nil, zero value otherwise.

### GetResponseActionsOk

`func (o *EsqlRuleUpdateProps) GetResponseActionsOk() (*[]ResponseAction, bool)`

GetResponseActionsOk returns a tuple with the ResponseActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseActions

`func (o *EsqlRuleUpdateProps) SetResponseActions(v []ResponseAction)`

SetResponseActions sets ResponseActions field to given value.

### HasResponseActions

`func (o *EsqlRuleUpdateProps) HasResponseActions() bool`

HasResponseActions returns a boolean if a field has been set.

### GetLanguage

`func (o *EsqlRuleUpdateProps) GetLanguage() EsqlQueryLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *EsqlRuleUpdateProps) GetLanguageOk() (*EsqlQueryLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *EsqlRuleUpdateProps) SetLanguage(v EsqlQueryLanguage)`

SetLanguage sets Language field to given value.


### GetQuery

`func (o *EsqlRuleUpdateProps) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *EsqlRuleUpdateProps) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *EsqlRuleUpdateProps) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetType

`func (o *EsqlRuleUpdateProps) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EsqlRuleUpdateProps) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EsqlRuleUpdateProps) SetType(v string)`

SetType sets Type field to given value.


### GetActions

`func (o *EsqlRuleUpdateProps) GetActions() []RuleAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *EsqlRuleUpdateProps) GetActionsOk() (*[]RuleAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *EsqlRuleUpdateProps) SetActions(v []RuleAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *EsqlRuleUpdateProps) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetAliasPurpose

`func (o *EsqlRuleUpdateProps) GetAliasPurpose() SavedObjectResolveAliasPurpose`

GetAliasPurpose returns the AliasPurpose field if non-nil, zero value otherwise.

### GetAliasPurposeOk

`func (o *EsqlRuleUpdateProps) GetAliasPurposeOk() (*SavedObjectResolveAliasPurpose, bool)`

GetAliasPurposeOk returns a tuple with the AliasPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasPurpose

`func (o *EsqlRuleUpdateProps) SetAliasPurpose(v SavedObjectResolveAliasPurpose)`

SetAliasPurpose sets AliasPurpose field to given value.

### HasAliasPurpose

`func (o *EsqlRuleUpdateProps) HasAliasPurpose() bool`

HasAliasPurpose returns a boolean if a field has been set.

### GetAliasTargetId

`func (o *EsqlRuleUpdateProps) GetAliasTargetId() string`

GetAliasTargetId returns the AliasTargetId field if non-nil, zero value otherwise.

### GetAliasTargetIdOk

`func (o *EsqlRuleUpdateProps) GetAliasTargetIdOk() (*string, bool)`

GetAliasTargetIdOk returns a tuple with the AliasTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasTargetId

`func (o *EsqlRuleUpdateProps) SetAliasTargetId(v string)`

SetAliasTargetId sets AliasTargetId field to given value.

### HasAliasTargetId

`func (o *EsqlRuleUpdateProps) HasAliasTargetId() bool`

HasAliasTargetId returns a boolean if a field has been set.

### GetAuthor

`func (o *EsqlRuleUpdateProps) GetAuthor() []string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *EsqlRuleUpdateProps) GetAuthorOk() (*[]string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *EsqlRuleUpdateProps) SetAuthor(v []string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *EsqlRuleUpdateProps) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetBuildingBlockType

`func (o *EsqlRuleUpdateProps) GetBuildingBlockType() string`

GetBuildingBlockType returns the BuildingBlockType field if non-nil, zero value otherwise.

### GetBuildingBlockTypeOk

`func (o *EsqlRuleUpdateProps) GetBuildingBlockTypeOk() (*string, bool)`

GetBuildingBlockTypeOk returns a tuple with the BuildingBlockType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildingBlockType

`func (o *EsqlRuleUpdateProps) SetBuildingBlockType(v string)`

SetBuildingBlockType sets BuildingBlockType field to given value.

### HasBuildingBlockType

`func (o *EsqlRuleUpdateProps) HasBuildingBlockType() bool`

HasBuildingBlockType returns a boolean if a field has been set.

### GetDescription

`func (o *EsqlRuleUpdateProps) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EsqlRuleUpdateProps) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EsqlRuleUpdateProps) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEnabled

`func (o *EsqlRuleUpdateProps) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EsqlRuleUpdateProps) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EsqlRuleUpdateProps) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *EsqlRuleUpdateProps) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExceptionsList

`func (o *EsqlRuleUpdateProps) GetExceptionsList() []RuleExceptionList`

GetExceptionsList returns the ExceptionsList field if non-nil, zero value otherwise.

### GetExceptionsListOk

`func (o *EsqlRuleUpdateProps) GetExceptionsListOk() (*[]RuleExceptionList, bool)`

GetExceptionsListOk returns a tuple with the ExceptionsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionsList

`func (o *EsqlRuleUpdateProps) SetExceptionsList(v []RuleExceptionList)`

SetExceptionsList sets ExceptionsList field to given value.

### HasExceptionsList

`func (o *EsqlRuleUpdateProps) HasExceptionsList() bool`

HasExceptionsList returns a boolean if a field has been set.

### GetFalsePositives

`func (o *EsqlRuleUpdateProps) GetFalsePositives() []string`

GetFalsePositives returns the FalsePositives field if non-nil, zero value otherwise.

### GetFalsePositivesOk

`func (o *EsqlRuleUpdateProps) GetFalsePositivesOk() (*[]string, bool)`

GetFalsePositivesOk returns a tuple with the FalsePositives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFalsePositives

`func (o *EsqlRuleUpdateProps) SetFalsePositives(v []string)`

SetFalsePositives sets FalsePositives field to given value.

### HasFalsePositives

`func (o *EsqlRuleUpdateProps) HasFalsePositives() bool`

HasFalsePositives returns a boolean if a field has been set.

### GetFrom

`func (o *EsqlRuleUpdateProps) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *EsqlRuleUpdateProps) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *EsqlRuleUpdateProps) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *EsqlRuleUpdateProps) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetId

`func (o *EsqlRuleUpdateProps) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EsqlRuleUpdateProps) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EsqlRuleUpdateProps) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EsqlRuleUpdateProps) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterval

`func (o *EsqlRuleUpdateProps) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *EsqlRuleUpdateProps) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *EsqlRuleUpdateProps) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *EsqlRuleUpdateProps) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetInvestigationFields

`func (o *EsqlRuleUpdateProps) GetInvestigationFields() InvestigationFields`

GetInvestigationFields returns the InvestigationFields field if non-nil, zero value otherwise.

### GetInvestigationFieldsOk

`func (o *EsqlRuleUpdateProps) GetInvestigationFieldsOk() (*InvestigationFields, bool)`

GetInvestigationFieldsOk returns a tuple with the InvestigationFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigationFields

`func (o *EsqlRuleUpdateProps) SetInvestigationFields(v InvestigationFields)`

SetInvestigationFields sets InvestigationFields field to given value.

### HasInvestigationFields

`func (o *EsqlRuleUpdateProps) HasInvestigationFields() bool`

HasInvestigationFields returns a boolean if a field has been set.

### GetLicense

`func (o *EsqlRuleUpdateProps) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *EsqlRuleUpdateProps) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *EsqlRuleUpdateProps) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *EsqlRuleUpdateProps) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetMaxSignals

`func (o *EsqlRuleUpdateProps) GetMaxSignals() int32`

GetMaxSignals returns the MaxSignals field if non-nil, zero value otherwise.

### GetMaxSignalsOk

`func (o *EsqlRuleUpdateProps) GetMaxSignalsOk() (*int32, bool)`

GetMaxSignalsOk returns a tuple with the MaxSignals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSignals

`func (o *EsqlRuleUpdateProps) SetMaxSignals(v int32)`

SetMaxSignals sets MaxSignals field to given value.

### HasMaxSignals

`func (o *EsqlRuleUpdateProps) HasMaxSignals() bool`

HasMaxSignals returns a boolean if a field has been set.

### GetMeta

`func (o *EsqlRuleUpdateProps) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *EsqlRuleUpdateProps) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *EsqlRuleUpdateProps) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *EsqlRuleUpdateProps) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetName

`func (o *EsqlRuleUpdateProps) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EsqlRuleUpdateProps) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EsqlRuleUpdateProps) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *EsqlRuleUpdateProps) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *EsqlRuleUpdateProps) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *EsqlRuleUpdateProps) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *EsqlRuleUpdateProps) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNote

`func (o *EsqlRuleUpdateProps) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *EsqlRuleUpdateProps) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *EsqlRuleUpdateProps) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *EsqlRuleUpdateProps) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetOutcome

`func (o *EsqlRuleUpdateProps) GetOutcome() SavedObjectResolveOutcome`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *EsqlRuleUpdateProps) GetOutcomeOk() (*SavedObjectResolveOutcome, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *EsqlRuleUpdateProps) SetOutcome(v SavedObjectResolveOutcome)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *EsqlRuleUpdateProps) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetOutputIndex

`func (o *EsqlRuleUpdateProps) GetOutputIndex() string`

GetOutputIndex returns the OutputIndex field if non-nil, zero value otherwise.

### GetOutputIndexOk

`func (o *EsqlRuleUpdateProps) GetOutputIndexOk() (*string, bool)`

GetOutputIndexOk returns a tuple with the OutputIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputIndex

`func (o *EsqlRuleUpdateProps) SetOutputIndex(v string)`

SetOutputIndex sets OutputIndex field to given value.

### HasOutputIndex

`func (o *EsqlRuleUpdateProps) HasOutputIndex() bool`

HasOutputIndex returns a boolean if a field has been set.

### GetReferences

`func (o *EsqlRuleUpdateProps) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *EsqlRuleUpdateProps) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *EsqlRuleUpdateProps) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *EsqlRuleUpdateProps) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetRelatedIntegrations

`func (o *EsqlRuleUpdateProps) GetRelatedIntegrations() []RelatedIntegration`

GetRelatedIntegrations returns the RelatedIntegrations field if non-nil, zero value otherwise.

### GetRelatedIntegrationsOk

`func (o *EsqlRuleUpdateProps) GetRelatedIntegrationsOk() (*[]RelatedIntegration, bool)`

GetRelatedIntegrationsOk returns a tuple with the RelatedIntegrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedIntegrations

`func (o *EsqlRuleUpdateProps) SetRelatedIntegrations(v []RelatedIntegration)`

SetRelatedIntegrations sets RelatedIntegrations field to given value.

### HasRelatedIntegrations

`func (o *EsqlRuleUpdateProps) HasRelatedIntegrations() bool`

HasRelatedIntegrations returns a boolean if a field has been set.

### GetRequiredFields

`func (o *EsqlRuleUpdateProps) GetRequiredFields() []RequiredFieldInput`

GetRequiredFields returns the RequiredFields field if non-nil, zero value otherwise.

### GetRequiredFieldsOk

`func (o *EsqlRuleUpdateProps) GetRequiredFieldsOk() (*[]RequiredFieldInput, bool)`

GetRequiredFieldsOk returns a tuple with the RequiredFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredFields

`func (o *EsqlRuleUpdateProps) SetRequiredFields(v []RequiredFieldInput)`

SetRequiredFields sets RequiredFields field to given value.

### HasRequiredFields

`func (o *EsqlRuleUpdateProps) HasRequiredFields() bool`

HasRequiredFields returns a boolean if a field has been set.

### GetRiskScore

`func (o *EsqlRuleUpdateProps) GetRiskScore() int32`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *EsqlRuleUpdateProps) GetRiskScoreOk() (*int32, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *EsqlRuleUpdateProps) SetRiskScore(v int32)`

SetRiskScore sets RiskScore field to given value.


### GetRiskScoreMapping

`func (o *EsqlRuleUpdateProps) GetRiskScoreMapping() []RiskScoreMappingInner`

GetRiskScoreMapping returns the RiskScoreMapping field if non-nil, zero value otherwise.

### GetRiskScoreMappingOk

`func (o *EsqlRuleUpdateProps) GetRiskScoreMappingOk() (*[]RiskScoreMappingInner, bool)`

GetRiskScoreMappingOk returns a tuple with the RiskScoreMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScoreMapping

`func (o *EsqlRuleUpdateProps) SetRiskScoreMapping(v []RiskScoreMappingInner)`

SetRiskScoreMapping sets RiskScoreMapping field to given value.

### HasRiskScoreMapping

`func (o *EsqlRuleUpdateProps) HasRiskScoreMapping() bool`

HasRiskScoreMapping returns a boolean if a field has been set.

### GetRuleId

`func (o *EsqlRuleUpdateProps) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *EsqlRuleUpdateProps) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *EsqlRuleUpdateProps) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *EsqlRuleUpdateProps) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetRuleNameOverride

`func (o *EsqlRuleUpdateProps) GetRuleNameOverride() string`

GetRuleNameOverride returns the RuleNameOverride field if non-nil, zero value otherwise.

### GetRuleNameOverrideOk

`func (o *EsqlRuleUpdateProps) GetRuleNameOverrideOk() (*string, bool)`

GetRuleNameOverrideOk returns a tuple with the RuleNameOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleNameOverride

`func (o *EsqlRuleUpdateProps) SetRuleNameOverride(v string)`

SetRuleNameOverride sets RuleNameOverride field to given value.

### HasRuleNameOverride

`func (o *EsqlRuleUpdateProps) HasRuleNameOverride() bool`

HasRuleNameOverride returns a boolean if a field has been set.

### GetSetup

`func (o *EsqlRuleUpdateProps) GetSetup() string`

GetSetup returns the Setup field if non-nil, zero value otherwise.

### GetSetupOk

`func (o *EsqlRuleUpdateProps) GetSetupOk() (*string, bool)`

GetSetupOk returns a tuple with the Setup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetup

`func (o *EsqlRuleUpdateProps) SetSetup(v string)`

SetSetup sets Setup field to given value.

### HasSetup

`func (o *EsqlRuleUpdateProps) HasSetup() bool`

HasSetup returns a boolean if a field has been set.

### GetSeverity

`func (o *EsqlRuleUpdateProps) GetSeverity() Severity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *EsqlRuleUpdateProps) GetSeverityOk() (*Severity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *EsqlRuleUpdateProps) SetSeverity(v Severity)`

SetSeverity sets Severity field to given value.


### GetSeverityMapping

`func (o *EsqlRuleUpdateProps) GetSeverityMapping() []SeverityMappingInner`

GetSeverityMapping returns the SeverityMapping field if non-nil, zero value otherwise.

### GetSeverityMappingOk

`func (o *EsqlRuleUpdateProps) GetSeverityMappingOk() (*[]SeverityMappingInner, bool)`

GetSeverityMappingOk returns a tuple with the SeverityMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityMapping

`func (o *EsqlRuleUpdateProps) SetSeverityMapping(v []SeverityMappingInner)`

SetSeverityMapping sets SeverityMapping field to given value.

### HasSeverityMapping

`func (o *EsqlRuleUpdateProps) HasSeverityMapping() bool`

HasSeverityMapping returns a boolean if a field has been set.

### GetTags

`func (o *EsqlRuleUpdateProps) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EsqlRuleUpdateProps) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EsqlRuleUpdateProps) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EsqlRuleUpdateProps) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThreat

`func (o *EsqlRuleUpdateProps) GetThreat() []Threat`

GetThreat returns the Threat field if non-nil, zero value otherwise.

### GetThreatOk

`func (o *EsqlRuleUpdateProps) GetThreatOk() (*[]Threat, bool)`

GetThreatOk returns a tuple with the Threat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreat

`func (o *EsqlRuleUpdateProps) SetThreat(v []Threat)`

SetThreat sets Threat field to given value.

### HasThreat

`func (o *EsqlRuleUpdateProps) HasThreat() bool`

HasThreat returns a boolean if a field has been set.

### GetThrottle

`func (o *EsqlRuleUpdateProps) GetThrottle() RuleActionThrottle`

GetThrottle returns the Throttle field if non-nil, zero value otherwise.

### GetThrottleOk

`func (o *EsqlRuleUpdateProps) GetThrottleOk() (*RuleActionThrottle, bool)`

GetThrottleOk returns a tuple with the Throttle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottle

`func (o *EsqlRuleUpdateProps) SetThrottle(v RuleActionThrottle)`

SetThrottle sets Throttle field to given value.

### HasThrottle

`func (o *EsqlRuleUpdateProps) HasThrottle() bool`

HasThrottle returns a boolean if a field has been set.

### GetTimelineId

`func (o *EsqlRuleUpdateProps) GetTimelineId() string`

GetTimelineId returns the TimelineId field if non-nil, zero value otherwise.

### GetTimelineIdOk

`func (o *EsqlRuleUpdateProps) GetTimelineIdOk() (*string, bool)`

GetTimelineIdOk returns a tuple with the TimelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineId

`func (o *EsqlRuleUpdateProps) SetTimelineId(v string)`

SetTimelineId sets TimelineId field to given value.

### HasTimelineId

`func (o *EsqlRuleUpdateProps) HasTimelineId() bool`

HasTimelineId returns a boolean if a field has been set.

### GetTimelineTitle

`func (o *EsqlRuleUpdateProps) GetTimelineTitle() string`

GetTimelineTitle returns the TimelineTitle field if non-nil, zero value otherwise.

### GetTimelineTitleOk

`func (o *EsqlRuleUpdateProps) GetTimelineTitleOk() (*string, bool)`

GetTimelineTitleOk returns a tuple with the TimelineTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineTitle

`func (o *EsqlRuleUpdateProps) SetTimelineTitle(v string)`

SetTimelineTitle sets TimelineTitle field to given value.

### HasTimelineTitle

`func (o *EsqlRuleUpdateProps) HasTimelineTitle() bool`

HasTimelineTitle returns a boolean if a field has been set.

### GetTimestampOverride

`func (o *EsqlRuleUpdateProps) GetTimestampOverride() string`

GetTimestampOverride returns the TimestampOverride field if non-nil, zero value otherwise.

### GetTimestampOverrideOk

`func (o *EsqlRuleUpdateProps) GetTimestampOverrideOk() (*string, bool)`

GetTimestampOverrideOk returns a tuple with the TimestampOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampOverride

`func (o *EsqlRuleUpdateProps) SetTimestampOverride(v string)`

SetTimestampOverride sets TimestampOverride field to given value.

### HasTimestampOverride

`func (o *EsqlRuleUpdateProps) HasTimestampOverride() bool`

HasTimestampOverride returns a boolean if a field has been set.

### GetTimestampOverrideFallbackDisabled

`func (o *EsqlRuleUpdateProps) GetTimestampOverrideFallbackDisabled() bool`

GetTimestampOverrideFallbackDisabled returns the TimestampOverrideFallbackDisabled field if non-nil, zero value otherwise.

### GetTimestampOverrideFallbackDisabledOk

`func (o *EsqlRuleUpdateProps) GetTimestampOverrideFallbackDisabledOk() (*bool, bool)`

GetTimestampOverrideFallbackDisabledOk returns a tuple with the TimestampOverrideFallbackDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampOverrideFallbackDisabled

`func (o *EsqlRuleUpdateProps) SetTimestampOverrideFallbackDisabled(v bool)`

SetTimestampOverrideFallbackDisabled sets TimestampOverrideFallbackDisabled field to given value.

### HasTimestampOverrideFallbackDisabled

`func (o *EsqlRuleUpdateProps) HasTimestampOverrideFallbackDisabled() bool`

HasTimestampOverrideFallbackDisabled returns a boolean if a field has been set.

### GetTo

`func (o *EsqlRuleUpdateProps) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EsqlRuleUpdateProps) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EsqlRuleUpdateProps) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *EsqlRuleUpdateProps) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetVersion

`func (o *EsqlRuleUpdateProps) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EsqlRuleUpdateProps) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EsqlRuleUpdateProps) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EsqlRuleUpdateProps) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


