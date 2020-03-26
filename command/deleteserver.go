package command

import (
	"net/http"

	"github.com/PNAP/bmc-api-sdk/client"
)

// DeleteServerCommand represents command for server deprovisioning
type DeleteServerCommand struct {
	requester client.Requester
	serverID  string
}

// Execute deprovisions the server
func (command *DeleteServerCommand) Execute() (*http.Response, error) {
	var req = command.requester
	return req.Delete("servers/" + command.serverID)
}

// SetRequester sets requester to the command
func (command *DeleteServerCommand) SetRequester(requester client.Requester) {
	command.requester = requester
}

// SetServerID sets server id to the command
func (command *DeleteServerCommand) SetServerID(id string) {
	command.serverID = id
}

//NewDeleteServerCommand constructs new commmand of this type
func NewDeleteServerCommand(requester client.Requester, serverID string) *DeleteServerCommand {

	return &DeleteServerCommand{requester, serverID}
}
