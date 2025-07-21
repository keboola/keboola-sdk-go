# CreateANewBigQueryBackend201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] 
**Region** | **string** |  | 
**Backend** | Pointer to **string** |  | [optional] 
**Credentials** | Pointer to [**CreateNewGoogleCloudStorage201ResponseGcsCredentials**](CreateNewGoogleCloudStorage201ResponseGcsCredentials.md) |  | [optional] 
**FolderId** | Pointer to **float32** |  | [optional] 

## Methods

### NewCreateANewBigQueryBackend201Response

`func NewCreateANewBigQueryBackend201Response(region string, ) *CreateANewBigQueryBackend201Response`

NewCreateANewBigQueryBackend201Response instantiates a new CreateANewBigQueryBackend201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateANewBigQueryBackend201ResponseWithDefaults

`func NewCreateANewBigQueryBackend201ResponseWithDefaults() *CreateANewBigQueryBackend201Response`

NewCreateANewBigQueryBackend201ResponseWithDefaults instantiates a new CreateANewBigQueryBackend201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateANewBigQueryBackend201Response) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateANewBigQueryBackend201Response) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateANewBigQueryBackend201Response) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *CreateANewBigQueryBackend201Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRegion

`func (o *CreateANewBigQueryBackend201Response) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateANewBigQueryBackend201Response) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateANewBigQueryBackend201Response) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetBackend

`func (o *CreateANewBigQueryBackend201Response) GetBackend() string`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *CreateANewBigQueryBackend201Response) GetBackendOk() (*string, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *CreateANewBigQueryBackend201Response) SetBackend(v string)`

SetBackend sets Backend field to given value.

### HasBackend

`func (o *CreateANewBigQueryBackend201Response) HasBackend() bool`

HasBackend returns a boolean if a field has been set.

### GetCredentials

`func (o *CreateANewBigQueryBackend201Response) GetCredentials() CreateNewGoogleCloudStorage201ResponseGcsCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *CreateANewBigQueryBackend201Response) GetCredentialsOk() (*CreateNewGoogleCloudStorage201ResponseGcsCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *CreateANewBigQueryBackend201Response) SetCredentials(v CreateNewGoogleCloudStorage201ResponseGcsCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *CreateANewBigQueryBackend201Response) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetFolderId

`func (o *CreateANewBigQueryBackend201Response) GetFolderId() float32`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *CreateANewBigQueryBackend201Response) GetFolderIdOk() (*float32, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *CreateANewBigQueryBackend201Response) SetFolderId(v float32)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *CreateANewBigQueryBackend201Response) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


