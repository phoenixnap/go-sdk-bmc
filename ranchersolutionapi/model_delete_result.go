/*
Rancher Solution API

Simplify enterprise-grade Kubernetes cluster operations and management with Rancher on Bare Metal Cloud. Deploy Kubernetes clusters using a few API calls.<br> <br> <span class='pnap-api-knowledge-base-link'> Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/rancher-bmc-integration-kubernetes' target='_blank'>here</a> </span><br> <br> <b>All URLs are relative to (https://api.phoenixnap.com/solutions/rancher/v1beta)</b> 

API version: 0.1
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ranchersolutionapi

import (
	"encoding/json"
)

// DeleteResult Result of a successful delete action.
type DeleteResult struct {
	// Solution cluster has been deleted.
	Result string `json:"result"`
	// The unique identifier of the solution cluster.
	ClusterId string `json:"clusterId"`
}

// NewDeleteResult instantiates a new DeleteResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteResult(result string, clusterId string) *DeleteResult {
	this := DeleteResult{}
	this.Result = result
	this.ClusterId = clusterId
	return &this
}

// NewDeleteResultWithDefaults instantiates a new DeleteResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteResultWithDefaults() *DeleteResult {
	this := DeleteResult{}
	return &this
}

// GetResult returns the Result field value
func (o *DeleteResult) GetResult() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *DeleteResult) GetResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *DeleteResult) SetResult(v string) {
	o.Result = v
}

// GetClusterId returns the ClusterId field value
func (o *DeleteResult) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *DeleteResult) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value
func (o *DeleteResult) SetClusterId(v string) {
	o.ClusterId = v
}

func (o DeleteResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["result"] = o.Result
	}
	if true {
		toSerialize["clusterId"] = o.ClusterId
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteResult struct {
	value *DeleteResult
	isSet bool
}

func (v NullableDeleteResult) Get() *DeleteResult {
	return v.value
}

func (v *NullableDeleteResult) Set(val *DeleteResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteResult(val *DeleteResult) *NullableDeleteResult {
	return &NullableDeleteResult{value: val, isSet: true}
}

func (v NullableDeleteResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


