package command

import (
	"net/http"

	"github.com/PNAP/bmc-api-sdk/client"
	"github.com/PNAP/bmc-api-sdk/dto"
)

// CreateServerCommand represents command for server provisioning
type CreateServerCommand struct {
	requester client.Requester
	server    dto.ProvisionedServer
}

// Execute provisions new server
func (command *CreateServerCommand) Execute() (*http.Response, error) {
	var req = command.requester
	return req.Post("servers", command.server.ToBytes())
}

// SetRequester sets requester to the command
func (command *CreateServerCommand) SetRequester(requester client.Requester) {
	command.requester = requester
}

// SetServer sets server details to the command
func (command *CreateServerCommand) SetServer(server dto.ProvisionedServer) {
	command.server = server
}
