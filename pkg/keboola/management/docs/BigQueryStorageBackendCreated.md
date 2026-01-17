# BigQueryStorageBackendCreated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] 
**Region** | **string** |  | 
**Backend** | Pointer to **string** |  | [optional] 
**Credentials** | Pointer to [**GCPCredentialsNoPK**](GCPCredentialsNoPK.md) |  | [optional] 
**FolderId** | Pointer to **float32** |  | [optional] 

## Methods

### NewBigQueryStorageBackendCreated

`func NewBigQueryStorageBackendCreated(region string, ) *BigQueryStorageBackendCreated`

NewBigQueryStorageBackendCreated instantiates a new BigQueryStorageBackendCreated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBigQueryStorageBackendCreatedWithDefaults

`func NewBigQueryStorageBackendCreatedWithDefaults() *BigQueryStorageBackendCreated`

NewBigQueryStorageBackendCreatedWithDefaults instantiates a new BigQueryStorageBackendCreated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BigQueryStorageBackendCreated) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BigQueryStorageBackendCreated) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BigQueryStorageBackendCreated) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *BigQueryStorageBackendCreated) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRegion

`func (o *BigQueryStorageBackendCreated) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *BigQueryStorageBackendCreated) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *BigQueryStorageBackendCreated) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetBackend

`func (o *BigQueryStorageBackendCreated) GetBackend() string`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *BigQueryStorageBackendCreated) GetBackendOk() (*string, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *BigQueryStorageBackendCreated) SetBackend(v string)`

SetBackend sets Backend field to given value.

### HasBackend

`func (o *BigQueryStorageBackendCreated) HasBackend() bool`

HasBackend returns a boolean if a field has been set.

### GetCredentials

`func (o *BigQueryStorageBackendCreated) GetCredentials() GCPCredentialsNoPK`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *BigQueryStorageBackendCreated) GetCredentialsOk() (*GCPCredentialsNoPK, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *BigQueryStorageBackendCreated) SetCredentials(v GCPCredentialsNoPK)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *BigQueryStorageBackendCreated) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetFolderId

`func (o *BigQueryStorageBackendCreated) GetFolderId() float32`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *BigQueryStorageBackendCreated) GetFolderIdOk() (*float32, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *BigQueryStorageBackendCreated) SetFolderId(v float32)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *BigQueryStorageBackendCreated) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


