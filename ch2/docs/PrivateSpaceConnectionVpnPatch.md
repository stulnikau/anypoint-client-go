# PrivateSpaceConnectionVpnPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**RemoteAsn** | Pointer to **float32** | The ASN of the customer&#39;s outside network, used by VPNs utilizing BGP for routes.  | [optional] 
**LocalAsn** | Pointer to **float32** | The ASN of the customer&#39;s cloudhub network, used by VPNs utilizing BGP for routes.  | [optional] 
**RemoteIpAddress** | **string** | The IP address of the customer&#39;s VPN device on their outside network  | 
**StaticRoutes** | **[]string** | The CIDRs that the Cloudhub VPC will send traffic to via the VPN.  | 
**VpnTunnels** | [**[]Items7**](Items7.md) | VPN tunnel configuration that the customer wants to be applied to their VPN connection.  | 

## Methods

### NewPrivateSpaceConnectionVpnPatch

`func NewPrivateSpaceConnectionVpnPatch(id string, remoteIpAddress string, staticRoutes []string, vpnTunnels []Items7, ) *PrivateSpaceConnectionVpnPatch`

NewPrivateSpaceConnectionVpnPatch instantiates a new PrivateSpaceConnectionVpnPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateSpaceConnectionVpnPatchWithDefaults

`func NewPrivateSpaceConnectionVpnPatchWithDefaults() *PrivateSpaceConnectionVpnPatch`

NewPrivateSpaceConnectionVpnPatchWithDefaults instantiates a new PrivateSpaceConnectionVpnPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrivateSpaceConnectionVpnPatch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateSpaceConnectionVpnPatch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateSpaceConnectionVpnPatch) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PrivateSpaceConnectionVpnPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateSpaceConnectionVpnPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateSpaceConnectionVpnPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PrivateSpaceConnectionVpnPatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRemoteAsn

`func (o *PrivateSpaceConnectionVpnPatch) GetRemoteAsn() float32`

GetRemoteAsn returns the RemoteAsn field if non-nil, zero value otherwise.

### GetRemoteAsnOk

`func (o *PrivateSpaceConnectionVpnPatch) GetRemoteAsnOk() (*float32, bool)`

GetRemoteAsnOk returns a tuple with the RemoteAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAsn

`func (o *PrivateSpaceConnectionVpnPatch) SetRemoteAsn(v float32)`

SetRemoteAsn sets RemoteAsn field to given value.

### HasRemoteAsn

`func (o *PrivateSpaceConnectionVpnPatch) HasRemoteAsn() bool`

HasRemoteAsn returns a boolean if a field has been set.

### GetLocalAsn

`func (o *PrivateSpaceConnectionVpnPatch) GetLocalAsn() float32`

GetLocalAsn returns the LocalAsn field if non-nil, zero value otherwise.

### GetLocalAsnOk

`func (o *PrivateSpaceConnectionVpnPatch) GetLocalAsnOk() (*float32, bool)`

GetLocalAsnOk returns a tuple with the LocalAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAsn

`func (o *PrivateSpaceConnectionVpnPatch) SetLocalAsn(v float32)`

SetLocalAsn sets LocalAsn field to given value.

### HasLocalAsn

`func (o *PrivateSpaceConnectionVpnPatch) HasLocalAsn() bool`

HasLocalAsn returns a boolean if a field has been set.

### GetRemoteIpAddress

`func (o *PrivateSpaceConnectionVpnPatch) GetRemoteIpAddress() string`

GetRemoteIpAddress returns the RemoteIpAddress field if non-nil, zero value otherwise.

### GetRemoteIpAddressOk

`func (o *PrivateSpaceConnectionVpnPatch) GetRemoteIpAddressOk() (*string, bool)`

GetRemoteIpAddressOk returns a tuple with the RemoteIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIpAddress

`func (o *PrivateSpaceConnectionVpnPatch) SetRemoteIpAddress(v string)`

SetRemoteIpAddress sets RemoteIpAddress field to given value.


### GetStaticRoutes

`func (o *PrivateSpaceConnectionVpnPatch) GetStaticRoutes() []string`

GetStaticRoutes returns the StaticRoutes field if non-nil, zero value otherwise.

### GetStaticRoutesOk

`func (o *PrivateSpaceConnectionVpnPatch) GetStaticRoutesOk() (*[]string, bool)`

GetStaticRoutesOk returns a tuple with the StaticRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticRoutes

`func (o *PrivateSpaceConnectionVpnPatch) SetStaticRoutes(v []string)`

SetStaticRoutes sets StaticRoutes field to given value.


### GetVpnTunnels

`func (o *PrivateSpaceConnectionVpnPatch) GetVpnTunnels() []Items7`

GetVpnTunnels returns the VpnTunnels field if non-nil, zero value otherwise.

### GetVpnTunnelsOk

`func (o *PrivateSpaceConnectionVpnPatch) GetVpnTunnelsOk() (*[]Items7, bool)`

GetVpnTunnelsOk returns a tuple with the VpnTunnels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnTunnels

`func (o *PrivateSpaceConnectionVpnPatch) SetVpnTunnels(v []Items7)`

SetVpnTunnels sets VpnTunnels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


