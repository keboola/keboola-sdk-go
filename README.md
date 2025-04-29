# Keboola SDK for Go

[![Build and test Go](https://github.com/keboola/keboola-sdk-go/actions/workflows/push.yml/badge.svg)](https://github.com/keboola/keboola-sdk-go/actions/workflows/push.yml)

The Keboola SDK for Go programming language provides developers with easy access to the Keboola platform and related services from
their code.

## Getting started

1. You will need a Keboola account and project - to get one, sign up [here](https://www.keboola.com/). Once you
have a project, get a Storage API token for that project. With this token, you can use all Keboola APIs that require a
Storage API token.
2. Install SDK dependencies:
    ```shell
    go get github.com/keboola/keboola-sdk-go
    ```
3. Write your code

### Create configuration

Configuration is a basic Keboola concept - it's a definition of work to be done by a component. In this example, we will
create a simple Python transformation that will print `Hello World` text.

```go
package your_project

import (
   "context"
   "fmt"
   "log"
   "encoding/json"
   "github.com/keboola/keboola-sdk-go/pkg/keboola"
   "github.com/keboola/go-utils/pkg/orderedmap"
)

func main() {
	// Create context
	ctx := context.Background()

	// Connect to Keboola API with our token
	// This is like logging into the Keboola website
	api, err := keboola.NewAuthorizedAPI(
		ctx,
		"<https://connection...>",   // The Keboola server address
		"<Storage API token>", // Replace with your token
	)
	// Check if we connected successfully
	if err != nil {
		log.Fatalf("Could not connect to Keboola: %v", err)
	}

	// Get the default branch
	branch, err := api.GetDefaultBranchRequest().Send(ctx)
	if err != nil {
		log.Fatalf("Could not get default branch: %v", err)
	}

	// Define our Python transformation configuration
	pythonConfig := `{
		"parameters": {
			"blocks": [
				{
					"name": "Block",
					"codes": [
						{
							"name": "Code",
							"script": ["print(\"Hello World\")"] 
						}
					]
				}
			]
		}
	}`

	// Keboola requires configurations to be encrypted before storing them
	encryptedData, err := api.EncryptRequest(
		<project ID>, // Your project ID
		"keboola.python-transformation-v2",
		map[string]string{"configuration": pythonConfig},
	).Send(ctx)
	if err != nil {
		log.Fatalf("Could not encrypt the configuration: %v", err)
	}

	// Convert the encrypted JSON to an OrderedMap
	// Keboola uses a special map type that keeps keys in order
	// First, get the encrypted configuration string
	encryptedConfigStr := (*encryptedData)["configuration"]

	// Convert the JSON string to a Go map
	var configMap map[string]any
	if err := json.Unmarshal([]byte(encryptedConfigStr), &configMap); err != nil {
		log.Fatalf("Could not parse the JSON: %v", err)
	}

	// Convert the Go map to an OrderedMap
	configContent := orderedmap.New()
	for key, value := range configMap {
		configContent.Set(key, value)
	}

	// Build the configuration object with all required fields
	newConfig := &keboola.ConfigWithRows{
		Config: &keboola.Config{
			ConfigKey: keboola.ConfigKey{
				BranchID:    branch.ID,
				ComponentID: "keboola.python-transformation-v2",
			},
			Name:        "Hello World",
			Content:     configContent,
		},
	}

	// Send the request to create the configuration
	response, err := api.CreateConfigRequest(newConfig, true).Send(ctx)
	if err != nil {
		log.Fatalf("Could not create the configuration: %v", err)
	}

	// Success! Print the ID of our new configuration
	fmt.Printf("Success! Created configuration with ID: %s\n", response.ID)
 
	// Run a job
	job, err := api.NewCreateJobRequest("keboola.python-transformation-v2").
		WithConfig(response.ID).
		Send(ctx)
	if err != nil {
		log.Fatalf("Could not run a job: %v", err)
	}
	fmt.Printf("Success! Created job with ID: %s\n", job.ID)
```

## Packages

### `request` package

The `request` package provides abstract and immutable definition of an HTTP request by the `request.HTTPRequest`.

One or more HTTP requests can be grouped together to the generic type `request.APIRequest[R]`,
where the `R` is a result type to which HTTP requests are mapped.

Request sending is provided by the `Sender` interface, the `client` package provides its default implementation.

### `client` package

The `client` package provides default implementation of the `request.Sender` interface based on the standard `net/http` package.

### `keboola` package

The `keboola` package provides the `keboola.API` implementation, it covers:
  - [Storage API](https://keboola.docs.apiary.io/#)
  - [Encryption API](https://keboolaencryption.docs.apiary.io/#)
  - [Jobs Queue API](https://app.swaggerhub.com/apis-docs/keboola/job-queue-api)
  - [Sandboxes API](https://sandboxes.keboola.com/documentation)
  - [Scheduler API](https://app.swaggerhub.com/apis/odinuv/scheduler)

Not all API requests are covered, API requests are extended as needed.

## Direct HTTP Requests

The `request` package provides a flexible way to make direct HTTP requests using the `NewHTTPRequest` function with any implementation of the `Sender` interface. This approach offers several advantages:

- **Immutability**: Each method call returns a new request instance, allowing for safe request modification.
- **Fluent API**: Chain method calls for a clean and readable request definition.
- **Type Safety**: Response mapping to Go structs with automatic JSON unmarshaling.
- **Error Handling**: Custom error types for structured error responses.
- **Middleware Support**: Callbacks for request/response processing.

The following example demonstrates how to use `NewHTTPRequest` with the default `client.Client` implementation:

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/keboola/keboola-sdk-go/pkg/client"
	"github.com/keboola/keboola-sdk-go/pkg/request"
)

// APIError represents an API error response
type APIError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// Error implements the error interface
func (e APIError) Error() string {
	return fmt.Sprintf("API error %d: %s", e.Code, e.Message)
}

func main() {
	// Create a context
	ctx := context.Background()

	// Create a client (implements the request.Sender interface)
	c := client.New().
		WithBaseURL("https://api.example.com").
		WithHeader("Authorization", "Bearer your-token-here")

	// Define a struct for the response
	type User struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	// Create and send a GET request
	var user User
	_, _, err := request.NewHTTPRequest(c).
		WithGet("/users/123").
		WithResult(&user).
		Send(ctx)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("User: %+v\n", user)

	// Create and send a POST request with JSON body
	newUser := User{Name: "John Doe"}
	var createdUser User
	_, _, err = request.NewHTTPRequest(c).
		WithPost("/users").
		WithJSONBody(newUser).
		WithResult(&createdUser).
		Send(ctx)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("Created user: %+v\n", createdUser)

	// Handle errors with custom error type
	var result User
	_, _, err = request.NewHTTPRequest(c).
		WithGet("/users/999").
		WithResult(&result).
		WithError(&APIError{}).
		Send(ctx)
	if err != nil {
		fmt.Printf("Error handled: %v\n", err)
	}

	// Using query parameters
	var users []User
	_, _, err = request.NewHTTPRequest(c).
		WithGet("/users").
		AndQueryParam("limit", "10").
		AndQueryParam("offset", "0").
		WithResult(&users).
		Send(ctx)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("Users: %+v\n", users)

	// Using path parameters
	var project User
	_, _, err = request.NewHTTPRequest(c).
		WithGet("/users/{userId}/projects/{projectId}").
		AndPathParam("userId", "123").
		AndPathParam("projectId", "456").
		WithResult(&project).
		Send(ctx)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("Project: %+v\n", project)

	// Using callbacks
	_, _, err = request.NewHTTPRequest(c).
		WithGet("/users/123").
		WithResult(&user).
		WithOnSuccess(func(ctx context.Context, response request.HTTPResponse) error {
			fmt.Println("Request succeeded with status:", response.StatusCode())
			return nil
		}).
		WithOnError(func(ctx context.Context, response request.HTTPResponse, err error) error {
			fmt.Println("Request failed with status:", response.StatusCode())
			return err
		}).
		Send(ctx)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Using form body
	var loginResponse struct {
		Token string `json:"token"`
	}
	_, _, err = request.NewHTTPRequest(c).
		WithPost("/login").
		WithFormBody(map[string]string{
			"username": "user@example.com",
			"password": "password123",
		}).
		WithResult(&loginResponse).
		Send(ctx)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("Login token: %s\n", loginResponse.Token)
}
```

## Development

Clone the repository and run the dev container:
```shell
docker-compose run --rm -u "$UID:$GID" --service-ports dev bash
```

Run lint and tests in the container:
```shell
task lint
task tests
```
*Note: See "Test projects" section before running tests*

Run the HTTP server with documentation:
```shell
task godoc
```

Open `http://localhost:6060/pkg/github.com/keboola/keboola-sdk-go/pkg/` in your browser.

### Test projects

To successfully run all tests you will need test projects.

1. Create a `projects.json` file
   ```shell
   cp ./build/ci/projects.json projects.json
   ```
2. Replace token strings for each of the projects

## License

MIT licensed, see [LICENSE](./LICENSE) file.
