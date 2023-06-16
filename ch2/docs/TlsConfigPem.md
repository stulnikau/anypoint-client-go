# TlsConfigPem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyStore** | [**KeyStorePem**](KeyStorePem.md) |  | 
**TrustStore** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewTlsConfigPem

`func NewTlsConfigPem(keyStore KeyStorePem, ) *TlsConfigPem`

NewTlsConfigPem instantiates a new TlsConfigPem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsConfigPemWithDefaults

`func NewTlsConfigPemWithDefaults() *TlsConfigPem`

NewTlsConfigPemWithDefaults instantiates a new TlsConfigPem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyStore

`func (o *TlsConfigPem) GetKeyStore() KeyStorePem`

GetKeyStore returns the KeyStore field if non-nil, zero value otherwise.

### GetKeyStoreOk

`func (o *TlsConfigPem) GetKeyStoreOk() (*KeyStorePem, bool)`

GetKeyStoreOk returns a tuple with the KeyStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStore

`func (o *TlsConfigPem) SetKeyStore(v KeyStorePem)`

SetKeyStore sets KeyStore field to given value.


### GetTrustStore

`func (o *TlsConfigPem) GetTrustStore() map[string]interface{}`

GetTrustStore returns the TrustStore field if non-nil, zero value otherwise.

### GetTrustStoreOk

`func (o *TlsConfigPem) GetTrustStoreOk() (*map[string]interface{}, bool)`

GetTrustStoreOk returns a tuple with the TrustStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStore

`func (o *TlsConfigPem) SetTrustStore(v map[string]interface{})`

SetTrustStore sets TrustStore field to given value.

### HasTrustStore

`func (o *TlsConfigPem) HasTrustStore() bool`

HasTrustStore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


