# KeyStoreJks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** |  | 
**KeystoreBase64** | **string** | base64 encoded file contents | 
**Alias** | **string** |  | 
**KeyPassphrase** | **string** |  | 
**StorePassphrase** | **string** |  | 
**KeystoreFileName** | Pointer to **string** |  | [optional] 

## Methods

### NewKeyStoreJks

`func NewKeyStoreJks(source string, keystoreBase64 string, alias string, keyPassphrase string, storePassphrase string, ) *KeyStoreJks`

NewKeyStoreJks instantiates a new KeyStoreJks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyStoreJksWithDefaults

`func NewKeyStoreJksWithDefaults() *KeyStoreJks`

NewKeyStoreJksWithDefaults instantiates a new KeyStoreJks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *KeyStoreJks) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *KeyStoreJks) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *KeyStoreJks) SetSource(v string)`

SetSource sets Source field to given value.


### GetKeystoreBase64

`func (o *KeyStoreJks) GetKeystoreBase64() string`

GetKeystoreBase64 returns the KeystoreBase64 field if non-nil, zero value otherwise.

### GetKeystoreBase64Ok

`func (o *KeyStoreJks) GetKeystoreBase64Ok() (*string, bool)`

GetKeystoreBase64Ok returns a tuple with the KeystoreBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoreBase64

`func (o *KeyStoreJks) SetKeystoreBase64(v string)`

SetKeystoreBase64 sets KeystoreBase64 field to given value.


### GetAlias

`func (o *KeyStoreJks) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *KeyStoreJks) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *KeyStoreJks) SetAlias(v string)`

SetAlias sets Alias field to given value.


### GetKeyPassphrase

`func (o *KeyStoreJks) GetKeyPassphrase() string`

GetKeyPassphrase returns the KeyPassphrase field if non-nil, zero value otherwise.

### GetKeyPassphraseOk

`func (o *KeyStoreJks) GetKeyPassphraseOk() (*string, bool)`

GetKeyPassphraseOk returns a tuple with the KeyPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPassphrase

`func (o *KeyStoreJks) SetKeyPassphrase(v string)`

SetKeyPassphrase sets KeyPassphrase field to given value.


### GetStorePassphrase

`func (o *KeyStoreJks) GetStorePassphrase() string`

GetStorePassphrase returns the StorePassphrase field if non-nil, zero value otherwise.

### GetStorePassphraseOk

`func (o *KeyStoreJks) GetStorePassphraseOk() (*string, bool)`

GetStorePassphraseOk returns a tuple with the StorePassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePassphrase

`func (o *KeyStoreJks) SetStorePassphrase(v string)`

SetStorePassphrase sets StorePassphrase field to given value.


### GetKeystoreFileName

`func (o *KeyStoreJks) GetKeystoreFileName() string`

GetKeystoreFileName returns the KeystoreFileName field if non-nil, zero value otherwise.

### GetKeystoreFileNameOk

`func (o *KeyStoreJks) GetKeystoreFileNameOk() (*string, bool)`

GetKeystoreFileNameOk returns a tuple with the KeystoreFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoreFileName

`func (o *KeyStoreJks) SetKeystoreFileName(v string)`

SetKeystoreFileName sets KeystoreFileName field to given value.

### HasKeystoreFileName

`func (o *KeyStoreJks) HasKeystoreFileName() bool`

HasKeystoreFileName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


