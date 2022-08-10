# \AgentsApi

All URIs are relative to *https://api.syntropystack.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1NetworkAgentsCreate**](AgentsApi.md#V1NetworkAgentsCreate) | **Post** /v1/network/agents | Create Agent
[**V1NetworkAgentsGet**](AgentsApi.md#V1NetworkAgentsGet) | **Get** /v1/network/agents | Get Agents
[**V1NetworkAgentsRemove**](AgentsApi.md#V1NetworkAgentsRemove) | **Post** /v1/network/agents/remove | Delete Agents
[**V1NetworkAgentsSearch**](AgentsApi.md#V1NetworkAgentsSearch) | **Post** /v1/network/agents/search | Search Agents
[**V1NetworkAgentsUpdate**](AgentsApi.md#V1NetworkAgentsUpdate) | **Patch** /v1/network/agents/{agent_id} | Updates Agent



## V1NetworkAgentsCreate

> V1NetworkAgentsCreateResponse V1NetworkAgentsCreate(ctx).V1NetworkAgentsCreateRequest(v1NetworkAgentsCreateRequest).Execute()

Create Agent



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    v1NetworkAgentsCreateRequest := *openapiclient.NewV1NetworkAgentsCreateRequest("my_agent_token", "aws_1") // V1NetworkAgentsCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentsApi.V1NetworkAgentsCreate(context.Background()).V1NetworkAgentsCreateRequest(v1NetworkAgentsCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentsApi.V1NetworkAgentsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1NetworkAgentsCreate`: V1NetworkAgentsCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `AgentsApi.V1NetworkAgentsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkAgentsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1NetworkAgentsCreateRequest** | [**V1NetworkAgentsCreateRequest**](V1NetworkAgentsCreateRequest.md) |  | 

### Return type

[**V1NetworkAgentsCreateResponse**](V1NetworkAgentsCreateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1NetworkAgentsGet

> V1NetworkAgentsGetResponse V1NetworkAgentsGet(ctx).Skip(skip).Take(take).Filter(filter).Execute()

Get Agents



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    skip := int32(56) // int32 | Skip the number of returned Agents. (optional) (default to 0)
    take := int32(56) // int32 | Limit the number of returned Agents. (optional) (default to 20)
    filter := "1,2,3" // string | Array of agent IDs. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentsApi.V1NetworkAgentsGet(context.Background()).Skip(skip).Take(take).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentsApi.V1NetworkAgentsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1NetworkAgentsGet`: V1NetworkAgentsGetResponse
    fmt.Fprintf(os.Stdout, "Response from `AgentsApi.V1NetworkAgentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkAgentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Skip the number of returned Agents. | [default to 0]
 **take** | **int32** | Limit the number of returned Agents. | [default to 20]
 **filter** | **string** | Array of agent IDs. | 

### Return type

[**V1NetworkAgentsGetResponse**](V1NetworkAgentsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1NetworkAgentsRemove

> V1NetworkAgentsRemove(ctx).V1NetworkAgentsRemoveRequest(v1NetworkAgentsRemoveRequest).Execute()

Delete Agents



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    v1NetworkAgentsRemoveRequest := *openapiclient.NewV1NetworkAgentsRemoveRequest([]int32{int32(1)}) // V1NetworkAgentsRemoveRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentsApi.V1NetworkAgentsRemove(context.Background()).V1NetworkAgentsRemoveRequest(v1NetworkAgentsRemoveRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentsApi.V1NetworkAgentsRemove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkAgentsRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1NetworkAgentsRemoveRequest** | [**V1NetworkAgentsRemoveRequest**](V1NetworkAgentsRemoveRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[accessToken](../README.md#accessToken), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1NetworkAgentsSearch

> V1NetworkAgentsSearchResponse V1NetworkAgentsSearch(ctx).V1NetworkAgentsSearchRequest(v1NetworkAgentsSearchRequest).Execute()

Search Agents



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    v1NetworkAgentsSearchRequest := *openapiclient.NewV1NetworkAgentsSearchRequest() // V1NetworkAgentsSearchRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentsApi.V1NetworkAgentsSearch(context.Background()).V1NetworkAgentsSearchRequest(v1NetworkAgentsSearchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentsApi.V1NetworkAgentsSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1NetworkAgentsSearch`: V1NetworkAgentsSearchResponse
    fmt.Fprintf(os.Stdout, "Response from `AgentsApi.V1NetworkAgentsSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkAgentsSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1NetworkAgentsSearchRequest** | [**V1NetworkAgentsSearchRequest**](V1NetworkAgentsSearchRequest.md) |  | 

### Return type

[**V1NetworkAgentsSearchResponse**](V1NetworkAgentsSearchResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1NetworkAgentsUpdate

> V1NetworkAgentsUpdate(ctx, agentId).V1NetworkAgentsUpdateRequest(v1NetworkAgentsUpdateRequest).Execute()

Updates Agent



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    agentId := int32(56) // int32 | 
    v1NetworkAgentsUpdateRequest := *openapiclient.NewV1NetworkAgentsUpdateRequest() // V1NetworkAgentsUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentsApi.V1NetworkAgentsUpdate(context.Background(), agentId).V1NetworkAgentsUpdateRequest(v1NetworkAgentsUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentsApi.V1NetworkAgentsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkAgentsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1NetworkAgentsUpdateRequest** | [**V1NetworkAgentsUpdateRequest**](V1NetworkAgentsUpdateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[accessToken](../README.md#accessToken), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

