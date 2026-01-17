# RegisterNewApplicationVersionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManifestUrl** | **string** | URL of manifest describing UI build | 
**Activate** | Pointer to **bool** | Activate version after registration | [optional] 

## Methods

### NewRegisterNewApplicationVersionRequest

`func NewRegisterNewApplicationVersionRequest(manifestUrl string, ) *RegisterNewApplicationVersionRequest`

NewRegisterNewApplicationVersionRequest instantiates a new RegisterNewApplicationVersionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterNewApplicationVersionRequestWithDefaults

`func NewRegisterNewApplicationVersionRequestWithDefaults() *RegisterNewApplicationVersionRequest`

NewRegisterNewApplicationVersionRequestWithDefaults instantiates a new RegisterNewApplicationVersionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManifestUrl

`func (o *RegisterNewApplicationVersionRequest) GetManifestUrl() string`

GetManifestUrl returns the ManifestUrl field if non-nil, zero value otherwise.

### GetManifestUrlOk

`func (o *RegisterNewApplicationVersionRequest) GetManifestUrlOk() (*string, bool)`

GetManifestUrlOk returns a tuple with the ManifestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestUrl

`func (o *RegisterNewApplicationVersionRequest) SetManifestUrl(v string)`

SetManifestUrl sets ManifestUrl field to given value.


### GetActivate

`func (o *RegisterNewApplicationVersionRequest) GetActivate() bool`

GetActivate returns the Activate field if non-nil, zero value otherwise.

### GetActivateOk

`func (o *RegisterNewApplicationVersionRequest) GetActivateOk() (*bool, bool)`

GetActivateOk returns a tuple with the Activate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivate

`func (o *RegisterNewApplicationVersionRequest) SetActivate(v bool)`

SetActivate sets Activate field to given value.

### HasActivate

`func (o *RegisterNewApplicationVersionRequest) HasActivate() bool`

HasActivate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


