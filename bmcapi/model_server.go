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
	"fmt"
	"time"
)

// checks if the Server type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Server{}

// Server Bare metal server.
type Server struct {
	// The unique identifier of the server.
	Id string `json:"id"`
	// The status of the server. Can have one of the following values: `creating` , `powered-on` , `powered-off` , `rebooting` , `resetting` , `deleting` , `reserved` , `error` or `reinstating`.
	Status string `json:"status"`
	// Hostname of server.
	Hostname string `json:"hostname"`
	// Description of server.
	Description *string `json:"description,omitempty"`
	// The server’s OS ID used when the server was created. Currently this field should be set to either `ubuntu/bionic`, `ubuntu/focal`, `ubuntu/jammy`, `ubuntu/jammy+pytorch`, `ubuntu/noble`, `centos/centos7`, `centos/centos8`, `windows/srv2019std`, `windows/srv2019dc`, `windows/srv2022std`, `windows/srv2022dc`, `windows/srv2025std`, `windows/srv2025dc`, `esxi/esxi70`, `esxi/esxi80`, `almalinux/almalinux8`, `rockylinux/rockylinux8`, `almalinux/almalinux9`, `rockylinux/rockylinux9`, `virtuozzo/virtuozzo7`, `oraclelinux/oraclelinux9`, `debian/bullseye`, `debian/bookworm`, `proxmox/bullseye`, `proxmox/proxmox8`, `netris/controller`, `netris/softgate_1g`, `netris/softgate_10g` or `netris/softgate_25g`.
	Os *string `json:"os,omitempty"`
	// Server type ID. Cannot be changed once a server is created. Currently this field should be set to either `s0.d1.small`, `s0.d1.medium`, `s1.c1.small`, `s1.c1.medium`, `s1.c2.medium`, `s1.c2.large`, `s1.e1.small`, `s1.e1.medium`, `s1.e1.large`, `s2.c1.small`, `s2.c1.medium`, `s2.c1.large`, `s2.c2.small`, `s2.c2.medium`, `s2.c2.large`, `d1.c4.small`, `d1.c4.medium`, `d1.c4.large`, `d1.m4.medium`, `d2.c1.medium`, `d2.c2.medium`, `d2.c3.medium`, `d2.c4.medium`, `d2.c5.medium`, `d2.c1.large`, `d2.c2.large`, `d2.c3.large`, `d2.c4.large`, `d2.c5.large`, `d2.m1.xlarge`, `d2.m2.xxlarge`, `d2.m3.xlarge`, `d2.m4.xlarge`, `d2.m5.xlarge`, `d2.c4.db1.pliops1`, `d3.m4.xlarge`, `d3.m5.xlarge`, `d3.m6.xlarge`, `a1.c5.large`, `a1.c5.xlarge`, `d3.s5.xlarge`, `d3.m4.xxlarge`, `d3.m5.xxlarge`,  `d3.m6.xxlarge`, `s3.c3.medium`, `s3.c3.large`, `d3.c4.medium`, `d3.c5.medium`, `d3.c6.medium`, `d3.c1.large`, `d3.c2.large`, `d3.c3.large`, `d3.m1.xlarge`, `d3.m2.xlarge`, `d3.m3.xlarge`, `d3.g2.c1.xlarge`, `d3.g2.c2.xlarge`, `d3.g2.c3.xlarge`, s4.x6.c6.large or s4.x6.m6.xlarge.
	Type string `json:"type"`
	// Server location ID. Cannot be changed once a server is created. Currently this field should be set to `PHX`, `ASH`, `SGP`, `NLD`, `CHI`, `SEA` or `AUS`.
	Location string `json:"location"`
	// A description of the machine CPU.
	Cpu string `json:"cpu"`
	// The number of CPUs available in the system.
	CpuCount int32 `json:"cpuCount"`
	// The number of physical cores present on each CPU.
	CoresPerCpu int32 `json:"coresPerCpu"`
	// The CPU frequency in GHz.
	CpuFrequency float32 `json:"cpuFrequency"`
	// A description of the machine RAM.
	Ram string `json:"ram"`
	// A description of the machine storage.
	Storage string `json:"storage"`
	// Private IP addresses assigned to server.
	PrivateIpAddresses []string `json:"privateIpAddresses"`
	// Public IP addresses assigned to server.
	PublicIpAddresses []string `json:"publicIpAddresses,omitempty"`
	// The reservation reference id if any.
	ReservationId *string `json:"reservationId,omitempty"`
	// The pricing model this server is being billed. Currently this field should be set to `HOURLY`, `ONE_MONTH_RESERVATION`, `TWELVE_MONTHS_RESERVATION`, `TWENTY_FOUR_MONTHS_RESERVATION` or `THIRTY_SIX_MONTHS_RESERVATION`.
	PricingModel string `json:"pricingModel"`
	// Auto-generated password set for user `Admin` on Windows server, user `root` on ESXi servers, user `root` on Proxmox server and user `netris` on Netris servers.<br> The password is not stored and therefore will only be returned in response to provisioning a server. Copy and save it for future reference.
	Password *string `json:"password,omitempty"`
	// The type of network configuration for this server. Currently this field should be set to `PUBLIC_AND_PRIVATE`, `PRIVATE_ONLY`, `PUBLIC_ONLY` or `NONE`.
	NetworkType *string `json:"networkType,omitempty"`
	// The cluster reference id if any.
	ClusterId *string `json:"clusterId,omitempty"`
	// The tags assigned if any.
	Tags []TagAssignment `json:"tags,omitempty"`
	// Date and time when server was provisioned.
	ProvisionedOn        *time.Time           `json:"provisionedOn,omitempty"`
	OsConfiguration      *OsConfiguration     `json:"osConfiguration,omitempty"`
	NetworkConfiguration NetworkConfiguration `json:"networkConfiguration"`
	StorageConfiguration StorageConfiguration `json:"storageConfiguration"`
	GpuConfiguration     *GpuConfiguration    `json:"gpuConfiguration,omitempty"`
	// Unique identifier of the server to which the reservation has been transferred.
	SupersededBy *string `json:"supersededBy,omitempty"`
	// Unique identifier of the server from which the reservation has been transferred.
	Supersedes           *string `json:"supersedes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Server Server

// NewServer instantiates a new Server object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServer(id string, status string, hostname string, type_ string, location string, cpu string, cpuCount int32, coresPerCpu int32, cpuFrequency float32, ram string, storage string, privateIpAddresses []string, pricingModel string, networkConfiguration NetworkConfiguration, storageConfiguration StorageConfiguration) *Server {
	this := Server{}
	this.Id = id
	this.Status = status
	this.Hostname = hostname
	this.Type = type_
	this.Location = location
	this.Cpu = cpu
	this.CpuCount = cpuCount
	this.CoresPerCpu = coresPerCpu
	this.CpuFrequency = cpuFrequency
	this.Ram = ram
	this.Storage = storage
	this.PrivateIpAddresses = privateIpAddresses
	this.PricingModel = pricingModel
	var networkType string = "PUBLIC_AND_PRIVATE"
	this.NetworkType = &networkType
	this.NetworkConfiguration = networkConfiguration
	this.StorageConfiguration = storageConfiguration
	return &this
}

// NewServerWithDefaults instantiates a new Server object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerWithDefaults() *Server {
	this := Server{}
	var pricingModel string = "HOURLY"
	this.PricingModel = pricingModel
	var networkType string = "PUBLIC_AND_PRIVATE"
	this.NetworkType = &networkType
	return &this
}

// GetId returns the Id field value
func (o *Server) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Server) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Server) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *Server) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Server) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Server) SetStatus(v string) {
	o.Status = v
}

// GetHostname returns the Hostname field value
func (o *Server) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *Server) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *Server) SetHostname(v string) {
	o.Hostname = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Server) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Server) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Server) SetDescription(v string) {
	o.Description = &v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *Server) GetOs() string {
	if o == nil || IsNil(o.Os) {
		var ret string
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetOsOk() (*string, bool) {
	if o == nil || IsNil(o.Os) {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *Server) HasOs() bool {
	if o != nil && !IsNil(o.Os) {
		return true
	}

	return false
}

// SetOs gets a reference to the given string and assigns it to the Os field.
func (o *Server) SetOs(v string) {
	o.Os = &v
}

// GetType returns the Type field value
func (o *Server) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Server) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Server) SetType(v string) {
	o.Type = v
}

// GetLocation returns the Location field value
func (o *Server) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *Server) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *Server) SetLocation(v string) {
	o.Location = v
}

// GetCpu returns the Cpu field value
func (o *Server) GetCpu() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value
// and a boolean to check if the value has been set.
func (o *Server) GetCpuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cpu, true
}

// SetCpu sets field value
func (o *Server) SetCpu(v string) {
	o.Cpu = v
}

// GetCpuCount returns the CpuCount field value
func (o *Server) GetCpuCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CpuCount
}

// GetCpuCountOk returns a tuple with the CpuCount field value
// and a boolean to check if the value has been set.
func (o *Server) GetCpuCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpuCount, true
}

// SetCpuCount sets field value
func (o *Server) SetCpuCount(v int32) {
	o.CpuCount = v
}

// GetCoresPerCpu returns the CoresPerCpu field value
func (o *Server) GetCoresPerCpu() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CoresPerCpu
}

// GetCoresPerCpuOk returns a tuple with the CoresPerCpu field value
// and a boolean to check if the value has been set.
func (o *Server) GetCoresPerCpuOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CoresPerCpu, true
}

// SetCoresPerCpu sets field value
func (o *Server) SetCoresPerCpu(v int32) {
	o.CoresPerCpu = v
}

// GetCpuFrequency returns the CpuFrequency field value
func (o *Server) GetCpuFrequency() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CpuFrequency
}

// GetCpuFrequencyOk returns a tuple with the CpuFrequency field value
// and a boolean to check if the value has been set.
func (o *Server) GetCpuFrequencyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpuFrequency, true
}

// SetCpuFrequency sets field value
func (o *Server) SetCpuFrequency(v float32) {
	o.CpuFrequency = v
}

// GetRam returns the Ram field value
func (o *Server) GetRam() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ram
}

// GetRamOk returns a tuple with the Ram field value
// and a boolean to check if the value has been set.
func (o *Server) GetRamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ram, true
}

// SetRam sets field value
func (o *Server) SetRam(v string) {
	o.Ram = v
}

// GetStorage returns the Storage field value
func (o *Server) GetStorage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *Server) GetStorageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *Server) SetStorage(v string) {
	o.Storage = v
}

// GetPrivateIpAddresses returns the PrivateIpAddresses field value
func (o *Server) GetPrivateIpAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PrivateIpAddresses
}

// GetPrivateIpAddressesOk returns a tuple with the PrivateIpAddresses field value
// and a boolean to check if the value has been set.
func (o *Server) GetPrivateIpAddressesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrivateIpAddresses, true
}

// SetPrivateIpAddresses sets field value
func (o *Server) SetPrivateIpAddresses(v []string) {
	o.PrivateIpAddresses = v
}

// GetPublicIpAddresses returns the PublicIpAddresses field value if set, zero value otherwise.
func (o *Server) GetPublicIpAddresses() []string {
	if o == nil || IsNil(o.PublicIpAddresses) {
		var ret []string
		return ret
	}
	return o.PublicIpAddresses
}

// GetPublicIpAddressesOk returns a tuple with the PublicIpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetPublicIpAddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.PublicIpAddresses) {
		return nil, false
	}
	return o.PublicIpAddresses, true
}

// HasPublicIpAddresses returns a boolean if a field has been set.
func (o *Server) HasPublicIpAddresses() bool {
	if o != nil && !IsNil(o.PublicIpAddresses) {
		return true
	}

	return false
}

// SetPublicIpAddresses gets a reference to the given []string and assigns it to the PublicIpAddresses field.
func (o *Server) SetPublicIpAddresses(v []string) {
	o.PublicIpAddresses = v
}

// GetReservationId returns the ReservationId field value if set, zero value otherwise.
func (o *Server) GetReservationId() string {
	if o == nil || IsNil(o.ReservationId) {
		var ret string
		return ret
	}
	return *o.ReservationId
}

// GetReservationIdOk returns a tuple with the ReservationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetReservationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReservationId) {
		return nil, false
	}
	return o.ReservationId, true
}

// HasReservationId returns a boolean if a field has been set.
func (o *Server) HasReservationId() bool {
	if o != nil && !IsNil(o.ReservationId) {
		return true
	}

	return false
}

// SetReservationId gets a reference to the given string and assigns it to the ReservationId field.
func (o *Server) SetReservationId(v string) {
	o.ReservationId = &v
}

// GetPricingModel returns the PricingModel field value
func (o *Server) GetPricingModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PricingModel
}

// GetPricingModelOk returns a tuple with the PricingModel field value
// and a boolean to check if the value has been set.
func (o *Server) GetPricingModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PricingModel, true
}

// SetPricingModel sets field value
func (o *Server) SetPricingModel(v string) {
	o.PricingModel = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *Server) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *Server) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *Server) SetPassword(v string) {
	o.Password = &v
}

// GetNetworkType returns the NetworkType field value if set, zero value otherwise.
func (o *Server) GetNetworkType() string {
	if o == nil || IsNil(o.NetworkType) {
		var ret string
		return ret
	}
	return *o.NetworkType
}

// GetNetworkTypeOk returns a tuple with the NetworkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetNetworkTypeOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkType) {
		return nil, false
	}
	return o.NetworkType, true
}

// HasNetworkType returns a boolean if a field has been set.
func (o *Server) HasNetworkType() bool {
	if o != nil && !IsNil(o.NetworkType) {
		return true
	}

	return false
}

// SetNetworkType gets a reference to the given string and assigns it to the NetworkType field.
func (o *Server) SetNetworkType(v string) {
	o.NetworkType = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *Server) GetClusterId() string {
	if o == nil || IsNil(o.ClusterId) {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetClusterIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterId) {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *Server) HasClusterId() bool {
	if o != nil && !IsNil(o.ClusterId) {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *Server) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Server) GetTags() []TagAssignment {
	if o == nil || IsNil(o.Tags) {
		var ret []TagAssignment
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetTagsOk() ([]TagAssignment, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Server) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagAssignment and assigns it to the Tags field.
func (o *Server) SetTags(v []TagAssignment) {
	o.Tags = v
}

// GetProvisionedOn returns the ProvisionedOn field value if set, zero value otherwise.
func (o *Server) GetProvisionedOn() time.Time {
	if o == nil || IsNil(o.ProvisionedOn) {
		var ret time.Time
		return ret
	}
	return *o.ProvisionedOn
}

// GetProvisionedOnOk returns a tuple with the ProvisionedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetProvisionedOnOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ProvisionedOn) {
		return nil, false
	}
	return o.ProvisionedOn, true
}

// HasProvisionedOn returns a boolean if a field has been set.
func (o *Server) HasProvisionedOn() bool {
	if o != nil && !IsNil(o.ProvisionedOn) {
		return true
	}

	return false
}

// SetProvisionedOn gets a reference to the given time.Time and assigns it to the ProvisionedOn field.
func (o *Server) SetProvisionedOn(v time.Time) {
	o.ProvisionedOn = &v
}

// GetOsConfiguration returns the OsConfiguration field value if set, zero value otherwise.
func (o *Server) GetOsConfiguration() OsConfiguration {
	if o == nil || IsNil(o.OsConfiguration) {
		var ret OsConfiguration
		return ret
	}
	return *o.OsConfiguration
}

// GetOsConfigurationOk returns a tuple with the OsConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetOsConfigurationOk() (*OsConfiguration, bool) {
	if o == nil || IsNil(o.OsConfiguration) {
		return nil, false
	}
	return o.OsConfiguration, true
}

// HasOsConfiguration returns a boolean if a field has been set.
func (o *Server) HasOsConfiguration() bool {
	if o != nil && !IsNil(o.OsConfiguration) {
		return true
	}

	return false
}

// SetOsConfiguration gets a reference to the given OsConfiguration and assigns it to the OsConfiguration field.
func (o *Server) SetOsConfiguration(v OsConfiguration) {
	o.OsConfiguration = &v
}

// GetNetworkConfiguration returns the NetworkConfiguration field value
func (o *Server) GetNetworkConfiguration() NetworkConfiguration {
	if o == nil {
		var ret NetworkConfiguration
		return ret
	}

	return o.NetworkConfiguration
}

// GetNetworkConfigurationOk returns a tuple with the NetworkConfiguration field value
// and a boolean to check if the value has been set.
func (o *Server) GetNetworkConfigurationOk() (*NetworkConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkConfiguration, true
}

// SetNetworkConfiguration sets field value
func (o *Server) SetNetworkConfiguration(v NetworkConfiguration) {
	o.NetworkConfiguration = v
}

// GetStorageConfiguration returns the StorageConfiguration field value
func (o *Server) GetStorageConfiguration() StorageConfiguration {
	if o == nil {
		var ret StorageConfiguration
		return ret
	}

	return o.StorageConfiguration
}

// GetStorageConfigurationOk returns a tuple with the StorageConfiguration field value
// and a boolean to check if the value has been set.
func (o *Server) GetStorageConfigurationOk() (*StorageConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageConfiguration, true
}

// SetStorageConfiguration sets field value
func (o *Server) SetStorageConfiguration(v StorageConfiguration) {
	o.StorageConfiguration = v
}

// GetGpuConfiguration returns the GpuConfiguration field value if set, zero value otherwise.
func (o *Server) GetGpuConfiguration() GpuConfiguration {
	if o == nil || IsNil(o.GpuConfiguration) {
		var ret GpuConfiguration
		return ret
	}
	return *o.GpuConfiguration
}

// GetGpuConfigurationOk returns a tuple with the GpuConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetGpuConfigurationOk() (*GpuConfiguration, bool) {
	if o == nil || IsNil(o.GpuConfiguration) {
		return nil, false
	}
	return o.GpuConfiguration, true
}

// HasGpuConfiguration returns a boolean if a field has been set.
func (o *Server) HasGpuConfiguration() bool {
	if o != nil && !IsNil(o.GpuConfiguration) {
		return true
	}

	return false
}

// SetGpuConfiguration gets a reference to the given GpuConfiguration and assigns it to the GpuConfiguration field.
func (o *Server) SetGpuConfiguration(v GpuConfiguration) {
	o.GpuConfiguration = &v
}

// GetSupersededBy returns the SupersededBy field value if set, zero value otherwise.
func (o *Server) GetSupersededBy() string {
	if o == nil || IsNil(o.SupersededBy) {
		var ret string
		return ret
	}
	return *o.SupersededBy
}

// GetSupersededByOk returns a tuple with the SupersededBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetSupersededByOk() (*string, bool) {
	if o == nil || IsNil(o.SupersededBy) {
		return nil, false
	}
	return o.SupersededBy, true
}

// HasSupersededBy returns a boolean if a field has been set.
func (o *Server) HasSupersededBy() bool {
	if o != nil && !IsNil(o.SupersededBy) {
		return true
	}

	return false
}

// SetSupersededBy gets a reference to the given string and assigns it to the SupersededBy field.
func (o *Server) SetSupersededBy(v string) {
	o.SupersededBy = &v
}

// GetSupersedes returns the Supersedes field value if set, zero value otherwise.
func (o *Server) GetSupersedes() string {
	if o == nil || IsNil(o.Supersedes) {
		var ret string
		return ret
	}
	return *o.Supersedes
}

// GetSupersedesOk returns a tuple with the Supersedes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetSupersedesOk() (*string, bool) {
	if o == nil || IsNil(o.Supersedes) {
		return nil, false
	}
	return o.Supersedes, true
}

// HasSupersedes returns a boolean if a field has been set.
func (o *Server) HasSupersedes() bool {
	if o != nil && !IsNil(o.Supersedes) {
		return true
	}

	return false
}

// SetSupersedes gets a reference to the given string and assigns it to the Supersedes field.
func (o *Server) SetSupersedes(v string) {
	o.Supersedes = &v
}

func (o Server) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Server) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["status"] = o.Status
	toSerialize["hostname"] = o.Hostname
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Os) {
		toSerialize["os"] = o.Os
	}
	toSerialize["type"] = o.Type
	toSerialize["location"] = o.Location
	toSerialize["cpu"] = o.Cpu
	toSerialize["cpuCount"] = o.CpuCount
	toSerialize["coresPerCpu"] = o.CoresPerCpu
	toSerialize["cpuFrequency"] = o.CpuFrequency
	toSerialize["ram"] = o.Ram
	toSerialize["storage"] = o.Storage
	toSerialize["privateIpAddresses"] = o.PrivateIpAddresses
	if !IsNil(o.PublicIpAddresses) {
		toSerialize["publicIpAddresses"] = o.PublicIpAddresses
	}
	if !IsNil(o.ReservationId) {
		toSerialize["reservationId"] = o.ReservationId
	}
	toSerialize["pricingModel"] = o.PricingModel
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.NetworkType) {
		toSerialize["networkType"] = o.NetworkType
	}
	if !IsNil(o.ClusterId) {
		toSerialize["clusterId"] = o.ClusterId
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.ProvisionedOn) {
		toSerialize["provisionedOn"] = o.ProvisionedOn
	}
	if !IsNil(o.OsConfiguration) {
		toSerialize["osConfiguration"] = o.OsConfiguration
	}
	toSerialize["networkConfiguration"] = o.NetworkConfiguration
	toSerialize["storageConfiguration"] = o.StorageConfiguration
	if !IsNil(o.GpuConfiguration) {
		toSerialize["gpuConfiguration"] = o.GpuConfiguration
	}
	if !IsNil(o.SupersededBy) {
		toSerialize["supersededBy"] = o.SupersededBy
	}
	if !IsNil(o.Supersedes) {
		toSerialize["supersedes"] = o.Supersedes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Server) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"status",
		"hostname",
		"type",
		"location",
		"cpu",
		"cpuCount",
		"coresPerCpu",
		"cpuFrequency",
		"ram",
		"storage",
		"privateIpAddresses",
		"pricingModel",
		"networkConfiguration",
		"storageConfiguration",
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

	varServer := _Server{}

	err = json.Unmarshal(data, &varServer)

	if err != nil {
		return err
	}

	*o = Server(varServer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "hostname")
		delete(additionalProperties, "description")
		delete(additionalProperties, "os")
		delete(additionalProperties, "type")
		delete(additionalProperties, "location")
		delete(additionalProperties, "cpu")
		delete(additionalProperties, "cpuCount")
		delete(additionalProperties, "coresPerCpu")
		delete(additionalProperties, "cpuFrequency")
		delete(additionalProperties, "ram")
		delete(additionalProperties, "storage")
		delete(additionalProperties, "privateIpAddresses")
		delete(additionalProperties, "publicIpAddresses")
		delete(additionalProperties, "reservationId")
		delete(additionalProperties, "pricingModel")
		delete(additionalProperties, "password")
		delete(additionalProperties, "networkType")
		delete(additionalProperties, "clusterId")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "provisionedOn")
		delete(additionalProperties, "osConfiguration")
		delete(additionalProperties, "networkConfiguration")
		delete(additionalProperties, "storageConfiguration")
		delete(additionalProperties, "gpuConfiguration")
		delete(additionalProperties, "supersededBy")
		delete(additionalProperties, "supersedes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServer struct {
	value *Server
	isSet bool
}

func (v NullableServer) Get() *Server {
	return v.value
}

func (v *NullableServer) Set(val *Server) {
	v.value = val
	v.isSet = true
}

func (v NullableServer) IsSet() bool {
	return v.isSet
}

func (v *NullableServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServer(val *Server) *NullableServer {
	return &NullableServer{value: val, isSet: true}
}

func (v NullableServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
