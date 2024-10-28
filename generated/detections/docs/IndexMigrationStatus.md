# IndexMigrationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **string** | A string that is not empty and does not contain only whitespace | 
**IsOutdated** | **bool** |  | 
**Migrations** | [**[]MigrationStatus**](MigrationStatus.md) |  | 
**SignalVersions** | [**[]AlertVersion**](AlertVersion.md) |  | 
**Version** | **int32** |  | 

## Methods

### NewIndexMigrationStatus

`func NewIndexMigrationStatus(index string, isOutdated bool, migrations []MigrationStatus, signalVersions []AlertVersion, version int32, ) *IndexMigrationStatus`

NewIndexMigrationStatus instantiates a new IndexMigrationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexMigrationStatusWithDefaults

`func NewIndexMigrationStatusWithDefaults() *IndexMigrationStatus`

NewIndexMigrationStatusWithDefaults instantiates a new IndexMigrationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *IndexMigrationStatus) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *IndexMigrationStatus) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *IndexMigrationStatus) SetIndex(v string)`

SetIndex sets Index field to given value.


### GetIsOutdated

`func (o *IndexMigrationStatus) GetIsOutdated() bool`

GetIsOutdated returns the IsOutdated field if non-nil, zero value otherwise.

### GetIsOutdatedOk

`func (o *IndexMigrationStatus) GetIsOutdatedOk() (*bool, bool)`

GetIsOutdatedOk returns a tuple with the IsOutdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOutdated

`func (o *IndexMigrationStatus) SetIsOutdated(v bool)`

SetIsOutdated sets IsOutdated field to given value.


### GetMigrations

`func (o *IndexMigrationStatus) GetMigrations() []MigrationStatus`

GetMigrations returns the Migrations field if non-nil, zero value otherwise.

### GetMigrationsOk

`func (o *IndexMigrationStatus) GetMigrationsOk() (*[]MigrationStatus, bool)`

GetMigrationsOk returns a tuple with the Migrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrations

`func (o *IndexMigrationStatus) SetMigrations(v []MigrationStatus)`

SetMigrations sets Migrations field to given value.


### GetSignalVersions

`func (o *IndexMigrationStatus) GetSignalVersions() []AlertVersion`

GetSignalVersions returns the SignalVersions field if non-nil, zero value otherwise.

### GetSignalVersionsOk

`func (o *IndexMigrationStatus) GetSignalVersionsOk() (*[]AlertVersion, bool)`

GetSignalVersionsOk returns a tuple with the SignalVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalVersions

`func (o *IndexMigrationStatus) SetSignalVersions(v []AlertVersion)`

SetSignalVersions sets SignalVersions field to given value.


### GetVersion

`func (o *IndexMigrationStatus) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IndexMigrationStatus) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IndexMigrationStatus) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


