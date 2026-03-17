# CreateANewBackend201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**Host** | Pointer to **string** | 37-demo.cmizbsfmzc6w.us-east-1.redshift.amazonaws.com (required) | [optional] 
**Backend** | **string** |  | 
**Region** | Pointer to **string** | east-1 | [optional] 

## Methods

### NewCreateANewBackend201Response

`func NewCreateANewBackend201Response(id float32, backend string, ) *CreateANewBackend201Response`

NewCreateANewBackend201Response instantiates a new CreateANewBackend201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateANewBackend201ResponseWithDefaults

`func NewCreateANewBackend201ResponseWithDefaults() *CreateANewBackend201Response`

NewCreateANewBackend201ResponseWithDefaults instantiates a new CreateANewBackend201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateANewBackend201Response) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateANewBackend201Response) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateANewBackend201Response) SetId(v float32)`

SetId sets Id field to given value.


### GetHost

`func (o *CreateANewBackend201Response) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CreateANewBackend201Response) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CreateANewBackend201Response) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *CreateANewBackend201Response) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetBackend

`func (o *CreateANewBackend201Response) GetBackend() string`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *CreateANewBackend201Response) GetBackendOk() (*string, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *CreateANewBackend201Response) SetBackend(v string)`

SetBackend sets Backend field to given value.


### GetRegion

`func (o *CreateANewBackend201Response) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateANewBackend201Response) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateANewBackend201Response) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CreateANewBackend201Response) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


