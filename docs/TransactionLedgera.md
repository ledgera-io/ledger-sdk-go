# \TransactionLedgera

All URIs are relative to *https://api.ledgera*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1TransactionIdGet**](TransactionLedgera.md#V1TransactionIdGet) | **Get** /v1/transaction/{id} | Get transaction
[**V1TransactionPost**](TransactionLedgera.md#V1TransactionPost) | **Post** /v1/transaction | Create transaction
[**V1TransactionRawPost**](TransactionLedgera.md#V1TransactionRawPost) | **Post** /v1/transaction/raw | Create raw transaction



## V1TransactionIdGet

> LedgerTransaction V1TransactionIdGet(ctx, id).Execute()

Get transaction



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ledgera-io/ledger-sdk-go"
)

func main() {
    id := "id_example" // string | Transaction ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionLedgera.V1TransactionIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionLedgera.V1TransactionIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1TransactionIdGet`: LedgerTransaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionLedgera.V1TransactionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1TransactionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LedgerTransaction**](LedgerTransaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1TransactionPost

> LedgerTransactionProcessed V1TransactionPost(ctx).Transaction(transaction).Execute()

Create transaction



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ledgera-io/ledger-sdk-go"
)

func main() {
    transaction := *openapiclient.NewLedgerTransactionParams("Amount_example", "Currency_example", "TransactionId_example", "TransactionProcess_example", "TransactionType_example") // LedgerTransactionParams | Transaction JSON

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionLedgera.V1TransactionPost(context.Background()).Transaction(transaction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionLedgera.V1TransactionPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1TransactionPost`: LedgerTransactionProcessed
    fmt.Fprintf(os.Stdout, "Response from `TransactionLedgera.V1TransactionPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TransactionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transaction** | [**LedgerTransactionParams**](LedgerTransactionParams.md) | Transaction JSON | 

### Return type

[**LedgerTransactionProcessed**](LedgerTransactionProcessed.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1TransactionRawPost

> LedgerTransactionProcessed V1TransactionRawPost(ctx).Transaction(transaction).Execute()

Create raw transaction



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ledgera-io/ledger-sdk-go"
)

func main() {
    transaction := *openapiclient.NewLedgerTransactionParams("Amount_example", "Currency_example", "TransactionId_example", "TransactionProcess_example", "TransactionType_example") // LedgerTransactionParams | Transaction JSON

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionLedgera.V1TransactionRawPost(context.Background()).Transaction(transaction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionLedgera.V1TransactionRawPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1TransactionRawPost`: LedgerTransactionProcessed
    fmt.Fprintf(os.Stdout, "Response from `TransactionLedgera.V1TransactionRawPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TransactionRawPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transaction** | [**LedgerTransactionParams**](LedgerTransactionParams.md) | Transaction JSON | 

### Return type

[**LedgerTransactionProcessed**](LedgerTransactionProcessed.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

