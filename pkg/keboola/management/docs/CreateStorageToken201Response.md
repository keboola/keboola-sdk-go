# CreateStorageToken201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 
**IsMasterToken** | Pointer to **bool** |  | [optional] 
**CanManageBuckets** | Pointer to **bool** |  | [optional] 
**CanManageTokens** | Pointer to **bool** |  | [optional] 
**CanReadAllFileUploads** | Pointer to **bool** |  | [optional] 
**CanPurgeTrash** | Pointer to **bool** |  | [optional] 
**Expires** | Pointer to **interface{}** |  | [optional] 
**IsExpired** | Pointer to **bool** |  | [optional] 
**IsDisabled** | Pointer to **bool** |  | [optional] 
**DailyCapacity** | Pointer to **float32** |  | [optional] 
**BucketPermissions** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCreateStorageToken201Response

`func NewCreateStorageToken201Response() *CreateStorageToken201Response`

NewCreateStorageToken201Response instantiates a new CreateStorageToken201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStorageToken201ResponseWithDefaults

`func NewCreateStorageToken201ResponseWithDefaults() *CreateStorageToken201Response`

NewCreateStorageToken201ResponseWithDefaults instantiates a new CreateStorageToken201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateStorageToken201Response) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateStorageToken201Response) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateStorageToken201Response) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *CreateStorageToken201Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetToken

`func (o *CreateStorageToken201Response) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateStorageToken201Response) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateStorageToken201Response) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateStorageToken201Response) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetCreated

`func (o *CreateStorageToken201Response) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CreateStorageToken201Response) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CreateStorageToken201Response) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CreateStorageToken201Response) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *CreateStorageToken201Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateStorageToken201Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateStorageToken201Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateStorageToken201Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUri

`func (o *CreateStorageToken201Response) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *CreateStorageToken201Response) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *CreateStorageToken201Response) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *CreateStorageToken201Response) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetIsMasterToken

`func (o *CreateStorageToken201Response) GetIsMasterToken() bool`

GetIsMasterToken returns the IsMasterToken field if non-nil, zero value otherwise.

### GetIsMasterTokenOk

`func (o *CreateStorageToken201Response) GetIsMasterTokenOk() (*bool, bool)`

GetIsMasterTokenOk returns a tuple with the IsMasterToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMasterToken

`func (o *CreateStorageToken201Response) SetIsMasterToken(v bool)`

SetIsMasterToken sets IsMasterToken field to given value.

### HasIsMasterToken

`func (o *CreateStorageToken201Response) HasIsMasterToken() bool`

HasIsMasterToken returns a boolean if a field has been set.

### GetCanManageBuckets

`func (o *CreateStorageToken201Response) GetCanManageBuckets() bool`

GetCanManageBuckets returns the CanManageBuckets field if non-nil, zero value otherwise.

### GetCanManageBucketsOk

`func (o *CreateStorageToken201Response) GetCanManageBucketsOk() (*bool, bool)`

GetCanManageBucketsOk returns a tuple with the CanManageBuckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageBuckets

`func (o *CreateStorageToken201Response) SetCanManageBuckets(v bool)`

SetCanManageBuckets sets CanManageBuckets field to given value.

### HasCanManageBuckets

`func (o *CreateStorageToken201Response) HasCanManageBuckets() bool`

HasCanManageBuckets returns a boolean if a field has been set.

### GetCanManageTokens

`func (o *CreateStorageToken201Response) GetCanManageTokens() bool`

GetCanManageTokens returns the CanManageTokens field if non-nil, zero value otherwise.

### GetCanManageTokensOk

`func (o *CreateStorageToken201Response) GetCanManageTokensOk() (*bool, bool)`

GetCanManageTokensOk returns a tuple with the CanManageTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageTokens

`func (o *CreateStorageToken201Response) SetCanManageTokens(v bool)`

SetCanManageTokens sets CanManageTokens field to given value.

### HasCanManageTokens

`func (o *CreateStorageToken201Response) HasCanManageTokens() bool`

HasCanManageTokens returns a boolean if a field has been set.

### GetCanReadAllFileUploads

`func (o *CreateStorageToken201Response) GetCanReadAllFileUploads() bool`

GetCanReadAllFileUploads returns the CanReadAllFileUploads field if non-nil, zero value otherwise.

### GetCanReadAllFileUploadsOk

`func (o *CreateStorageToken201Response) GetCanReadAllFileUploadsOk() (*bool, bool)`

GetCanReadAllFileUploadsOk returns a tuple with the CanReadAllFileUploads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanReadAllFileUploads

`func (o *CreateStorageToken201Response) SetCanReadAllFileUploads(v bool)`

SetCanReadAllFileUploads sets CanReadAllFileUploads field to given value.

### HasCanReadAllFileUploads

`func (o *CreateStorageToken201Response) HasCanReadAllFileUploads() bool`

HasCanReadAllFileUploads returns a boolean if a field has been set.

### GetCanPurgeTrash

`func (o *CreateStorageToken201Response) GetCanPurgeTrash() bool`

GetCanPurgeTrash returns the CanPurgeTrash field if non-nil, zero value otherwise.

### GetCanPurgeTrashOk

`func (o *CreateStorageToken201Response) GetCanPurgeTrashOk() (*bool, bool)`

GetCanPurgeTrashOk returns a tuple with the CanPurgeTrash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanPurgeTrash

`func (o *CreateStorageToken201Response) SetCanPurgeTrash(v bool)`

SetCanPurgeTrash sets CanPurgeTrash field to given value.

### HasCanPurgeTrash

`func (o *CreateStorageToken201Response) HasCanPurgeTrash() bool`

HasCanPurgeTrash returns a boolean if a field has been set.

### GetExpires

`func (o *CreateStorageToken201Response) GetExpires() interface{}`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *CreateStorageToken201Response) GetExpiresOk() (*interface{}, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *CreateStorageToken201Response) SetExpires(v interface{})`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *CreateStorageToken201Response) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *CreateStorageToken201Response) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *CreateStorageToken201Response) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetIsExpired

`func (o *CreateStorageToken201Response) GetIsExpired() bool`

GetIsExpired returns the IsExpired field if non-nil, zero value otherwise.

### GetIsExpiredOk

`func (o *CreateStorageToken201Response) GetIsExpiredOk() (*bool, bool)`

GetIsExpiredOk returns a tuple with the IsExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpired

`func (o *CreateStorageToken201Response) SetIsExpired(v bool)`

SetIsExpired sets IsExpired field to given value.

### HasIsExpired

`func (o *CreateStorageToken201Response) HasIsExpired() bool`

HasIsExpired returns a boolean if a field has been set.

### GetIsDisabled

`func (o *CreateStorageToken201Response) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *CreateStorageToken201Response) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *CreateStorageToken201Response) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *CreateStorageToken201Response) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetDailyCapacity

`func (o *CreateStorageToken201Response) GetDailyCapacity() float32`

GetDailyCapacity returns the DailyCapacity field if non-nil, zero value otherwise.

### GetDailyCapacityOk

`func (o *CreateStorageToken201Response) GetDailyCapacityOk() (*float32, bool)`

GetDailyCapacityOk returns a tuple with the DailyCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyCapacity

`func (o *CreateStorageToken201Response) SetDailyCapacity(v float32)`

SetDailyCapacity sets DailyCapacity field to given value.

### HasDailyCapacity

`func (o *CreateStorageToken201Response) HasDailyCapacity() bool`

HasDailyCapacity returns a boolean if a field has been set.

### GetBucketPermissions

`func (o *CreateStorageToken201Response) GetBucketPermissions() map[string]interface{}`

GetBucketPermissions returns the BucketPermissions field if non-nil, zero value otherwise.

### GetBucketPermissionsOk

`func (o *CreateStorageToken201Response) GetBucketPermissionsOk() (*map[string]interface{}, bool)`

GetBucketPermissionsOk returns a tuple with the BucketPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketPermissions

`func (o *CreateStorageToken201Response) SetBucketPermissions(v map[string]interface{})`

SetBucketPermissions sets BucketPermissions field to given value.

### HasBucketPermissions

`func (o *CreateStorageToken201Response) HasBucketPermissions() bool`

HasBucketPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


