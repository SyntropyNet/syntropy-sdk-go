/*
Syntropy network API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
Contact: support@syntropynet.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syntropy

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// AgentsApiService AgentsApi service
type AgentsApiService service

type ApiV1NetworkAgentsCreateRequest struct {
	ctx                          context.Context
	ApiService                   *AgentsApiService
	v1NetworkAgentsCreateRequest *V1NetworkAgentsCreateRequest
}

func (r ApiV1NetworkAgentsCreateRequest) V1NetworkAgentsCreateRequest(v1NetworkAgentsCreateRequest V1NetworkAgentsCreateRequest) ApiV1NetworkAgentsCreateRequest {
	r.v1NetworkAgentsCreateRequest = &v1NetworkAgentsCreateRequest
	return r
}

func (r ApiV1NetworkAgentsCreateRequest) Execute() (*V1NetworkAgentsCreateResponse, *http.Response, error) {
	return r.ApiService.V1NetworkAgentsCreateExecute(r)
}

/*
V1NetworkAgentsCreate Create Agent

Creates new agent. Only `VIRTUAL` agents creation is supported. Not `VIRTUAL` agents are automatically created when first WebSocket
connection is established.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1NetworkAgentsCreateRequest
*/
func (a *AgentsApiService) V1NetworkAgentsCreate(ctx context.Context) ApiV1NetworkAgentsCreateRequest {
	return ApiV1NetworkAgentsCreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return V1NetworkAgentsCreateResponse
func (a *AgentsApiService) V1NetworkAgentsCreateExecute(r ApiV1NetworkAgentsCreateRequest) (*V1NetworkAgentsCreateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *V1NetworkAgentsCreateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentsApiService.V1NetworkAgentsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/network/agents"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.v1NetworkAgentsCreateRequest == nil {
		return localVarReturnValue, nil, reportError("v1NetworkAgentsCreateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.v1NetworkAgentsCreateRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["accessToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v V1ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v V1ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiV1NetworkAgentsGetRequest struct {
	ctx        context.Context
	ApiService *AgentsApiService
	skip       *int32
	take       *int32
	filter     *string
}

// Skip the number of returned Agents.
func (r ApiV1NetworkAgentsGetRequest) Skip(skip int32) ApiV1NetworkAgentsGetRequest {
	r.skip = &skip
	return r
}

// Limit the number of returned Agents.
func (r ApiV1NetworkAgentsGetRequest) Take(take int32) ApiV1NetworkAgentsGetRequest {
	r.take = &take
	return r
}

// Array of agent IDs.
func (r ApiV1NetworkAgentsGetRequest) Filter(filter string) ApiV1NetworkAgentsGetRequest {
	r.filter = &filter
	return r
}

func (r ApiV1NetworkAgentsGetRequest) Execute() (*V1NetworkAgentsGetResponse, *http.Response, error) {
	return r.ApiService.V1NetworkAgentsGetExecute(r)
}

/*
V1NetworkAgentsGet Get Agents

Gets a list of all created Agents, filtered and ordered according to the filter conditions provided in the request. The response may contain a fewer than take number of results. Filter is an `array of agents IDs`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1NetworkAgentsGetRequest
*/
func (a *AgentsApiService) V1NetworkAgentsGet(ctx context.Context) ApiV1NetworkAgentsGetRequest {
	return ApiV1NetworkAgentsGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return V1NetworkAgentsGetResponse
func (a *AgentsApiService) V1NetworkAgentsGetExecute(r ApiV1NetworkAgentsGetRequest) (*V1NetworkAgentsGetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *V1NetworkAgentsGetResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentsApiService.V1NetworkAgentsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/network/agents"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		localVarQueryParams.Add("skip", parameterToString(*r.skip, ""))
	}
	if r.take != nil {
		localVarQueryParams.Add("take", parameterToString(*r.take, ""))
	}
	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["accessToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v V1ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v V1ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiV1NetworkAgentsRemoveRequest struct {
	ctx                          context.Context
	ApiService                   *AgentsApiService
	v1NetworkAgentsRemoveRequest *V1NetworkAgentsRemoveRequest
}

func (r ApiV1NetworkAgentsRemoveRequest) V1NetworkAgentsRemoveRequest(v1NetworkAgentsRemoveRequest V1NetworkAgentsRemoveRequest) ApiV1NetworkAgentsRemoveRequest {
	r.v1NetworkAgentsRemoveRequest = &v1NetworkAgentsRemoveRequest
	return r
}

func (r ApiV1NetworkAgentsRemoveRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1NetworkAgentsRemoveExecute(r)
}

/*
V1NetworkAgentsRemove Delete Agents

Deletes Agents with existing connections. Learn more about platform agents [here](https://docs.syntropystack.com/docs/what-is-syntropy-agent).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1NetworkAgentsRemoveRequest
*/
func (a *AgentsApiService) V1NetworkAgentsRemove(ctx context.Context) ApiV1NetworkAgentsRemoveRequest {
	return ApiV1NetworkAgentsRemoveRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *AgentsApiService) V1NetworkAgentsRemoveExecute(r ApiV1NetworkAgentsRemoveRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentsApiService.V1NetworkAgentsRemove")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/network/agents/remove"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.v1NetworkAgentsRemoveRequest == nil {
		return nil, reportError("v1NetworkAgentsRemoveRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.v1NetworkAgentsRemoveRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["accessToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v V1ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v V1ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiV1NetworkAgentsUpdateRequest struct {
	ctx                          context.Context
	ApiService                   *AgentsApiService
	agentId                      int32
	v1NetworkAgentsUpdateRequest *V1NetworkAgentsUpdateRequest
}

func (r ApiV1NetworkAgentsUpdateRequest) V1NetworkAgentsUpdateRequest(v1NetworkAgentsUpdateRequest V1NetworkAgentsUpdateRequest) ApiV1NetworkAgentsUpdateRequest {
	r.v1NetworkAgentsUpdateRequest = &v1NetworkAgentsUpdateRequest
	return r
}

func (r ApiV1NetworkAgentsUpdateRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1NetworkAgentsUpdateExecute(r)
}

/*
V1NetworkAgentsUpdate Updates Agent

Updates the agent properties by `agent_id` (please make sure that none of the properties that you want to update are not locked: `GET` the agent and see `agent_locked_fields` ). You can add `tags` and also update the `provider` for the agent - please see the full [list of available providers](https://docs.syntropystack.com/docs/syntropy-agent-variables#syntropy-provider).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param agentId
 @return ApiV1NetworkAgentsUpdateRequest
*/
func (a *AgentsApiService) V1NetworkAgentsUpdate(ctx context.Context, agentId int32) ApiV1NetworkAgentsUpdateRequest {
	return ApiV1NetworkAgentsUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		agentId:    agentId,
	}
}

// Execute executes the request
func (a *AgentsApiService) V1NetworkAgentsUpdateExecute(r ApiV1NetworkAgentsUpdateRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentsApiService.V1NetworkAgentsUpdate")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/network/agents/{agent_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"agent_id"+"}", url.PathEscape(parameterToString(r.agentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.agentId < 1 {
		return nil, reportError("agentId must be greater than 1")
	}
	if r.agentId > 2147483647 {
		return nil, reportError("agentId must be less than 2147483647")
	}
	if r.v1NetworkAgentsUpdateRequest == nil {
		return nil, reportError("v1NetworkAgentsUpdateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.v1NetworkAgentsUpdateRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["accessToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v V1ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v V1ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
