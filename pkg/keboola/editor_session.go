package keboola

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v5"
	"github.com/hashicorp/go-multierror"
	"github.com/relvacode/iso8601"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/keboola/keboola-sdk-go/v2/pkg/request"
)

// EditorSessionID is the unique identifier of an editor session.
type EditorSessionID string

func (v EditorSessionID) String() string {
	return string(v)
}

// EditorSessionStatus is the lifecycle status of an editor session.
type EditorSessionStatus string

const (
	EditorSessionStatusInitializing EditorSessionStatus = "initializing"
	EditorSessionStatusReady        EditorSessionStatus = "ready"
)

// EditorSessionBackendType is the backend database type of an editor session.
type EditorSessionBackendType string

// EditorSessionLoginType is the login/authentication type for an editor session.
type EditorSessionLoginType string

// EditorSessionWorkspaceLoadJob is a reference to a workspace load job.
type EditorSessionWorkspaceLoadJob struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// EditorSession represents a SQL editor session returned by the editor service.
// It replaces the deprecated Snowflake and BigQuery sandbox workspaces.
// Use EditorSession.WorkspaceID with StorageWorkspaceLoadDataRequest / StorageWorkspaceUnloadRequest
// to load or unload data from the session workspace.
type EditorSession struct {
	ID                    EditorSessionID                   `json:"id"`
	CreatedAt             iso8601.Time                      `json:"createdAt"`
	UpdatedAt             iso8601.Time                      `json:"updatedAt"`
	Status                EditorSessionStatus               `json:"status"`
	UserID                string                            `json:"userId"`
	BranchID              string                            `json:"branchId"`
	ComponentID           string                            `json:"componentId"`
	ConfigurationID       string                            `json:"configurationId"`
	WorkspaceSchema       string                            `json:"workspaceSchema"`
	WorkspaceDatabase     string                            `json:"workspaceDatabase"`
	WorkspaceID           string                            `json:"workspaceId"`
	WorkspaceCreateJobID  string                            `json:"workspaceCreateJobId"`
	WorkspaceLoadJobs     [][]EditorSessionWorkspaceLoadJob `json:"workspaceLoadJobs"`
	SnowflakePrivateKey   string                            `json:"snowflakePrivateKey"` //nolint: gosec
	ReadOnlyStorageAccess bool                              `json:"readOnlyStorageAccess"`
	BackendType           EditorSessionBackendType          `json:"backendType"`
	BackendSize           string                            `json:"backendSize"`
	LastError             string                            `json:"lastError"`
	Shared                bool                              `json:"shared"`
	WorkspaceLoginType    EditorSessionLoginType            `json:"workspaceLoginType"`
	LastLoadedAt          *iso8601.Time                     `json:"lastLoadedAt"`
}

// EditorSessionWithConfig pairs an EditorSession with its backing keboola.sandboxes config.
// Use EditorSession.WorkspaceID with StorageWorkspaceLoadDataRequest / StorageWorkspaceUnloadRequest
// to load or unload data, and Config.ID when calling DeleteEditorSession.
type EditorSessionWithConfig struct {
	EditorSession *EditorSession
	Config        *Config
}

// CreateEditorSessionPayload is the request body for POST /sql/sessions.
type CreateEditorSessionPayload struct {
	BranchID        string `json:"branchId"`
	ComponentID     string `json:"componentId,omitempty"`
	ConfigurationID string `json:"configurationId,omitempty"`
	LoadMode        string `json:"loadMode,omitempty"`
	LoginType       string `json:"loginType,omitempty"`
}

// WaitForEditorSession polls the editor session until its status is "ready" or the context deadline is exceeded.
func (a *AuthorizedAPI) WaitForEditorSession(ctx context.Context, id EditorSessionID) (err error) {
	_, ok := ctx.Deadline()
	if !ok {
		return fmt.Errorf("timeout for the editor session was not set")
	}

	// Telemetry
	parentSpan := trace.SpanFromContext(ctx)
	var span trace.Span
	ctx, span = parentSpan.TracerProvider().Tracer(appName).Start(ctx, "keboola.go.api.client.waitFor.editorSession")
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	retry := newEditorSessionBackoff()
	for {
		session, err := a.GetEditorSessionRequest(id).Send(ctx)
		if err != nil {
			return err
		}

		if session.Status == EditorSessionStatusReady {
			return nil
		}

		if session.LastError != "" {
			return fmt.Errorf(`editor session "%s" (backend: %s, branch: %s) failed: %s`, id, session.BackendType, session.BranchID, session.LastError)
		}

		select {
		case <-ctx.Done():
			return fmt.Errorf(`error while waiting for the editor session "%s" to become ready: %w`, id, ctx.Err())
		case <-time.After(retry.NextBackOff()):
			// try again
		}
	}
}

// newEditorSessionBackoff creates retry backoff for WaitForEditorSession.
func newEditorSessionBackoff() *backoff.ExponentialBackOff {
	b := backoff.NewExponentialBackOff()
	b.RandomizationFactor = 0
	b.InitialInterval = 3 * time.Second
	b.Multiplier = 2
	b.MaxInterval = 5 * time.Second
	b.Reset()
	return b
}

// waitForEditorSessionDeleted polls GetEditorSessionRequest until the session returns 404.
func (a *AuthorizedAPI) waitForEditorSessionDeleted(ctx context.Context, id EditorSessionID) error {
	retry := newEditorSessionBackoff()
	for {
		_, err := a.GetEditorSessionRequest(id).Send(ctx)
		if err != nil {
			var editorErr *EditorError
			if errors.As(err, &editorErr) && editorErr.StatusCode() == http.StatusNotFound {
				return nil
			}
			return err
		}
		select {
		case <-ctx.Done():
			return fmt.Errorf("timed out waiting for editor session %q to be deleted: %w", id, ctx.Err())
		case <-time.After(retry.NextBackOff()):
		}
	}
}

// waitForStorageWorkspaceDeleted polls StorageWorkspaceDetailRequest until the workspace
// returns 404, confirming the async storage workspaceDrop job triggered by the editor service
// has completed. This prevents CleanStorageWorkspaces from racing and issuing a second drop.
func (a *AuthorizedAPI) waitForStorageWorkspaceDeleted(ctx context.Context, branchID BranchID, workspaceID uint64) error {
	retry := newEditorSessionBackoff()
	for {
		_, err := a.StorageWorkspaceDetailRequest(branchID, workspaceID).Send(ctx)
		if err != nil {
			var storageErr *StorageError
			if errors.As(err, &storageErr) && storageErr.StatusCode() == http.StatusNotFound {
				return nil
			}
			return err
		}
		select {
		case <-ctx.Done():
			return fmt.Errorf("timed out waiting for storage workspace %d to be deleted: %w", workspaceID, ctx.Err())
		case <-time.After(retry.NextBackOff()):
		}
	}
}

// ListEditorSessionsRequest returns a list of all editor sessions.
// https://api.keboola.com/?service=editor#get-/sql/sessions
func (a *AuthorizedAPI) ListEditorSessionsRequest() request.APIRequest[*[]*EditorSession] {
	result := make([]*EditorSession, 0)
	req := a.newRequest(EditorAPI).
		WithResult(&result).
		WithGet(EditorAPISessions)
	return request.NewAPIRequest(&result, req)
}

// GetEditorSessionRequest returns a single editor session by ID.
// https://api.keboola.com/?service=editor#get-/sql/sessions/-sessionId-
func (a *AuthorizedAPI) GetEditorSessionRequest(id EditorSessionID) request.APIRequest[*EditorSession] {
	result := &EditorSession{}
	req := a.newRequest(EditorAPI).
		WithResult(result).
		WithGet(EditorAPISession).
		AndPathParam("sessionId", id.String())
	return request.NewAPIRequest(result, req)
}

// CreateEditorSessionRequest creates a new editor session.
// The session starts with status "initializing" and this request waits until status is "ready".
// https://api.keboola.com/?service=editor#post-/sql/sessions
func (a *AuthorizedAPI) CreateEditorSessionRequest(payload CreateEditorSessionPayload) request.APIRequest[*EditorSession] {
	result := &EditorSession{}
	httpReq := a.newRequest(EditorAPI).
		WithResult(result).
		WithPost(EditorAPISessions).
		WithJSONBody(payload)
	return request.NewAPIRequest(result, httpReq).
		WithOnSuccess(func(ctx context.Context, session *EditorSession) error {
			waitCtx, cancelFn := context.WithTimeout(ctx, a.onSuccessTimeout)
			defer cancelFn()
			if err := a.WaitForEditorSession(waitCtx, session.ID); err != nil {
				return err
			}
			// Refresh result with final session state after it becomes ready.
			// Use waitCtx so the whole OnSuccess path is bounded by onSuccessTimeout.
			finalSession, err := a.GetEditorSessionRequest(session.ID).Send(waitCtx)
			if err != nil {
				return err
			}
			*result = *finalSession
			return nil
		})
}

// EditorSessionCredentials holds the rotated keypair returned by reset-credentials.
type EditorSessionCredentials struct {
	PrivateKey string `json:"privateKey"` //nolint: gosec
	PublicKey  string `json:"publicKey"`
}

// ResetEditorSessionCredentialsRequest rotates the credentials of a running editor session.
// https://api.keboola.com/?service=editor#post-/sql/sessions/-id-/reset-credentials
func (a *AuthorizedAPI) ResetEditorSessionCredentialsRequest(id EditorSessionID) request.APIRequest[*EditorSessionCredentials] {
	result := &EditorSessionCredentials{}
	req := a.newRequest(EditorAPI).
		WithResult(result).
		WithPost(EditorAPISessionResetCredentials).
		AndPathParam("sessionId", id.String())
	return request.NewAPIRequest(result, req)
}

// DeleteEditorSessionRequest deletes an editor session by ID.
// https://api.keboola.com/?service=editor#delete-/sql/sessions/-sessionId-
func (a *AuthorizedAPI) DeleteEditorSessionRequest(id EditorSessionID) request.APIRequest[request.NoResult] {
	req := a.newRequest(EditorAPI).
		WithDelete(EditorAPISession).
		AndPathParam("sessionId", id.String())
	return request.NewAPIRequest(request.NoResult{}, req)
}

// CreateEditorSession creates a keboola.sandboxes config and then starts a new editor session
// linked to that config. It waits until the session is ready and returns both together.
func (a *AuthorizedAPI) CreateEditorSession(ctx context.Context, branchID BranchID, workspaceName string) (*EditorSessionWithConfig, error) {
	// Step 1: Create keboola.sandboxes config.
	emptyConfig, err := a.CreateSandboxWorkspaceConfigRequest(branchID, workspaceName).Send(ctx)
	if err != nil {
		return nil, err
	}

	// Step 2: Start the editor session linked to that config.
	session, err := a.CreateEditorSessionRequest(CreateEditorSessionPayload{
		BranchID:        branchID.String(),
		ComponentID:     SandboxWorkspacesComponent,
		ConfigurationID: emptyConfig.ID.String(),
	}).Send(ctx)
	if err != nil {
		// If the POST succeeded but WaitForEditorSession timed out, session is non-nil.
		// Best-effort: delete the orphaned session to avoid leaving it running.
		if session != nil {
			if delErr := a.DeleteEditorSessionRequest(session.ID).SendOrErr(ctx); delErr != nil {
				err = multierror.Append(err, fmt.Errorf("failed to clean up orphaned session %s: %w", session.ID, delErr))
			}
		}
		// Best-effort: delete the orphaned config so it doesn't litter the UI.
		if _, cleanupErr := a.DeleteSandboxWorkspaceConfigRequest(branchID, emptyConfig.ID).Send(ctx); cleanupErr != nil {
			err = multierror.Append(err, fmt.Errorf("failed to clean up orphaned config %s: %w", emptyConfig.ID, cleanupErr))
		}
		return nil, err
	}

	return &EditorSessionWithConfig{EditorSession: session, Config: emptyConfig.Config}, nil
}

// DeleteEditorSession deletes an editor session and its backing keboola.sandboxes config.
// Both deletions are always attempted; any errors are combined and returned together.
func (a *AuthorizedAPI) DeleteEditorSession(ctx context.Context, branchID BranchID, configID ConfigID, sessionID EditorSessionID) error {
	var err error
	if e := a.DeleteEditorSessionRequest(sessionID).SendOrErr(ctx); e != nil {
		err = multierror.Append(err, e)
	}
	if _, e := a.DeleteSandboxWorkspaceConfigRequest(branchID, configID).Send(ctx); e != nil {
		err = multierror.Append(err, e)
	}
	return err
}

// CleanEditorSessions deletes all editor sessions in the project.
// It is used for test project cleanup alongside CleanSandboxWorkspaceInstances.
func (a *AuthorizedAPI) CleanEditorSessions(ctx context.Context) error {
	sessions, err := a.ListEditorSessionsRequest().Send(ctx)
	if err != nil {
		return err
	}

	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}

	for _, s := range *sessions {
		wg.Go(func() {
			if e := a.DeleteEditorSessionRequest(s.ID).SendOrErr(ctx); e != nil {
				m.Lock()
				err = multierror.Append(err, e)
				m.Unlock()
				// Don't return — still attempt to delete the backing config below.
			} else {
				// Wait for the editor session to be fully removed from the editor API.
				if e := a.waitForEditorSessionDeleted(ctx, s.ID); e != nil {
					m.Lock()
					err = multierror.Append(err, e)
					m.Unlock()
				} else if s.WorkspaceID != "" && s.BranchID != "" {
					// The editor service drops the backing storage workspace asynchronously.
					// Wait for it to disappear from the storage API so CleanStorageWorkspaces
					// won't race and trigger a second workspaceDrop on the same workspace.
					workspaceID, parseErr := strconv.ParseUint(s.WorkspaceID, 10, 64)
					branchIDInt, parseBranchErr := strconv.Atoi(s.BranchID)
					if parseErr == nil && parseBranchErr == nil {
						if e := a.waitForStorageWorkspaceDeleted(ctx, BranchID(branchIDInt), workspaceID); e != nil {
							m.Lock()
							err = multierror.Append(err, e)
							m.Unlock()
						}
					}
				}
			}
			// Also delete the backing keboola.sandboxes config created by CreateEditorSession.
			// Guard on ComponentID so we never touch configs belonging to other components.
			if s.ComponentID == SandboxWorkspacesComponent && s.ConfigurationID != "" && s.BranchID != "" {
				branchIDInt, parseErr := strconv.Atoi(s.BranchID)
				if parseErr == nil {
					if _, e := a.DeleteSandboxWorkspaceConfigRequest(BranchID(branchIDInt), ConfigID(s.ConfigurationID)).Send(ctx); e != nil {
						m.Lock()
						err = multierror.Append(err, e)
						m.Unlock()
					}
				}
			}
		})
	}

	wg.Wait()
	return err
}
