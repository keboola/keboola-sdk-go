# \SUPERFileStorageManagementAPI

All URIs are relative to *https://connection.keboola.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewAWSS3Storage**](SUPERFileStorageManagementAPI.md#CreateNewAWSS3Storage) | **Post** /manage/file-storage-s3 | Create new AWS S3 Storage
[**CreateNewAzureBlobStorage**](SUPERFileStorageManagementAPI.md#CreateNewAzureBlobStorage) | **Post** /manage/file-storage-abs | Create new Azure Blob Storage
[**CreateNewGoogleCloudStorage**](SUPERFileStorageManagementAPI.md#CreateNewGoogleCloudStorage) | **Post** /manage/file-storage-gcs | Create new Google Cloud Storage
[**ListAzureBlobStorage**](SUPERFileStorageManagementAPI.md#ListAzureBlobStorage) | **Get** /manage/file-storage-abs | List Azure Blob Storage
[**ListGoogleCloudStorage**](SUPERFileStorageManagementAPI.md#ListGoogleCloudStorage) | **Get** /manage/file-storage-gcs | List Google Cloud Storage
[**ListStorages**](SUPERFileStorageManagementAPI.md#ListStorages) | **Get** /manage/file-storage-s3 | List storages
[**SetBlobStorageAsDefault**](SUPERFileStorageManagementAPI.md#SetBlobStorageAsDefault) | **Post** /manage/file-storage-abs/{file_storage_id}/default | Set Blob Storage as default
[**SetS3StorageAsDefault**](SUPERFileStorageManagementAPI.md#SetS3StorageAsDefault) | **Post** /manage/file-storage-s3/{file_storage_id}/default | Set S3 Storage as default



## CreateNewAWSS3Storage

> CreateNewAWSS3Storage201Response CreateNewAWSS3Storage(ctx).CreateNewAWSS3StorageRequest(createNewAWSS3StorageRequest).Execute()

Create new AWS S3 Storage



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/pkg/keboola/management"
)

func main() {
	createNewAWSS3StorageRequest := *openapiclient.NewCreateNewAWSS3StorageRequest("AwsKey_example", "FilesBucket_example", "Region_example", "Owner_example", "AwsSecret_example") // CreateNewAWSS3StorageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SUPERFileStorageManagementAPI.CreateNewAWSS3Storage(context.Background()).CreateNewAWSS3StorageRequest(createNewAWSS3StorageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERFileStorageManagementAPI.CreateNewAWSS3Storage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewAWSS3Storage`: CreateNewAWSS3Storage201Response
	fmt.Fprintf(os.Stdout, "Response from `SUPERFileStorageManagementAPI.CreateNewAWSS3Storage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewAWSS3StorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNewAWSS3StorageRequest** | [**CreateNewAWSS3StorageRequest**](CreateNewAWSS3StorageRequest.md) |  | 

### Return type

[**CreateNewAWSS3Storage201Response**](CreateNewAWSS3Storage201Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNewAzureBlobStorage

> CreateNewAzureBlobStorage201Response CreateNewAzureBlobStorage(ctx).CreateNewAzureBlobStorageRequest(createNewAzureBlobStorageRequest).Execute()

Create new Azure Blob Storage



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/pkg/keboola/management"
)

func main() {
	createNewAzureBlobStorageRequest := *openapiclient.NewCreateNewAzureBlobStorageRequest("AccountName_example", "Owner_example", "AccountKey_example") // CreateNewAzureBlobStorageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SUPERFileStorageManagementAPI.CreateNewAzureBlobStorage(context.Background()).CreateNewAzureBlobStorageRequest(createNewAzureBlobStorageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERFileStorageManagementAPI.CreateNewAzureBlobStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewAzureBlobStorage`: CreateNewAzureBlobStorage201Response
	fmt.Fprintf(os.Stdout, "Response from `SUPERFileStorageManagementAPI.CreateNewAzureBlobStorage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewAzureBlobStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNewAzureBlobStorageRequest** | [**CreateNewAzureBlobStorageRequest**](CreateNewAzureBlobStorageRequest.md) |  | 

### Return type

[**CreateNewAzureBlobStorage201Response**](CreateNewAzureBlobStorage201Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNewGoogleCloudStorage

> CreateNewGoogleCloudStorage201Response CreateNewGoogleCloudStorage(ctx).CreateNewGoogleCloudStorageRequest(createNewGoogleCloudStorageRequest).Execute()

Create new Google Cloud Storage



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/pkg/keboola/management"
)

func main() {
	createNewGoogleCloudStorageRequest := *openapiclient.NewCreateNewGoogleCloudStorageRequest("FilesBucket_example", "Owner_example", "Region_example") // CreateNewGoogleCloudStorageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SUPERFileStorageManagementAPI.CreateNewGoogleCloudStorage(context.Background()).CreateNewGoogleCloudStorageRequest(createNewGoogleCloudStorageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERFileStorageManagementAPI.CreateNewGoogleCloudStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewGoogleCloudStorage`: CreateNewGoogleCloudStorage201Response
	fmt.Fprintf(os.Stdout, "Response from `SUPERFileStorageManagementAPI.CreateNewGoogleCloudStorage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewGoogleCloudStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNewGoogleCloudStorageRequest** | [**CreateNewGoogleCloudStorageRequest**](CreateNewGoogleCloudStorageRequest.md) |  | 

### Return type

[**CreateNewGoogleCloudStorage201Response**](CreateNewGoogleCloudStorage201Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAzureBlobStorage

> []interface{} ListAzureBlobStorage(ctx).Execute()

List Azure Blob Storage



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/pkg/keboola/management"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SUPERFileStorageManagementAPI.ListAzureBlobStorage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERFileStorageManagementAPI.ListAzureBlobStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAzureBlobStorage`: []interface{}
	fmt.Fprintf(os.Stdout, "Response from `SUPERFileStorageManagementAPI.ListAzureBlobStorage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAzureBlobStorageRequest struct via the builder pattern


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


## ListGoogleCloudStorage

> []interface{} ListGoogleCloudStorage(ctx).Execute()

List Google Cloud Storage



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/pkg/keboola/management"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SUPERFileStorageManagementAPI.ListGoogleCloudStorage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERFileStorageManagementAPI.ListGoogleCloudStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGoogleCloudStorage`: []interface{}
	fmt.Fprintf(os.Stdout, "Response from `SUPERFileStorageManagementAPI.ListGoogleCloudStorage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListGoogleCloudStorageRequest struct via the builder pattern


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


## ListStorages

> []interface{} ListStorages(ctx).Execute()

List storages



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/pkg/keboola/management"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SUPERFileStorageManagementAPI.ListStorages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERFileStorageManagementAPI.ListStorages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStorages`: []interface{}
	fmt.Fprintf(os.Stdout, "Response from `SUPERFileStorageManagementAPI.ListStorages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListStoragesRequest struct via the builder pattern


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


## SetBlobStorageAsDefault

> SetBlobStorageAsDefault200Response SetBlobStorageAsDefault(ctx, fileStorageId).Execute()

Set Blob Storage as default



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/pkg/keboola/management"
)

func main() {
	fileStorageId := "fileStorageId_example" // string | File Storage ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SUPERFileStorageManagementAPI.SetBlobStorageAsDefault(context.Background(), fileStorageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERFileStorageManagementAPI.SetBlobStorageAsDefault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetBlobStorageAsDefault`: SetBlobStorageAsDefault200Response
	fmt.Fprintf(os.Stdout, "Response from `SUPERFileStorageManagementAPI.SetBlobStorageAsDefault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileStorageId** | **string** | File Storage ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetBlobStorageAsDefaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SetBlobStorageAsDefault200Response**](SetBlobStorageAsDefault200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetS3StorageAsDefault

> CreateNewAWSS3Storage201Response SetS3StorageAsDefault(ctx, fileStorageId).Execute()

Set S3 Storage as default



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/keboola/keboola-sdk-go/pkg/keboola/management"
)

func main() {
	fileStorageId := "fileStorageId_example" // string | File Storage ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SUPERFileStorageManagementAPI.SetS3StorageAsDefault(context.Background(), fileStorageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SUPERFileStorageManagementAPI.SetS3StorageAsDefault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetS3StorageAsDefault`: CreateNewAWSS3Storage201Response
	fmt.Fprintf(os.Stdout, "Response from `SUPERFileStorageManagementAPI.SetS3StorageAsDefault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileStorageId** | **string** | File Storage ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetS3StorageAsDefaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateNewAWSS3Storage201Response**](CreateNewAWSS3Storage201Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

