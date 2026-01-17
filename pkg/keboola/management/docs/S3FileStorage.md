# S3FileStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsKey** | **string** |  | 
**FilesBucket** | **string** |  | 
**Region** | **string** |  | 
**Owner** | **string** | Associated AWS account owner | 
**Id** | **string** |  | 
**IsDefault** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** | 06-17T15:07:48+01:00 | [optional] 
**Creator** | Pointer to [**ABSCreatedAllOfCreator**](ABSCreatedAllOfCreator.md) |  | [optional] 

## Methods

### NewS3FileStorage

`func NewS3FileStorage(awsKey string, filesBucket string, region string, owner string, id string, ) *S3FileStorage`

NewS3FileStorage instantiates a new S3FileStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3FileStorageWithDefaults

`func NewS3FileStorageWithDefaults() *S3FileStorage`

NewS3FileStorageWithDefaults instantiates a new S3FileStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsKey

`func (o *S3FileStorage) GetAwsKey() string`

GetAwsKey returns the AwsKey field if non-nil, zero value otherwise.

### GetAwsKeyOk

`func (o *S3FileStorage) GetAwsKeyOk() (*string, bool)`

GetAwsKeyOk returns a tuple with the AwsKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsKey

`func (o *S3FileStorage) SetAwsKey(v string)`

SetAwsKey sets AwsKey field to given value.


### GetFilesBucket

`func (o *S3FileStorage) GetFilesBucket() string`

GetFilesBucket returns the FilesBucket field if non-nil, zero value otherwise.

### GetFilesBucketOk

`func (o *S3FileStorage) GetFilesBucketOk() (*string, bool)`

GetFilesBucketOk returns a tuple with the FilesBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesBucket

`func (o *S3FileStorage) SetFilesBucket(v string)`

SetFilesBucket sets FilesBucket field to given value.


### GetRegion

`func (o *S3FileStorage) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *S3FileStorage) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *S3FileStorage) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetOwner

`func (o *S3FileStorage) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *S3FileStorage) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *S3FileStorage) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetId

`func (o *S3FileStorage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *S3FileStorage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *S3FileStorage) SetId(v string)`

SetId sets Id field to given value.


### GetIsDefault

`func (o *S3FileStorage) GetIsDefault() string`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *S3FileStorage) GetIsDefaultOk() (*string, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *S3FileStorage) SetIsDefault(v string)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *S3FileStorage) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetProvider

`func (o *S3FileStorage) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *S3FileStorage) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *S3FileStorage) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *S3FileStorage) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetCreated

`func (o *S3FileStorage) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *S3FileStorage) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *S3FileStorage) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *S3FileStorage) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreator

`func (o *S3FileStorage) GetCreator() ABSCreatedAllOfCreator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *S3FileStorage) GetCreatorOk() (*ABSCreatedAllOfCreator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *S3FileStorage) SetCreator(v ABSCreatedAllOfCreator)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *S3FileStorage) HasCreator() bool`

HasCreator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


