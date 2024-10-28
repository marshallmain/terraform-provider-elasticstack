# ThreatMatchRulePatchProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to **string** |  | [optional] 
**ThreatIndex** | Pointer to **[]string** |  | [optional] 
**ThreatMapping** | Pointer to [**[]ThreatMappingInner**](ThreatMappingInner.md) |  | [optional] 
**ThreatQuery** | Pointer to **string** | Query to run | [optional] 
**Type** | Pointer to **string** | Rule type | [optional] 
**AlertSuppression** | Pointer to [**AlertSuppression**](AlertSuppression.md) |  | [optional] 
**ConcurrentSearches** | Pointer to **int32** |  | [optional] 
**DataViewId** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to **[]interface{}** |  | [optional] 
**Index** | Pointer to **[]string** |  | [optional] 
**ItemsPerSearch** | Pointer to **int32** |  | [optional] 
**SavedId** | Pointer to **string** |  | [optional] 
**ThreatFilters** | Pointer to **[]interface{}** |  | [optional] 
**ThreatIndicatorPath** | Pointer to **string** | Defines the path to the threat indicator in the indicator documents (optional) | [optional] 
**ThreatLanguage** | Pointer to [**KqlQueryLanguage**](KqlQueryLanguage.md) |  | [optional] 
**Language** | Pointer to [**KqlQueryLanguage**](KqlQueryLanguage.md) |  | [optional] 
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
**License** | Pointer to **string** | The rule&#39;s license. | [optional] 
**MaxSignals** | Pointer to **int32** |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** | Has no effect. | [optional] 
**Note** | Pointer to **string** | Notes to help investigate alerts produced by the rule. | [optional] 
**Outcome** | Pointer to [**SavedObjectResolveOutcome**](SavedObjectResolveOutcome.md) |  | [optional] 
**OutputIndex** | Pointer to **string** | (deprecated) Has no effect. | [optional] 
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
**Version** | Pointer to **int32** | The rule&#39;s version number. | [optional] 

## Methods

### NewThreatMatchRulePatchProps

`func NewThreatMatchRulePatchProps() *ThreatMatchRulePatchProps`

NewThreatMatchRulePatchProps instantiates a new ThreatMatchRulePatchProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreatMatchRulePatchPropsWithDefaults

`func NewThreatMatchRulePatchPropsWithDefaults() *ThreatMatchRulePatchProps`

NewThreatMatchRulePatchPropsWithDefaults instantiates a new ThreatMatchRulePatchProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *ThreatMatchRulePatchProps) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ThreatMatchRulePatchProps) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ThreatMatchRulePatchProps) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *ThreatMatchRulePatchProps) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetThreatIndex

`func (o *ThreatMatchRulePatchProps) GetThreatIndex() []string`

GetThreatIndex returns the ThreatIndex field if non-nil, zero value otherwise.

### GetThreatIndexOk

`func (o *ThreatMatchRulePatchProps) GetThreatIndexOk() (*[]string, bool)`

GetThreatIndexOk returns a tuple with the ThreatIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatIndex

`func (o *ThreatMatchRulePatchProps) SetThreatIndex(v []string)`

SetThreatIndex sets ThreatIndex field to given value.

### HasThreatIndex

`func (o *ThreatMatchRulePatchProps) HasThreatIndex() bool`

HasThreatIndex returns a boolean if a field has been set.

### GetThreatMapping

`func (o *ThreatMatchRulePatchProps) GetThreatMapping() []ThreatMappingInner`

GetThreatMapping returns the ThreatMapping field if non-nil, zero value otherwise.

### GetThreatMappingOk

`func (o *ThreatMatchRulePatchProps) GetThreatMappingOk() (*[]ThreatMappingInner, bool)`

GetThreatMappingOk returns a tuple with the ThreatMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatMapping

`func (o *ThreatMatchRulePatchProps) SetThreatMapping(v []ThreatMappingInner)`

SetThreatMapping sets ThreatMapping field to given value.

### HasThreatMapping

`func (o *ThreatMatchRulePatchProps) HasThreatMapping() bool`

HasThreatMapping returns a boolean if a field has been set.

### GetThreatQuery

`func (o *ThreatMatchRulePatchProps) GetThreatQuery() string`

GetThreatQuery returns the ThreatQuery field if non-nil, zero value otherwise.

### GetThreatQueryOk

`func (o *ThreatMatchRulePatchProps) GetThreatQueryOk() (*string, bool)`

GetThreatQueryOk returns a tuple with the ThreatQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatQuery

`func (o *ThreatMatchRulePatchProps) SetThreatQuery(v string)`

SetThreatQuery sets ThreatQuery field to given value.

### HasThreatQuery

`func (o *ThreatMatchRulePatchProps) HasThreatQuery() bool`

HasThreatQuery returns a boolean if a field has been set.

### GetType

`func (o *ThreatMatchRulePatchProps) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ThreatMatchRulePatchProps) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ThreatMatchRulePatchProps) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ThreatMatchRulePatchProps) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAlertSuppression

`func (o *ThreatMatchRulePatchProps) GetAlertSuppression() AlertSuppression`

GetAlertSuppression returns the AlertSuppression field if non-nil, zero value otherwise.

### GetAlertSuppressionOk

`func (o *ThreatMatchRulePatchProps) GetAlertSuppressionOk() (*AlertSuppression, bool)`

GetAlertSuppressionOk returns a tuple with the AlertSuppression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertSuppression

`func (o *ThreatMatchRulePatchProps) SetAlertSuppression(v AlertSuppression)`

SetAlertSuppression sets AlertSuppression field to given value.

### HasAlertSuppression

`func (o *ThreatMatchRulePatchProps) HasAlertSuppression() bool`

HasAlertSuppression returns a boolean if a field has been set.

### GetConcurrentSearches

`func (o *ThreatMatchRulePatchProps) GetConcurrentSearches() int32`

GetConcurrentSearches returns the ConcurrentSearches field if non-nil, zero value otherwise.

### GetConcurrentSearchesOk

`func (o *ThreatMatchRulePatchProps) GetConcurrentSearchesOk() (*int32, bool)`

GetConcurrentSearchesOk returns a tuple with the ConcurrentSearches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrentSearches

`func (o *ThreatMatchRulePatchProps) SetConcurrentSearches(v int32)`

SetConcurrentSearches sets ConcurrentSearches field to given value.

### HasConcurrentSearches

`func (o *ThreatMatchRulePatchProps) HasConcurrentSearches() bool`

HasConcurrentSearches returns a boolean if a field has been set.

### GetDataViewId

`func (o *ThreatMatchRulePatchProps) GetDataViewId() string`

GetDataViewId returns the DataViewId field if non-nil, zero value otherwise.

### GetDataViewIdOk

`func (o *ThreatMatchRulePatchProps) GetDataViewIdOk() (*string, bool)`

GetDataViewIdOk returns a tuple with the DataViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataViewId

`func (o *ThreatMatchRulePatchProps) SetDataViewId(v string)`

SetDataViewId sets DataViewId field to given value.

### HasDataViewId

`func (o *ThreatMatchRulePatchProps) HasDataViewId() bool`

HasDataViewId returns a boolean if a field has been set.

### GetFilters

`func (o *ThreatMatchRulePatchProps) GetFilters() []interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ThreatMatchRulePatchProps) GetFiltersOk() (*[]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ThreatMatchRulePatchProps) SetFilters(v []interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ThreatMatchRulePatchProps) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetIndex

`func (o *ThreatMatchRulePatchProps) GetIndex() []string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ThreatMatchRulePatchProps) GetIndexOk() (*[]string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ThreatMatchRulePatchProps) SetIndex(v []string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *ThreatMatchRulePatchProps) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetItemsPerSearch

`func (o *ThreatMatchRulePatchProps) GetItemsPerSearch() int32`

GetItemsPerSearch returns the ItemsPerSearch field if non-nil, zero value otherwise.

### GetItemsPerSearchOk

`func (o *ThreatMatchRulePatchProps) GetItemsPerSearchOk() (*int32, bool)`

GetItemsPerSearchOk returns a tuple with the ItemsPerSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsPerSearch

`func (o *ThreatMatchRulePatchProps) SetItemsPerSearch(v int32)`

SetItemsPerSearch sets ItemsPerSearch field to given value.

### HasItemsPerSearch

`func (o *ThreatMatchRulePatchProps) HasItemsPerSearch() bool`

HasItemsPerSearch returns a boolean if a field has been set.

### GetSavedId

`func (o *ThreatMatchRulePatchProps) GetSavedId() string`

GetSavedId returns the SavedId field if non-nil, zero value otherwise.

### GetSavedIdOk

`func (o *ThreatMatchRulePatchProps) GetSavedIdOk() (*string, bool)`

GetSavedIdOk returns a tuple with the SavedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedId

`func (o *ThreatMatchRulePatchProps) SetSavedId(v string)`

SetSavedId sets SavedId field to given value.

### HasSavedId

`func (o *ThreatMatchRulePatchProps) HasSavedId() bool`

HasSavedId returns a boolean if a field has been set.

### GetThreatFilters

`func (o *ThreatMatchRulePatchProps) GetThreatFilters() []interface{}`

GetThreatFilters returns the ThreatFilters field if non-nil, zero value otherwise.

### GetThreatFiltersOk

`func (o *ThreatMatchRulePatchProps) GetThreatFiltersOk() (*[]interface{}, bool)`

GetThreatFiltersOk returns a tuple with the ThreatFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatFilters

`func (o *ThreatMatchRulePatchProps) SetThreatFilters(v []interface{})`

SetThreatFilters sets ThreatFilters field to given value.

### HasThreatFilters

`func (o *ThreatMatchRulePatchProps) HasThreatFilters() bool`

HasThreatFilters returns a boolean if a field has been set.

### GetThreatIndicatorPath

`func (o *ThreatMatchRulePatchProps) GetThreatIndicatorPath() string`

GetThreatIndicatorPath returns the ThreatIndicatorPath field if non-nil, zero value otherwise.

### GetThreatIndicatorPathOk

`func (o *ThreatMatchRulePatchProps) GetThreatIndicatorPathOk() (*string, bool)`

GetThreatIndicatorPathOk returns a tuple with the ThreatIndicatorPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatIndicatorPath

`func (o *ThreatMatchRulePatchProps) SetThreatIndicatorPath(v string)`

SetThreatIndicatorPath sets ThreatIndicatorPath field to given value.

### HasThreatIndicatorPath

`func (o *ThreatMatchRulePatchProps) HasThreatIndicatorPath() bool`

HasThreatIndicatorPath returns a boolean if a field has been set.

### GetThreatLanguage

`func (o *ThreatMatchRulePatchProps) GetThreatLanguage() KqlQueryLanguage`

GetThreatLanguage returns the ThreatLanguage field if non-nil, zero value otherwise.

### GetThreatLanguageOk

`func (o *ThreatMatchRulePatchProps) GetThreatLanguageOk() (*KqlQueryLanguage, bool)`

GetThreatLanguageOk returns a tuple with the ThreatLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatLanguage

`func (o *ThreatMatchRulePatchProps) SetThreatLanguage(v KqlQueryLanguage)`

SetThreatLanguage sets ThreatLanguage field to given value.

### HasThreatLanguage

`func (o *ThreatMatchRulePatchProps) HasThreatLanguage() bool`

HasThreatLanguage returns a boolean if a field has been set.

### GetLanguage

`func (o *ThreatMatchRulePatchProps) GetLanguage() KqlQueryLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ThreatMatchRulePatchProps) GetLanguageOk() (*KqlQueryLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ThreatMatchRulePatchProps) SetLanguage(v KqlQueryLanguage)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ThreatMatchRulePatchProps) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetActions

`func (o *ThreatMatchRulePatchProps) GetActions() []RuleAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ThreatMatchRulePatchProps) GetActionsOk() (*[]RuleAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ThreatMatchRulePatchProps) SetActions(v []RuleAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ThreatMatchRulePatchProps) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetAliasPurpose

`func (o *ThreatMatchRulePatchProps) GetAliasPurpose() SavedObjectResolveAliasPurpose`

GetAliasPurpose returns the AliasPurpose field if non-nil, zero value otherwise.

### GetAliasPurposeOk

`func (o *ThreatMatchRulePatchProps) GetAliasPurposeOk() (*SavedObjectResolveAliasPurpose, bool)`

GetAliasPurposeOk returns a tuple with the AliasPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasPurpose

`func (o *ThreatMatchRulePatchProps) SetAliasPurpose(v SavedObjectResolveAliasPurpose)`

SetAliasPurpose sets AliasPurpose field to given value.

### HasAliasPurpose

`func (o *ThreatMatchRulePatchProps) HasAliasPurpose() bool`

HasAliasPurpose returns a boolean if a field has been set.

### GetAliasTargetId

`func (o *ThreatMatchRulePatchProps) GetAliasTargetId() string`

GetAliasTargetId returns the AliasTargetId field if non-nil, zero value otherwise.

### GetAliasTargetIdOk

`func (o *ThreatMatchRulePatchProps) GetAliasTargetIdOk() (*string, bool)`

GetAliasTargetIdOk returns a tuple with the AliasTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasTargetId

`func (o *ThreatMatchRulePatchProps) SetAliasTargetId(v string)`

SetAliasTargetId sets AliasTargetId field to given value.

### HasAliasTargetId

`func (o *ThreatMatchRulePatchProps) HasAliasTargetId() bool`

HasAliasTargetId returns a boolean if a field has been set.

### GetAuthor

`func (o *ThreatMatchRulePatchProps) GetAuthor() []string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ThreatMatchRulePatchProps) GetAuthorOk() (*[]string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ThreatMatchRulePatchProps) SetAuthor(v []string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ThreatMatchRulePatchProps) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetBuildingBlockType

`func (o *ThreatMatchRulePatchProps) GetBuildingBlockType() string`

GetBuildingBlockType returns the BuildingBlockType field if non-nil, zero value otherwise.

### GetBuildingBlockTypeOk

`func (o *ThreatMatchRulePatchProps) GetBuildingBlockTypeOk() (*string, bool)`

GetBuildingBlockTypeOk returns a tuple with the BuildingBlockType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildingBlockType

`func (o *ThreatMatchRulePatchProps) SetBuildingBlockType(v string)`

SetBuildingBlockType sets BuildingBlockType field to given value.

### HasBuildingBlockType

`func (o *ThreatMatchRulePatchProps) HasBuildingBlockType() bool`

HasBuildingBlockType returns a boolean if a field has been set.

### GetDescription

`func (o *ThreatMatchRulePatchProps) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ThreatMatchRulePatchProps) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ThreatMatchRulePatchProps) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ThreatMatchRulePatchProps) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ThreatMatchRulePatchProps) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ThreatMatchRulePatchProps) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ThreatMatchRulePatchProps) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ThreatMatchRulePatchProps) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExceptionsList

`func (o *ThreatMatchRulePatchProps) GetExceptionsList() []RuleExceptionList`

GetExceptionsList returns the ExceptionsList field if non-nil, zero value otherwise.

### GetExceptionsListOk

`func (o *ThreatMatchRulePatchProps) GetExceptionsListOk() (*[]RuleExceptionList, bool)`

GetExceptionsListOk returns a tuple with the ExceptionsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionsList

`func (o *ThreatMatchRulePatchProps) SetExceptionsList(v []RuleExceptionList)`

SetExceptionsList sets ExceptionsList field to given value.

### HasExceptionsList

`func (o *ThreatMatchRulePatchProps) HasExceptionsList() bool`

HasExceptionsList returns a boolean if a field has been set.

### GetFalsePositives

`func (o *ThreatMatchRulePatchProps) GetFalsePositives() []string`

GetFalsePositives returns the FalsePositives field if non-nil, zero value otherwise.

### GetFalsePositivesOk

`func (o *ThreatMatchRulePatchProps) GetFalsePositivesOk() (*[]string, bool)`

GetFalsePositivesOk returns a tuple with the FalsePositives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFalsePositives

`func (o *ThreatMatchRulePatchProps) SetFalsePositives(v []string)`

SetFalsePositives sets FalsePositives field to given value.

### HasFalsePositives

`func (o *ThreatMatchRulePatchProps) HasFalsePositives() bool`

HasFalsePositives returns a boolean if a field has been set.

### GetFrom

`func (o *ThreatMatchRulePatchProps) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ThreatMatchRulePatchProps) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ThreatMatchRulePatchProps) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ThreatMatchRulePatchProps) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetId

`func (o *ThreatMatchRulePatchProps) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThreatMatchRulePatchProps) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThreatMatchRulePatchProps) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ThreatMatchRulePatchProps) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterval

`func (o *ThreatMatchRulePatchProps) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *ThreatMatchRulePatchProps) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *ThreatMatchRulePatchProps) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *ThreatMatchRulePatchProps) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetInvestigationFields

`func (o *ThreatMatchRulePatchProps) GetInvestigationFields() InvestigationFields`

GetInvestigationFields returns the InvestigationFields field if non-nil, zero value otherwise.

### GetInvestigationFieldsOk

`func (o *ThreatMatchRulePatchProps) GetInvestigationFieldsOk() (*InvestigationFields, bool)`

GetInvestigationFieldsOk returns a tuple with the InvestigationFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigationFields

`func (o *ThreatMatchRulePatchProps) SetInvestigationFields(v InvestigationFields)`

SetInvestigationFields sets InvestigationFields field to given value.

### HasInvestigationFields

`func (o *ThreatMatchRulePatchProps) HasInvestigationFields() bool`

HasInvestigationFields returns a boolean if a field has been set.

### GetLicense

`func (o *ThreatMatchRulePatchProps) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *ThreatMatchRulePatchProps) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *ThreatMatchRulePatchProps) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *ThreatMatchRulePatchProps) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetMaxSignals

`func (o *ThreatMatchRulePatchProps) GetMaxSignals() int32`

GetMaxSignals returns the MaxSignals field if non-nil, zero value otherwise.

### GetMaxSignalsOk

`func (o *ThreatMatchRulePatchProps) GetMaxSignalsOk() (*int32, bool)`

GetMaxSignalsOk returns a tuple with the MaxSignals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSignals

`func (o *ThreatMatchRulePatchProps) SetMaxSignals(v int32)`

SetMaxSignals sets MaxSignals field to given value.

### HasMaxSignals

`func (o *ThreatMatchRulePatchProps) HasMaxSignals() bool`

HasMaxSignals returns a boolean if a field has been set.

### GetMeta

`func (o *ThreatMatchRulePatchProps) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ThreatMatchRulePatchProps) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ThreatMatchRulePatchProps) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ThreatMatchRulePatchProps) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetName

`func (o *ThreatMatchRulePatchProps) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ThreatMatchRulePatchProps) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ThreatMatchRulePatchProps) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ThreatMatchRulePatchProps) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *ThreatMatchRulePatchProps) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ThreatMatchRulePatchProps) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ThreatMatchRulePatchProps) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ThreatMatchRulePatchProps) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNote

`func (o *ThreatMatchRulePatchProps) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ThreatMatchRulePatchProps) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ThreatMatchRulePatchProps) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *ThreatMatchRulePatchProps) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetOutcome

`func (o *ThreatMatchRulePatchProps) GetOutcome() SavedObjectResolveOutcome`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *ThreatMatchRulePatchProps) GetOutcomeOk() (*SavedObjectResolveOutcome, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *ThreatMatchRulePatchProps) SetOutcome(v SavedObjectResolveOutcome)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *ThreatMatchRulePatchProps) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetOutputIndex

`func (o *ThreatMatchRulePatchProps) GetOutputIndex() string`

GetOutputIndex returns the OutputIndex field if non-nil, zero value otherwise.

### GetOutputIndexOk

`func (o *ThreatMatchRulePatchProps) GetOutputIndexOk() (*string, bool)`

GetOutputIndexOk returns a tuple with the OutputIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputIndex

`func (o *ThreatMatchRulePatchProps) SetOutputIndex(v string)`

SetOutputIndex sets OutputIndex field to given value.

### HasOutputIndex

`func (o *ThreatMatchRulePatchProps) HasOutputIndex() bool`

HasOutputIndex returns a boolean if a field has been set.

### GetReferences

`func (o *ThreatMatchRulePatchProps) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ThreatMatchRulePatchProps) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ThreatMatchRulePatchProps) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *ThreatMatchRulePatchProps) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetRelatedIntegrations

`func (o *ThreatMatchRulePatchProps) GetRelatedIntegrations() []RelatedIntegration`

GetRelatedIntegrations returns the RelatedIntegrations field if non-nil, zero value otherwise.

### GetRelatedIntegrationsOk

`func (o *ThreatMatchRulePatchProps) GetRelatedIntegrationsOk() (*[]RelatedIntegration, bool)`

GetRelatedIntegrationsOk returns a tuple with the RelatedIntegrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedIntegrations

`func (o *ThreatMatchRulePatchProps) SetRelatedIntegrations(v []RelatedIntegration)`

SetRelatedIntegrations sets RelatedIntegrations field to given value.

### HasRelatedIntegrations

`func (o *ThreatMatchRulePatchProps) HasRelatedIntegrations() bool`

HasRelatedIntegrations returns a boolean if a field has been set.

### GetRequiredFields

`func (o *ThreatMatchRulePatchProps) GetRequiredFields() []RequiredFieldInput`

GetRequiredFields returns the RequiredFields field if non-nil, zero value otherwise.

### GetRequiredFieldsOk

`func (o *ThreatMatchRulePatchProps) GetRequiredFieldsOk() (*[]RequiredFieldInput, bool)`

GetRequiredFieldsOk returns a tuple with the RequiredFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredFields

`func (o *ThreatMatchRulePatchProps) SetRequiredFields(v []RequiredFieldInput)`

SetRequiredFields sets RequiredFields field to given value.

### HasRequiredFields

`func (o *ThreatMatchRulePatchProps) HasRequiredFields() bool`

HasRequiredFields returns a boolean if a field has been set.

### GetRiskScore

`func (o *ThreatMatchRulePatchProps) GetRiskScore() int32`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *ThreatMatchRulePatchProps) GetRiskScoreOk() (*int32, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *ThreatMatchRulePatchProps) SetRiskScore(v int32)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *ThreatMatchRulePatchProps) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### GetRiskScoreMapping

`func (o *ThreatMatchRulePatchProps) GetRiskScoreMapping() []RiskScoreMappingInner`

GetRiskScoreMapping returns the RiskScoreMapping field if non-nil, zero value otherwise.

### GetRiskScoreMappingOk

`func (o *ThreatMatchRulePatchProps) GetRiskScoreMappingOk() (*[]RiskScoreMappingInner, bool)`

GetRiskScoreMappingOk returns a tuple with the RiskScoreMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScoreMapping

`func (o *ThreatMatchRulePatchProps) SetRiskScoreMapping(v []RiskScoreMappingInner)`

SetRiskScoreMapping sets RiskScoreMapping field to given value.

### HasRiskScoreMapping

`func (o *ThreatMatchRulePatchProps) HasRiskScoreMapping() bool`

HasRiskScoreMapping returns a boolean if a field has been set.

### GetRuleId

`func (o *ThreatMatchRulePatchProps) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *ThreatMatchRulePatchProps) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *ThreatMatchRulePatchProps) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *ThreatMatchRulePatchProps) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetRuleNameOverride

`func (o *ThreatMatchRulePatchProps) GetRuleNameOverride() string`

GetRuleNameOverride returns the RuleNameOverride field if non-nil, zero value otherwise.

### GetRuleNameOverrideOk

`func (o *ThreatMatchRulePatchProps) GetRuleNameOverrideOk() (*string, bool)`

GetRuleNameOverrideOk returns a tuple with the RuleNameOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleNameOverride

`func (o *ThreatMatchRulePatchProps) SetRuleNameOverride(v string)`

SetRuleNameOverride sets RuleNameOverride field to given value.

### HasRuleNameOverride

`func (o *ThreatMatchRulePatchProps) HasRuleNameOverride() bool`

HasRuleNameOverride returns a boolean if a field has been set.

### GetSetup

`func (o *ThreatMatchRulePatchProps) GetSetup() string`

GetSetup returns the Setup field if non-nil, zero value otherwise.

### GetSetupOk

`func (o *ThreatMatchRulePatchProps) GetSetupOk() (*string, bool)`

GetSetupOk returns a tuple with the Setup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetup

`func (o *ThreatMatchRulePatchProps) SetSetup(v string)`

SetSetup sets Setup field to given value.

### HasSetup

`func (o *ThreatMatchRulePatchProps) HasSetup() bool`

HasSetup returns a boolean if a field has been set.

### GetSeverity

`func (o *ThreatMatchRulePatchProps) GetSeverity() Severity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ThreatMatchRulePatchProps) GetSeverityOk() (*Severity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ThreatMatchRulePatchProps) SetSeverity(v Severity)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ThreatMatchRulePatchProps) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSeverityMapping

`func (o *ThreatMatchRulePatchProps) GetSeverityMapping() []SeverityMappingInner`

GetSeverityMapping returns the SeverityMapping field if non-nil, zero value otherwise.

### GetSeverityMappingOk

`func (o *ThreatMatchRulePatchProps) GetSeverityMappingOk() (*[]SeverityMappingInner, bool)`

GetSeverityMappingOk returns a tuple with the SeverityMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityMapping

`func (o *ThreatMatchRulePatchProps) SetSeverityMapping(v []SeverityMappingInner)`

SetSeverityMapping sets SeverityMapping field to given value.

### HasSeverityMapping

`func (o *ThreatMatchRulePatchProps) HasSeverityMapping() bool`

HasSeverityMapping returns a boolean if a field has been set.

### GetTags

`func (o *ThreatMatchRulePatchProps) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ThreatMatchRulePatchProps) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ThreatMatchRulePatchProps) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ThreatMatchRulePatchProps) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThreat

`func (o *ThreatMatchRulePatchProps) GetThreat() []Threat`

GetThreat returns the Threat field if non-nil, zero value otherwise.

### GetThreatOk

`func (o *ThreatMatchRulePatchProps) GetThreatOk() (*[]Threat, bool)`

GetThreatOk returns a tuple with the Threat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreat

`func (o *ThreatMatchRulePatchProps) SetThreat(v []Threat)`

SetThreat sets Threat field to given value.

### HasThreat

`func (o *ThreatMatchRulePatchProps) HasThreat() bool`

HasThreat returns a boolean if a field has been set.

### GetThrottle

`func (o *ThreatMatchRulePatchProps) GetThrottle() RuleActionThrottle`

GetThrottle returns the Throttle field if non-nil, zero value otherwise.

### GetThrottleOk

`func (o *ThreatMatchRulePatchProps) GetThrottleOk() (*RuleActionThrottle, bool)`

GetThrottleOk returns a tuple with the Throttle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottle

`func (o *ThreatMatchRulePatchProps) SetThrottle(v RuleActionThrottle)`

SetThrottle sets Throttle field to given value.

### HasThrottle

`func (o *ThreatMatchRulePatchProps) HasThrottle() bool`

HasThrottle returns a boolean if a field has been set.

### GetTimelineId

`func (o *ThreatMatchRulePatchProps) GetTimelineId() string`

GetTimelineId returns the TimelineId field if non-nil, zero value otherwise.

### GetTimelineIdOk

`func (o *ThreatMatchRulePatchProps) GetTimelineIdOk() (*string, bool)`

GetTimelineIdOk returns a tuple with the TimelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineId

`func (o *ThreatMatchRulePatchProps) SetTimelineId(v string)`

SetTimelineId sets TimelineId field to given value.

### HasTimelineId

`func (o *ThreatMatchRulePatchProps) HasTimelineId() bool`

HasTimelineId returns a boolean if a field has been set.

### GetTimelineTitle

`func (o *ThreatMatchRulePatchProps) GetTimelineTitle() string`

GetTimelineTitle returns the TimelineTitle field if non-nil, zero value otherwise.

### GetTimelineTitleOk

`func (o *ThreatMatchRulePatchProps) GetTimelineTitleOk() (*string, bool)`

GetTimelineTitleOk returns a tuple with the TimelineTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineTitle

`func (o *ThreatMatchRulePatchProps) SetTimelineTitle(v string)`

SetTimelineTitle sets TimelineTitle field to given value.

### HasTimelineTitle

`func (o *ThreatMatchRulePatchProps) HasTimelineTitle() bool`

HasTimelineTitle returns a boolean if a field has been set.

### GetTimestampOverride

`func (o *ThreatMatchRulePatchProps) GetTimestampOverride() string`

GetTimestampOverride returns the TimestampOverride field if non-nil, zero value otherwise.

### GetTimestampOverrideOk

`func (o *ThreatMatchRulePatchProps) GetTimestampOverrideOk() (*string, bool)`

GetTimestampOverrideOk returns a tuple with the TimestampOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampOverride

`func (o *ThreatMatchRulePatchProps) SetTimestampOverride(v string)`

SetTimestampOverride sets TimestampOverride field to given value.

### HasTimestampOverride

`func (o *ThreatMatchRulePatchProps) HasTimestampOverride() bool`

HasTimestampOverride returns a boolean if a field has been set.

### GetTimestampOverrideFallbackDisabled

`func (o *ThreatMatchRulePatchProps) GetTimestampOverrideFallbackDisabled() bool`

GetTimestampOverrideFallbackDisabled returns the TimestampOverrideFallbackDisabled field if non-nil, zero value otherwise.

### GetTimestampOverrideFallbackDisabledOk

`func (o *ThreatMatchRulePatchProps) GetTimestampOverrideFallbackDisabledOk() (*bool, bool)`

GetTimestampOverrideFallbackDisabledOk returns a tuple with the TimestampOverrideFallbackDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampOverrideFallbackDisabled

`func (o *ThreatMatchRulePatchProps) SetTimestampOverrideFallbackDisabled(v bool)`

SetTimestampOverrideFallbackDisabled sets TimestampOverrideFallbackDisabled field to given value.

### HasTimestampOverrideFallbackDisabled

`func (o *ThreatMatchRulePatchProps) HasTimestampOverrideFallbackDisabled() bool`

HasTimestampOverrideFallbackDisabled returns a boolean if a field has been set.

### GetTo

`func (o *ThreatMatchRulePatchProps) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ThreatMatchRulePatchProps) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ThreatMatchRulePatchProps) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *ThreatMatchRulePatchProps) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetVersion

`func (o *ThreatMatchRulePatchProps) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ThreatMatchRulePatchProps) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ThreatMatchRulePatchProps) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ThreatMatchRulePatchProps) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


