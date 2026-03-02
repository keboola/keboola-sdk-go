# CreateSnowflakeStorageBackendResponse

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
**Edition** | Pointer to **string** |  | [optional] 
**SqlTemplate** | Pointer to **string** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**UserPublicKey** | Pointer to **string** |  | [optional] 
**SecurityIntegrationKey** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateSnowflakeStorageBackendResponse

`func NewCreateSnowflakeStorageBackendResponse() *CreateSnowflakeStorageBackendResponse`

NewCreateSnowflakeStorageBackendResponse instantiates a new CreateSnowflakeStorageBackendResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSnowflakeStorageBackendResponseWithDefaults

`func NewCreateSnowflakeStorageBackendResponseWithDefaults() *CreateSnowflakeStorageBackendResponse`

NewCreateSnowflakeStorageBackendResponseWithDefaults instantiates a new CreateSnowflakeStorageBackendResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateSnowflakeStorageBackendResponse) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateSnowflakeStorageBackendResponse) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateSnowflakeStorageBackendResponse) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *CreateSnowflakeStorageBackendResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHost

`func (o *CreateSnowflakeStorageBackendResponse) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CreateSnowflakeStorageBackendResponse) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CreateSnowflakeStorageBackendResponse) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *CreateSnowflakeStorageBackendResponse) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetWarehouse

`func (o *CreateSnowflakeStorageBackendResponse) GetWarehouse() string`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *CreateSnowflakeStorageBackendResponse) GetWarehouseOk() (*string, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *CreateSnowflakeStorageBackendResponse) SetWarehouse(v string)`

SetWarehouse sets Warehouse field to given value.

### HasWarehouse

`func (o *CreateSnowflakeStorageBackendResponse) HasWarehouse() bool`

HasWarehouse returns a boolean if a field has been set.

### GetUsername

`func (o *CreateSnowflakeStorageBackendResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateSnowflakeStorageBackendResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateSnowflakeStorageBackendResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateSnowflakeStorageBackendResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetOwner

`func (o *CreateSnowflakeStorageBackendResponse) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateSnowflakeStorageBackendResponse) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateSnowflakeStorageBackendResponse) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CreateSnowflakeStorageBackendResponse) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTechnicalOwner

`func (o *CreateSnowflakeStorageBackendResponse) GetTechnicalOwner() string`

GetTechnicalOwner returns the TechnicalOwner field if non-nil, zero value otherwise.

### GetTechnicalOwnerOk

`func (o *CreateSnowflakeStorageBackendResponse) GetTechnicalOwnerOk() (*string, bool)`

GetTechnicalOwnerOk returns a tuple with the TechnicalOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalOwner

`func (o *CreateSnowflakeStorageBackendResponse) SetTechnicalOwner(v string)`

SetTechnicalOwner sets TechnicalOwner field to given value.

### HasTechnicalOwner

`func (o *CreateSnowflakeStorageBackendResponse) HasTechnicalOwner() bool`

HasTechnicalOwner returns a boolean if a field has been set.

### GetTechnicalOwnerContactEmails

`func (o *CreateSnowflakeStorageBackendResponse) GetTechnicalOwnerContactEmails() []string`

GetTechnicalOwnerContactEmails returns the TechnicalOwnerContactEmails field if non-nil, zero value otherwise.

### GetTechnicalOwnerContactEmailsOk

`func (o *CreateSnowflakeStorageBackendResponse) GetTechnicalOwnerContactEmailsOk() (*[]string, bool)`

GetTechnicalOwnerContactEmailsOk returns a tuple with the TechnicalOwnerContactEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalOwnerContactEmails

`func (o *CreateSnowflakeStorageBackendResponse) SetTechnicalOwnerContactEmails(v []string)`

SetTechnicalOwnerContactEmails sets TechnicalOwnerContactEmails field to given value.

### HasTechnicalOwnerContactEmails

`func (o *CreateSnowflakeStorageBackendResponse) HasTechnicalOwnerContactEmails() bool`

HasTechnicalOwnerContactEmails returns a boolean if a field has been set.

### GetRegion

`func (o *CreateSnowflakeStorageBackendResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateSnowflakeStorageBackendResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateSnowflakeStorageBackendResponse) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CreateSnowflakeStorageBackendResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCreated

`func (o *CreateSnowflakeStorageBackendResponse) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CreateSnowflakeStorageBackendResponse) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CreateSnowflakeStorageBackendResponse) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CreateSnowflakeStorageBackendResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatorName

`func (o *CreateSnowflakeStorageBackendResponse) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *CreateSnowflakeStorageBackendResponse) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *CreateSnowflakeStorageBackendResponse) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.

### HasCreatorName

`func (o *CreateSnowflakeStorageBackendResponse) HasCreatorName() bool`

HasCreatorName returns a boolean if a field has been set.

### GetCreatorId

`func (o *CreateSnowflakeStorageBackendResponse) GetCreatorId() float32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *CreateSnowflakeStorageBackendResponse) GetCreatorIdOk() (*float32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *CreateSnowflakeStorageBackendResponse) SetCreatorId(v float32)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *CreateSnowflakeStorageBackendResponse) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetIsMaintenance

`func (o *CreateSnowflakeStorageBackendResponse) GetIsMaintenance() bool`

GetIsMaintenance returns the IsMaintenance field if non-nil, zero value otherwise.

### GetIsMaintenanceOk

`func (o *CreateSnowflakeStorageBackendResponse) GetIsMaintenanceOk() (*bool, bool)`

GetIsMaintenanceOk returns a tuple with the IsMaintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaintenance

`func (o *CreateSnowflakeStorageBackendResponse) SetIsMaintenance(v bool)`

SetIsMaintenance sets IsMaintenance field to given value.

### HasIsMaintenance

`func (o *CreateSnowflakeStorageBackendResponse) HasIsMaintenance() bool`

HasIsMaintenance returns a boolean if a field has been set.

### GetUseDynamicBackends

`func (o *CreateSnowflakeStorageBackendResponse) GetUseDynamicBackends() bool`

GetUseDynamicBackends returns the UseDynamicBackends field if non-nil, zero value otherwise.

### GetUseDynamicBackendsOk

`func (o *CreateSnowflakeStorageBackendResponse) GetUseDynamicBackendsOk() (*bool, bool)`

GetUseDynamicBackendsOk returns a tuple with the UseDynamicBackends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDynamicBackends

`func (o *CreateSnowflakeStorageBackendResponse) SetUseDynamicBackends(v bool)`

SetUseDynamicBackends sets UseDynamicBackends field to given value.

### HasUseDynamicBackends

`func (o *CreateSnowflakeStorageBackendResponse) HasUseDynamicBackends() bool`

HasUseDynamicBackends returns a boolean if a field has been set.

### GetUseNetworkPolicies

`func (o *CreateSnowflakeStorageBackendResponse) GetUseNetworkPolicies() bool`

GetUseNetworkPolicies returns the UseNetworkPolicies field if non-nil, zero value otherwise.

### GetUseNetworkPoliciesOk

`func (o *CreateSnowflakeStorageBackendResponse) GetUseNetworkPoliciesOk() (*bool, bool)`

GetUseNetworkPoliciesOk returns a tuple with the UseNetworkPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNetworkPolicies

`func (o *CreateSnowflakeStorageBackendResponse) SetUseNetworkPolicies(v bool)`

SetUseNetworkPolicies sets UseNetworkPolicies field to given value.

### HasUseNetworkPolicies

`func (o *CreateSnowflakeStorageBackendResponse) HasUseNetworkPolicies() bool`

HasUseNetworkPolicies returns a boolean if a field has been set.

### GetUseSso

`func (o *CreateSnowflakeStorageBackendResponse) GetUseSso() bool`

GetUseSso returns the UseSso field if non-nil, zero value otherwise.

### GetUseSsoOk

`func (o *CreateSnowflakeStorageBackendResponse) GetUseSsoOk() (*bool, bool)`

GetUseSsoOk returns a tuple with the UseSso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSso

`func (o *CreateSnowflakeStorageBackendResponse) SetUseSso(v bool)`

SetUseSso sets UseSso field to given value.

### HasUseSso

`func (o *CreateSnowflakeStorageBackendResponse) HasUseSso() bool`

HasUseSso returns a boolean if a field has been set.

### GetIsSsoEnabled

`func (o *CreateSnowflakeStorageBackendResponse) GetIsSsoEnabled() bool`

GetIsSsoEnabled returns the IsSsoEnabled field if non-nil, zero value otherwise.

### GetIsSsoEnabledOk

`func (o *CreateSnowflakeStorageBackendResponse) GetIsSsoEnabledOk() (*bool, bool)`

GetIsSsoEnabledOk returns a tuple with the IsSsoEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSsoEnabled

`func (o *CreateSnowflakeStorageBackendResponse) SetIsSsoEnabled(v bool)`

SetIsSsoEnabled sets IsSsoEnabled field to given value.

### HasIsSsoEnabled

`func (o *CreateSnowflakeStorageBackendResponse) HasIsSsoEnabled() bool`

HasIsSsoEnabled returns a boolean if a field has been set.

### GetIsSsoConfigured

`func (o *CreateSnowflakeStorageBackendResponse) GetIsSsoConfigured() bool`

GetIsSsoConfigured returns the IsSsoConfigured field if non-nil, zero value otherwise.

### GetIsSsoConfiguredOk

`func (o *CreateSnowflakeStorageBackendResponse) GetIsSsoConfiguredOk() (*bool, bool)`

GetIsSsoConfiguredOk returns a tuple with the IsSsoConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSsoConfigured

`func (o *CreateSnowflakeStorageBackendResponse) SetIsSsoConfigured(v bool)`

SetIsSsoConfigured sets IsSsoConfigured field to given value.

### HasIsSsoConfigured

`func (o *CreateSnowflakeStorageBackendResponse) HasIsSsoConfigured() bool`

HasIsSsoConfigured returns a boolean if a field has been set.

### GetEdition

`func (o *CreateSnowflakeStorageBackendResponse) GetEdition() string`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *CreateSnowflakeStorageBackendResponse) GetEditionOk() (*string, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *CreateSnowflakeStorageBackendResponse) SetEdition(v string)`

SetEdition sets Edition field to given value.

### HasEdition

`func (o *CreateSnowflakeStorageBackendResponse) HasEdition() bool`

HasEdition returns a boolean if a field has been set.

### GetSqlTemplate

`func (o *CreateSnowflakeStorageBackendResponse) GetSqlTemplate() string`

GetSqlTemplate returns the SqlTemplate field if non-nil, zero value otherwise.

### GetSqlTemplateOk

`func (o *CreateSnowflakeStorageBackendResponse) GetSqlTemplateOk() (*string, bool)`

GetSqlTemplateOk returns a tuple with the SqlTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlTemplate

`func (o *CreateSnowflakeStorageBackendResponse) SetSqlTemplate(v string)`

SetSqlTemplate sets SqlTemplate field to given value.

### HasSqlTemplate

`func (o *CreateSnowflakeStorageBackendResponse) HasSqlTemplate() bool`

HasSqlTemplate returns a boolean if a field has been set.

### GetIsEnabled

`func (o *CreateSnowflakeStorageBackendResponse) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *CreateSnowflakeStorageBackendResponse) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *CreateSnowflakeStorageBackendResponse) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *CreateSnowflakeStorageBackendResponse) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetUserPublicKey

`func (o *CreateSnowflakeStorageBackendResponse) GetUserPublicKey() string`

GetUserPublicKey returns the UserPublicKey field if non-nil, zero value otherwise.

### GetUserPublicKeyOk

`func (o *CreateSnowflakeStorageBackendResponse) GetUserPublicKeyOk() (*string, bool)`

GetUserPublicKeyOk returns a tuple with the UserPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPublicKey

`func (o *CreateSnowflakeStorageBackendResponse) SetUserPublicKey(v string)`

SetUserPublicKey sets UserPublicKey field to given value.

### HasUserPublicKey

`func (o *CreateSnowflakeStorageBackendResponse) HasUserPublicKey() bool`

HasUserPublicKey returns a boolean if a field has been set.

### GetSecurityIntegrationKey

`func (o *CreateSnowflakeStorageBackendResponse) GetSecurityIntegrationKey() string`

GetSecurityIntegrationKey returns the SecurityIntegrationKey field if non-nil, zero value otherwise.

### GetSecurityIntegrationKeyOk

`func (o *CreateSnowflakeStorageBackendResponse) GetSecurityIntegrationKeyOk() (*string, bool)`

GetSecurityIntegrationKeyOk returns a tuple with the SecurityIntegrationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityIntegrationKey

`func (o *CreateSnowflakeStorageBackendResponse) SetSecurityIntegrationKey(v string)`

SetSecurityIntegrationKey sets SecurityIntegrationKey field to given value.

### HasSecurityIntegrationKey

`func (o *CreateSnowflakeStorageBackendResponse) HasSecurityIntegrationKey() bool`

HasSecurityIntegrationKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


