# V1ConnectionServiceAgentService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentServiceSubnets** | [**[]ConnectionServiceAgentServiceSubnet**](ConnectionServiceAgentServiceSubnet.md) |  | 
**AgentServiceUdpPorts** | **[]int32** |  | 
**AgentServiceTcpPorts** | **[]int32** |  | 
**AgentServiceType** | [**AgentServiceTypes**](AgentServiceTypes.md) |  | 
**AgentServiceName** | **string** |  | 
**AgentServiceIsActive** | **bool** |  | 
**AgentServiceId** | **int32** |  | 
**AgentServiceCreatedAt** | **string** |  | 
**AgentServiceUpdatedAt** | **string** |  | 

## Methods

### NewV1ConnectionServiceAgentService

`func NewV1ConnectionServiceAgentService(agentServiceSubnets []ConnectionServiceAgentServiceSubnet, agentServiceUdpPorts []int32, agentServiceTcpPorts []int32, agentServiceType AgentServiceTypes, agentServiceName string, agentServiceIsActive bool, agentServiceId int32, agentServiceCreatedAt string, agentServiceUpdatedAt string, ) *V1ConnectionServiceAgentService`

NewV1ConnectionServiceAgentService instantiates a new V1ConnectionServiceAgentService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ConnectionServiceAgentServiceWithDefaults

`func NewV1ConnectionServiceAgentServiceWithDefaults() *V1ConnectionServiceAgentService`

NewV1ConnectionServiceAgentServiceWithDefaults instantiates a new V1ConnectionServiceAgentService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentServiceSubnets

`func (o *V1ConnectionServiceAgentService) GetAgentServiceSubnets() []ConnectionServiceAgentServiceSubnet`

GetAgentServiceSubnets returns the AgentServiceSubnets field if non-nil, zero value otherwise.

### GetAgentServiceSubnetsOk

`func (o *V1ConnectionServiceAgentService) GetAgentServiceSubnetsOk() (*[]ConnectionServiceAgentServiceSubnet, bool)`

GetAgentServiceSubnetsOk returns a tuple with the AgentServiceSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentServiceSubnets

`func (o *V1ConnectionServiceAgentService) SetAgentServiceSubnets(v []ConnectionServiceAgentServiceSubnet)`

SetAgentServiceSubnets sets AgentServiceSubnets field to given value.


### GetAgentServiceUdpPorts

`func (o *V1ConnectionServiceAgentService) GetAgentServiceUdpPorts() []int32`

GetAgentServiceUdpPorts returns the AgentServiceUdpPorts field if non-nil, zero value otherwise.

### GetAgentServiceUdpPortsOk

`func (o *V1ConnectionServiceAgentService) GetAgentServiceUdpPortsOk() (*[]int32, bool)`

GetAgentServiceUdpPortsOk returns a tuple with the AgentServiceUdpPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentServiceUdpPorts

`func (o *V1ConnectionServiceAgentService) SetAgentServiceUdpPorts(v []int32)`

SetAgentServiceUdpPorts sets AgentServiceUdpPorts field to given value.


### GetAgentServiceTcpPorts

`func (o *V1ConnectionServiceAgentService) GetAgentServiceTcpPorts() []int32`

GetAgentServiceTcpPorts returns the AgentServiceTcpPorts field if non-nil, zero value otherwise.

### GetAgentServiceTcpPortsOk

`func (o *V1ConnectionServiceAgentService) GetAgentServiceTcpPortsOk() (*[]int32, bool)`

GetAgentServiceTcpPortsOk returns a tuple with the AgentServiceTcpPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentServiceTcpPorts

`func (o *V1ConnectionServiceAgentService) SetAgentServiceTcpPorts(v []int32)`

SetAgentServiceTcpPorts sets AgentServiceTcpPorts field to given value.


### GetAgentServiceType

`func (o *V1ConnectionServiceAgentService) GetAgentServiceType() AgentServiceTypes`

GetAgentServiceType returns the AgentServiceType field if non-nil, zero value otherwise.

### GetAgentServiceTypeOk

`func (o *V1ConnectionServiceAgentService) GetAgentServiceTypeOk() (*AgentServiceTypes, bool)`

GetAgentServiceTypeOk returns a tuple with the AgentServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentServiceType

`func (o *V1ConnectionServiceAgentService) SetAgentServiceType(v AgentServiceTypes)`

SetAgentServiceType sets AgentServiceType field to given value.


### GetAgentServiceName

`func (o *V1ConnectionServiceAgentService) GetAgentServiceName() string`

GetAgentServiceName returns the AgentServiceName field if non-nil, zero value otherwise.

### GetAgentServiceNameOk

`func (o *V1ConnectionServiceAgentService) GetAgentServiceNameOk() (*string, bool)`

GetAgentServiceNameOk returns a tuple with the AgentServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentServiceName

`func (o *V1ConnectionServiceAgentService) SetAgentServiceName(v string)`

SetAgentServiceName sets AgentServiceName field to given value.


### GetAgentServiceIsActive

`func (o *V1ConnectionServiceAgentService) GetAgentServiceIsActive() bool`

GetAgentServiceIsActive returns the AgentServiceIsActive field if non-nil, zero value otherwise.

### GetAgentServiceIsActiveOk

`func (o *V1ConnectionServiceAgentService) GetAgentServiceIsActiveOk() (*bool, bool)`

GetAgentServiceIsActiveOk returns a tuple with the AgentServiceIsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentServiceIsActive

`func (o *V1ConnectionServiceAgentService) SetAgentServiceIsActive(v bool)`

SetAgentServiceIsActive sets AgentServiceIsActive field to given value.


### GetAgentServiceId

`func (o *V1ConnectionServiceAgentService) GetAgentServiceId() int32`

GetAgentServiceId returns the AgentServiceId field if non-nil, zero value otherwise.

### GetAgentServiceIdOk

`func (o *V1ConnectionServiceAgentService) GetAgentServiceIdOk() (*int32, bool)`

GetAgentServiceIdOk returns a tuple with the AgentServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentServiceId

`func (o *V1ConnectionServiceAgentService) SetAgentServiceId(v int32)`

SetAgentServiceId sets AgentServiceId field to given value.


### GetAgentServiceCreatedAt

`func (o *V1ConnectionServiceAgentService) GetAgentServiceCreatedAt() string`

GetAgentServiceCreatedAt returns the AgentServiceCreatedAt field if non-nil, zero value otherwise.

### GetAgentServiceCreatedAtOk

`func (o *V1ConnectionServiceAgentService) GetAgentServiceCreatedAtOk() (*string, bool)`

GetAgentServiceCreatedAtOk returns a tuple with the AgentServiceCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentServiceCreatedAt

`func (o *V1ConnectionServiceAgentService) SetAgentServiceCreatedAt(v string)`

SetAgentServiceCreatedAt sets AgentServiceCreatedAt field to given value.


### GetAgentServiceUpdatedAt

`func (o *V1ConnectionServiceAgentService) GetAgentServiceUpdatedAt() string`

GetAgentServiceUpdatedAt returns the AgentServiceUpdatedAt field if non-nil, zero value otherwise.

### GetAgentServiceUpdatedAtOk

`func (o *V1ConnectionServiceAgentService) GetAgentServiceUpdatedAtOk() (*string, bool)`

GetAgentServiceUpdatedAtOk returns a tuple with the AgentServiceUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentServiceUpdatedAt

`func (o *V1ConnectionServiceAgentService) SetAgentServiceUpdatedAt(v string)`

SetAgentServiceUpdatedAt sets AgentServiceUpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


