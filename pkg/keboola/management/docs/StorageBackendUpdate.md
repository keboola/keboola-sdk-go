# StorageBackendUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**UseDynamicBackends** | Pointer to **bool** | Only for backends supporting dynamic sizing (Snowflake). When enabled, new projects get dynamic backends assigned automatically. | [optional] 

## Methods

### NewStorageBackendUpdate

`func NewStorageBackendUpdate() *StorageBackendUpdate`

NewStorageBackendUpdate instantiates a new StorageBackendUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBackendUpdateWithDefaults

`func NewStorageBackendUpdateWithDefaults() *StorageBackendUpdate`

NewStorageBackendUpdateWithDefaults instantiates a new StorageBackendUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *StorageBackendUpdate) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *StorageBackendUpdate) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *StorageBackendUpdate) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *StorageBackendUpdate) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *StorageBackendUpdate) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *StorageBackendUpdate) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *StorageBackendUpdate) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *StorageBackendUpdate) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUseDynamicBackends

`func (o *StorageBackendUpdate) GetUseDynamicBackends() bool`

GetUseDynamicBackends returns the UseDynamicBackends field if non-nil, zero value otherwise.

### GetUseDynamicBackendsOk

`func (o *StorageBackendUpdate) GetUseDynamicBackendsOk() (*bool, bool)`

GetUseDynamicBackendsOk returns a tuple with the UseDynamicBackends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDynamicBackends

`func (o *StorageBackendUpdate) SetUseDynamicBackends(v bool)`

SetUseDynamicBackends sets UseDynamicBackends field to given value.

### HasUseDynamicBackends

`func (o *StorageBackendUpdate) HasUseDynamicBackends() bool`

HasUseDynamicBackends returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


