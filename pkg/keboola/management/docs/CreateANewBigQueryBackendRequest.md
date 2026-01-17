# CreateANewBigQueryBackendRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | **string** | associated GCP account owner | 
**FolderId** | **string** | where is service account created | 
**Region** | **string** |  | 
**Credentials** | Pointer to [**CreateNewGoogleCloudStorageRequestGcsCredentials**](CreateNewGoogleCloudStorageRequestGcsCredentials.md) |  | [optional] 

## Methods

### NewCreateANewBigQueryBackendRequest

`func NewCreateANewBigQueryBackendRequest(owner string, folderId string, region string, ) *CreateANewBigQueryBackendRequest`

NewCreateANewBigQueryBackendRequest instantiates a new CreateANewBigQueryBackendRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateANewBigQueryBackendRequestWithDefaults

`func NewCreateANewBigQueryBackendRequestWithDefaults() *CreateANewBigQueryBackendRequest`

NewCreateANewBigQueryBackendRequestWithDefaults instantiates a new CreateANewBigQueryBackendRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *CreateANewBigQueryBackendRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateANewBigQueryBackendRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateANewBigQueryBackendRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetFolderId

`func (o *CreateANewBigQueryBackendRequest) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *CreateANewBigQueryBackendRequest) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *CreateANewBigQueryBackendRequest) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.


### GetRegion

`func (o *CreateANewBigQueryBackendRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateANewBigQueryBackendRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateANewBigQueryBackendRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetCredentials

`func (o *CreateANewBigQueryBackendRequest) GetCredentials() CreateNewGoogleCloudStorageRequestGcsCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *CreateANewBigQueryBackendRequest) GetCredentialsOk() (*CreateNewGoogleCloudStorageRequestGcsCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *CreateANewBigQueryBackendRequest) SetCredentials(v CreateNewGoogleCloudStorageRequestGcsCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *CreateANewBigQueryBackendRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


