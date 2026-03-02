# \SUPERStorageBackendsManagementAPI

All URIs are relative to *https://connection.keboola.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateSnowflakeBackend**](SUPERStorageBackendsManagementAPI.md#ActivateSnowflakeBackend) | **Post** /manage/storage-backend/{backend_id}/activate | Activate Snowflake backend
[**BackendDetail**](SUPERStorageBackendsManagementAPI.md#BackendDetail) | **Get** /manage/storage-backend/{backend_id} | Backend detail
[**CreateANewBackend**](SUPERStorageBackendsManagementAPI.md#CreateANewBackend) | **Post** /manage/storage-backend | Create a new backend
[**CreateANewBigQueryBackend**](SUPERStorageBackendsManagementAPI.md#CreateANewBigQueryBackend) | **Post** /manage/storage-backend/bigquery | Create a new BigQuery backend
[**CreateSnowflakeBackendWithCert**](SUPERStorageBackendsManagementAPI.md#CreateSnowflakeBackendWithCert) | **Post** /manage/storage-backend/snowflake | Create a new Snowflake backend with certificate authentication
[**DeleteBackend**](SUPERStorageBackendsManagementAPI.md#DeleteBackend) | **Delete** /manage/storage-backend/{backend_id} | Delete backend
[**ListBackends**](SUPERStorageBackendsManagementAPI.md#ListBackends) | **Get** /manage/storage-backend | List backends
[**UpdateBackend**](SUPERStorageBackendsManagementAPI.md#UpdateBackend) | **Patch** /manage/storage-backend/{backend_id} | Update backend
[**UpdateBigQueryBackend**](SUPERStorageBackendsManagementAPI.md#UpdateBigQueryBackend) | **Patch** /manage/storage-backend/bigquery/{backend_id} | Update BigQuery backend



## ActivateSnowflakeBackend

> ActivateSnowflakeStorageBackendResponse ActivateSnowflakeBackend(ctx, backendId).Execute()

Activate Snowflake backend



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	backendId := "backendId_example" // string | Id of storage backend

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SUPERStorageBackendsManagementAPI.ActivateSnowflakeBackend(context.Background(), backendId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERStorageBackendsManagementAPI.ActivateSnowflakeBackend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivateSnowflakeBackend`: ActivateSnowflakeStorageBackendResponse
	fmt.Fprintf(os.Stdout, "Response from `SUPERStorageBackendsManagementAPI.ActivateSnowflakeBackend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backendId** | **string** | Id of storage backend | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateSnowflakeBackendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActivateSnowflakeStorageBackendResponse**](ActivateSnowflakeStorageBackendResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackendDetail

> []interface{} BackendDetail(ctx, backendId).Execute()

Backend detail



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	backendId := "backendId_example" // string | Backend ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SUPERStorageBackendsManagementAPI.BackendDetail(context.Background(), backendId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERStorageBackendsManagementAPI.BackendDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackendDetail`: []interface{}
	fmt.Fprintf(os.Stdout, "Response from `SUPERStorageBackendsManagementAPI.BackendDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backendId** | **string** | Backend ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackendDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateANewBackend

> CreateANewBackend201Response CreateANewBackend(ctx).CreateANewBackendRequest(createANewBackendRequest).Execute()

Create a new backend



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	createANewBackendRequest := *openapiclient.NewCreateANewBackendRequest("Backend_example", "Host_example", "Username_example", "Password_example", "Region_example", "Owner_example") // CreateANewBackendRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SUPERStorageBackendsManagementAPI.CreateANewBackend(context.Background()).CreateANewBackendRequest(createANewBackendRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERStorageBackendsManagementAPI.CreateANewBackend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateANewBackend`: CreateANewBackend201Response
	fmt.Fprintf(os.Stdout, "Response from `SUPERStorageBackendsManagementAPI.CreateANewBackend`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateANewBackendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createANewBackendRequest** | [**CreateANewBackendRequest**](CreateANewBackendRequest.md) |  | 

### Return type

[**CreateANewBackend201Response**](CreateANewBackend201Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateANewBigQueryBackend

> CreateANewBigQueryBackend201Response CreateANewBigQueryBackend(ctx).CreateANewBigQueryBackendRequest(createANewBigQueryBackendRequest).Execute()

Create a new BigQuery backend



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	createANewBigQueryBackendRequest := *openapiclient.NewCreateANewBigQueryBackendRequest("Owner_example", "FolderId_example", "Region_example") // CreateANewBigQueryBackendRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SUPERStorageBackendsManagementAPI.CreateANewBigQueryBackend(context.Background()).CreateANewBigQueryBackendRequest(createANewBigQueryBackendRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERStorageBackendsManagementAPI.CreateANewBigQueryBackend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateANewBigQueryBackend`: CreateANewBigQueryBackend201Response
	fmt.Fprintf(os.Stdout, "Response from `SUPERStorageBackendsManagementAPI.CreateANewBigQueryBackend`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateANewBigQueryBackendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createANewBigQueryBackendRequest** | [**CreateANewBigQueryBackendRequest**](CreateANewBigQueryBackendRequest.md) |  | 

### Return type

[**CreateANewBigQueryBackend201Response**](CreateANewBigQueryBackend201Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSnowflakeBackendWithCert

> CreateSnowflakeStorageBackendResponse CreateSnowflakeBackendWithCert(ctx).CreateSnowflakeBackendWithCertRequest(createSnowflakeBackendWithCertRequest).Execute()

Create a new Snowflake backend with certificate authentication



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	createSnowflakeBackendWithCertRequest := *openapiclient.NewCreateSnowflakeBackendWithCertRequest("demo.snowflakecomputing.com", "KEBOOLA", "us-east-1", "keboola", "keboola") // CreateSnowflakeBackendWithCertRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SUPERStorageBackendsManagementAPI.CreateSnowflakeBackendWithCert(context.Background()).CreateSnowflakeBackendWithCertRequest(createSnowflakeBackendWithCertRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERStorageBackendsManagementAPI.CreateSnowflakeBackendWithCert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSnowflakeBackendWithCert`: CreateSnowflakeStorageBackendResponse
	fmt.Fprintf(os.Stdout, "Response from `SUPERStorageBackendsManagementAPI.CreateSnowflakeBackendWithCert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSnowflakeBackendWithCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSnowflakeBackendWithCertRequest** | [**CreateSnowflakeBackendWithCertRequest**](CreateSnowflakeBackendWithCertRequest.md) |  | 

### Return type

[**CreateSnowflakeStorageBackendResponse**](CreateSnowflakeStorageBackendResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBackend

> DeleteBackend(ctx, backendId).Execute()

Delete backend



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	backendId := "backendId_example" // string | Id of storage backend

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SUPERStorageBackendsManagementAPI.DeleteBackend(context.Background(), backendId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERStorageBackendsManagementAPI.DeleteBackend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backendId** | **string** | Id of storage backend | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBackendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBackends

> []interface{} ListBackends(ctx).Execute()

List backends



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SUPERStorageBackendsManagementAPI.ListBackends(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERStorageBackendsManagementAPI.ListBackends``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBackends`: []interface{}
	fmt.Fprintf(os.Stdout, "Response from `SUPERStorageBackendsManagementAPI.ListBackends`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBackendsRequest struct via the builder pattern


### Return type

**[]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBackend

> CreateANewBackend201Response UpdateBackend(ctx, backendId).StorageBackendUpdate(storageBackendUpdate).Execute()

Update backend



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	backendId := "backendId_example" // string | Backend ID
	storageBackendUpdate := *openapiclient.NewStorageBackendUpdate() // StorageBackendUpdate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SUPERStorageBackendsManagementAPI.UpdateBackend(context.Background(), backendId).StorageBackendUpdate(storageBackendUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERStorageBackendsManagementAPI.UpdateBackend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBackend`: CreateANewBackend201Response
	fmt.Fprintf(os.Stdout, "Response from `SUPERStorageBackendsManagementAPI.UpdateBackend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backendId** | **string** | Backend ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBackendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storageBackendUpdate** | [**StorageBackendUpdate**](StorageBackendUpdate.md) |  | 

### Return type

[**CreateANewBackend201Response**](CreateANewBackend201Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBigQueryBackend

> CreateANewBigQueryBackend201Response UpdateBigQueryBackend(ctx, backendId).BigQueryStorageBackendUpdate(bigQueryStorageBackendUpdate).Execute()

Update BigQuery backend



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	backendId := "backendId_example" // string | Backend ID
	bigQueryStorageBackendUpdate := *openapiclient.NewBigQueryStorageBackendUpdate() // BigQueryStorageBackendUpdate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SUPERStorageBackendsManagementAPI.UpdateBigQueryBackend(context.Background(), backendId).BigQueryStorageBackendUpdate(bigQueryStorageBackendUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERStorageBackendsManagementAPI.UpdateBigQueryBackend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBigQueryBackend`: CreateANewBigQueryBackend201Response
	fmt.Fprintf(os.Stdout, "Response from `SUPERStorageBackendsManagementAPI.UpdateBigQueryBackend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backendId** | **string** | Backend ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBigQueryBackendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bigQueryStorageBackendUpdate** | [**BigQueryStorageBackendUpdate**](BigQueryStorageBackendUpdate.md) |  | 

### Return type

[**CreateANewBigQueryBackend201Response**](CreateANewBigQueryBackend201Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

