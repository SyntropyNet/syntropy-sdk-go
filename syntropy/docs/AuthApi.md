# \AuthApi

All URIs are relative to *https://api.syntropystack.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1NetworkAuthAccessTokensCreate**](AuthApi.md#V1NetworkAuthAccessTokensCreate) | **Post** /v1/network/auth/access-tokens | Create Access token
[**V1NetworkAuthAccessTokensDelete**](AuthApi.md#V1NetworkAuthAccessTokensDelete) | **Delete** /v1/network/auth/access-tokens/{access_token_id} | Delete Access token
[**V1NetworkAuthAccessTokensGet**](AuthApi.md#V1NetworkAuthAccessTokensGet) | **Get** /v1/network/auth/access-tokens | Get Access tokens
[**V1NetworkAuthAccessTokensPermissionsGet**](AuthApi.md#V1NetworkAuthAccessTokensPermissionsGet) | **Get** /v1/network/auth/access-tokens/permissions | Get Access token permissions
[**V1NetworkAuthAgentTokensCreate**](AuthApi.md#V1NetworkAuthAgentTokensCreate) | **Post** /v1/network/auth/agent-tokens | Create Agent Token
[**V1NetworkAuthAgentTokensDelete**](AuthApi.md#V1NetworkAuthAgentTokensDelete) | **Delete** /v1/network/auth/agent-tokens/{agent_token_id} | Delete Agent Token
[**V1NetworkAuthAgentTokensGet**](AuthApi.md#V1NetworkAuthAgentTokensGet) | **Get** /v1/network/auth/agent-tokens | Get Agent Tokens



## V1NetworkAuthAccessTokensCreate

> V1NetworkAuthAccessTokensCreateResponse V1NetworkAuthAccessTokensCreate(ctx).V1NetworkAuthAccessTokensCreateRequest(v1NetworkAuthAccessTokensCreateRequest).Execute()

Create Access token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    v1NetworkAuthAccessTokensCreateRequest := *openapiclient.NewV1NetworkAuthAccessTokensCreateRequest([]string{"Permissions_example"}, time.Now(), "AccessTokenName_example") // V1NetworkAuthAccessTokensCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.V1NetworkAuthAccessTokensCreate(context.Background()).V1NetworkAuthAccessTokensCreateRequest(v1NetworkAuthAccessTokensCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.V1NetworkAuthAccessTokensCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1NetworkAuthAccessTokensCreate`: V1NetworkAuthAccessTokensCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.V1NetworkAuthAccessTokensCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkAuthAccessTokensCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1NetworkAuthAccessTokensCreateRequest** | [**V1NetworkAuthAccessTokensCreateRequest**](V1NetworkAuthAccessTokensCreateRequest.md) |  | 

### Return type

[**V1NetworkAuthAccessTokensCreateResponse**](V1NetworkAuthAccessTokensCreateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1NetworkAuthAccessTokensDelete

> V1NetworkAuthAccessTokensDelete(ctx, accessTokenId).Execute()

Delete Access token



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
    accessTokenId := "accessTokenId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.V1NetworkAuthAccessTokensDelete(context.Background(), accessTokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.V1NetworkAuthAccessTokensDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessTokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkAuthAccessTokensDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[accessToken](../README.md#accessToken), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1NetworkAuthAccessTokensGet

> V1NetworkAuthAccessTokensGetResponse V1NetworkAuthAccessTokensGet(ctx).Skip(skip).Take(take).Order(order).Execute()

Get Access tokens



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
    skip := int32(56) // int32 | Number of tokens to skip. (optional) (default to 0)
    take := int32(56) // int32 | Limit of returned access tokens. (optional) (default to 20)
    order := "access_token_name:ASC" // string | Order: ascending or descending  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.V1NetworkAuthAccessTokensGet(context.Background()).Skip(skip).Take(take).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.V1NetworkAuthAccessTokensGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1NetworkAuthAccessTokensGet`: V1NetworkAuthAccessTokensGetResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.V1NetworkAuthAccessTokensGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkAuthAccessTokensGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Number of tokens to skip. | [default to 0]
 **take** | **int32** | Limit of returned access tokens. | [default to 20]
 **order** | **string** | Order: ascending or descending  | 

### Return type

[**V1NetworkAuthAccessTokensGetResponse**](V1NetworkAuthAccessTokensGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1NetworkAuthAccessTokensPermissionsGet

> V1NetworkAuthAccessTokensPermissionsGetResponse V1NetworkAuthAccessTokensPermissionsGet(ctx).Skip(skip).Take(take).Execute()

Get Access token permissions



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
    skip := int32(56) // int32 | Number of permission to be skipped. (optional) (default to 0)
    take := int32(56) // int32 | Limit the number of returned permissions. (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.V1NetworkAuthAccessTokensPermissionsGet(context.Background()).Skip(skip).Take(take).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.V1NetworkAuthAccessTokensPermissionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1NetworkAuthAccessTokensPermissionsGet`: V1NetworkAuthAccessTokensPermissionsGetResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.V1NetworkAuthAccessTokensPermissionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkAuthAccessTokensPermissionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Number of permission to be skipped. | [default to 0]
 **take** | **int32** | Limit the number of returned permissions. | [default to 20]

### Return type

[**V1NetworkAuthAccessTokensPermissionsGetResponse**](V1NetworkAuthAccessTokensPermissionsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1NetworkAuthAgentTokensCreate

> V1NetworkAuthAgentTokensCreateResponse V1NetworkAuthAgentTokensCreate(ctx).V1NetworkAuthAgentTokensCreateRequest(v1NetworkAuthAgentTokensCreateRequest).Execute()

Create Agent Token



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
    v1NetworkAuthAgentTokensCreateRequest := *openapiclient.NewV1NetworkAuthAgentTokensCreateRequest("AgentTokenName_example") // V1NetworkAuthAgentTokensCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.V1NetworkAuthAgentTokensCreate(context.Background()).V1NetworkAuthAgentTokensCreateRequest(v1NetworkAuthAgentTokensCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.V1NetworkAuthAgentTokensCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1NetworkAuthAgentTokensCreate`: V1NetworkAuthAgentTokensCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.V1NetworkAuthAgentTokensCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkAuthAgentTokensCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1NetworkAuthAgentTokensCreateRequest** | [**V1NetworkAuthAgentTokensCreateRequest**](V1NetworkAuthAgentTokensCreateRequest.md) |  | 

### Return type

[**V1NetworkAuthAgentTokensCreateResponse**](V1NetworkAuthAgentTokensCreateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1NetworkAuthAgentTokensDelete

> V1NetworkAuthAgentTokensDelete(ctx, agentTokenId).Execute()

Delete Agent Token



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
    agentTokenId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.V1NetworkAuthAgentTokensDelete(context.Background(), agentTokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.V1NetworkAuthAgentTokensDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentTokenId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkAuthAgentTokensDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[accessToken](../README.md#accessToken), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1NetworkAuthAgentTokensGet

> V1NetworkAuthAgentTokensGetResponse V1NetworkAuthAgentTokensGet(ctx).Skip(skip).Take(take).Order(order).Filter(filter).Execute()

Get Agent Tokens



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
    skip := int32(56) // int32 | Number of tokens to skip. (optional) (default to 0)
    take := int32(56) // int32 | Limit of returned tokens. (optional) (default to 20)
    order := "order_example" // string | string: 'ASC' | 'DESC' (optional)
    filter := "1,2,3" // string | Array of agent_token_ids (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.V1NetworkAuthAgentTokensGet(context.Background()).Skip(skip).Take(take).Order(order).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.V1NetworkAuthAgentTokensGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1NetworkAuthAgentTokensGet`: V1NetworkAuthAgentTokensGetResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.V1NetworkAuthAgentTokensGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkAuthAgentTokensGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Number of tokens to skip. | [default to 0]
 **take** | **int32** | Limit of returned tokens. | [default to 20]
 **order** | **string** | string: &#39;ASC&#39; | &#39;DESC&#39; | 
 **filter** | **string** | Array of agent_token_ids | 

### Return type

[**V1NetworkAuthAgentTokensGetResponse**](V1NetworkAuthAgentTokensGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

