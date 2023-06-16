# PrivateSpaceTlsContextPemPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**TlsConfig** | [**TlsConfigPem**](TlsConfigPem.md) |  | 
**Ciphers** | [**Ciphers**](Ciphers.md) |  | 

## Methods

### NewPrivateSpaceTlsContextPemPost

`func NewPrivateSpaceTlsContextPemPost(name string, tlsConfig TlsConfigPem, ciphers Ciphers, ) *PrivateSpaceTlsContextPemPost`

NewPrivateSpaceTlsContextPemPost instantiates a new PrivateSpaceTlsContextPemPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateSpaceTlsContextPemPostWithDefaults

`func NewPrivateSpaceTlsContextPemPostWithDefaults() *PrivateSpaceTlsContextPemPost`

NewPrivateSpaceTlsContextPemPostWithDefaults instantiates a new PrivateSpaceTlsContextPemPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PrivateSpaceTlsContextPemPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateSpaceTlsContextPemPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateSpaceTlsContextPemPost) SetName(v string)`

SetName sets Name field to given value.


### GetTlsConfig

`func (o *PrivateSpaceTlsContextPemPost) GetTlsConfig() TlsConfigPem`

GetTlsConfig returns the TlsConfig field if non-nil, zero value otherwise.

### GetTlsConfigOk

`func (o *PrivateSpaceTlsContextPemPost) GetTlsConfigOk() (*TlsConfigPem, bool)`

GetTlsConfigOk returns a tuple with the TlsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsConfig

`func (o *PrivateSpaceTlsContextPemPost) SetTlsConfig(v TlsConfigPem)`

SetTlsConfig sets TlsConfig field to given value.


### GetCiphers

`func (o *PrivateSpaceTlsContextPemPost) GetCiphers() Ciphers`

GetCiphers returns the Ciphers field if non-nil, zero value otherwise.

### GetCiphersOk

`func (o *PrivateSpaceTlsContextPemPost) GetCiphersOk() (*Ciphers, bool)`

GetCiphersOk returns a tuple with the Ciphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphers

`func (o *PrivateSpaceTlsContextPemPost) SetCiphers(v Ciphers)`

SetCiphers sets Ciphers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


