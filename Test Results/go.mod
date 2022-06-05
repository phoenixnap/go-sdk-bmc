module tests

go 1.17

require (
	github.com/phoenixnap/go-sdk-bmc/auditapi v1.0.3
	github.com/phoenixnap/go-sdk-bmc/bmcapi v1.2.1
	github.com/phoenixnap/go-sdk-bmc/ipapi v1.1.1
	github.com/phoenixnap/go-sdk-bmc/networkapi v1.1.2
	github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi v1.1.2
	github.com/phoenixnap/go-sdk-bmc/tagapi v1.1.2
	github.com/stretchr/testify v1.4.0
)

replace github.com/phoenixnap/go-sdk-bmc/auditapi => ../auditapi/

replace github.com/phoenixnap/go-sdk-bmc/ipapi => ../ipapi/

replace github.com/phoenixnap/go-sdk-bmc/bmcapi => ../bmcapi/

replace github.com/phoenixnap/go-sdk-bmc/tagapi => ../tagapi/

replace github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi => ../ranchersolutionapi/

replace github.com/phoenixnap/go-sdk-bmc/networkapi => ../networkapi/
