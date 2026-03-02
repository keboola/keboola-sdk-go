# ActivateSnowflakeStorageBackendResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Warehouse** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**TechnicalOwner** | Pointer to **string** |  | [optional] 
**TechnicalOwnerContactEmails** | Pointer to **[]string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**CreatorName** | Pointer to **string** |  | [optional] 
**CreatorId** | Pointer to **float32** |  | [optional] 
**IsMaintenance** | Pointer to **bool** |  | [optional] 
**UseDynamicBackends** | Pointer to **bool** |  | [optional] 
**UseNetworkPolicies** | Pointer to **bool** |  | [optional] 
**UseSso** | Pointer to **bool** |  | [optional] 
**IsSsoEnabled** | Pointer to **bool** |  | [optional] 
**IsSsoConfigured** | Pointer to **bool** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewActivateSnowflakeStorageBackendResponse

`func NewActivateSnowflakeStorageBackendResponse() *ActivateSnowflakeStorageBackendResponse`

NewActivateSnowflakeStorageBackendResponse instantiates a new ActivateSnowflakeStorageBackendResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivateSnowflakeStorageBackendResponseWithDefaults

`func NewActivateSnowflakeStorageBackendResponseWithDefaults() *ActivateSnowflakeStorageBackendResponse`

NewActivateSnowflakeStorageBackendResponseWithDefaults instantiates a new ActivateSnowflakeStorageBackendResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActivateSnowflakeStorageBackendResponse) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivateSnowflakeStorageBackendResponse) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivateSnowflakeStorageBackendResponse) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *ActivateSnowflakeStorageBackendResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHost

`func (o *ActivateSnowflakeStorageBackendResponse) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ActivateSnowflakeStorageBackendResponse) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ActivateSnowflakeStorageBackendResponse) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ActivateSnowflakeStorageBackendResponse) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetWarehouse

`func (o *ActivateSnowflakeStorageBackendResponse) GetWarehouse() string`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *ActivateSnowflakeStorageBackendResponse) GetWarehouseOk() (*string, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *ActivateSnowflakeStorageBackendResponse) SetWarehouse(v string)`

SetWarehouse sets Warehouse field to given value.

### HasWarehouse

`func (o *ActivateSnowflakeStorageBackendResponse) HasWarehouse() bool`

HasWarehouse returns a boolean if a field has been set.

### GetUsername

`func (o *ActivateSnowflakeStorageBackendResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ActivateSnowflakeStorageBackendResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ActivateSnowflakeStorageBackendResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ActivateSnowflakeStorageBackendResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetOwner

`func (o *ActivateSnowflakeStorageBackendResponse) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ActivateSnowflakeStorageBackendResponse) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ActivateSnowflakeStorageBackendResponse) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ActivateSnowflakeStorageBackendResponse) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTechnicalOwner

`func (o *ActivateSnowflakeStorageBackendResponse) GetTechnicalOwner() string`

GetTechnicalOwner returns the TechnicalOwner field if non-nil, zero value otherwise.

### GetTechnicalOwnerOk

`func (o *ActivateSnowflakeStorageBackendResponse) GetTechnicalOwnerOk() (*string, bool)`

GetTechnicalOwnerOk returns a tuple with the TechnicalOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalOwner

`func (o *ActivateSnowflakeStorageBackendResponse) SetTechnicalOwner(v string)`

SetTechnicalOwner sets TechnicalOwner field to given value.

### HasTechnicalOwner

`func (o *ActivateSnowflakeStorageBackendResponse) HasTechnicalOwner() bool`

HasTechnicalOwner returns a boolean if a field has been set.

### GetTechnicalOwnerContactEmails

`func (o *ActivateSnowflakeStorageBackendResponse) GetTechnicalOwnerContactEmails() []string`

GetTechnicalOwnerContactEmails returns the TechnicalOwnerContactEmails field if non-nil, zero value otherwise.

### GetTechnicalOwnerContactEmailsOk

`func (o *ActivateSnowflakeStorageBackendResponse) GetTechnicalOwnerContactEmailsOk() (*[]string, bool)`

GetTechnicalOwnerContactEmailsOk returns a tuple with the TechnicalOwnerContactEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalOwnerContactEmails

`func (o *ActivateSnowflakeStorageBackendResponse) SetTechnicalOwnerContactEmails(v []string)`

SetTechnicalOwnerContactEmails sets TechnicalOwnerContactEmails field to given value.

### HasTechnicalOwnerContactEmails

`func (o *ActivateSnowflakeStorageBackendResponse) HasTechnicalOwnerContactEmails() bool`

HasTechnicalOwnerContactEmails returns a boolean if a field has been set.

### GetRegion

`func (o *ActivateSnowflakeStorageBackendResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ActivateSnowflakeStorageBackendResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ActivateSnowflakeStorageBackendResponse) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ActivateSnowflakeStorageBackendResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCreated

`func (o *ActivateSnowflakeStorageBackendResponse) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ActivateSnowflakeStorageBackendResponse) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ActivateSnowflakeStorageBackendResponse) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ActivateSnowflakeStorageBackendResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatorName

`func (o *ActivateSnowflakeStorageBackendResponse) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *ActivateSnowflakeStorageBackendResponse) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *ActivateSnowflakeStorageBackendResponse) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.

### HasCreatorName

`func (o *ActivateSnowflakeStorageBackendResponse) HasCreatorName() bool`

HasCreatorName returns a boolean if a field has been set.

### GetCreatorId

`func (o *ActivateSnowflakeStorageBackendResponse) GetCreatorId() float32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *ActivateSnowflakeStorageBackendResponse) GetCreatorIdOk() (*float32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *ActivateSnowflakeStorageBackendResponse) SetCreatorId(v float32)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *ActivateSnowflakeStorageBackendResponse) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetIsMaintenance

`func (o *ActivateSnowflakeStorageBackendResponse) GetIsMaintenance() bool`

GetIsMaintenance returns the IsMaintenance field if non-nil, zero value otherwise.

### GetIsMaintenanceOk

`func (o *ActivateSnowflakeStorageBackendResponse) GetIsMaintenanceOk() (*bool, bool)`

GetIsMaintenanceOk returns a tuple with the IsMaintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaintenance

`func (o *ActivateSnowflakeStorageBackendResponse) SetIsMaintenance(v bool)`

SetIsMaintenance sets IsMaintenance field to given value.

### HasIsMaintenance

`func (o *ActivateSnowflakeStorageBackendResponse) HasIsMaintenance() bool`

HasIsMaintenance returns a boolean if a field has been set.

### GetUseDynamicBackends

`func (o *ActivateSnowflakeStorageBackendResponse) GetUseDynamicBackends() bool`

GetUseDynamicBackends returns the UseDynamicBackends field if non-nil, zero value otherwise.

### GetUseDynamicBackendsOk

`func (o *ActivateSnowflakeStorageBackendResponse) GetUseDynamicBackendsOk() (*bool, bool)`

GetUseDynamicBackendsOk returns a tuple with the UseDynamicBackends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDynamicBackends

`func (o *ActivateSnowflakeStorageBackendResponse) SetUseDynamicBackends(v bool)`

SetUseDynamicBackends sets UseDynamicBackends field to given value.

### HasUseDynamicBackends

`func (o *ActivateSnowflakeStorageBackendResponse) HasUseDynamicBackends() bool`

HasUseDynamicBackends returns a boolean if a field has been set.

### GetUseNetworkPolicies

`func (o *ActivateSnowflakeStorageBackendResponse) GetUseNetworkPolicies() bool`

GetUseNetworkPolicies returns the UseNetworkPolicies field if non-nil, zero value otherwise.

### GetUseNetworkPoliciesOk

`func (o *ActivateSnowflakeStorageBackendResponse) GetUseNetworkPoliciesOk() (*bool, bool)`

GetUseNetworkPoliciesOk returns a tuple with the UseNetworkPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNetworkPolicies

`func (o *ActivateSnowflakeStorageBackendResponse) SetUseNetworkPolicies(v bool)`

SetUseNetworkPolicies sets UseNetworkPolicies field to given value.

### HasUseNetworkPolicies

`func (o *ActivateSnowflakeStorageBackendResponse) HasUseNetworkPolicies() bool`

HasUseNetworkPolicies returns a boolean if a field has been set.

### GetUseSso

`func (o *ActivateSnowflakeStorageBackendResponse) GetUseSso() bool`

GetUseSso returns the UseSso field if non-nil, zero value otherwise.

### GetUseSsoOk

`func (o *ActivateSnowflakeStorageBackendResponse) GetUseSsoOk() (*bool, bool)`

GetUseSsoOk returns a tuple with the UseSso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSso

`func (o *ActivateSnowflakeStorageBackendResponse) SetUseSso(v bool)`

SetUseSso sets UseSso field to given value.

### HasUseSso

`func (o *ActivateSnowflakeStorageBackendResponse) HasUseSso() bool`

HasUseSso returns a boolean if a field has been set.

### GetIsSsoEnabled

`func (o *ActivateSnowflakeStorageBackendResponse) GetIsSsoEnabled() bool`

GetIsSsoEnabled returns the IsSsoEnabled field if non-nil, zero value otherwise.

### GetIsSsoEnabledOk

`func (o *ActivateSnowflakeStorageBackendResponse) GetIsSsoEnabledOk() (*bool, bool)`

GetIsSsoEnabledOk returns a tuple with the IsSsoEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSsoEnabled

`func (o *ActivateSnowflakeStorageBackendResponse) SetIsSsoEnabled(v bool)`

SetIsSsoEnabled sets IsSsoEnabled field to given value.

### HasIsSsoEnabled

`func (o *ActivateSnowflakeStorageBackendResponse) HasIsSsoEnabled() bool`

HasIsSsoEnabled returns a boolean if a field has been set.

### GetIsSsoConfigured

`func (o *ActivateSnowflakeStorageBackendResponse) GetIsSsoConfigured() bool`

GetIsSsoConfigured returns the IsSsoConfigured field if non-nil, zero value otherwise.

### GetIsSsoConfiguredOk

`func (o *ActivateSnowflakeStorageBackendResponse) GetIsSsoConfiguredOk() (*bool, bool)`

GetIsSsoConfiguredOk returns a tuple with the IsSsoConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSsoConfigured

`func (o *ActivateSnowflakeStorageBackendResponse) SetIsSsoConfigured(v bool)`

SetIsSsoConfigured sets IsSsoConfigured field to given value.

### HasIsSsoConfigured

`func (o *ActivateSnowflakeStorageBackendResponse) HasIsSsoConfigured() bool`

HasIsSsoConfigured returns a boolean if a field has been set.

### GetIsEnabled

`func (o *ActivateSnowflakeStorageBackendResponse) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ActivateSnowflakeStorageBackendResponse) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ActivateSnowflakeStorageBackendResponse) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *ActivateSnowflakeStorageBackendResponse) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


