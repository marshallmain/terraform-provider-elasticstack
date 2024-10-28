# RuleExceptionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A string that is not empty and does not contain only whitespace | 
**ListId** | **string** | A string that is not empty and does not contain only whitespace | 
**NamespaceType** | **string** | Determines the exceptions validity in rule&#39;s Kibana space | 
**Type** | [**ExceptionListType**](ExceptionListType.md) |  | 

## Methods

### NewRuleExceptionList

`func NewRuleExceptionList(id string, listId string, namespaceType string, type_ ExceptionListType, ) *RuleExceptionList`

NewRuleExceptionList instantiates a new RuleExceptionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleExceptionListWithDefaults

`func NewRuleExceptionListWithDefaults() *RuleExceptionList`

NewRuleExceptionListWithDefaults instantiates a new RuleExceptionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RuleExceptionList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuleExceptionList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuleExceptionList) SetId(v string)`

SetId sets Id field to given value.


### GetListId

`func (o *RuleExceptionList) GetListId() string`

GetListId returns the ListId field if non-nil, zero value otherwise.

### GetListIdOk

`func (o *RuleExceptionList) GetListIdOk() (*string, bool)`

GetListIdOk returns a tuple with the ListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListId

`func (o *RuleExceptionList) SetListId(v string)`

SetListId sets ListId field to given value.


### GetNamespaceType

`func (o *RuleExceptionList) GetNamespaceType() string`

GetNamespaceType returns the NamespaceType field if non-nil, zero value otherwise.

### GetNamespaceTypeOk

`func (o *RuleExceptionList) GetNamespaceTypeOk() (*string, bool)`

GetNamespaceTypeOk returns a tuple with the NamespaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceType

`func (o *RuleExceptionList) SetNamespaceType(v string)`

SetNamespaceType sets NamespaceType field to given value.


### GetType

`func (o *RuleExceptionList) GetType() ExceptionListType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RuleExceptionList) GetTypeOk() (*ExceptionListType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RuleExceptionList) SetType(v ExceptionListType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


