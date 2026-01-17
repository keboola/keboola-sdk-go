# CreateStorageTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Token description | 
**CanManageBuckets** | Pointer to **bool** | Token has full permissions on tabular storage | [optional] 
**CanManageTokens** | Pointer to **bool** | deprecated &amp; will be ignored - Token has permission to create tokens in project | [optional] 
**CanReadAllFileUploads** | Pointer to **bool** | Token has full permissions to files staging | [optional] 
**CanPurgeTrash** | Pointer to **bool** | Allows permanently remove deleted configurations. | [optional] 
**ExpiresIn** | Pointer to **float32** | Token lifetime | [optional] 
**BucketPermissions** | Pointer to [**CreateStorageTokenRequestBucketPermissions**](CreateStorageTokenRequestBucketPermissions.md) |  | [optional] 
**ComponentAccess** | Pointer to **string** | Grants access for component configurations. Allowed values are [valid component IDs](https://components.keboola.com/components). | [optional] 

## Methods

### NewCreateStorageTokenRequest

`func NewCreateStorageTokenRequest(description string, ) *CreateStorageTokenRequest`

NewCreateStorageTokenRequest instantiates a new CreateStorageTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStorageTokenRequestWithDefaults

`func NewCreateStorageTokenRequestWithDefaults() *CreateStorageTokenRequest`

NewCreateStorageTokenRequestWithDefaults instantiates a new CreateStorageTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateStorageTokenRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateStorageTokenRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateStorageTokenRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCanManageBuckets

`func (o *CreateStorageTokenRequest) GetCanManageBuckets() bool`

GetCanManageBuckets returns the CanManageBuckets field if non-nil, zero value otherwise.

### GetCanManageBucketsOk

`func (o *CreateStorageTokenRequest) GetCanManageBucketsOk() (*bool, bool)`

GetCanManageBucketsOk returns a tuple with the CanManageBuckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageBuckets

`func (o *CreateStorageTokenRequest) SetCanManageBuckets(v bool)`

SetCanManageBuckets sets CanManageBuckets field to given value.

### HasCanManageBuckets

`func (o *CreateStorageTokenRequest) HasCanManageBuckets() bool`

HasCanManageBuckets returns a boolean if a field has been set.

### GetCanManageTokens

`func (o *CreateStorageTokenRequest) GetCanManageTokens() bool`

GetCanManageTokens returns the CanManageTokens field if non-nil, zero value otherwise.

### GetCanManageTokensOk

`func (o *CreateStorageTokenRequest) GetCanManageTokensOk() (*bool, bool)`

GetCanManageTokensOk returns a tuple with the CanManageTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageTokens

`func (o *CreateStorageTokenRequest) SetCanManageTokens(v bool)`

SetCanManageTokens sets CanManageTokens field to given value.

### HasCanManageTokens

`func (o *CreateStorageTokenRequest) HasCanManageTokens() bool`

HasCanManageTokens returns a boolean if a field has been set.

### GetCanReadAllFileUploads

`func (o *CreateStorageTokenRequest) GetCanReadAllFileUploads() bool`

GetCanReadAllFileUploads returns the CanReadAllFileUploads field if non-nil, zero value otherwise.

### GetCanReadAllFileUploadsOk

`func (o *CreateStorageTokenRequest) GetCanReadAllFileUploadsOk() (*bool, bool)`

GetCanReadAllFileUploadsOk returns a tuple with the CanReadAllFileUploads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanReadAllFileUploads

`func (o *CreateStorageTokenRequest) SetCanReadAllFileUploads(v bool)`

SetCanReadAllFileUploads sets CanReadAllFileUploads field to given value.

### HasCanReadAllFileUploads

`func (o *CreateStorageTokenRequest) HasCanReadAllFileUploads() bool`

HasCanReadAllFileUploads returns a boolean if a field has been set.

### GetCanPurgeTrash

`func (o *CreateStorageTokenRequest) GetCanPurgeTrash() bool`

GetCanPurgeTrash returns the CanPurgeTrash field if non-nil, zero value otherwise.

### GetCanPurgeTrashOk

`func (o *CreateStorageTokenRequest) GetCanPurgeTrashOk() (*bool, bool)`

GetCanPurgeTrashOk returns a tuple with the CanPurgeTrash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanPurgeTrash

`func (o *CreateStorageTokenRequest) SetCanPurgeTrash(v bool)`

SetCanPurgeTrash sets CanPurgeTrash field to given value.

### HasCanPurgeTrash

`func (o *CreateStorageTokenRequest) HasCanPurgeTrash() bool`

HasCanPurgeTrash returns a boolean if a field has been set.

### GetExpiresIn

`func (o *CreateStorageTokenRequest) GetExpiresIn() float32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *CreateStorageTokenRequest) GetExpiresInOk() (*float32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *CreateStorageTokenRequest) SetExpiresIn(v float32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *CreateStorageTokenRequest) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetBucketPermissions

`func (o *CreateStorageTokenRequest) GetBucketPermissions() CreateStorageTokenRequestBucketPermissions`

GetBucketPermissions returns the BucketPermissions field if non-nil, zero value otherwise.

### GetBucketPermissionsOk

`func (o *CreateStorageTokenRequest) GetBucketPermissionsOk() (*CreateStorageTokenRequestBucketPermissions, bool)`

GetBucketPermissionsOk returns a tuple with the BucketPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketPermissions

`func (o *CreateStorageTokenRequest) SetBucketPermissions(v CreateStorageTokenRequestBucketPermissions)`

SetBucketPermissions sets BucketPermissions field to given value.

### HasBucketPermissions

`func (o *CreateStorageTokenRequest) HasBucketPermissions() bool`

HasBucketPermissions returns a boolean if a field has been set.

### GetComponentAccess

`func (o *CreateStorageTokenRequest) GetComponentAccess() string`

GetComponentAccess returns the ComponentAccess field if non-nil, zero value otherwise.

### GetComponentAccessOk

`func (o *CreateStorageTokenRequest) GetComponentAccessOk() (*string, bool)`

GetComponentAccessOk returns a tuple with the ComponentAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentAccess

`func (o *CreateStorageTokenRequest) SetComponentAccess(v string)`

SetComponentAccess sets ComponentAccess field to given value.

### HasComponentAccess

`func (o *CreateStorageTokenRequest) HasComponentAccess() bool`

HasComponentAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


