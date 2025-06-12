# OrganizationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**AllowAutoJoin** | Pointer to **bool** |  | [optional] 
**CrmId** | Pointer to **string** |  | [optional] 
**ActivityCenterProjectId** | Pointer to **float32** |  | [optional] 
**MfaRequired** | Pointer to **bool** |  | [optional] 
**Projects** | Pointer to [**[]RetrieveAnOrganization200ResponseProjectsInner**](RetrieveAnOrganization200ResponseProjectsInner.md) |  | [optional] 

## Methods

### NewOrganizationModel

`func NewOrganizationModel() *OrganizationModel`

NewOrganizationModel instantiates a new OrganizationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationModelWithDefaults

`func NewOrganizationModelWithDefaults() *OrganizationModel`

NewOrganizationModelWithDefaults instantiates a new OrganizationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationModel) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationModel) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationModel) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *OrganizationModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OrganizationModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreated

`func (o *OrganizationModel) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *OrganizationModel) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *OrganizationModel) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *OrganizationModel) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetAllowAutoJoin

`func (o *OrganizationModel) GetAllowAutoJoin() bool`

GetAllowAutoJoin returns the AllowAutoJoin field if non-nil, zero value otherwise.

### GetAllowAutoJoinOk

`func (o *OrganizationModel) GetAllowAutoJoinOk() (*bool, bool)`

GetAllowAutoJoinOk returns a tuple with the AllowAutoJoin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAutoJoin

`func (o *OrganizationModel) SetAllowAutoJoin(v bool)`

SetAllowAutoJoin sets AllowAutoJoin field to given value.

### HasAllowAutoJoin

`func (o *OrganizationModel) HasAllowAutoJoin() bool`

HasAllowAutoJoin returns a boolean if a field has been set.

### GetCrmId

`func (o *OrganizationModel) GetCrmId() string`

GetCrmId returns the CrmId field if non-nil, zero value otherwise.

### GetCrmIdOk

`func (o *OrganizationModel) GetCrmIdOk() (*string, bool)`

GetCrmIdOk returns a tuple with the CrmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrmId

`func (o *OrganizationModel) SetCrmId(v string)`

SetCrmId sets CrmId field to given value.

### HasCrmId

`func (o *OrganizationModel) HasCrmId() bool`

HasCrmId returns a boolean if a field has been set.

### GetActivityCenterProjectId

`func (o *OrganizationModel) GetActivityCenterProjectId() float32`

GetActivityCenterProjectId returns the ActivityCenterProjectId field if non-nil, zero value otherwise.

### GetActivityCenterProjectIdOk

`func (o *OrganizationModel) GetActivityCenterProjectIdOk() (*float32, bool)`

GetActivityCenterProjectIdOk returns a tuple with the ActivityCenterProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityCenterProjectId

`func (o *OrganizationModel) SetActivityCenterProjectId(v float32)`

SetActivityCenterProjectId sets ActivityCenterProjectId field to given value.

### HasActivityCenterProjectId

`func (o *OrganizationModel) HasActivityCenterProjectId() bool`

HasActivityCenterProjectId returns a boolean if a field has been set.

### GetMfaRequired

`func (o *OrganizationModel) GetMfaRequired() bool`

GetMfaRequired returns the MfaRequired field if non-nil, zero value otherwise.

### GetMfaRequiredOk

`func (o *OrganizationModel) GetMfaRequiredOk() (*bool, bool)`

GetMfaRequiredOk returns a tuple with the MfaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaRequired

`func (o *OrganizationModel) SetMfaRequired(v bool)`

SetMfaRequired sets MfaRequired field to given value.

### HasMfaRequired

`func (o *OrganizationModel) HasMfaRequired() bool`

HasMfaRequired returns a boolean if a field has been set.

### GetProjects

`func (o *OrganizationModel) GetProjects() []RetrieveAnOrganization200ResponseProjectsInner`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *OrganizationModel) GetProjectsOk() (*[]RetrieveAnOrganization200ResponseProjectsInner, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *OrganizationModel) SetProjects(v []RetrieveAnOrganization200ResponseProjectsInner)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *OrganizationModel) HasProjects() bool`

HasProjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


