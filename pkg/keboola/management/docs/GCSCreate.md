# GCSCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GcsCredentials** | Pointer to [**GCPCredentials**](GCPCredentials.md) |  | [optional] 
**FilesBucket** | **string** | name of the bucket | 
**Owner** | **string** |  | 
**Region** | **string** |  | 

## Methods

### NewGCSCreate

`func NewGCSCreate(filesBucket string, owner string, region string, ) *GCSCreate`

NewGCSCreate instantiates a new GCSCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGCSCreateWithDefaults

`func NewGCSCreateWithDefaults() *GCSCreate`

NewGCSCreateWithDefaults instantiates a new GCSCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGcsCredentials

`func (o *GCSCreate) GetGcsCredentials() GCPCredentials`

GetGcsCredentials returns the GcsCredentials field if non-nil, zero value otherwise.

### GetGcsCredentialsOk

`func (o *GCSCreate) GetGcsCredentialsOk() (*GCPCredentials, bool)`

GetGcsCredentialsOk returns a tuple with the GcsCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcsCredentials

`func (o *GCSCreate) SetGcsCredentials(v GCPCredentials)`

SetGcsCredentials sets GcsCredentials field to given value.

### HasGcsCredentials

`func (o *GCSCreate) HasGcsCredentials() bool`

HasGcsCredentials returns a boolean if a field has been set.

### GetFilesBucket

`func (o *GCSCreate) GetFilesBucket() string`

GetFilesBucket returns the FilesBucket field if non-nil, zero value otherwise.

### GetFilesBucketOk

`func (o *GCSCreate) GetFilesBucketOk() (*string, bool)`

GetFilesBucketOk returns a tuple with the FilesBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesBucket

`func (o *GCSCreate) SetFilesBucket(v string)`

SetFilesBucket sets FilesBucket field to given value.


### GetOwner

`func (o *GCSCreate) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GCSCreate) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GCSCreate) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetRegion

`func (o *GCSCreate) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GCSCreate) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GCSCreate) SetRegion(v string)`

SetRegion sets Region field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


