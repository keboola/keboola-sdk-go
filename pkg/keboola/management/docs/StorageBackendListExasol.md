# StorageBackendListExasol

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

## Methods

### NewStorageBackendListExasol

`func NewStorageBackendListExasol(id float32, owner string, created string, host string, backend string, ) *StorageBackendListExasol`

NewStorageBackendListExasol instantiates a new StorageBackendListExasol object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBackendListExasolWithDefaults

`func NewStorageBackendListExasolWithDefaults() *StorageBackendListExasol`

NewStorageBackendListExasolWithDefaults instantiates a new StorageBackendListExasol object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorageBackendListExasol) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageBackendListExasol) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageBackendListExasol) SetId(v float32)`

SetId sets Id field to given value.


### GetRegion

`func (o *StorageBackendListExasol) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *StorageBackendListExasol) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *StorageBackendListExasol) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *StorageBackendListExasol) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetOwner

`func (o *StorageBackendListExasol) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *StorageBackendListExasol) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *StorageBackendListExasol) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetUsername

`func (o *StorageBackendListExasol) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *StorageBackendListExasol) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *StorageBackendListExasol) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *StorageBackendListExasol) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetStats

`func (o *StorageBackendListExasol) GetStats() StorageBackendListBaseStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *StorageBackendListExasol) GetStatsOk() (*StorageBackendListBaseStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *StorageBackendListExasol) SetStats(v StorageBackendListBaseStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *StorageBackendListExasol) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetCreated

`func (o *StorageBackendListExasol) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *StorageBackendListExasol) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *StorageBackendListExasol) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetCreator

`func (o *StorageBackendListExasol) GetCreator() StorageBackendListBaseCreator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *StorageBackendListExasol) GetCreatorOk() (*StorageBackendListBaseCreator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *StorageBackendListExasol) SetCreator(v StorageBackendListBaseCreator)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *StorageBackendListExasol) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetHost

`func (o *StorageBackendListExasol) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *StorageBackendListExasol) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *StorageBackendListExasol) SetHost(v string)`

SetHost sets Host field to given value.


### GetBackend

`func (o *StorageBackendListExasol) GetBackend() string`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *StorageBackendListExasol) GetBackendOk() (*string, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *StorageBackendListExasol) SetBackend(v string)`

SetBackend sets Backend field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


