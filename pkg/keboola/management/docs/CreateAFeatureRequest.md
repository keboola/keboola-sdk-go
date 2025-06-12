# CreateAFeatureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Feature name (programmatic name) | 
**Type** | Pointer to **string** |  | [optional] 
**Title** | **string** | Feature title (display name) | 
**Description** | **string** | Short description of the feature | 
**CanBeManageByAdmin** | Pointer to **bool** | If true, the feature can be assigned by a user without the superadmin role | [optional] 
**CanBeManagedViaAPI** | Pointer to **bool** | Enables or disables the ability to assign the project/admin feature using the API | [optional] 

## Methods

### NewCreateAFeatureRequest

`func NewCreateAFeatureRequest(name string, title string, description string, ) *CreateAFeatureRequest`

NewCreateAFeatureRequest instantiates a new CreateAFeatureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAFeatureRequestWithDefaults

`func NewCreateAFeatureRequestWithDefaults() *CreateAFeatureRequest`

NewCreateAFeatureRequestWithDefaults instantiates a new CreateAFeatureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAFeatureRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAFeatureRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAFeatureRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CreateAFeatureRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateAFeatureRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateAFeatureRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateAFeatureRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTitle

`func (o *CreateAFeatureRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateAFeatureRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateAFeatureRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *CreateAFeatureRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAFeatureRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAFeatureRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCanBeManageByAdmin

`func (o *CreateAFeatureRequest) GetCanBeManageByAdmin() bool`

GetCanBeManageByAdmin returns the CanBeManageByAdmin field if non-nil, zero value otherwise.

### GetCanBeManageByAdminOk

`func (o *CreateAFeatureRequest) GetCanBeManageByAdminOk() (*bool, bool)`

GetCanBeManageByAdminOk returns a tuple with the CanBeManageByAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanBeManageByAdmin

`func (o *CreateAFeatureRequest) SetCanBeManageByAdmin(v bool)`

SetCanBeManageByAdmin sets CanBeManageByAdmin field to given value.

### HasCanBeManageByAdmin

`func (o *CreateAFeatureRequest) HasCanBeManageByAdmin() bool`

HasCanBeManageByAdmin returns a boolean if a field has been set.

### GetCanBeManagedViaAPI

`func (o *CreateAFeatureRequest) GetCanBeManagedViaAPI() bool`

GetCanBeManagedViaAPI returns the CanBeManagedViaAPI field if non-nil, zero value otherwise.

### GetCanBeManagedViaAPIOk

`func (o *CreateAFeatureRequest) GetCanBeManagedViaAPIOk() (*bool, bool)`

GetCanBeManagedViaAPIOk returns a tuple with the CanBeManagedViaAPI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanBeManagedViaAPI

`func (o *CreateAFeatureRequest) SetCanBeManagedViaAPI(v bool)`

SetCanBeManagedViaAPI sets CanBeManagedViaAPI field to given value.

### HasCanBeManagedViaAPI

`func (o *CreateAFeatureRequest) HasCanBeManagedViaAPI() bool`

HasCanBeManagedViaAPI returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


