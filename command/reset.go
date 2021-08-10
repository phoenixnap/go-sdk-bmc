package command

import (
	"net/http"

	"github.com/PNAP/bmc-api-sdk/client"
	"github.com/PNAP/bmc-api-sdk/dto"
)

// ResetCommand represents command that resets the server
type ResetCommand struct {
	requester client.Requester
	server    dto.ProvisionedServer
}

// Execute reset on the server
func (command *ResetCommand) Execute() (*http.Response, error) {
	var req = command.requester
	var apiPrefix = "bmc/v1/"
	return req.Post(apiPrefix+"servers/"+command.server.ID+"/actions/reset", command.server.ToBytes())
}

// SetRequester sets requester to the command
func (command *ResetCommand) SetRequester(requester client.Requester) {
	command.requester = requester
}

// SetServer sets server details to the command
func (command *ResetCommand) SetServer(server dto.ProvisionedServer) {
	command.server = server
}

//NewResetCommand constructs new commmand of this type
func NewResetCommand(requester client.Requester, server dto.ProvisionedServer) *ResetCommand {

	return &ResetCommand{requester, server}
}
