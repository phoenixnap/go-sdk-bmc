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

// OsConfiguration OS specific configuration properties.
type OsConfiguration struct {
	Windows *OsConfigurationWindows `json:"windows,omitempty"`
	// Password set for user root on an ESXi server which will only be returned in response to provisioning a server.
	RootPassword *string `json:"rootPassword,omitempty"`
	// The URL of the management UI which will only be returned in response to provisioning a server.
	ManagementUiUrl *string `json:"managementUiUrl,omitempty"`
	// List of IPs allowed to access the Management UI. Supported in single IP, CIDR and range format. When undefined, Management UI is disabled. This will only be returned in response to provisioning a server.
	ManagementAccessAllowedIps []string `json:"managementAccessAllowedIps,omitempty"`
	// If true, OS will be installed to and booted from the server's RAM. On restart RAM OS will be lost and the server will not be reachable unless a custom bootable OS has been deployed. Only supported for ubuntu/focal.
	InstallOsToRam *bool                     `json:"installOsToRam,omitempty"`
	CloudInit      *OsConfigurationCloudInit `json:"cloudInit,omitempty"`
}

// NewOsConfiguration instantiates a new OsConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsConfiguration() *OsConfiguration {
	this := OsConfiguration{}
	var installOsToRam bool = false
	this.InstallOsToRam = &installOsToRam
	return &this
}

// NewOsConfigurationWithDefaults instantiates a new OsConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsConfigurationWithDefaults() *OsConfiguration {
	this := OsConfiguration{}
	var installOsToRam bool = false
	this.InstallOsToRam = &installOsToRam
	return &this
}

// GetWindows returns the Windows field value if set, zero value otherwise.
func (o *OsConfiguration) GetWindows() OsConfigurationWindows {
	if o == nil || o.Windows == nil {
		var ret OsConfigurationWindows
		return ret
	}
	return *o.Windows
}

// GetWindowsOk returns a tuple with the Windows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsConfiguration) GetWindowsOk() (*OsConfigurationWindows, bool) {
	if o == nil || o.Windows == nil {
		return nil, false
	}
	return o.Windows, true
}

// HasWindows returns a boolean if a field has been set.
func (o *OsConfiguration) HasWindows() bool {
	if o != nil && o.Windows != nil {
		return true
	}

	return false
}

// SetWindows gets a reference to the given OsConfigurationWindows and assigns it to the Windows field.
func (o *OsConfiguration) SetWindows(v OsConfigurationWindows) {
	o.Windows = &v
}

// GetRootPassword returns the RootPassword field value if set, zero value otherwise.
func (o *OsConfiguration) GetRootPassword() string {
	if o == nil || o.RootPassword == nil {
		var ret string
		return ret
	}
	return *o.RootPassword
}

// GetRootPasswordOk returns a tuple with the RootPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsConfiguration) GetRootPasswordOk() (*string, bool) {
	if o == nil || o.RootPassword == nil {
		return nil, false
	}
	return o.RootPassword, true
}

// HasRootPassword returns a boolean if a field has been set.
func (o *OsConfiguration) HasRootPassword() bool {
	if o != nil && o.RootPassword != nil {
		return true
	}

	return false
}

// SetRootPassword gets a reference to the given string and assigns it to the RootPassword field.
func (o *OsConfiguration) SetRootPassword(v string) {
	o.RootPassword = &v
}

// GetManagementUiUrl returns the ManagementUiUrl field value if set, zero value otherwise.
func (o *OsConfiguration) GetManagementUiUrl() string {
	if o == nil || o.ManagementUiUrl == nil {
		var ret string
		return ret
	}
	return *o.ManagementUiUrl
}

// GetManagementUiUrlOk returns a tuple with the ManagementUiUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsConfiguration) GetManagementUiUrlOk() (*string, bool) {
	if o == nil || o.ManagementUiUrl == nil {
		return nil, false
	}
	return o.ManagementUiUrl, true
}

// HasManagementUiUrl returns a boolean if a field has been set.
func (o *OsConfiguration) HasManagementUiUrl() bool {
	if o != nil && o.ManagementUiUrl != nil {
		return true
	}

	return false
}

// SetManagementUiUrl gets a reference to the given string and assigns it to the ManagementUiUrl field.
func (o *OsConfiguration) SetManagementUiUrl(v string) {
	o.ManagementUiUrl = &v
}

// GetManagementAccessAllowedIps returns the ManagementAccessAllowedIps field value if set, zero value otherwise.
func (o *OsConfiguration) GetManagementAccessAllowedIps() []string {
	if o == nil || o.ManagementAccessAllowedIps == nil {
		var ret []string
		return ret
	}
	return o.ManagementAccessAllowedIps
}

// GetManagementAccessAllowedIpsOk returns a tuple with the ManagementAccessAllowedIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsConfiguration) GetManagementAccessAllowedIpsOk() ([]string, bool) {
	if o == nil || o.ManagementAccessAllowedIps == nil {
		return nil, false
	}
	return o.ManagementAccessAllowedIps, true
}

// HasManagementAccessAllowedIps returns a boolean if a field has been set.
func (o *OsConfiguration) HasManagementAccessAllowedIps() bool {
	if o != nil && o.ManagementAccessAllowedIps != nil {
		return true
	}

	return false
}

// SetManagementAccessAllowedIps gets a reference to the given []string and assigns it to the ManagementAccessAllowedIps field.
func (o *OsConfiguration) SetManagementAccessAllowedIps(v []string) {
	o.ManagementAccessAllowedIps = v
}

// GetInstallOsToRam returns the InstallOsToRam field value if set, zero value otherwise.
func (o *OsConfiguration) GetInstallOsToRam() bool {
	if o == nil || o.InstallOsToRam == nil {
		var ret bool
		return ret
	}
	return *o.InstallOsToRam
}

// GetInstallOsToRamOk returns a tuple with the InstallOsToRam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsConfiguration) GetInstallOsToRamOk() (*bool, bool) {
	if o == nil || o.InstallOsToRam == nil {
		return nil, false
	}
	return o.InstallOsToRam, true
}

// HasInstallOsToRam returns a boolean if a field has been set.
func (o *OsConfiguration) HasInstallOsToRam() bool {
	if o != nil && o.InstallOsToRam != nil {
		return true
	}

	return false
}

// SetInstallOsToRam gets a reference to the given bool and assigns it to the InstallOsToRam field.
func (o *OsConfiguration) SetInstallOsToRam(v bool) {
	o.InstallOsToRam = &v
}

// GetCloudInit returns the CloudInit field value if set, zero value otherwise.
func (o *OsConfiguration) GetCloudInit() OsConfigurationCloudInit {
	if o == nil || o.CloudInit == nil {
		var ret OsConfigurationCloudInit
		return ret
	}
	return *o.CloudInit
}

// GetCloudInitOk returns a tuple with the CloudInit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsConfiguration) GetCloudInitOk() (*OsConfigurationCloudInit, bool) {
	if o == nil || o.CloudInit == nil {
		return nil, false
	}
	return o.CloudInit, true
}

// HasCloudInit returns a boolean if a field has been set.
func (o *OsConfiguration) HasCloudInit() bool {
	if o != nil && o.CloudInit != nil {
		return true
	}

	return false
}

// SetCloudInit gets a reference to the given OsConfigurationCloudInit and assigns it to the CloudInit field.
func (o *OsConfiguration) SetCloudInit(v OsConfigurationCloudInit) {
	o.CloudInit = &v
}

func (o OsConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Windows != nil {
		toSerialize["windows"] = o.Windows
	}
	if o.RootPassword != nil {
		toSerialize["rootPassword"] = o.RootPassword
	}
	if o.ManagementUiUrl != nil {
		toSerialize["managementUiUrl"] = o.ManagementUiUrl
	}
	if o.ManagementAccessAllowedIps != nil {
		toSerialize["managementAccessAllowedIps"] = o.ManagementAccessAllowedIps
	}
	if o.InstallOsToRam != nil {
		toSerialize["installOsToRam"] = o.InstallOsToRam
	}
	if o.CloudInit != nil {
		toSerialize["cloudInit"] = o.CloudInit
	}
	return json.Marshal(toSerialize)
}

type NullableOsConfiguration struct {
	value *OsConfiguration
	isSet bool
}

func (v NullableOsConfiguration) Get() *OsConfiguration {
	return v.value
}

func (v *NullableOsConfiguration) Set(val *OsConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableOsConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableOsConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsConfiguration(val *OsConfiguration) *NullableOsConfiguration {
	return &NullableOsConfiguration{value: val, isSet: true}
}

func (v NullableOsConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
