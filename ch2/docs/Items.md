# Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the VPN.  | 
**VpnId** | Pointer to **string** | The id of the VPN.  | [optional] 
**ConnectionId** | Pointer to **string** | The connectionId associated to the vpn, to which it belongs.  | [optional] 
**ConnectionName** | Pointer to **string** | The connectionName associated to the vpn, to which it belongs.  | [optional] 
**VpnConnectionStatus** | Pointer to **string** | The status of vpn(pending | available | deleting | deleted).  | [optional] 
**RemoteAsn** | Pointer to **float32** | The ASN of the customer&#39;s outside network, used by VPNs utilizing BGP for routes.  | [optional] 
**LocalAsn** | Pointer to **float32** | The ASN of the customer&#39;s cloudhub network, used by VPNs utilizing BGP for routes.  | [optional] 
**RemoteIpAddress** | **string** | The IP address of the customer&#39;s VPN device on their outside network  | 
**StaticRoutes** | **[]string** | The CIDRs that the Cloudhub VPC will send traffic to via the VPN.  | 
**VpnTunnels** | [**[]Items7**](Items7.md) | VPN tunnel actual state of customer&#39;s VPN connection. | 

## Methods

### NewItems

`func NewItems(name string, remoteIpAddress string, staticRoutes []string, vpnTunnels []Items7, ) *Items`

NewItems instantiates a new Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemsWithDefaults

`func NewItemsWithDefaults() *Items`

NewItemsWithDefaults instantiates a new Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Items) SetName(v string)`

SetName sets Name field to given value.


### GetVpnId

`func (o *Items) GetVpnId() string`

GetVpnId returns the VpnId field if non-nil, zero value otherwise.

### GetVpnIdOk

`func (o *Items) GetVpnIdOk() (*string, bool)`

GetVpnIdOk returns a tuple with the VpnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnId

`func (o *Items) SetVpnId(v string)`

SetVpnId sets VpnId field to given value.

### HasVpnId

`func (o *Items) HasVpnId() bool`

HasVpnId returns a boolean if a field has been set.

### GetConnectionId

`func (o *Items) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *Items) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *Items) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *Items) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetConnectionName

`func (o *Items) GetConnectionName() string`

GetConnectionName returns the ConnectionName field if non-nil, zero value otherwise.

### GetConnectionNameOk

`func (o *Items) GetConnectionNameOk() (*string, bool)`

GetConnectionNameOk returns a tuple with the ConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionName

`func (o *Items) SetConnectionName(v string)`

SetConnectionName sets ConnectionName field to given value.

### HasConnectionName

`func (o *Items) HasConnectionName() bool`

HasConnectionName returns a boolean if a field has been set.

### GetVpnConnectionStatus

`func (o *Items) GetVpnConnectionStatus() string`

GetVpnConnectionStatus returns the VpnConnectionStatus field if non-nil, zero value otherwise.

### GetVpnConnectionStatusOk

`func (o *Items) GetVpnConnectionStatusOk() (*string, bool)`

GetVpnConnectionStatusOk returns a tuple with the VpnConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnConnectionStatus

`func (o *Items) SetVpnConnectionStatus(v string)`

SetVpnConnectionStatus sets VpnConnectionStatus field to given value.

### HasVpnConnectionStatus

`func (o *Items) HasVpnConnectionStatus() bool`

HasVpnConnectionStatus returns a boolean if a field has been set.

### GetRemoteAsn

`func (o *Items) GetRemoteAsn() float32`

GetRemoteAsn returns the RemoteAsn field if non-nil, zero value otherwise.

### GetRemoteAsnOk

`func (o *Items) GetRemoteAsnOk() (*float32, bool)`

GetRemoteAsnOk returns a tuple with the RemoteAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAsn

`func (o *Items) SetRemoteAsn(v float32)`

SetRemoteAsn sets RemoteAsn field to given value.

### HasRemoteAsn

`func (o *Items) HasRemoteAsn() bool`

HasRemoteAsn returns a boolean if a field has been set.

### GetLocalAsn

`func (o *Items) GetLocalAsn() float32`

GetLocalAsn returns the LocalAsn field if non-nil, zero value otherwise.

### GetLocalAsnOk

`func (o *Items) GetLocalAsnOk() (*float32, bool)`

GetLocalAsnOk returns a tuple with the LocalAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAsn

`func (o *Items) SetLocalAsn(v float32)`

SetLocalAsn sets LocalAsn field to given value.

### HasLocalAsn

`func (o *Items) HasLocalAsn() bool`

HasLocalAsn returns a boolean if a field has been set.

### GetRemoteIpAddress

`func (o *Items) GetRemoteIpAddress() string`

GetRemoteIpAddress returns the RemoteIpAddress field if non-nil, zero value otherwise.

### GetRemoteIpAddressOk

`func (o *Items) GetRemoteIpAddressOk() (*string, bool)`

GetRemoteIpAddressOk returns a tuple with the RemoteIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIpAddress

`func (o *Items) SetRemoteIpAddress(v string)`

SetRemoteIpAddress sets RemoteIpAddress field to given value.


### GetStaticRoutes

`func (o *Items) GetStaticRoutes() []string`

GetStaticRoutes returns the StaticRoutes field if non-nil, zero value otherwise.

### GetStaticRoutesOk

`func (o *Items) GetStaticRoutesOk() (*[]string, bool)`

GetStaticRoutesOk returns a tuple with the StaticRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticRoutes

`func (o *Items) SetStaticRoutes(v []string)`

SetStaticRoutes sets StaticRoutes field to given value.


### GetVpnTunnels

`func (o *Items) GetVpnTunnels() []Items7`

GetVpnTunnels returns the VpnTunnels field if non-nil, zero value otherwise.

### GetVpnTunnelsOk

`func (o *Items) GetVpnTunnelsOk() (*[]Items7, bool)`

GetVpnTunnelsOk returns a tuple with the VpnTunnels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnTunnels

`func (o *Items) SetVpnTunnels(v []Items7)`

SetVpnTunnels sets VpnTunnels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


