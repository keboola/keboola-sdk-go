# CommandExecute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | **string** | command to execute | 
**Parameters** | Pointer to **[]interface{}** | command parameters | [optional] 

## Methods

### NewCommandExecute

`func NewCommandExecute(command string, ) *CommandExecute`

NewCommandExecute instantiates a new CommandExecute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandExecuteWithDefaults

`func NewCommandExecuteWithDefaults() *CommandExecute`

NewCommandExecuteWithDefaults instantiates a new CommandExecute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *CommandExecute) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *CommandExecute) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *CommandExecute) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetParameters

`func (o *CommandExecute) GetParameters() []interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *CommandExecute) GetParametersOk() (*[]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *CommandExecute) SetParameters(v []interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *CommandExecute) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


