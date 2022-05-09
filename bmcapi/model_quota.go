/*
Bare Metal Cloud API

Create, power off, power on, reset, reboot, or shut down your server with the Bare Metal Cloud API.  Deprovision servers, get or edit SSH key details, assign public IPs, assign servers to networks and a lot more.  Manage your infrastructure more efficiently using just a few simple API calls.<br> <br> <span class='pnap-api-knowledge-base-link'> Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/how-to-deploy-bare-metal-cloud-server' target='_blank'>here</a> </span><br> <br> <b>All URLs are relative to (https://api.phoenixnap.com/bmc/v1/)</b>

API version: 0.1
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bmcapi

import (
	"encoding/json"
)

// Quota Quota.
type Quota struct {
	// The ID of the Quota.
	Id string `json:"id"`
	// The name of the Quota.
	Name string `json:"name"`
	// The Quota description.
	Description string `json:"description"`
	// The status of the quota resource usage.
	Status string `json:"status"`
	// The limit set for the quota.
	Limit int32 `json:"limit"`
	// An enum field describing what the limit is measured in.
	Unit string `json:"unit"`
	// The quota used expressed as a number.
	Used                         int32                          `json:"used"`
	QuotaEditLimitRequestDetails []QuotaEditLimitRequestDetails `json:"quotaEditLimitRequestDetails"`
}

// NewQuota instantiates a new Quota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuota(id string, name string, description string, status string, limit int32, unit string, used int32, quotaEditLimitRequestDetails []QuotaEditLimitRequestDetails) *Quota {
	this := Quota{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.Status = status
	this.Limit = limit
	this.Unit = unit
	this.Used = used
	this.QuotaEditLimitRequestDetails = quotaEditLimitRequestDetails
	return &this
}

// NewQuotaWithDefaults instantiates a new Quota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotaWithDefaults() *Quota {
	this := Quota{}
	return &this
}

// GetId returns the Id field value
func (o *Quota) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Quota) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Quota) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Quota) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Quota) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Quota) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *Quota) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Quota) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Quota) SetDescription(v string) {
	o.Description = v
}

// GetStatus returns the Status field value
func (o *Quota) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Quota) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Quota) SetStatus(v string) {
	o.Status = v
}

// GetLimit returns the Limit field value
func (o *Quota) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *Quota) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *Quota) SetLimit(v int32) {
	o.Limit = v
}

// GetUnit returns the Unit field value
func (o *Quota) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *Quota) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *Quota) SetUnit(v string) {
	o.Unit = v
}

// GetUsed returns the Used field value
func (o *Quota) GetUsed() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Used
}

// GetUsedOk returns a tuple with the Used field value
// and a boolean to check if the value has been set.
func (o *Quota) GetUsedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Used, true
}

// SetUsed sets field value
func (o *Quota) SetUsed(v int32) {
	o.Used = v
}

// GetQuotaEditLimitRequestDetails returns the QuotaEditLimitRequestDetails field value
func (o *Quota) GetQuotaEditLimitRequestDetails() []QuotaEditLimitRequestDetails {
	if o == nil {
		var ret []QuotaEditLimitRequestDetails
		return ret
	}

	return o.QuotaEditLimitRequestDetails
}

// GetQuotaEditLimitRequestDetailsOk returns a tuple with the QuotaEditLimitRequestDetails field value
// and a boolean to check if the value has been set.
func (o *Quota) GetQuotaEditLimitRequestDetailsOk() (*[]QuotaEditLimitRequestDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuotaEditLimitRequestDetails, true
}

// SetQuotaEditLimitRequestDetails sets field value
func (o *Quota) SetQuotaEditLimitRequestDetails(v []QuotaEditLimitRequestDetails) {
	o.QuotaEditLimitRequestDetails = v
}

func (o Quota) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["limit"] = o.Limit
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	if true {
		toSerialize["used"] = o.Used
	}
	if true {
		toSerialize["quotaEditLimitRequestDetails"] = o.QuotaEditLimitRequestDetails
	}
	return json.Marshal(toSerialize)
}

type NullableQuota struct {
	value *Quota
	isSet bool
}

func (v NullableQuota) Get() *Quota {
	return v.value
}

func (v *NullableQuota) Set(val *Quota) {
	v.value = val
	v.isSet = true
}

func (v NullableQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuota(val *Quota) *NullableQuota {
	return &NullableQuota{value: val, isSet: true}
}

func (v NullableQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
