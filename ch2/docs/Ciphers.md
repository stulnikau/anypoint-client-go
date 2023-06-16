# Ciphers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aes128GcmSha256** | **bool** |  | [default to false]
**Aes128Sha256** | **bool** |  | [default to false]
**Aes256GcmSha384** | **bool** |  | [default to false]
**Aes256Sha256** | **bool** |  | [default to false]
**DheRsaAes128Sha256** | **bool** |  | [default to false]
**DheRsaAes256GcmSha384** | **bool** |  | [default to false]
**DheRsaAes256Sha256** | **bool** |  | [default to false]
**EcdheEcdsaAes128GcmSha256** | **bool** |  | [default to true]
**EcdheEcdsaAes128Sha1** | **bool** |  | [default to false]
**EcdheEcdsaAes256GcmSha384** | **bool** |  | [default to true]
**EcdheEcdsaAes256Sha1** | **bool** |  | [default to false]
**EcdheRsaAes128GcmSha256** | **bool** |  | [default to true]
**EcdheRsaAes128Sha1** | **bool** |  | [default to false]
**EcdheRsaAes256GcmSha384** | **bool** |  | [default to true]
**EcdheRsaAes256Sha1** | **bool** |  | [default to false]
**EcdheEcdsaChacha20Poly1305** | **bool** |  | [default to false]
**EcdheRsaChacha20Poly1305** | **bool** |  | [default to false]
**DheRsaChacha20Poly1305** | **bool** |  | [default to false]
**TlsAes256GcmSha384** | **bool** |  | [default to true]
**TlsChacha20Poly1305Sha256** | **bool** |  | [default to true]
**TlsAes128GcmSha256** | **bool** |  | [default to true]

## Methods

### NewCiphers

`func NewCiphers(aes128GcmSha256 bool, aes128Sha256 bool, aes256GcmSha384 bool, aes256Sha256 bool, dheRsaAes128Sha256 bool, dheRsaAes256GcmSha384 bool, dheRsaAes256Sha256 bool, ecdheEcdsaAes128GcmSha256 bool, ecdheEcdsaAes128Sha1 bool, ecdheEcdsaAes256GcmSha384 bool, ecdheEcdsaAes256Sha1 bool, ecdheRsaAes128GcmSha256 bool, ecdheRsaAes128Sha1 bool, ecdheRsaAes256GcmSha384 bool, ecdheRsaAes256Sha1 bool, ecdheEcdsaChacha20Poly1305 bool, ecdheRsaChacha20Poly1305 bool, dheRsaChacha20Poly1305 bool, tlsAes256GcmSha384 bool, tlsChacha20Poly1305Sha256 bool, tlsAes128GcmSha256 bool, ) *Ciphers`

NewCiphers instantiates a new Ciphers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCiphersWithDefaults

`func NewCiphersWithDefaults() *Ciphers`

NewCiphersWithDefaults instantiates a new Ciphers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAes128GcmSha256

`func (o *Ciphers) GetAes128GcmSha256() bool`

GetAes128GcmSha256 returns the Aes128GcmSha256 field if non-nil, zero value otherwise.

### GetAes128GcmSha256Ok

`func (o *Ciphers) GetAes128GcmSha256Ok() (*bool, bool)`

GetAes128GcmSha256Ok returns a tuple with the Aes128GcmSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAes128GcmSha256

`func (o *Ciphers) SetAes128GcmSha256(v bool)`

SetAes128GcmSha256 sets Aes128GcmSha256 field to given value.


### GetAes128Sha256

`func (o *Ciphers) GetAes128Sha256() bool`

GetAes128Sha256 returns the Aes128Sha256 field if non-nil, zero value otherwise.

### GetAes128Sha256Ok

`func (o *Ciphers) GetAes128Sha256Ok() (*bool, bool)`

GetAes128Sha256Ok returns a tuple with the Aes128Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAes128Sha256

`func (o *Ciphers) SetAes128Sha256(v bool)`

SetAes128Sha256 sets Aes128Sha256 field to given value.


### GetAes256GcmSha384

`func (o *Ciphers) GetAes256GcmSha384() bool`

GetAes256GcmSha384 returns the Aes256GcmSha384 field if non-nil, zero value otherwise.

### GetAes256GcmSha384Ok

`func (o *Ciphers) GetAes256GcmSha384Ok() (*bool, bool)`

GetAes256GcmSha384Ok returns a tuple with the Aes256GcmSha384 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAes256GcmSha384

`func (o *Ciphers) SetAes256GcmSha384(v bool)`

SetAes256GcmSha384 sets Aes256GcmSha384 field to given value.


### GetAes256Sha256

`func (o *Ciphers) GetAes256Sha256() bool`

GetAes256Sha256 returns the Aes256Sha256 field if non-nil, zero value otherwise.

### GetAes256Sha256Ok

`func (o *Ciphers) GetAes256Sha256Ok() (*bool, bool)`

GetAes256Sha256Ok returns a tuple with the Aes256Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAes256Sha256

`func (o *Ciphers) SetAes256Sha256(v bool)`

SetAes256Sha256 sets Aes256Sha256 field to given value.


### GetDheRsaAes128Sha256

`func (o *Ciphers) GetDheRsaAes128Sha256() bool`

GetDheRsaAes128Sha256 returns the DheRsaAes128Sha256 field if non-nil, zero value otherwise.

### GetDheRsaAes128Sha256Ok

`func (o *Ciphers) GetDheRsaAes128Sha256Ok() (*bool, bool)`

GetDheRsaAes128Sha256Ok returns a tuple with the DheRsaAes128Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDheRsaAes128Sha256

`func (o *Ciphers) SetDheRsaAes128Sha256(v bool)`

SetDheRsaAes128Sha256 sets DheRsaAes128Sha256 field to given value.


### GetDheRsaAes256GcmSha384

`func (o *Ciphers) GetDheRsaAes256GcmSha384() bool`

GetDheRsaAes256GcmSha384 returns the DheRsaAes256GcmSha384 field if non-nil, zero value otherwise.

### GetDheRsaAes256GcmSha384Ok

`func (o *Ciphers) GetDheRsaAes256GcmSha384Ok() (*bool, bool)`

GetDheRsaAes256GcmSha384Ok returns a tuple with the DheRsaAes256GcmSha384 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDheRsaAes256GcmSha384

`func (o *Ciphers) SetDheRsaAes256GcmSha384(v bool)`

SetDheRsaAes256GcmSha384 sets DheRsaAes256GcmSha384 field to given value.


### GetDheRsaAes256Sha256

`func (o *Ciphers) GetDheRsaAes256Sha256() bool`

GetDheRsaAes256Sha256 returns the DheRsaAes256Sha256 field if non-nil, zero value otherwise.

### GetDheRsaAes256Sha256Ok

`func (o *Ciphers) GetDheRsaAes256Sha256Ok() (*bool, bool)`

GetDheRsaAes256Sha256Ok returns a tuple with the DheRsaAes256Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDheRsaAes256Sha256

`func (o *Ciphers) SetDheRsaAes256Sha256(v bool)`

SetDheRsaAes256Sha256 sets DheRsaAes256Sha256 field to given value.


### GetEcdheEcdsaAes128GcmSha256

`func (o *Ciphers) GetEcdheEcdsaAes128GcmSha256() bool`

GetEcdheEcdsaAes128GcmSha256 returns the EcdheEcdsaAes128GcmSha256 field if non-nil, zero value otherwise.

### GetEcdheEcdsaAes128GcmSha256Ok

`func (o *Ciphers) GetEcdheEcdsaAes128GcmSha256Ok() (*bool, bool)`

GetEcdheEcdsaAes128GcmSha256Ok returns a tuple with the EcdheEcdsaAes128GcmSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdheEcdsaAes128GcmSha256

`func (o *Ciphers) SetEcdheEcdsaAes128GcmSha256(v bool)`

SetEcdheEcdsaAes128GcmSha256 sets EcdheEcdsaAes128GcmSha256 field to given value.


### GetEcdheEcdsaAes128Sha1

`func (o *Ciphers) GetEcdheEcdsaAes128Sha1() bool`

GetEcdheEcdsaAes128Sha1 returns the EcdheEcdsaAes128Sha1 field if non-nil, zero value otherwise.

### GetEcdheEcdsaAes128Sha1Ok

`func (o *Ciphers) GetEcdheEcdsaAes128Sha1Ok() (*bool, bool)`

GetEcdheEcdsaAes128Sha1Ok returns a tuple with the EcdheEcdsaAes128Sha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdheEcdsaAes128Sha1

`func (o *Ciphers) SetEcdheEcdsaAes128Sha1(v bool)`

SetEcdheEcdsaAes128Sha1 sets EcdheEcdsaAes128Sha1 field to given value.


### GetEcdheEcdsaAes256GcmSha384

`func (o *Ciphers) GetEcdheEcdsaAes256GcmSha384() bool`

GetEcdheEcdsaAes256GcmSha384 returns the EcdheEcdsaAes256GcmSha384 field if non-nil, zero value otherwise.

### GetEcdheEcdsaAes256GcmSha384Ok

`func (o *Ciphers) GetEcdheEcdsaAes256GcmSha384Ok() (*bool, bool)`

GetEcdheEcdsaAes256GcmSha384Ok returns a tuple with the EcdheEcdsaAes256GcmSha384 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdheEcdsaAes256GcmSha384

`func (o *Ciphers) SetEcdheEcdsaAes256GcmSha384(v bool)`

SetEcdheEcdsaAes256GcmSha384 sets EcdheEcdsaAes256GcmSha384 field to given value.


### GetEcdheEcdsaAes256Sha1

`func (o *Ciphers) GetEcdheEcdsaAes256Sha1() bool`

GetEcdheEcdsaAes256Sha1 returns the EcdheEcdsaAes256Sha1 field if non-nil, zero value otherwise.

### GetEcdheEcdsaAes256Sha1Ok

`func (o *Ciphers) GetEcdheEcdsaAes256Sha1Ok() (*bool, bool)`

GetEcdheEcdsaAes256Sha1Ok returns a tuple with the EcdheEcdsaAes256Sha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdheEcdsaAes256Sha1

`func (o *Ciphers) SetEcdheEcdsaAes256Sha1(v bool)`

SetEcdheEcdsaAes256Sha1 sets EcdheEcdsaAes256Sha1 field to given value.


### GetEcdheRsaAes128GcmSha256

`func (o *Ciphers) GetEcdheRsaAes128GcmSha256() bool`

GetEcdheRsaAes128GcmSha256 returns the EcdheRsaAes128GcmSha256 field if non-nil, zero value otherwise.

### GetEcdheRsaAes128GcmSha256Ok

`func (o *Ciphers) GetEcdheRsaAes128GcmSha256Ok() (*bool, bool)`

GetEcdheRsaAes128GcmSha256Ok returns a tuple with the EcdheRsaAes128GcmSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdheRsaAes128GcmSha256

`func (o *Ciphers) SetEcdheRsaAes128GcmSha256(v bool)`

SetEcdheRsaAes128GcmSha256 sets EcdheRsaAes128GcmSha256 field to given value.


### GetEcdheRsaAes128Sha1

`func (o *Ciphers) GetEcdheRsaAes128Sha1() bool`

GetEcdheRsaAes128Sha1 returns the EcdheRsaAes128Sha1 field if non-nil, zero value otherwise.

### GetEcdheRsaAes128Sha1Ok

`func (o *Ciphers) GetEcdheRsaAes128Sha1Ok() (*bool, bool)`

GetEcdheRsaAes128Sha1Ok returns a tuple with the EcdheRsaAes128Sha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdheRsaAes128Sha1

`func (o *Ciphers) SetEcdheRsaAes128Sha1(v bool)`

SetEcdheRsaAes128Sha1 sets EcdheRsaAes128Sha1 field to given value.


### GetEcdheRsaAes256GcmSha384

`func (o *Ciphers) GetEcdheRsaAes256GcmSha384() bool`

GetEcdheRsaAes256GcmSha384 returns the EcdheRsaAes256GcmSha384 field if non-nil, zero value otherwise.

### GetEcdheRsaAes256GcmSha384Ok

`func (o *Ciphers) GetEcdheRsaAes256GcmSha384Ok() (*bool, bool)`

GetEcdheRsaAes256GcmSha384Ok returns a tuple with the EcdheRsaAes256GcmSha384 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdheRsaAes256GcmSha384

`func (o *Ciphers) SetEcdheRsaAes256GcmSha384(v bool)`

SetEcdheRsaAes256GcmSha384 sets EcdheRsaAes256GcmSha384 field to given value.


### GetEcdheRsaAes256Sha1

`func (o *Ciphers) GetEcdheRsaAes256Sha1() bool`

GetEcdheRsaAes256Sha1 returns the EcdheRsaAes256Sha1 field if non-nil, zero value otherwise.

### GetEcdheRsaAes256Sha1Ok

`func (o *Ciphers) GetEcdheRsaAes256Sha1Ok() (*bool, bool)`

GetEcdheRsaAes256Sha1Ok returns a tuple with the EcdheRsaAes256Sha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdheRsaAes256Sha1

`func (o *Ciphers) SetEcdheRsaAes256Sha1(v bool)`

SetEcdheRsaAes256Sha1 sets EcdheRsaAes256Sha1 field to given value.


### GetEcdheEcdsaChacha20Poly1305

`func (o *Ciphers) GetEcdheEcdsaChacha20Poly1305() bool`

GetEcdheEcdsaChacha20Poly1305 returns the EcdheEcdsaChacha20Poly1305 field if non-nil, zero value otherwise.

### GetEcdheEcdsaChacha20Poly1305Ok

`func (o *Ciphers) GetEcdheEcdsaChacha20Poly1305Ok() (*bool, bool)`

GetEcdheEcdsaChacha20Poly1305Ok returns a tuple with the EcdheEcdsaChacha20Poly1305 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdheEcdsaChacha20Poly1305

`func (o *Ciphers) SetEcdheEcdsaChacha20Poly1305(v bool)`

SetEcdheEcdsaChacha20Poly1305 sets EcdheEcdsaChacha20Poly1305 field to given value.


### GetEcdheRsaChacha20Poly1305

`func (o *Ciphers) GetEcdheRsaChacha20Poly1305() bool`

GetEcdheRsaChacha20Poly1305 returns the EcdheRsaChacha20Poly1305 field if non-nil, zero value otherwise.

### GetEcdheRsaChacha20Poly1305Ok

`func (o *Ciphers) GetEcdheRsaChacha20Poly1305Ok() (*bool, bool)`

GetEcdheRsaChacha20Poly1305Ok returns a tuple with the EcdheRsaChacha20Poly1305 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdheRsaChacha20Poly1305

`func (o *Ciphers) SetEcdheRsaChacha20Poly1305(v bool)`

SetEcdheRsaChacha20Poly1305 sets EcdheRsaChacha20Poly1305 field to given value.


### GetDheRsaChacha20Poly1305

`func (o *Ciphers) GetDheRsaChacha20Poly1305() bool`

GetDheRsaChacha20Poly1305 returns the DheRsaChacha20Poly1305 field if non-nil, zero value otherwise.

### GetDheRsaChacha20Poly1305Ok

`func (o *Ciphers) GetDheRsaChacha20Poly1305Ok() (*bool, bool)`

GetDheRsaChacha20Poly1305Ok returns a tuple with the DheRsaChacha20Poly1305 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDheRsaChacha20Poly1305

`func (o *Ciphers) SetDheRsaChacha20Poly1305(v bool)`

SetDheRsaChacha20Poly1305 sets DheRsaChacha20Poly1305 field to given value.


### GetTlsAes256GcmSha384

`func (o *Ciphers) GetTlsAes256GcmSha384() bool`

GetTlsAes256GcmSha384 returns the TlsAes256GcmSha384 field if non-nil, zero value otherwise.

### GetTlsAes256GcmSha384Ok

`func (o *Ciphers) GetTlsAes256GcmSha384Ok() (*bool, bool)`

GetTlsAes256GcmSha384Ok returns a tuple with the TlsAes256GcmSha384 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsAes256GcmSha384

`func (o *Ciphers) SetTlsAes256GcmSha384(v bool)`

SetTlsAes256GcmSha384 sets TlsAes256GcmSha384 field to given value.


### GetTlsChacha20Poly1305Sha256

`func (o *Ciphers) GetTlsChacha20Poly1305Sha256() bool`

GetTlsChacha20Poly1305Sha256 returns the TlsChacha20Poly1305Sha256 field if non-nil, zero value otherwise.

### GetTlsChacha20Poly1305Sha256Ok

`func (o *Ciphers) GetTlsChacha20Poly1305Sha256Ok() (*bool, bool)`

GetTlsChacha20Poly1305Sha256Ok returns a tuple with the TlsChacha20Poly1305Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsChacha20Poly1305Sha256

`func (o *Ciphers) SetTlsChacha20Poly1305Sha256(v bool)`

SetTlsChacha20Poly1305Sha256 sets TlsChacha20Poly1305Sha256 field to given value.


### GetTlsAes128GcmSha256

`func (o *Ciphers) GetTlsAes128GcmSha256() bool`

GetTlsAes128GcmSha256 returns the TlsAes128GcmSha256 field if non-nil, zero value otherwise.

### GetTlsAes128GcmSha256Ok

`func (o *Ciphers) GetTlsAes128GcmSha256Ok() (*bool, bool)`

GetTlsAes128GcmSha256Ok returns a tuple with the TlsAes128GcmSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsAes128GcmSha256

`func (o *Ciphers) SetTlsAes128GcmSha256(v bool)`

SetTlsAes128GcmSha256 sets TlsAes128GcmSha256 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


