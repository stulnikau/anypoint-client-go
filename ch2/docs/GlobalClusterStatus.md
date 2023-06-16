# GlobalClusterStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Infra** | Pointer to [**Infra**](Infra.md) |  | [optional] 
**Fabric** | Pointer to [**Fabric**](Fabric.md) |  | [optional] 
**Ingress** | Pointer to [**Ingress**](Ingress.md) |  | [optional] 

## Methods

### NewGlobalClusterStatus

`func NewGlobalClusterStatus() *GlobalClusterStatus`

NewGlobalClusterStatus instantiates a new GlobalClusterStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalClusterStatusWithDefaults

`func NewGlobalClusterStatusWithDefaults() *GlobalClusterStatus`

NewGlobalClusterStatusWithDefaults instantiates a new GlobalClusterStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfra

`func (o *GlobalClusterStatus) GetInfra() Infra`

GetInfra returns the Infra field if non-nil, zero value otherwise.

### GetInfraOk

`func (o *GlobalClusterStatus) GetInfraOk() (*Infra, bool)`

GetInfraOk returns a tuple with the Infra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfra

`func (o *GlobalClusterStatus) SetInfra(v Infra)`

SetInfra sets Infra field to given value.

### HasInfra

`func (o *GlobalClusterStatus) HasInfra() bool`

HasInfra returns a boolean if a field has been set.

### GetFabric

`func (o *GlobalClusterStatus) GetFabric() Fabric`

GetFabric returns the Fabric field if non-nil, zero value otherwise.

### GetFabricOk

`func (o *GlobalClusterStatus) GetFabricOk() (*Fabric, bool)`

GetFabricOk returns a tuple with the Fabric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabric

`func (o *GlobalClusterStatus) SetFabric(v Fabric)`

SetFabric sets Fabric field to given value.

### HasFabric

`func (o *GlobalClusterStatus) HasFabric() bool`

HasFabric returns a boolean if a field has been set.

### GetIngress

`func (o *GlobalClusterStatus) GetIngress() Ingress`

GetIngress returns the Ingress field if non-nil, zero value otherwise.

### GetIngressOk

`func (o *GlobalClusterStatus) GetIngressOk() (*Ingress, bool)`

GetIngressOk returns a tuple with the Ingress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngress

`func (o *GlobalClusterStatus) SetIngress(v Ingress)`

SetIngress sets Ingress field to given value.

### HasIngress

`func (o *GlobalClusterStatus) HasIngress() bool`

HasIngress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


