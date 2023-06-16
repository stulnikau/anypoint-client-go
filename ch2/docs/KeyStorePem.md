# KeyStorePem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** |  | 
**Certificate** | **string** |  | 
**Key** | **string** |  | 
**Capath** | Pointer to **string** |  | [optional] 
**CertificateFileName** | Pointer to **string** |  | [optional] 
**KeyFileName** | Pointer to **string** |  | [optional] 
**CapathFileName** | Pointer to **string** |  | [optional] 
**KeyPassphrase** | **string** |  | 

## Methods

### NewKeyStorePem

`func NewKeyStorePem(source string, certificate string, key string, keyPassphrase string, ) *KeyStorePem`

NewKeyStorePem instantiates a new KeyStorePem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyStorePemWithDefaults

`func NewKeyStorePemWithDefaults() *KeyStorePem`

NewKeyStorePemWithDefaults instantiates a new KeyStorePem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *KeyStorePem) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *KeyStorePem) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *KeyStorePem) SetSource(v string)`

SetSource sets Source field to given value.


### GetCertificate

`func (o *KeyStorePem) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *KeyStorePem) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *KeyStorePem) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetKey

`func (o *KeyStorePem) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *KeyStorePem) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *KeyStorePem) SetKey(v string)`

SetKey sets Key field to given value.


### GetCapath

`func (o *KeyStorePem) GetCapath() string`

GetCapath returns the Capath field if non-nil, zero value otherwise.

### GetCapathOk

`func (o *KeyStorePem) GetCapathOk() (*string, bool)`

GetCapathOk returns a tuple with the Capath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapath

`func (o *KeyStorePem) SetCapath(v string)`

SetCapath sets Capath field to given value.

### HasCapath

`func (o *KeyStorePem) HasCapath() bool`

HasCapath returns a boolean if a field has been set.

### GetCertificateFileName

`func (o *KeyStorePem) GetCertificateFileName() string`

GetCertificateFileName returns the CertificateFileName field if non-nil, zero value otherwise.

### GetCertificateFileNameOk

`func (o *KeyStorePem) GetCertificateFileNameOk() (*string, bool)`

GetCertificateFileNameOk returns a tuple with the CertificateFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateFileName

`func (o *KeyStorePem) SetCertificateFileName(v string)`

SetCertificateFileName sets CertificateFileName field to given value.

### HasCertificateFileName

`func (o *KeyStorePem) HasCertificateFileName() bool`

HasCertificateFileName returns a boolean if a field has been set.

### GetKeyFileName

`func (o *KeyStorePem) GetKeyFileName() string`

GetKeyFileName returns the KeyFileName field if non-nil, zero value otherwise.

### GetKeyFileNameOk

`func (o *KeyStorePem) GetKeyFileNameOk() (*string, bool)`

GetKeyFileNameOk returns a tuple with the KeyFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyFileName

`func (o *KeyStorePem) SetKeyFileName(v string)`

SetKeyFileName sets KeyFileName field to given value.

### HasKeyFileName

`func (o *KeyStorePem) HasKeyFileName() bool`

HasKeyFileName returns a boolean if a field has been set.

### GetCapathFileName

`func (o *KeyStorePem) GetCapathFileName() string`

GetCapathFileName returns the CapathFileName field if non-nil, zero value otherwise.

### GetCapathFileNameOk

`func (o *KeyStorePem) GetCapathFileNameOk() (*string, bool)`

GetCapathFileNameOk returns a tuple with the CapathFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapathFileName

`func (o *KeyStorePem) SetCapathFileName(v string)`

SetCapathFileName sets CapathFileName field to given value.

### HasCapathFileName

`func (o *KeyStorePem) HasCapathFileName() bool`

HasCapathFileName returns a boolean if a field has been set.

### GetKeyPassphrase

`func (o *KeyStorePem) GetKeyPassphrase() string`

GetKeyPassphrase returns the KeyPassphrase field if non-nil, zero value otherwise.

### GetKeyPassphraseOk

`func (o *KeyStorePem) GetKeyPassphraseOk() (*string, bool)`

GetKeyPassphraseOk returns a tuple with the KeyPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPassphrase

`func (o *KeyStorePem) SetKeyPassphrase(v string)`

SetKeyPassphrase sets KeyPassphrase field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


