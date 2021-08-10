package command

import (
	"net/http"

	"github.com/phoenixnap/go-sdk-bmc/client"
	"github.com/phoenixnap/go-sdk-bmc/dto"
)

// CreateServerCommand represents command for server provisioning. Use NewCreateServerCommand to initilize command properly.
type CreateServerCommand struct {
	requester client.Requester
	server    dto.ProvisionedServer
}

// Execute provisions new server
func (command *CreateServerCommand) Execute() (*http.Response, error) {
	var req = command.requester
	var apiPrefix = "bmc/v1/"
	return req.Post(apiPrefix + "servers", command.server.ToBytes())
}

// SetRequester sets requester to the command
func (command *CreateServerCommand) SetRequester(requester client.Requester) {
	command.requester = requester
}

// SetServer sets server details to the command
func (command *CreateServerCommand) SetServer(server dto.ProvisionedServer) {
	command.server = server
}

//NewCreateServerCommand constructs new commmand of this type
func NewCreateServerCommand(requester client.Requester, server dto.ProvisionedServer) *CreateServerCommand {

	return &CreateServerCommand{requester, server}
}
