# \TransactionRuleLedgera

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TransactionRuleBulkPost**](TransactionRuleLedgera.md#TransactionRuleBulkPost) | **Post** /transaction/rule/bulk | Create transaction rule
[**TransactionRuleGet**](TransactionRuleLedgera.md#TransactionRuleGet) | **Get** /transaction/rule | List transaction rule
[**TransactionRulePost**](TransactionRuleLedgera.md#TransactionRulePost) | **Post** /transaction/rule | Create transaction rule



## TransactionRuleBulkPost

> []LedgerTransactionRule TransactionRuleBulkPost(ctx).Transaction(transaction).Execute()

Create transaction rule



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
    transaction := []openapiclient.LedgerTransactionRule{*openapiclient.NewLedgerTransactionRule()} // []LedgerTransactionRule | Transaction rule JSON

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionRuleLedgera.TransactionRuleBulkPost(context.Background()).Transaction(transaction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionRuleLedgera.TransactionRuleBulkPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransactionRuleBulkPost`: []LedgerTransactionRule
    fmt.Fprintf(os.Stdout, "Response from `TransactionRuleLedgera.TransactionRuleBulkPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionRuleBulkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transaction** | [**[]LedgerTransactionRule**](LedgerTransactionRule.md) | Transaction rule JSON | 

### Return type

[**[]LedgerTransactionRule**](LedgerTransactionRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransactionRuleGet

> []LedgerTransactionRule TransactionRuleGet(ctx).Execute()

List transaction rule



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
    resp, r, err := apiClient.TransactionRuleLedgera.TransactionRuleGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionRuleLedgera.TransactionRuleGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransactionRuleGet`: []LedgerTransactionRule
    fmt.Fprintf(os.Stdout, "Response from `TransactionRuleLedgera.TransactionRuleGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTransactionRuleGetRequest struct via the builder pattern


### Return type

[**[]LedgerTransactionRule**](LedgerTransactionRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransactionRulePost

> LedgerTransactionRule TransactionRulePost(ctx).Transaction(transaction).Execute()

Create transaction rule



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
    transaction := *openapiclient.NewLedgerTransactionRuleParams("CreditAccountBalanceType_example", "CreditAccountSubType_example", "CreditAccountType_example", "DebitAccountBalanceType_example", "DebitAccountSubType_example", "DebitAccountType_example", int32(123), "EntryType_example", openapiclient.ledger.MessageType("single"), openapiclient.ledger.ProcessType("execution"), "TransactionType_example") // LedgerTransactionRuleParams | Transaction rule JSON

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionRuleLedgera.TransactionRulePost(context.Background()).Transaction(transaction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionRuleLedgera.TransactionRulePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransactionRulePost`: LedgerTransactionRule
    fmt.Fprintf(os.Stdout, "Response from `TransactionRuleLedgera.TransactionRulePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionRulePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transaction** | [**LedgerTransactionRuleParams**](LedgerTransactionRuleParams.md) | Transaction rule JSON | 

### Return type

[**LedgerTransactionRule**](LedgerTransactionRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

