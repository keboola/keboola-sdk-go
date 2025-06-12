# S3FileStorageBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsKey** | **string** |  | 
**FilesBucket** | **string** |  | 
**Region** | **string** |  | 
**Owner** | **string** | Associated AWS account owner | 

## Methods

### NewS3FileStorageBase

`func NewS3FileStorageBase(awsKey string, filesBucket string, region string, owner string, ) *S3FileStorageBase`

NewS3FileStorageBase instantiates a new S3FileStorageBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3FileStorageBaseWithDefaults

`func NewS3FileStorageBaseWithDefaults() *S3FileStorageBase`

NewS3FileStorageBaseWithDefaults instantiates a new S3FileStorageBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsKey

`func (o *S3FileStorageBase) GetAwsKey() string`

GetAwsKey returns the AwsKey field if non-nil, zero value otherwise.

### GetAwsKeyOk

`func (o *S3FileStorageBase) GetAwsKeyOk() (*string, bool)`

GetAwsKeyOk returns a tuple with the AwsKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsKey

`func (o *S3FileStorageBase) SetAwsKey(v string)`

SetAwsKey sets AwsKey field to given value.


### GetFilesBucket

`func (o *S3FileStorageBase) GetFilesBucket() string`

GetFilesBucket returns the FilesBucket field if non-nil, zero value otherwise.

### GetFilesBucketOk

`func (o *S3FileStorageBase) GetFilesBucketOk() (*string, bool)`

GetFilesBucketOk returns a tuple with the FilesBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesBucket

`func (o *S3FileStorageBase) SetFilesBucket(v string)`

SetFilesBucket sets FilesBucket field to given value.


### GetRegion

`func (o *S3FileStorageBase) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *S3FileStorageBase) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *S3FileStorageBase) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetOwner

`func (o *S3FileStorageBase) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *S3FileStorageBase) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *S3FileStorageBase) SetOwner(v string)`

SetOwner sets Owner field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


