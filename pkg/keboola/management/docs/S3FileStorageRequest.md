# S3FileStorageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsKey** | **string** |  | 
**FilesBucket** | **string** |  | 
**Region** | **string** |  | 
**Owner** | **string** | Associated AWS account owner | 
**AwsSecret** | **string** |  | 

## Methods

### NewS3FileStorageRequest

`func NewS3FileStorageRequest(awsKey string, filesBucket string, region string, owner string, awsSecret string, ) *S3FileStorageRequest`

NewS3FileStorageRequest instantiates a new S3FileStorageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3FileStorageRequestWithDefaults

`func NewS3FileStorageRequestWithDefaults() *S3FileStorageRequest`

NewS3FileStorageRequestWithDefaults instantiates a new S3FileStorageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsKey

`func (o *S3FileStorageRequest) GetAwsKey() string`

GetAwsKey returns the AwsKey field if non-nil, zero value otherwise.

### GetAwsKeyOk

`func (o *S3FileStorageRequest) GetAwsKeyOk() (*string, bool)`

GetAwsKeyOk returns a tuple with the AwsKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsKey

`func (o *S3FileStorageRequest) SetAwsKey(v string)`

SetAwsKey sets AwsKey field to given value.


### GetFilesBucket

`func (o *S3FileStorageRequest) GetFilesBucket() string`

GetFilesBucket returns the FilesBucket field if non-nil, zero value otherwise.

### GetFilesBucketOk

`func (o *S3FileStorageRequest) GetFilesBucketOk() (*string, bool)`

GetFilesBucketOk returns a tuple with the FilesBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesBucket

`func (o *S3FileStorageRequest) SetFilesBucket(v string)`

SetFilesBucket sets FilesBucket field to given value.


### GetRegion

`func (o *S3FileStorageRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *S3FileStorageRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *S3FileStorageRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetOwner

`func (o *S3FileStorageRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *S3FileStorageRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *S3FileStorageRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetAwsSecret

`func (o *S3FileStorageRequest) GetAwsSecret() string`

GetAwsSecret returns the AwsSecret field if non-nil, zero value otherwise.

### GetAwsSecretOk

`func (o *S3FileStorageRequest) GetAwsSecretOk() (*string, bool)`

GetAwsSecretOk returns a tuple with the AwsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecret

`func (o *S3FileStorageRequest) SetAwsSecret(v string)`

SetAwsSecret sets AwsSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


