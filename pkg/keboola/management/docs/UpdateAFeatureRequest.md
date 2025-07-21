# UpdateAFeatureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Feature title (display name) | [optional] 
**Description** | Pointer to **string** | Short description of the feature | [optional] 
**CanBeManageByAdmin** | Pointer to **bool** | If true, the feature can be assigned by a user without the superadmin role | [optional] 
**CanBeManagedViaAPI** | Pointer to **bool** | Enables or disables the ability to assign the project/admin feature using the API | [optional] 

## Methods

### NewUpdateAFeatureRequest

`func NewUpdateAFeatureRequest() *UpdateAFeatureRequest`

NewUpdateAFeatureRequest instantiates a new UpdateAFeatureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAFeatureRequestWithDefaults

`func NewUpdateAFeatureRequestWithDefaults() *UpdateAFeatureRequest`

NewUpdateAFeatureRequestWithDefaults instantiates a new UpdateAFeatureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *UpdateAFeatureRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateAFeatureRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateAFeatureRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateAFeatureRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateAFeatureRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateAFeatureRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateAFeatureRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateAFeatureRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCanBeManageByAdmin

`func (o *UpdateAFeatureRequest) GetCanBeManageByAdmin() bool`

GetCanBeManageByAdmin returns the CanBeManageByAdmin field if non-nil, zero value otherwise.

### GetCanBeManageByAdminOk

`func (o *UpdateAFeatureRequest) GetCanBeManageByAdminOk() (*bool, bool)`

GetCanBeManageByAdminOk returns a tuple with the CanBeManageByAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanBeManageByAdmin

`func (o *UpdateAFeatureRequest) SetCanBeManageByAdmin(v bool)`

SetCanBeManageByAdmin sets CanBeManageByAdmin field to given value.

### HasCanBeManageByAdmin

`func (o *UpdateAFeatureRequest) HasCanBeManageByAdmin() bool`

HasCanBeManageByAdmin returns a boolean if a field has been set.

### GetCanBeManagedViaAPI

`func (o *UpdateAFeatureRequest) GetCanBeManagedViaAPI() bool`

GetCanBeManagedViaAPI returns the CanBeManagedViaAPI field if non-nil, zero value otherwise.

### GetCanBeManagedViaAPIOk

`func (o *UpdateAFeatureRequest) GetCanBeManagedViaAPIOk() (*bool, bool)`

GetCanBeManagedViaAPIOk returns a tuple with the CanBeManagedViaAPI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanBeManagedViaAPI

`func (o *UpdateAFeatureRequest) SetCanBeManagedViaAPI(v bool)`

SetCanBeManagedViaAPI sets CanBeManagedViaAPI field to given value.

### HasCanBeManagedViaAPI

`func (o *UpdateAFeatureRequest) HasCanBeManagedViaAPI() bool`

HasCanBeManagedViaAPI returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


