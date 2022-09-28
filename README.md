<h1 align="center">
  <br>
  <a href="https://phoenixnap.com/bare-metal-cloud"><img src="https://user-images.githubusercontent.com/78744488/109779287-16da8600-7c06-11eb-81a1-97bf44983d33.png" alt="phoenixnap Bare Metal Cloud" width="300"></a>
  <br>
  Go SDK for Bare Metal Cloud
  <br>
</h1>

<p align="center">
This SDK allows you to provision and manage Bare Metal Cloud servers via API with Go.
</p>

<p align="center">
  <a href="https://phoenixnap.com/bare-metal-cloud">Bare Metal Cloud</a> •
  <a href="https://developers.phoenixnap.com/apis">API</a> •
  <a href="https://developers.phoenixnap.com/">Developers Portal</a> •
  <a href="http://phoenixnap.com/kb">Knowledge Base</a> •
  <a href="https://developers.phoenixnap.com/support">Support</a>
</p>

## Requirements

- [Bare Metal Cloud](https://bmc.phoenixnap.com) account
- [Go](https://golang.org/dl/)

## Creating a Bare Metal Cloud account

1. Go to the [Bare Metal Cloud signup page](https://support.phoenixnap.com/wap-jpost3/bmcSignup).
2. Follow the prompts to set up your account.
3. Use your credentials to [log in to Bare Metal Cloud portal](https://bmc.phoenixnap.com).

:arrow_forward: **Video tutorial:** [How to Create a Bare Metal Cloud Account in Minutes](https://www.youtube.com/watch?v=hPR60XWOSsQ)
<br>

:arrow_forward: **Video tutorial:** [How to Deploy a Bare Metal Server in a Minute](https://www.youtube.com/watch?v=BzBBwLxR80o)

## SDK Usage Example

Bare Metal Cloud Go SDK provides a library of resources and APIs that can be used to interact with the platform. Below is a sample code to create dependancy: 

```go
package main

import (
	"context"

	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"golang.org/x/oauth2/clientcredentials"
)

func main() {

	configuration := bmcapiclient.NewConfiguration()

	config := clientcredentials.Config{
		ClientID:     "<Client ID>",
		ClientSecret: "<Client Secret>",
		TokenURL:     "<Token URL>",
		Scopes:       []string{"bmc", "bmc.read"},
	}

	configuration.HTTPClient = config.Client(context.Background())

	api_client := bmcapiclient.NewAPIClient(configuration)

	resp, r, err := api_client.ServersApi.ServersGet(context.Background()).Execute()

```

:bulb: For each submodule, API-specific documentation is available as follows: 

- [Audit Logs API](auditapi/README.md): read audit log entries and track API calls and activities in the BMC Portal
- [Billing API](billingapi/README.md): automate infrastructure billing, reserve server instances, retrieve your servers' rated usage, and enable or disable auto-renewals.
- [BMC API](bmcapi/README.md): create, power on, power off, reset, reboot or shut down your servers. Deprovision your servers, get SSH keys, and more. 
- [Network API](networkapi/README.md): create, list, edit, and delete private networks. 
- [Network Storage API](networkstorageapi/README.md): Create, list, edit, and delete storage networks.
- [Rancher solutions API](ranchersolutionapi/README.md): deploy Kubernetes clusters faster using BMC integration with Rancher. 
- [Tag API](tagapi/README.md): assign tags to relevant resources in your BMC portal to group and categorize them. 
- [IP API](ipapi/README.md): request and delete IP Blocks.

## Bare Metal Cloud community

Become part of the Bare Metal Cloud community to get updates on new features, help us improve the platform, and engage with developers and other users.

- Follow [@phoenixNAP on Twitter](https://twitter.com/phoenixnap)
- Join the [official Slack channel](https://phoenixnap.slack.com)
- Sign up for our [Developers Monthly newsletter](https://phoenixnap.com/developers-monthly-newsletter)

### Resources

- [Product page](https://phoenixnap.com/bare-metal-cloud)
- [Instance pricing](https://phoenixnap.com/bare-metal-cloud/instances)
- [YouTube tutorials](https://www.youtube.com/watch?v=8TLsqgLDMN4&list=PLWcrQnFWd54WwkHM0oPpR1BrAhxlsy1Rc&ab_channel=PhoenixNAPGlobalITServices)
- [Developers Portal](https://developers.phoenixnap.com)
- [Knowledge Base](https://phoenixnap.com/kb)
- [Blog](https:/phoenixnap.com/blog)

### Documentation

- [API documentation](https://developers.phoenixnap.com/apis)

### Contact phoenixNAP

Get in touch with us if you have questions or need help with Bare Metal Cloud.

<p align="left">
  <a href="https://twitter.com/phoenixNAP">Twitter</a> •
  <a href="https://www.facebook.com/phoenixnap">Facebook</a> •
  <a href="https://www.linkedin.com/company/phoenix-nap">LinkedIn</a> •
  <a href="https://www.instagram.com/phoenixnap">Instagram</a> •
  <a href="https://www.youtube.com/user/PhoenixNAPdatacenter">YouTube</a> •
  <a href="https://developers.phoenixnap.com/support">Email</a> 
</p>

<p align="center">
  <br>
  <a href="https://phoenixnap.com/bare-metal-cloud"><img src="https://user-images.githubusercontent.com/81640346/115243282-0c773b80-a123-11eb-9de7-59e3934a5712.jpg" alt="phoenixnap Bare Metal Cloud"></a>
</p>
