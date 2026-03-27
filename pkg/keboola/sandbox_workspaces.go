package keboola

const SandboxWorkspacesComponent = "keboola.sandboxes"

const (
	SandboxWorkspaceSizeSmall  = "small"
	SandboxWorkspaceSizeMedium = "medium"
	SandboxWorkspaceSizeLarge  = "large"
)
const (
	SandboxWorkspaceTypeSnowflake = SandboxWorkspaceType("snowflake")
	SandboxWorkspaceTypeBigQuery  = SandboxWorkspaceType("bigquery")
	SandboxWorkspaceTypePython    = SandboxWorkspaceType("python")
	SandboxWorkspaceTypeR         = SandboxWorkspaceType("r")
)

type SandboxWorkspaceType string

func (s SandboxWorkspaceType) String() string {
	return string(s)
}

func SandboxWorkspaceSizesOrdered() []string {
	return []string{
		SandboxWorkspaceSizeSmall,
		SandboxWorkspaceSizeMedium,
		SandboxWorkspaceSizeLarge,
	}
}

func SandboxWorkspaceSizesMap() map[string]bool {
	return map[string]bool{
		SandboxWorkspaceSizeSmall:  true,
		SandboxWorkspaceSizeMedium: true,
		SandboxWorkspaceSizeLarge:  true,
	}
}

func SandboxWorkspaceTypesOrdered() []SandboxWorkspaceType {
	return []SandboxWorkspaceType{
		SandboxWorkspaceTypeSnowflake,
		SandboxWorkspaceTypeBigQuery,
		SandboxWorkspaceTypePython,
		SandboxWorkspaceTypeR,
	}
}

func SandboxWorkspaceTypesMap() map[SandboxWorkspaceType]bool {
	return map[SandboxWorkspaceType]bool{
		SandboxWorkspaceTypeSnowflake: true,
		SandboxWorkspaceTypeBigQuery:  true,
		SandboxWorkspaceTypePython:    true,
		SandboxWorkspaceTypeR:         true,
	}
}

func SandboxWorkspaceSupportsSizes(typ SandboxWorkspaceType) bool {
	switch typ {
	case SandboxWorkspaceTypePython:
		return true
	case SandboxWorkspaceTypeR:
		return true
	default:
		return false
	}
}
