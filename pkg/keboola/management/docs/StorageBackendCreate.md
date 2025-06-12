# StorageBackendCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backend** | **string** | can be redshift, snowflake ,synapse, exasol, teradata | 
**Host** | **string** |  | 
**Warehouse** | Pointer to **string** | required only for Snowflake | [optional] 
**Username** | **string** |  | 
**Password** | **string** |  | 
**Region** | **string** |  | 
**Owner** | **string** | associated AWS account owner | 
**Database** | Pointer to **string** | required for Synapse and Teradata | [optional] 
**UseSynapseManagedIdentity** | Pointer to **string** | optional for Synapse | [optional] 
**UseDynamicBackends** | Pointer to **bool** | Only for backends supporting dynamic sizing (Snowflake). When enabled, new projects get dynamic backends assigned automatically. | [optional] 

## Methods

### NewStorageBackendCreate

`func NewStorageBackendCreate(backend string, host string, username string, password string, region string, owner string, ) *StorageBackendCreate`

NewStorageBackendCreate instantiates a new StorageBackendCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBackendCreateWithDefaults

`func NewStorageBackendCreateWithDefaults() *StorageBackendCreate`

NewStorageBackendCreateWithDefaults instantiates a new StorageBackendCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackend

`func (o *StorageBackendCreate) GetBackend() string`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *StorageBackendCreate) GetBackendOk() (*string, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *StorageBackendCreate) SetBackend(v string)`

SetBackend sets Backend field to given value.


### GetHost

`func (o *StorageBackendCreate) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *StorageBackendCreate) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *StorageBackendCreate) SetHost(v string)`

SetHost sets Host field to given value.


### GetWarehouse

`func (o *StorageBackendCreate) GetWarehouse() string`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *StorageBackendCreate) GetWarehouseOk() (*string, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *StorageBackendCreate) SetWarehouse(v string)`

SetWarehouse sets Warehouse field to given value.

### HasWarehouse

`func (o *StorageBackendCreate) HasWarehouse() bool`

HasWarehouse returns a boolean if a field has been set.

### GetUsername

`func (o *StorageBackendCreate) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *StorageBackendCreate) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *StorageBackendCreate) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *StorageBackendCreate) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *StorageBackendCreate) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *StorageBackendCreate) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetRegion

`func (o *StorageBackendCreate) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *StorageBackendCreate) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *StorageBackendCreate) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetOwner

`func (o *StorageBackendCreate) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *StorageBackendCreate) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *StorageBackendCreate) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetDatabase

`func (o *StorageBackendCreate) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *StorageBackendCreate) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *StorageBackendCreate) SetDatabase(v string)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *StorageBackendCreate) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetUseSynapseManagedIdentity

`func (o *StorageBackendCreate) GetUseSynapseManagedIdentity() string`

GetUseSynapseManagedIdentity returns the UseSynapseManagedIdentity field if non-nil, zero value otherwise.

### GetUseSynapseManagedIdentityOk

`func (o *StorageBackendCreate) GetUseSynapseManagedIdentityOk() (*string, bool)`

GetUseSynapseManagedIdentityOk returns a tuple with the UseSynapseManagedIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSynapseManagedIdentity

`func (o *StorageBackendCreate) SetUseSynapseManagedIdentity(v string)`

SetUseSynapseManagedIdentity sets UseSynapseManagedIdentity field to given value.

### HasUseSynapseManagedIdentity

`func (o *StorageBackendCreate) HasUseSynapseManagedIdentity() bool`

HasUseSynapseManagedIdentity returns a boolean if a field has been set.

### GetUseDynamicBackends

`func (o *StorageBackendCreate) GetUseDynamicBackends() bool`

GetUseDynamicBackends returns the UseDynamicBackends field if non-nil, zero value otherwise.

### GetUseDynamicBackendsOk

`func (o *StorageBackendCreate) GetUseDynamicBackendsOk() (*bool, bool)`

GetUseDynamicBackendsOk returns a tuple with the UseDynamicBackends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDynamicBackends

`func (o *StorageBackendCreate) SetUseDynamicBackends(v bool)`

SetUseDynamicBackends sets UseDynamicBackends field to given value.

### HasUseDynamicBackends

`func (o *StorageBackendCreate) HasUseDynamicBackends() bool`

HasUseDynamicBackends returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


