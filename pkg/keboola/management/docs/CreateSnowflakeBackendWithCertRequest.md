# CreateSnowflakeBackendWithCertRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | Snowflake host name | 
**Warehouse** | **string** | Snowflake warehouse name | 
**Username** | Pointer to **string** | Snowflake username | [optional] 
**Region** | **string** | Backend region | 
**Owner** | **string** | Associated legal owner (string, mostly: keboola, client-&lt;name&gt;) | 
**TechnicalOwner** | **string** | Associated technical owner. Enum: keboola, internal, kbdb, byodb | 
**TechnicalOwnerContactEmails** | Pointer to **[]string** | Array of technical owner contact emails | [optional] 
**UseDynamicBackends** | Pointer to **bool** | Enable dynamic backends | [optional] 
**UseNetworkPolicies** | Pointer to **bool** | Enable network policies | [optional] 
**UseSso** | Pointer to **bool** | Enable SSO | [optional] 
**Edition** | Pointer to **string** | Snowflake edition type. Enum: enterprise, standard | [optional] 

## Methods

### NewCreateSnowflakeBackendWithCertRequest

`func NewCreateSnowflakeBackendWithCertRequest(host string, warehouse string, region string, owner string, technicalOwner string, ) *CreateSnowflakeBackendWithCertRequest`

NewCreateSnowflakeBackendWithCertRequest instantiates a new CreateSnowflakeBackendWithCertRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSnowflakeBackendWithCertRequestWithDefaults

`func NewCreateSnowflakeBackendWithCertRequestWithDefaults() *CreateSnowflakeBackendWithCertRequest`

NewCreateSnowflakeBackendWithCertRequestWithDefaults instantiates a new CreateSnowflakeBackendWithCertRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *CreateSnowflakeBackendWithCertRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CreateSnowflakeBackendWithCertRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CreateSnowflakeBackendWithCertRequest) SetHost(v string)`

SetHost sets Host field to given value.


### GetWarehouse

`func (o *CreateSnowflakeBackendWithCertRequest) GetWarehouse() string`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *CreateSnowflakeBackendWithCertRequest) GetWarehouseOk() (*string, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *CreateSnowflakeBackendWithCertRequest) SetWarehouse(v string)`

SetWarehouse sets Warehouse field to given value.


### GetUsername

`func (o *CreateSnowflakeBackendWithCertRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateSnowflakeBackendWithCertRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateSnowflakeBackendWithCertRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateSnowflakeBackendWithCertRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetRegion

`func (o *CreateSnowflakeBackendWithCertRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateSnowflakeBackendWithCertRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateSnowflakeBackendWithCertRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetOwner

`func (o *CreateSnowflakeBackendWithCertRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateSnowflakeBackendWithCertRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateSnowflakeBackendWithCertRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetTechnicalOwner

`func (o *CreateSnowflakeBackendWithCertRequest) GetTechnicalOwner() string`

GetTechnicalOwner returns the TechnicalOwner field if non-nil, zero value otherwise.

### GetTechnicalOwnerOk

`func (o *CreateSnowflakeBackendWithCertRequest) GetTechnicalOwnerOk() (*string, bool)`

GetTechnicalOwnerOk returns a tuple with the TechnicalOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalOwner

`func (o *CreateSnowflakeBackendWithCertRequest) SetTechnicalOwner(v string)`

SetTechnicalOwner sets TechnicalOwner field to given value.


### GetTechnicalOwnerContactEmails

`func (o *CreateSnowflakeBackendWithCertRequest) GetTechnicalOwnerContactEmails() []string`

GetTechnicalOwnerContactEmails returns the TechnicalOwnerContactEmails field if non-nil, zero value otherwise.

### GetTechnicalOwnerContactEmailsOk

`func (o *CreateSnowflakeBackendWithCertRequest) GetTechnicalOwnerContactEmailsOk() (*[]string, bool)`

GetTechnicalOwnerContactEmailsOk returns a tuple with the TechnicalOwnerContactEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalOwnerContactEmails

`func (o *CreateSnowflakeBackendWithCertRequest) SetTechnicalOwnerContactEmails(v []string)`

SetTechnicalOwnerContactEmails sets TechnicalOwnerContactEmails field to given value.

### HasTechnicalOwnerContactEmails

`func (o *CreateSnowflakeBackendWithCertRequest) HasTechnicalOwnerContactEmails() bool`

HasTechnicalOwnerContactEmails returns a boolean if a field has been set.

### GetUseDynamicBackends

`func (o *CreateSnowflakeBackendWithCertRequest) GetUseDynamicBackends() bool`

GetUseDynamicBackends returns the UseDynamicBackends field if non-nil, zero value otherwise.

### GetUseDynamicBackendsOk

`func (o *CreateSnowflakeBackendWithCertRequest) GetUseDynamicBackendsOk() (*bool, bool)`

GetUseDynamicBackendsOk returns a tuple with the UseDynamicBackends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDynamicBackends

`func (o *CreateSnowflakeBackendWithCertRequest) SetUseDynamicBackends(v bool)`

SetUseDynamicBackends sets UseDynamicBackends field to given value.

### HasUseDynamicBackends

`func (o *CreateSnowflakeBackendWithCertRequest) HasUseDynamicBackends() bool`

HasUseDynamicBackends returns a boolean if a field has been set.

### GetUseNetworkPolicies

`func (o *CreateSnowflakeBackendWithCertRequest) GetUseNetworkPolicies() bool`

GetUseNetworkPolicies returns the UseNetworkPolicies field if non-nil, zero value otherwise.

### GetUseNetworkPoliciesOk

`func (o *CreateSnowflakeBackendWithCertRequest) GetUseNetworkPoliciesOk() (*bool, bool)`

GetUseNetworkPoliciesOk returns a tuple with the UseNetworkPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNetworkPolicies

`func (o *CreateSnowflakeBackendWithCertRequest) SetUseNetworkPolicies(v bool)`

SetUseNetworkPolicies sets UseNetworkPolicies field to given value.

### HasUseNetworkPolicies

`func (o *CreateSnowflakeBackendWithCertRequest) HasUseNetworkPolicies() bool`

HasUseNetworkPolicies returns a boolean if a field has been set.

### GetUseSso

`func (o *CreateSnowflakeBackendWithCertRequest) GetUseSso() bool`

GetUseSso returns the UseSso field if non-nil, zero value otherwise.

### GetUseSsoOk

`func (o *CreateSnowflakeBackendWithCertRequest) GetUseSsoOk() (*bool, bool)`

GetUseSsoOk returns a tuple with the UseSso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSso

`func (o *CreateSnowflakeBackendWithCertRequest) SetUseSso(v bool)`

SetUseSso sets UseSso field to given value.

### HasUseSso

`func (o *CreateSnowflakeBackendWithCertRequest) HasUseSso() bool`

HasUseSso returns a boolean if a field has been set.

### GetEdition

`func (o *CreateSnowflakeBackendWithCertRequest) GetEdition() string`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *CreateSnowflakeBackendWithCertRequest) GetEditionOk() (*string, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *CreateSnowflakeBackendWithCertRequest) SetEdition(v string)`

SetEdition sets Edition field to given value.

### HasEdition

`func (o *CreateSnowflakeBackendWithCertRequest) HasEdition() bool`

HasEdition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


