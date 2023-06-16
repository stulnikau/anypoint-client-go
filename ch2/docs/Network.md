# Network

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **string** | The region of the Private Space. Required when create a Private Space network. | [optional] 
**VpcId** | Pointer to **string** | The existing VPC where the Private Space will be created. Required when create a Private Space in existing Anypoint VPC.  | [optional] 
**CidrBlock** | Pointer to **string** | The CIDR block of the VPC where the Private Space will be created. Required when create a Private Space in a new VPC. Default is 10.0.0.0/16.  | [optional] [default to "10.0.0.0/16"]
**InternalDns** | Pointer to [**InternalDns**](InternalDns.md) |  | [optional] 
**InboundStaticIps** | Pointer to **[]string** | The inbound static IPs of the Private Space network.  | [optional] 
**OutboundStaticIps** | Pointer to **[]string** | The outbound static IPs of the Private Space network.  | [optional] 
**DnsTarget** | Pointer to **string** | The dns target of the Private Space network.  | [optional] 
**InternalDnsTarget** | Pointer to **string** | The internal dns target of the VPC that the Private Space network uses.  | [optional] 

## Methods

### NewNetwork

`func NewNetwork() *Network`

NewNetwork instantiates a new Network object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkWithDefaults

`func NewNetworkWithDefaults() *Network`

NewNetworkWithDefaults instantiates a new Network object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *Network) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Network) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Network) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Network) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetVpcId

`func (o *Network) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *Network) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *Network) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *Network) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### GetCidrBlock

`func (o *Network) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *Network) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *Network) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.

### HasCidrBlock

`func (o *Network) HasCidrBlock() bool`

HasCidrBlock returns a boolean if a field has been set.

### GetInternalDns

`func (o *Network) GetInternalDns() InternalDns`

GetInternalDns returns the InternalDns field if non-nil, zero value otherwise.

### GetInternalDnsOk

`func (o *Network) GetInternalDnsOk() (*InternalDns, bool)`

GetInternalDnsOk returns a tuple with the InternalDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalDns

`func (o *Network) SetInternalDns(v InternalDns)`

SetInternalDns sets InternalDns field to given value.

### HasInternalDns

`func (o *Network) HasInternalDns() bool`

HasInternalDns returns a boolean if a field has been set.

### GetInboundStaticIps

`func (o *Network) GetInboundStaticIps() []string`

GetInboundStaticIps returns the InboundStaticIps field if non-nil, zero value otherwise.

### GetInboundStaticIpsOk

`func (o *Network) GetInboundStaticIpsOk() (*[]string, bool)`

GetInboundStaticIpsOk returns a tuple with the InboundStaticIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundStaticIps

`func (o *Network) SetInboundStaticIps(v []string)`

SetInboundStaticIps sets InboundStaticIps field to given value.

### HasInboundStaticIps

`func (o *Network) HasInboundStaticIps() bool`

HasInboundStaticIps returns a boolean if a field has been set.

### GetOutboundStaticIps

`func (o *Network) GetOutboundStaticIps() []string`

GetOutboundStaticIps returns the OutboundStaticIps field if non-nil, zero value otherwise.

### GetOutboundStaticIpsOk

`func (o *Network) GetOutboundStaticIpsOk() (*[]string, bool)`

GetOutboundStaticIpsOk returns a tuple with the OutboundStaticIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundStaticIps

`func (o *Network) SetOutboundStaticIps(v []string)`

SetOutboundStaticIps sets OutboundStaticIps field to given value.

### HasOutboundStaticIps

`func (o *Network) HasOutboundStaticIps() bool`

HasOutboundStaticIps returns a boolean if a field has been set.

### GetDnsTarget

`func (o *Network) GetDnsTarget() string`

GetDnsTarget returns the DnsTarget field if non-nil, zero value otherwise.

### GetDnsTargetOk

`func (o *Network) GetDnsTargetOk() (*string, bool)`

GetDnsTargetOk returns a tuple with the DnsTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsTarget

`func (o *Network) SetDnsTarget(v string)`

SetDnsTarget sets DnsTarget field to given value.

### HasDnsTarget

`func (o *Network) HasDnsTarget() bool`

HasDnsTarget returns a boolean if a field has been set.

### GetInternalDnsTarget

`func (o *Network) GetInternalDnsTarget() string`

GetInternalDnsTarget returns the InternalDnsTarget field if non-nil, zero value otherwise.

### GetInternalDnsTargetOk

`func (o *Network) GetInternalDnsTargetOk() (*string, bool)`

GetInternalDnsTargetOk returns a tuple with the InternalDnsTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalDnsTarget

`func (o *Network) SetInternalDnsTarget(v string)`

SetInternalDnsTarget sets InternalDnsTarget field to given value.

### HasInternalDnsTarget

`func (o *Network) HasInternalDnsTarget() bool`

HasInternalDnsTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


