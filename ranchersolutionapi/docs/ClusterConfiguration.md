# ClusterConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** | Shared secret used to join a server or agent to a cluster. | [optional] 
**TlsSan** | Pointer to **string** | This maps to ranchers &#x60;tls-san&#x60;. Add additional hostname or IP as a Subject Alternative Name in the TLS cert. | [optional] 
**EtcdSnapshotScheduleCron** | Pointer to **string** | This maps to ranchers &#x60;etcd-snapshot-schedule-cron&#x60;. Snapshot interval time in cron spec. eg. every 5 hours ‘0 *_/5 * * *’. Default: at 12 am/pm | [optional] [default to "0 0,12 * * *"]
**EtcdSnapshotRetention** | Pointer to **int32** | This maps to ranchers &#x60;etcd-snapshot-retention&#x60;. Number of snapshots to retain. | [optional] [default to 5]
**NodeTaint** | Pointer to **string** | This maps to ranchers &#x60;node-taint&#x60;. Registering kubelet with set of taints. By default, server nodes will be schedulable and thus your workloads can get launched on them. If you wish to have a dedicated control plane where no user workloads will run, you can use taints. | [optional] 
**ClusterDomain** | Pointer to **string** | This maps to ranchers &#x60;cluster-domain&#x60;. Cluster Domain. | [optional] 
**Certificates** | Pointer to [**RancherClusterConfigCertificates**](RancherClusterConfigCertificates.md) |  | [optional] 

## Methods

### NewClusterConfiguration

`func NewClusterConfiguration() *ClusterConfiguration`

NewClusterConfiguration instantiates a new ClusterConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterConfigurationWithDefaults

`func NewClusterConfigurationWithDefaults() *ClusterConfiguration`

NewClusterConfigurationWithDefaults instantiates a new ClusterConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *ClusterConfiguration) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ClusterConfiguration) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ClusterConfiguration) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ClusterConfiguration) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTlsSan

`func (o *ClusterConfiguration) GetTlsSan() string`

GetTlsSan returns the TlsSan field if non-nil, zero value otherwise.

### GetTlsSanOk

`func (o *ClusterConfiguration) GetTlsSanOk() (*string, bool)`

GetTlsSanOk returns a tuple with the TlsSan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsSan

`func (o *ClusterConfiguration) SetTlsSan(v string)`

SetTlsSan sets TlsSan field to given value.

### HasTlsSan

`func (o *ClusterConfiguration) HasTlsSan() bool`

HasTlsSan returns a boolean if a field has been set.

### GetEtcdSnapshotScheduleCron

`func (o *ClusterConfiguration) GetEtcdSnapshotScheduleCron() string`

GetEtcdSnapshotScheduleCron returns the EtcdSnapshotScheduleCron field if non-nil, zero value otherwise.

### GetEtcdSnapshotScheduleCronOk

`func (o *ClusterConfiguration) GetEtcdSnapshotScheduleCronOk() (*string, bool)`

GetEtcdSnapshotScheduleCronOk returns a tuple with the EtcdSnapshotScheduleCron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtcdSnapshotScheduleCron

`func (o *ClusterConfiguration) SetEtcdSnapshotScheduleCron(v string)`

SetEtcdSnapshotScheduleCron sets EtcdSnapshotScheduleCron field to given value.

### HasEtcdSnapshotScheduleCron

`func (o *ClusterConfiguration) HasEtcdSnapshotScheduleCron() bool`

HasEtcdSnapshotScheduleCron returns a boolean if a field has been set.

### GetEtcdSnapshotRetention

`func (o *ClusterConfiguration) GetEtcdSnapshotRetention() int32`

GetEtcdSnapshotRetention returns the EtcdSnapshotRetention field if non-nil, zero value otherwise.

### GetEtcdSnapshotRetentionOk

`func (o *ClusterConfiguration) GetEtcdSnapshotRetentionOk() (*int32, bool)`

GetEtcdSnapshotRetentionOk returns a tuple with the EtcdSnapshotRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtcdSnapshotRetention

`func (o *ClusterConfiguration) SetEtcdSnapshotRetention(v int32)`

SetEtcdSnapshotRetention sets EtcdSnapshotRetention field to given value.

### HasEtcdSnapshotRetention

`func (o *ClusterConfiguration) HasEtcdSnapshotRetention() bool`

HasEtcdSnapshotRetention returns a boolean if a field has been set.

### GetNodeTaint

`func (o *ClusterConfiguration) GetNodeTaint() string`

GetNodeTaint returns the NodeTaint field if non-nil, zero value otherwise.

### GetNodeTaintOk

`func (o *ClusterConfiguration) GetNodeTaintOk() (*string, bool)`

GetNodeTaintOk returns a tuple with the NodeTaint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeTaint

`func (o *ClusterConfiguration) SetNodeTaint(v string)`

SetNodeTaint sets NodeTaint field to given value.

### HasNodeTaint

`func (o *ClusterConfiguration) HasNodeTaint() bool`

HasNodeTaint returns a boolean if a field has been set.

### GetClusterDomain

`func (o *ClusterConfiguration) GetClusterDomain() string`

GetClusterDomain returns the ClusterDomain field if non-nil, zero value otherwise.

### GetClusterDomainOk

`func (o *ClusterConfiguration) GetClusterDomainOk() (*string, bool)`

GetClusterDomainOk returns a tuple with the ClusterDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterDomain

`func (o *ClusterConfiguration) SetClusterDomain(v string)`

SetClusterDomain sets ClusterDomain field to given value.

### HasClusterDomain

`func (o *ClusterConfiguration) HasClusterDomain() bool`

HasClusterDomain returns a boolean if a field has been set.

### GetCertificates

`func (o *ClusterConfiguration) GetCertificates() RancherClusterConfigCertificates`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *ClusterConfiguration) GetCertificatesOk() (*RancherClusterConfigCertificates, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *ClusterConfiguration) SetCertificates(v RancherClusterConfigCertificates)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *ClusterConfiguration) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


