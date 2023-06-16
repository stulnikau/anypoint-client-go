# Items10

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issuer** | [**Issuer**](Issuer.md) |  | 
**Subject** | [**Subject**](Subject.md) |  | 
**SubjectAlternativeName** | Pointer to **[]string** | Collection of subject alternative names from the SubjectAltName x509 extension | [optional] 
**Version** | **string** |  | 
**SerialNumber** | **string** | Serial number assigned by the CA to this certificate, in hex format | 
**SignatureAlgorithm** | **string** | Name of the signature algorithm | 
**PublicKeyAlgorithm** | **string** | The standard algorithm name for the public key of this certificate | 
**BasicConstraints** | Pointer to [**Items10BasicConstraints**](Items10BasicConstraints.md) |  | [optional] 
**Validity** | [**Validity**](Validity.md) |  | 
**KeyUsage** | Pointer to **[]string** | A list of values defining the purpose of the public key i.e. the key usage extensions from this certificate | [optional] 
**ExtendedKeyUsage** | Pointer to **[]string** | A list of values providing details about the extended key usage extensions from this certificate. | [optional] 
**CertificateType** | **string** | The type of this certificate | 

## Methods

### NewItems10

`func NewItems10(issuer Issuer, subject Subject, version string, serialNumber string, signatureAlgorithm string, publicKeyAlgorithm string, validity Validity, certificateType string, ) *Items10`

NewItems10 instantiates a new Items10 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItems10WithDefaults

`func NewItems10WithDefaults() *Items10`

NewItems10WithDefaults instantiates a new Items10 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuer

`func (o *Items10) GetIssuer() Issuer`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *Items10) GetIssuerOk() (*Issuer, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *Items10) SetIssuer(v Issuer)`

SetIssuer sets Issuer field to given value.


### GetSubject

`func (o *Items10) GetSubject() Subject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Items10) GetSubjectOk() (*Subject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Items10) SetSubject(v Subject)`

SetSubject sets Subject field to given value.


### GetSubjectAlternativeName

`func (o *Items10) GetSubjectAlternativeName() []string`

GetSubjectAlternativeName returns the SubjectAlternativeName field if non-nil, zero value otherwise.

### GetSubjectAlternativeNameOk

`func (o *Items10) GetSubjectAlternativeNameOk() (*[]string, bool)`

GetSubjectAlternativeNameOk returns a tuple with the SubjectAlternativeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAlternativeName

`func (o *Items10) SetSubjectAlternativeName(v []string)`

SetSubjectAlternativeName sets SubjectAlternativeName field to given value.

### HasSubjectAlternativeName

`func (o *Items10) HasSubjectAlternativeName() bool`

HasSubjectAlternativeName returns a boolean if a field has been set.

### GetVersion

`func (o *Items10) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Items10) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Items10) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetSerialNumber

`func (o *Items10) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *Items10) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *Items10) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetSignatureAlgorithm

`func (o *Items10) GetSignatureAlgorithm() string`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *Items10) GetSignatureAlgorithmOk() (*string, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *Items10) SetSignatureAlgorithm(v string)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.


### GetPublicKeyAlgorithm

`func (o *Items10) GetPublicKeyAlgorithm() string`

GetPublicKeyAlgorithm returns the PublicKeyAlgorithm field if non-nil, zero value otherwise.

### GetPublicKeyAlgorithmOk

`func (o *Items10) GetPublicKeyAlgorithmOk() (*string, bool)`

GetPublicKeyAlgorithmOk returns a tuple with the PublicKeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyAlgorithm

`func (o *Items10) SetPublicKeyAlgorithm(v string)`

SetPublicKeyAlgorithm sets PublicKeyAlgorithm field to given value.


### GetBasicConstraints

`func (o *Items10) GetBasicConstraints() Items10BasicConstraints`

GetBasicConstraints returns the BasicConstraints field if non-nil, zero value otherwise.

### GetBasicConstraintsOk

`func (o *Items10) GetBasicConstraintsOk() (*Items10BasicConstraints, bool)`

GetBasicConstraintsOk returns a tuple with the BasicConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicConstraints

`func (o *Items10) SetBasicConstraints(v Items10BasicConstraints)`

SetBasicConstraints sets BasicConstraints field to given value.

### HasBasicConstraints

`func (o *Items10) HasBasicConstraints() bool`

HasBasicConstraints returns a boolean if a field has been set.

### GetValidity

`func (o *Items10) GetValidity() Validity`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *Items10) GetValidityOk() (*Validity, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *Items10) SetValidity(v Validity)`

SetValidity sets Validity field to given value.


### GetKeyUsage

`func (o *Items10) GetKeyUsage() []string`

GetKeyUsage returns the KeyUsage field if non-nil, zero value otherwise.

### GetKeyUsageOk

`func (o *Items10) GetKeyUsageOk() (*[]string, bool)`

GetKeyUsageOk returns a tuple with the KeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsage

`func (o *Items10) SetKeyUsage(v []string)`

SetKeyUsage sets KeyUsage field to given value.

### HasKeyUsage

`func (o *Items10) HasKeyUsage() bool`

HasKeyUsage returns a boolean if a field has been set.

### GetExtendedKeyUsage

`func (o *Items10) GetExtendedKeyUsage() []string`

GetExtendedKeyUsage returns the ExtendedKeyUsage field if non-nil, zero value otherwise.

### GetExtendedKeyUsageOk

`func (o *Items10) GetExtendedKeyUsageOk() (*[]string, bool)`

GetExtendedKeyUsageOk returns a tuple with the ExtendedKeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedKeyUsage

`func (o *Items10) SetExtendedKeyUsage(v []string)`

SetExtendedKeyUsage sets ExtendedKeyUsage field to given value.

### HasExtendedKeyUsage

`func (o *Items10) HasExtendedKeyUsage() bool`

HasExtendedKeyUsage returns a boolean if a field has been set.

### GetCertificateType

`func (o *Items10) GetCertificateType() string`

GetCertificateType returns the CertificateType field if non-nil, zero value otherwise.

### GetCertificateTypeOk

`func (o *Items10) GetCertificateTypeOk() (*string, bool)`

GetCertificateTypeOk returns a tuple with the CertificateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateType

`func (o *Items10) SetCertificateType(v string)`

SetCertificateType sets CertificateType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


