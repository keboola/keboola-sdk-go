# CreateNewAWSS3StorageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsKey** | **string** |  | 
**FilesBucket** | **string** |  | 
**Region** | **string** |  | 
**Owner** | **string** | Associated AWS account owner | 
**AwsSecret** | **string** |  | 

## Methods

### NewCreateNewAWSS3StorageRequest

`func NewCreateNewAWSS3StorageRequest(awsKey string, filesBucket string, region string, owner string, awsSecret string, ) *CreateNewAWSS3StorageRequest`

NewCreateNewAWSS3StorageRequest instantiates a new CreateNewAWSS3StorageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNewAWSS3StorageRequestWithDefaults

`func NewCreateNewAWSS3StorageRequestWithDefaults() *CreateNewAWSS3StorageRequest`

NewCreateNewAWSS3StorageRequestWithDefaults instantiates a new CreateNewAWSS3StorageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsKey

`func (o *CreateNewAWSS3StorageRequest) GetAwsKey() string`

GetAwsKey returns the AwsKey field if non-nil, zero value otherwise.

### GetAwsKeyOk

`func (o *CreateNewAWSS3StorageRequest) GetAwsKeyOk() (*string, bool)`

GetAwsKeyOk returns a tuple with the AwsKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsKey

`func (o *CreateNewAWSS3StorageRequest) SetAwsKey(v string)`

SetAwsKey sets AwsKey field to given value.


### GetFilesBucket

`func (o *CreateNewAWSS3StorageRequest) GetFilesBucket() string`

GetFilesBucket returns the FilesBucket field if non-nil, zero value otherwise.

### GetFilesBucketOk

`func (o *CreateNewAWSS3StorageRequest) GetFilesBucketOk() (*string, bool)`

GetFilesBucketOk returns a tuple with the FilesBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesBucket

`func (o *CreateNewAWSS3StorageRequest) SetFilesBucket(v string)`

SetFilesBucket sets FilesBucket field to given value.


### GetRegion

`func (o *CreateNewAWSS3StorageRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateNewAWSS3StorageRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateNewAWSS3StorageRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetOwner

`func (o *CreateNewAWSS3StorageRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateNewAWSS3StorageRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateNewAWSS3StorageRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetAwsSecret

`func (o *CreateNewAWSS3StorageRequest) GetAwsSecret() string`

GetAwsSecret returns the AwsSecret field if non-nil, zero value otherwise.

### GetAwsSecretOk

`func (o *CreateNewAWSS3StorageRequest) GetAwsSecretOk() (*string, bool)`

GetAwsSecretOk returns a tuple with the AwsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecret

`func (o *CreateNewAWSS3StorageRequest) SetAwsSecret(v string)`

SetAwsSecret sets AwsSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


