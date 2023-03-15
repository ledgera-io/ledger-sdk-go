# \AccountApi

All URIs are relative to *https://api.ledgera.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountGet**](AccountApi.md#AccountGet) | **Get** /account | 
[**AccountPost**](AccountApi.md#AccountPost) | **Post** /account | 
[**V1AccountIdJournalGet**](AccountApi.md#V1AccountIdJournalGet) | **Get** /v1/account/{id}/journal | 



## AccountGet

> Account AccountGet(ctx).Execute()





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
    resp, r, err := apiClient.AccountApi.AccountGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.AccountGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountGet`: Account
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.AccountGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAccountGetRequest struct via the builder pattern


### Return type

[**Account**](Account.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountPost

> Account AccountPost(ctx).AccountPostRequest(accountPostRequest).Execute()





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
    accountPostRequest := *openapiclient.NewAccountPostRequest(openapiclient.accountType("liability"), "Label_example") // AccountPostRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountApi.AccountPost(context.Background()).AccountPostRequest(accountPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.AccountPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountPost`: Account
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.AccountPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountPostRequest** | [**AccountPostRequest**](AccountPostRequest.md) |  | 

### Return type

[**Account**](Account.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json:
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountIdJournalGet

> V1AccountIdJournalGet200Response V1AccountIdJournalGet(ctx, id).Limit(limit).Cursor(cursor).DateFrom(dateFrom).DateTo(dateTo).EntryType(entryType).Execute()



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
    id := "684936a4-f1f5-48ed-9378-ca9e31cf76a7" // string | account ID
    limit := int32(1) // int32 |  (optional)
    cursor := "MjAyMy0wMy0wMiAxNzowMzozOS4zOTU3NzYrMDAwMzdlOTk4NS05ZmM0LTQwMDgtOWFjMy04YmE4NzA4MTViYjE=" // string |  (optional)
    dateFrom := "2023-03-13T21:33:07.202Z" // string |  (optional)
    dateTo := "2023-03-13T21:33:07.202Z" // string |  (optional)
    entryType := "DEBIT" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountApi.V1AccountIdJournalGet(context.Background(), id).Limit(limit).Cursor(cursor).DateFrom(dateFrom).DateTo(dateTo).EntryType(entryType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.V1AccountIdJournalGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AccountIdJournalGet`: V1AccountIdJournalGet200Response
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.V1AccountIdJournalGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountIdJournalGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **cursor** | **string** |  | 
 **dateFrom** | **string** |  | 
 **dateTo** | **string** |  | 
 **entryType** | **string** |  | 

### Return type

[**V1AccountIdJournalGet200Response**](V1AccountIdJournalGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

