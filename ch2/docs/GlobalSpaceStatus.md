# GlobalSpaceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Space** | [**Space**](Space.md) |  | 
**Cluster** | [**[]GlobalClusterStatus**](GlobalClusterStatus.md) | The status of the space cluster component. | 
**Network** | [**Network8**](Network8.md) |  | 

## Methods

### NewGlobalSpaceStatus

`func NewGlobalSpaceStatus(space Space, cluster []GlobalClusterStatus, network Network8, ) *GlobalSpaceStatus`

NewGlobalSpaceStatus instantiates a new GlobalSpaceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalSpaceStatusWithDefaults

`func NewGlobalSpaceStatusWithDefaults() *GlobalSpaceStatus`

NewGlobalSpaceStatusWithDefaults instantiates a new GlobalSpaceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpace

`func (o *GlobalSpaceStatus) GetSpace() Space`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *GlobalSpaceStatus) GetSpaceOk() (*Space, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *GlobalSpaceStatus) SetSpace(v Space)`

SetSpace sets Space field to given value.


### GetCluster

`func (o *GlobalSpaceStatus) GetCluster() []GlobalClusterStatus`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *GlobalSpaceStatus) GetClusterOk() (*[]GlobalClusterStatus, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *GlobalSpaceStatus) SetCluster(v []GlobalClusterStatus)`

SetCluster sets Cluster field to given value.


### GetNetwork

`func (o *GlobalSpaceStatus) GetNetwork() Network8`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GlobalSpaceStatus) GetNetworkOk() (*Network8, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GlobalSpaceStatus) SetNetwork(v Network8)`

SetNetwork sets Network field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


