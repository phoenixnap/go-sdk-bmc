module tests

go 1.23.0

toolchain go1.24.3

require (
	github.com/phoenixnap/go-sdk-bmc/auditapi v0.0.0
	github.com/phoenixnap/go-sdk-bmc/billingapi v0.0.0
	github.com/phoenixnap/go-sdk-bmc/bmcapi v0.0.0
	github.com/phoenixnap/go-sdk-bmc/invoicingapi v0.0.0
	github.com/phoenixnap/go-sdk-bmc/ipapi v0.0.0
	github.com/phoenixnap/go-sdk-bmc/locationapi v0.0.0
	github.com/phoenixnap/go-sdk-bmc/networkapi v0.0.0
	github.com/phoenixnap/go-sdk-bmc/networkstorageapi v0.0.0
	github.com/phoenixnap/go-sdk-bmc/paymentsapi v0.0.0
	github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi v0.0.0
	github.com/phoenixnap/go-sdk-bmc/tagapi v0.0.0
	github.com/stretchr/testify v1.10.0
	golang.org/x/oauth2 v0.27.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/phoenixnap/go-sdk-bmc/auditapi => ../auditapi/

replace github.com/phoenixnap/go-sdk-bmc/ipapi => ../ipapi/

replace github.com/phoenixnap/go-sdk-bmc/bmcapi => ../bmcapi/

replace github.com/phoenixnap/go-sdk-bmc/tagapi => ../tagapi/

replace github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi => ../ranchersolutionapi/

replace github.com/phoenixnap/go-sdk-bmc/networkapi => ../networkapi/

replace github.com/phoenixnap/go-sdk-bmc/billingapi => ../billingapi/

replace github.com/phoenixnap/go-sdk-bmc/networkstorageapi => ../networkstorageapi/

replace github.com/phoenixnap/go-sdk-bmc/locationapi => ../locationapi/

replace github.com/phoenixnap/go-sdk-bmc/invoicingapi => ../invoicingapi/

replace github.com/phoenixnap/go-sdk-bmc/paymentsapi => ../paymentsapi/
