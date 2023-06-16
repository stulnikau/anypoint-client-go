# KeyStoreResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Cn** | **string** |  | 
**San** | **[]string** |  | 
**Alias** | Pointer to **string** |  | [optional] 
**FileName** | **string** | name of the key store file | 
**KeyFileName** | Pointer to **string** |  | [optional] 
**CapathFileName** | Pointer to **string** |  | [optional] 
**ExpirationDate** | Pointer to **string** |  | [optional] 

## Methods

### NewKeyStoreResponse

`func NewKeyStoreResponse(type_ string, cn string, san []string, fileName string, ) *KeyStoreResponse`

NewKeyStoreResponse instantiates a new KeyStoreResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyStoreResponseWithDefaults

`func NewKeyStoreResponseWithDefaults() *KeyStoreResponse`

NewKeyStoreResponseWithDefaults instantiates a new KeyStoreResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *KeyStoreResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KeyStoreResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KeyStoreResponse) SetType(v string)`

SetType sets Type field to given value.


### GetCn

`func (o *KeyStoreResponse) GetCn() string`

GetCn returns the Cn field if non-nil, zero value otherwise.

### GetCnOk

`func (o *KeyStoreResponse) GetCnOk() (*string, bool)`

GetCnOk returns a tuple with the Cn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCn

`func (o *KeyStoreResponse) SetCn(v string)`

SetCn sets Cn field to given value.


### GetSan

`func (o *KeyStoreResponse) GetSan() []string`

GetSan returns the San field if non-nil, zero value otherwise.

### GetSanOk

`func (o *KeyStoreResponse) GetSanOk() (*[]string, bool)`

GetSanOk returns a tuple with the San field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSan

`func (o *KeyStoreResponse) SetSan(v []string)`

SetSan sets San field to given value.


### GetAlias

`func (o *KeyStoreResponse) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *KeyStoreResponse) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *KeyStoreResponse) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *KeyStoreResponse) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetFileName

`func (o *KeyStoreResponse) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *KeyStoreResponse) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *KeyStoreResponse) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetKeyFileName

`func (o *KeyStoreResponse) GetKeyFileName() string`

GetKeyFileName returns the KeyFileName field if non-nil, zero value otherwise.

### GetKeyFileNameOk

`func (o *KeyStoreResponse) GetKeyFileNameOk() (*string, bool)`

GetKeyFileNameOk returns a tuple with the KeyFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyFileName

`func (o *KeyStoreResponse) SetKeyFileName(v string)`

SetKeyFileName sets KeyFileName field to given value.

### HasKeyFileName

`func (o *KeyStoreResponse) HasKeyFileName() bool`

HasKeyFileName returns a boolean if a field has been set.

### GetCapathFileName

`func (o *KeyStoreResponse) GetCapathFileName() string`

GetCapathFileName returns the CapathFileName field if non-nil, zero value otherwise.

### GetCapathFileNameOk

`func (o *KeyStoreResponse) GetCapathFileNameOk() (*string, bool)`

GetCapathFileNameOk returns a tuple with the CapathFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapathFileName

`func (o *KeyStoreResponse) SetCapathFileName(v string)`

SetCapathFileName sets CapathFileName field to given value.

### HasCapathFileName

`func (o *KeyStoreResponse) HasCapathFileName() bool`

HasCapathFileName returns a boolean if a field has been set.

### GetExpirationDate

`func (o *KeyStoreResponse) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *KeyStoreResponse) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *KeyStoreResponse) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *KeyStoreResponse) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


