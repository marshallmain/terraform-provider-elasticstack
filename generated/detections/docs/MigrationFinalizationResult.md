# MigrationFinalizationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Completed** | **bool** |  | 
**DestinationIndex** | **string** |  | 
**Error** | Pointer to [**MigrationCleanupResultError**](MigrationCleanupResultError.md) |  | [optional] 
**Id** | **string** |  | 
**SourceIndex** | **string** |  | 
**Status** | **string** |  | 
**Updated** | **time.Time** |  | 
**Version** | **string** |  | 

## Methods

### NewMigrationFinalizationResult

`func NewMigrationFinalizationResult(completed bool, destinationIndex string, id string, sourceIndex string, status string, updated time.Time, version string, ) *MigrationFinalizationResult`

NewMigrationFinalizationResult instantiates a new MigrationFinalizationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationFinalizationResultWithDefaults

`func NewMigrationFinalizationResultWithDefaults() *MigrationFinalizationResult`

NewMigrationFinalizationResultWithDefaults instantiates a new MigrationFinalizationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompleted

`func (o *MigrationFinalizationResult) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *MigrationFinalizationResult) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *MigrationFinalizationResult) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.


### GetDestinationIndex

`func (o *MigrationFinalizationResult) GetDestinationIndex() string`

GetDestinationIndex returns the DestinationIndex field if non-nil, zero value otherwise.

### GetDestinationIndexOk

`func (o *MigrationFinalizationResult) GetDestinationIndexOk() (*string, bool)`

GetDestinationIndexOk returns a tuple with the DestinationIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIndex

`func (o *MigrationFinalizationResult) SetDestinationIndex(v string)`

SetDestinationIndex sets DestinationIndex field to given value.


### GetError

`func (o *MigrationFinalizationResult) GetError() MigrationCleanupResultError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MigrationFinalizationResult) GetErrorOk() (*MigrationCleanupResultError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MigrationFinalizationResult) SetError(v MigrationCleanupResultError)`

SetError sets Error field to given value.

### HasError

`func (o *MigrationFinalizationResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetId

`func (o *MigrationFinalizationResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MigrationFinalizationResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MigrationFinalizationResult) SetId(v string)`

SetId sets Id field to given value.


### GetSourceIndex

`func (o *MigrationFinalizationResult) GetSourceIndex() string`

GetSourceIndex returns the SourceIndex field if non-nil, zero value otherwise.

### GetSourceIndexOk

`func (o *MigrationFinalizationResult) GetSourceIndexOk() (*string, bool)`

GetSourceIndexOk returns a tuple with the SourceIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIndex

`func (o *MigrationFinalizationResult) SetSourceIndex(v string)`

SetSourceIndex sets SourceIndex field to given value.


### GetStatus

`func (o *MigrationFinalizationResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MigrationFinalizationResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MigrationFinalizationResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdated

`func (o *MigrationFinalizationResult) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *MigrationFinalizationResult) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *MigrationFinalizationResult) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.


### GetVersion

`func (o *MigrationFinalizationResult) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MigrationFinalizationResult) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MigrationFinalizationResult) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


