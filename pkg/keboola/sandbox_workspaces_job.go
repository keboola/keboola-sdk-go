package keboola

import (
	"context"

	"github.com/keboola/keboola-sdk-go/v2/pkg/request"
)

type params struct {
	Type             SandboxWorkspaceType
	Shared           bool
	ExpireAfterHours uint64
	Size             string
	ImageVersion     string
	PublicKey        string
	LoginType        string
}

type CreateSandboxWorkspaceOption func(p *params)

func WithShared(v bool) CreateSandboxWorkspaceOption {
	return func(p *params) { p.Shared = v }
}

func WithExpireAfterHours(v uint64) CreateSandboxWorkspaceOption {
	return func(p *params) { p.ExpireAfterHours = v }
}

func WithSize(v string) CreateSandboxWorkspaceOption {
	return func(p *params) { p.Size = v }
}

func WithImageVersion(v string) CreateSandboxWorkspaceOption {
	return func(p *params) { p.ImageVersion = v }
}

func WithPublicKey(v string) CreateSandboxWorkspaceOption {
	return func(p *params) {
		p.PublicKey = v
		p.LoginType = "snowflake-person-keypair"
	}
}

func newParams(type_ SandboxWorkspaceType, opts ...CreateSandboxWorkspaceOption) params {
	p := params{
		Type:             type_,
		Shared:           false,
		ExpireAfterHours: 0,
	}
	for _, opt := range opts {
		opt(&p)
	}
	return p
}

func (p params) toMap() map[string]any {
	m := map[string]any{
		"task":                 "create",
		"type":                 p.Type,
		"shared":               p.Shared,
		"expirationAfterHours": p.ExpireAfterHours,
	}
	if len(p.Size) > 0 {
		m["size"] = p.Size
	}
	if len(p.ImageVersion) > 0 {
		m["imageVersion"] = p.ImageVersion
	}
	return m
}

func (a *AuthorizedAPI) CreateSandboxWorkspaceJobRequest(
	configID ConfigID,
	workspaceType SandboxWorkspaceType,
	opts ...CreateSandboxWorkspaceOption,
) request.APIRequest[request.NoResult] {
	params := newParams(workspaceType, opts...)
	req := a.NewCreateJobRequest(SandboxWorkspacesComponent).
		WithConfig(configID).
		WithConfigData(map[string]any{"parameters": params.toMap()}).
		Build().
		WithOnSuccess(func(ctx context.Context, result *QueueJob) error {
			return a.WaitForQueueJob(ctx, result.ID)
		})
	return request.NewAPIRequest(request.NoResult{}, req)
}

func (a *AuthorizedAPI) DeleteSandboxWorkspaceJobRequest(workspaceID SandboxWorkspaceID) request.APIRequest[request.NoResult] {
	configData := map[string]any{
		"parameters": map[string]any{
			"task": "delete",
			"id":   workspaceID.String(),
		},
	}
	req := a.NewCreateJobRequest(SandboxWorkspacesComponent).
		WithConfigData(configData).
		Build().WithOnSuccess(func(ctx context.Context, result *QueueJob) error {
		return a.WaitForQueueJob(ctx, result.ID)
	})
	return request.NewAPIRequest(request.NoResult{}, req)
}
