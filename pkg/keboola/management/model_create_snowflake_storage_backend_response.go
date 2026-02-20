/*
Manage API

API version: 1.0.0
*/

package management

import (
	"encoding/json"
)

// checks if the CreateSnowflakeStorageBackendResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSnowflakeStorageBackendResponse{}

// CreateSnowflakeStorageBackendResponse struct for CreateSnowflakeStorageBackendResponse
type CreateSnowflakeStorageBackendResponse struct {
	Id *float32 `json:"id,omitempty"`
	Host *string `json:"host,omitempty"`
	Warehouse *string `json:"warehouse,omitempty"`
	Username *string `json:"username,omitempty"`
	Owner *string `json:"owner,omitempty"`
	TechnicalOwner *string `json:"technicalOwner,omitempty"`
	TechnicalOwnerContactEmails []string `json:"technicalOwnerContactEmails,omitempty"`
	Region *string `json:"region,omitempty"`
	Created *string `json:"created,omitempty"`
	CreatorName *string `json:"creatorName,omitempty"`
	CreatorId *float32 `json:"creatorId,omitempty"`
	IsMaintenance *bool `json:"isMaintenance,omitempty"`
	UseDynamicBackends *bool `json:"useDynamicBackends,omitempty"`
	UseNetworkPolicies *bool `json:"useNetworkPolicies,omitempty"`
	UseSso *bool `json:"useSso,omitempty"`
	IsSsoEnabled *bool `json:"isSsoEnabled,omitempty"`
	IsSsoConfigured *bool `json:"isSsoConfigured,omitempty"`
	Edition *string `json:"edition,omitempty"`
	SqlTemplate *string `json:"sqlTemplate,omitempty"`
	IsEnabled *bool `json:"isEnabled,omitempty"`
	UserPublicKey *string `json:"userPublicKey,omitempty"`
	SecurityIntegrationKey *string `json:"securityIntegrationKey,omitempty"`
}

// NewCreateSnowflakeStorageBackendResponse instantiates a new CreateSnowflakeStorageBackendResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSnowflakeStorageBackendResponse() *CreateSnowflakeStorageBackendResponse {
	this := CreateSnowflakeStorageBackendResponse{}
	return &this
}

// NewCreateSnowflakeStorageBackendResponseWithDefaults instantiates a new CreateSnowflakeStorageBackendResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSnowflakeStorageBackendResponseWithDefaults() *CreateSnowflakeStorageBackendResponse {
	this := CreateSnowflakeStorageBackendResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateSnowflakeStorageBackendResponse) GetId() float32 {
	if o == nil || IsNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnowflakeStorageBackendResponse) GetIdOk() (*float32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateSnowflakeStorageBackendResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *CreateSnowflakeStorageBackendResponse) SetId(v float32) {
	o.Id = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *CreateSnowflakeStorageBackendResponse) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnowflakeStorageBackendResponse) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *CreateSnowflakeStorageBackendResponse) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *CreateSnowflakeStorageBackendResponse) SetHost(v string) {
	o.Host = &v
}

// GetWarehouse returns the Warehouse field value if set, zero value otherwise.
func (o *CreateSnowflakeStorageBackendResponse) GetWarehouse() string {
	if o == nil || IsNil(o.Warehouse) {
		var ret string
		return ret
	}
	return *o.Warehouse
}

// GetWarehouseOk returns a tuple with the Warehouse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnowflakeStorageBackendResponse) GetWarehouseOk() (*string, bool) {
	if o == nil || IsNil(o.Warehouse) {
		return nil, false
	}
	return o.Warehouse, true
}

// HasWarehouse returns a boolean if a field has been set.
func (o *CreateSnowflakeStorageBackendResponse) HasWarehouse() bool {
	if o != nil && !IsNil(o.Warehouse) {
		return true
	}

	return false
}

// SetWarehouse gets a reference to the given string and assigns it to the Warehouse field.
func (o *CreateSnowflakeStorageBackendResponse) SetWarehouse(v string) {
	o.Warehouse = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *CreateSnowflakeStorageBackendResponse) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnowflakeStorageBackendResponse) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *CreateSnowflakeStorageBackendResponse) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *CreateSnowflakeStorageBackendResponse) SetUsername(v string) {
	o.Username = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *CreateSnowflakeStorageBackendResponse) GetOwner() string {
	if o == nil || IsNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnowflakeStorageBackendResponse) GetOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *CreateSnowflakeStorageBackendResponse) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *CreateSnowflakeStorageBackendResponse) SetOwner(v string) {
	o.Owner = &v
}

// GetTechnicalOwner returns the TechnicalOwner field value if set, zero value otherwise.
func (o *CreateSnowflakeStorageBackendResponse) GetTechnicalOwner() string {
	if o == nil || IsNil(o.TechnicalOwner) {
		var ret string
		return ret
	}
	return *o.TechnicalOwner
}

// GetTechnicalOwnerOk returns a tuple with the TechnicalOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnowflakeStorageBackendResponse) GetTechnicalOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.TechnicalOwner) {
		return nil, false
	}
	return o.TechnicalOwner, true
}

// HasTechnicalOwner returns a boolean if a field has been set.
func (o *CreateSnowflakeStorageBackendResponse) HasTechnicalOwner() bool {
	if o != nil && !IsNil(o.TechnicalOwner) {
		return true
	}

	return false
}

// SetTechnicalOwner gets a reference to the given string and assigns it to the TechnicalOwner field.
func (o *CreateSnowflakeStorageBackendResponse) SetTechnicalOwner(v string) {
	o.TechnicalOwner = &v
}

// GetTechnicalOwnerContactEmails returns the TechnicalOwnerContactEmails field value if set, zero value otherwise.
func (o *CreateSnowflakeStorageBackendResponse) GetTechnicalOwnerContactEmails() []string {
	if o == nil || IsNil(o.TechnicalOwnerContactEmails) {
		var ret []string
		return ret
	}
	return o.TechnicalOwnerContactEmails
}

// GetTechnicalOwnerContactEmailsOk returns a tuple with the TechnicalOwnerContactEmails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnowflakeStorageBackendResponse) GetTechnicalOwnerContactEmailsOk() ([]string, bool) {
	if o == nil || IsNil(o.TechnicalOwnerContactEmails) {
		return nil, false
	}
	return o.TechnicalOwnerContactEmails, true
}

// HasTechnicalOwnerContactEmails returns a boolean if a field has been set.
func (o *CreateSnowflakeStorageBackendResponse) HasTechnicalOwnerContactEmails() bool {
	if o != nil && !IsNil(o.TechnicalOwnerContactEmails) {
		return true
	}

	return false
}

// SetTechnicalOwnerContactEmails gets a reference to the given []string and assigns it to the TechnicalOwnerContactEmails field.
func (o *CreateSnowflakeStorageBackendResponse) SetTechnicalOwnerContactEmails(v []string) {
	o.TechnicalOwnerContactEmails = v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *CreateSnowflakeStorageBackendResponse) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnowflakeStorageBackendResponse) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *CreateSnowflakeStorageBackendResponse) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *CreateSnowflakeStorageBackendResponse) SetRegion(v string) {
	o.Region = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *CreateSnowflakeStorageBackendResponse) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnowflakeStorageBackendResponse) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *CreateSnowflakeStorageBackendResponse) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *CreateSnowflakeStorageBackendResponse) SetCreated(v string) {
	o.Created = &v
}

// GetCreatorName returns the CreatorName field value if set, zero value otherwise.
func (o *CreateSnowflakeStorageBackendResponse) GetCreatorName() string {
	if o == nil || IsNil(o.CreatorName) {
		var ret string
		return ret
	}
	return *o.CreatorName
}

// GetCreatorNameOk returns a tuple with the CreatorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnowflakeStorageBackendResponse) GetCreatorNameOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorName) {
		return nil, false
	}
	return o.CreatorName, true
}

// HasCreatorName returns a boolean if a field has been set.
func (o *CreateSnowflakeStorageBackendResponse) HasCreatorName() bool {
	if o != nil && !IsNil(o.CreatorName) {
		return true
	}

	return false
}

// SetCreatorName gets a reference to the given string and assigns it to the CreatorName field.
func (o *CreateSnowflakeStorageBackendResponse) SetCreatorName(v string) {
	o.CreatorName = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *CreateSnowflakeStorageBackendResponse) GetCreatorId() float32 {
	if o == nil || IsNil(o.CreatorId) {
		var ret float32
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnowflakeStorageBackendResponse) GetCreatorIdOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatorId) {
		return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *CreateSnowflakeStorageBackendResponse) HasCreatorId() bool {
	if o != nil && !IsNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given float32 and assigns it to the CreatorId field.
func (o *CreateSnowflakeStorageBackendResponse) SetCreatorId(v float32) {
	o.CreatorId = &v
}

// GetIsMaintenance returns the IsMaintenance field value if set, zero value otherwise.
func (o *CreateSnowflakeStorageBackendResponse) GetIsMaintenance() bool {
	if o == nil || IsNil(o.IsMaintenance) {
		var ret bool
		return ret
	}
	return *o.IsMaintenance
}

// GetIsMaintenanceOk returns a tuple with the IsMaintenance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnowflakeStorageBackendResponse) GetIsMaintenanceOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMaintenance) {
		return nil, false
	}
	return o.IsMaintenance, true
}

// HasIsMaintenance returns a boolean if a field has been set.
func (o *CreateSnowflakeStorageBackendResponse) HasIsMaintenance() bool {
	if o != nil && !IsNil(o.IsMaintenance) {
		return true
	}

	return false
}

// SetIsMaintenance gets a reference to the given bool and assigns it to the IsMaintenance field.
func (o *CreateSnowflakeStorageBackendResponse) SetIsMaintenance(v bool) {
	o.IsMaintenance = &v
}

// GetUseDynamicBackends returns the UseDynamicBackends field value if set, zero value otherwise.
func (o *CreateSnowflakeStorageBackendResponse) GetUseDynamicBackends() bool {
	if o == nil || IsNil(o.UseDynamicBackends) {
		var ret bool
		return ret
	}
	return *o.UseDynamicBackends
}

// GetUseDynamicBackendsOk returns a tuple with the UseDynamicBackends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnowflakeStorageBackendResponse) GetUseDynamicBackendsOk() (*bool, bool) {
	if o == nil || IsNil(o.UseDynamicBackends) {
		return nil, false
	}
	return o.UseDynamicBackends, true
}

// HasUseDynamicBackends returns a boolean if a field has been set.
func (o *CreateSnowflakeStorageBackendResponse) HasUseDynamicBackends() bool {
	if o != nil && !IsNil(o.UseDynamicBackends) {
		return true
	}

	return false
}

// SetUseDynamicBackends gets a reference to the given bool and assigns it to the UseDynamicBackends field.
func (o *CreateSnowflakeStorageBackendResponse) SetUseDynamicBackends(v bool) {
	o.UseDynamicBackends = &v
}

// GetUseNetworkPolicies returns the UseNetworkPolicies field value if set, zero value otherwise.
func (o *CreateSnowflakeStorageBackendResponse) GetUseNetworkPolicies() bool {
	if o == nil || IsNil(o.UseNetworkPolicies) {
		var ret bool
		return ret
	}
	return *o.UseNetworkPolicies
}

// GetUseNetworkPoliciesOk returns a tuple with the UseNetworkPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnowflakeStorageBackendResponse) GetUseNetworkPoliciesOk() (*bool, bool) {
	if o == nil || IsNil(o.UseNetworkPolicies) {
		return nil, false
	}
	return o.UseNetworkPolicies, true
}

// HasUseNetworkPolicies returns a boolean if a field has been set.
func (o *CreateSnowflakeStorageBackendResponse) HasUseNetworkPolicies() bool {
	if o != nil && !IsNil(o.UseNetworkPolicies) {
		return true
	}

	return false
}

// SetUseNetworkPolicies gets a reference to the given bool and assigns it to the UseNetworkPolicies field.
func (o *CreateSnowflakeStorageBackendResponse) SetUseNetworkPolicies(v bool) {
	o.UseNetworkPolicies = &v
}

// GetUseSso returns the UseSso field value if set, zero value otherwise.
func (o *CreateSnowflakeStorageBackendResponse) GetUseSso() bool {
	if o == nil || IsNil(o.UseSso) {
		var ret bool
		return ret
	}
	return *o.UseSso
}

// GetUseSsoOk returns a tuple with the UseSso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnowflakeStorageBackendResponse) GetUseSsoOk() (*bool, bool) {
	if o == nil || IsNil(o.UseSso) {
		return nil, false
	}
	return o.UseSso, true
}

// HasUseSso returns a boolean if a field has been set.
func (o *CreateSnowflakeStorageBackendResponse) HasUseSso() bool {
	if o != nil && !IsNil(o.UseSso) {
		return true
	}

	return false
}

// SetUseSso gets a reference to the given bool and assigns it to the UseSso field.
func (o *CreateSnowflakeStorageBackendResponse) SetUseSso(v bool) {
	o.UseSso = &v
}

// GetIsSsoEnabled returns the IsSsoEnabled field value if set, zero value otherwise.
func (o *CreateSnowflakeStorageBackendResponse) GetIsSsoEnabled() bool {
	if o == nil || IsNil(o.IsSsoEnabled) {
		var ret bool
		return ret
	}
	return *o.IsSsoEnabled
}

// GetIsSsoEnabledOk returns a tuple with the IsSsoEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnowflakeStorageBackendResponse) GetIsSsoEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSsoEnabled) {
		return nil, false
	}
	return o.IsSsoEnabled, true
}

// HasIsSsoEnabled returns a boolean if a field has been set.
func (o *CreateSnowflakeStorageBackendResponse) HasIsSsoEnabled() bool {
	if o != nil && !IsNil(o.IsSsoEnabled) {
		return true
	}

	return false
}

// SetIsSsoEnabled gets a reference to the given bool and assigns it to the IsSsoEnabled field.
func (o *CreateSnowflakeStorageBackendResponse) SetIsSsoEnabled(v bool) {
	o.IsSsoEnabled = &v
}

// GetIsSsoConfigured returns the IsSsoConfigured field value if set, zero value otherwise.
func (o *CreateSnowflakeStorageBackendResponse) GetIsSsoConfigured() bool {
	if o == nil || IsNil(o.IsSsoConfigured) {
		var ret bool
		return ret
	}
	return *o.IsSsoConfigured
}

// GetIsSsoConfiguredOk returns a tuple with the IsSsoConfigured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnowflakeStorageBackendResponse) GetIsSsoConfiguredOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSsoConfigured) {
		return nil, false
	}
	return o.IsSsoConfigured, true
}

// HasIsSsoConfigured returns a boolean if a field has been set.
func (o *CreateSnowflakeStorageBackendResponse) HasIsSsoConfigured() bool {
	if o != nil && !IsNil(o.IsSsoConfigured) {
		return true
	}

	return false
}

// SetIsSsoConfigured gets a reference to the given bool and assigns it to the IsSsoConfigured field.
func (o *CreateSnowflakeStorageBackendResponse) SetIsSsoConfigured(v bool) {
	o.IsSsoConfigured = &v
}

// GetEdition returns the Edition field value if set, zero value otherwise.
func (o *CreateSnowflakeStorageBackendResponse) GetEdition() string {
	if o == nil || IsNil(o.Edition) {
		var ret string
		return ret
	}
	return *o.Edition
}

// GetEditionOk returns a tuple with the Edition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnowflakeStorageBackendResponse) GetEditionOk() (*string, bool) {
	if o == nil || IsNil(o.Edition) {
		return nil, false
	}
	return o.Edition, true
}

// HasEdition returns a boolean if a field has been set.
func (o *CreateSnowflakeStorageBackendResponse) HasEdition() bool {
	if o != nil && !IsNil(o.Edition) {
		return true
	}

	return false
}

// SetEdition gets a reference to the given string and assigns it to the Edition field.
func (o *CreateSnowflakeStorageBackendResponse) SetEdition(v string) {
	o.Edition = &v
}

// GetSqlTemplate returns the SqlTemplate field value if set, zero value otherwise.
func (o *CreateSnowflakeStorageBackendResponse) GetSqlTemplate() string {
	if o == nil || IsNil(o.SqlTemplate) {
		var ret string
		return ret
	}
	return *o.SqlTemplate
}

// GetSqlTemplateOk returns a tuple with the SqlTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnowflakeStorageBackendResponse) GetSqlTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.SqlTemplate) {
		return nil, false
	}
	return o.SqlTemplate, true
}

// HasSqlTemplate returns a boolean if a field has been set.
func (o *CreateSnowflakeStorageBackendResponse) HasSqlTemplate() bool {
	if o != nil && !IsNil(o.SqlTemplate) {
		return true
	}

	return false
}

// SetSqlTemplate gets a reference to the given string and assigns it to the SqlTemplate field.
func (o *CreateSnowflakeStorageBackendResponse) SetSqlTemplate(v string) {
	o.SqlTemplate = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *CreateSnowflakeStorageBackendResponse) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnowflakeStorageBackendResponse) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *CreateSnowflakeStorageBackendResponse) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *CreateSnowflakeStorageBackendResponse) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetUserPublicKey returns the UserPublicKey field value if set, zero value otherwise.
func (o *CreateSnowflakeStorageBackendResponse) GetUserPublicKey() string {
	if o == nil || IsNil(o.UserPublicKey) {
		var ret string
		return ret
	}
	return *o.UserPublicKey
}

// GetUserPublicKeyOk returns a tuple with the UserPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnowflakeStorageBackendResponse) GetUserPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.UserPublicKey) {
		return nil, false
	}
	return o.UserPublicKey, true
}

// HasUserPublicKey returns a boolean if a field has been set.
func (o *CreateSnowflakeStorageBackendResponse) HasUserPublicKey() bool {
	if o != nil && !IsNil(o.UserPublicKey) {
		return true
	}

	return false
}

// SetUserPublicKey gets a reference to the given string and assigns it to the UserPublicKey field.
func (o *CreateSnowflakeStorageBackendResponse) SetUserPublicKey(v string) {
	o.UserPublicKey = &v
}

// GetSecurityIntegrationKey returns the SecurityIntegrationKey field value if set, zero value otherwise.
func (o *CreateSnowflakeStorageBackendResponse) GetSecurityIntegrationKey() string {
	if o == nil || IsNil(o.SecurityIntegrationKey) {
		var ret string
		return ret
	}
	return *o.SecurityIntegrationKey
}

// GetSecurityIntegrationKeyOk returns a tuple with the SecurityIntegrationKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnowflakeStorageBackendResponse) GetSecurityIntegrationKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SecurityIntegrationKey) {
		return nil, false
	}
	return o.SecurityIntegrationKey, true
}

// HasSecurityIntegrationKey returns a boolean if a field has been set.
func (o *CreateSnowflakeStorageBackendResponse) HasSecurityIntegrationKey() bool {
	if o != nil && !IsNil(o.SecurityIntegrationKey) {
		return true
	}

	return false
}

// SetSecurityIntegrationKey gets a reference to the given string and assigns it to the SecurityIntegrationKey field.
func (o *CreateSnowflakeStorageBackendResponse) SetSecurityIntegrationKey(v string) {
	o.SecurityIntegrationKey = &v
}

func (o CreateSnowflakeStorageBackendResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSnowflakeStorageBackendResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Warehouse) {
		toSerialize["warehouse"] = o.Warehouse
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.TechnicalOwner) {
		toSerialize["technicalOwner"] = o.TechnicalOwner
	}
	if !IsNil(o.TechnicalOwnerContactEmails) {
		toSerialize["technicalOwnerContactEmails"] = o.TechnicalOwnerContactEmails
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.CreatorName) {
		toSerialize["creatorName"] = o.CreatorName
	}
	if !IsNil(o.CreatorId) {
		toSerialize["creatorId"] = o.CreatorId
	}
	if !IsNil(o.IsMaintenance) {
		toSerialize["isMaintenance"] = o.IsMaintenance
	}
	if !IsNil(o.UseDynamicBackends) {
		toSerialize["useDynamicBackends"] = o.UseDynamicBackends
	}
	if !IsNil(o.UseNetworkPolicies) {
		toSerialize["useNetworkPolicies"] = o.UseNetworkPolicies
	}
	if !IsNil(o.UseSso) {
		toSerialize["useSso"] = o.UseSso
	}
	if !IsNil(o.IsSsoEnabled) {
		toSerialize["isSsoEnabled"] = o.IsSsoEnabled
	}
	if !IsNil(o.IsSsoConfigured) {
		toSerialize["isSsoConfigured"] = o.IsSsoConfigured
	}
	if !IsNil(o.Edition) {
		toSerialize["edition"] = o.Edition
	}
	if !IsNil(o.SqlTemplate) {
		toSerialize["sqlTemplate"] = o.SqlTemplate
	}
	if !IsNil(o.IsEnabled) {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if !IsNil(o.UserPublicKey) {
		toSerialize["userPublicKey"] = o.UserPublicKey
	}
	if !IsNil(o.SecurityIntegrationKey) {
		toSerialize["securityIntegrationKey"] = o.SecurityIntegrationKey
	}
	return toSerialize, nil
}

type NullableCreateSnowflakeStorageBackendResponse struct {
	value *CreateSnowflakeStorageBackendResponse
	isSet bool
}

func (v NullableCreateSnowflakeStorageBackendResponse) Get() *CreateSnowflakeStorageBackendResponse {
	return v.value
}

func (v *NullableCreateSnowflakeStorageBackendResponse) Set(val *CreateSnowflakeStorageBackendResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSnowflakeStorageBackendResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSnowflakeStorageBackendResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSnowflakeStorageBackendResponse(val *CreateSnowflakeStorageBackendResponse) *NullableCreateSnowflakeStorageBackendResponse {
	return &NullableCreateSnowflakeStorageBackendResponse{value: val, isSet: true}
}

func (v NullableCreateSnowflakeStorageBackendResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSnowflakeStorageBackendResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
