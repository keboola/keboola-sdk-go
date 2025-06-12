# ListMyJoinRequests200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**Created** | Pointer to **string** |  | [optional] 
**Expires** | **NullableString** |  | 
**Reason** | **string** |  | 
**Project** | [**TokenVerification200ResponseCreator**](TokenVerification200ResponseCreator.md) |  | 
**Crated** | Pointer to **string** |  | [optional] 

## Methods

### NewListMyJoinRequests200ResponseInner

`func NewListMyJoinRequests200ResponseInner(id float32, expires NullableString, reason string, project TokenVerification200ResponseCreator, ) *ListMyJoinRequests200ResponseInner`

NewListMyJoinRequests200ResponseInner instantiates a new ListMyJoinRequests200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMyJoinRequests200ResponseInnerWithDefaults

`func NewListMyJoinRequests200ResponseInnerWithDefaults() *ListMyJoinRequests200ResponseInner`

NewListMyJoinRequests200ResponseInnerWithDefaults instantiates a new ListMyJoinRequests200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListMyJoinRequests200ResponseInner) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListMyJoinRequests200ResponseInner) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListMyJoinRequests200ResponseInner) SetId(v float32)`

SetId sets Id field to given value.


### GetCreated

`func (o *ListMyJoinRequests200ResponseInner) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ListMyJoinRequests200ResponseInner) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ListMyJoinRequests200ResponseInner) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ListMyJoinRequests200ResponseInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExpires

`func (o *ListMyJoinRequests200ResponseInner) GetExpires() string`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *ListMyJoinRequests200ResponseInner) GetExpiresOk() (*string, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *ListMyJoinRequests200ResponseInner) SetExpires(v string)`

SetExpires sets Expires field to given value.


### SetExpiresNil

`func (o *ListMyJoinRequests200ResponseInner) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *ListMyJoinRequests200ResponseInner) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetReason

`func (o *ListMyJoinRequests200ResponseInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ListMyJoinRequests200ResponseInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ListMyJoinRequests200ResponseInner) SetReason(v string)`

SetReason sets Reason field to given value.


### GetProject

`func (o *ListMyJoinRequests200ResponseInner) GetProject() TokenVerification200ResponseCreator`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ListMyJoinRequests200ResponseInner) GetProjectOk() (*TokenVerification200ResponseCreator, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ListMyJoinRequests200ResponseInner) SetProject(v TokenVerification200ResponseCreator)`

SetProject sets Project field to given value.


### GetCrated

`func (o *ListMyJoinRequests200ResponseInner) GetCrated() string`

GetCrated returns the Crated field if non-nil, zero value otherwise.

### GetCratedOk

`func (o *ListMyJoinRequests200ResponseInner) GetCratedOk() (*string, bool)`

GetCratedOk returns a tuple with the Crated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrated

`func (o *ListMyJoinRequests200ResponseInner) SetCrated(v string)`

SetCrated sets Crated field to given value.

### HasCrated

`func (o *ListMyJoinRequests200ResponseInner) HasCrated() bool`

HasCrated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


