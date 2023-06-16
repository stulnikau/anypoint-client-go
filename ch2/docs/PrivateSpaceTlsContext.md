# PrivateSpaceTlsContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**TrustStore** | Pointer to [**TrustStore**](TrustStore.md) |  | [optional] 
**KeyStore** | [**KeyStoreResponse**](KeyStoreResponse.md) |  | 
**Ciphers** | [**Ciphers**](Ciphers.md) |  | 
**Type** | Pointer to **string** | \&quot;custom\&quot; TLS context is managed by users, whereas \&quot;default\&quot; is managed by Cloudhub 2.0  | [optional] 

## Methods

### NewPrivateSpaceTlsContext

`func NewPrivateSpaceTlsContext(id string, name string, keyStore KeyStoreResponse, ciphers Ciphers, ) *PrivateSpaceTlsContext`

NewPrivateSpaceTlsContext instantiates a new PrivateSpaceTlsContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateSpaceTlsContextWithDefaults

`func NewPrivateSpaceTlsContextWithDefaults() *PrivateSpaceTlsContext`

NewPrivateSpaceTlsContextWithDefaults instantiates a new PrivateSpaceTlsContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrivateSpaceTlsContext) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateSpaceTlsContext) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateSpaceTlsContext) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PrivateSpaceTlsContext) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateSpaceTlsContext) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateSpaceTlsContext) SetName(v string)`

SetName sets Name field to given value.


### GetTrustStore

`func (o *PrivateSpaceTlsContext) GetTrustStore() TrustStore`

GetTrustStore returns the TrustStore field if non-nil, zero value otherwise.

### GetTrustStoreOk

`func (o *PrivateSpaceTlsContext) GetTrustStoreOk() (*TrustStore, bool)`

GetTrustStoreOk returns a tuple with the TrustStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStore

`func (o *PrivateSpaceTlsContext) SetTrustStore(v TrustStore)`

SetTrustStore sets TrustStore field to given value.

### HasTrustStore

`func (o *PrivateSpaceTlsContext) HasTrustStore() bool`

HasTrustStore returns a boolean if a field has been set.

### GetKeyStore

`func (o *PrivateSpaceTlsContext) GetKeyStore() KeyStoreResponse`

GetKeyStore returns the KeyStore field if non-nil, zero value otherwise.

### GetKeyStoreOk

`func (o *PrivateSpaceTlsContext) GetKeyStoreOk() (*KeyStoreResponse, bool)`

GetKeyStoreOk returns a tuple with the KeyStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStore

`func (o *PrivateSpaceTlsContext) SetKeyStore(v KeyStoreResponse)`

SetKeyStore sets KeyStore field to given value.


### GetCiphers

`func (o *PrivateSpaceTlsContext) GetCiphers() Ciphers`

GetCiphers returns the Ciphers field if non-nil, zero value otherwise.

### GetCiphersOk

`func (o *PrivateSpaceTlsContext) GetCiphersOk() (*Ciphers, bool)`

GetCiphersOk returns a tuple with the Ciphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphers

`func (o *PrivateSpaceTlsContext) SetCiphers(v Ciphers)`

SetCiphers sets Ciphers field to given value.


### GetType

`func (o *PrivateSpaceTlsContext) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PrivateSpaceTlsContext) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PrivateSpaceTlsContext) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PrivateSpaceTlsContext) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


