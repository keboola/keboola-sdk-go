package keboola

import (
	"context"

	"github.com/keboola/keboola-sdk-go/v2/pkg/request"
)

type params struct {
	Type             string
	Shared           bool
	ExpireAfterHours uint64
	Size             string
	ImageVersion     string
}

type CreateWorkspaceOption func(p *params)

func WithShared(v bool) CreateWorkspaceOption {
	return func(p *params) { p.Shared = v }
}

func WithExpireAfterHours(v uint64) CreateWorkspaceOption {
	return func(p *params) { p.ExpireAfterHours = v }
}

func WithSize(v string) CreateWorkspaceOption {
	return func(p *params) { p.Size = v }
}

func WithImageVersion(v string) CreateWorkspaceOption {
	return func(p *params) { p.ImageVersion = v }
}

func newParams(type_ string, opts ...CreateWorkspaceOption) params {
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

func (a *AuthorizedAPI) CreateWorkspaceJobRequest(configID ConfigID, workspaceType string, opts ...CreateWorkspaceOption) request.APIRequest[request.NoResult] {
	params := newParams(workspaceType, opts...)
	req := a.CreateQueueJobConfigDataRequest(WorkspacesComponent, configID, map[string]any{"parameters": params.toMap()}).
		WithOnSuccess(func(ctx context.Context, result *QueueJob) error {
			return a.WaitForQueueJob(ctx, result.ID)
		})
	return request.NewAPIRequest(request.NoResult{}, req)
}

func (a *AuthorizedAPI) DeleteWorkspaceJobRequest(workspaceID WorkspaceID) request.APIRequest[request.NoResult] {
	configData := map[string]any{
		"parameters": map[string]any{
			"task": "delete",
			"id":   workspaceID.String(),
		},
	}
	req := a.CreateQueueJobConfigDataRequest(WorkspacesComponent, "", configData).
		WithOnSuccess(func(ctx context.Context, result *QueueJob) error {
			return a.WaitForQueueJob(ctx, result.ID)
		})
	return request.NewAPIRequest(request.NoResult{}, req)
}
