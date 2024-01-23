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

// ServerPrivateNetwork Private network details of bare metal server.
type ServerPrivateNetwork struct {
	// The network identifier.
	Id string `json:"id"`
	// IPs to configure/configured on the server.<br> Valid IP formats are single IPv4 addresses or IPv4 ranges. IPs must be within the network's range. Should be null or empty list if DHCP is true. <br> If field is undefined and DHCP is false, next available IP in network will be automatically allocated.<br> If the network contains a membership of type 'storage', the first twelve IPs are already reserved by BMC and not usable.<br> Setting the `force` query parameter to `true` allows you to:<ul> <li> Assign no specific IP addresses by designating an empty array of IPs. Note that at least one IP is required for the gateway address to be selected from this network. <li> Assign one or more IP addresses which are already configured on other resource(s) in network. <li> Assign IP addresses which are considered as reserved in network.</ul>
	Ips []string `json:"ips,omitempty"`
	// Determines whether DHCP is enabled for this server. Should be false if any IPs are provided. Not supported for Proxmox OS.
	Dhcp *bool `json:"dhcp,omitempty"`
	// (Read-only) The status of the network.
	StatusDescription *string `json:"statusDescription,omitempty"`
}

// NewServerPrivateNetwork instantiates a new ServerPrivateNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerPrivateNetwork(id string) *ServerPrivateNetwork {
	this := ServerPrivateNetwork{}
	this.Id = id
	var dhcp bool = false
	this.Dhcp = &dhcp
	return &this
}

// NewServerPrivateNetworkWithDefaults instantiates a new ServerPrivateNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerPrivateNetworkWithDefaults() *ServerPrivateNetwork {
	this := ServerPrivateNetwork{}
	var dhcp bool = false
	this.Dhcp = &dhcp
	return &this
}

// GetId returns the Id field value
func (o *ServerPrivateNetwork) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServerPrivateNetwork) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServerPrivateNetwork) SetId(v string) {
	o.Id = v
}

// GetIps returns the Ips field value if set, zero value otherwise.
func (o *ServerPrivateNetwork) GetIps() []string {
	if o == nil || o.Ips == nil {
		var ret []string
		return ret
	}
	return o.Ips
}

// GetIpsOk returns a tuple with the Ips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerPrivateNetwork) GetIpsOk() ([]string, bool) {
	if o == nil || o.Ips == nil {
		return nil, false
	}
	return o.Ips, true
}

// HasIps returns a boolean if a field has been set.
func (o *ServerPrivateNetwork) HasIps() bool {
	if o != nil && o.Ips != nil {
		return true
	}

	return false
}

// SetIps gets a reference to the given []string and assigns it to the Ips field.
func (o *ServerPrivateNetwork) SetIps(v []string) {
	o.Ips = v
}

// GetDhcp returns the Dhcp field value if set, zero value otherwise.
func (o *ServerPrivateNetwork) GetDhcp() bool {
	if o == nil || o.Dhcp == nil {
		var ret bool
		return ret
	}
	return *o.Dhcp
}

// GetDhcpOk returns a tuple with the Dhcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerPrivateNetwork) GetDhcpOk() (*bool, bool) {
	if o == nil || o.Dhcp == nil {
		return nil, false
	}
	return o.Dhcp, true
}

// HasDhcp returns a boolean if a field has been set.
func (o *ServerPrivateNetwork) HasDhcp() bool {
	if o != nil && o.Dhcp != nil {
		return true
	}

	return false
}

// SetDhcp gets a reference to the given bool and assigns it to the Dhcp field.
func (o *ServerPrivateNetwork) SetDhcp(v bool) {
	o.Dhcp = &v
}

// GetStatusDescription returns the StatusDescription field value if set, zero value otherwise.
func (o *ServerPrivateNetwork) GetStatusDescription() string {
	if o == nil || o.StatusDescription == nil {
		var ret string
		return ret
	}
	return *o.StatusDescription
}

// GetStatusDescriptionOk returns a tuple with the StatusDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerPrivateNetwork) GetStatusDescriptionOk() (*string, bool) {
	if o == nil || o.StatusDescription == nil {
		return nil, false
	}
	return o.StatusDescription, true
}

// HasStatusDescription returns a boolean if a field has been set.
func (o *ServerPrivateNetwork) HasStatusDescription() bool {
	if o != nil && o.StatusDescription != nil {
		return true
	}

	return false
}

// SetStatusDescription gets a reference to the given string and assigns it to the StatusDescription field.
func (o *ServerPrivateNetwork) SetStatusDescription(v string) {
	o.StatusDescription = &v
}

func (o ServerPrivateNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Ips != nil {
		toSerialize["ips"] = o.Ips
	}
	if o.Dhcp != nil {
		toSerialize["dhcp"] = o.Dhcp
	}
	if o.StatusDescription != nil {
		toSerialize["statusDescription"] = o.StatusDescription
	}
	return json.Marshal(toSerialize)
}

type NullableServerPrivateNetwork struct {
	value *ServerPrivateNetwork
	isSet bool
}

func (v NullableServerPrivateNetwork) Get() *ServerPrivateNetwork {
	return v.value
}

func (v *NullableServerPrivateNetwork) Set(val *ServerPrivateNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableServerPrivateNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableServerPrivateNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerPrivateNetwork(val *ServerPrivateNetwork) *NullableServerPrivateNetwork {
	return &NullableServerPrivateNetwork{value: val, isSet: true}
}

func (v NullableServerPrivateNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerPrivateNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
