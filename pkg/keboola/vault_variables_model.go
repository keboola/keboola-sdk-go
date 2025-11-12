package keboola

type VaultVariableHash string

func (v VaultVariableHash) String() string {
	return string(v)
}

type VaultVariableKey struct {
	Hash VaultVariableHash `json:"hash" validate:"required"`
}

type VaultVariable struct {
	Key        string                 `json:"key"`
	Value      string                 `json:"value"`
	Flags      []string               `json:"flags,omitempty"`
	Group      string                 `json:"group,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Hash       VaultVariableHash      `json:"hash"`
}

type VaultVariableCreatePayload struct {
	Key        string                 `json:"key"`
	Value      string                 `json:"value"`
	Flags      []string               `json:"flags,omitempty"`
	Group      string                 `json:"group,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}

type VaultVariableListOptions struct {
	Key        string                 `json:"key,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Offset     int                    `json:"offset,omitempty"`
	Limit      int                    `json:"limit,omitempty"`
}
