# \ConnectionsApi

All URIs are relative to *https://api.syntropystack.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1NetworkConnectionsCreateMesh**](ConnectionsApi.md#V1NetworkConnectionsCreateMesh) | **Post** /v1/network/connections/mesh | Create Connections Mesh
[**V1NetworkConnectionsCreateP2P**](ConnectionsApi.md#V1NetworkConnectionsCreateP2P) | **Post** /v1/network/connections/point-to-point | Create P2P Connections
[**V1NetworkConnectionsGet**](ConnectionsApi.md#V1NetworkConnectionsGet) | **Get** /v1/network/connections | Get Connections
[**V1NetworkConnectionsRemove**](ConnectionsApi.md#V1NetworkConnectionsRemove) | **Post** /v1/network/connections/remove | Delete Connections
[**V1NetworkConnectionsSearch**](ConnectionsApi.md#V1NetworkConnectionsSearch) | **Post** /v1/network/connections/search | Search Connections
[**V1NetworkConnectionsServicesGet**](ConnectionsApi.md#V1NetworkConnectionsServicesGet) | **Get** /v1/network/connections/services | Get Connection services
[**V1NetworkConnectionsServicesUpdate**](ConnectionsApi.md#V1NetworkConnectionsServicesUpdate) | **Patch** /v1/network/connections/services | Update Connection services
[**V1NetworkConnectionsUpdate**](ConnectionsApi.md#V1NetworkConnectionsUpdate) | **Patch** /v1/network/connections | Update Connections



## V1NetworkConnectionsCreateMesh

> V1NetworkConnectionsCreateResponse V1NetworkConnectionsCreateMesh(ctx).V1NetworkConnectionsCreateMeshRequest(v1NetworkConnectionsCreateMeshRequest).Execute()

Create Connections Mesh



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
    v1NetworkConnectionsCreateMeshRequest := *openapiclient.NewV1NetworkConnectionsCreateMeshRequest([]openapiclient.V1NetworkConnectionsCreateMeshRequestAgentIds{*openapiclient.NewV1NetworkConnectionsCreateMeshRequestAgentIds(int32(1))}) // V1NetworkConnectionsCreateMeshRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionsApi.V1NetworkConnectionsCreateMesh(context.Background()).V1NetworkConnectionsCreateMeshRequest(v1NetworkConnectionsCreateMeshRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.V1NetworkConnectionsCreateMesh``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1NetworkConnectionsCreateMesh`: V1NetworkConnectionsCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.V1NetworkConnectionsCreateMesh`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkConnectionsCreateMeshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1NetworkConnectionsCreateMeshRequest** | [**V1NetworkConnectionsCreateMeshRequest**](V1NetworkConnectionsCreateMeshRequest.md) |  | 

### Return type

[**V1NetworkConnectionsCreateResponse**](V1NetworkConnectionsCreateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1NetworkConnectionsCreateP2P

> V1NetworkConnectionsCreateResponse V1NetworkConnectionsCreateP2P(ctx).V1NetworkConnectionsCreateP2PRequest(v1NetworkConnectionsCreateP2PRequest).Execute()

Create P2P Connections



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
    v1NetworkConnectionsCreateP2PRequest := *openapiclient.NewV1NetworkConnectionsCreateP2PRequest([]openapiclient.V1NetworkConnectionsCreateP2PRequestAgentPairs{*openapiclient.NewV1NetworkConnectionsCreateP2PRequestAgentPairs(int32(1), int32(1))}) // V1NetworkConnectionsCreateP2PRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionsApi.V1NetworkConnectionsCreateP2P(context.Background()).V1NetworkConnectionsCreateP2PRequest(v1NetworkConnectionsCreateP2PRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.V1NetworkConnectionsCreateP2P``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1NetworkConnectionsCreateP2P`: V1NetworkConnectionsCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.V1NetworkConnectionsCreateP2P`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkConnectionsCreateP2PRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1NetworkConnectionsCreateP2PRequest** | [**V1NetworkConnectionsCreateP2PRequest**](V1NetworkConnectionsCreateP2PRequest.md) |  | 

### Return type

[**V1NetworkConnectionsCreateResponse**](V1NetworkConnectionsCreateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1NetworkConnectionsGet

> V1NetworkConnectionsGetResponse V1NetworkConnectionsGet(ctx).Skip(skip).Take(take).Filter(filter).Execute()

Get Connections



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
    skip := int32(56) // int32 |  (optional) (default to 0)
    take := int32(56) // int32 |  (optional) (default to 20)
    filter := "1,2,3" // string | array of Agent Connection ids (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionsApi.V1NetworkConnectionsGet(context.Background()).Skip(skip).Take(take).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.V1NetworkConnectionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1NetworkConnectionsGet`: V1NetworkConnectionsGetResponse
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.V1NetworkConnectionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkConnectionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** |  | [default to 0]
 **take** | **int32** |  | [default to 20]
 **filter** | **string** | array of Agent Connection ids | 

### Return type

[**V1NetworkConnectionsGetResponse**](V1NetworkConnectionsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1NetworkConnectionsRemove

> V1NetworkConnectionsRemove(ctx).V1NetworkConnectionsRemoveRequest(v1NetworkConnectionsRemoveRequest).Execute()

Delete Connections



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
    v1NetworkConnectionsRemoveRequest := *openapiclient.NewV1NetworkConnectionsRemoveRequest([]int32{int32(1)}) // V1NetworkConnectionsRemoveRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionsApi.V1NetworkConnectionsRemove(context.Background()).V1NetworkConnectionsRemoveRequest(v1NetworkConnectionsRemoveRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.V1NetworkConnectionsRemove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkConnectionsRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1NetworkConnectionsRemoveRequest** | [**V1NetworkConnectionsRemoveRequest**](V1NetworkConnectionsRemoveRequest.md) |  | 

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


## V1NetworkConnectionsSearch

> V1NetworkConnectionsSearchResponse V1NetworkConnectionsSearch(ctx).V1NetworkConnectionsSearchRequest(v1NetworkConnectionsSearchRequest).Execute()

Search Connections



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
    v1NetworkConnectionsSearchRequest := *openapiclient.NewV1NetworkConnectionsSearchRequest() // V1NetworkConnectionsSearchRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionsApi.V1NetworkConnectionsSearch(context.Background()).V1NetworkConnectionsSearchRequest(v1NetworkConnectionsSearchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.V1NetworkConnectionsSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1NetworkConnectionsSearch`: V1NetworkConnectionsSearchResponse
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.V1NetworkConnectionsSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkConnectionsSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1NetworkConnectionsSearchRequest** | [**V1NetworkConnectionsSearchRequest**](V1NetworkConnectionsSearchRequest.md) |  | 

### Return type

[**V1NetworkConnectionsSearchResponse**](V1NetworkConnectionsSearchResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1NetworkConnectionsServicesGet

> V1NetworkConnectionsServicesGetResponse V1NetworkConnectionsServicesGet(ctx).Filter(filter).Skip(skip).Take(take).Execute()

Get Connection services



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
    filter := "1,2,3" // string | array of Agent Connection group ids (optional)
    skip := int32(56) // int32 |  (optional) (default to 0)
    take := int32(56) // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionsApi.V1NetworkConnectionsServicesGet(context.Background()).Filter(filter).Skip(skip).Take(take).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.V1NetworkConnectionsServicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1NetworkConnectionsServicesGet`: V1NetworkConnectionsServicesGetResponse
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.V1NetworkConnectionsServicesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkConnectionsServicesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | array of Agent Connection group ids | 
 **skip** | **int32** |  | [default to 0]
 **take** | **int32** |  | [default to 20]

### Return type

[**V1NetworkConnectionsServicesGetResponse**](V1NetworkConnectionsServicesGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1NetworkConnectionsServicesUpdate

> V1NetworkConnectionsServicesUpdate(ctx).V1NetworkConnectionsServicesUpdateRequest(v1NetworkConnectionsServicesUpdateRequest).Execute()

Update Connection services



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
    v1NetworkConnectionsServicesUpdateRequest := *openapiclient.NewV1NetworkConnectionsServicesUpdateRequest([]openapiclient.AgentServicesUpdateChanges{*openapiclient.NewAgentServicesUpdateChanges(int32(1), false)}) // V1NetworkConnectionsServicesUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionsApi.V1NetworkConnectionsServicesUpdate(context.Background()).V1NetworkConnectionsServicesUpdateRequest(v1NetworkConnectionsServicesUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.V1NetworkConnectionsServicesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkConnectionsServicesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1NetworkConnectionsServicesUpdateRequest** | [**V1NetworkConnectionsServicesUpdateRequest**](V1NetworkConnectionsServicesUpdateRequest.md) |  | 

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


## V1NetworkConnectionsUpdate

> V1NetworkConnectionsUpdate(ctx).V1NetworkConnectionsUpdateRequest(v1NetworkConnectionsUpdateRequest).Execute()

Update Connections



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
    v1NetworkConnectionsUpdateRequest := *openapiclient.NewV1NetworkConnectionsUpdateRequest([]openapiclient.V1ConnectionUpdateChange{*openapiclient.NewV1ConnectionUpdateChange(int32(1), false)}) // V1NetworkConnectionsUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionsApi.V1NetworkConnectionsUpdate(context.Background()).V1NetworkConnectionsUpdateRequest(v1NetworkConnectionsUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.V1NetworkConnectionsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkConnectionsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1NetworkConnectionsUpdateRequest** | [**V1NetworkConnectionsUpdateRequest**](V1NetworkConnectionsUpdateRequest.md) |  | 

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

