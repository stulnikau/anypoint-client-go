# PrivateSpaceTransitGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Spec** | [**TgwSpec**](TgwSpec.md) |  | 
**Status** | [**TgwStatus**](TgwStatus.md) |  | 

## Methods

### NewPrivateSpaceTransitGateway

`func NewPrivateSpaceTransitGateway(id string, name string, spec TgwSpec, status TgwStatus, ) *PrivateSpaceTransitGateway`

NewPrivateSpaceTransitGateway instantiates a new PrivateSpaceTransitGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateSpaceTransitGatewayWithDefaults

`func NewPrivateSpaceTransitGatewayWithDefaults() *PrivateSpaceTransitGateway`

NewPrivateSpaceTransitGatewayWithDefaults instantiates a new PrivateSpaceTransitGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrivateSpaceTransitGateway) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateSpaceTransitGateway) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateSpaceTransitGateway) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PrivateSpaceTransitGateway) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateSpaceTransitGateway) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateSpaceTransitGateway) SetName(v string)`

SetName sets Name field to given value.


### GetSpec

`func (o *PrivateSpaceTransitGateway) GetSpec() TgwSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *PrivateSpaceTransitGateway) GetSpecOk() (*TgwSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *PrivateSpaceTransitGateway) SetSpec(v TgwSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *PrivateSpaceTransitGateway) GetStatus() TgwStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrivateSpaceTransitGateway) GetStatusOk() (*TgwStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrivateSpaceTransitGateway) SetStatus(v TgwStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


