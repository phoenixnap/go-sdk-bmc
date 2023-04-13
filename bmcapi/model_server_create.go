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

// ServerCreate Provision bare metal server.
type ServerCreate struct {
	// Hostname of server.
	Hostname string `json:"hostname"`
	// Description of server.
	Description *string `json:"description,omitempty"`
	// The server’s OS ID used when the server was created. Currently this field should be set to either `ubuntu/bionic`, `ubuntu/focal`, `ubuntu/jammy`, `centos/centos7`, `centos/centos8`, `windows/srv2019std`, `windows/srv2019dc`, `esxi/esxi70`, `debian/bullseye`, `proxmox/bullseye`, `netris/controller`, `netris/softgate_1g` or `netris/softgate_10g`.
	Os string `json:"os"`
	// Server type ID. Cannot be changed once a server is created. Currently this field should be set to either `s0.d1.small`, `s0.d1.medium`, `s1.c1.small`, `s1.c1.medium`, `s1.c2.medium`, `s1.c2.large`, `s1.e1.small`, `s1.e1.medium`, `s1.e1.large`, `s2.c1.small`, `s2.c1.medium`, `s2.c1.large`, `s2.c2.small`, `s2.c2.medium`, `s2.c2.large`, `d1.c1.small`, `d1.c2.small`, `d1.c3.small`, `d1.c4.small`, `d1.c1.medium`, `d1.c2.medium`, `d1.c3.medium`, `d1.c4.medium`, `d1.c1.large`, `d1.c2.large`, `d1.c3.large`, `d1.c4.large`, `d1.m1.medium`, `d1.m2.medium`, `d1.m3.medium`, `d1.m4.medium`, `d2.c1.medium`, `d2.c2.medium`, `d2.c3.medium`, `d2.c4.medium`, `d2.c5.medium`, `d2.c1.large`, `d2.c2.large`, `d2.c3.large`, `d2.c4.large`, `d2.c5.large`, `d2.m1.medium`, `d2.m1.large`, `d2.m2.medium`, `d2.m2.large`, `d2.m2.xlarge`, `d2.c4.storage.pliops1`, `d3.m4.xlarge`, `d3.m5.xlarge` or `d3.m6.xlarge`.
	Type string `json:"type"`
	// Server location ID. Cannot be changed once a server is created. Currently this field should be set to `PHX`, `ASH`, `SGP`, `NLD`, `CHI`, `SEA` or `AUS`.
	Location string `json:"location"`
	// Whether or not to install SSH keys marked as default in addition to any SSH keys specified in this request.
	InstallDefaultSshKeys *bool `json:"installDefaultSshKeys,omitempty"`
	// A list of SSH keys that will be installed on the server.
	SshKeys []string `json:"sshKeys,omitempty"`
	// A list of SSH key IDs that will be installed on the server in addition to any SSH keys specified in this request.
	SshKeyIds []string `json:"sshKeyIds,omitempty"`
	// Server reservation ID.
	ReservationId *string `json:"reservationId,omitempty"`
	// Server pricing model. Currently this field should be set to `HOURLY`, `ONE_MONTH_RESERVATION`, `TWELVE_MONTHS_RESERVATION`, `TWENTY_FOUR_MONTHS_RESERVATION` or `THIRTY_SIX_MONTHS_RESERVATION`.
	PricingModel *string `json:"pricingModel,omitempty"`
	// The type of network configuration for this server. Currently this field should be set to `PUBLIC_AND_PRIVATE` or `PRIVATE_ONLY`.
	NetworkType     *string          `json:"networkType,omitempty"`
	OsConfiguration *OsConfiguration `json:"osConfiguration,omitempty"`
	// Tags to set to the server. To create a new tag or list all the existing tags that you can use, refer to [Tags API](https://developers.phoenixnap.com/docs/tags/1/overview).
	Tags                 []TagAssignmentRequest `json:"tags,omitempty"`
	NetworkConfiguration *NetworkConfiguration  `json:"networkConfiguration,omitempty"`
}

// NewServerCreate instantiates a new ServerCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerCreate(hostname string, os string, type_ string, location string) *ServerCreate {
	this := ServerCreate{}
	this.Hostname = hostname
	this.Os = os
	this.Type = type_
	this.Location = location
	var installDefaultSshKeys bool = true
	this.InstallDefaultSshKeys = &installDefaultSshKeys
	var pricingModel string = "HOURLY"
	this.PricingModel = &pricingModel
	var networkType string = "PUBLIC_AND_PRIVATE"
	this.NetworkType = &networkType
	return &this
}

// NewServerCreateWithDefaults instantiates a new ServerCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerCreateWithDefaults() *ServerCreate {
	this := ServerCreate{}
	var installDefaultSshKeys bool = true
	this.InstallDefaultSshKeys = &installDefaultSshKeys
	var pricingModel string = "HOURLY"
	this.PricingModel = &pricingModel
	var networkType string = "PUBLIC_AND_PRIVATE"
	this.NetworkType = &networkType
	return &this
}

// GetHostname returns the Hostname field value
func (o *ServerCreate) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *ServerCreate) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *ServerCreate) SetHostname(v string) {
	o.Hostname = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServerCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServerCreate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServerCreate) SetDescription(v string) {
	o.Description = &v
}

// GetOs returns the Os field value
func (o *ServerCreate) GetOs() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Os
}

// GetOsOk returns a tuple with the Os field value
// and a boolean to check if the value has been set.
func (o *ServerCreate) GetOsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Os, true
}

// SetOs sets field value
func (o *ServerCreate) SetOs(v string) {
	o.Os = v
}

// GetType returns the Type field value
func (o *ServerCreate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ServerCreate) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ServerCreate) SetType(v string) {
	o.Type = v
}

// GetLocation returns the Location field value
func (o *ServerCreate) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *ServerCreate) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *ServerCreate) SetLocation(v string) {
	o.Location = v
}

// GetInstallDefaultSshKeys returns the InstallDefaultSshKeys field value if set, zero value otherwise.
func (o *ServerCreate) GetInstallDefaultSshKeys() bool {
	if o == nil || o.InstallDefaultSshKeys == nil {
		var ret bool
		return ret
	}
	return *o.InstallDefaultSshKeys
}

// GetInstallDefaultSshKeysOk returns a tuple with the InstallDefaultSshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerCreate) GetInstallDefaultSshKeysOk() (*bool, bool) {
	if o == nil || o.InstallDefaultSshKeys == nil {
		return nil, false
	}
	return o.InstallDefaultSshKeys, true
}

// HasInstallDefaultSshKeys returns a boolean if a field has been set.
func (o *ServerCreate) HasInstallDefaultSshKeys() bool {
	if o != nil && o.InstallDefaultSshKeys != nil {
		return true
	}

	return false
}

// SetInstallDefaultSshKeys gets a reference to the given bool and assigns it to the InstallDefaultSshKeys field.
func (o *ServerCreate) SetInstallDefaultSshKeys(v bool) {
	o.InstallDefaultSshKeys = &v
}

// GetSshKeys returns the SshKeys field value if set, zero value otherwise.
func (o *ServerCreate) GetSshKeys() []string {
	if o == nil || o.SshKeys == nil {
		var ret []string
		return ret
	}
	return o.SshKeys
}

// GetSshKeysOk returns a tuple with the SshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerCreate) GetSshKeysOk() ([]string, bool) {
	if o == nil || o.SshKeys == nil {
		return nil, false
	}
	return o.SshKeys, true
}

// HasSshKeys returns a boolean if a field has been set.
func (o *ServerCreate) HasSshKeys() bool {
	if o != nil && o.SshKeys != nil {
		return true
	}

	return false
}

// SetSshKeys gets a reference to the given []string and assigns it to the SshKeys field.
func (o *ServerCreate) SetSshKeys(v []string) {
	o.SshKeys = v
}

// GetSshKeyIds returns the SshKeyIds field value if set, zero value otherwise.
func (o *ServerCreate) GetSshKeyIds() []string {
	if o == nil || o.SshKeyIds == nil {
		var ret []string
		return ret
	}
	return o.SshKeyIds
}

// GetSshKeyIdsOk returns a tuple with the SshKeyIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerCreate) GetSshKeyIdsOk() ([]string, bool) {
	if o == nil || o.SshKeyIds == nil {
		return nil, false
	}
	return o.SshKeyIds, true
}

// HasSshKeyIds returns a boolean if a field has been set.
func (o *ServerCreate) HasSshKeyIds() bool {
	if o != nil && o.SshKeyIds != nil {
		return true
	}

	return false
}

// SetSshKeyIds gets a reference to the given []string and assigns it to the SshKeyIds field.
func (o *ServerCreate) SetSshKeyIds(v []string) {
	o.SshKeyIds = v
}

// GetReservationId returns the ReservationId field value if set, zero value otherwise.
func (o *ServerCreate) GetReservationId() string {
	if o == nil || o.ReservationId == nil {
		var ret string
		return ret
	}
	return *o.ReservationId
}

// GetReservationIdOk returns a tuple with the ReservationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerCreate) GetReservationIdOk() (*string, bool) {
	if o == nil || o.ReservationId == nil {
		return nil, false
	}
	return o.ReservationId, true
}

// HasReservationId returns a boolean if a field has been set.
func (o *ServerCreate) HasReservationId() bool {
	if o != nil && o.ReservationId != nil {
		return true
	}

	return false
}

// SetReservationId gets a reference to the given string and assigns it to the ReservationId field.
func (o *ServerCreate) SetReservationId(v string) {
	o.ReservationId = &v
}

// GetPricingModel returns the PricingModel field value if set, zero value otherwise.
func (o *ServerCreate) GetPricingModel() string {
	if o == nil || o.PricingModel == nil {
		var ret string
		return ret
	}
	return *o.PricingModel
}

// GetPricingModelOk returns a tuple with the PricingModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerCreate) GetPricingModelOk() (*string, bool) {
	if o == nil || o.PricingModel == nil {
		return nil, false
	}
	return o.PricingModel, true
}

// HasPricingModel returns a boolean if a field has been set.
func (o *ServerCreate) HasPricingModel() bool {
	if o != nil && o.PricingModel != nil {
		return true
	}

	return false
}

// SetPricingModel gets a reference to the given string and assigns it to the PricingModel field.
func (o *ServerCreate) SetPricingModel(v string) {
	o.PricingModel = &v
}

// GetNetworkType returns the NetworkType field value if set, zero value otherwise.
func (o *ServerCreate) GetNetworkType() string {
	if o == nil || o.NetworkType == nil {
		var ret string
		return ret
	}
	return *o.NetworkType
}

// GetNetworkTypeOk returns a tuple with the NetworkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerCreate) GetNetworkTypeOk() (*string, bool) {
	if o == nil || o.NetworkType == nil {
		return nil, false
	}
	return o.NetworkType, true
}

// HasNetworkType returns a boolean if a field has been set.
func (o *ServerCreate) HasNetworkType() bool {
	if o != nil && o.NetworkType != nil {
		return true
	}

	return false
}

// SetNetworkType gets a reference to the given string and assigns it to the NetworkType field.
func (o *ServerCreate) SetNetworkType(v string) {
	o.NetworkType = &v
}

// GetOsConfiguration returns the OsConfiguration field value if set, zero value otherwise.
func (o *ServerCreate) GetOsConfiguration() OsConfiguration {
	if o == nil || o.OsConfiguration == nil {
		var ret OsConfiguration
		return ret
	}
	return *o.OsConfiguration
}

// GetOsConfigurationOk returns a tuple with the OsConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerCreate) GetOsConfigurationOk() (*OsConfiguration, bool) {
	if o == nil || o.OsConfiguration == nil {
		return nil, false
	}
	return o.OsConfiguration, true
}

// HasOsConfiguration returns a boolean if a field has been set.
func (o *ServerCreate) HasOsConfiguration() bool {
	if o != nil && o.OsConfiguration != nil {
		return true
	}

	return false
}

// SetOsConfiguration gets a reference to the given OsConfiguration and assigns it to the OsConfiguration field.
func (o *ServerCreate) SetOsConfiguration(v OsConfiguration) {
	o.OsConfiguration = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ServerCreate) GetTags() []TagAssignmentRequest {
	if o == nil || o.Tags == nil {
		var ret []TagAssignmentRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerCreate) GetTagsOk() ([]TagAssignmentRequest, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ServerCreate) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagAssignmentRequest and assigns it to the Tags field.
func (o *ServerCreate) SetTags(v []TagAssignmentRequest) {
	o.Tags = v
}

// GetNetworkConfiguration returns the NetworkConfiguration field value if set, zero value otherwise.
func (o *ServerCreate) GetNetworkConfiguration() NetworkConfiguration {
	if o == nil || o.NetworkConfiguration == nil {
		var ret NetworkConfiguration
		return ret
	}
	return *o.NetworkConfiguration
}

// GetNetworkConfigurationOk returns a tuple with the NetworkConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerCreate) GetNetworkConfigurationOk() (*NetworkConfiguration, bool) {
	if o == nil || o.NetworkConfiguration == nil {
		return nil, false
	}
	return o.NetworkConfiguration, true
}

// HasNetworkConfiguration returns a boolean if a field has been set.
func (o *ServerCreate) HasNetworkConfiguration() bool {
	if o != nil && o.NetworkConfiguration != nil {
		return true
	}

	return false
}

// SetNetworkConfiguration gets a reference to the given NetworkConfiguration and assigns it to the NetworkConfiguration field.
func (o *ServerCreate) SetNetworkConfiguration(v NetworkConfiguration) {
	o.NetworkConfiguration = &v
}

func (o ServerCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["hostname"] = o.Hostname
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["os"] = o.Os
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["location"] = o.Location
	}
	if o.InstallDefaultSshKeys != nil {
		toSerialize["installDefaultSshKeys"] = o.InstallDefaultSshKeys
	}
	if o.SshKeys != nil {
		toSerialize["sshKeys"] = o.SshKeys
	}
	if o.SshKeyIds != nil {
		toSerialize["sshKeyIds"] = o.SshKeyIds
	}
	if o.ReservationId != nil {
		toSerialize["reservationId"] = o.ReservationId
	}
	if o.PricingModel != nil {
		toSerialize["pricingModel"] = o.PricingModel
	}
	if o.NetworkType != nil {
		toSerialize["networkType"] = o.NetworkType
	}
	if o.OsConfiguration != nil {
		toSerialize["osConfiguration"] = o.OsConfiguration
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.NetworkConfiguration != nil {
		toSerialize["networkConfiguration"] = o.NetworkConfiguration
	}
	return json.Marshal(toSerialize)
}

type NullableServerCreate struct {
	value *ServerCreate
	isSet bool
}

func (v NullableServerCreate) Get() *ServerCreate {
	return v.value
}

func (v *NullableServerCreate) Set(val *ServerCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableServerCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableServerCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerCreate(val *ServerCreate) *NullableServerCreate {
	return &NullableServerCreate{value: val, isSet: true}
}

func (v NullableServerCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
