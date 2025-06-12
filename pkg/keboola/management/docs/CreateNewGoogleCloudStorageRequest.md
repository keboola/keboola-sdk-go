# CreateNewGoogleCloudStorageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GcsCredentials** | Pointer to [**CreateNewGoogleCloudStorageRequestGcsCredentials**](CreateNewGoogleCloudStorageRequestGcsCredentials.md) |  | [optional] 
**FilesBucket** | **string** | name of the bucket | 
**Owner** | **string** |  | 
**Region** | **string** |  | 

## Methods

### NewCreateNewGoogleCloudStorageRequest

`func NewCreateNewGoogleCloudStorageRequest(filesBucket string, owner string, region string, ) *CreateNewGoogleCloudStorageRequest`

NewCreateNewGoogleCloudStorageRequest instantiates a new CreateNewGoogleCloudStorageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNewGoogleCloudStorageRequestWithDefaults

`func NewCreateNewGoogleCloudStorageRequestWithDefaults() *CreateNewGoogleCloudStorageRequest`

NewCreateNewGoogleCloudStorageRequestWithDefaults instantiates a new CreateNewGoogleCloudStorageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGcsCredentials

`func (o *CreateNewGoogleCloudStorageRequest) GetGcsCredentials() CreateNewGoogleCloudStorageRequestGcsCredentials`

GetGcsCredentials returns the GcsCredentials field if non-nil, zero value otherwise.

### GetGcsCredentialsOk

`func (o *CreateNewGoogleCloudStorageRequest) GetGcsCredentialsOk() (*CreateNewGoogleCloudStorageRequestGcsCredentials, bool)`

GetGcsCredentialsOk returns a tuple with the GcsCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcsCredentials

`func (o *CreateNewGoogleCloudStorageRequest) SetGcsCredentials(v CreateNewGoogleCloudStorageRequestGcsCredentials)`

SetGcsCredentials sets GcsCredentials field to given value.

### HasGcsCredentials

`func (o *CreateNewGoogleCloudStorageRequest) HasGcsCredentials() bool`

HasGcsCredentials returns a boolean if a field has been set.

### GetFilesBucket

`func (o *CreateNewGoogleCloudStorageRequest) GetFilesBucket() string`

GetFilesBucket returns the FilesBucket field if non-nil, zero value otherwise.

### GetFilesBucketOk

`func (o *CreateNewGoogleCloudStorageRequest) GetFilesBucketOk() (*string, bool)`

GetFilesBucketOk returns a tuple with the FilesBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesBucket

`func (o *CreateNewGoogleCloudStorageRequest) SetFilesBucket(v string)`

SetFilesBucket sets FilesBucket field to given value.


### GetOwner

`func (o *CreateNewGoogleCloudStorageRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateNewGoogleCloudStorageRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateNewGoogleCloudStorageRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetRegion

`func (o *CreateNewGoogleCloudStorageRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateNewGoogleCloudStorageRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateNewGoogleCloudStorageRequest) SetRegion(v string)`

SetRegion sets Region field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


