/*
Manage API

API version: 1.0.0
*/

package management

import (
	"encoding/json"
)

// checks if the ActivateSnowflakeStorageBackendResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivateSnowflakeStorageBackendResponse{}

// ActivateSnowflakeStorageBackendResponse struct for ActivateSnowflakeStorageBackendResponse
type ActivateSnowflakeStorageBackendResponse struct {
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
	IsEnabled *bool `json:"isEnabled,omitempty"`
}

// NewActivateSnowflakeStorageBackendResponse instantiates a new ActivateSnowflakeStorageBackendResponse object
func NewActivateSnowflakeStorageBackendResponse() *ActivateSnowflakeStorageBackendResponse {
	this := ActivateSnowflakeStorageBackendResponse{}
	return &this
}

// NewActivateSnowflakeStorageBackendResponseWithDefaults instantiates a new ActivateSnowflakeStorageBackendResponse object
func NewActivateSnowflakeStorageBackendResponseWithDefaults() *ActivateSnowflakeStorageBackendResponse {
	this := ActivateSnowflakeStorageBackendResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ActivateSnowflakeStorageBackendResponse) GetId() float32 {
	if o == nil || IsNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivateSnowflakeStorageBackendResponse) GetIdOk() (*float32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ActivateSnowflakeStorageBackendResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *ActivateSnowflakeStorageBackendResponse) SetId(v float32) {
	o.Id = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *ActivateSnowflakeStorageBackendResponse) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivateSnowflakeStorageBackendResponse) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *ActivateSnowflakeStorageBackendResponse) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *ActivateSnowflakeStorageBackendResponse) SetHost(v string) {
	o.Host = &v
}

// GetWarehouse returns the Warehouse field value if set, zero value otherwise.
func (o *ActivateSnowflakeStorageBackendResponse) GetWarehouse() string {
	if o == nil || IsNil(o.Warehouse) {
		var ret string
		return ret
	}
	return *o.Warehouse
}

// GetWarehouseOk returns a tuple with the Warehouse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivateSnowflakeStorageBackendResponse) GetWarehouseOk() (*string, bool) {
	if o == nil || IsNil(o.Warehouse) {
		return nil, false
	}
	return o.Warehouse, true
}

// HasWarehouse returns a boolean if a field has been set.
func (o *ActivateSnowflakeStorageBackendResponse) HasWarehouse() bool {
	if o != nil && !IsNil(o.Warehouse) {
		return true
	}

	return false
}

// SetWarehouse gets a reference to the given string and assigns it to the Warehouse field.
func (o *ActivateSnowflakeStorageBackendResponse) SetWarehouse(v string) {
	o.Warehouse = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ActivateSnowflakeStorageBackendResponse) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivateSnowflakeStorageBackendResponse) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ActivateSnowflakeStorageBackendResponse) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ActivateSnowflakeStorageBackendResponse) SetUsername(v string) {
	o.Username = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ActivateSnowflakeStorageBackendResponse) GetOwner() string {
	if o == nil || IsNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivateSnowflakeStorageBackendResponse) GetOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ActivateSnowflakeStorageBackendResponse) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *ActivateSnowflakeStorageBackendResponse) SetOwner(v string) {
	o.Owner = &v
}

// GetTechnicalOwner returns the TechnicalOwner field value if set, zero value otherwise.
func (o *ActivateSnowflakeStorageBackendResponse) GetTechnicalOwner() string {
	if o == nil || IsNil(o.TechnicalOwner) {
		var ret string
		return ret
	}
	return *o.TechnicalOwner
}

// GetTechnicalOwnerOk returns a tuple with the TechnicalOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivateSnowflakeStorageBackendResponse) GetTechnicalOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.TechnicalOwner) {
		return nil, false
	}
	return o.TechnicalOwner, true
}

// HasTechnicalOwner returns a boolean if a field has been set.
func (o *ActivateSnowflakeStorageBackendResponse) HasTechnicalOwner() bool {
	if o != nil && !IsNil(o.TechnicalOwner) {
		return true
	}

	return false
}

// SetTechnicalOwner gets a reference to the given string and assigns it to the TechnicalOwner field.
func (o *ActivateSnowflakeStorageBackendResponse) SetTechnicalOwner(v string) {
	o.TechnicalOwner = &v
}

// GetTechnicalOwnerContactEmails returns the TechnicalOwnerContactEmails field value if set, zero value otherwise.
func (o *ActivateSnowflakeStorageBackendResponse) GetTechnicalOwnerContactEmails() []string {
	if o == nil || IsNil(o.TechnicalOwnerContactEmails) {
		var ret []string
		return ret
	}
	return o.TechnicalOwnerContactEmails
}

// GetTechnicalOwnerContactEmailsOk returns a tuple with the TechnicalOwnerContactEmails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivateSnowflakeStorageBackendResponse) GetTechnicalOwnerContactEmailsOk() ([]string, bool) {
	if o == nil || IsNil(o.TechnicalOwnerContactEmails) {
		return nil, false
	}
	return o.TechnicalOwnerContactEmails, true
}

// HasTechnicalOwnerContactEmails returns a boolean if a field has been set.
func (o *ActivateSnowflakeStorageBackendResponse) HasTechnicalOwnerContactEmails() bool {
	if o != nil && !IsNil(o.TechnicalOwnerContactEmails) {
		return true
	}

	return false
}

// SetTechnicalOwnerContactEmails gets a reference to the given []string and assigns it to the TechnicalOwnerContactEmails field.
func (o *ActivateSnowflakeStorageBackendResponse) SetTechnicalOwnerContactEmails(v []string) {
	o.TechnicalOwnerContactEmails = v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *ActivateSnowflakeStorageBackendResponse) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivateSnowflakeStorageBackendResponse) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *ActivateSnowflakeStorageBackendResponse) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *ActivateSnowflakeStorageBackendResponse) SetRegion(v string) {
	o.Region = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ActivateSnowflakeStorageBackendResponse) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivateSnowflakeStorageBackendResponse) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ActivateSnowflakeStorageBackendResponse) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *ActivateSnowflakeStorageBackendResponse) SetCreated(v string) {
	o.Created = &v
}

// GetCreatorName returns the CreatorName field value if set, zero value otherwise.
func (o *ActivateSnowflakeStorageBackendResponse) GetCreatorName() string {
	if o == nil || IsNil(o.CreatorName) {
		var ret string
		return ret
	}
	return *o.CreatorName
}

// GetCreatorNameOk returns a tuple with the CreatorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivateSnowflakeStorageBackendResponse) GetCreatorNameOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorName) {
		return nil, false
	}
	return o.CreatorName, true
}

// HasCreatorName returns a boolean if a field has been set.
func (o *ActivateSnowflakeStorageBackendResponse) HasCreatorName() bool {
	if o != nil && !IsNil(o.CreatorName) {
		return true
	}

	return false
}

// SetCreatorName gets a reference to the given string and assigns it to the CreatorName field.
func (o *ActivateSnowflakeStorageBackendResponse) SetCreatorName(v string) {
	o.CreatorName = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *ActivateSnowflakeStorageBackendResponse) GetCreatorId() float32 {
	if o == nil || IsNil(o.CreatorId) {
		var ret float32
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivateSnowflakeStorageBackendResponse) GetCreatorIdOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatorId) {
		return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *ActivateSnowflakeStorageBackendResponse) HasCreatorId() bool {
	if o != nil && !IsNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given float32 and assigns it to the CreatorId field.
func (o *ActivateSnowflakeStorageBackendResponse) SetCreatorId(v float32) {
	o.CreatorId = &v
}

// GetIsMaintenance returns the IsMaintenance field value if set, zero value otherwise.
func (o *ActivateSnowflakeStorageBackendResponse) GetIsMaintenance() bool {
	if o == nil || IsNil(o.IsMaintenance) {
		var ret bool
		return ret
	}
	return *o.IsMaintenance
}

// GetIsMaintenanceOk returns a tuple with the IsMaintenance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivateSnowflakeStorageBackendResponse) GetIsMaintenanceOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMaintenance) {
		return nil, false
	}
	return o.IsMaintenance, true
}

// HasIsMaintenance returns a boolean if a field has been set.
func (o *ActivateSnowflakeStorageBackendResponse) HasIsMaintenance() bool {
	if o != nil && !IsNil(o.IsMaintenance) {
		return true
	}

	return false
}

// SetIsMaintenance gets a reference to the given bool and assigns it to the IsMaintenance field.
func (o *ActivateSnowflakeStorageBackendResponse) SetIsMaintenance(v bool) {
	o.IsMaintenance = &v
}

// GetUseDynamicBackends returns the UseDynamicBackends field value if set, zero value otherwise.
func (o *ActivateSnowflakeStorageBackendResponse) GetUseDynamicBackends() bool {
	if o == nil || IsNil(o.UseDynamicBackends) {
		var ret bool
		return ret
	}
	return *o.UseDynamicBackends
}

// GetUseDynamicBackendsOk returns a tuple with the UseDynamicBackends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivateSnowflakeStorageBackendResponse) GetUseDynamicBackendsOk() (*bool, bool) {
	if o == nil || IsNil(o.UseDynamicBackends) {
		return nil, false
	}
	return o.UseDynamicBackends, true
}

// HasUseDynamicBackends returns a boolean if a field has been set.
func (o *ActivateSnowflakeStorageBackendResponse) HasUseDynamicBackends() bool {
	if o != nil && !IsNil(o.UseDynamicBackends) {
		return true
	}

	return false
}

// SetUseDynamicBackends gets a reference to the given bool and assigns it to the UseDynamicBackends field.
func (o *ActivateSnowflakeStorageBackendResponse) SetUseDynamicBackends(v bool) {
	o.UseDynamicBackends = &v
}

// GetUseNetworkPolicies returns the UseNetworkPolicies field value if set, zero value otherwise.
func (o *ActivateSnowflakeStorageBackendResponse) GetUseNetworkPolicies() bool {
	if o == nil || IsNil(o.UseNetworkPolicies) {
		var ret bool
		return ret
	}
	return *o.UseNetworkPolicies
}

// GetUseNetworkPoliciesOk returns a tuple with the UseNetworkPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivateSnowflakeStorageBackendResponse) GetUseNetworkPoliciesOk() (*bool, bool) {
	if o == nil || IsNil(o.UseNetworkPolicies) {
		return nil, false
	}
	return o.UseNetworkPolicies, true
}

// HasUseNetworkPolicies returns a boolean if a field has been set.
func (o *ActivateSnowflakeStorageBackendResponse) HasUseNetworkPolicies() bool {
	if o != nil && !IsNil(o.UseNetworkPolicies) {
		return true
	}

	return false
}

// SetUseNetworkPolicies gets a reference to the given bool and assigns it to the UseNetworkPolicies field.
func (o *ActivateSnowflakeStorageBackendResponse) SetUseNetworkPolicies(v bool) {
	o.UseNetworkPolicies = &v
}

// GetUseSso returns the UseSso field value if set, zero value otherwise.
func (o *ActivateSnowflakeStorageBackendResponse) GetUseSso() bool {
	if o == nil || IsNil(o.UseSso) {
		var ret bool
		return ret
	}
	return *o.UseSso
}

// GetUseSsoOk returns a tuple with the UseSso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivateSnowflakeStorageBackendResponse) GetUseSsoOk() (*bool, bool) {
	if o == nil || IsNil(o.UseSso) {
		return nil, false
	}
	return o.UseSso, true
}

// HasUseSso returns a boolean if a field has been set.
func (o *ActivateSnowflakeStorageBackendResponse) HasUseSso() bool {
	if o != nil && !IsNil(o.UseSso) {
		return true
	}

	return false
}

// SetUseSso gets a reference to the given bool and assigns it to the UseSso field.
func (o *ActivateSnowflakeStorageBackendResponse) SetUseSso(v bool) {
	o.UseSso = &v
}

// GetIsSsoEnabled returns the IsSsoEnabled field value if set, zero value otherwise.
func (o *ActivateSnowflakeStorageBackendResponse) GetIsSsoEnabled() bool {
	if o == nil || IsNil(o.IsSsoEnabled) {
		var ret bool
		return ret
	}
	return *o.IsSsoEnabled
}

// GetIsSsoEnabledOk returns a tuple with the IsSsoEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivateSnowflakeStorageBackendResponse) GetIsSsoEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSsoEnabled) {
		return nil, false
	}
	return o.IsSsoEnabled, true
}

// HasIsSsoEnabled returns a boolean if a field has been set.
func (o *ActivateSnowflakeStorageBackendResponse) HasIsSsoEnabled() bool {
	if o != nil && !IsNil(o.IsSsoEnabled) {
		return true
	}

	return false
}

// SetIsSsoEnabled gets a reference to the given bool and assigns it to the IsSsoEnabled field.
func (o *ActivateSnowflakeStorageBackendResponse) SetIsSsoEnabled(v bool) {
	o.IsSsoEnabled = &v
}

// GetIsSsoConfigured returns the IsSsoConfigured field value if set, zero value otherwise.
func (o *ActivateSnowflakeStorageBackendResponse) GetIsSsoConfigured() bool {
	if o == nil || IsNil(o.IsSsoConfigured) {
		var ret bool
		return ret
	}
	return *o.IsSsoConfigured
}

// GetIsSsoConfiguredOk returns a tuple with the IsSsoConfigured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivateSnowflakeStorageBackendResponse) GetIsSsoConfiguredOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSsoConfigured) {
		return nil, false
	}
	return o.IsSsoConfigured, true
}

// HasIsSsoConfigured returns a boolean if a field has been set.
func (o *ActivateSnowflakeStorageBackendResponse) HasIsSsoConfigured() bool {
	if o != nil && !IsNil(o.IsSsoConfigured) {
		return true
	}

	return false
}

// SetIsSsoConfigured gets a reference to the given bool and assigns it to the IsSsoConfigured field.
func (o *ActivateSnowflakeStorageBackendResponse) SetIsSsoConfigured(v bool) {
	o.IsSsoConfigured = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *ActivateSnowflakeStorageBackendResponse) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivateSnowflakeStorageBackendResponse) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *ActivateSnowflakeStorageBackendResponse) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *ActivateSnowflakeStorageBackendResponse) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

func (o ActivateSnowflakeStorageBackendResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivateSnowflakeStorageBackendResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.IsEnabled) {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	return toSerialize, nil
}

type NullableActivateSnowflakeStorageBackendResponse struct {
	value *ActivateSnowflakeStorageBackendResponse
	isSet bool
}

func (v NullableActivateSnowflakeStorageBackendResponse) Get() *ActivateSnowflakeStorageBackendResponse {
	return v.value
}

func (v *NullableActivateSnowflakeStorageBackendResponse) Set(val *ActivateSnowflakeStorageBackendResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableActivateSnowflakeStorageBackendResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableActivateSnowflakeStorageBackendResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivateSnowflakeStorageBackendResponse(val *ActivateSnowflakeStorageBackendResponse) *NullableActivateSnowflakeStorageBackendResponse {
	return &NullableActivateSnowflakeStorageBackendResponse{value: val, isSet: true}
}

func (v NullableActivateSnowflakeStorageBackendResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivateSnowflakeStorageBackendResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
