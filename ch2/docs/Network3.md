# Network3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **string** | The region of the Private Space. Required when create a Private Space network. | [optional] 
**VpcId** | Pointer to **string** | The existing VPC where the Private Space will be created. Required when create a Private Space in existing Anypoint VPC.  | [optional] 
**CidrBlock** | Pointer to **string** | The CIDR block of the VPC where the Private Space will be created. Required when create a Private Space in a new VPC. Default is 10.0.0.0/16.  | [optional] [default to "10.0.0.0/16"]
**InternalDns** | Pointer to [**InternalDns9**](InternalDns9.md) |  | [optional] 

## Methods

### NewNetwork3

`func NewNetwork3() *Network3`

NewNetwork3 instantiates a new Network3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetwork3WithDefaults

`func NewNetwork3WithDefaults() *Network3`

NewNetwork3WithDefaults instantiates a new Network3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *Network3) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Network3) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Network3) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Network3) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetVpcId

`func (o *Network3) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *Network3) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *Network3) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *Network3) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### GetCidrBlock

`func (o *Network3) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *Network3) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *Network3) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.

### HasCidrBlock

`func (o *Network3) HasCidrBlock() bool`

HasCidrBlock returns a boolean if a field has been set.

### GetInternalDns

`func (o *Network3) GetInternalDns() InternalDns9`

GetInternalDns returns the InternalDns field if non-nil, zero value otherwise.

### GetInternalDnsOk

`func (o *Network3) GetInternalDnsOk() (*InternalDns9, bool)`

GetInternalDnsOk returns a tuple with the InternalDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalDns

`func (o *Network3) SetInternalDns(v InternalDns9)`

SetInternalDns sets InternalDns field to given value.

### HasInternalDns

`func (o *Network3) HasInternalDns() bool`

HasInternalDns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


