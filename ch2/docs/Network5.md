# Network5

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **string** | The region of the Private Space. Required when create a Private Space network. | [optional] 
**VpcId** | Pointer to **string** | The existing VPC where the Private Space will be created. Required when create a Private Space in existing Anypoint VPC.  | [optional] 
**CidrBlock** | Pointer to **string** | The CIDR block of the VPC where the Private Space will be created. Required when create a Private Space in a new VPC. Default is 10.0.0.0/16.  | [optional] [default to "10.0.0.0/16"]
**InternalDns** | Pointer to [**InternalDns9**](InternalDns9.md) |  | [optional] 

## Methods

### NewNetwork5

`func NewNetwork5() *Network5`

NewNetwork5 instantiates a new Network5 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetwork5WithDefaults

`func NewNetwork5WithDefaults() *Network5`

NewNetwork5WithDefaults instantiates a new Network5 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *Network5) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Network5) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Network5) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Network5) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetVpcId

`func (o *Network5) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *Network5) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *Network5) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *Network5) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### GetCidrBlock

`func (o *Network5) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *Network5) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *Network5) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.

### HasCidrBlock

`func (o *Network5) HasCidrBlock() bool`

HasCidrBlock returns a boolean if a field has been set.

### GetInternalDns

`func (o *Network5) GetInternalDns() InternalDns9`

GetInternalDns returns the InternalDns field if non-nil, zero value otherwise.

### GetInternalDnsOk

`func (o *Network5) GetInternalDnsOk() (*InternalDns9, bool)`

GetInternalDnsOk returns a tuple with the InternalDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalDns

`func (o *Network5) SetInternalDns(v InternalDns9)`

SetInternalDns sets InternalDns field to given value.

### HasInternalDns

`func (o *Network5) HasInternalDns() bool`

HasInternalDns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


