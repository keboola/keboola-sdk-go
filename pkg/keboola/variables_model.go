package keboola

type VariableHash string

func (v VariableHash) String() string {
	return string(v)
}

type VariableKey struct {
	Hash VariableHash `json:"hash" validate:"required"`
}

type Variable struct {
	Key        string                 `json:"key"`
	Value      string                 `json:"value"`
	Flags      []string               `json:"flags,omitempty"`
	Group      string                 `json:"group,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Hash       VariableHash           `json:"hash"`
}

type VariableCreatePayload struct {
	Key        string                 `json:"key"`
	Value      string                 `json:"value"`
	Flags      []string               `json:"flags,omitempty"`
	Group      string                 `json:"group,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}

type VariableListOptions struct {
	Key        string                 `json:"key,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Offset     int                    `json:"offset,omitempty"`
	Limit      int                    `json:"limit,omitempty"`
}
