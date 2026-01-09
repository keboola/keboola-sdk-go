# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build and Test Commands

```bash
# Run linter (go vet + golangci-lint)
task lint

# Auto-fix lint issues
task fix

# Run all tests
task tests

# Run a single test
go test -v -run TestFunctionName ./pkg/keboola/...

# Run tests in a specific file
go test -v ./pkg/keboola/... -run "TestSearchJobs"

# Run tests with race detection (default in task tests)
go test -race -v ./pkg/keboola/...

# Start godoc server
task godoc  # then open http://localhost:6060
```

## Test Project Setup

Integration tests require real Keboola projects. Before running tests:

1. Copy `build/ci/projects.json` to `projects.json` in the root
2. Replace token placeholders with actual Storage API tokens for each project
3. Tests need projects across AWS, Azure, and GCP backends

Use `APIClientForAnEmptyProject(t, ctx)` or `APIClientForRandomProject(t, ctx)` helpers in tests to get pre-configured API clients.

## Architecture

### Package Structure

- **`pkg/request`** - Abstract HTTP request definitions. `HTTPRequest` is immutable and fluent. `APIRequest[R]` wraps one or more HTTP requests with a typed result.

- **`pkg/client`** - HTTP client implementation of `request.Sender` interface. Handles retries, tracing, transport configuration.

- **`pkg/keboola`** - Keboola API implementations covering Storage API, Queue API, Encryption API, Scheduler API, and Workspaces API.

### Key Patterns

**Request Building**: All API requests follow a fluent builder pattern:
```go
api.NewCreateJobRequest("component-id").
    WithConfig(configID).
    WithBranch(branchID).
    Send(ctx)
```

**Functional Options**: Use `WithXxx` functions for optional parameters:
```go
api.SearchJobsRequest(
    WithSearchJobsBranch(branchID),
    WithSearchJobsLimit(10),
)
```

**Path Parameters**: Use `WithGet("path/{param}")` + `AndPathParam("param", value)` instead of `fmt.Sprintf`.

### API Types

- `PublicAPI` - Unauthenticated requests (index, service discovery)
- `AuthorizedAPI` - Authenticated requests requiring Storage API token
- Service types: `StorageAPI`, `QueueAPI`, `EncryptionAPI`, `SchedulerAPI`, `WorkspacesAPI`

### Testing Conventions

- Use `t.Parallel()` at the start of each test
- Use `t.Context()` instead of `context.Background()` in tests
- Use `require.NoError` for error checks where subsequent code depends on the result
- Use `assert` for non-critical assertions
- Time durations: `5*time.Minute` (number before unit)
- Mock HTTP with `client.NewMockedClient()` and `httpmock.RegisterResponder`

## Module Path

```
github.com/keboola/keboola-sdk-go/v2
```

Import packages as:
```go
import "github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
```
