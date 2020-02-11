package command

import (
	"net/http"
)

// Executor is the interface used to define contract for commands.
type Executor interface {
	Execute() (*http.Response, error)
}
