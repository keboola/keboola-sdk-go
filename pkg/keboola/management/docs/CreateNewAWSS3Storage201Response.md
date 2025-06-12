# CreateNewAWSS3Storage201Response

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
**Creator** | Pointer to [**CreateNewAWSS3Storage201ResponseCreator**](CreateNewAWSS3Storage201ResponseCreator.md) |  | [optional] 

## Methods

### NewCreateNewAWSS3Storage201Response

`func NewCreateNewAWSS3Storage201Response(awsKey string, filesBucket string, region string, owner string, id string, ) *CreateNewAWSS3Storage201Response`

NewCreateNewAWSS3Storage201Response instantiates a new CreateNewAWSS3Storage201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNewAWSS3Storage201ResponseWithDefaults

`func NewCreateNewAWSS3Storage201ResponseWithDefaults() *CreateNewAWSS3Storage201Response`

NewCreateNewAWSS3Storage201ResponseWithDefaults instantiates a new CreateNewAWSS3Storage201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsKey

`func (o *CreateNewAWSS3Storage201Response) GetAwsKey() string`

GetAwsKey returns the AwsKey field if non-nil, zero value otherwise.

### GetAwsKeyOk

`func (o *CreateNewAWSS3Storage201Response) GetAwsKeyOk() (*string, bool)`

GetAwsKeyOk returns a tuple with the AwsKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsKey

`func (o *CreateNewAWSS3Storage201Response) SetAwsKey(v string)`

SetAwsKey sets AwsKey field to given value.


### GetFilesBucket

`func (o *CreateNewAWSS3Storage201Response) GetFilesBucket() string`

GetFilesBucket returns the FilesBucket field if non-nil, zero value otherwise.

### GetFilesBucketOk

`func (o *CreateNewAWSS3Storage201Response) GetFilesBucketOk() (*string, bool)`

GetFilesBucketOk returns a tuple with the FilesBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesBucket

`func (o *CreateNewAWSS3Storage201Response) SetFilesBucket(v string)`

SetFilesBucket sets FilesBucket field to given value.


### GetRegion

`func (o *CreateNewAWSS3Storage201Response) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateNewAWSS3Storage201Response) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateNewAWSS3Storage201Response) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetOwner

`func (o *CreateNewAWSS3Storage201Response) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateNewAWSS3Storage201Response) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateNewAWSS3Storage201Response) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetId

`func (o *CreateNewAWSS3Storage201Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateNewAWSS3Storage201Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateNewAWSS3Storage201Response) SetId(v string)`

SetId sets Id field to given value.


### GetIsDefault

`func (o *CreateNewAWSS3Storage201Response) GetIsDefault() string`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CreateNewAWSS3Storage201Response) GetIsDefaultOk() (*string, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CreateNewAWSS3Storage201Response) SetIsDefault(v string)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *CreateNewAWSS3Storage201Response) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetProvider

`func (o *CreateNewAWSS3Storage201Response) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CreateNewAWSS3Storage201Response) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CreateNewAWSS3Storage201Response) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CreateNewAWSS3Storage201Response) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetCreated

`func (o *CreateNewAWSS3Storage201Response) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CreateNewAWSS3Storage201Response) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CreateNewAWSS3Storage201Response) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CreateNewAWSS3Storage201Response) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreator

`func (o *CreateNewAWSS3Storage201Response) GetCreator() CreateNewAWSS3Storage201ResponseCreator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *CreateNewAWSS3Storage201Response) GetCreatorOk() (*CreateNewAWSS3Storage201ResponseCreator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *CreateNewAWSS3Storage201Response) SetCreator(v CreateNewAWSS3Storage201ResponseCreator)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *CreateNewAWSS3Storage201Response) HasCreator() bool`

HasCreator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


