module tests

go 1.18

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
	github.com/stretchr/testify v1.4.0
	golang.org/x/oauth2 v0.0.0-20210323180902-22b0adad7558
)

require (
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.23.0 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.2.2 // indirect
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
