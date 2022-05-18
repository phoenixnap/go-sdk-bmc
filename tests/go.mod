module tests

go 1.17

require (
	github.com/phoenixnap/go-sdk-bmc/tagapi v1.1.1
	github.com/stretchr/testify v1.4.0
)

replace github.com/phoenixnap/go-sdk-bmc/tagapi v1.1.1 => ../tagapi/
