package command

import (
	"net/http"

	"github.com/phoenixnap/go-sdk-bmc/client"
	"github.com/phoenixnap/go-sdk-bmc/dto"
)

// ReserveCommand represents command that reserves the server
type ReserveCommand struct {
	requester client.Requester
	server    dto.ProvisionedServer
}

// Execute reserve on the server
func (command *ReserveCommand) Execute() (*http.Response, error) {
	var req = command.requester
	var apiPrefix = "bmc/v1/"
	return req.Post(apiPrefix+"servers/"+command.server.ID+"/actions/reserve", command.server.ToBytes())
}

// SetRequester sets requester to the command
func (command *ReserveCommand) SetRequester(requester client.Requester) {
	command.requester = requester
}

// SetServer sets server details to the command
func (command *ReserveCommand) SetServer(server dto.ProvisionedServer) {
	command.server = server
}

//NewReserveCommand constructs new commmand of this type
func NewReserveCommand(requester client.Requester, server dto.ProvisionedServer) *ReserveCommand {
	return &ReserveCommand{requester, server}
}
