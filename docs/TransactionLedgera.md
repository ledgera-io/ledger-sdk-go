# \TransactionLedgera

All URIs are relative to *https://api.ledgera*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TransactionIdGet**](TransactionLedgera.md#TransactionIdGet) | **Get** /transaction/{id} | Get transaction
[**TransactionPost**](TransactionLedgera.md#TransactionPost) | **Post** /transaction | Create transaction
[**TransactionRawPost**](TransactionLedgera.md#TransactionRawPost) | **Post** /transaction/raw | Create raw transaction



## TransactionIdGet

> LedgerTransaction TransactionIdGet(ctx, id).Execute()

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
    resp, r, err := apiClient.TransactionLedgera.TransactionIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionLedgera.TransactionIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransactionIdGet`: LedgerTransaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionLedgera.TransactionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransactionIdGetRequest struct via the builder pattern


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


## TransactionPost

> LedgerTransactionProcessed TransactionPost(ctx).Transaction(transaction).Execute()

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
    resp, r, err := apiClient.TransactionLedgera.TransactionPost(context.Background()).Transaction(transaction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionLedgera.TransactionPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransactionPost`: LedgerTransactionProcessed
    fmt.Fprintf(os.Stdout, "Response from `TransactionLedgera.TransactionPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionPostRequest struct via the builder pattern


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


## TransactionRawPost

> LedgerTransactionProcessed TransactionRawPost(ctx).Transaction(transaction).Execute()

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
    resp, r, err := apiClient.TransactionLedgera.TransactionRawPost(context.Background()).Transaction(transaction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionLedgera.TransactionRawPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransactionRawPost`: LedgerTransactionProcessed
    fmt.Fprintf(os.Stdout, "Response from `TransactionLedgera.TransactionRawPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionRawPostRequest struct via the builder pattern


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

