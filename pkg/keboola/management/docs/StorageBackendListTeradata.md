# StorageBackendListTeradata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**Region** | Pointer to **string** | east-1 (required) | [optional] 
**Owner** | **string** |  | 
**Username** | Pointer to **string** |  | [optional] 
**Stats** | Pointer to [**StorageBackendListBaseStats**](StorageBackendListBaseStats.md) |  | [optional] 
**Created** | **string** | The datetime (ISO8601 format) &#39;Y-m-d\\TH:i:sO&#39; | 
**Creator** | Pointer to [**StorageBackendListBaseCreator**](StorageBackendListBaseCreator.md) |  | [optional] 
**Host** | **string** |  | 
**Backend** | **string** |  | 
**Database** | Pointer to **string** |  | [optional] 

## Methods

### NewStorageBackendListTeradata

`func NewStorageBackendListTeradata(id float32, owner string, created string, host string, backend string, ) *StorageBackendListTeradata`

NewStorageBackendListTeradata instantiates a new StorageBackendListTeradata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBackendListTeradataWithDefaults

`func NewStorageBackendListTeradataWithDefaults() *StorageBackendListTeradata`

NewStorageBackendListTeradataWithDefaults instantiates a new StorageBackendListTeradata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorageBackendListTeradata) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageBackendListTeradata) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageBackendListTeradata) SetId(v float32)`

SetId sets Id field to given value.


### GetRegion

`func (o *StorageBackendListTeradata) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *StorageBackendListTeradata) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *StorageBackendListTeradata) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *StorageBackendListTeradata) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetOwner

`func (o *StorageBackendListTeradata) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *StorageBackendListTeradata) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *StorageBackendListTeradata) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetUsername

`func (o *StorageBackendListTeradata) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *StorageBackendListTeradata) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *StorageBackendListTeradata) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *StorageBackendListTeradata) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetStats

`func (o *StorageBackendListTeradata) GetStats() StorageBackendListBaseStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *StorageBackendListTeradata) GetStatsOk() (*StorageBackendListBaseStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *StorageBackendListTeradata) SetStats(v StorageBackendListBaseStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *StorageBackendListTeradata) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetCreated

`func (o *StorageBackendListTeradata) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *StorageBackendListTeradata) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *StorageBackendListTeradata) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetCreator

`func (o *StorageBackendListTeradata) GetCreator() StorageBackendListBaseCreator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *StorageBackendListTeradata) GetCreatorOk() (*StorageBackendListBaseCreator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *StorageBackendListTeradata) SetCreator(v StorageBackendListBaseCreator)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *StorageBackendListTeradata) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetHost

`func (o *StorageBackendListTeradata) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *StorageBackendListTeradata) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *StorageBackendListTeradata) SetHost(v string)`

SetHost sets Host field to given value.


### GetBackend

`func (o *StorageBackendListTeradata) GetBackend() string`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *StorageBackendListTeradata) GetBackendOk() (*string, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *StorageBackendListTeradata) SetBackend(v string)`

SetBackend sets Backend field to given value.


### GetDatabase

`func (o *StorageBackendListTeradata) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *StorageBackendListTeradata) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *StorageBackendListTeradata) SetDatabase(v string)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *StorageBackendListTeradata) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


