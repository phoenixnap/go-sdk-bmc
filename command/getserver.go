package command

import (
	"net/http"

	"github.com/PNAP/bmc-api-sdk/client"
)

// GetServerCommand represents command that pulls details about specific server
type GetServerCommand struct {
	requester client.Requester
	serverID  string
}

// Execute pulls details about specific server
func (command *GetServerCommand) Execute() (*http.Response, error) {
	var req = command.requester
	return req.Get("servers/" + command.serverID)
}

// SetRequester sets requester to the command
func (command *GetServerCommand) SetRequester(requester client.Requester) {
	command.requester = requester
}

// SetServerID sets server id to the command
func (command *GetServerCommand) SetServerID(id string) {
	command.serverID = id
}

//NewGetServerCommand constructs new commmand of this type
func NewGetServerCommand(requester client.Requester, serverID string) *GetServerCommand {

	return &GetServerCommand{requester, serverID}
}
