/*
Rancher Solution API

Create and manage Rancher clusters. </br></br>**All URLs are relative to (https://api.phoenixnap.com/solutions/rancher/v0)**

API version: 0.1
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ranchersolutionapi

import (
	"encoding/json"
)

// Cluster Cluster details.
type Cluster struct {
	// (Read-only) The Cluster identifier.
	Id *string `json:"id,omitempty"`
	// Cluster name. This field is autogenerated if not provided.
	Name *string `json:"name,omitempty"`
	// Cluster description.
	Description *string `json:"description,omitempty"`
	// Deployment location.
	Location string `json:"location"`
	// (Read-only) The Rancher version that was installed on the cluster during the first creation process.
	InitialClusterVersion *string `json:"initialClusterVersion,omitempty"`
	// The node pools associated with the cluster.
	NodePools     *[]NodePool            `json:"nodePools,omitempty"`
	Configuration *RancherClusterConfig  `json:"configuration,omitempty"`
	Metadata      *RancherServerMetadata `json:"metadata,omitempty"`
	// The cluster status
	StatusDescription *string `json:"statusDescription,omitempty"`
}

// NewCluster instantiates a new Cluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCluster(location string) *Cluster {
	this := Cluster{}
	this.Location = location
	return &this
}

// NewClusterWithDefaults instantiates a new Cluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterWithDefaults() *Cluster {
	this := Cluster{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Cluster) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Cluster) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Cluster) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Cluster) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Cluster) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Cluster) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Cluster) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Cluster) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Cluster) SetDescription(v string) {
	o.Description = &v
}

// GetLocation returns the Location field value
func (o *Cluster) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *Cluster) SetLocation(v string) {
	o.Location = v
}

// GetInitialClusterVersion returns the InitialClusterVersion field value if set, zero value otherwise.
func (o *Cluster) GetInitialClusterVersion() string {
	if o == nil || o.InitialClusterVersion == nil {
		var ret string
		return ret
	}
	return *o.InitialClusterVersion
}

// GetInitialClusterVersionOk returns a tuple with the InitialClusterVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetInitialClusterVersionOk() (*string, bool) {
	if o == nil || o.InitialClusterVersion == nil {
		return nil, false
	}
	return o.InitialClusterVersion, true
}

// HasInitialClusterVersion returns a boolean if a field has been set.
func (o *Cluster) HasInitialClusterVersion() bool {
	if o != nil && o.InitialClusterVersion != nil {
		return true
	}

	return false
}

// SetInitialClusterVersion gets a reference to the given string and assigns it to the InitialClusterVersion field.
func (o *Cluster) SetInitialClusterVersion(v string) {
	o.InitialClusterVersion = &v
}

// GetNodePools returns the NodePools field value if set, zero value otherwise.
func (o *Cluster) GetNodePools() []NodePool {
	if o == nil || o.NodePools == nil {
		var ret []NodePool
		return ret
	}
	return *o.NodePools
}

// GetNodePoolsOk returns a tuple with the NodePools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetNodePoolsOk() (*[]NodePool, bool) {
	if o == nil || o.NodePools == nil {
		return nil, false
	}
	return o.NodePools, true
}

// HasNodePools returns a boolean if a field has been set.
func (o *Cluster) HasNodePools() bool {
	if o != nil && o.NodePools != nil {
		return true
	}

	return false
}

// SetNodePools gets a reference to the given []NodePool and assigns it to the NodePools field.
func (o *Cluster) SetNodePools(v []NodePool) {
	o.NodePools = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *Cluster) GetConfiguration() RancherClusterConfig {
	if o == nil || o.Configuration == nil {
		var ret RancherClusterConfig
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetConfigurationOk() (*RancherClusterConfig, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *Cluster) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given RancherClusterConfig and assigns it to the Configuration field.
func (o *Cluster) SetConfiguration(v RancherClusterConfig) {
	o.Configuration = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Cluster) GetMetadata() RancherServerMetadata {
	if o == nil || o.Metadata == nil {
		var ret RancherServerMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetMetadataOk() (*RancherServerMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Cluster) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given RancherServerMetadata and assigns it to the Metadata field.
func (o *Cluster) SetMetadata(v RancherServerMetadata) {
	o.Metadata = &v
}

// GetStatusDescription returns the StatusDescription field value if set, zero value otherwise.
func (o *Cluster) GetStatusDescription() string {
	if o == nil || o.StatusDescription == nil {
		var ret string
		return ret
	}
	return *o.StatusDescription
}

// GetStatusDescriptionOk returns a tuple with the StatusDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetStatusDescriptionOk() (*string, bool) {
	if o == nil || o.StatusDescription == nil {
		return nil, false
	}
	return o.StatusDescription, true
}

// HasStatusDescription returns a boolean if a field has been set.
func (o *Cluster) HasStatusDescription() bool {
	if o != nil && o.StatusDescription != nil {
		return true
	}

	return false
}

// SetStatusDescription gets a reference to the given string and assigns it to the StatusDescription field.
func (o *Cluster) SetStatusDescription(v string) {
	o.StatusDescription = &v
}

func (o Cluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["location"] = o.Location
	}
	if o.InitialClusterVersion != nil {
		toSerialize["initialClusterVersion"] = o.InitialClusterVersion
	}
	if o.NodePools != nil {
		toSerialize["nodePools"] = o.NodePools
	}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.StatusDescription != nil {
		toSerialize["statusDescription"] = o.StatusDescription
	}
	return json.Marshal(toSerialize)
}

type NullableCluster struct {
	value *Cluster
	isSet bool
}

func (v NullableCluster) Get() *Cluster {
	return v.value
}

func (v *NullableCluster) Set(val *Cluster) {
	v.value = val
	v.isSet = true
}

func (v NullableCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCluster(val *Cluster) *NullableCluster {
	return &NullableCluster{value: val, isSet: true}
}

func (v NullableCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
