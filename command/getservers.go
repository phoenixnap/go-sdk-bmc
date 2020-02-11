package command

import (
	"net/http"

	"github.com/PNAP/bmc-api-sdk/client"
)

// GetServersCommand represents command for server retrieval
type GetServersCommand struct {
	requester client.Requester
}

// Execute retrieves all servers
func (command *GetServersCommand) Execute() (*http.Response, error) {
	var req = command.requester
	return req.Get("servers")
}

// SetRequester sets requester to the command
func (command *GetServersCommand) SetRequester(requester client.Requester) {
	command.requester = requester
}
