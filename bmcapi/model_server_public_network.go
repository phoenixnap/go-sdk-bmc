/*
Bare Metal Cloud API

Create, power off, power on, reset, reboot, or shut down your server with the Bare Metal Cloud API.  Deprovision servers, get or edit SSH key details, assign public IPs, assign servers to networks and a lot more.  Manage your infrastructure more efficiently using just a few simple API calls.<br> <br> <span class='pnap-api-knowledge-base-link'> Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/how-to-deploy-bare-metal-cloud-server' target='_blank'>here</a> </span><br> <br> <b>All URLs are relative to (https://api.phoenixnap.com/bmc/v1/)</b>

API version: 0.1
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bmcapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ServerPublicNetwork type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerPublicNetwork{}

// ServerPublicNetwork Public network details of bare metal server.
type ServerPublicNetwork struct {
	// The network identifier.
	Id string `json:"id"`
	// Configurable/configured IPs on the server.<br> At least 1 IP address is required. Valid IP formats are single IPv4 addresses or IPv4 ranges. All IPs must be within the network's range.<br> Setting the `force` query parameter to `true` allows you to:<ul> <li> Assign no specific IP addresses by designating an empty array of IPs. Note that at least one IP is required for the gateway address to be selected from this network. <li> Assign one or more IP addresses which are already configured on other resource(s) in network.</ul>
	Ips []string `json:"ips,omitempty"`
	// (Read-only) The status of the assignment to the network.
	StatusDescription *string `json:"statusDescription,omitempty"`
}

type _ServerPublicNetwork ServerPublicNetwork

// NewServerPublicNetwork instantiates a new ServerPublicNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerPublicNetwork(id string) *ServerPublicNetwork {
	this := ServerPublicNetwork{}
	this.Id = id
	return &this
}

// NewServerPublicNetworkWithDefaults instantiates a new ServerPublicNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerPublicNetworkWithDefaults() *ServerPublicNetwork {
	this := ServerPublicNetwork{}
	return &this
}

// GetId returns the Id field value
func (o *ServerPublicNetwork) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServerPublicNetwork) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServerPublicNetwork) SetId(v string) {
	o.Id = v
}

// GetIps returns the Ips field value if set, zero value otherwise.
func (o *ServerPublicNetwork) GetIps() []string {
	if o == nil || IsNil(o.Ips) {
		var ret []string
		return ret
	}
	return o.Ips
}

// GetIpsOk returns a tuple with the Ips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerPublicNetwork) GetIpsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ips) {
		return nil, false
	}
	return o.Ips, true
}

// HasIps returns a boolean if a field has been set.
func (o *ServerPublicNetwork) HasIps() bool {
	if o != nil && !IsNil(o.Ips) {
		return true
	}

	return false
}

// SetIps gets a reference to the given []string and assigns it to the Ips field.
func (o *ServerPublicNetwork) SetIps(v []string) {
	o.Ips = v
}

// GetStatusDescription returns the StatusDescription field value if set, zero value otherwise.
func (o *ServerPublicNetwork) GetStatusDescription() string {
	if o == nil || IsNil(o.StatusDescription) {
		var ret string
		return ret
	}
	return *o.StatusDescription
}

// GetStatusDescriptionOk returns a tuple with the StatusDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerPublicNetwork) GetStatusDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.StatusDescription) {
		return nil, false
	}
	return o.StatusDescription, true
}

// HasStatusDescription returns a boolean if a field has been set.
func (o *ServerPublicNetwork) HasStatusDescription() bool {
	if o != nil && !IsNil(o.StatusDescription) {
		return true
	}

	return false
}

// SetStatusDescription gets a reference to the given string and assigns it to the StatusDescription field.
func (o *ServerPublicNetwork) SetStatusDescription(v string) {
	o.StatusDescription = &v
}

func (o ServerPublicNetwork) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerPublicNetwork) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Ips) {
		toSerialize["ips"] = o.Ips
	}
	if !IsNil(o.StatusDescription) {
		toSerialize["statusDescription"] = o.StatusDescription
	}
	return toSerialize, nil
}

func (o *ServerPublicNetwork) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varServerPublicNetwork := _ServerPublicNetwork{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServerPublicNetwork)

	if err != nil {
		return err
	}

	*o = ServerPublicNetwork(varServerPublicNetwork)

	return err
}

type NullableServerPublicNetwork struct {
	value *ServerPublicNetwork
	isSet bool
}

func (v NullableServerPublicNetwork) Get() *ServerPublicNetwork {
	return v.value
}

func (v *NullableServerPublicNetwork) Set(val *ServerPublicNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableServerPublicNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableServerPublicNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerPublicNetwork(val *ServerPublicNetwork) *NullableServerPublicNetwork {
	return &NullableServerPublicNetwork{value: val, isSet: true}
}

func (v NullableServerPublicNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerPublicNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
