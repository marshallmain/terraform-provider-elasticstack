# SavedQueryRuleCreateProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SavedId** | **string** |  | 
**Type** | **string** | Rule type | 
**AlertSuppression** | Pointer to [**AlertSuppression**](AlertSuppression.md) |  | [optional] 
**DataViewId** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to **[]interface{}** |  | [optional] 
**Index** | Pointer to **[]string** |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**ResponseActions** | Pointer to [**[]ResponseAction**](ResponseAction.md) |  | [optional] 
**Language** | Pointer to [**KqlQueryLanguage**](KqlQueryLanguage.md) |  | [optional] 
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

## Methods

### NewSavedQueryRuleCreateProps

`func NewSavedQueryRuleCreateProps(savedId string, type_ string, description string, name string, riskScore int32, severity Severity, ) *SavedQueryRuleCreateProps`

NewSavedQueryRuleCreateProps instantiates a new SavedQueryRuleCreateProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavedQueryRuleCreatePropsWithDefaults

`func NewSavedQueryRuleCreatePropsWithDefaults() *SavedQueryRuleCreateProps`

NewSavedQueryRuleCreatePropsWithDefaults instantiates a new SavedQueryRuleCreateProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSavedId

`func (o *SavedQueryRuleCreateProps) GetSavedId() string`

GetSavedId returns the SavedId field if non-nil, zero value otherwise.

### GetSavedIdOk

`func (o *SavedQueryRuleCreateProps) GetSavedIdOk() (*string, bool)`

GetSavedIdOk returns a tuple with the SavedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedId

`func (o *SavedQueryRuleCreateProps) SetSavedId(v string)`

SetSavedId sets SavedId field to given value.


### GetType

`func (o *SavedQueryRuleCreateProps) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SavedQueryRuleCreateProps) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SavedQueryRuleCreateProps) SetType(v string)`

SetType sets Type field to given value.


### GetAlertSuppression

`func (o *SavedQueryRuleCreateProps) GetAlertSuppression() AlertSuppression`

GetAlertSuppression returns the AlertSuppression field if non-nil, zero value otherwise.

### GetAlertSuppressionOk

`func (o *SavedQueryRuleCreateProps) GetAlertSuppressionOk() (*AlertSuppression, bool)`

GetAlertSuppressionOk returns a tuple with the AlertSuppression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertSuppression

`func (o *SavedQueryRuleCreateProps) SetAlertSuppression(v AlertSuppression)`

SetAlertSuppression sets AlertSuppression field to given value.

### HasAlertSuppression

`func (o *SavedQueryRuleCreateProps) HasAlertSuppression() bool`

HasAlertSuppression returns a boolean if a field has been set.

### GetDataViewId

`func (o *SavedQueryRuleCreateProps) GetDataViewId() string`

GetDataViewId returns the DataViewId field if non-nil, zero value otherwise.

### GetDataViewIdOk

`func (o *SavedQueryRuleCreateProps) GetDataViewIdOk() (*string, bool)`

GetDataViewIdOk returns a tuple with the DataViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataViewId

`func (o *SavedQueryRuleCreateProps) SetDataViewId(v string)`

SetDataViewId sets DataViewId field to given value.

### HasDataViewId

`func (o *SavedQueryRuleCreateProps) HasDataViewId() bool`

HasDataViewId returns a boolean if a field has been set.

### GetFilters

`func (o *SavedQueryRuleCreateProps) GetFilters() []interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *SavedQueryRuleCreateProps) GetFiltersOk() (*[]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *SavedQueryRuleCreateProps) SetFilters(v []interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *SavedQueryRuleCreateProps) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetIndex

`func (o *SavedQueryRuleCreateProps) GetIndex() []string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *SavedQueryRuleCreateProps) GetIndexOk() (*[]string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *SavedQueryRuleCreateProps) SetIndex(v []string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *SavedQueryRuleCreateProps) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetQuery

`func (o *SavedQueryRuleCreateProps) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SavedQueryRuleCreateProps) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SavedQueryRuleCreateProps) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *SavedQueryRuleCreateProps) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetResponseActions

`func (o *SavedQueryRuleCreateProps) GetResponseActions() []ResponseAction`

GetResponseActions returns the ResponseActions field if non-nil, zero value otherwise.

### GetResponseActionsOk

`func (o *SavedQueryRuleCreateProps) GetResponseActionsOk() (*[]ResponseAction, bool)`

GetResponseActionsOk returns a tuple with the ResponseActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseActions

`func (o *SavedQueryRuleCreateProps) SetResponseActions(v []ResponseAction)`

SetResponseActions sets ResponseActions field to given value.

### HasResponseActions

`func (o *SavedQueryRuleCreateProps) HasResponseActions() bool`

HasResponseActions returns a boolean if a field has been set.

### GetLanguage

`func (o *SavedQueryRuleCreateProps) GetLanguage() KqlQueryLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *SavedQueryRuleCreateProps) GetLanguageOk() (*KqlQueryLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *SavedQueryRuleCreateProps) SetLanguage(v KqlQueryLanguage)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *SavedQueryRuleCreateProps) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetActions

`func (o *SavedQueryRuleCreateProps) GetActions() []RuleAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *SavedQueryRuleCreateProps) GetActionsOk() (*[]RuleAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *SavedQueryRuleCreateProps) SetActions(v []RuleAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *SavedQueryRuleCreateProps) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetAliasPurpose

`func (o *SavedQueryRuleCreateProps) GetAliasPurpose() SavedObjectResolveAliasPurpose`

GetAliasPurpose returns the AliasPurpose field if non-nil, zero value otherwise.

### GetAliasPurposeOk

`func (o *SavedQueryRuleCreateProps) GetAliasPurposeOk() (*SavedObjectResolveAliasPurpose, bool)`

GetAliasPurposeOk returns a tuple with the AliasPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasPurpose

`func (o *SavedQueryRuleCreateProps) SetAliasPurpose(v SavedObjectResolveAliasPurpose)`

SetAliasPurpose sets AliasPurpose field to given value.

### HasAliasPurpose

`func (o *SavedQueryRuleCreateProps) HasAliasPurpose() bool`

HasAliasPurpose returns a boolean if a field has been set.

### GetAliasTargetId

`func (o *SavedQueryRuleCreateProps) GetAliasTargetId() string`

GetAliasTargetId returns the AliasTargetId field if non-nil, zero value otherwise.

### GetAliasTargetIdOk

`func (o *SavedQueryRuleCreateProps) GetAliasTargetIdOk() (*string, bool)`

GetAliasTargetIdOk returns a tuple with the AliasTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasTargetId

`func (o *SavedQueryRuleCreateProps) SetAliasTargetId(v string)`

SetAliasTargetId sets AliasTargetId field to given value.

### HasAliasTargetId

`func (o *SavedQueryRuleCreateProps) HasAliasTargetId() bool`

HasAliasTargetId returns a boolean if a field has been set.

### GetAuthor

`func (o *SavedQueryRuleCreateProps) GetAuthor() []string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *SavedQueryRuleCreateProps) GetAuthorOk() (*[]string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *SavedQueryRuleCreateProps) SetAuthor(v []string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *SavedQueryRuleCreateProps) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetBuildingBlockType

`func (o *SavedQueryRuleCreateProps) GetBuildingBlockType() string`

GetBuildingBlockType returns the BuildingBlockType field if non-nil, zero value otherwise.

### GetBuildingBlockTypeOk

`func (o *SavedQueryRuleCreateProps) GetBuildingBlockTypeOk() (*string, bool)`

GetBuildingBlockTypeOk returns a tuple with the BuildingBlockType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildingBlockType

`func (o *SavedQueryRuleCreateProps) SetBuildingBlockType(v string)`

SetBuildingBlockType sets BuildingBlockType field to given value.

### HasBuildingBlockType

`func (o *SavedQueryRuleCreateProps) HasBuildingBlockType() bool`

HasBuildingBlockType returns a boolean if a field has been set.

### GetDescription

`func (o *SavedQueryRuleCreateProps) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SavedQueryRuleCreateProps) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SavedQueryRuleCreateProps) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEnabled

`func (o *SavedQueryRuleCreateProps) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SavedQueryRuleCreateProps) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SavedQueryRuleCreateProps) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SavedQueryRuleCreateProps) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExceptionsList

`func (o *SavedQueryRuleCreateProps) GetExceptionsList() []RuleExceptionList`

GetExceptionsList returns the ExceptionsList field if non-nil, zero value otherwise.

### GetExceptionsListOk

`func (o *SavedQueryRuleCreateProps) GetExceptionsListOk() (*[]RuleExceptionList, bool)`

GetExceptionsListOk returns a tuple with the ExceptionsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionsList

`func (o *SavedQueryRuleCreateProps) SetExceptionsList(v []RuleExceptionList)`

SetExceptionsList sets ExceptionsList field to given value.

### HasExceptionsList

`func (o *SavedQueryRuleCreateProps) HasExceptionsList() bool`

HasExceptionsList returns a boolean if a field has been set.

### GetFalsePositives

`func (o *SavedQueryRuleCreateProps) GetFalsePositives() []string`

GetFalsePositives returns the FalsePositives field if non-nil, zero value otherwise.

### GetFalsePositivesOk

`func (o *SavedQueryRuleCreateProps) GetFalsePositivesOk() (*[]string, bool)`

GetFalsePositivesOk returns a tuple with the FalsePositives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFalsePositives

`func (o *SavedQueryRuleCreateProps) SetFalsePositives(v []string)`

SetFalsePositives sets FalsePositives field to given value.

### HasFalsePositives

`func (o *SavedQueryRuleCreateProps) HasFalsePositives() bool`

HasFalsePositives returns a boolean if a field has been set.

### GetFrom

`func (o *SavedQueryRuleCreateProps) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SavedQueryRuleCreateProps) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SavedQueryRuleCreateProps) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *SavedQueryRuleCreateProps) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetInterval

`func (o *SavedQueryRuleCreateProps) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *SavedQueryRuleCreateProps) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *SavedQueryRuleCreateProps) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *SavedQueryRuleCreateProps) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetInvestigationFields

`func (o *SavedQueryRuleCreateProps) GetInvestigationFields() InvestigationFields`

GetInvestigationFields returns the InvestigationFields field if non-nil, zero value otherwise.

### GetInvestigationFieldsOk

`func (o *SavedQueryRuleCreateProps) GetInvestigationFieldsOk() (*InvestigationFields, bool)`

GetInvestigationFieldsOk returns a tuple with the InvestigationFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigationFields

`func (o *SavedQueryRuleCreateProps) SetInvestigationFields(v InvestigationFields)`

SetInvestigationFields sets InvestigationFields field to given value.

### HasInvestigationFields

`func (o *SavedQueryRuleCreateProps) HasInvestigationFields() bool`

HasInvestigationFields returns a boolean if a field has been set.

### GetLicense

`func (o *SavedQueryRuleCreateProps) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *SavedQueryRuleCreateProps) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *SavedQueryRuleCreateProps) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *SavedQueryRuleCreateProps) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetMaxSignals

`func (o *SavedQueryRuleCreateProps) GetMaxSignals() int32`

GetMaxSignals returns the MaxSignals field if non-nil, zero value otherwise.

### GetMaxSignalsOk

`func (o *SavedQueryRuleCreateProps) GetMaxSignalsOk() (*int32, bool)`

GetMaxSignalsOk returns a tuple with the MaxSignals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSignals

`func (o *SavedQueryRuleCreateProps) SetMaxSignals(v int32)`

SetMaxSignals sets MaxSignals field to given value.

### HasMaxSignals

`func (o *SavedQueryRuleCreateProps) HasMaxSignals() bool`

HasMaxSignals returns a boolean if a field has been set.

### GetMeta

`func (o *SavedQueryRuleCreateProps) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SavedQueryRuleCreateProps) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SavedQueryRuleCreateProps) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SavedQueryRuleCreateProps) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetName

`func (o *SavedQueryRuleCreateProps) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SavedQueryRuleCreateProps) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SavedQueryRuleCreateProps) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *SavedQueryRuleCreateProps) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *SavedQueryRuleCreateProps) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *SavedQueryRuleCreateProps) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *SavedQueryRuleCreateProps) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNote

`func (o *SavedQueryRuleCreateProps) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *SavedQueryRuleCreateProps) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *SavedQueryRuleCreateProps) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *SavedQueryRuleCreateProps) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetOutcome

`func (o *SavedQueryRuleCreateProps) GetOutcome() SavedObjectResolveOutcome`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *SavedQueryRuleCreateProps) GetOutcomeOk() (*SavedObjectResolveOutcome, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *SavedQueryRuleCreateProps) SetOutcome(v SavedObjectResolveOutcome)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *SavedQueryRuleCreateProps) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetOutputIndex

`func (o *SavedQueryRuleCreateProps) GetOutputIndex() string`

GetOutputIndex returns the OutputIndex field if non-nil, zero value otherwise.

### GetOutputIndexOk

`func (o *SavedQueryRuleCreateProps) GetOutputIndexOk() (*string, bool)`

GetOutputIndexOk returns a tuple with the OutputIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputIndex

`func (o *SavedQueryRuleCreateProps) SetOutputIndex(v string)`

SetOutputIndex sets OutputIndex field to given value.

### HasOutputIndex

`func (o *SavedQueryRuleCreateProps) HasOutputIndex() bool`

HasOutputIndex returns a boolean if a field has been set.

### GetReferences

`func (o *SavedQueryRuleCreateProps) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *SavedQueryRuleCreateProps) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *SavedQueryRuleCreateProps) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *SavedQueryRuleCreateProps) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetRelatedIntegrations

`func (o *SavedQueryRuleCreateProps) GetRelatedIntegrations() []RelatedIntegration`

GetRelatedIntegrations returns the RelatedIntegrations field if non-nil, zero value otherwise.

### GetRelatedIntegrationsOk

`func (o *SavedQueryRuleCreateProps) GetRelatedIntegrationsOk() (*[]RelatedIntegration, bool)`

GetRelatedIntegrationsOk returns a tuple with the RelatedIntegrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedIntegrations

`func (o *SavedQueryRuleCreateProps) SetRelatedIntegrations(v []RelatedIntegration)`

SetRelatedIntegrations sets RelatedIntegrations field to given value.

### HasRelatedIntegrations

`func (o *SavedQueryRuleCreateProps) HasRelatedIntegrations() bool`

HasRelatedIntegrations returns a boolean if a field has been set.

### GetRequiredFields

`func (o *SavedQueryRuleCreateProps) GetRequiredFields() []RequiredFieldInput`

GetRequiredFields returns the RequiredFields field if non-nil, zero value otherwise.

### GetRequiredFieldsOk

`func (o *SavedQueryRuleCreateProps) GetRequiredFieldsOk() (*[]RequiredFieldInput, bool)`

GetRequiredFieldsOk returns a tuple with the RequiredFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredFields

`func (o *SavedQueryRuleCreateProps) SetRequiredFields(v []RequiredFieldInput)`

SetRequiredFields sets RequiredFields field to given value.

### HasRequiredFields

`func (o *SavedQueryRuleCreateProps) HasRequiredFields() bool`

HasRequiredFields returns a boolean if a field has been set.

### GetRiskScore

`func (o *SavedQueryRuleCreateProps) GetRiskScore() int32`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *SavedQueryRuleCreateProps) GetRiskScoreOk() (*int32, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *SavedQueryRuleCreateProps) SetRiskScore(v int32)`

SetRiskScore sets RiskScore field to given value.


### GetRiskScoreMapping

`func (o *SavedQueryRuleCreateProps) GetRiskScoreMapping() []RiskScoreMappingInner`

GetRiskScoreMapping returns the RiskScoreMapping field if non-nil, zero value otherwise.

### GetRiskScoreMappingOk

`func (o *SavedQueryRuleCreateProps) GetRiskScoreMappingOk() (*[]RiskScoreMappingInner, bool)`

GetRiskScoreMappingOk returns a tuple with the RiskScoreMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScoreMapping

`func (o *SavedQueryRuleCreateProps) SetRiskScoreMapping(v []RiskScoreMappingInner)`

SetRiskScoreMapping sets RiskScoreMapping field to given value.

### HasRiskScoreMapping

`func (o *SavedQueryRuleCreateProps) HasRiskScoreMapping() bool`

HasRiskScoreMapping returns a boolean if a field has been set.

### GetRuleId

`func (o *SavedQueryRuleCreateProps) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *SavedQueryRuleCreateProps) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *SavedQueryRuleCreateProps) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *SavedQueryRuleCreateProps) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetRuleNameOverride

`func (o *SavedQueryRuleCreateProps) GetRuleNameOverride() string`

GetRuleNameOverride returns the RuleNameOverride field if non-nil, zero value otherwise.

### GetRuleNameOverrideOk

`func (o *SavedQueryRuleCreateProps) GetRuleNameOverrideOk() (*string, bool)`

GetRuleNameOverrideOk returns a tuple with the RuleNameOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleNameOverride

`func (o *SavedQueryRuleCreateProps) SetRuleNameOverride(v string)`

SetRuleNameOverride sets RuleNameOverride field to given value.

### HasRuleNameOverride

`func (o *SavedQueryRuleCreateProps) HasRuleNameOverride() bool`

HasRuleNameOverride returns a boolean if a field has been set.

### GetSetup

`func (o *SavedQueryRuleCreateProps) GetSetup() string`

GetSetup returns the Setup field if non-nil, zero value otherwise.

### GetSetupOk

`func (o *SavedQueryRuleCreateProps) GetSetupOk() (*string, bool)`

GetSetupOk returns a tuple with the Setup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetup

`func (o *SavedQueryRuleCreateProps) SetSetup(v string)`

SetSetup sets Setup field to given value.

### HasSetup

`func (o *SavedQueryRuleCreateProps) HasSetup() bool`

HasSetup returns a boolean if a field has been set.

### GetSeverity

`func (o *SavedQueryRuleCreateProps) GetSeverity() Severity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *SavedQueryRuleCreateProps) GetSeverityOk() (*Severity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *SavedQueryRuleCreateProps) SetSeverity(v Severity)`

SetSeverity sets Severity field to given value.


### GetSeverityMapping

`func (o *SavedQueryRuleCreateProps) GetSeverityMapping() []SeverityMappingInner`

GetSeverityMapping returns the SeverityMapping field if non-nil, zero value otherwise.

### GetSeverityMappingOk

`func (o *SavedQueryRuleCreateProps) GetSeverityMappingOk() (*[]SeverityMappingInner, bool)`

GetSeverityMappingOk returns a tuple with the SeverityMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityMapping

`func (o *SavedQueryRuleCreateProps) SetSeverityMapping(v []SeverityMappingInner)`

SetSeverityMapping sets SeverityMapping field to given value.

### HasSeverityMapping

`func (o *SavedQueryRuleCreateProps) HasSeverityMapping() bool`

HasSeverityMapping returns a boolean if a field has been set.

### GetTags

`func (o *SavedQueryRuleCreateProps) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SavedQueryRuleCreateProps) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SavedQueryRuleCreateProps) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SavedQueryRuleCreateProps) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThreat

`func (o *SavedQueryRuleCreateProps) GetThreat() []Threat`

GetThreat returns the Threat field if non-nil, zero value otherwise.

### GetThreatOk

`func (o *SavedQueryRuleCreateProps) GetThreatOk() (*[]Threat, bool)`

GetThreatOk returns a tuple with the Threat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreat

`func (o *SavedQueryRuleCreateProps) SetThreat(v []Threat)`

SetThreat sets Threat field to given value.

### HasThreat

`func (o *SavedQueryRuleCreateProps) HasThreat() bool`

HasThreat returns a boolean if a field has been set.

### GetThrottle

`func (o *SavedQueryRuleCreateProps) GetThrottle() RuleActionThrottle`

GetThrottle returns the Throttle field if non-nil, zero value otherwise.

### GetThrottleOk

`func (o *SavedQueryRuleCreateProps) GetThrottleOk() (*RuleActionThrottle, bool)`

GetThrottleOk returns a tuple with the Throttle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottle

`func (o *SavedQueryRuleCreateProps) SetThrottle(v RuleActionThrottle)`

SetThrottle sets Throttle field to given value.

### HasThrottle

`func (o *SavedQueryRuleCreateProps) HasThrottle() bool`

HasThrottle returns a boolean if a field has been set.

### GetTimelineId

`func (o *SavedQueryRuleCreateProps) GetTimelineId() string`

GetTimelineId returns the TimelineId field if non-nil, zero value otherwise.

### GetTimelineIdOk

`func (o *SavedQueryRuleCreateProps) GetTimelineIdOk() (*string, bool)`

GetTimelineIdOk returns a tuple with the TimelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineId

`func (o *SavedQueryRuleCreateProps) SetTimelineId(v string)`

SetTimelineId sets TimelineId field to given value.

### HasTimelineId

`func (o *SavedQueryRuleCreateProps) HasTimelineId() bool`

HasTimelineId returns a boolean if a field has been set.

### GetTimelineTitle

`func (o *SavedQueryRuleCreateProps) GetTimelineTitle() string`

GetTimelineTitle returns the TimelineTitle field if non-nil, zero value otherwise.

### GetTimelineTitleOk

`func (o *SavedQueryRuleCreateProps) GetTimelineTitleOk() (*string, bool)`

GetTimelineTitleOk returns a tuple with the TimelineTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineTitle

`func (o *SavedQueryRuleCreateProps) SetTimelineTitle(v string)`

SetTimelineTitle sets TimelineTitle field to given value.

### HasTimelineTitle

`func (o *SavedQueryRuleCreateProps) HasTimelineTitle() bool`

HasTimelineTitle returns a boolean if a field has been set.

### GetTimestampOverride

`func (o *SavedQueryRuleCreateProps) GetTimestampOverride() string`

GetTimestampOverride returns the TimestampOverride field if non-nil, zero value otherwise.

### GetTimestampOverrideOk

`func (o *SavedQueryRuleCreateProps) GetTimestampOverrideOk() (*string, bool)`

GetTimestampOverrideOk returns a tuple with the TimestampOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampOverride

`func (o *SavedQueryRuleCreateProps) SetTimestampOverride(v string)`

SetTimestampOverride sets TimestampOverride field to given value.

### HasTimestampOverride

`func (o *SavedQueryRuleCreateProps) HasTimestampOverride() bool`

HasTimestampOverride returns a boolean if a field has been set.

### GetTimestampOverrideFallbackDisabled

`func (o *SavedQueryRuleCreateProps) GetTimestampOverrideFallbackDisabled() bool`

GetTimestampOverrideFallbackDisabled returns the TimestampOverrideFallbackDisabled field if non-nil, zero value otherwise.

### GetTimestampOverrideFallbackDisabledOk

`func (o *SavedQueryRuleCreateProps) GetTimestampOverrideFallbackDisabledOk() (*bool, bool)`

GetTimestampOverrideFallbackDisabledOk returns a tuple with the TimestampOverrideFallbackDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampOverrideFallbackDisabled

`func (o *SavedQueryRuleCreateProps) SetTimestampOverrideFallbackDisabled(v bool)`

SetTimestampOverrideFallbackDisabled sets TimestampOverrideFallbackDisabled field to given value.

### HasTimestampOverrideFallbackDisabled

`func (o *SavedQueryRuleCreateProps) HasTimestampOverrideFallbackDisabled() bool`

HasTimestampOverrideFallbackDisabled returns a boolean if a field has been set.

### GetTo

`func (o *SavedQueryRuleCreateProps) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SavedQueryRuleCreateProps) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SavedQueryRuleCreateProps) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *SavedQueryRuleCreateProps) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetVersion

`func (o *SavedQueryRuleCreateProps) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SavedQueryRuleCreateProps) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SavedQueryRuleCreateProps) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SavedQueryRuleCreateProps) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


