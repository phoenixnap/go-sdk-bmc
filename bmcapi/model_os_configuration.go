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

// checks if the OsConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OsConfiguration{}

// OsConfiguration OS specific configuration properties.
type OsConfiguration struct {
	NetrisController *OsConfigurationNetrisController `json:"netrisController,omitempty"`
	NetrisSoftgate   *OsConfigurationNetrisSoftgate   `json:"netrisSoftgate,omitempty"`
	Windows          *OsConfigurationWindows          `json:"windows,omitempty"`
	// (Read-only) Auto-generated password set for user 'root' on an ESXi or Proxmox server.<br>  The password is not stored and therefore will only be returned in response to provisioning a server. Copy and save it for future reference.
	RootPassword *string `json:"rootPassword,omitempty"`
	// (Read-only) The URL of the management UI which will only be returned in response to provisioning a server.
	ManagementUiUrl *string `json:"managementUiUrl,omitempty"`
	// List of IPs allowed to access the Management UI. Supported in single IP, CIDR and range format. When undefined, Management UI is disabled. This will only be returned in response to provisioning a server.
	ManagementAccessAllowedIps []string `json:"managementAccessAllowedIps,omitempty"`
	// If true, OS will be installed to and booted from the server's RAM. On restart RAM OS will be lost and the server will not be reachable unless a custom bootable OS has been deployed. Follow the <a href='https://phoenixnap.com/kb/bmc-custom-os' target='_blank'>instructions</a> on how to install custom OS on BMC. Only supported for ubuntu/focal and ubuntu/jammy.
	InstallOsToRam       *bool                     `json:"installOsToRam,omitempty"`
	Esxi                 *EsxiOsConfiguration      `json:"esxi,omitempty"`
	CloudInit            *OsConfigurationCloudInit `json:"cloudInit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OsConfiguration OsConfiguration

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

// GetNetrisController returns the NetrisController field value if set, zero value otherwise.
func (o *OsConfiguration) GetNetrisController() OsConfigurationNetrisController {
	if o == nil || IsNil(o.NetrisController) {
		var ret OsConfigurationNetrisController
		return ret
	}
	return *o.NetrisController
}

// GetNetrisControllerOk returns a tuple with the NetrisController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsConfiguration) GetNetrisControllerOk() (*OsConfigurationNetrisController, bool) {
	if o == nil || IsNil(o.NetrisController) {
		return nil, false
	}
	return o.NetrisController, true
}

// HasNetrisController returns a boolean if a field has been set.
func (o *OsConfiguration) HasNetrisController() bool {
	if o != nil && !IsNil(o.NetrisController) {
		return true
	}

	return false
}

// SetNetrisController gets a reference to the given OsConfigurationNetrisController and assigns it to the NetrisController field.
func (o *OsConfiguration) SetNetrisController(v OsConfigurationNetrisController) {
	o.NetrisController = &v
}

// GetNetrisSoftgate returns the NetrisSoftgate field value if set, zero value otherwise.
func (o *OsConfiguration) GetNetrisSoftgate() OsConfigurationNetrisSoftgate {
	if o == nil || IsNil(o.NetrisSoftgate) {
		var ret OsConfigurationNetrisSoftgate
		return ret
	}
	return *o.NetrisSoftgate
}

// GetNetrisSoftgateOk returns a tuple with the NetrisSoftgate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsConfiguration) GetNetrisSoftgateOk() (*OsConfigurationNetrisSoftgate, bool) {
	if o == nil || IsNil(o.NetrisSoftgate) {
		return nil, false
	}
	return o.NetrisSoftgate, true
}

// HasNetrisSoftgate returns a boolean if a field has been set.
func (o *OsConfiguration) HasNetrisSoftgate() bool {
	if o != nil && !IsNil(o.NetrisSoftgate) {
		return true
	}

	return false
}

// SetNetrisSoftgate gets a reference to the given OsConfigurationNetrisSoftgate and assigns it to the NetrisSoftgate field.
func (o *OsConfiguration) SetNetrisSoftgate(v OsConfigurationNetrisSoftgate) {
	o.NetrisSoftgate = &v
}

// GetWindows returns the Windows field value if set, zero value otherwise.
func (o *OsConfiguration) GetWindows() OsConfigurationWindows {
	if o == nil || IsNil(o.Windows) {
		var ret OsConfigurationWindows
		return ret
	}
	return *o.Windows
}

// GetWindowsOk returns a tuple with the Windows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsConfiguration) GetWindowsOk() (*OsConfigurationWindows, bool) {
	if o == nil || IsNil(o.Windows) {
		return nil, false
	}
	return o.Windows, true
}

// HasWindows returns a boolean if a field has been set.
func (o *OsConfiguration) HasWindows() bool {
	if o != nil && !IsNil(o.Windows) {
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
	if o == nil || IsNil(o.RootPassword) {
		var ret string
		return ret
	}
	return *o.RootPassword
}

// GetRootPasswordOk returns a tuple with the RootPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsConfiguration) GetRootPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.RootPassword) {
		return nil, false
	}
	return o.RootPassword, true
}

// HasRootPassword returns a boolean if a field has been set.
func (o *OsConfiguration) HasRootPassword() bool {
	if o != nil && !IsNil(o.RootPassword) {
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
	if o == nil || IsNil(o.ManagementUiUrl) {
		var ret string
		return ret
	}
	return *o.ManagementUiUrl
}

// GetManagementUiUrlOk returns a tuple with the ManagementUiUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsConfiguration) GetManagementUiUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ManagementUiUrl) {
		return nil, false
	}
	return o.ManagementUiUrl, true
}

// HasManagementUiUrl returns a boolean if a field has been set.
func (o *OsConfiguration) HasManagementUiUrl() bool {
	if o != nil && !IsNil(o.ManagementUiUrl) {
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
	if o == nil || IsNil(o.ManagementAccessAllowedIps) {
		var ret []string
		return ret
	}
	return o.ManagementAccessAllowedIps
}

// GetManagementAccessAllowedIpsOk returns a tuple with the ManagementAccessAllowedIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsConfiguration) GetManagementAccessAllowedIpsOk() ([]string, bool) {
	if o == nil || IsNil(o.ManagementAccessAllowedIps) {
		return nil, false
	}
	return o.ManagementAccessAllowedIps, true
}

// HasManagementAccessAllowedIps returns a boolean if a field has been set.
func (o *OsConfiguration) HasManagementAccessAllowedIps() bool {
	if o != nil && !IsNil(o.ManagementAccessAllowedIps) {
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
	if o == nil || IsNil(o.InstallOsToRam) {
		var ret bool
		return ret
	}
	return *o.InstallOsToRam
}

// GetInstallOsToRamOk returns a tuple with the InstallOsToRam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsConfiguration) GetInstallOsToRamOk() (*bool, bool) {
	if o == nil || IsNil(o.InstallOsToRam) {
		return nil, false
	}
	return o.InstallOsToRam, true
}

// HasInstallOsToRam returns a boolean if a field has been set.
func (o *OsConfiguration) HasInstallOsToRam() bool {
	if o != nil && !IsNil(o.InstallOsToRam) {
		return true
	}

	return false
}

// SetInstallOsToRam gets a reference to the given bool and assigns it to the InstallOsToRam field.
func (o *OsConfiguration) SetInstallOsToRam(v bool) {
	o.InstallOsToRam = &v
}

// GetEsxi returns the Esxi field value if set, zero value otherwise.
func (o *OsConfiguration) GetEsxi() EsxiOsConfiguration {
	if o == nil || IsNil(o.Esxi) {
		var ret EsxiOsConfiguration
		return ret
	}
	return *o.Esxi
}

// GetEsxiOk returns a tuple with the Esxi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsConfiguration) GetEsxiOk() (*EsxiOsConfiguration, bool) {
	if o == nil || IsNil(o.Esxi) {
		return nil, false
	}
	return o.Esxi, true
}

// HasEsxi returns a boolean if a field has been set.
func (o *OsConfiguration) HasEsxi() bool {
	if o != nil && !IsNil(o.Esxi) {
		return true
	}

	return false
}

// SetEsxi gets a reference to the given EsxiOsConfiguration and assigns it to the Esxi field.
func (o *OsConfiguration) SetEsxi(v EsxiOsConfiguration) {
	o.Esxi = &v
}

// GetCloudInit returns the CloudInit field value if set, zero value otherwise.
func (o *OsConfiguration) GetCloudInit() OsConfigurationCloudInit {
	if o == nil || IsNil(o.CloudInit) {
		var ret OsConfigurationCloudInit
		return ret
	}
	return *o.CloudInit
}

// GetCloudInitOk returns a tuple with the CloudInit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsConfiguration) GetCloudInitOk() (*OsConfigurationCloudInit, bool) {
	if o == nil || IsNil(o.CloudInit) {
		return nil, false
	}
	return o.CloudInit, true
}

// HasCloudInit returns a boolean if a field has been set.
func (o *OsConfiguration) HasCloudInit() bool {
	if o != nil && !IsNil(o.CloudInit) {
		return true
	}

	return false
}

// SetCloudInit gets a reference to the given OsConfigurationCloudInit and assigns it to the CloudInit field.
func (o *OsConfiguration) SetCloudInit(v OsConfigurationCloudInit) {
	o.CloudInit = &v
}

func (o OsConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OsConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NetrisController) {
		toSerialize["netrisController"] = o.NetrisController
	}
	if !IsNil(o.NetrisSoftgate) {
		toSerialize["netrisSoftgate"] = o.NetrisSoftgate
	}
	if !IsNil(o.Windows) {
		toSerialize["windows"] = o.Windows
	}
	if !IsNil(o.RootPassword) {
		toSerialize["rootPassword"] = o.RootPassword
	}
	if !IsNil(o.ManagementUiUrl) {
		toSerialize["managementUiUrl"] = o.ManagementUiUrl
	}
	if !IsNil(o.ManagementAccessAllowedIps) {
		toSerialize["managementAccessAllowedIps"] = o.ManagementAccessAllowedIps
	}
	if !IsNil(o.InstallOsToRam) {
		toSerialize["installOsToRam"] = o.InstallOsToRam
	}
	if !IsNil(o.Esxi) {
		toSerialize["esxi"] = o.Esxi
	}
	if !IsNil(o.CloudInit) {
		toSerialize["cloudInit"] = o.CloudInit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OsConfiguration) UnmarshalJSON(data []byte) (err error) {
	varOsConfiguration := _OsConfiguration{}

	err = json.Unmarshal(data, &varOsConfiguration)

	if err != nil {
		return err
	}

	*o = OsConfiguration(varOsConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "netrisController")
		delete(additionalProperties, "netrisSoftgate")
		delete(additionalProperties, "windows")
		delete(additionalProperties, "rootPassword")
		delete(additionalProperties, "managementUiUrl")
		delete(additionalProperties, "managementAccessAllowedIps")
		delete(additionalProperties, "installOsToRam")
		delete(additionalProperties, "esxi")
		delete(additionalProperties, "cloudInit")
		o.AdditionalProperties = additionalProperties
	}

	return err
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
