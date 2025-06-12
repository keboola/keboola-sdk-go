# StorageBackend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**Host** | Pointer to **string** | 37-demo.cmizbsfmzc6w.us-east-1.redshift.amazonaws.com (required) | [optional] 
**Backend** | **string** |  | 
**Region** | Pointer to **string** | east-1 | [optional] 

## Methods

### NewStorageBackend

`func NewStorageBackend(id float32, backend string, ) *StorageBackend`

NewStorageBackend instantiates a new StorageBackend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBackendWithDefaults

`func NewStorageBackendWithDefaults() *StorageBackend`

NewStorageBackendWithDefaults instantiates a new StorageBackend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorageBackend) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageBackend) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageBackend) SetId(v float32)`

SetId sets Id field to given value.


### GetHost

`func (o *StorageBackend) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *StorageBackend) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *StorageBackend) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *StorageBackend) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetBackend

`func (o *StorageBackend) GetBackend() string`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *StorageBackend) GetBackendOk() (*string, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *StorageBackend) SetBackend(v string)`

SetBackend sets Backend field to given value.


### GetRegion

`func (o *StorageBackend) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *StorageBackend) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *StorageBackend) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *StorageBackend) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


