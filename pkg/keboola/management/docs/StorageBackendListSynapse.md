# StorageBackendListSynapse

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
**Host** | Pointer to **string** | synapse.database.windows.net (required) | [optional] 
**Backend** | **string** |  | 
**Database** | Pointer to **string** |  | [optional] 
**UseSynapseManagedIdentity** | Pointer to **string** |  | [optional] 

## Methods

### NewStorageBackendListSynapse

`func NewStorageBackendListSynapse(id float32, owner string, created string, backend string, ) *StorageBackendListSynapse`

NewStorageBackendListSynapse instantiates a new StorageBackendListSynapse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBackendListSynapseWithDefaults

`func NewStorageBackendListSynapseWithDefaults() *StorageBackendListSynapse`

NewStorageBackendListSynapseWithDefaults instantiates a new StorageBackendListSynapse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorageBackendListSynapse) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageBackendListSynapse) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageBackendListSynapse) SetId(v float32)`

SetId sets Id field to given value.


### GetRegion

`func (o *StorageBackendListSynapse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *StorageBackendListSynapse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *StorageBackendListSynapse) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *StorageBackendListSynapse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetOwner

`func (o *StorageBackendListSynapse) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *StorageBackendListSynapse) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *StorageBackendListSynapse) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetUsername

`func (o *StorageBackendListSynapse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *StorageBackendListSynapse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *StorageBackendListSynapse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *StorageBackendListSynapse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetStats

`func (o *StorageBackendListSynapse) GetStats() StorageBackendListBaseStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *StorageBackendListSynapse) GetStatsOk() (*StorageBackendListBaseStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *StorageBackendListSynapse) SetStats(v StorageBackendListBaseStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *StorageBackendListSynapse) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetCreated

`func (o *StorageBackendListSynapse) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *StorageBackendListSynapse) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *StorageBackendListSynapse) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetCreator

`func (o *StorageBackendListSynapse) GetCreator() StorageBackendListBaseCreator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *StorageBackendListSynapse) GetCreatorOk() (*StorageBackendListBaseCreator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *StorageBackendListSynapse) SetCreator(v StorageBackendListBaseCreator)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *StorageBackendListSynapse) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetHost

`func (o *StorageBackendListSynapse) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *StorageBackendListSynapse) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *StorageBackendListSynapse) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *StorageBackendListSynapse) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetBackend

`func (o *StorageBackendListSynapse) GetBackend() string`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *StorageBackendListSynapse) GetBackendOk() (*string, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *StorageBackendListSynapse) SetBackend(v string)`

SetBackend sets Backend field to given value.


### GetDatabase

`func (o *StorageBackendListSynapse) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *StorageBackendListSynapse) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *StorageBackendListSynapse) SetDatabase(v string)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *StorageBackendListSynapse) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetUseSynapseManagedIdentity

`func (o *StorageBackendListSynapse) GetUseSynapseManagedIdentity() string`

GetUseSynapseManagedIdentity returns the UseSynapseManagedIdentity field if non-nil, zero value otherwise.

### GetUseSynapseManagedIdentityOk

`func (o *StorageBackendListSynapse) GetUseSynapseManagedIdentityOk() (*string, bool)`

GetUseSynapseManagedIdentityOk returns a tuple with the UseSynapseManagedIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSynapseManagedIdentity

`func (o *StorageBackendListSynapse) SetUseSynapseManagedIdentity(v string)`

SetUseSynapseManagedIdentity sets UseSynapseManagedIdentity field to given value.

### HasUseSynapseManagedIdentity

`func (o *StorageBackendListSynapse) HasUseSynapseManagedIdentity() bool`

HasUseSynapseManagedIdentity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


