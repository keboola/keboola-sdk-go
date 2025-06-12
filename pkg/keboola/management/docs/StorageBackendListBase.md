# StorageBackendListBase

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

## Methods

### NewStorageBackendListBase

`func NewStorageBackendListBase(id float32, owner string, created string, ) *StorageBackendListBase`

NewStorageBackendListBase instantiates a new StorageBackendListBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBackendListBaseWithDefaults

`func NewStorageBackendListBaseWithDefaults() *StorageBackendListBase`

NewStorageBackendListBaseWithDefaults instantiates a new StorageBackendListBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorageBackendListBase) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageBackendListBase) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageBackendListBase) SetId(v float32)`

SetId sets Id field to given value.


### GetRegion

`func (o *StorageBackendListBase) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *StorageBackendListBase) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *StorageBackendListBase) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *StorageBackendListBase) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetOwner

`func (o *StorageBackendListBase) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *StorageBackendListBase) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *StorageBackendListBase) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetUsername

`func (o *StorageBackendListBase) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *StorageBackendListBase) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *StorageBackendListBase) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *StorageBackendListBase) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetStats

`func (o *StorageBackendListBase) GetStats() StorageBackendListBaseStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *StorageBackendListBase) GetStatsOk() (*StorageBackendListBaseStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *StorageBackendListBase) SetStats(v StorageBackendListBaseStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *StorageBackendListBase) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetCreated

`func (o *StorageBackendListBase) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *StorageBackendListBase) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *StorageBackendListBase) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetCreator

`func (o *StorageBackendListBase) GetCreator() StorageBackendListBaseCreator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *StorageBackendListBase) GetCreatorOk() (*StorageBackendListBaseCreator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *StorageBackendListBase) SetCreator(v StorageBackendListBaseCreator)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *StorageBackendListBase) HasCreator() bool`

HasCreator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


