package command

import (
	"net/http"

	"github.com/phoenixnap/go-sdk-bmc/client"
)

// GetServersCommand represents command for server retrieval
type GetServersCommand struct {
	requester client.Requester
}

// Execute retrieves all servers
func (command *GetServersCommand) Execute() (*http.Response, error) {
	var req = command.requester
	var apiPrefix = "bmc/v1/"
	return req.Get(apiPrefix + "servers")
}

// SetRequester sets requester to the command
func (command *GetServersCommand) SetRequester(requester client.Requester) {
	command.requester = requester
}

//NewGetServersCommand constructs new commmand of this type
func NewGetServersCommand(requester client.Requester) *GetServersCommand {

	return &GetServersCommand{requester}
}
