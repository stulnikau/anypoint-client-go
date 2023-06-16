# Items7

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpnConnectionId** | Pointer to **string** | Connection id for the vpn.  | [optional] 
**Psk** | Pointer to **string** | The pre-shared key that the customer will use to make a connection.  | [optional] 
**PtpCidr** | Pointer to **string** | The Point-To-Point CIDR range that the customer will use for connections.  | [optional] 
**LocalExternalIpAddress** | Pointer to **string** | Local external ip address for the tunnel.  | [optional] 
**LocalPtpIpAddress** | Pointer to **string** | Local Point-To-Point ip address for tunnel.  | [optional] 
**RemotePtpIpAddress** | Pointer to **string** | Remote Point-To-Point ip address for tunnel.  | [optional] 
**StartupAction** | Pointer to **string** | The action to take when the establishing the tunnel for the VPN connection. By default, CGW must initiate the IKE negotiation and bring up the tunnel. Specify start for AWS to initiate the IKE negotiation.  | [optional] 
**DPDTimeoutAction** | Pointer to **string** | The action to take after Dead Peer Detection timeout occurs. Specify restart to restart the IKE initiation. Specify clear to end the IKE session.  | [optional] 
**RekeyMarginInSeconds** | Pointer to **int32** | The margin time, in seconds, before the phase 2 lifetime expires, during which the AWS side of the VPN connection performs an IKE rekey.  | [optional] 
**RekeyFuzz** | Pointer to **int32** | The percentage of the rekey window (determined by RekeyMarginTimeSeconds) during which the rekey time is randomly selected.  | [optional] 
**IkeVersions** | Pointer to **[]string** | The IKE versions that are permitted for the VPN tunnel. Valid values: ikev1,ikev2  | [optional] 
**Phase1DhGroups** | Pointer to **[]int32** | One or more Diffie-Hellman group numbers that are permitted for the VPN tunnel for phase 1 IKE negotiations. Valid values: 2,14,15,16,17,18,19,20,21,22,23,24  | [optional] 
**Phase2DhGroups** | Pointer to **[]int32** | One or more Diffie-Hellman group numbers that are permitted for the VPN tunnel for phase 2 IKE negotiations. Valid values: 2,5,14,15,16,17,18,19,20,21,22,23,24  | [optional] 
**Phase1EncryptionAlgorithms** | Pointer to **[]string** | One or more encryption algorithms that are permitted for the VPN tunnel for phase 1 IKE negotiations. Valid values: AES128,AES256,AES128-GCM-16,AES256-GCM-16  | [optional] 
**Phase2EncryptionAlgorithms** | Pointer to **[]string** | One or more encryption algorithms that are permitted for the VPN tunnel for phase 2 IKE negotiations. Valid values: AES128,AES256,AES128-GCM-16,AES256-GCM-16  | [optional] 
**Phase1IntegrityAlgorithms** | Pointer to **[]string** | One or more integrity algorithms that are permitted for the VPN tunnel for phase 1 IKE negotiations. Valid values: SHA1,SHA2-256,SHA2-384,SHA2-512  | [optional] 
**Phase2IntegrityAlgorithms** | Pointer to **[]string** | One or more integrity algorithms that are permitted for the VPN tunnel for phase 2 IKE negotiations. Valid values: SHA1,SHA2-256,SHA2-384,SHA2-512  | [optional] 
**AcceptedRouteCount** | Pointer to **float32** | Accepted route count for the tunnel.  | [optional] 
**LastStatusChange** | Pointer to **string** | Last status change for the tunnel.  | [optional] 
**Status** | Pointer to **string** | Status for the tunnel.  | [optional] 
**StatusMessage** | Pointer to **string** | Status message for the tunnel.  | [optional] 

## Methods

### NewItems7

`func NewItems7() *Items7`

NewItems7 instantiates a new Items7 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItems7WithDefaults

`func NewItems7WithDefaults() *Items7`

NewItems7WithDefaults instantiates a new Items7 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpnConnectionId

`func (o *Items7) GetVpnConnectionId() string`

GetVpnConnectionId returns the VpnConnectionId field if non-nil, zero value otherwise.

### GetVpnConnectionIdOk

`func (o *Items7) GetVpnConnectionIdOk() (*string, bool)`

GetVpnConnectionIdOk returns a tuple with the VpnConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnConnectionId

`func (o *Items7) SetVpnConnectionId(v string)`

SetVpnConnectionId sets VpnConnectionId field to given value.

### HasVpnConnectionId

`func (o *Items7) HasVpnConnectionId() bool`

HasVpnConnectionId returns a boolean if a field has been set.

### GetPsk

`func (o *Items7) GetPsk() string`

GetPsk returns the Psk field if non-nil, zero value otherwise.

### GetPskOk

`func (o *Items7) GetPskOk() (*string, bool)`

GetPskOk returns a tuple with the Psk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsk

`func (o *Items7) SetPsk(v string)`

SetPsk sets Psk field to given value.

### HasPsk

`func (o *Items7) HasPsk() bool`

HasPsk returns a boolean if a field has been set.

### GetPtpCidr

`func (o *Items7) GetPtpCidr() string`

GetPtpCidr returns the PtpCidr field if non-nil, zero value otherwise.

### GetPtpCidrOk

`func (o *Items7) GetPtpCidrOk() (*string, bool)`

GetPtpCidrOk returns a tuple with the PtpCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtpCidr

`func (o *Items7) SetPtpCidr(v string)`

SetPtpCidr sets PtpCidr field to given value.

### HasPtpCidr

`func (o *Items7) HasPtpCidr() bool`

HasPtpCidr returns a boolean if a field has been set.

### GetLocalExternalIpAddress

`func (o *Items7) GetLocalExternalIpAddress() string`

GetLocalExternalIpAddress returns the LocalExternalIpAddress field if non-nil, zero value otherwise.

### GetLocalExternalIpAddressOk

`func (o *Items7) GetLocalExternalIpAddressOk() (*string, bool)`

GetLocalExternalIpAddressOk returns a tuple with the LocalExternalIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalExternalIpAddress

`func (o *Items7) SetLocalExternalIpAddress(v string)`

SetLocalExternalIpAddress sets LocalExternalIpAddress field to given value.

### HasLocalExternalIpAddress

`func (o *Items7) HasLocalExternalIpAddress() bool`

HasLocalExternalIpAddress returns a boolean if a field has been set.

### GetLocalPtpIpAddress

`func (o *Items7) GetLocalPtpIpAddress() string`

GetLocalPtpIpAddress returns the LocalPtpIpAddress field if non-nil, zero value otherwise.

### GetLocalPtpIpAddressOk

`func (o *Items7) GetLocalPtpIpAddressOk() (*string, bool)`

GetLocalPtpIpAddressOk returns a tuple with the LocalPtpIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPtpIpAddress

`func (o *Items7) SetLocalPtpIpAddress(v string)`

SetLocalPtpIpAddress sets LocalPtpIpAddress field to given value.

### HasLocalPtpIpAddress

`func (o *Items7) HasLocalPtpIpAddress() bool`

HasLocalPtpIpAddress returns a boolean if a field has been set.

### GetRemotePtpIpAddress

`func (o *Items7) GetRemotePtpIpAddress() string`

GetRemotePtpIpAddress returns the RemotePtpIpAddress field if non-nil, zero value otherwise.

### GetRemotePtpIpAddressOk

`func (o *Items7) GetRemotePtpIpAddressOk() (*string, bool)`

GetRemotePtpIpAddressOk returns a tuple with the RemotePtpIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePtpIpAddress

`func (o *Items7) SetRemotePtpIpAddress(v string)`

SetRemotePtpIpAddress sets RemotePtpIpAddress field to given value.

### HasRemotePtpIpAddress

`func (o *Items7) HasRemotePtpIpAddress() bool`

HasRemotePtpIpAddress returns a boolean if a field has been set.

### GetStartupAction

`func (o *Items7) GetStartupAction() string`

GetStartupAction returns the StartupAction field if non-nil, zero value otherwise.

### GetStartupActionOk

`func (o *Items7) GetStartupActionOk() (*string, bool)`

GetStartupActionOk returns a tuple with the StartupAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupAction

`func (o *Items7) SetStartupAction(v string)`

SetStartupAction sets StartupAction field to given value.

### HasStartupAction

`func (o *Items7) HasStartupAction() bool`

HasStartupAction returns a boolean if a field has been set.

### GetDPDTimeoutAction

`func (o *Items7) GetDPDTimeoutAction() string`

GetDPDTimeoutAction returns the DPDTimeoutAction field if non-nil, zero value otherwise.

### GetDPDTimeoutActionOk

`func (o *Items7) GetDPDTimeoutActionOk() (*string, bool)`

GetDPDTimeoutActionOk returns a tuple with the DPDTimeoutAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDPDTimeoutAction

`func (o *Items7) SetDPDTimeoutAction(v string)`

SetDPDTimeoutAction sets DPDTimeoutAction field to given value.

### HasDPDTimeoutAction

`func (o *Items7) HasDPDTimeoutAction() bool`

HasDPDTimeoutAction returns a boolean if a field has been set.

### GetRekeyMarginInSeconds

`func (o *Items7) GetRekeyMarginInSeconds() int32`

GetRekeyMarginInSeconds returns the RekeyMarginInSeconds field if non-nil, zero value otherwise.

### GetRekeyMarginInSecondsOk

`func (o *Items7) GetRekeyMarginInSecondsOk() (*int32, bool)`

GetRekeyMarginInSecondsOk returns a tuple with the RekeyMarginInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRekeyMarginInSeconds

`func (o *Items7) SetRekeyMarginInSeconds(v int32)`

SetRekeyMarginInSeconds sets RekeyMarginInSeconds field to given value.

### HasRekeyMarginInSeconds

`func (o *Items7) HasRekeyMarginInSeconds() bool`

HasRekeyMarginInSeconds returns a boolean if a field has been set.

### GetRekeyFuzz

`func (o *Items7) GetRekeyFuzz() int32`

GetRekeyFuzz returns the RekeyFuzz field if non-nil, zero value otherwise.

### GetRekeyFuzzOk

`func (o *Items7) GetRekeyFuzzOk() (*int32, bool)`

GetRekeyFuzzOk returns a tuple with the RekeyFuzz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRekeyFuzz

`func (o *Items7) SetRekeyFuzz(v int32)`

SetRekeyFuzz sets RekeyFuzz field to given value.

### HasRekeyFuzz

`func (o *Items7) HasRekeyFuzz() bool`

HasRekeyFuzz returns a boolean if a field has been set.

### GetIkeVersions

`func (o *Items7) GetIkeVersions() []string`

GetIkeVersions returns the IkeVersions field if non-nil, zero value otherwise.

### GetIkeVersionsOk

`func (o *Items7) GetIkeVersionsOk() (*[]string, bool)`

GetIkeVersionsOk returns a tuple with the IkeVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeVersions

`func (o *Items7) SetIkeVersions(v []string)`

SetIkeVersions sets IkeVersions field to given value.

### HasIkeVersions

`func (o *Items7) HasIkeVersions() bool`

HasIkeVersions returns a boolean if a field has been set.

### GetPhase1DhGroups

`func (o *Items7) GetPhase1DhGroups() []int32`

GetPhase1DhGroups returns the Phase1DhGroups field if non-nil, zero value otherwise.

### GetPhase1DhGroupsOk

`func (o *Items7) GetPhase1DhGroupsOk() (*[]int32, bool)`

GetPhase1DhGroupsOk returns a tuple with the Phase1DhGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1DhGroups

`func (o *Items7) SetPhase1DhGroups(v []int32)`

SetPhase1DhGroups sets Phase1DhGroups field to given value.

### HasPhase1DhGroups

`func (o *Items7) HasPhase1DhGroups() bool`

HasPhase1DhGroups returns a boolean if a field has been set.

### GetPhase2DhGroups

`func (o *Items7) GetPhase2DhGroups() []int32`

GetPhase2DhGroups returns the Phase2DhGroups field if non-nil, zero value otherwise.

### GetPhase2DhGroupsOk

`func (o *Items7) GetPhase2DhGroupsOk() (*[]int32, bool)`

GetPhase2DhGroupsOk returns a tuple with the Phase2DhGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2DhGroups

`func (o *Items7) SetPhase2DhGroups(v []int32)`

SetPhase2DhGroups sets Phase2DhGroups field to given value.

### HasPhase2DhGroups

`func (o *Items7) HasPhase2DhGroups() bool`

HasPhase2DhGroups returns a boolean if a field has been set.

### GetPhase1EncryptionAlgorithms

`func (o *Items7) GetPhase1EncryptionAlgorithms() []string`

GetPhase1EncryptionAlgorithms returns the Phase1EncryptionAlgorithms field if non-nil, zero value otherwise.

### GetPhase1EncryptionAlgorithmsOk

`func (o *Items7) GetPhase1EncryptionAlgorithmsOk() (*[]string, bool)`

GetPhase1EncryptionAlgorithmsOk returns a tuple with the Phase1EncryptionAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1EncryptionAlgorithms

`func (o *Items7) SetPhase1EncryptionAlgorithms(v []string)`

SetPhase1EncryptionAlgorithms sets Phase1EncryptionAlgorithms field to given value.

### HasPhase1EncryptionAlgorithms

`func (o *Items7) HasPhase1EncryptionAlgorithms() bool`

HasPhase1EncryptionAlgorithms returns a boolean if a field has been set.

### GetPhase2EncryptionAlgorithms

`func (o *Items7) GetPhase2EncryptionAlgorithms() []string`

GetPhase2EncryptionAlgorithms returns the Phase2EncryptionAlgorithms field if non-nil, zero value otherwise.

### GetPhase2EncryptionAlgorithmsOk

`func (o *Items7) GetPhase2EncryptionAlgorithmsOk() (*[]string, bool)`

GetPhase2EncryptionAlgorithmsOk returns a tuple with the Phase2EncryptionAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2EncryptionAlgorithms

`func (o *Items7) SetPhase2EncryptionAlgorithms(v []string)`

SetPhase2EncryptionAlgorithms sets Phase2EncryptionAlgorithms field to given value.

### HasPhase2EncryptionAlgorithms

`func (o *Items7) HasPhase2EncryptionAlgorithms() bool`

HasPhase2EncryptionAlgorithms returns a boolean if a field has been set.

### GetPhase1IntegrityAlgorithms

`func (o *Items7) GetPhase1IntegrityAlgorithms() []string`

GetPhase1IntegrityAlgorithms returns the Phase1IntegrityAlgorithms field if non-nil, zero value otherwise.

### GetPhase1IntegrityAlgorithmsOk

`func (o *Items7) GetPhase1IntegrityAlgorithmsOk() (*[]string, bool)`

GetPhase1IntegrityAlgorithmsOk returns a tuple with the Phase1IntegrityAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1IntegrityAlgorithms

`func (o *Items7) SetPhase1IntegrityAlgorithms(v []string)`

SetPhase1IntegrityAlgorithms sets Phase1IntegrityAlgorithms field to given value.

### HasPhase1IntegrityAlgorithms

`func (o *Items7) HasPhase1IntegrityAlgorithms() bool`

HasPhase1IntegrityAlgorithms returns a boolean if a field has been set.

### GetPhase2IntegrityAlgorithms

`func (o *Items7) GetPhase2IntegrityAlgorithms() []string`

GetPhase2IntegrityAlgorithms returns the Phase2IntegrityAlgorithms field if non-nil, zero value otherwise.

### GetPhase2IntegrityAlgorithmsOk

`func (o *Items7) GetPhase2IntegrityAlgorithmsOk() (*[]string, bool)`

GetPhase2IntegrityAlgorithmsOk returns a tuple with the Phase2IntegrityAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2IntegrityAlgorithms

`func (o *Items7) SetPhase2IntegrityAlgorithms(v []string)`

SetPhase2IntegrityAlgorithms sets Phase2IntegrityAlgorithms field to given value.

### HasPhase2IntegrityAlgorithms

`func (o *Items7) HasPhase2IntegrityAlgorithms() bool`

HasPhase2IntegrityAlgorithms returns a boolean if a field has been set.

### GetAcceptedRouteCount

`func (o *Items7) GetAcceptedRouteCount() float32`

GetAcceptedRouteCount returns the AcceptedRouteCount field if non-nil, zero value otherwise.

### GetAcceptedRouteCountOk

`func (o *Items7) GetAcceptedRouteCountOk() (*float32, bool)`

GetAcceptedRouteCountOk returns a tuple with the AcceptedRouteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedRouteCount

`func (o *Items7) SetAcceptedRouteCount(v float32)`

SetAcceptedRouteCount sets AcceptedRouteCount field to given value.

### HasAcceptedRouteCount

`func (o *Items7) HasAcceptedRouteCount() bool`

HasAcceptedRouteCount returns a boolean if a field has been set.

### GetLastStatusChange

`func (o *Items7) GetLastStatusChange() string`

GetLastStatusChange returns the LastStatusChange field if non-nil, zero value otherwise.

### GetLastStatusChangeOk

`func (o *Items7) GetLastStatusChangeOk() (*string, bool)`

GetLastStatusChangeOk returns a tuple with the LastStatusChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatusChange

`func (o *Items7) SetLastStatusChange(v string)`

SetLastStatusChange sets LastStatusChange field to given value.

### HasLastStatusChange

`func (o *Items7) HasLastStatusChange() bool`

HasLastStatusChange returns a boolean if a field has been set.

### GetStatus

`func (o *Items7) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Items7) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Items7) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Items7) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *Items7) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *Items7) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *Items7) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *Items7) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


