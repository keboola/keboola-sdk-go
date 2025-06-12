# RunCommandRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | **string** | command to execute | 
**Parameters** | Pointer to **[]interface{}** | command parameters | [optional] 

## Methods

### NewRunCommandRequest

`func NewRunCommandRequest(command string, ) *RunCommandRequest`

NewRunCommandRequest instantiates a new RunCommandRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunCommandRequestWithDefaults

`func NewRunCommandRequestWithDefaults() *RunCommandRequest`

NewRunCommandRequestWithDefaults instantiates a new RunCommandRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *RunCommandRequest) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *RunCommandRequest) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *RunCommandRequest) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetParameters

`func (o *RunCommandRequest) GetParameters() []interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *RunCommandRequest) GetParametersOk() (*[]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *RunCommandRequest) SetParameters(v []interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *RunCommandRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


