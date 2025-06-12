# StorageBackendListSnowflake

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
**Warehouse** | Pointer to **string** |  | [optional] 
**Saml2Configuration** | Pointer to [**StorageBackendListSnowflakeAllOfSaml2Configuration**](StorageBackendListSnowflakeAllOfSaml2Configuration.md) |  | [optional] 

## Methods

### NewStorageBackendListSnowflake

`func NewStorageBackendListSnowflake(id float32, owner string, created string, host string, backend string, ) *StorageBackendListSnowflake`

NewStorageBackendListSnowflake instantiates a new StorageBackendListSnowflake object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBackendListSnowflakeWithDefaults

`func NewStorageBackendListSnowflakeWithDefaults() *StorageBackendListSnowflake`

NewStorageBackendListSnowflakeWithDefaults instantiates a new StorageBackendListSnowflake object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorageBackendListSnowflake) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageBackendListSnowflake) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageBackendListSnowflake) SetId(v float32)`

SetId sets Id field to given value.


### GetRegion

`func (o *StorageBackendListSnowflake) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *StorageBackendListSnowflake) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *StorageBackendListSnowflake) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *StorageBackendListSnowflake) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetOwner

`func (o *StorageBackendListSnowflake) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *StorageBackendListSnowflake) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *StorageBackendListSnowflake) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetUsername

`func (o *StorageBackendListSnowflake) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *StorageBackendListSnowflake) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *StorageBackendListSnowflake) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *StorageBackendListSnowflake) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetStats

`func (o *StorageBackendListSnowflake) GetStats() StorageBackendListBaseStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *StorageBackendListSnowflake) GetStatsOk() (*StorageBackendListBaseStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *StorageBackendListSnowflake) SetStats(v StorageBackendListBaseStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *StorageBackendListSnowflake) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetCreated

`func (o *StorageBackendListSnowflake) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *StorageBackendListSnowflake) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *StorageBackendListSnowflake) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetCreator

`func (o *StorageBackendListSnowflake) GetCreator() StorageBackendListBaseCreator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *StorageBackendListSnowflake) GetCreatorOk() (*StorageBackendListBaseCreator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *StorageBackendListSnowflake) SetCreator(v StorageBackendListBaseCreator)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *StorageBackendListSnowflake) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetHost

`func (o *StorageBackendListSnowflake) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *StorageBackendListSnowflake) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *StorageBackendListSnowflake) SetHost(v string)`

SetHost sets Host field to given value.


### GetBackend

`func (o *StorageBackendListSnowflake) GetBackend() string`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *StorageBackendListSnowflake) GetBackendOk() (*string, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *StorageBackendListSnowflake) SetBackend(v string)`

SetBackend sets Backend field to given value.


### GetWarehouse

`func (o *StorageBackendListSnowflake) GetWarehouse() string`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *StorageBackendListSnowflake) GetWarehouseOk() (*string, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *StorageBackendListSnowflake) SetWarehouse(v string)`

SetWarehouse sets Warehouse field to given value.

### HasWarehouse

`func (o *StorageBackendListSnowflake) HasWarehouse() bool`

HasWarehouse returns a boolean if a field has been set.

### GetSaml2Configuration

`func (o *StorageBackendListSnowflake) GetSaml2Configuration() StorageBackendListSnowflakeAllOfSaml2Configuration`

GetSaml2Configuration returns the Saml2Configuration field if non-nil, zero value otherwise.

### GetSaml2ConfigurationOk

`func (o *StorageBackendListSnowflake) GetSaml2ConfigurationOk() (*StorageBackendListSnowflakeAllOfSaml2Configuration, bool)`

GetSaml2ConfigurationOk returns a tuple with the Saml2Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml2Configuration

`func (o *StorageBackendListSnowflake) SetSaml2Configuration(v StorageBackendListSnowflakeAllOfSaml2Configuration)`

SetSaml2Configuration sets Saml2Configuration field to given value.

### HasSaml2Configuration

`func (o *StorageBackendListSnowflake) HasSaml2Configuration() bool`

HasSaml2Configuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


